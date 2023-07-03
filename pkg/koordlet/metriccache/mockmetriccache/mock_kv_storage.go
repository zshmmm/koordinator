// /*
// Copyright 2022 The Koordinator Authors.
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
// */
//

// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/koordlet/metriccache/kv_storage.go

// Package mock_metriccache is a generated GoMock package.
package mock_metriccache

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockKVStorage is a mock of KVStorage interface.
type MockKVStorage struct {
	ctrl     *gomock.Controller
	recorder *MockKVStorageMockRecorder
}

// MockKVStorageMockRecorder is the mock recorder for MockKVStorage.
type MockKVStorageMockRecorder struct {
	mock *MockKVStorage
}

// NewMockKVStorage creates a new mock instance.
func NewMockKVStorage(ctrl *gomock.Controller) *MockKVStorage {
	mock := &MockKVStorage{ctrl: ctrl}
	mock.recorder = &MockKVStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKVStorage) EXPECT() *MockKVStorageMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockKVStorage) Get(key interface{}) (interface{}, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockKVStorageMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKVStorage)(nil).Get), key)
}

// Set mocks base method.
func (m *MockKVStorage) Set(key, value interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", key, value)
}

// Set indicates an expected call of Set.
func (mr *MockKVStorageMockRecorder) Set(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockKVStorage)(nil).Set), key, value)
}
