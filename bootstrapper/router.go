package bootstrapper

import (
	"github.com/jaijiv/go-social-api/controllers"

	"github.com/gorilla/mux"
)

func SetupWebRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/hello/{name:[0-9]+}", controllers.Hello).Methods("GET", "OPTIONS")
	router.HandleFunc("/hello", controllers.GetAllHello).Methods("GET", "OPTIONS")

	return router
}
