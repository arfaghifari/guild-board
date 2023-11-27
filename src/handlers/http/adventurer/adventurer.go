package adventurer

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "github.com/arfaghifari/guild-board/src/model/adventurer"
	repo "github.com/arfaghifari/guild-board/src/repository/adventurer"
)

func CreateAdventurer(w http.ResponseWriter, r *http.Request) {
	var adventurer model.Adventurer

	if err := json.NewDecoder(r.Body).Decode(&adventurer); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	err := repo.CreateAdventurer(adventurer)

	if err != nil {
		fmt.Fprintf(w, "success")
	}
}

func UpdateAdventurerRank(w http.ResponseWriter, r *http.Request) {
	var adventurer model.Adventurer

	if err := json.NewDecoder(r.Body).Decode(&adventurer); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	err := repo.UpdateAdventurerRank(adventurer)

	if err != nil {
		fmt.Fprintf(w, "success")
	}
}
