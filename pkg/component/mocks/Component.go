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

// Code generated by mockery v1.0.0
package mocks

import ast "github.com/google/go-jsonnet/ast"
import component "github.com/marsfun/ksonnet/pkg/component"
import mock "github.com/stretchr/testify/mock"

// Component is an autogenerated mock type for the Component type
type Component struct {
	mock.Mock
}

// DeleteParam provides a mock function with given fields: path
func (_m *Component) DeleteParam(path []string) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Name provides a mock function with given fields: wantsNamedSpaced
func (_m *Component) Name(wantsNamedSpaced bool) string {
	ret := _m.Called(wantsNamedSpaced)

	var r0 string
	if rf, ok := ret.Get(0).(func(bool) string); ok {
		r0 = rf(wantsNamedSpaced)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Params provides a mock function with given fields: envName
func (_m *Component) Params(envName string) ([]component.ModuleParameter, error) {
	ret := _m.Called(envName)

	var r0 []component.ModuleParameter
	if rf, ok := ret.Get(0).(func(string) []component.ModuleParameter); ok {
		r0 = rf(envName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]component.ModuleParameter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(envName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetParam provides a mock function with given fields: path, value
func (_m *Component) SetParam(path []string, value interface{}) error {
	ret := _m.Called(path, value)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string, interface{}) error); ok {
		r0 = rf(path, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Summarize provides a mock function with given fields:
func (_m *Component) Summarize() (component.Summary, error) {
	ret := _m.Called()

	var r0 component.Summary
	if rf, ok := ret.Get(0).(func() component.Summary); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(component.Summary)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToNode provides a mock function with given fields: envName
func (_m *Component) ToNode(envName string) (string, ast.Node, error) {
	ret := _m.Called(envName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(envName)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 ast.Node
	if rf, ok := ret.Get(1).(func(string) ast.Node); ok {
		r1 = rf(envName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(ast.Node)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(envName)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Type provides a mock function with given fields:
func (_m *Component) Type() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
