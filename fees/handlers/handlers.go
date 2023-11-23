package handlers

import (
	"encoding/json"
	"fees/db"
	"fees/students"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func QueryAddStudentHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal("Cant get form Data")
	}

	SID := r.FormValue("sid")
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	parent := r.FormValue("parent")
	contact := r.FormValue("contact")
	acadyear := r.FormValue("acadyear")
	class := r.FormValue("class")
	section := r.FormValue("section")
	busfee := r.FormValue("busfee")
	tutionfee := r.FormValue("tutionfee")
	totalmonths := r.FormValue("totalmonths")

	sid, err := strconv.ParseInt(SID, 10, 32)
	if err != nil {
		log.Fatal("Unable to convert Student ID")
	}

	cls, err := strconv.ParseInt(class, 10, 32)
	if err != nil {
		log.Fatal("Unable to convert Class")
	}

	months, err := strconv.ParseInt(totalmonths, 10, 32)
	if err != nil {
		log.Fatal("Unable to convert Months")
	}

	st := students.NewStudent(int(sid), fname, lname, parent, contact, acadyear, int(cls), section, busfee, tutionfee, int(months))

	client := db.GetClient()
	result, err := db.QueryAddStudent(client, *st)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Cannot insert student"))
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

func QueryPayFeeHanlder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		log.Fatal("Unable to get Student ID")
	}
	amount := vars["amount"]

	client := db.GetClient()

	result, err := db.QueryPayFee(client, int(id), amount)
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
