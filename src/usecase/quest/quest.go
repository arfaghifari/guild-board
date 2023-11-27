package quest

import (
	"errors"

	model "github.com/arfaghifari/guild-board/src/model/quest"
	repoAdv "github.com/arfaghifari/guild-board/src/repository/adventurer"
	repo "github.com/arfaghifari/guild-board/src/repository/quest"
)

func TakeQuest(quest_id, adventurer_id int64) error {
	quest, err := repo.GetQuest(quest_id)
	if err != nil {
		return err
	}
	if quest.Status != 0 {
		return errors.New("quest have been taken")
	}
	adv, err := repoAdv.GetAdventurer(adventurer_id)
	if err != nil {
		return err
	}
	if adv.Rank < quest.MinimumRank {
		return errors.New("not capable adventurer rank")
	}
	err = repo.CreateTakenBy(quest_id, adventurer_id)
	if err != nil {
		return err
	}
	quest.Status = 1
	err = repo.UpdateQuestStatus(quest)

	return err
}

func ReportQuest(quest_id, adventurer_id int64) error {
	err := repoAdv.AddCompletedQuest(adventurer_id)
	if err != nil {
		return err
	}
	quest := model.Quest{
		ID:     quest_id,
		Status: 2,
	}
	err = repo.UpdateQuestStatus(quest)
	return err
}
