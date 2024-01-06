package handlers

import (
	"encoding/json"
	"fees/db"
	"fees/students"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Request struct {
	SID int
}

func QueryAddStudentHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := r.ParseForm()
	if err != nil {
		log.Fatal("Cant get form Data")
	}

	fmt.Println("Here", r.PostForm)

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

	bfee, err := strconv.ParseFloat(busfee, 32)
	if err != nil {
		log.Fatal("Unable to convert BusFee")
	}

	tfee, err := strconv.ParseFloat(tutionfee, 32)
	if err != nil {
		log.Fatal("Unable to convert TutionFee")
	}

	st := students.NewStudent(int(sid), fname, lname, parent, contact, acadyear, int(cls), section, float32(bfee), float32(tfee), int(months))

	client := db.GetClient()
	result, err := db.QueryAddStudent(client, *st)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Cannot insert student"))
	}

	fmt.Println(result)

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

	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := r.ParseForm()
	if err != nil {
		log.Fatal("Cant get form Data")
	}

	fmt.Println("Here", r.PostForm)

	SID := r.FormValue("sid")
	fmt.Println("SID: ", SID)

	sid, err := strconv.ParseInt(SID, 10, 32)
	if err != nil {
		log.Fatal("Unable to convert Student ID")
	}

	client := db.GetClient()

	result, err := db.QueryStudent(client, int(sid))
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

	fmt.Println(result)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func QueryPayFeeHanlder(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := r.ParseForm()
	if err != nil {
		log.Fatal("Cant get form Data")
	}

	// fmt.Println("Here", r.PostForm)

	SID := r.FormValue("sid")
	busfee := r.FormValue("busfee")
	tutionfee := r.FormValue("tutionfee")

	sid, err := strconv.ParseInt(SID, 10, 32)
	if err != nil {
		log.Fatal("Unable to convert Student ID")
	}

	bfee, err := strconv.ParseFloat(busfee, 32)
	if err != nil {
		log.Fatal("Unable to convert BusFee")
	}

	tfee, err := strconv.ParseFloat(tutionfee, 32)
	if err != nil {
		log.Fatal("Unable to convert TutionFee")
	}

	client := db.GetClient()

	result, err := db.QueryPayFee(client, int(sid), float32(bfee+tfee))

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	if result.SID == 0 {
		w.WriteHeader(500)
		w.Write([]byte("No such student exists"))
		return
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
