package adventurer

import (
	"database/sql"

	"github.com/arfaghifari/guild-board/src/database"
	model "github.com/arfaghifari/guild-board/src/model/adventurer"
)

type Repository interface {
	Close()
	CreateAdventurer(model.Adventurer) error
	UpdateAdventurerRank(model.Adventurer) error
	GetAdventurer(int64) (model.Adventurer, error)
	AddCompletedQuest(int64) error
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

func (r *repository) CreateAdventurer(adventurer model.Adventurer) error {
	db := r.db
	query := `INSERT INTO adventurer(name, rank)
	VALUES($1, $2)`
	createForm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	createForm.Exec(adventurer.Name, adventurer.Rank)
	defer createForm.Close()
	return err
}

func (r *repository) UpdateAdventurerRank(adventurer model.Adventurer) error {
	db := r.db
	query := `UPDATE adventurer
	SET rank = $1
	WHERE id = $2`
	updateForm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	updateForm.Exec(adventurer.Rank, adventurer.ID)
	defer updateForm.Close()
	return err
}

func (r *repository) GetAdventurer(id int64) (adventurer model.Adventurer, err error) {
	db := r.db
	query := `SELECT name, rank, completed_quest
	FROM adventurer
	WHERE id = $1`
	adventurer.ID = id
	err = db.QueryRow(query, id).Scan(&adventurer.Name, &adventurer.Rank, &adventurer.CompletedQuest)
	return
}

func (r *repository) AddCompletedQuest(id int64) error {
	db := r.db
	query := `UPDATE adventurer
		SET completed_quest = completed_quest + 1
		WHERE id = $1`
	addForm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	addForm.Exec(id)
	defer addForm.Close()
	return err
}
