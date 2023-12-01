package quest

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	constant "github.com/arfaghifari/guild-board/src/constant"
	modelAdv "github.com/arfaghifari/guild-board/src/model/adventurer"
	model "github.com/arfaghifari/guild-board/src/model/quest"
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

func TestNewHandlers(t *testing.T) {
	res, err := NewHandlers()
	assert.NoError(t, err)
	assert.NotNil(t, res)
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

	type fields struct {
		u *MockUsecase
	}
	type requests struct {
		is          bool
		statusQuery string
	}
	type responses struct {
		body []model.GetQuestByStatus
	}
	tests := []struct {
		name           string
		fields         fields
		req            requests
		resp           responses
		mock           func(*MockUsecase)
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "success get available quest",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				is:          true,
				statusQuery: "0",
			},
			resp: responses{
				body: bulkQuestByStatus[0:1],
			},
			mock: func(usecase *MockUsecase) {
				usecase.EXPECT().GetQuestByStatus(int32(constant.AvailableQuest)).Return(bulkQuestByStatus[0:1], nil).Times(1)
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "failed get available quest usecase",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				is:          true,
				statusQuery: "0",
			},
			resp: responses{
				body: []model.GetQuestByStatus{},
			},
			mock: func(usecase *MockUsecase) {
				usecase.EXPECT().GetQuestByStatus(int32(constant.AvailableQuest)).Return([]model.GetQuestByStatus{}, errors.New("any error")).Times(1)
			},
			wantStatusCode: http.StatusInternalServerError,
			wantErr:        true,
		},
		{
			name: "success get completed quest",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				is:          true,
				statusQuery: "2",
			},
			resp: responses{
				body: bulkQuestByStatus[2:],
			},
			mock: func(usecase *MockUsecase) {
				usecase.EXPECT().GetQuestByStatus(int32(constant.CompletedQuest)).Return(bulkQuestByStatus[2:], nil).Times(1)
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "query empty",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				is:          false,
				statusQuery: "",
			},
			resp: responses{
				body: []model.GetQuestByStatus{},
			},
			mock: func(usecase *MockUsecase) {
			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "query not int",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				is:          true,
				statusQuery: "a",
			},
			resp: responses{
				body: []model.GetQuestByStatus{},
			},
			mock: func(usecase *MockUsecase) {
			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "query not valid",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				is:          true,
				statusQuery: "1",
			},
			resp: responses{
				body: []model.GetQuestByStatus{},
			},
			mock: func(usecase *MockUsecase) {
			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			router := mux.NewRouter()
			h := &handlers{
				usecase: tt.fields.u,
			}
			router.HandleFunc("/quest-status", h.GetQuestByStatus).Methods(http.MethodGet)
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/quest-status", strings.NewReader(``))
			if tt.req.is {
				values := request.URL.Query()
				values.Add("status", tt.req.statusQuery)
				request.URL.RawQuery = values.Encode()
			}
			request = request.WithContext(ctx)
			tt.mock(tt.fields.u)
			router.ServeHTTP(recorder, request)
			var resp GetQuestByStatusResponse
			assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &resp))
			assert.Equal(t, tt.wantStatusCode, recorder.Code, "error code")
			assert.Equal(t, tt.resp.body, resp.Data)
			if tt.wantErr {
				assert.NotEqual(t, "", resp.Header.Error, "error message")
			} else {
				assert.Equal(t, "", resp.Header.Error, "error message")
			}
		})
	}
}

func TestCreateQuest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		u *MockUsecase
	}
	type requests struct {
		body string
	}
	type responses struct {
		body model.Quest
	}
	tests := []struct {
		name           string
		fields         fields
		req            requests
		resp           responses
		mock           func(*MockUsecase)
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "success created a quest",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"name" : "menyelamatkan kucing",  "description" : "menyelamatkan kucing yang terjebak di atas pohon" , "minimum_rank" : 11, "reward_number" : 200000}`,
			},
			resp: responses{
				body: bulkQuest[0],
			},
			mock: func(usecase *MockUsecase) {
				quest := bulkQuest[0]
				quest.ID = 0
				usecase.EXPECT().CreateQuest(quest).Return(bulkQuest[0], nil).Times(1)
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "json failed",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{`,
			},
			resp: responses{
				body: model.Quest{},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "empty name",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"description" : "menyelamatkan kucing yang terjebak di atas pohon" , "minimum_rank" : 11, "reward_number" : 200000}`,
			},
			resp: responses{
				body: model.Quest{},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "invalid name",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"name" : "",  "description" : "menyelamatkan kucing yang terjebak di atas pohon" , "minimum_rank" : 11, "reward_number" : 200000}`,
			},
			resp: responses{
				body: model.Quest{},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "empty rank",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"name" : "menyelamatkan kucing",  "description" : "menyelamatkan kucing yang terjebak di atas pohon" ,  "reward_number" : 200000}`,
			},
			resp: responses{
				body: model.Quest{},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "invalid rank",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"name" : "menyelamatkan kucing",  "description" : "menyelamatkan kucing yang terjebak di atas pohon" , "minimum_rank" : -1, "reward_number" : 200000}`,
			},
			resp: responses{
				body: model.Quest{},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "empty reward",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"name" : "menyelamatkan kucing",  "description" : "menyelamatkan kucing yang terjebak di atas pohon" , "minimum_rank" : 11}`,
			},
			resp: responses{
				body: model.Quest{},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "invalid reward",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"name" : "menyelamatkan kucing",  "description" : "menyelamatkan kucing yang terjebak di atas pohon" , "minimum_rank" : 11, "reward_number" : -1}`,
			},
			resp: responses{
				body: model.Quest{},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "error at layer usecase",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"name" : "menyelamatkan kucing",  "description" : "menyelamatkan kucing yang terjebak di atas pohon" , "minimum_rank" : 11, "reward_number" : 200000}`,
			},
			resp: responses{
				body: model.Quest{},
			},
			mock: func(usecase *MockUsecase) {
				quest := bulkQuest[0]
				quest.ID = 0
				usecase.EXPECT().CreateQuest(quest).Return(model.Quest{}, errors.New("any error")).Times(1)
			},
			wantStatusCode: http.StatusInternalServerError,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			router := mux.NewRouter()
			h := &handlers{
				usecase: tt.fields.u,
			}
			router.HandleFunc("/quest", h.CreateQuest).Methods(http.MethodPost)
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("POST", "/quest", strings.NewReader(tt.req.body))
			request = request.WithContext(ctx)
			tt.mock(tt.fields.u)
			router.ServeHTTP(recorder, request)
			var resp QuestResponse
			assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &resp))
			assert.Equal(t, tt.wantStatusCode, recorder.Code, "error code")
			assert.Equal(t, tt.resp.body, resp.Data)
			if tt.wantErr {
				assert.NotEqual(t, "", resp.Header.Error, "error message")
			} else {
				assert.Equal(t, "", resp.Header.Error, "error message")
			}
		})
	}
}

func TestDeleteQuest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		u *MockUsecase
	}
	type requests struct {
		body string
	}
	type responses struct {
		body SuccesMessage
	}
	quest := model.Quest{
		ID: 1,
	}
	tests := []struct {
		name           string
		fields         fields
		req            requests
		resp           responses
		mock           func(*MockUsecase)
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "success delete quest",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id" : 1 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: true},
			},
			mock: func(usecase *MockUsecase) {
				usecase.EXPECT().DeleteQuest(quest).Return(nil).Times(1)
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "json failed",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "empty id",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{ }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "invalid id",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id" : -1 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "error at layer usecase",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id" : 1 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {
				usecase.EXPECT().DeleteQuest(quest).Return(errors.New("any error")).Times(1)
			},
			wantStatusCode: http.StatusInternalServerError,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			router := mux.NewRouter()
			h := &handlers{
				usecase: tt.fields.u,
			}
			router.HandleFunc("/quest", h.DeleteQuest).Methods(http.MethodDelete)
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("DELETE", "/quest", strings.NewReader(tt.req.body))
			request = request.WithContext(ctx)
			tt.mock(tt.fields.u)
			router.ServeHTTP(recorder, request)
			var resp MessageResponse
			assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &resp))
			assert.Equal(t, tt.wantStatusCode, recorder.Code, "error code")
			assert.Equal(t, tt.resp.body, resp.Data)
			if tt.wantErr {
				assert.NotEqual(t, "", resp.Header.Error, "error message")
			} else {
				assert.Equal(t, "", resp.Header.Error, "error message")
			}
		})
	}
}

func TestUpdateQuestRank(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		u *MockUsecase
	}
	type requests struct {
		body string
	}
	type responses struct {
		body SuccesMessage
	}
	quest := model.Quest{
		ID:          1,
		MinimumRank: 12,
	}
	tests := []struct {
		name           string
		fields         fields
		req            requests
		resp           responses
		mock           func(*MockUsecase)
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "success update rank quest",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id" : 1, "minimum_rank" : 12 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: true},
			},
			mock: func(usecase *MockUsecase) {
				usecase.EXPECT().UpdateQuestRank(quest).Return(nil).Times(1)
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "json failed",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "empty id",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{ "minimum_rank" : 11 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "invalid id",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id" : -1, "minimum_rank": 11 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "empty rank",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id" : 1 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "invalid rank",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id" : 1, "minimum_rank": -2 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "error at layer usecase",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id" : 1, "minimum_rank" : 12 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {
				usecase.EXPECT().UpdateQuestRank(quest).Return(errors.New("any error")).Times(1)
			},
			wantStatusCode: http.StatusInternalServerError,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			router := mux.NewRouter()
			h := &handlers{
				usecase: tt.fields.u,
			}
			router.HandleFunc("/quest-rank", h.UpdateQuestRank).Methods(http.MethodPatch)
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("PATCH", "/quest-rank", strings.NewReader(tt.req.body))
			request = request.WithContext(ctx)
			tt.mock(tt.fields.u)
			router.ServeHTTP(recorder, request)
			var resp MessageResponse
			assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &resp))
			assert.Equal(t, tt.wantStatusCode, recorder.Code, "error code")
			assert.Equal(t, tt.resp.body, resp.Data)
			if tt.wantErr {
				assert.NotEqual(t, "", resp.Header.Error, "error message")
			} else {
				assert.Equal(t, "", resp.Header.Error, "error message")
			}
		})
	}
}

func TestUpdateQuestReward(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		u *MockUsecase
	}
	type requests struct {
		body string
	}
	type responses struct {
		body SuccesMessage
	}
	quest := model.Quest{
		ID:           1,
		RewardNumber: 250000,
	}
	tests := []struct {
		name           string
		fields         fields
		req            requests
		resp           responses
		mock           func(*MockUsecase)
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "success update rank quest",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id" : 1, "reward_number" : 250000 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: true},
			},
			mock: func(usecase *MockUsecase) {
				usecase.EXPECT().UpdateQuestReward(quest).Return(nil).Times(1)
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "json failed",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "empty id",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{ "reward_number" : 250000 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "invalid id",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id" : -1, "reward_number" : 250000 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "empty reward",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id" : 1 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "invalid reward",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id" : 1, "reward_number" : -1 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "error at layer usecase",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id" : 1, "reward_number" : 250000 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {
				usecase.EXPECT().UpdateQuestReward(quest).Return(errors.New("any error")).Times(1)
			},
			wantStatusCode: http.StatusInternalServerError,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			router := mux.NewRouter()
			h := &handlers{
				usecase: tt.fields.u,
			}
			router.HandleFunc("/quest-reward", h.UpdateQuestReward).Methods(http.MethodPatch)
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("PATCH", "/quest-reward", strings.NewReader(tt.req.body))
			request = request.WithContext(ctx)
			tt.mock(tt.fields.u)
			router.ServeHTTP(recorder, request)
			var resp MessageResponse
			assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &resp))
			assert.Equal(t, tt.wantStatusCode, recorder.Code, "error code")
			assert.Equal(t, tt.resp.body, resp.Data)
			if tt.wantErr {
				assert.NotEqual(t, "", resp.Header.Error, "error message")
			} else {
				assert.Equal(t, "", resp.Header.Error, "error message")
			}
		})
	}
}

func TestTakeQuest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		u *MockUsecase
	}
	type requests struct {
		body string
	}
	type responses struct {
		body SuccesMessage
	}
	tests := []struct {
		name           string
		fields         fields
		req            requests
		resp           responses
		mock           func(*MockUsecase)
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "success took a quest",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id": 1, "adv_id" : 1}`,
			},
			resp: responses{
				body: SuccesMessage{Success: true},
			},
			mock: func(usecase *MockUsecase) {
				quest := bulkQuest[0]
				quest.ID = 0
				usecase.EXPECT().TakeQuest(bulkQuest[0].ID, adv.ID).Return(nil).Times(1)
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "json failed",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "empty quest_id",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{ "adv_id" : 1}`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "invalid quest_id",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id": -1, "adv_id" : 1}`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "empty adv_id",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id": 1}`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "invalid adv_id",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id": 1, "adv_id" : -1}`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "error at layer usecase",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id": 1, "adv_id" : 1}`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {
				quest := bulkQuest[0]
				quest.ID = 0
				usecase.EXPECT().TakeQuest(bulkQuest[0].ID, adv.ID).Return(errors.New("any error")).Times(1)
			},
			wantStatusCode: http.StatusInternalServerError,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			router := mux.NewRouter()
			h := &handlers{
				usecase: tt.fields.u,
			}
			router.HandleFunc("/take-quest", h.TakeQuest).Methods(http.MethodPost)
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("POST", "/take-quest", strings.NewReader(tt.req.body))
			request = request.WithContext(ctx)
			tt.mock(tt.fields.u)
			router.ServeHTTP(recorder, request)
			var resp MessageResponse
			assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &resp))
			assert.Equal(t, tt.wantStatusCode, recorder.Code, "error code")
			assert.Equal(t, tt.resp.body, resp.Data)
			if tt.wantErr {
				assert.NotEqual(t, "", resp.Header.Error, "error message")
			} else {
				assert.Equal(t, "", resp.Header.Error, "error message")
			}
		})
	}
}

func TestReportQuest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		u *MockUsecase
	}
	type requests struct {
		body string
	}
	type responses struct {
		body SuccesMessage
	}
	tests := []struct {
		name           string
		fields         fields
		req            requests
		resp           responses
		mock           func(*MockUsecase)
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "success report a quest",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id": 1, "adv_id" : 1, "is_completed" : true}`,
			},
			resp: responses{
				body: SuccesMessage{Success: true},
			},
			mock: func(usecase *MockUsecase) {
				quest := bulkQuest[0]
				quest.ID = 0
				usecase.EXPECT().ReportQuest(bulkQuest[0].ID, adv.ID, true).Return(nil).Times(1)
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "success  uncompleted a quest",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id": 1, "adv_id" : 1, "is_completed" : false}`,
			},
			resp: responses{
				body: SuccesMessage{Success: true},
			},
			mock: func(usecase *MockUsecase) {
				quest := bulkQuest[0]
				quest.ID = 0
				usecase.EXPECT().ReportQuest(bulkQuest[0].ID, adv.ID, false).Return(nil).Times(1)
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "json failed",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "empty quest_id",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{ "adv_id" : 1,  "is_completed" : true}`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "invalid quest_id",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id": -1, "adv_id" : 1,  "is_completed" : true}`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "empty adv_id",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id": 1,  "is_completed" : true}`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "invalid adv_id",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id": 1, "adv_id" : -1,  "is_completed" : true}`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "empty is_completed",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id": 1, "adv_id" : 1}`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "error at layer usecase",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"quest_id": 1, "adv_id" : 1, "is_completed" : true}`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *MockUsecase) {
				quest := bulkQuest[0]
				quest.ID = 0
				usecase.EXPECT().ReportQuest(bulkQuest[0].ID, adv.ID, true).Return(errors.New("any error")).Times(1)
			},
			wantStatusCode: http.StatusInternalServerError,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			router := mux.NewRouter()
			h := &handlers{
				usecase: tt.fields.u,
			}
			router.HandleFunc("/report-quest", h.ReportQuest).Methods(http.MethodPost)
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("POST", "/report-quest", strings.NewReader(tt.req.body))
			request = request.WithContext(ctx)
			tt.mock(tt.fields.u)
			router.ServeHTTP(recorder, request)
			var resp MessageResponse
			assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &resp))
			assert.Equal(t, tt.wantStatusCode, recorder.Code, "error code")
			assert.Equal(t, tt.resp.body, resp.Data)
			if tt.wantErr {
				assert.NotEqual(t, "", resp.Header.Error, "error message")
			} else {
				assert.Equal(t, "", resp.Header.Error, "error message")
			}
		})
	}
}

func TestGetActiveAdventurer(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		u *MockUsecase
	}
	type requests struct {
		is     bool
		adv_id string
	}
	type responses struct {
		body []model.Quest
	}
	tests := []struct {
		name           string
		fields         fields
		req            requests
		resp           responses
		mock           func(*MockUsecase)
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "success get available quest",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				is:     true,
				adv_id: "1",
			},
			resp: responses{
				body: bulkQuest[:1],
			},
			mock: func(usecase *MockUsecase) {
				usecase.EXPECT().GetQuestActiveAdventurer(adv.ID).Return(bulkQuest[:1], nil).Times(1)
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "failed get available quest usecase",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				is:     true,
				adv_id: "1",
			},
			resp: responses{
				body: []model.Quest{},
			},
			mock: func(usecase *MockUsecase) {
				usecase.EXPECT().GetQuestActiveAdventurer(adv.ID).Return([]model.Quest{}, errors.New("any error")).Times(1)
			},
			wantStatusCode: http.StatusInternalServerError,
			wantErr:        true,
		},
		{
			name: "query empty",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				is:     false,
				adv_id: "",
			},
			resp: responses{
				body: []model.Quest{},
			},
			mock: func(usecase *MockUsecase) {
			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "query not int",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				is:     true,
				adv_id: "a",
			},
			resp: responses{
				body: []model.Quest{},
			},
			mock: func(usecase *MockUsecase) {
			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "query not valid",
			fields: fields{
				u: NewMockUsecase(mockCtrl),
			},
			req: requests{
				is:     true,
				adv_id: "-1",
			},
			resp: responses{
				body: []model.Quest{},
			},
			mock: func(usecase *MockUsecase) {
			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			router := mux.NewRouter()
			h := &handlers{
				usecase: tt.fields.u,
			}
			router.HandleFunc("/quest-active-adv", h.GetQuestActiveAdventurer).Methods(http.MethodGet)
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/quest-active-adv", strings.NewReader(``))
			if tt.req.is {
				values := request.URL.Query()
				values.Add("adv_id", tt.req.adv_id)
				request.URL.RawQuery = values.Encode()
			}
			request = request.WithContext(ctx)
			tt.mock(tt.fields.u)
			router.ServeHTTP(recorder, request)
			var resp GetQuestActiveAdventurer
			assert.NoError(t, json.Unmarshal(recorder.Body.Bytes(), &resp))
			assert.Equal(t, tt.wantStatusCode, recorder.Code, "error code")
			assert.Equal(t, tt.resp.body, resp.Data)
			if tt.wantErr {
				assert.NotEqual(t, "", resp.Header.Error, "error message")
			} else {
				assert.Equal(t, "", resp.Header.Error, "error message")
			}
		})
	}
}
