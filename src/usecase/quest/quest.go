package quest

import (
	"errors"

	constant "github.com/arfaghifari/guild-board/src/constant"
	model "github.com/arfaghifari/guild-board/src/model/quest"
	repoAdv "github.com/arfaghifari/guild-board/src/repository/adventurer"
	repo "github.com/arfaghifari/guild-board/src/repository/quest"
)

type Usecase interface {
	GetQuestByStatus(int32) ([]model.GetQuestByStatus, error)
	CreateQuest(model.Quest) (model.Quest, error)
	DeleteQuest(model.Quest) error
	UpdateQuestRank(model.Quest) error
	UpdateQuestReward(model.Quest) error
	TakeQuest(int64, int64) error
	ReportQuest(int64, int64, bool) error
	GetQuestActiveAdventurer(int64) ([]model.Quest, error)
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
	if status == constant.AvailableQuest {
		return u.repo.GetAllAvailableQuest()
	} else {
		return u.repo.GetAllCompletedQuest()
	}
}

func (u *usecase) CreateQuest(quest model.Quest) (model.Quest, error) {
	return u.repo.CreateQuest(quest)
}

func (u *usecase) DeleteQuest(quest model.Quest) error {
	return u.repo.DeleteQuest(quest)
}

func (u *usecase) UpdateQuestReward(quest model.Quest) error {
	return u.repo.UpdateQuestReward(quest)
}

func (u *usecase) UpdateQuestRank(quest model.Quest) error {
	return u.repo.UpdateQuestRank(quest)
}

func (u *usecase) TakeQuest(quest_id, adventurer_id int64) error {
	quest, err := u.repo.GetQuest(quest_id)
	if err != nil {
		return err
	}
	if quest.Status != constant.AvailableQuest {
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
	quest.Status = constant.WorkingQuest
	err = u.repo.UpdateQuestStatus(quest)

	return err
}

func (u *usecase) ReportQuest(quest_id, adventurer_id int64, is bool) error {
	if err := u.repo.IsExistTakenBy(quest_id, adventurer_id); err != nil {
		return err
	}
	quest, err := u.repo.GetQuest(quest_id)
	if err != nil {
		return err
	}
	if quest.Status != constant.WorkingQuest {
		return errors.New("quest have not taken")
	}
	if !is {
		quest := model.Quest{
			ID:     quest_id,
			Status: constant.AvailableQuest,
		}
		return u.repo.UpdateQuestStatus(quest)
	} else {
		err := u.repoAdv.AddCompletedQuest(adventurer_id)
		if err != nil {
			return err
		}
		quest := model.Quest{
			ID:     quest_id,
			Status: constant.CompletedQuest,
		}
		err = u.repo.UpdateQuestStatus(quest)
		return err
	}
}

func (u *usecase) GetQuestActiveAdventurer(adv_id int64) ([]model.Quest, error) {
	return u.repo.GetQuestActiveAdventurer(adv_id)
}
