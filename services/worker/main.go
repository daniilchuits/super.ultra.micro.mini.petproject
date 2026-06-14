package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"worker/database"
	"worker/internal/handlers"
	jobsrepo "worker/internal/jobs_repo"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {

	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	cnnStr := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=jobs-postgres sslmode=disable",
		dbName, user, password,
	)

	db, err := sql.Open("postgres", cnnStr)
	if err != nil {
		log.Println("Error openning db:", err)
		log.Println(cnnStr)
		return
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Println("Error pinging db:", err)
		return
	}

	dbManager := database.NewDbManager(db)
	if err = dbManager.CreateJobsTable(); err != nil {
		log.Println("Error creating jobs table:", err)
		return
	}

	repoManager := jobsrepo.NewRepoManager(db)

	check := repoManager
	insert := repoManager
	getJobs := repoManager
	// getJob := repoManager

	postJobHandler := handlers.NewPostJobHandler(check, insert)
	getJobsHandler := handlers.NewGetJobHandler(getJobs)
	// getJobHandler := handlers.NewGetOneNoteHandler(getJob)

	r := chi.NewRouter()

	r.Post("/jobs", postJobHandler.PostNote)
	r.Get("/jobs", getJobsHandler.GetJobs)
	// r.Get("/jobs/{id}", getJobHandler.GetOneJob)
	// POST /upload

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("worker ok"))
	}) // endpoints for worker

	log.Fatal(http.ListenAndServe(":8082", r)) // zachekinil, works ok, do handlers
}
