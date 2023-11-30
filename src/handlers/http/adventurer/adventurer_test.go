package adventurer

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	model "github.com/arfaghifari/guild-board/src/model/adventurer"
	usecase "github.com/arfaghifari/guild-board/src/usecase/adventurer"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var adv = model.Adventurer{
	Name: "andi",
	Rank: 11,
}
var adv2 = model.Adventurer{
	ID:   1,
	Rank: 12,
}

func TestCreateAdventurer(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		u *usecase.MockUsecase
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
		mock           func(*usecase.MockUsecase)
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "success created an adventurer",
			fields: fields{
				u: usecase.NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"name" : "andi", "rank" : 11 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: true},
			},
			mock: func(usecase *usecase.MockUsecase) {
				usecase.EXPECT().CreateAdventurer(adv).Return(nil).Times(1)
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "json failed",
			fields: fields{
				u: usecase.NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *usecase.MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "empty name",
			fields: fields{
				u: usecase.NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{ "rank" : 11 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *usecase.MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "invalid name",
			fields: fields{
				u: usecase.NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"name" : "" , "rank" : 11}`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *usecase.MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "empty rank",
			fields: fields{
				u: usecase.NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"name" : "andi" }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *usecase.MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "invalid rank",
			fields: fields{
				u: usecase.NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"rank" : -2, "name" : "andi" }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *usecase.MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "error at layer usecase",
			fields: fields{
				u: usecase.NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"name" : "andi", "rank" : 11 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *usecase.MockUsecase) {
				usecase.EXPECT().CreateAdventurer(adv).Return(errors.New("any error")).Times(1)
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
			router.HandleFunc("/adventurer", h.CreateAdventurer).Methods(http.MethodPost)
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("POST", "/adventurer", strings.NewReader(tt.req.body))
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

func TestUpdateAdventureRank(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	type fields struct {
		u *usecase.MockUsecase
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
		mock           func(*usecase.MockUsecase)
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "success updated an adventurer rank",
			fields: fields{
				u: usecase.NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"id" : 1, "rank" : 12 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: true},
			},
			mock: func(usecase *usecase.MockUsecase) {
				usecase.EXPECT().UpdateAdventurerRank(adv2).Return(nil).Times(1)
			},
			wantStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name: "json failed",
			fields: fields{
				u: usecase.NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *usecase.MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "empty id",
			fields: fields{
				u: usecase.NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{ "rank" : 11 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *usecase.MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "invalid id",
			fields: fields{
				u: usecase.NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"id" : -1, "rank": 11 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *usecase.MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "empty rank",
			fields: fields{
				u: usecase.NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"id" : 1 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *usecase.MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "invalid rank",
			fields: fields{
				u: usecase.NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"id" : 1, "rank": -2 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *usecase.MockUsecase) {

			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
		},
		{
			name: "error at layer usecase",
			fields: fields{
				u: usecase.NewMockUsecase(mockCtrl),
			},
			req: requests{
				body: `{"id" : 1, "rank" : 12 }`,
			},
			resp: responses{
				body: SuccesMessage{Success: false},
			},
			mock: func(usecase *usecase.MockUsecase) {
				usecase.EXPECT().UpdateAdventurerRank(adv2).Return(errors.New("any error")).Times(1)
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
			router.HandleFunc("/adventurer-rank", h.UpdateAdventurerRank).Methods(http.MethodPatch)
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("PATCH", "/adventurer-rank", strings.NewReader(tt.req.body))
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
