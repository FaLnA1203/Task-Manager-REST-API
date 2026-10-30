package api

import (
	"TaskManager/task"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type HTTPHandler struct {
	TasksList *task.TasksList
}

func NewHTTPhandler(tasks *task.TasksList) *HTTPHandler {
	return &HTTPHandler{TasksList: tasks}
}

/*
pattern:/tasks
mettod: post
body: json

failed:
status code: 500, 400, 409
body: json error + time
succesed:
status code: 201
body: json
*/

func (h *HTTPHandler) HandlerCreateTask(w http.ResponseWriter, r *http.Request) {
	var DTOTask DtoTask
	if err := json.NewDecoder(r.Body).Decode(&DTOTask); err != nil {
		errDTO := DtoError{
			Error: err.Error(),
			Time:  time.Now(),
		}

		http.Error(w, errDTO.ErrorToString(), http.StatusBadRequest)
		return
	}
	if err := DTOTask.Validate(); err != nil {
		errDTO := DtoError{
			Error: err.Error(),
			Time:  time.Now(),
		}
		http.Error(w, errDTO.ErrorToString(), http.StatusBadRequest)
		return
	}

	t := task.CreateTask(DTOTask.Name, DTOTask.Description)

	if err := h.TasksList.AddTask(&t); err != nil {
		errDTO := DtoError{
			Error: err.Error(),
			Time:  time.Now(),
		}
		if errors.Is(err, task.TaskAlreadyExists) {
			http.Error(w, errDTO.ErrorToString(), http.StatusConflict)
		} else {
			http.Error(w, errDTO.ErrorToString(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println(err)
		return
	}
}

/*
pattern:/tasks/{task_id}
mettod: get
body: -

failed:
status code: 500, 400, 404
body: json error + time
succesed:
status code: 200
body: json
*/
func (h *HTTPHandler) HandlerGetTask(w http.ResponseWriter, r *http.Request) {
	title := chi.URLParam(r, "title")
	t, err := h.TasksList.GetTask(title)
	if err != nil {
		errDTO := DtoError{
			Error: err.Error(),
			Time:  time.Now(),
		}
		if errors.Is(err, task.TaskNotFound) {
			http.Error(w, errDTO.ErrorToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ErrorToString(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println(err)
		return
	}
}

/*
pattern:/tasks/
mettod: get
body: -

failed:
status code: 500, 400, 404
body: json error + time
succesed:
status code: 200
body: json
*/
func (h *HTTPHandler) HandlerGetAllTasks(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("complete")
	var lt map[string]task.Task
	switch q {
	case "true":
		lt = h.TasksList.ListCompletedTasks()
	case "false":
		lt = h.TasksList.ListUncompletedTasks()
	default:
		lt = h.TasksList.ListTasks()
	}

	b, err := json.MarshalIndent(lt, "", "  ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println(err)
		return
	}
}

/*
pattern:/tasks/{task_id}
mettod: patch
body: json

failed:
status code: 500, 400, 404, 409
body: json error + time
succesed:
status code: 200
body: json
*/
func (h *HTTPHandler) HandlerCompleteTask(w http.ResponseWriter, r *http.Request) {
	var DTOCompleted DtoCompleted
	if err := json.NewDecoder(r.Body).Decode(&DTOCompleted); err != nil {
		errDTO := DtoError{
			Error: err.Error(),
			Time:  time.Now(),
		}
		http.Error(w, errDTO.ErrorToString(), http.StatusBadRequest)
		return
	}
	title := chi.URLParam(r, "title")
	err := h.TasksList.CompliteTask(title, DTOCompleted.Completed)
	if err != nil {
		errDTO := DtoError{
			Error: err.Error(),
			Time:  time.Now(),
		}
		if errors.Is(err, task.TaskAlreadyExists) {
			http.Error(w, errDTO.ErrorToString(), http.StatusConflict)
		} else {
			http.Error(w, errDTO.ErrorToString(), http.StatusInternalServerError)
		}
		return
	}
	t, _ := h.TasksList.GetTask(title)
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println(err)
		return
	}
}

/*
pattern:/tasks/{task_id}
mettod: Delete
body: -

failed:
status code: 500, 400, 404
body: json error + time
succesed:
status code: 204
body: -
*/
func (h *HTTPHandler) HandlerDeleteTask(w http.ResponseWriter, r *http.Request) {
	title := chi.URLParam(r, "title")
	if err := h.TasksList.DeleteTask(title); err != nil {
		errDTO := DtoError{
			Error: err.Error(),
			Time:  time.Now(),
		}
		if errors.Is(err, task.TaskNotFound) {
			http.Error(w, errDTO.ErrorToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ErrorToString(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
