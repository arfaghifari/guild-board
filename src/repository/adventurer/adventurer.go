package adventurer

import (
	"github.com/arfaghifari/guild-board/src/database"
	model "github.com/arfaghifari/guild-board/src/model/adventurer"
)

func CreateAdventurer(adventurer model.Adventurer) error {
	db := database.GetDB()
	query := `INSERT INTO adventurer(name, rank)
	VALUES('$1', '$2')`
	_, err := db.Query(query, adventurer.Name, adventurer.Rank)
	return err
}

func UpdateAdventurerRank(adventurer model.Adventurer) error {
	db := database.GetDB()
	query := `UPDATE adventurer
	SET rank = $1,
	WHERE id= $2;`
	_, err := db.Query(query, adventurer.Rank, adventurer.ID)
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
	return
}

func AddCompletedQuest(id int64) error {
	db := database.GetDB()
	query := `UPDATE adventurer
		SET completed_quest = completed_quest + 1
		WHERE id = $1;`
	_, err := db.Query(query, id)

	return err
}
