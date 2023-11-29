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

type MessageResponse struct {
	Header `json:"header"`
	Data   SuccesMessage `json:"data"`
}

type SuccesMessage struct {
	Success bool `json:"success"`
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
	var (
		statusCode = http.StatusBadRequest
		resp       GetQuestByStatusResponse
	)
	defer func() {
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
	status, err := strconv.Atoi(r.URL.Query().Get("status"))
	if err != nil {
		resp.Header.Error = err.Error()
		return
	}
	if status != 0 && status != 2 {
		resp.Header.Error = "Invalid status number"
		return
	}
	res, err := h.usecase.GetQuestByStatus(int32(status))
	if err != nil {
		statusCode = 500
		resp.Header.Error = err.Error()
		return
	}
	statusCode = 200
	resp.Data = res

}

func (h *handlers) CreateQuest(w http.ResponseWriter, r *http.Request) {
	var (
		statusCode = http.StatusBadRequest
		resp       MessageResponse
		quest      model.Quest
	)
	defer func() {
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

	if err := json.NewDecoder(r.Body).Decode(&quest); err != nil {
		resp.Header.Error = err.Error()
		return
	}
	if quest.Name == "" || quest.MinimumRank <= 0 || quest.RewardNumber <= 0 {
		resp.Header.Error = "name, minimum_rank and reward_number are required and must be valid"
		return
	}

	err := h.usecase.CreateQuest(quest)

	if err != nil {
		statusCode = 500
		resp.Header.Error = err.Error()
		return
	}
	statusCode = 200
	resp.Data.Success = true

}

func (h *handlers) DeleteQuest(w http.ResponseWriter, r *http.Request) {
	var (
		statusCode = http.StatusBadRequest
		resp       MessageResponse
		quest      model.Quest
	)
	defer func() {
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

	if err := json.NewDecoder(r.Body).Decode(&quest); err != nil {
		resp.Header.Error = err.Error()
		return
	}

	if quest.ID <= 0 {
		resp.Header.Error = "quest_id is required and must be valid"
		return
	}

	err := h.usecase.DeleteQuest(quest)

	if err != nil {
		statusCode = 500
		resp.Header.Error = err.Error()
		return
	}
	statusCode = 200
	resp.Data.Success = true
}

func (h *handlers) UpdateQuestRank(w http.ResponseWriter, r *http.Request) {
	var (
		statusCode = http.StatusBadRequest
		resp       MessageResponse
		quest      model.Quest
	)
	defer func() {
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

	if err := json.NewDecoder(r.Body).Decode(&quest); err != nil {
		resp.Header.Error = err.Error()
		return
	}

	if quest.ID <= 0 || quest.MinimumRank <= 0 {
		resp.Header.Error = "quest_id and minimum_rank are required and must be valid"
		return
	}

	err := h.usecase.UpdateQuestRank(quest)

	if err != nil {
		statusCode = 500
		resp.Header.Error = err.Error()
		return
	}
	statusCode = 200
	resp.Data.Success = true
}

func (h *handlers) UpdateQuestReward(w http.ResponseWriter, r *http.Request) {
	var (
		statusCode = http.StatusBadRequest
		resp       MessageResponse
		quest      model.Quest
	)
	defer func() {
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

	if err := json.NewDecoder(r.Body).Decode(&quest); err != nil {
		resp.Header.Error = err.Error()
		return
	}

	if quest.ID <= 0 || quest.RewardNumber <= 0 {
		resp.Header.Error = "quest_id and reward_number are required and must be valid"
		return
	}

	err := h.usecase.UpdateQuestReward(quest)

	if err != nil {
		statusCode = 500
		resp.Header.Error = err.Error()
		return
	}
	statusCode = 200
	resp.Data.Success = true
}

func (h *handlers) TakeQuest(w http.ResponseWriter, r *http.Request) {
	var (
		statusCode    = http.StatusBadRequest
		resp          MessageResponse
		takeByRequest model.TakenBy
	)
	defer func() {
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

	if err := json.NewDecoder(r.Body).Decode(&takeByRequest); err != nil {
		resp.Header.Error = err.Error()
		return
	}

	if takeByRequest.AdventurerID <= 0 || takeByRequest.QuestID <= 0 {
		resp.Header.Error = "adv_id and quest_id are required and must be valid"
		return
	}

	err := h.usecase.TakeQuest(takeByRequest.QuestID, takeByRequest.AdventurerID)

	if err != nil {
		statusCode = 500
		resp.Header.Error = err.Error()
		return
	}
	statusCode = 200
	resp.Data.Success = true
}

func (h *handlers) ReportQuest(w http.ResponseWriter, r *http.Request) {
	var (
		statusCode  = http.StatusBadRequest
		resp        MessageResponse
		reportQuest model.ReportQuest
	)
	defer func() {
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

	if err := json.NewDecoder(r.Body).Decode(&reportQuest); err != nil {
		resp.Header.Error = err.Error()
		return
	}

	if reportQuest.AdventurerID <= 0 || reportQuest.QuestID <= 0 {
		resp.Header.Error = "adv_id and quest_id are required and must be valid"
		return
	}

	err := h.usecase.ReportQuest(reportQuest.QuestID, reportQuest.AdventurerID, reportQuest.IsCompleted)

	if err != nil {
		statusCode = 500
		resp.Header.Error = err.Error()
		return
	}
	statusCode = 200
	resp.Data.Success = true
}
