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

package templates

import (
	"bytes"
	"regexp"
	"text/template"
)

var inputsPrefixRegexp = regexp.MustCompile(`{{\s*(inputs.)`)

// ApplyBinding renders an input template using bindings in Go templating format
func ApplyBinding(input string, bindings map[string]string) (string, error) {
	if input == "" {
		return "", nil
	}

	if bindings == nil {
		bindings = make(map[string]string)
	}

	// Support both `.inputs.foo` and `inputs.foo`
	input = inputsPrefixRegexp.ReplaceAllString(input, "{{ .inputs.")

	tmpl, err := template.New("chainloop").Option("missingkey=zero").Parse(input)
	if err != nil {
		return "", err
	}

	// Only support placeholders that are prefixed with "inputs.", ex `{{ inputs.foo }}
	namespacedBinding := map[string]any{"inputs": bindings}

	buffer := new(bytes.Buffer)
	err = tmpl.Execute(buffer, namespacedBinding)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
