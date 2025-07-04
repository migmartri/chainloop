//
// Copyright 2023 The Chainloop Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	pb "github.com/chainloop-dev/chainloop/app/controlplane/api/controlplane/v1"
	conf "github.com/chainloop-dev/chainloop/app/controlplane/internal/conf/controlplane/config/v1"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/biz"
	"github.com/chainloop-dev/chainloop/internal/oauth"
	casJWT "github.com/chainloop-dev/chainloop/internal/robotaccount/cas"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
)

const (
	getDownloadURLOperation = "/controlplane.v1.CASRedirectService/DownloadRedirect"

	// bearerFormat authorization token format
	bearerFormat string = "Bearer %s"

	// authorizationKey holds the key used to store the JWT Token in the request tokenHeader.
	authorizationKey string = "Authorization"
)

type CASRedirectService struct {
	pb.UnimplementedCASRedirectServiceServer
	*service

	casMappingUC    *biz.CASMappingUseCase
	casCredsUseCase *biz.CASCredentialsUseCase
	casServerConf   *conf.Bootstrap_CASServer
}

func NewCASRedirectService(casmUC *biz.CASMappingUseCase, casCredsUC *biz.CASCredentialsUseCase, conf *conf.Bootstrap_CASServer, opts ...NewOpt) (*CASRedirectService, error) {
	if conf == nil || conf.GetDownloadUrl() == "" {
		return nil, errors.New("CASServer.downloadURL configuration is missing")
	}

	return &CASRedirectService{
		service:         newService(opts...),
		casMappingUC:    casmUC,
		casCredsUseCase: casCredsUC,
		casServerConf:   conf,
	}, nil
}

// GetDownloadURL returns a signed URL to download the artifact from the CAS backend
// The URL includes a JWT token that is used to authenticate the request, this token has all the information required to validate the request
// The result would look like "https://cas.chainloop.dev/download/sha256:[DIGEST]?t=tokenJWT
func (s *CASRedirectService) GetDownloadURL(ctx context.Context, req *pb.GetDownloadURLRequest) (*pb.GetDownloadURLResponse, error) {
	currentUser, currentAPIToken, err := requireCurrentUserOrAPIToken(ctx)
	if err != nil {
		return nil, err
	}

	currentOrg, err := requireCurrentOrg(ctx)
	if err != nil {
		return nil, err
	}

	var mapping *biz.CASMapping
	if currentUser != nil {
		mapping, err = s.casMappingUC.FindCASMappingForDownloadByUser(ctx, req.Digest, currentUser.ID)
	} else if currentAPIToken != nil {
		var orgID uuid.UUID
		orgID, err = uuid.Parse(currentOrg.ID)
		if err != nil {
			return nil, handleUseCaseErr(err, s.log)
		}
		mapping, err = s.casMappingUC.FindCASMappingForDownloadByOrg(ctx, req.Digest, []uuid.UUID{orgID}, nil)
	}

	if err != nil {
		// We don't want to leak the fact that the asset exists but the user does not have permissions
		// that's why we return a generic 404 in unauthorized scenarios too
		if biz.IsNotFound(err) {
			return nil, kerrors.NotFound("not found", "artifact not found")
		} else if biz.IsErrValidation(err) {
			return nil, kerrors.BadRequest("invalid", err.Error())
		}

		return nil, handleUseCaseErr(err, s.log)
	}

	backend := mapping.CASBackend

	// inline backends don't have a download URL
	if backend.Inline {
		return nil, kerrors.NotFound("not found", "CAS backend is inline")
	}

	// Create an URL to download the artifact from the CAS backend
	downloadBase, err := url.Parse(s.casServerConf.GetDownloadUrl())
	if err != nil {
		return nil, handleUseCaseErr(err, s.log)
	}

	// 1 - append the digest /download/[digest]
	downloadURL := downloadBase.JoinPath(req.Digest)

	// 2- add authentication token to the query params ?t=[token]
	if backend.SecretName != "" {
		ref := &biz.CASCredsOpts{BackendType: string(backend.Provider), SecretPath: backend.SecretName, Role: casJWT.Downloader, MaxBytes: backend.Limits.MaxBytes}
		t, err := s.casCredsUseCase.GenerateTemporaryCredentials(ref)
		if err != nil {
			return nil, handleUseCaseErr(err, s.log)
		}

		q := downloadURL.Query()
		q.Set("t", t)
		downloadURL.RawQuery = q.Encode()
	}

	return &pb.GetDownloadURLResponse{Result: &pb.GetDownloadURLResponse_Result{Url: downloadURL.String()}}, nil
}

// Custom HTTP handler that handles redirect to the CAS download URL
// This handler does
// 1 - Force the user to do an oauth login dance
// 2 - calls GetDownloadURL to generate a proper URL
// 3 - redirects to the generated URL
// NOTE: This code is similar to the one that gets autogenerated if we happened to use proto.http but we can't use it because we need custom redirection
func (s *CASRedirectService) HTTPDownload(ctx khttp.Context) error {
	if redirected := authenticateUser(ctx.Response(), ctx.Request()); redirected {
		return nil
	}

	// Arguments marshalling
	var in pb.GetDownloadURLRequest
	if err := ctx.BindVars(&in); err != nil {
		return err
	}

	// Run the downloadURLCode within the middleware chain
	// Set identifier so it can be skipped, filtered, selected during the middleware chain
	khttp.SetOperation(ctx, getDownloadURLOperation)
	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetDownloadURL(ctx, req.(*pb.GetDownloadURLRequest))
	})

	out, err := h(ctx, &in)
	if err != nil {
		return err
	}

	urlResponse := out.(*pb.GetDownloadURLResponse)
	if urlResponse.GetResult().GetUrl() == "" {
		return kerrors.BadRequest("invalid URL", "the URL returned by the server is empty")
	}

	// if the client is not a browser (i.e curl), we perform a redirect
	acceptHeader := ctx.Request().Header.Get("Accept")
	if !strings.Contains(acceptHeader, "text/http") {
		ctx.Response().Header().Set("Location", urlResponse.Result.Url)
		ctx.Response().WriteHeader(http.StatusFound)
		return nil
	}

	// perform redirect
	// In order ton show a redirect message in the page before the redirection,
	// we need to use a Refresh Header instead of a vanilla redirect (via location change)
	// We want to show a message because in some cases the download will be shown as downloadable and hence
	// the UX for the user is like the browser got stuck
	ctx.Response().Header().Set("Refresh", "1;url="+urlResponse.Result.Url)
	ctx.Response().WriteHeader(http.StatusFound)
	fmt.Fprintln(ctx.Response(), "Your download will begin shortly...")

	return nil
}

// authenticateUser checks if the user is authenticated, if not it will redirect to the login page
// TODO: move to middleware
func authenticateUser(w http.ResponseWriter, r *http.Request) (redirected bool) {
	// Check if the token is already in the header
	if r.Header.Get(authorizationKey) != "" {
		return
	}

	// Check if it's in the query params and add it to the header
	// This scenario happens when the user is redirected from the oauth callback
	// so the next middleware will pick it up
	if userToken := r.URL.Query().Get("t"); userToken != "" {
		// Set the token in the request, this is required to call GetDownloadURL
		r.Header.Set(authorizationKey, fmt.Sprintf(bearerFormat, userToken))
		return
	}

	// redirect to login
	loginURL, _ := url.Parse(AuthLoginPath)
	// Append local callback URL
	q := loginURL.Query()
	q.Set(oauth.QueryParamCallback, r.URL.String())
	loginURL.RawQuery = q.Encode()

	// Get the current user auth token and set it in the request context
	http.Redirect(w, r, loginURL.String(), http.StatusTemporaryRedirect)
	return true
}
