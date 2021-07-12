package handlers

import (
	"encoding/json"
	. "fmt"
	"net/http"
	"server/models"
	"strconv"
)

var jobs []models.Job

func init() {
	///add some job to the slice
	jobs = append(jobs, models.Job{ID: 1, Name: "Accounting"})
	jobs = append(jobs, models.Job{ID: 2, Name: "Programming"})
}

func GetJobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}

func AddJobs(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var model models.Job
	err := decoder.Decode(&model)
	if err != nil {
		panic(err)
	}
	jobs = append(jobs, models.Job{ID: model.ID, Name: model.Name})
	Println("Job added")
	json.NewEncoder(w).Encode(jobs)
}

func RemoveIndex(s []models.Job, index int) []models.Job {
	return append(s[:index], s[index+1:]...)
}

func DeleteJob(w http.ResponseWriter, r *http.Request) {
	key, ok := r.URL.Query()["ID"]
	value, err := strconv.Atoi(key[0])
	Print(err)
	if !ok {
		Println("Url Param is missing")
		return
	}
	for i, job := range jobs {
		if value == job.ID {
			jobs = RemoveIndex(jobs, i)
		}
	}
	json.NewEncoder(w).Encode(jobs)
}
