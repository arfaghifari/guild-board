package quest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	model "github.com/arfaghifari/guild-board/src/model/quest"
	repo "github.com/arfaghifari/guild-board/src/repository/quest"
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

func GetHello(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello World")
	fmt.Fprintf(w, "HELLO NAKAMA")
}

func GetQuestByStatus(w http.ResponseWriter, r *http.Request) {
	status, _ := strconv.Atoi(r.URL.Query().Get("status"))
	log.Println(r.URL.Query().Get("status"))
	if status == 0 {
		res, err := repo.GetAllAvailableQuest()
		if err != nil {
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

	if status == 1 {
		res, err := repo.GetAllCompletedQuest()
		if err != nil {
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
}

func CreateQuest(w http.ResponseWriter, r *http.Request) {
	var quest model.Quest

	if err := json.NewDecoder(r.Body).Decode(&quest); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	err := repo.CreateQuest(quest)

	if err != nil {
		fmt.Fprintf(w, "success")
	}
}

func DeleteQuest(w http.ResponseWriter, r *http.Request) {
	var quest model.Quest

	if err := json.NewDecoder(r.Body).Decode(&quest); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	err := repo.DeleteQuest(quest)

	if err != nil {
		fmt.Fprintf(w, "success")
	}
}

func UpdateQuestRank(w http.ResponseWriter, r *http.Request) {
	var quest model.Quest

	if err := json.NewDecoder(r.Body).Decode(&quest); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	err := repo.UpdateQuestRank(quest)

	if err != nil {
		fmt.Fprintf(w, "success")
	}
}

func UpdateQuestReward(w http.ResponseWriter, r *http.Request) {
	var quest model.Quest

	if err := json.NewDecoder(r.Body).Decode(&quest); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	err := repo.UpdateQuestReward(quest)

	if err != nil {
		fmt.Fprintf(w, "success")
	}
}

func TakeQuest(w http.ResponseWriter, r *http.Request) {
	var takeByRequest model.TakenBy

	if err := json.NewDecoder(r.Body).Decode(&takeByRequest); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	err := usecase.TakeQuest(takeByRequest.QuestID, takeByRequest.AdventurerID)

	if err != nil {
		fmt.Fprintf(w, "success")
	}
}

func ReportQuest(w http.ResponseWriter, r *http.Request) {
	var reportQuest model.ReportQuest

	if err := json.NewDecoder(r.Body).Decode(&reportQuest); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	err := usecase.ReportQuest(reportQuest.QuestID, reportQuest.AdventurerID, reportQuest.IsCompleted)

	if err != nil {
		fmt.Fprintf(w, "success")
	}
}
