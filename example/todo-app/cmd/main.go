package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/fossildev/todo-app/pkg/app"
	"github.com/fossildev/todo-app/pkg/database"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func init() {
	flag.StringVar(&database.Host, "database-host", "127.0.0.1", "Database host")
	flag.StringVar(&database.User, "database-user", "todoapp", "Database user")
	flag.StringVar(&database.Password, "database-password", "secret", "Database password")
	flag.StringVar(&database.Name, "database-name", "todoapp", "Database name")
}

func main() {
	flag.Parse()

	todo := app.NewTodoHandler()
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/todos", todo.ListTodos).Methods("GET")
	api.HandleFunc("/todos", todo.CreateTodos).Methods("POST")
	api.HandleFunc("/todos/{todo_id}", todo.UpdateTodos).Methods("PUT")
	api.HandleFunc("/todos/{todo_id}", todo.DeleteTodos).Methods("DELETE")

	log.Print("Serve app on port 9000")

	handler := handlers.LoggingHandler(os.Stdout, router)
	if err := http.ListenAndServe(":9000", handler); err != nil {
		log.Fatal(err)
	}
}
