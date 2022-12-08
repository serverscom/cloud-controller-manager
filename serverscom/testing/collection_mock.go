package serverscom_testing

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	serverscom "github.com/serverscom/serverscom-go-client/pkg"
)

// MockCollection is a mock of Collection interface.
type MockCollection[K any] struct {
	ctrl     *gomock.Controller
	recorder *MockCollectionMockRecorder[K]
}

// MockCollectionMockRecorder is the mock recorder for MockCollection.
type MockCollectionMockRecorder[K any] struct {
	mock *MockCollection[K]
}

// NewMockCollection creates a new mock instance.
func NewMockCollection[K any](ctrl *gomock.Controller) *MockCollection[K] {
	mock := &MockCollection[K]{ctrl: ctrl}
	mock.recorder = &MockCollectionMockRecorder[K]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCollection[K]) EXPECT() *MockCollectionMockRecorder[K] {
	return m.recorder
}

// Collect mocks base method.
func (m *MockCollection[K]) Collect(ctx context.Context) ([]K, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Collect", ctx)
	ret0, _ := ret[0].([]K)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Collect indicates an expected call of Collect.
func (mr *MockCollectionMockRecorder[K]) Collect(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collect", reflect.TypeOf((*MockCollection[K])(nil).Collect), ctx)
}

// FirstPage mocks base method.
func (m *MockCollection[K]) FirstPage(ctx context.Context) ([]K, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FirstPage", ctx)
	ret0, _ := ret[0].([]K)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FirstPage indicates an expected call of FirstPage.
func (mr *MockCollectionMockRecorder[K]) FirstPage(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirstPage", reflect.TypeOf((*MockCollection[K])(nil).FirstPage), ctx)
}

// HasFirstPage mocks base method.
func (m *MockCollection[K]) HasFirstPage() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasFirstPage")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasFirstPage indicates an expected call of HasFirstPage.
func (mr *MockCollectionMockRecorder[K]) HasFirstPage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasFirstPage", reflect.TypeOf((*MockCollection[K])(nil).HasFirstPage))
}

// HasLastPage mocks base method.
func (m *MockCollection[K]) HasLastPage() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasLastPage")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasLastPage indicates an expected call of HasLastPage.
func (mr *MockCollectionMockRecorder[K]) HasLastPage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasLastPage", reflect.TypeOf((*MockCollection[K])(nil).HasLastPage))
}

// HasNextPage mocks base method.
func (m *MockCollection[K]) HasNextPage() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNextPage")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNextPage indicates an expected call of HasNextPage.
func (mr *MockCollectionMockRecorder[K]) HasNextPage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNextPage", reflect.TypeOf((*MockCollection[K])(nil).HasNextPage))
}

// HasPreviousPage mocks base method.
func (m *MockCollection[K]) HasPreviousPage() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPreviousPage")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasPreviousPage indicates an expected call of HasPreviousPage.
func (mr *MockCollectionMockRecorder[K]) HasPreviousPage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPreviousPage", reflect.TypeOf((*MockCollection[K])(nil).HasPreviousPage))
}

// IsClean mocks base method.
func (m *MockCollection[K]) IsClean() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClean")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsClean indicates an expected call of IsClean.
func (mr *MockCollectionMockRecorder[K]) IsClean() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClean", reflect.TypeOf((*MockCollection[K])(nil).IsClean))
}

// LastPage mocks base method.
func (m *MockCollection[K]) LastPage(ctx context.Context) ([]K, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastPage", ctx)
	ret0, _ := ret[0].([]K)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastPage indicates an expected call of LastPage.
func (mr *MockCollectionMockRecorder[K]) LastPage(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastPage", reflect.TypeOf((*MockCollection[K])(nil).LastPage), ctx)
}

// List mocks base method.
func (m *MockCollection[K]) List(ctx context.Context) ([]K, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]K)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCollectionMockRecorder[K]) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCollection[K])(nil).List), ctx)
}

// NextPage mocks base method.
func (m *MockCollection[K]) NextPage(ctx context.Context) ([]K, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextPage", ctx)
	ret0, _ := ret[0].([]K)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextPage indicates an expected call of NextPage.
func (mr *MockCollectionMockRecorder[K]) NextPage(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextPage", reflect.TypeOf((*MockCollection[K])(nil).NextPage), ctx)
}

// PreviousPage mocks base method.
func (m *MockCollection[K]) PreviousPage(ctx context.Context) ([]K, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreviousPage", ctx)
	ret0, _ := ret[0].([]K)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PreviousPage indicates an expected call of PreviousPage.
func (mr *MockCollectionMockRecorder[K]) PreviousPage(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreviousPage", reflect.TypeOf((*MockCollection[K])(nil).PreviousPage), ctx)
}

// Refresh mocks base method.
func (m *MockCollection[K]) Refresh(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Refresh indicates an expected call of Refresh.
func (mr *MockCollectionMockRecorder[K]) Refresh(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockCollection[K])(nil).Refresh), ctx)
}

// SetPage mocks base method.
func (m *MockCollection[K]) SetPage(page int) serverscom.Collection[K] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPage", page)
	ret0, _ := ret[0].(serverscom.Collection[K])
	return ret0
}

// SetPage indicates an expected call of SetPage.
func (mr *MockCollectionMockRecorder[K]) SetPage(page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPage", reflect.TypeOf((*MockCollection[K])(nil).SetPage), page)
}

// SetParam mocks base method.
func (m *MockCollection[K]) SetParam(name, value string) serverscom.Collection[K] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetParam", name, value)
	ret0, _ := ret[0].(serverscom.Collection[K])
	return ret0
}

// SetParam indicates an expected call of SetParam.
func (mr *MockCollectionMockRecorder[K]) SetParam(name, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParam", reflect.TypeOf((*MockCollection[K])(nil).SetParam), name, value)
}

// SetPerPage mocks base method.
func (m *MockCollection[K]) SetPerPage(perPage int) serverscom.Collection[K] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPerPage", perPage)
	ret0, _ := ret[0].(serverscom.Collection[K])
	return ret0
}

// SetPerPage indicates an expected call of SetPerPage.
func (mr *MockCollectionMockRecorder[K]) SetPerPage(perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPerPage", reflect.TypeOf((*MockCollection[K])(nil).SetPerPage), perPage)
}
