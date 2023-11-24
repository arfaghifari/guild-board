package quest

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "github.com/arfaghifari/guild-board/src/model/quest"
	repo "github.com/arfaghifari/guild-board/src/repository/quest"
)

type Header struct {
	Error      string `json:"error_code"`
	StatusCode int    `json:"status_code"`
}

type GetQuestByStatusResponse struct {
	Header `json:"header"`
	Data   []model.GetQuestByStatus `json:"data"`
}

func GetQuestByStatus(w http.ResponseWriter, r *http.Request) {
	status, _ := strconv.Atoi(r.URL.Query().Get("status"))
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
}
