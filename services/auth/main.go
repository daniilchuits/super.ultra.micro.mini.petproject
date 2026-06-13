package main

import (
	"auth/database"
	"auth/internal/handlers"
	usersrepo "auth/internal/users_repo"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

func main() {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	cnnStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable host=users-postgres",
		user, password, dbname,
	)

	db, err := sql.Open("postgres", cnnStr)
	if err != nil {
		log.Println("Openning db err:", err)
		return
	}

	if err = db.Ping(); err != nil {
		log.Println("Pinging db err:", err)
		return
	}

	if err = database.CreateUsersTable(db); err != nil {
		log.Println("Err in creating repo users:", err)
		return
	}

	repoManager := usersrepo.NewRepoManager(db)

	getNote := repoManager
	insertNote := repoManager

	registerHandler := handlers.NewRegisterHandler(getNote, insertNote)

	r := chi.NewRouter()

	r.Post("/register", registerHandler.RegisterUser)

	//register
	//login

	log.Fatal(http.ListenAndServe(":8081", r))
}
