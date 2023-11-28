package quest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	model "github.com/arfaghifari/guild-board/src/model/quest"
	usecase "github.com/arfaghifari/guild-board/src/usecase/quest"
)

type Header struct {
	Error      string `json:"error_code"`
	StatusCode int    `json:"status_code"`
}

type GetQuestByStatusResponse struct {
	Header `json:"header"`
	Data   []model.GetQuestByStatus `json:"data"`
}

type Handlers interface {
	GetQuestByStatus(http.ResponseWriter, *http.Request)
	CreateQuest(http.ResponseWriter, *http.Request)
	DeleteQuest(http.ResponseWriter, *http.Request)
	UpdateQuestRank(http.ResponseWriter, *http.Request)
	UpdateQuestReward(http.ResponseWriter, *http.Request)
	TakeQuest(http.ResponseWriter, *http.Request)
	ReportQuest(http.ResponseWriter, *http.Request)
}

type handlers struct {
	usecase usecase.Usecase
}

func NewHandlers() (Handlers, error) {
	usecase, _ := usecase.NewUsecase()

	return &handlers{usecase}, nil
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello World")
	fmt.Fprintf(w, "HELLO NAKAMA")
}

func (h *handlers) GetQuestByStatus(w http.ResponseWriter, r *http.Request) {
	status, err := strconv.Atoi(r.URL.Query().Get("status"))
	if err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	res, err := h.usecase.GetQuestByStatus(int32(status))
	if err != nil {
		http.Error(w, "internal server error", 500)
		return
	}
	resp, _ := json.Marshal(GetQuestByStatusResponse{
		Header: Header{
			Error:      "",
			StatusCode: 200,
		},
		Data: res,
	})

	w.Write(resp)

}

func (h *handlers) CreateQuest(w http.ResponseWriter, r *http.Request) {
	var quest model.Quest

	if err := json.NewDecoder(r.Body).Decode(&quest); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	err := h.usecase.CreateQuest(quest)

	if err != nil {
		http.Error(w, "internal server error", 500)
		return
	}

	fmt.Fprintf(w, "success")

}

func (h *handlers) DeleteQuest(w http.ResponseWriter, r *http.Request) {
	var quest model.Quest

	if err := json.NewDecoder(r.Body).Decode(&quest); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	err := h.usecase.DeleteQuest(quest)

	if err != nil {
		http.Error(w, "internal server error", 500)
		return
	}

	fmt.Fprintf(w, "success")
}

func (h *handlers) UpdateQuestRank(w http.ResponseWriter, r *http.Request) {
	var quest model.Quest

	if err := json.NewDecoder(r.Body).Decode(&quest); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	err := h.usecase.UpdateQuestRank(quest)

	if err != nil {
		http.Error(w, "internal server error", 500)
		return
	}

	fmt.Fprintf(w, "success")
}

func (h *handlers) UpdateQuestReward(w http.ResponseWriter, r *http.Request) {
	var quest model.Quest

	if err := json.NewDecoder(r.Body).Decode(&quest); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	err := h.usecase.UpdateQuestReward(quest)

	if err != nil {
		http.Error(w, "internal server error", 500)
		return
	}

	fmt.Fprintf(w, "success")
}

func (h *handlers) TakeQuest(w http.ResponseWriter, r *http.Request) {
	var takeByRequest model.TakenBy

	if err := json.NewDecoder(r.Body).Decode(&takeByRequest); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	err := h.usecase.TakeQuest(takeByRequest.QuestID, takeByRequest.AdventurerID)

	if err != nil {
		http.Error(w, "internal server error", 500)
		return
	}

	fmt.Fprintf(w, "success")
}

func (h *handlers) ReportQuest(w http.ResponseWriter, r *http.Request) {
	var reportQuest model.ReportQuest

	if err := json.NewDecoder(r.Body).Decode(&reportQuest); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	err := h.usecase.ReportQuest(reportQuest.QuestID, reportQuest.AdventurerID, reportQuest.IsCompleted)

	if err != nil {
		http.Error(w, "internal server error", 500)
		return
	}

	fmt.Fprintf(w, "success")
}
