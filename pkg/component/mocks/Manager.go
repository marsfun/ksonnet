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

import "github.com/marsfun/ksonnet/pkg/app"
import "github.com/marsfun/ksonnet/pkg/component"
import mock "github.com/stretchr/testify/mock"
import params "github.com/marsfun/ksonnet/metadata/params"
import "github.com/marsfun/ksonnet/pkg/prototype"

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// Component provides a mock function with given fields: ksApp, module, componentName
func (_m *Manager) Component(ksApp app.App, module string, componentName string) (component.Component, error) {
	ret := _m.Called(ksApp, module, componentName)

	var r0 component.Component
	if rf, ok := ret.Get(0).(func(app.App, string, string) component.Component); ok {
		r0 = rf(ksApp, module, componentName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(component.Component)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(app.App, string, string) error); ok {
		r1 = rf(ksApp, module, componentName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Components provides a mock function with given fields: ns
func (_m *Manager) Components(ns component.Module) ([]component.Component, error) {
	ret := _m.Called(ns)

	var r0 []component.Component
	if rf, ok := ret.Get(0).(func(component.Module) []component.Component); ok {
		r0 = rf(ns)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]component.Component)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(component.Module) error); ok {
		r1 = rf(ns)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateComponent provides a mock function with given fields: ksApp, name, text, _a3, templateType
func (_m *Manager) CreateComponent(ksApp app.App, name string, text string, _a3 params.Params, templateType prototype.TemplateType) (string, error) {
	ret := _m.Called(ksApp, name, text, _a3, templateType)

	var r0 string
	if rf, ok := ret.Get(0).(func(app.App, string, string, params.Params, prototype.TemplateType) string); ok {
		r0 = rf(ksApp, name, text, _a3, templateType)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(app.App, string, string, params.Params, prototype.TemplateType) error); ok {
		r1 = rf(ksApp, name, text, _a3, templateType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateModule provides a mock function with given fields: ksApp, name
func (_m *Manager) CreateModule(ksApp app.App, name string) error {
	ret := _m.Called(ksApp, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(app.App, string) error); ok {
		r0 = rf(ksApp, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Module provides a mock function with given fields: ksApp, module
func (_m *Manager) Module(ksApp app.App, module string) (component.Module, error) {
	ret := _m.Called(ksApp, module)

	var r0 component.Module
	if rf, ok := ret.Get(0).(func(app.App, string) component.Module); ok {
		r0 = rf(ksApp, module)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(component.Module)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(app.App, string) error); ok {
		r1 = rf(ksApp, module)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Modules provides a mock function with given fields: ksApp, envName
func (_m *Manager) Modules(ksApp app.App, envName string) ([]component.Module, error) {
	ret := _m.Called(ksApp, envName)

	var r0 []component.Module
	if rf, ok := ret.Get(0).(func(app.App, string) []component.Module); ok {
		r0 = rf(ksApp, envName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]component.Module)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(app.App, string) error); ok {
		r1 = rf(ksApp, envName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NSResolveParams provides a mock function with given fields: ns
func (_m *Manager) NSResolveParams(ns component.Module) (string, error) {
	ret := _m.Called(ns)

	var r0 string
	if rf, ok := ret.Get(0).(func(component.Module) string); ok {
		r0 = rf(ns)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(component.Module) error); ok {
		r1 = rf(ns)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResolvePath provides a mock function with given fields: ksApp, path
func (_m *Manager) ResolvePath(ksApp app.App, path string) (component.Module, component.Component, error) {
	ret := _m.Called(ksApp, path)

	var r0 component.Module
	if rf, ok := ret.Get(0).(func(app.App, string) component.Module); ok {
		r0 = rf(ksApp, path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(component.Module)
		}
	}

	var r1 component.Component
	if rf, ok := ret.Get(1).(func(app.App, string) component.Component); ok {
		r1 = rf(ksApp, path)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(component.Component)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(app.App, string) error); ok {
		r2 = rf(ksApp, path)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
