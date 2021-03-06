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

package schema

import (
	"strings"

	rice "github.com/GeertJohan/go.rice"
	"github.com/ksonnet/ksonnet-lib/ksonnet-gen/astext"
	"github.com/marsfun/ksonnet/pkg/node"
	jsonnetutil "github.com/marsfun/ksonnet/pkg/util/jsonnet"
	"github.com/pkg/errors"
)

//go:generate rice embed-go

var valueExtractor *ValueExtractor

// ValueExtractorFactory returns a value extractor.
func ValueExtractorFactory() (*ValueExtractor, error) {
	if valueExtractor != nil {
		return valueExtractor, nil
	}

	assetsBox, err := rice.FindBox("assets")
	if err != nil {
		return nil, err
	}

	source, err := assetsBox.String("k8s.libsonnet")
	if err != nil {
		return nil, err
	}

	obj, err := jsonnetutil.Parse("k8s.libsonnet", source)
	if err != nil {
		return nil, err
	}

	valueExtractor = NewValueExtractor(obj)
	return valueExtractor, nil
}

// Values are values extracted from a manifest.
type Values struct {
	Lookup []string
	Setter string
	Value  interface{}
}

// ValueExtractor extracts Values from a manifest.
type ValueExtractor struct {
	object *node.Node
}

// NewValueExtractor creates an instance of ValueExtractor.
func NewValueExtractor(root *astext.Object) *ValueExtractor {
	return &ValueExtractor{
		object: node.New("root", root),
	}
}

// Extract extracts values from an object.
func (ve *ValueExtractor) Extract(gvk GVK, props Properties) (map[string]Values, error) {
	m := make(map[string]Values)
	cache := make(map[string]bool)

	paths := props.Paths(gvk)
	for _, path := range paths {
		item, err := ve.object.Search(path.Path...)
		if err != nil {
			continue
		}

		var manifestPath []string
		var found bool
		for _, p := range item.Path {
			if p == gvk.Kind {
				found = true
				continue
			}

			if !found {
				continue
			}

			manifestPath = append(manifestPath, p)
		}

		cachedPath := strings.Join(manifestPath, ".")
		if _, ok := cache[cachedPath]; ok {
			continue
		}

		cache[cachedPath] = true

		v, err := props.Value(manifestPath)
		if err != nil {
			return nil, errors.Wrapf(err, "retrieve values for %s", strings.Join(manifestPath, "."))
		}

		lookupPath := manifestPath
		if manifestPath[0] == "mixin" {
			lookupPath = manifestPath[1:]
		}

		p := strings.Join(item.Path, ".")
		dv := Values{
			Lookup: lookupPath,
			Setter: item.Name,
			Value:  v,
		}

		m[p] = dv
	}

	return m, nil
}
