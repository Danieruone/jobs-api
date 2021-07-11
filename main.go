package main

import (
	"log"
	"net/http"
	"server/handlers"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	router.Get("/api/jobs", handlers.GetJobs)
	router.Post("/api/jobs", handlers.AddJobs)
	//run it on port 8080
	err := http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
