// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/apiserver/pkg/authorization/authorizer (interfaces: Authorizer)
//
// Generated by this command:
//
//	mockgen -package authorizer -destination=mocks.go k8s.io/apiserver/pkg/authorization/authorizer Authorizer
//

// Package authorizer is a generated GoMock package.
package authorizer

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	authorizer "k8s.io/apiserver/pkg/authorization/authorizer"
)

// MockAuthorizer is a mock of Authorizer interface.
type MockAuthorizer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizerMockRecorder
	isgomock struct{}
}

// MockAuthorizerMockRecorder is the mock recorder for MockAuthorizer.
type MockAuthorizerMockRecorder struct {
	mock *MockAuthorizer
}

// NewMockAuthorizer creates a new mock instance.
func NewMockAuthorizer(ctrl *gomock.Controller) *MockAuthorizer {
	mock := &MockAuthorizer{ctrl: ctrl}
	mock.recorder = &MockAuthorizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizer) EXPECT() *MockAuthorizerMockRecorder {
	return m.recorder
}

// Authorize mocks base method.
func (m *MockAuthorizer) Authorize(ctx context.Context, a authorizer.Attributes) (authorizer.Decision, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorize", ctx, a)
	ret0, _ := ret[0].(authorizer.Decision)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Authorize indicates an expected call of Authorize.
func (mr *MockAuthorizerMockRecorder) Authorize(ctx, a any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*MockAuthorizer)(nil).Authorize), ctx, a)
}
