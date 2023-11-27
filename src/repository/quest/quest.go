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
	WHERE status = 2
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
	defer db.Close()
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
	defer db.Close()
	return
}

func CreateQuest(quest model.Quest) error {
	db := database.GetDB()
	query := `INSERT INTO quest(name, description, minimum_rank, reward_number)
	VALUES($1, $2, $3, $4)`
	createForm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	createForm.Exec(quest.Name, quest.Description, quest.MinimumRank, quest.RewardNumber)
	defer db.Close()
	return err
}

func UpdateQuestRank(quest model.Quest) error {
	db := database.GetDB()
	query := `UPDATE quest
	SET minimum_rank = $1,
	WHERE quest_id = $2;`
	updateForm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	updateForm.Exec(quest.MinimumRank, quest.ID)
	defer db.Close()
	return err
}

func UpdateQuestReward(quest model.Quest) error {
	db := database.GetDB()
	query := `UPDATE quest
	SET reward_number = $1,
	WHERE quest_id = $2;`
	updateForm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	updateForm.Exec(quest.RewardNumber, quest.ID)
	defer db.Close()
	return err
}

func UpdateQuestStatus(quest model.Quest) error {
	db := database.GetDB()
	query := `UPDATE quest
	SET status = $1,
	WHERE quest_id = $2;`
	updateForm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	updateForm.Exec(quest.Status, quest.ID)
	defer db.Close()
	return err
}

func DeleteQuest(quest model.Quest) error {
	db := database.GetDB()
	query := `DELETE FROM quest
	WHERE quest_id = $1`
	deleteForm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	deleteForm.Exec(quest.ID)
	defer db.Close()
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
	defer db.Close()
	return
}

func CreateTakenBy(quest_id, adventurer_id int64) error {
	db := database.GetDB()
	query := `INSERT INTO taken_by(quest_id, adv_id)
	VALUES($1, $2)`
	createForm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	createForm.Exec(quest_id, adventurer_id)
	defer db.Close()
	return err
}
