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

package actions

import (
	"io"

	"github.com/ksonnet/ksonnet/pkg/app"
	"github.com/ksonnet/ksonnet/pkg/client"
	"github.com/ksonnet/ksonnet/pkg/cluster"
)

type runShowFn func(cluster.ShowConfig, ...cluster.ShowOpts) error

// RunShow runs `show`.
func RunShow(m map[string]interface{}) error {
	a, err := newShow(m)
	if err != nil {
		return err
	}

	return a.run()
}

type showOpt func(*Show)

// Show shows objects.
type Show struct {
	app            app.App
	clientConfig   *client.Config
	componentNames []string
	envName        string
	format         string

	out       io.Writer
	runShowFn runShowFn
}

// RunShow runs `show`
func newShow(m map[string]interface{}, opts ...showOpt) (*Show, error) {
	ol := newOptionLoader(m)

	s := &Show{
		app:            ol.LoadApp(),
		componentNames: ol.LoadStringSlice(OptionComponentNames),
		format:         ol.LoadString(OptionFormat),

		out:       ol.LoadBuffer(OptionBuffer),
		runShowFn: cluster.RunShow,
	}

	if ol.err != nil {
		return nil, ol.err
	}

	for _, opt := range opts {
		opt(s)
	}

	if err := setCurrentEnv(s.app, s, ol); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Show) run() error {
	config := cluster.ShowConfig{
		App:            s.app,
		ComponentNames: s.componentNames,
		EnvName:        s.envName,
		Format:         s.format,
		Out:            s.out,
	}

	return s.runShowFn(config)
}

func (s *Show) setCurrentEnv(name string) {
	s.envName = name
}
