package adventurer

import (
	reflect "reflect"

	adventurer "github.com/arfaghifari/guild-board/src/model/adventurer"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddCompletedQuest mocks base method.
func (m *MockRepository) AddCompletedQuest(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCompletedQuest", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCompletedQuest indicates an expected call of AddCompletedQuest.
func (mr *MockRepositoryMockRecorder) AddCompletedQuest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCompletedQuest", reflect.TypeOf((*MockRepository)(nil).AddCompletedQuest), arg0)
}

// Close mocks base method.
func (m *MockRepository) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockRepositoryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRepository)(nil).Close))
}

// CreateAdventurer mocks base method.
func (m *MockRepository) CreateAdventurer(arg0 adventurer.Adventurer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAdventurer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAdventurer indicates an expected call of CreateAdventurer.
func (mr *MockRepositoryMockRecorder) CreateAdventurer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdventurer", reflect.TypeOf((*MockRepository)(nil).CreateAdventurer), arg0)
}

// GetAdventurer mocks base method.
func (m *MockRepository) GetAdventurer(arg0 int64) (adventurer.Adventurer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdventurer", arg0)
	ret0, _ := ret[0].(adventurer.Adventurer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdventurer indicates an expected call of GetAdventurer.
func (mr *MockRepositoryMockRecorder) GetAdventurer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdventurer", reflect.TypeOf((*MockRepository)(nil).GetAdventurer), arg0)
}

// UpdateAdventurerRank mocks base method.
func (m *MockRepository) UpdateAdventurerRank(arg0 adventurer.Adventurer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdventurerRank", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAdventurerRank indicates an expected call of UpdateAdventurerRank.
func (mr *MockRepositoryMockRecorder) UpdateAdventurerRank(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdventurerRank", reflect.TypeOf((*MockRepository)(nil).UpdateAdventurerRank), arg0)
}
