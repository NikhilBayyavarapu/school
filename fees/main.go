package main

import (
	"fees/db"
	"fees/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes(r *mux.Router) {
	r.HandleFunc("/", homePageHandler)
	r.HandleFunc("/class/{num}", handlers.QueryClassHandler).Methods("GET")
	r.HandleFunc("/student/{id}", handlers.QueryStudentHandler).Methods("GET")
	r.HandleFunc("/pay/student/{id}/{amount}", handlers.QueryPayFeeHanlder).Methods("POST")
	r.HandleFunc("/add/student", handlers.QueryAddStudentHandler).Methods("POST")
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
	http.ListenAndServe(":8080", r)
}

// Finish testing adding a new student using postman
