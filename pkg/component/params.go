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
	"regexp"

	"github.com/marsfun/ksonnet/pkg/util/jsonnet"
	"github.com/sirupsen/logrus"
)

func applyGlobals(params string) (string, error) {
	vm := jsonnet.NewVM()

	vm.ExtCode("params", params)
	return vm.EvaluateSnippet("applyGlobals", snippetMapGlobal)
}

var snippetMapGlobal = `
local params = std.extVar("params");
local applyGlobal = function(key, value) std.mergePatch(value, params.global);

{
	components: std.mapWithKey(applyGlobal, params.components)
}
`

var (
	reParamSwap = regexp.MustCompile(`(?m)import "\.\.\/\.\.\/components\/params\.libsonnet"`)
)

// upgradeParams replaces relative params imports with an extVar to handle
// multiple component namespaces.
// NOTE: It warns when it makes a change. This serves as a temporary fix until
// ksonnet generates the correct file.
func upgradeParams(envName, in string) string {
	if reParamSwap.MatchString(in) {
		logrus.Warnf("rewriting %q environment params to not use relative paths", envName)
		return reParamSwap.ReplaceAllLiteralString(in, `std.extVar("__ksonnet/params")`)
	}

	return in
}
