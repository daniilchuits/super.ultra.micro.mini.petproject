package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"worker/database"
	goworker "worker/internal/go_worker"
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

	wd, _ := os.Getwd()
	fmt.Println(wd)

	repoManager := jobsrepo.NewRepoManager(db)

	check := repoManager
	insert := repoManager
	getJobs := repoManager
	getJob := repoManager
	update := repoManager

	chooseForGoWorker := repoManager
	updForGoWorker := repoManager

	goWorker := goworker.NewGoWorker(chooseForGoWorker, updForGoWorker)

	go goWorker.Process()

	postJobHandler := handlers.NewPostJobHandler(check, insert)
	getJobsHandler := handlers.NewGetJobHandler(getJobs)
	getJobHandler := handlers.NewGetOneNoteHandler(getJob)
	updateHandler := handlers.NewUpdateHandler(check, update, getJob)

	r := chi.NewRouter()

	r.Post("/jobs", postJobHandler.PostNote)
	r.Get("/jobs", getJobsHandler.GetJobs)
	r.Get("/jobs/{id}", getJobHandler.GetOneJob)
	r.Post("/upload", updateHandler.UpdateFile)

	log.Fatal(http.ListenAndServe(":8082", r))
	// вроде все, опять же, при тестах выявится много багов, сейчас напишу ридми,
	// хезе пока как его писать... ну разберемся
}
