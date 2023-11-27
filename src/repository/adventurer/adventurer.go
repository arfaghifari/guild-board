package adventurer

import (
	"github.com/arfaghifari/guild-board/src/database"
	model "github.com/arfaghifari/guild-board/src/model/adventurer"
)

func CreateAdventurer(adventurer model.Adventurer) error {
	db := database.GetDB()
	query := `INSERT INTO adventurer(name, rank)
	VALUES($1, $2)`
	createForm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	createForm.Exec(adventurer.Name, adventurer.Rank)
	defer db.Close()
	return err
}

func UpdateAdventurerRank(adventurer model.Adventurer) error {
	db := database.GetDB()
	query := `UPDATE adventurer
	SET rank = $1,
	WHERE id= $2;`
	updateForm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	updateForm.Exec(adventurer.Rank, adventurer.ID)
	defer db.Close()
	return err
}

func GetAdventurer(id int64) (adventurer model.Adventurer, err error) {
	db := database.GetDB()
	query := `SELECT name, rank, completed_quest
	FROM adventurer
	WHERE id = $1`
	rows, err := db.Query(query, id)
	if err != nil {
		return
	}
	err = rows.Scan(&adventurer.Name, &adventurer.Rank, &adventurer.CompletedQuest)
	defer db.Close()
	return
}

func AddCompletedQuest(id int64) error {
	db := database.GetDB()
	query := `UPDATE adventurer
		SET completed_quest = completed_quest + 1
		WHERE id = $1;`
	addForm, err := db.Prepare(query)
	if err != nil {
		return err
	}
	addForm.Exec(id)
	defer db.Close()
	return err
}
