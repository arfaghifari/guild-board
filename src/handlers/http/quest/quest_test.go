package quest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	constant "github.com/arfaghifari/guild-board/src/constant"
	modelAdv "github.com/arfaghifari/guild-board/src/model/adventurer"
	model "github.com/arfaghifari/guild-board/src/model/quest"
	usecase "github.com/arfaghifari/guild-board/src/usecase/quest"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var adv = modelAdv.Adventurer{
	ID:             1,
	Name:           "andi",
	Rank:           11,
	CompletedQuest: 1,
}

var bulkQuest = []model.Quest{
	{
		ID:           1,
		Name:         "menyelamatkan kucing",
		Description:  "menyelamatkan kucing yang terjebak di atas pohon",
		MinimumRank:  11,
		RewardNumber: 200000,
		Status:       constant.AvailableQuest,
	},
	{
		ID:           2,
		Name:         "membersihkan selokan",
		Description:  "membersihkan selokan penuh dengan lumut",
		MinimumRank:  12,
		RewardNumber: 200000,
		Status:       constant.AvailableQuest,
	},
	{
		ID:           3,
		Name:         "Supir perjalanan",
		Description:  "Mengantar pulang pergi dan keliling kota, Jakarta-Bandung, Sudah di kasih makan",
		MinimumRank:  13,
		RewardNumber: 600000,
		Status:       constant.CompletedQuest,
	},
}

var bulkQuestByStatus = []model.GetQuestByStatus{
	{
		ID:           1,
		Name:         "menyelamatkan kucing",
		Description:  "menyelamatkan kucing yang terjebak di atas pohon",
		MinimumRank:  11,
		RewardNumber: 200000,
	},
	{
		ID:           2,
		Name:         "membersihkan selokan",
		Description:  "membersihkan selokan penuh dengan lumut",
		MinimumRank:  11,
		RewardNumber: 200000,
	},
	{
		ID:           3,
		Name:         "Supir perjalanan",
		Description:  "Mengantar pulang pergi dan keliling kota, Jakarta-Bandung, Sudah di kasih makan",
		MinimumRank:  13,
		RewardNumber: 600000,
	},
}

func TestGetHello(t *testing.T) {

	ctx := context.Background()
	t.Log("Hello World")
	router := mux.NewRouter()
	router.HandleFunc("/hello", GetHello).Methods(http.MethodGet)
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/hello", strings.NewReader(``))
	request = request.WithContext(ctx)
	router.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code, "error code")
	assert.Equal(t, "HELLO NAKAMA", string(recorder.Body.String()))

}

func TestGetQuestByStatus(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsecase := usecase.NewMockUsecase(mockCtrl)
	testHandlers := &handlers{usecase: mockUsecase}
	mockUsecase.EXPECT().GetQuestByStatus(int32(constant.AvailableQuest)).Return(bulkQuestByStatus[0:1], nil).Times(1)
	ctx := context.Background()
	router := mux.NewRouter()
	router.HandleFunc("/quest-status", testHandlers.GetQuestByStatus).Methods(http.MethodGet)
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/quest-status", strings.NewReader(``))
	values := request.URL.Query()

	values.Add("status", strconv.Itoa(constant.AvailableQuest))
	request.URL.RawQuery = values.Encode()
	request = request.WithContext(ctx)
	router.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code, "error code")
	var bulk GetQuestByStatusResponse
	json.Unmarshal(recorder.Body.Bytes(), &bulk)
	assert.Equal(t, bulkQuestByStatus[0:1], bulk.Data)
}

func TestGetQuestByStatusC(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsecase := usecase.NewMockUsecase(mockCtrl)
	testHandlers := &handlers{usecase: mockUsecase}
	mockUsecase.EXPECT().GetQuestByStatus(int32(constant.CompletedQuest)).Return(bulkQuestByStatus[2:], nil).Times(1)
	ctx := context.Background()
	router := mux.NewRouter()
	router.HandleFunc("/quest-status", testHandlers.GetQuestByStatus).Methods(http.MethodGet)
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/quest-status", strings.NewReader(``))
	values := request.URL.Query()
	values.Add("status", strconv.Itoa(constant.CompletedQuest))
	request.URL.RawQuery = values.Encode()
	request = request.WithContext(ctx)
	router.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code, "error code")
	var resp GetQuestByStatusResponse
	assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &resp))
	assert.Equal(t, bulkQuestByStatus[2:], resp.Data, resp.Header.Error)
}

func TestCreateQuest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsecase := usecase.NewMockUsecase(mockCtrl)
	testHandlers := &handlers{usecase: mockUsecase}
	quest := bulkQuest[0]
	quest.ID = 0
	mockUsecase.EXPECT().CreateQuest(quest).Return(nil).Times(1)
	ctx := context.Background()
	router := mux.NewRouter()
	router.HandleFunc("/quest", testHandlers.CreateQuest).Methods(http.MethodPost)
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/quest", strings.NewReader(`{"name" : "menyelamatkan kucing",  "description" : "menyelamatkan kucing yang terjebak di atas pohon" , "minimum_rank" : 11, "reward_number" : 200000}`))
	request = request.WithContext(ctx)
	router.ServeHTTP(recorder, request)
	var resp MessageResponse
	assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &resp))
	assert.Equal(t, http.StatusOK, recorder.Code, "error code")
	assert.Equal(t, SuccesMessage{true}, resp.Data)
}

func TestUpdateRankQuest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsecase := usecase.NewMockUsecase(mockCtrl)
	testHandlers := &handlers{usecase: mockUsecase}
	quest := model.Quest{
		ID:          1,
		MinimumRank: 12,
	}
	mockUsecase.EXPECT().UpdateQuestRank(quest).Return(nil).Times(1)
	ctx := context.Background()
	router := mux.NewRouter()
	router.HandleFunc("/quest-rank", testHandlers.UpdateQuestRank).Methods(http.MethodPatch)
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/quest-rank", strings.NewReader(`{"quest_id": 1, "minimum_rank" : 12}`))
	request = request.WithContext(ctx)
	router.ServeHTTP(recorder, request)
	var resp MessageResponse
	assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &resp))
	assert.Equal(t, http.StatusOK, recorder.Code, "error code")
	assert.Equal(t, SuccesMessage{true}, resp.Data)
}

func TestUpdateRewardQuest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsecase := usecase.NewMockUsecase(mockCtrl)
	testHandlers := &handlers{usecase: mockUsecase}
	quest := model.Quest{
		ID:           1,
		RewardNumber: 250000,
	}
	mockUsecase.EXPECT().UpdateQuestReward(quest).Return(nil).Times(1)
	ctx := context.Background()
	router := mux.NewRouter()
	router.HandleFunc("/quest-reward", testHandlers.UpdateQuestReward).Methods(http.MethodPatch)
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/quest-reward", strings.NewReader(`{"quest_id": 1, "reward_number" : 250000}`))
	request = request.WithContext(ctx)
	router.ServeHTTP(recorder, request)
	var resp MessageResponse
	assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &resp))
	assert.Equal(t, http.StatusOK, recorder.Code, "error code")
	assert.Equal(t, SuccesMessage{true}, resp.Data)
}

func TestTakeQuest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsecase := usecase.NewMockUsecase(mockCtrl)
	testHandlers := &handlers{usecase: mockUsecase}
	mockUsecase.EXPECT().TakeQuest(bulkQuest[0].ID, adv.ID).Return(nil).Times(1)
	ctx := context.Background()
	router := mux.NewRouter()
	router.HandleFunc("/take-quest", testHandlers.TakeQuest).Methods(http.MethodPost)
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/take-quest", strings.NewReader(`{"quest_id": 1, "adv_id" : 1}`))
	request = request.WithContext(ctx)
	router.ServeHTTP(recorder, request)
	var resp MessageResponse
	assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &resp))
	assert.Equal(t, http.StatusOK, recorder.Code, "error code")
	assert.Equal(t, SuccesMessage{true}, resp.Data)
}

func TestReportQuest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsecase := usecase.NewMockUsecase(mockCtrl)
	testHandlers := &handlers{usecase: mockUsecase}
	mockUsecase.EXPECT().ReportQuest(bulkQuest[0].ID, adv.ID, true).Return(nil).Times(1)
	ctx := context.Background()
	router := mux.NewRouter()
	router.HandleFunc("/report-quest", testHandlers.ReportQuest).Methods(http.MethodPost)
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/report-quest", strings.NewReader(`{"quest_id": 1, "adv_id" : 1, "is_completed" : true}`))
	request = request.WithContext(ctx)
	router.ServeHTTP(recorder, request)
	var resp MessageResponse
	assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &resp))
	assert.Equal(t, http.StatusOK, recorder.Code, "error code")
	assert.Equal(t, SuccesMessage{true}, resp.Data)
}

func TestReportQuestF(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsecase := usecase.NewMockUsecase(mockCtrl)
	testHandlers := &handlers{usecase: mockUsecase}
	mockUsecase.EXPECT().ReportQuest(bulkQuest[0].ID, adv.ID, false).Return(nil).Times(1)
	ctx := context.Background()
	router := mux.NewRouter()
	router.HandleFunc("/report-quest", testHandlers.ReportQuest).Methods(http.MethodPost)
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/report-quest", strings.NewReader(`{"quest_id": 1, "adv_id" : 1, "is_completed" : false}`))
	request = request.WithContext(ctx)
	router.ServeHTTP(recorder, request)
	var resp MessageResponse
	assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &resp))
	assert.Equal(t, http.StatusOK, recorder.Code, "error code")
	assert.Equal(t, SuccesMessage{true}, resp.Data)
}
