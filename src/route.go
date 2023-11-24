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
	router.HandleFunc("/get-quest-status", questHandlers.GetQuestByStatus).Methods(http.MethodGet)
	// router.HandleFunc("/update-product", handlers.UpdateProduct).Methods(http.MethodPatch)
	// router.HandleFunc("/delete-product", handlers.DeleteProduct).Methods(http.MethodDelete)

	// router.HandleFunc("/banner", handlers.UploadBanner).Methods(http.MethodPost)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         8000,
	}
	server.Serve(serverConfig, router)
}
