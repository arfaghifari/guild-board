package adventurer

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	model "github.com/arfaghifari/guild-board/src/model/adventurer"
	usecase "github.com/arfaghifari/guild-board/src/usecase/adventurer"
)

type Header struct {
	Error      string `json:"error_code"`
	StatusCode int    `json:"status_code"`
}

type MessageResponse struct {
	Header `json:"header"`
	Data   SuccesMessage `json:"data"`
}

type AdvResponse struct {
	Header `json:"header"`
	Data   model.Adventurer `json:"data"`
}

type SuccesMessage struct {
	Success bool `json:"success"`
}

type Handlers interface {
	CreateAdventurer(http.ResponseWriter, *http.Request)
	UpdateAdventurerRank(http.ResponseWriter, *http.Request)
	GetAdventurer(http.ResponseWriter, *http.Request)
}

type handlers struct {
	usecase usecase.Usecase
}

func NewHandlers() (Handlers, error) {
	usecase, _ := usecase.NewUsecase()

	return &handlers{usecase}, nil
}

func (h *handlers) CreateAdventurer(w http.ResponseWriter, r *http.Request) {
	var (
		statusCode = http.StatusBadRequest
		resp       AdvResponse
		adventurer model.Adventurer
	)
	defer func() {
		w.Header().Set("Content-Type", "application/json")
		resp.StatusCode = statusCode
		responseWriter, err := json.Marshal(resp)
		if err != nil {
			log.Fatal("Failed build response")
		}
		if statusCode == http.StatusOK {
			w.Write(responseWriter)
		} else {
			http.Error(w, string(responseWriter), statusCode)
		}
	}()
	resp.Data = model.Adventurer{}
	if err := json.NewDecoder(r.Body).Decode(&adventurer); err != nil {
		resp.Header.Error = err.Error()
		return
	}
	if adventurer.Name == "" || adventurer.Rank <= 0 {
		resp.Header.Error = "name and rank are required and must be valid"
		return
	}

	res, err := h.usecase.CreateAdventurer(adventurer)
	if err != nil {
		statusCode = http.StatusInternalServerError
		resp.Header.Error = err.Error()
		return
	}
	statusCode = http.StatusOK
	resp.Data = res
}

func (h *handlers) UpdateAdventurerRank(w http.ResponseWriter, r *http.Request) {
	var (
		statusCode = http.StatusBadRequest
		resp       MessageResponse
		adventurer model.Adventurer
	)
	defer func() {
		w.Header().Set("Content-Type", "application/json")
		resp.StatusCode = statusCode
		responseWriter, err := json.Marshal(resp)
		if err != nil {
			log.Fatal("Failed build response")
		}
		if statusCode == http.StatusOK {
			w.Write(responseWriter)
		} else {
			http.Error(w, string(responseWriter), statusCode)
		}
	}()

	if err := json.NewDecoder(r.Body).Decode(&adventurer); err != nil {
		resp.Header.Error = err.Error()
		return
	}

	if adventurer.ID <= 0 || adventurer.Rank <= 0 {
		resp.Header.Error = "id and rank are required and must be valid"
		return
	}

	err := h.usecase.UpdateAdventurerRank(adventurer)

	if err != nil {
		statusCode = http.StatusInternalServerError
		resp.Header.Error = err.Error()
		return
	}
	statusCode = http.StatusOK
	resp.Data.Success = true
}

func (h *handlers) GetAdventurer(w http.ResponseWriter, r *http.Request) {
	var (
		statusCode = http.StatusBadRequest
		resp       AdvResponse
	)
	defer func() {
		w.Header().Set("Content-Type", "application/json")
		resp.StatusCode = statusCode
		responseWriter, err := json.Marshal(resp)
		if err != nil {
			log.Fatal("Failed build response")
		}
		if statusCode == http.StatusOK {
			w.Write(responseWriter)
		} else {
			http.Error(w, string(responseWriter), statusCode)
		}
	}()
	resp.Data = model.Adventurer{}
	adv_id, err := strconv.Atoi(r.URL.Query().Get("adv_id"))
	if err != nil {
		resp.Header.Error = err.Error()
		return
	}
	if adv_id <= 0 {
		resp.Header.Error = "Invalid id"
		return
	}

	res, err := h.usecase.GetAdventurer(int64(adv_id))
	if err != nil {
		statusCode = http.StatusInternalServerError
		resp.Header.Error = err.Error()
		return
	}
	statusCode = http.StatusOK
	resp.Data = res
}
