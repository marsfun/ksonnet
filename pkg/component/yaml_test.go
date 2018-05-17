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

package component

import (
	"io/ioutil"
	"testing"

	"github.com/ksonnet/ksonnet-lib/ksonnet-gen/astext"

	"github.com/marsfun/ksonnet/pkg/app/mocks"
	"github.com/marsfun/ksonnet/pkg/util/test"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestYAML_Type(t *testing.T) {
	test.WithApp(t, "/app", func(a *mocks.App, fs afero.Fs) {

		test.StageFile(t, fs, "params-mixed.libsonnet", "/app/components/params.libsonnet")
		test.StageFile(t, fs, "deployment.yaml", "/app/components/deployment.yaml")

		y := NewYAML(a, "", "/app/components/deployment.yaml", "/app/components/params.libsonnet")

		require.Equal(t, TypeYAML, y.Type())
	})
}

func TestYAML_Name(t *testing.T) {
	test.WithApp(t, "/app", func(a *mocks.App, fs afero.Fs) {

		test.StageFile(t, fs, "params-mixed.libsonnet", "/app/components/params.libsonnet")
		test.StageFile(t, fs, "deployment.yaml", "/app/components/deployment.yaml")

		y := NewYAML(a, "", "/app/components/deployment.yaml", "/app/components/params.libsonnet")

		cases := []struct {
			name         string
			isNameSpaced bool
			expected     string
		}{
			{
				name:         "wants namespaced",
				isNameSpaced: true,
				expected:     "deployment",
			},
			{
				name:         "no namespace",
				isNameSpaced: false,
				expected:     "deployment",
			},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				require.Equal(t, tc.expected, y.Name(tc.isNameSpaced))
			})
		}
	})

}

func TestYAML_Params(t *testing.T) {
	test.WithApp(t, "/app", func(a *mocks.App, fs afero.Fs) {

		test.StageFile(t, fs, "params-mixed.libsonnet", "/app/components/params.libsonnet")
		test.StageFile(t, fs, "deployment.yaml", "/app/components/deployment.yaml")

		y := NewYAML(a, "", "/app/components/deployment.yaml", "/app/components/params.libsonnet")
		params, err := y.Params("")
		require.NoError(t, err)

		require.Len(t, params, 1)

		param := params[0]
		expected := ModuleParameter{
			Component: "deployment",
			Key:       "metadata.labels",
			Value:     `{"label1":"label1","label2":"label2"}`,
		}
		require.Equal(t, expected, param)
	})
}

func TestYAML_Params_literal(t *testing.T) {
	test.WithApp(t, "/app", func(a *mocks.App, fs afero.Fs) {

		test.StageFile(t, fs, "params-mixed.libsonnet", "/params.libsonnet")
		test.StageFile(t, fs, "clusterrole-cert-manager.yaml", "/clusterrole-cert-manager.yaml")

		y := NewYAML(a, "", "/clusterrole-cert-manager.yaml", "/params.libsonnet")
		params, err := y.Params("")
		require.NoError(t, err)

		require.Len(t, params, 1)

		param := params[0]
		expected := ModuleParameter{
			Component: "clusterrole-cert-manager",
			Key:       "metadata.name",
			Value:     "cert-manager2",
		}
		require.Equal(t, expected, param)
	})
}

func TestYAML_Params_extra_entries(t *testing.T) {
	test.WithApp(t, "/app", func(a *mocks.App, fs afero.Fs) {

		test.StageFile(t, fs, "deployment.yaml", "/deployment.yaml")
		test.StageFile(t, fs, "params-annotations.libsonnet", "/params.libsonnet")

		y := NewYAML(a, "", "/deployment.yaml", "/params.libsonnet")
		params, err := y.Params("")
		require.NoError(t, err)

		expected := []ModuleParameter{
			{
				Component: "deployment",
				Key:       "metadata.annotations",
				Value:     `{"size":"large"}`,
			},
		}
		require.Equal(t, expected, params)
	})
}

func TestYAML_SetParam(t *testing.T) {
	test.WithApp(t, "/app", func(a *mocks.App, fs afero.Fs) {

		test.StageFile(t, fs, "certificate-crd.yaml", "/certificate-crd.yaml")
		test.StageFile(t, fs, "params-no-entry.libsonnet", "/params.libsonnet")

		y := NewYAML(a, "", "/certificate-crd.yaml", "/params.libsonnet")

		err := y.SetParam([]string{"spec", "version"}, "v2")
		require.NoError(t, err)

		b, err := afero.ReadFile(fs, "/params.libsonnet")
		require.NoError(t, err)

		expected := testdata(t, "updated-yaml-params.libsonnet")

		require.Equal(t, string(expected), string(b))
	})
}

func TestYAML_DeleteParam(t *testing.T) {
	test.WithApp(t, "/app", func(a *mocks.App, fs afero.Fs) {

		test.StageFile(t, fs, "certificate-crd.yaml", "/certificate-crd.yaml")
		test.StageFile(t, fs, "params-with-entry.libsonnet", "/params.libsonnet")

		y := NewYAML(a, "", "/certificate-crd.yaml", "/params.libsonnet")

		err := y.DeleteParam([]string{"spec", "version"})
		require.NoError(t, err)

		b, err := afero.ReadFile(fs, "/params.libsonnet")
		require.NoError(t, err)

		expected := testdata(t, "params-delete-entry.libsonnet")

		require.Equal(t, string(expected), string(b))
	})
}

func TestYAML_Summarize(t *testing.T) {
	test.WithApp(t, "/app", func(a *mocks.App, fs afero.Fs) {

		test.StageFile(t, fs, "clusterrole-cert-manager.yaml", "/components/clusterrole-cert-manager.yaml")
		test.StageFile(t, fs, "params-no-entry.libsonnet", "/components/params.libsonnet")

		y := NewYAML(a, "", "/components/clusterrole-cert-manager.yaml", "/components/params.libsonnet")

		list, err := y.Summarize()
		require.NoError(t, err)

		expected := Summary{
			ComponentName: "clusterrole-cert-manager",
			Type:          "yaml",
			APIVersion:    "rbac.authorization.k8s.io/v1beta1",
			Kind:          "ClusterRole",
			Name:          "cert-manager",
		}

		require.Equal(t, expected, list)
	})
}

func TestYAML_Summarize_json(t *testing.T) {
	test.WithApp(t, "/app", func(a *mocks.App, fs afero.Fs) {

		test.StageFile(t, fs, "certificate-crd.json", "/components/certificate-crd.json")
		test.StageFile(t, fs, "params-no-entry.libsonnet", "/components/params.libsonnet")

		y := NewYAML(a, "", "/components/certificate-crd.json", "/components/params.libsonnet")

		list, err := y.Summarize()
		require.NoError(t, err)

		expected := Summary{
			ComponentName: "certificate-crd",
			Type:          "json",
			APIVersion:    "apiextensions.k8s.io/v1beta1",
			Kind:          "CustomResourceDefinition",
			Name:          "certificates_certmanager_k8s_io",
		}

		require.Equal(t, expected, list)
	})
}

func Test_mapToPaths(t *testing.T) {
	m := map[string]interface{}{
		"metadata": map[string]interface{}{
			"name": "name",
			"labels": map[string]interface{}{
				"label1": "label1",
			},
		},
	}

	lookup := map[string]bool{
		"metadata.name":   true,
		"metadata.labels": true,
	}

	got := mapToPaths(m, lookup, nil)

	expected := []paramPath{
		{path: []string{"metadata", "labels"}, value: map[string]interface{}{"label1": "label1"}},
		{path: []string{"metadata", "name"}, value: "name"},
	}

	require.Equal(t, expected, got)

}

func testdata(t *testing.T, name string) []byte {
	b, err := ioutil.ReadFile("testdata/" + name)
	require.NoError(t, err, "read testdata %s", name)
	return b
}

func TestYAML_ToNode(t *testing.T) {
	type fields struct {
		paramsPath    string
		componentPath string
	}
	type args struct {
		envName string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		nodeName string
		wantErr  bool
	}{
		{
			name: "in general",
			fields: fields{
				paramsPath:    "params-mixed.libsonnet",
				componentPath: "deployment.yaml",
			},
			nodeName: "deployment",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test.WithApp(t, "/app", func(a *mocks.App, fs afero.Fs) {
				test.StageFile(t, fs, tt.fields.paramsPath, "/app/components/params.libsonnet")
				test.StageFile(t, fs, tt.fields.componentPath, "/app/components/deployment.yaml")

				y := NewYAML(a, "", "/app/components/deployment.yaml", "/app/components/params.libsonnet")

				nodeName, node, err := y.ToNode(tt.args.envName)
				if tt.wantErr {
					require.Error(t, err)
					return
				}

				assert.Equal(t, tt.nodeName, nodeName, "node name was not expected")
				assert.IsType(t, &astext.Object{}, node)
			})
		})
	}
}
