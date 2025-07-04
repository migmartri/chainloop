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
	"time"

	pb "github.com/chainloop-dev/chainloop/app/controlplane/api/controlplane/v1"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type APITokenService struct {
	pb.UnimplementedAPITokenServiceServer
	*service

	APITokenUseCase *biz.APITokenUseCase
}

func NewAPITokenService(uc *biz.APITokenUseCase, opts ...NewOpt) *APITokenService {
	return &APITokenService{
		service:         newService(opts...),
		APITokenUseCase: uc,
	}
}

func (s *APITokenService) Create(ctx context.Context, req *pb.APITokenServiceCreateRequest) (*pb.APITokenServiceCreateResponse, error) {
	currentOrg, err := requireCurrentOrg(ctx)
	if err != nil {
		return nil, err
	}

	var expiresIn *time.Duration
	if req.ExpiresIn != nil {
		expiresIn = new(time.Duration)
		*expiresIn = req.ExpiresIn.AsDuration()
	}

	token, err := s.APITokenUseCase.Create(ctx, req.Name, req.Description, expiresIn, currentOrg.ID)
	if err != nil {
		return nil, handleUseCaseErr(err, s.log)
	}

	return &pb.APITokenServiceCreateResponse{
		Result: &pb.APITokenServiceCreateResponse_APITokenFull{
			Item: apiTokenBizToPb(token),
			Jwt:  token.JWT,
		},
	}, nil
}

func (s *APITokenService) List(ctx context.Context, req *pb.APITokenServiceListRequest) (*pb.APITokenServiceListResponse, error) {
	currentOrg, err := requireCurrentOrg(ctx)
	if err != nil {
		return nil, err
	}

	// Only expose system tokens
	tokens, err := s.APITokenUseCase.List(ctx, currentOrg.ID, req.IncludeRevoked, biz.APITokenShowOnlySystemTokens(true))
	if err != nil {
		return nil, handleUseCaseErr(err, s.log)
	}

	result := make([]*pb.APITokenItem, 0, len(tokens))
	for _, p := range tokens {
		result = append(result, apiTokenBizToPb(p))
	}

	return &pb.APITokenServiceListResponse{Result: result}, nil
}

func (s *APITokenService) Revoke(ctx context.Context, req *pb.APITokenServiceRevokeRequest) (*pb.APITokenServiceRevokeResponse, error) {
	currentOrg, err := requireCurrentOrg(ctx)
	if err != nil {
		return nil, err
	}

	t, err := s.APITokenUseCase.FindByNameInOrg(ctx, currentOrg.ID, req.Name)
	if err != nil {
		return nil, handleUseCaseErr(err, s.log)
	}

	if err := s.APITokenUseCase.Revoke(ctx, currentOrg.ID, t.ID.String()); err != nil {
		return nil, handleUseCaseErr(err, s.log)
	}

	return &pb.APITokenServiceRevokeResponse{}, nil
}

func apiTokenBizToPb(in *biz.APIToken) *pb.APITokenItem {
	res := &pb.APITokenItem{
		Id:               in.ID.String(),
		Name:             in.Name,
		OrganizationId:   in.OrganizationID.String(),
		OrganizationName: in.OrganizationName,
		Description:      in.Description,
		CreatedAt:        timestamppb.New(*in.CreatedAt),
	}

	if in.ExpiresAt != nil {
		res.ExpiresAt = timestamppb.New(*in.ExpiresAt)
	}

	if in.RevokedAt != nil {
		res.RevokedAt = timestamppb.New(*in.RevokedAt)
	}

	if in.LastUsedAt != nil {
		res.LastUsedAt = timestamppb.New(*in.LastUsedAt)
	}

	if in.ProjectID != nil {
		res.ProjectId = in.ProjectID.String()
	}

	if in.ProjectName != nil {
		res.ProjectName = *in.ProjectName
	}

	return res
}
