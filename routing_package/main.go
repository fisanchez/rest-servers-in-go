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
	router.Use(setJSONContentTypeHeader)

	// root
	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Nothing to see here, move along."))
	})
	// Get(pattern string, h http.HandlerFunc)
	router.Get("/tasks", getTasks)
	// Get a task from an ID
	router.Get("/tasks/{id}", getTask)
	// Create a task
	router.Post("/tasks", createTask)
	// Delete a task
	router.Post("/tasks/{id}", deleteTask)

	http.ListenAndServe(":8080", router)
}

// Writing a middleware function
// TODO: Read up more about this. I'm not 100% sure what's going on here.
func setJSONContentTypeHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func getTasks(writer http.ResponseWriter, request *http.Request) {
	jsonTasks, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	writer.Write(jsonTasks)
}

func getTask(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Getting a task!"))
}

// This works without explictly returning; why??
func createTask(writer http.ResponseWriter, request *http.Request) {
	var task Task

	// Take the request body and decode it into the task struct
	err := json.NewDecoder(request.Body).Decode(&task)
	if err != nil {
		panic(err)
	}

	// Create an id for the new task
	task.Id = len(tasks) + 1
	tasks = append(tasks, task)
	json.NewEncoder(writer).Encode(task)
}

func deleteTask(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Deleting a task!"))
}
