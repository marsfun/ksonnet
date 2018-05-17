// Copyright 2018 The ksonnet authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package params

import (
	"path/filepath"
	"testing"

	"github.com/marsfun/ksonnet/metadata/params"
	"github.com/marsfun/ksonnet/pkg/util/test"
	"github.com/stretchr/testify/require"
)

func TestEnvGlobalsSet(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		output string
		params params.Params
	}{
		{
			name:   "in general",
			input:  filepath.Join("env", "globals", "set-global", "in.libsonnet"),
			output: filepath.Join("env", "globals", "set-global", "out.libsonnet"),
			params: params.Params{
				"group": "dev",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			snippet := test.ReadTestData(t, tc.input)

			egs := NewEnvGlobalsSet()

			got, err := egs.Set(snippet, tc.params)
			require.NoError(t, err)

			expected := test.ReadTestData(t, tc.output)
			require.Equal(t, expected, got)
		})
	}
}
