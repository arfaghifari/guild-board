package src

import (
	"net/http"
	"time"

	advHandlers "github.com/arfaghifari/guild-board/src/handlers/http/adventurer"
	qstHandlers "github.com/arfaghifari/guild-board/src/handlers/http/quest"
	server "github.com/arfaghifari/guild-board/src/server"
	"github.com/gorilla/mux"
)

func Main() {

	// Init database connection
	// database.InitDB()

	// Init serve HTTP
	router := mux.NewRouter()
	questHandlers, _ := qstHandlers.NewHandlers()
	adventurerHandlers, _ := advHandlers.NewHandlers()
	// routes http
	router.HandleFunc("/hello", qstHandlers.GetHello).Methods(http.MethodGet)

	router.HandleFunc("/quest-status", questHandlers.GetQuestByStatus).Methods(http.MethodGet)
	router.HandleFunc("/quest", questHandlers.CreateQuest).Methods(http.MethodPost)
	router.HandleFunc("/quest", questHandlers.DeleteQuest).Methods(http.MethodDelete)
	router.HandleFunc("/quest-rank", questHandlers.UpdateQuestRank).Methods(http.MethodPatch)
	router.HandleFunc("/quest-reward", questHandlers.UpdateQuestReward).Methods(http.MethodPatch)

	router.HandleFunc("/adventurer", adventurerHandlers.CreateAdventurer).Methods(http.MethodPost)
	router.HandleFunc("/adventurer-rank", adventurerHandlers.UpdateAdventurerRank).Methods(http.MethodPatch)

	router.HandleFunc("/take-quest", questHandlers.TakeQuest).Methods(http.MethodPost)
	router.HandleFunc("/done-quest", questHandlers.ReportQuest).Methods(http.MethodPost)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         8000,
	}
	server.Serve(serverConfig, router)
}
