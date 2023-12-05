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

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	setupRoutes(r)
	http.ListenAndServe(":8080", r)
}

// Finish testing adding a new student using postman
