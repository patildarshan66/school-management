package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"schoolmanagement/internal/models"
	"strconv"
	"strings"
	"sync"
)

var (
	teachers = make(map[int]models.Teacher)
	mutex    = &sync.Mutex{}
	nextID   = 1
)

func init() {
	teachers[nextID] = models.Teacher{ID: nextID, FirstName: "John", LastName: "Doe", Subject: "Math", Class: "10A"}
	nextID++
	teachers[nextID] = models.Teacher{ID: nextID, FirstName: "Jane", LastName: "Smith", Subject: "Science", Class: "10B"}
	nextID++
	teachers[nextID] = models.Teacher{ID: nextID, FirstName: "Jane", LastName: "Doe", Subject: "Biology", Class: "10C"}
	nextID++
}

func TeachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTeachersHandler(w, r)
	case http.MethodPost:
		addTeacherHandler(w, r)
	case http.MethodDelete:
		w.Write([]byte("DELETE request received at /teachers"))
	case http.MethodPut:
		w.Write([]byte("PUT request received at /teachers"))
	case http.MethodPatch:
		w.Write([]byte("PATCH request received at /teachers"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")
	if idStr == "" {
		firstName := r.URL.Query().Get("first_name")
		lastName := r.URL.Query().Get("last_name")

		teachersList := make([]models.Teacher, 0, len(teachers))
		for _, t := range teachers {
			if (firstName == "" || firstName == t.FirstName) && (lastName == "" || lastName == t.LastName) {
				teachersList = append(teachersList, t)
			}
		}
		response := struct {
			Status string           `json:"status"`
			Count  int              `json:"count"`
			Data   []models.Teacher `json:"data"`
		}{
			Status: "success",
			Count:  len(teachersList),
			Data:   teachersList,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	//convert string id to int
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println(err)
		return
	}

	t, exists := teachers[id]

	if !exists {
		err := fmt.Sprintf("Teacher not found with the ID: %d", id)
		http.Error(w, err, http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(t)
}

func addTeacherHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	var newTeachers []models.Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	addedTeachers := make([]models.Teacher, len(newTeachers))
	for i, teacher := range newTeachers {
		teacher.ID = nextID
		teachers[nextID] = teacher
		addedTeachers[i] = teacher
		nextID++
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(addedTeachers),
		Data:   addedTeachers,
	}
	json.NewEncoder(w).Encode(response)
}
