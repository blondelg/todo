package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(task)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status Deleted"
	jsonResp, _ := json.Marshal(resp)
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 8)
	id8 := uint8(id)
	// Delete element from id
	var index = getTaskIndexFromId(id8)
	if index >= 0 {
		Tasks = append(Tasks[:index], Tasks[1+index:]...)
	}
	// w.Write(jsonResp)
	json.NewEncoder(w).Encode(jsonResp)
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/tasks", returnAllTasks).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/tasks/create", createOrReplaceTask).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/tasks/update", createOrReplaceTask).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/tasks/{id}", deleteTask).Methods(http.MethodDelete, http.MethodOptions)
	r.Use(mux.CORSMethodMiddleware(r))
	log.Fatal(http.ListenAndServe(":10000", r))
}

func main() {
	fmt.Println("Server running on http://localhost:10000")
	handleRequests()
}
