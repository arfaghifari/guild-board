package quest

import (
	"database/sql"

	constant "github.com/arfaghifari/guild-board/src/constant"
	"github.com/arfaghifari/guild-board/src/database"
	model "github.com/arfaghifari/guild-board/src/model/quest"
)

type Repository interface {
	Close()
	GetAllCompletedQuest() ([]model.GetQuestByStatus, error)
	GetAllAvailableQuest() ([]model.GetQuestByStatus, error)
	CreateQuest(model.Quest) (model.Quest, error)
	UpdateQuestRank(model.Quest) error
	UpdateQuestStatus(model.Quest) error
	UpdateQuestReward(model.Quest) error
	DeleteQuest(model.Quest) error
	GetQuest(int64) (model.Quest, error)
	CreateTakenBy(int64, int64) error
	IsExistTakenBy(int64, int64) error
	GetQuestActiveAdventurer(int64) ([]model.Quest, error)
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
	completedStatus := constant.CompletedQuest
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
	availableStatus := constant.AvailableQuest
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

func (r *repository) CreateQuest(quest model.Quest) (qst model.Quest, err error) {
	db := r.db
	query := `INSERT INTO quest(name, description, minimum_rank, reward_number)
	VALUES($1, $2, $3, $4) RETURNING quest_id`
	createForm, err := db.Prepare(query)
	qst = quest
	if err != nil {
		return model.Quest{}, err
	}
	err = createForm.QueryRow(quest.Name, quest.Description, quest.MinimumRank, quest.RewardNumber).Scan(&qst.ID)
	if err != nil {
		return model.Quest{}, err
	}
	qst.Status = 0
	defer createForm.Close()
	return
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
	_, err = deleteForm.Exec(quest.ID)
	// res.LastInsertId()
	// res.RowsAffected()
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

func (r *repository) IsExistTakenBy(quest_id, adv_id int64) error {
	var one int
	db := r.db
	query := `SELECT 1
	FROM taken_by
	WHERE quest_id = $1 AND adv_id = $2`
	return db.QueryRow(query, quest_id, adv_id).Scan(&one)

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

func (r *repository) GetQuestActiveAdventurer(id int64) (quests []model.Quest, err error) {
	db := r.db

	query := `
	SELECT quest_id, name, description, minimum_rank, reward_number, status
	FROM quest NATURAL JOIN taken_by
	WHERE status = $1 AND adv_id = $2
	`
	quests = []model.Quest{}
	rows, err := db.Query(query, constant.WorkingQuest, id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		quest := model.Quest{}
		if err = rows.Scan(&quest.ID, &quest.Name, &quest.Description, &quest.MinimumRank, &quest.RewardNumber, &quest.Status); err != nil {
			return
		}
		quests = append(quests, quest)
	}

	return
}
