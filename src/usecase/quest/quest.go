package quest

import (
	"errors"

	model "github.com/arfaghifari/guild-board/src/model/quest"
	repoAdv "github.com/arfaghifari/guild-board/src/repository/adventurer"
	repo "github.com/arfaghifari/guild-board/src/repository/quest"
)

type Usecase interface {
	GetQuestByStatus(int32) ([]model.GetQuestByStatus, error)
	CreateQuest(model.Quest) error
	DeleteQuest(model.Quest) error
	UpdateQuestRank(model.Quest) error
	UpdateQuestReward(model.Quest) error
	TakeQuest(int64, int64) error
	ReportQuest(int64, int64, bool) error
}

type usecase struct {
	repo    repo.Repository
	repoAdv repoAdv.Repository
}

func NewUsecase() (Usecase, error) {
	repo, _ := repo.NewRepository()
	repoAdv, _ := repoAdv.NewRepository()

	return &usecase{repo, repoAdv}, nil
}

func (u *usecase) GetQuestByStatus(status int32) ([]model.GetQuestByStatus, error) {
	if status == 0 {
		return u.repo.GetAllAvailableQuest()
	} else {
		return u.repo.GetAllCompletedQuest()
	}
}

func (u *usecase) CreateQuest(model.Quest) error {
	return u.repo.CreateQuest(model.Quest{})
}

func (u *usecase) DeleteQuest(model.Quest) error {
	return u.repo.DeleteQuest(model.Quest{})
}

func (u *usecase) UpdateQuestReward(model.Quest) error {
	return u.repo.UpdateQuestReward(model.Quest{})
}

func (u *usecase) UpdateQuestRank(model.Quest) error {
	return u.repo.UpdateQuestRank(model.Quest{})
}

func (u *usecase) TakeQuest(quest_id, adventurer_id int64) error {
	quest, err := u.repo.GetQuest(quest_id)
	if err != nil {
		return err
	}
	if quest.Status != 0 {
		return errors.New("quest have been taken")
	}
	adv, err := u.repoAdv.GetAdventurer(adventurer_id)
	if err != nil {
		return err
	}
	if adv.Rank < quest.MinimumRank {
		return errors.New("not capable adventurer rank")
	}
	err = u.repo.CreateTakenBy(quest_id, adventurer_id)
	if err != nil {
		return err
	}
	quest.Status = 1
	err = u.repo.UpdateQuestStatus(quest)

	return err
}

func (u *usecase) ReportQuest(quest_id, adventurer_id int64, is bool) error {
	if !is {
		quest := model.Quest{
			ID:     quest_id,
			Status: 0,
		}
		return u.repo.UpdateQuestStatus(quest)
	} else {
		err := u.repoAdv.AddCompletedQuest(adventurer_id)
		if err != nil {
			return err
		}
		quest := model.Quest{
			ID:     quest_id,
			Status: 2,
		}
		err = u.repo.UpdateQuestStatus(quest)
		return err
	}
}
