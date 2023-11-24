package quest

import (
	"github.com/arfaghifari/guild-board/src/database"
	model "github.com/arfaghifari/guild-board/src/model/quest"
)

func GetAllCompletedQuest() (quests []model.GetQuestByStatus, err error) {
	db := database.GetDB()

	query := `
	SELECT quest_id, name, description, minimum_rank, reward_number
	FROM quest
	WHERE status = 1
	`

	quests = []model.GetQuestByStatus{}
	rows, err := db.Query(query)
	for rows.Next() {
		quest := model.GetQuestByStatus{}
		if err = rows.Scan(&quest.ID, &quest.Name, &quest.Description, &quest.MinimumRank, &quest.RewardNumber); err != nil {
			return
		}
		quests = append(quests, quest)
	}

	return
}

func GetAllAvailableQuest() (quests []model.GetQuestByStatus, err error) {
	db := database.GetDB()

	query := `
	SELECT quest_id, name, description, minimum_rank, reward_number
	FROM quest
	WHERE status = 0
	`

	quests = []model.GetQuestByStatus{}
	rows, err := db.Query(query)
	for rows.Next() {
		quest := model.GetQuestByStatus{}
		if err = rows.Scan(&quest.ID, &quest.Name, &quest.Description, &quest.MinimumRank, &quest.RewardNumber); err != nil {
			return
		}
		quests = append(quests, quest)
	}

	return
}
