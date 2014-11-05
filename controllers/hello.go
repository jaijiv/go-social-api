package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/jaijiv/go-social-api/services/hellosvc"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("Hello...")
	// record call in db
	go hellosvc.Post(vars["name"])
	fmt.Fprint(w, "Hello ", vars["name"])
}

func GetAllHello(w http.ResponseWriter, r *http.Request) {
	hellosvc.GetAll()
}
