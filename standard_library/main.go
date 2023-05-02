package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Similar to a controller?
type tasksHandler struct{}
type rootHandler struct{}

// Struct to handle tasks
type Task struct {
	Id   int       `json:"id"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

var tasks []Task

// Setting up mux server
func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &rootHandler{})
	mux.Handle("/tasks/", &tasksHandler{})
	// Start and listen on port 8080
	http.ListenAndServe(":8080", mux)
}

func (h *tasksHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		// Return tasks
		tasks, err := loadTasks()
		if err != nil {
			panic(err)
		}
		jsonData, err := json.Marshal(tasks)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	case http.MethodPost:
		// Add a new task
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		var task Task
		err = json.Unmarshal(body, &task)
		if err != nil {
			panic(err)
		}

		// Load tasks from the db
		tasks, err := loadTasks()
		if err != nil {
			panic(err)
		}

		// Add new task to that list
		tasks = append(tasks, task)

		// Save task to db
		err = saveTasks(tasks)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusCreated)
		// How to show created task?
	}
}

func loadTasks() ([]Task, error) {
	var tasks []Task
	file, err := os.Open("db.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	file, err := os.Create("db.json")
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		return err
	}

	return nil
}

func (h *rootHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, Home Page!"))
}

// TODO:
//  Determine how to create a JSON file to store our data
