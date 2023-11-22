package handlers

import (
	"encoding/json"
	"fees/db"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func QueryClassHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	class, err := strconv.ParseInt(vars["num"], 10, 32)
	if err != nil {
		log.Fatal("Unable to get Class")
	}

	client := db.GetClient()

	result, err := db.QueryClass(client, int(class))
	if err != nil {
		w.WriteHeader(500)
	}

	// Convert the array to JSON
	jsonData, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func QueryStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		log.Fatal("Unable to get Student ID")
	}

	client := db.GetClient()

	result, err := db.QueryStudent(client, int(id))
	if result.SID == 0 {
		w.WriteHeader(500)
		w.Write([]byte("No such student exists"))
		return
	}

	if err != nil {
		w.WriteHeader(500)
	}

	// Convert the array to JSON
	jsonData, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
