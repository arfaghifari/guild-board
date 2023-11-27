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

func CreateQuest(quest model.Quest) error {
	db := database.GetDB()
	query := `INSERT INTO quest(name, description, minimum_rank, reward_number)
	VALUES('$1', '$2', '$3', $4)`
	_, err := db.Query(query, quest.Name, quest.Description, quest.MinimumRank, quest.RewardNumber)
	return err
}

func UpdateQuestRank(quest model.Quest) error {
	db := database.GetDB()
	query := `UPDATE quest
	SET minimum_rank = $1,
	WHERE quest_id = $2;`
	_, err := db.Query(query, quest.MinimumRank, quest.ID)
	return err
}

func UpdateQuestReward(quest model.Quest) error {
	db := database.GetDB()
	query := `UPDATE quest
	SET reward_number = $1,
	WHERE quest_id = $2;`
	_, err := db.Query(query, quest.RewardNumber, quest.ID)
	return err
}

func UpdateQuestStatus(quest model.Quest) error {
	db := database.GetDB()
	query := `UPDATE quest
	SET status = $1,
	WHERE quest_id = $2;`
	_, err := db.Query(query, quest.Status, quest.ID)
	return err
}

func DeleteQuest(quest model.Quest) error {
	db := database.GetDB()
	query := `DELETE FROM quest
	WHERE quest_id = $1`
	_, err := db.Query(query, quest.ID)
	return err
}

func GetQuest(id int64) (quest model.Quest, err error) {
	db := database.GetDB()
	query := `SELECT name, description, minimum_rank, reward_number, status 
	FROM quest
	WHERE quest_id = $1`
	rows, err := db.Query(query, id)
	if err != nil {
		return
	}
	err = rows.Scan(&quest.Name, &quest.Description, &quest.MinimumRank, &quest.RewardNumber, &quest.Status)
	return
}

func CreateTakenBy(quest_id, adventurer_id int64) error {
	db := database.GetDB()
	query := `INSERT INTO taken_by(quest_id, adv_id)
	VALUES('$1', '$2')`
	_, err := db.Query(query, quest_id, adventurer_id)
	return err
}
