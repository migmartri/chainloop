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

package action

import (
	"context"

	"github.com/chainloop-dev/chainloop/app/cli/action"
	pb "github.com/chainloop-dev/chainloop/app/controlplane/api/controlplane/v1"
)

type WorkflowContractDelete struct {
	cfg *action.ActionsOpts
}

func NewWorkflowContractDelete(cfg *action.ActionsOpts) *WorkflowContractDelete {
	return &WorkflowContractDelete{cfg}
}

func (action *WorkflowContractDelete) Run(contractID string) error {
	client := pb.NewWorkflowContractServiceClient(action.cfg.CPConnection)
	if _, err := client.Delete(context.Background(), &pb.WorkflowContractServiceDeleteRequest{Id: contractID}); err != nil {
		action.cfg.Logger.Debug().Err(err).Msg("making the API request")
		return err
	}

	return nil
}
