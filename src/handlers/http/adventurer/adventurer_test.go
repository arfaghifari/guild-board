package adventurer

import (
	"context"
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
	mockUsecase := usecase.NewMockUsecase(mockCtrl)
	testHandlers := &handlers{usecase: mockUsecase}
	mockUsecase.EXPECT().CreateAdventurer(adv).Return(nil).Times(1)
	ctx := context.Background()
	router := mux.NewRouter()
	router.HandleFunc("/adventurer", testHandlers.CreateAdventurer).Methods(http.MethodPost)
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/adventurer", strings.NewReader(`{"name" : "andi", "rank" : 11 }`))
	request = request.WithContext(ctx)
	router.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code, "error code")
	assert.Equal(t, "success", string(recorder.Body.String()))
}

func TestUpdateAdventureRank(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUsecase := usecase.NewMockUsecase(mockCtrl)
	testHandlers := &handlers{usecase: mockUsecase}

	mockUsecase.EXPECT().UpdateAdventurerRank(adv2).Return(nil).Times(1)
	ctx := context.Background()
	router := mux.NewRouter()
	router.HandleFunc("/adventurer-rank", testHandlers.UpdateAdventurerRank).Methods(http.MethodPatch)
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/adventurer-rank", strings.NewReader(`{"id" : 1, "rank" : 12 }`))
	request = request.WithContext(ctx)
	router.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code, "error code")
	assert.Equal(t, "success", string(recorder.Body.String()))
}
