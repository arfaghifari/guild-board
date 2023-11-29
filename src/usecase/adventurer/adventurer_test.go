package adventurer

import (
	"testing"

	model "github.com/arfaghifari/guild-board/src/model/adventurer"
	repo "github.com/arfaghifari/guild-board/src/repository/adventurer"
	"github.com/golang/mock/gomock"
)

var adv = model.Adventurer{
	ID:             1,
	Name:           "andi",
	Rank:           11,
	CompletedQuest: 1,
}

func TestCreateAdventurer(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := repo.NewMockRepository(mockCtrl)
	testUsecase := &usecase{repo: mockRepo}
	mockRepo.EXPECT().CreateAdventurer(adv).Return(nil).Times(1)
	testUsecase.CreateAdventurer(adv)
}

func TestUpdateAdventurerRank(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := repo.NewMockRepository(mockCtrl)
	testUsecase := &usecase{repo: mockRepo}
	mockRepo.EXPECT().UpdateAdventurerRank(adv).Return(nil).Times(1)
	testUsecase.UpdateAdventurerRank(adv)
}


