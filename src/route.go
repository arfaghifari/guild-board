package src

import (
	"net/http"
	"time"

	questHandlers "github.com/arfaghifari/guild-board/src/handlers/http/quest"
	server "github.com/arfaghifari/guild-board/src/server"
	"github.com/gorilla/mux"
)

func Main() {

	// Init database connection
	// database.InitDB()

	// Init serve HTTP
	router := mux.NewRouter()

	// routes http
	router.HandleFunc("/get-hello", questHandlers.GetHello).Methods(http.MethodGet)
	router.HandleFunc("/get-quest-status", questHandlers.GetQuestByStatus).Methods(http.MethodGet)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         8000,
	}
	server.Serve(serverConfig, router)
}
