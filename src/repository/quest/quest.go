package quest

import (
	"database/sql"

	"github.com/arfaghifari/guild-board/src/database"
	model "github.com/arfaghifari/guild-board/src/model/quest"
)

type Repository interface {
	Close()
	GetAllCompletedQuest() ([]model.GetQuestByStatus, error)
	GetAllAvailableQuest() ([]model.GetQuestByStatus, error)
	CreateQuest(model.Quest) error
	UpdateQuestRank(model.Quest) error
	UpdateQuestStatus(model.Quest) error
	UpdateQuestReward(model.Quest) error
	DeleteQuest(model.Quest) error
	GetQuest(int64) (model.Quest, error)
	CreateTakenBy(int64, int64) error
}

type repository struct {
	db *sql.DB
}

func NewRepository() (Repository, error) {
	db := database.GetDB()

	return &repository{db}, nil
}

func (r *repository) Close() {
	r.db.Close()
}

func (r *repository) GetAllCompletedQuest() (quests []model.GetQuestByStatus, err error) {

	db := r.db

	query := `
	SELECT quest_id, name, description, minimum_rank, reward_number
	FROM quest
	WHERE status = $1
	`
	completedStatus := 2
	quests = []model.GetQuestByStatus{}
	rows, err := db.Query(query, completedStatus)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		quest := model.GetQuestByStatus{}
		if err = rows.Scan(&quest.ID, &quest.Name, &quest.Description, &quest.MinimumRank, &quest.RewardNumber); err != nil {
			return
		}
		quests = append(quests, quest)
	}

	return
}

func (r *repository) GetAllAvailableQuest() (quests []model.GetQuestByStatus, err error) {
	db := r.db

	query := `
	SELECT quest_id, name, description, minimum_rank, reward_number
	FROM quest
	WHERE status = $1
	`
	availableStatus := 0
	quests = []model.GetQuestByStatus{}
	rows, err := db.Query(query, availableStatus)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		quest := model.GetQuestByStatus{}
		if err = rows.Scan(&quest.ID, &quest.Name, &quest.Description, &quest.MinimumRank, &quest.RewardNumber); err != nil {
			return
		}
		quests = append(quests, quest)
	}

	return
}

func (r *repository) CreateQuest(quest model.Quest) error {
	db := r.db
	query := `INSERT INTO quest(name, description, minimum_rank, reward_number)
	VALUES($1, $2, $3, $4)`
	createForm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	createForm.Exec(quest.Name, quest.Description, quest.MinimumRank, quest.RewardNumber)
	defer createForm.Close()
	return err
}

func (r *repository) UpdateQuestRank(quest model.Quest) error {
	db := r.db
	query := `UPDATE quest
	SET minimum_rank = $1
	WHERE quest_id = $2`
	updateForm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	updateForm.Exec(quest.MinimumRank, quest.ID)
	defer updateForm.Close()
	return err
}

func (r *repository) UpdateQuestReward(quest model.Quest) error {
	db := r.db
	query := `UPDATE quest
	SET reward_number = $1
	WHERE quest_id = $2`
	updateForm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	updateForm.Exec(quest.RewardNumber, quest.ID)
	defer updateForm.Close()
	return err
}

func (r *repository) UpdateQuestStatus(quest model.Quest) error {
	db := r.db
	query := `UPDATE quest
	SET status = $1
	WHERE quest_id = $2`
	updateForm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	updateForm.Exec(quest.Status, quest.ID)
	defer updateForm.Close()
	return err
}

func (r *repository) DeleteQuest(quest model.Quest) error {
	db := r.db
	query := `DELETE FROM quest
	WHERE quest_id = $1`
	deleteForm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	deleteForm.Exec(quest.ID)
	defer deleteForm.Close()
	return err
}

func (r *repository) GetQuest(id int64) (quest model.Quest, err error) {
	db := r.db
	query := `SELECT name, description, minimum_rank, reward_number, status 
	FROM quest
	WHERE quest_id = $1`
	quest.ID = id
	err = db.QueryRow(query, id).Scan(&quest.Name, &quest.Description, &quest.MinimumRank, &quest.RewardNumber, &quest.Status)
	return
}

func (r *repository) CreateTakenBy(quest_id, adventurer_id int64) error {
	db := r.db
	query := `INSERT INTO taken_by(quest_id, adv_id)
	VALUES($1, $2)`
	createForm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	createForm.Exec(quest_id, adventurer_id)
	defer createForm.Close()
	return err
}
