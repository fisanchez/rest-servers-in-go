package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Task struct {
	Id   int       `json:"id"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

var tasks = []Task{
	{
		Id:   1,
		Text: "Task 1",
		Tags: []string{"tag1", "tag2"},
		Due:  time.Now(),
	},
}

func main() {
	router := chi.NewRouter()

	// Logs server requests to terminal
	router.Use(middleware.Logger)

	// router
	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World!"))
	})

	// Get(pattern string, h http.HandlerFunc)
	router.Get("/tasks", getTasks)
	router.Get("/tasks/{id}", getTask)
	router.Post("/tasks", postTask)
	router.Post("/tasks/{id}", deleteTask)

	http.ListenAndServe(":8080", router)
}

func getTasks(writer http.ResponseWriter, request *http.Request) {
	// writer.Write([]byte("Getting all tasks!"))
	jsonTasks, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jsonTasks)
}

func getTask(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Getting a task!"))
}

func postTask(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Posting a task!"))
}

func deleteTask(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Deleting a task!"))
}
