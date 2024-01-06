package main

import (
	"fees/db"
	"fees/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes(r *mux.Router) {
	r.HandleFunc("/", homePageHandler).Methods("OPTIONS")
	r.HandleFunc("/class/{num}", handlers.QueryClassHandler).Methods("GET")
	r.HandleFunc("/student", handlers.QueryStudentHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/pay/student", handlers.QueryPayFeeHanlder).Methods("POST", "OPTIONS")
	r.HandleFunc("/add/student", handlers.QueryAddStudentHandler).Methods("POST", "OPTIONS")
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home page")
}

func main() {

	if err := db.Connect(); err != nil {
		fmt.Println("Cannt connect to DB", err)
		return
	}
	r := mux.NewRouter()

	setupRoutes(r)
	r.Use(mux.CORSMethodMiddleware(r))
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Finish testing adding a new student using postman

// Pay fee working fine. Add errors for edge cases.
