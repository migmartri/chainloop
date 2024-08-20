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

package biz

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIdentifyAndValidateRawContract(t *testing.T) {
	testData := []struct {
		filename          string
		wantFormat        ContractRawFormat
		wantValidationErr bool
		wantFormatErr     bool
	}{
		{
			filename:   "contract.cue",
			wantFormat: ContractRawFormatCUE,
		},
		{
			filename:   "contract.json",
			wantFormat: ContractRawFormatJSON,
		},
		{
			filename:          "invalid_contract.json",
			wantValidationErr: true,
		},
		{
			filename:   "contract.yaml",
			wantFormat: ContractRawFormatYAML,
		},
		{
			filename:          "invalid_contract.yaml",
			wantValidationErr: true,
		},
		{
			filename:      "invalid_format.json",
			wantFormatErr: true,
		},
	}

	for _, tc := range testData {
		t.Run(tc.filename, func(t *testing.T) {
			// load file from testdata/contracts
			data, err := os.ReadFile("testdata/contracts/" + tc.filename)
			require.NoError(t, err)

			contract, err := identifyUnMarshalAndValidateRawContract(data)
			if tc.wantValidationErr {
				assert.Error(t, err)
				assert.True(t, IsErrValidation(err))
				return
			} else if tc.wantFormatErr {
				assert.Error(t, err)
				assert.False(t, IsErrValidation(err))
				return
			}

			require.NoError(t, err)

			assert.Equal(t, tc.wantFormat, contract.Format)
			assert.Equal(t, data, contract.Raw)
		})
	}
}

func TestIdentifyFormat(t *testing.T) {
	testData := []struct {
		filename   string
		wantFormat ContractRawFormat
		wantErr    bool
	}{
		{
			filename:   "contract.cue",
			wantFormat: ContractRawFormatCUE,
		},
		{
			filename:   "contract.json",
			wantFormat: ContractRawFormatJSON,
		},
		{
			filename:   "invalid_contract.json",
			wantFormat: ContractRawFormatJSON,
		},
		{
			filename:   "contract.yaml",
			wantFormat: ContractRawFormatYAML,
		},
		{
			filename:   "invalid_contract.yaml",
			wantFormat: ContractRawFormatYAML,
		},
		{
			filename: "invalid_format.json",
			wantErr:  true,
		},
	}

	for _, tt := range testData {
		t.Run(tt.filename, func(t *testing.T) {
			// load file from testdata/contracts
			data, err := os.ReadFile("testdata/contracts/" + tt.filename)
			require.NoError(t, err)

			format, err := identifyFormat(data)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			assert.Equal(t, tt.wantFormat, format)
		})
	}
}