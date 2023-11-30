package quest

import (
	reflect "reflect"

	adventurer "github.com/arfaghifari/guild-board/src/model/adventurer"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type AdvMockRepository struct {
	ctrl     *gomock.Controller
	recorder *AdvMockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type AdvMockRepositoryMockRecorder struct {
	mock *AdvMockRepository
}

// NewMockRepository creates a new mock instance.
func NewAdvMockRepository(ctrl *gomock.Controller) *AdvMockRepository {
	mock := &AdvMockRepository{ctrl: ctrl}
	mock.recorder = &AdvMockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *AdvMockRepository) EXPECT() *AdvMockRepositoryMockRecorder {
	return m.recorder
}

// AddCompletedQuest mocks base method.
func (m *AdvMockRepository) AddCompletedQuest(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCompletedQuest", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCompletedQuest indicates an expected call of AddCompletedQuest.
func (mr *AdvMockRepositoryMockRecorder) AddCompletedQuest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCompletedQuest", reflect.TypeOf((*AdvMockRepository)(nil).AddCompletedQuest), arg0)
}

// Close mocks base method.
func (m *AdvMockRepository) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *AdvMockRepositoryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRepository)(nil).Close))
}

// CreateAdventurer mocks base method.
func (m *AdvMockRepository) CreateAdventurer(arg0 adventurer.Adventurer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAdventurer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAdventurer indicates an expected call of CreateAdventurer.
func (mr *AdvMockRepositoryMockRecorder) CreateAdventurer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdventurer", reflect.TypeOf((*AdvMockRepository)(nil).CreateAdventurer), arg0)
}

// GetAdventurer mocks base method.
func (m *AdvMockRepository) GetAdventurer(arg0 int64) (adventurer.Adventurer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdventurer", arg0)
	ret0, _ := ret[0].(adventurer.Adventurer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdventurer indicates an expected call of GetAdventurer.
func (mr *AdvMockRepositoryMockRecorder) GetAdventurer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdventurer", reflect.TypeOf((*AdvMockRepository)(nil).GetAdventurer), arg0)
}

// UpdateAdventurerRank mocks base method.
func (m *AdvMockRepository) UpdateAdventurerRank(arg0 adventurer.Adventurer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdventurerRank", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAdventurerRank indicates an expected call of UpdateAdventurerRank.
func (mr *AdvMockRepositoryMockRecorder) UpdateAdventurerRank(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdventurerRank", reflect.TypeOf((*AdvMockRepository)(nil).UpdateAdventurerRank), arg0)
}
