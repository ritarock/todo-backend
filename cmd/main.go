package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ritarock/todo-backend/api"
	"github.com/ritarock/todo-backend/api/rest"
	"github.com/ritarock/todo-backend/repository"
)

const timeout = 2 * time.Second

func main() {
	db, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTodoRepository(db)
	handler := api.NewHandler(repo)
	srv, err := rest.NewServer(handler)
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
