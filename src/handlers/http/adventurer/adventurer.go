package adventurer

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "github.com/arfaghifari/guild-board/src/model/adventurer"
	usecase "github.com/arfaghifari/guild-board/src/usecase/adventurer"
)

type Handlers interface {
	CreateAdventurer(http.ResponseWriter, *http.Request)
	UpdateAdventurerRank(http.ResponseWriter, *http.Request)
}

type handlers struct {
	usecase usecase.Usecase
}

func NewHandlers() (Handlers, error) {
	usecase, _ := usecase.NewUsecase()

	return &handlers{usecase}, nil
}

func (h *handlers) CreateAdventurer(w http.ResponseWriter, r *http.Request) {
	var adventurer model.Adventurer

	if err := json.NewDecoder(r.Body).Decode(&adventurer); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	err := h.usecase.CreateAdventurer(adventurer)

	if err == nil {
		fmt.Fprintf(w, "success")
	}
}

func (h *handlers) UpdateAdventurerRank(w http.ResponseWriter, r *http.Request) {
	var adventurer model.Adventurer

	if err := json.NewDecoder(r.Body).Decode(&adventurer); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	err := h.usecase.UpdateAdventurerRank(adventurer)

	if err == nil {
		fmt.Fprintf(w, "success")
	}
}
