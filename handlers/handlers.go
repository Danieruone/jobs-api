package handlers

import (
	"encoding/json"
	. "fmt"
	"net/http"
	"server/models"
)

//A handler to fetch all the jobs
var jobs []models.Job

func GetJobs(w http.ResponseWriter, r *http.Request) {
	//make a slice to hold our jobs data
	///add some job to the slice
	Println(r)
	jobs = append(jobs, models.Job{ID: 1, Name: "Accounting"})
	jobs = append(jobs, models.Job{ID: 2, Name: "Programming"})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}
