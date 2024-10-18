// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/client/kubernetes (interfaces: Applier)
//
// Generated by this command:
//
//	mockgen -package mock -destination=mocks_applier.go github.com/gardener/gardener/pkg/client/kubernetes Applier
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	kubernetes "github.com/gardener/gardener/pkg/client/kubernetes"
	gomock "go.uber.org/mock/gomock"
)

// MockApplier is a mock of Applier interface.
type MockApplier struct {
	ctrl     *gomock.Controller
	recorder *MockApplierMockRecorder
	isgomock struct{}
}

// MockApplierMockRecorder is the mock recorder for MockApplier.
type MockApplierMockRecorder struct {
	mock *MockApplier
}

// NewMockApplier creates a new mock instance.
func NewMockApplier(ctrl *gomock.Controller) *MockApplier {
	mock := &MockApplier{ctrl: ctrl}
	mock.recorder = &MockApplierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplier) EXPECT() *MockApplierMockRecorder {
	return m.recorder
}

// ApplyManifest mocks base method.
func (m *MockApplier) ApplyManifest(ctx context.Context, unstructured kubernetes.UnstructuredReader, options kubernetes.MergeFuncs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyManifest", ctx, unstructured, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyManifest indicates an expected call of ApplyManifest.
func (mr *MockApplierMockRecorder) ApplyManifest(ctx, unstructured, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyManifest", reflect.TypeOf((*MockApplier)(nil).ApplyManifest), ctx, unstructured, options)
}

// DeleteManifest mocks base method.
func (m *MockApplier) DeleteManifest(ctx context.Context, unstructured kubernetes.UnstructuredReader, opts ...kubernetes.DeleteManifestOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, unstructured}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteManifest", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteManifest indicates an expected call of DeleteManifest.
func (mr *MockApplierMockRecorder) DeleteManifest(ctx, unstructured any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, unstructured}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteManifest", reflect.TypeOf((*MockApplier)(nil).DeleteManifest), varargs...)
}
