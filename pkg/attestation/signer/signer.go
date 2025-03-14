//
// Copyright 2024 The Chainloop Authors.
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

package signer

import (
	"fmt"
	"strings"

	pb "github.com/chainloop-dev/chainloop/app/controlplane/api/controlplane/v1"
	"github.com/chainloop-dev/chainloop/pkg/attestation/signer/chainloop"
	"github.com/chainloop-dev/chainloop/pkg/attestation/signer/cosign"
	"github.com/chainloop-dev/chainloop/pkg/attestation/signer/signserver"
	"github.com/rs/zerolog"
	sigstoresigner "github.com/sigstore/sigstore/pkg/signature"
)

type Opts struct {
	SignServerOpts *SignServerOpts
	Vaultclient    pb.SigningServiceClient
}

type SignServerOpts struct {
	// CA certificate for TLS connection
	CAPath string
	// (optional) Client cert and passphrase for mutual TLS authentication
	AuthClientCertPath, AuthClientCertPass string
}

// GetSigner creates a new Signer based on input parameters
func GetSigner(keyPath string, logger zerolog.Logger, opts *Opts) (sigstoresigner.Signer, error) {
	var signer sigstoresigner.Signer
	if keyPath != "" {
		if strings.HasPrefix(keyPath, signserver.ReferenceScheme) {
			if opts.SignServerOpts == nil {
				// initialize empty opts (no custom CA, no client cert, no passphrase)
				opts.SignServerOpts = &SignServerOpts{}
			}
			host, worker, err := signserver.ParseKeyReference(keyPath)
			if err != nil {
				return nil, fmt.Errorf("failed to parse key: %w", err)
			}
			signer = signserver.NewSigner(host, worker,
				signserver.WithCAPath(opts.SignServerOpts.CAPath),
				signserver.WithClientCertPath(opts.SignServerOpts.AuthClientCertPath),
				signserver.WithClientCertPass(opts.SignServerOpts.AuthClientCertPass))
		} else {
			signer = cosign.NewSigner(keyPath, logger)
		}
	} else {
		signer = chainloop.NewSigner(opts.Vaultclient, logger)
	}

	return signer, nil
}
