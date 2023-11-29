package quest

import (
	"errors"
	"testing"

	modelAdv "github.com/arfaghifari/guild-board/src/model/adventurer"
	model "github.com/arfaghifari/guild-board/src/model/quest"
	advRepo "github.com/arfaghifari/guild-board/src/repository/adventurer"
	repo "github.com/arfaghifari/guild-board/src/repository/quest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var adv = modelAdv.Adventurer{
	ID:             1,
	Name:           "andi",
	Rank:           11,
	CompletedQuest: 1,
}

var bulkQuest = []model.Quest{
	{
		ID:           1,
		Name:         "menyelamatkan kucing",
		Description:  "menyelamatkan kucing yang terjebak di atas pohon",
		MinimumRank:  11,
		RewardNumber: 200000,
		Status:       0,
	},
	{
		ID:           2,
		Name:         "membersihkan selokan",
		Description:  "membersihkan selokan penuh dengan lumut",
		MinimumRank:  12,
		RewardNumber: 200000,
		Status:       0,
	},
	{
		ID:           3,
		Name:         "Supir perjalanan",
		Description:  "Mengantar pulang pergi dan keliling kota, Jakarta-Bandung, Sudah di kasih makan",
		MinimumRank:  13,
		RewardNumber: 600000,
		Status:       2,
	},
}

var bulkQuestByStatus = []model.GetQuestByStatus{
	{
		ID:           1,
		Name:         "menyelamatkan kucing",
		Description:  "menyelamatkan kucing yang terjebak di atas pohon",
		MinimumRank:  11,
		RewardNumber: 200000,
	},
	{
		ID:           2,
		Name:         "membersihkan selokan",
		Description:  "membersihkan selokan penuh dengan lumut",
		MinimumRank:  11,
		RewardNumber: 200000,
	},
	{
		ID:           3,
		Name:         "Supir perjalanan",
		Description:  "Mengantar pulang pergi dan keliling kota, Jakarta-Bandung, Sudah di kasih makan",
		MinimumRank:  13,
		RewardNumber: 600000,
	},
}

func TestCreateQuest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := repo.NewMockRepository(mockCtrl)
	testUsecase := &usecase{repo: mockRepo}
	mockRepo.EXPECT().CreateQuest(bulkQuest[0]).Return(nil).Times(1)
	err := testUsecase.CreateQuest(bulkQuest[0])
	assert.Nil(t, err)
}

func TestUpdateQuestReward(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := repo.NewMockRepository(mockCtrl)
	testUsecase := &usecase{repo: mockRepo}
	mockRepo.EXPECT().UpdateQuestReward(bulkQuest[0]).Return(nil).Times(1)
	err := testUsecase.UpdateQuestReward(bulkQuest[0])
	assert.Nil(t, err)
}

func TestUpdateQuestRank(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := repo.NewMockRepository(mockCtrl)
	testUsecase := &usecase{repo: mockRepo}
	mockRepo.EXPECT().UpdateQuestRank(bulkQuest[0]).Return(nil).Times(1)
	err := testUsecase.UpdateQuestRank(bulkQuest[0])
	assert.Nil(t, err)
}

func TestDeleteQuest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := repo.NewMockRepository(mockCtrl)
	testUsecase := &usecase{repo: mockRepo}
	mockRepo.EXPECT().DeleteQuest(bulkQuest[0]).Return(nil).Times(1)
	err := testUsecase.DeleteQuest(bulkQuest[0])
	assert.Nil(t, err)
}

func TestGetQuestByStatusA(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := repo.NewMockRepository(mockCtrl)
	testUsecase := &usecase{repo: mockRepo}
	mockRepo.EXPECT().GetAllAvailableQuest().Return(bulkQuestByStatus[0:1], nil).Times(1)
	res, err := testUsecase.GetQuestByStatus(0)
	assert.Equal(t, bulkQuestByStatus[0:1], res)
	assert.Nil(t, err)
}

func TestGetQuestByStatusC(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := repo.NewMockRepository(mockCtrl)
	testUsecase := &usecase{repo: mockRepo}
	mockRepo.EXPECT().GetAllCompletedQuest().Return(bulkQuestByStatus[2:], nil).Times(1)
	res, err := testUsecase.GetQuestByStatus(2)
	assert.Equal(t, bulkQuestByStatus[2:], res)
	assert.Nil(t, err)
}

func TestTakeQuest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := repo.NewMockRepository(mockCtrl)
	mockAdvRepo := advRepo.NewMockRepository(mockCtrl)
	testUsecase := &usecase{repo: mockRepo, repoAdv: mockAdvRepo}
	mockRepo.EXPECT().GetQuest(int64(1)).Return(bulkQuest[0], nil).Times(1)
	mockAdvRepo.EXPECT().GetAdventurer(int64(1)).Return(adv, nil).Times(1)
	mockRepo.EXPECT().CreateTakenBy(int64(1), int64(1)).Return(nil).Times(1)
	copyQuest := bulkQuest[0]
	copyQuest.Status = 1
	mockRepo.EXPECT().UpdateQuestStatus(copyQuest).Return(nil).Times(1)
	err := testUsecase.TakeQuest(int64(1), int64(1))
	assert.Nil(t, err)
}

func TestTakeQuestStatus(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := repo.NewMockRepository(mockCtrl)
	mockAdvRepo := advRepo.NewMockRepository(mockCtrl)
	testUsecase := &usecase{repo: mockRepo, repoAdv: mockAdvRepo}
	mockRepo.EXPECT().GetQuest(int64(3)).Return(bulkQuest[2], nil).Times(1)
	err := testUsecase.TakeQuest(int64(3), int64(1))
	assert.NotNil(t, err)
}

func TestTakeQuestRank(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := repo.NewMockRepository(mockCtrl)
	mockAdvRepo := advRepo.NewMockRepository(mockCtrl)
	testUsecase := &usecase{repo: mockRepo, repoAdv: mockAdvRepo}
	mockRepo.EXPECT().GetQuest(int64(2)).Return(bulkQuest[1], nil).Times(1)
	mockAdvRepo.EXPECT().GetAdventurer(int64(1)).Return(adv, nil).Times(1)
	err := testUsecase.TakeQuest(int64(2), int64(1))
	assert.NotNil(t, err)
}

func TestTakeQuestNoQuest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := repo.NewMockRepository(mockCtrl)
	mockAdvRepo := advRepo.NewMockRepository(mockCtrl)
	testUsecase := &usecase{repo: mockRepo, repoAdv: mockAdvRepo}
	mockRepo.EXPECT().GetQuest(int64(1)).Return(model.Quest{}, errors.New("err")).Times(1)
	err := testUsecase.TakeQuest(int64(1), int64(1))
	assert.NotNil(t, err)
}

func TestTakeQuestNoAdv(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := repo.NewMockRepository(mockCtrl)
	mockAdvRepo := advRepo.NewMockRepository(mockCtrl)
	testUsecase := &usecase{repo: mockRepo, repoAdv: mockAdvRepo}
	mockRepo.EXPECT().GetQuest(int64(1)).Return(bulkQuest[0], nil).Times(1)
	mockAdvRepo.EXPECT().GetAdventurer(int64(1)).Return(modelAdv.Adventurer{}, errors.New("err")).Times(1)
	err := testUsecase.TakeQuest(int64(1), int64(1))
	assert.NotNil(t, err)
}

func TestReportQuestIs(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := repo.NewMockRepository(mockCtrl)
	mockAdvRepo := advRepo.NewMockRepository(mockCtrl)
	testUsecase := &usecase{repo: mockRepo, repoAdv: mockAdvRepo}
	mockRepo.EXPECT().UpdateQuestStatus(model.Quest{ID: 1, Status: 2}).Return(nil).Times(1)
	mockAdvRepo.EXPECT().AddCompletedQuest(int64(1)).Return(nil).Times(1)
	err := testUsecase.ReportQuest(int64(1), int64(1), true)
	assert.Nil(t, err)
}

func TestReportQuestNot(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := repo.NewMockRepository(mockCtrl)
	mockAdvRepo := advRepo.NewMockRepository(mockCtrl)
	testUsecase := &usecase{repo: mockRepo, repoAdv: mockAdvRepo}
	mockRepo.EXPECT().UpdateQuestStatus(model.Quest{ID: 1, Status: 0}).Return(nil).Times(1)
	err := testUsecase.ReportQuest(1, 1, false)
	assert.Nil(t, err)
}
