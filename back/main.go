package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Task struct {
	Id    uint8  `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var Tasks = []Task{
	{Id: 1, Title: "Tache 1", Done: false},
	{Id: 2, Title: "Tache 2", Done: false},
}

func getTaskIndexFromId(id uint8) (result int) {
	result = -1
	for index, v := range Tasks {
		if v.Id == id {
			result = index
		}
	}
	return result
}

func getNextIndex() (result uint8) {
	result = 0
	for _, v := range Tasks {
		if v.Id > result {
			result = v.Id
		}
	}
	return result + 1
}

func returnAllTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllTasks")
	json.NewEncoder(w).Encode(Tasks)
}

func createOrReplaceTask(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var task Task
	json.Unmarshal(reqBody, &task)
	if getTaskIndexFromId(task.Id) == -1 {
		// No index OR index not existing
		task.Id = getNextIndex()
		Tasks = append(Tasks, task)
	} else {
		// check for update
		Tasks[getTaskIndexFromId(task.Id)] = task
	}
	json.NewEncoder(w).Encode(task)
}

func handleRequests() {
	http.HandleFunc("/tasks", returnAllTasks)
	http.HandleFunc("/tasks/create", createOrReplaceTask)
	http.HandleFunc("/tasks/update", createOrReplaceTask)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	fmt.Println("Server running on http://localhost:10000")
	handleRequests()
}
