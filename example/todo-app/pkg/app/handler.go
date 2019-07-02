package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Handler impement todo handler
type Handler struct {
	DS TodoDatastore
}

// NewTodoHandler create new todo handler instance
func NewTodoHandler() *Handler {
	handler := new(Handler)
	handler.DS = NewTodoDatastore()
	return handler
}

// ListTodos handler
func (h *Handler) ListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.DS.GetAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	res := new(TodoList)
	res.Todos = todos

	b, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// CreateTodos handler
func (h *Handler) CreateTodos(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)

	id, err := h.DS.Insert(&todo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	res, err := h.DS.Get(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	b, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// UpdateTodos handler
func (h *Handler) UpdateTodos(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	todoid, _ := strconv.ParseInt(mux.Vars(r)["todo_id"], 10, 64)
	json.NewDecoder(r.Body).Decode(&todo)

	if err := h.DS.Update(todoid, &todo); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	b, _ := json.Marshal(todo)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// DeleteTodos handler
func (h *Handler) DeleteTodos(w http.ResponseWriter, r *http.Request) {
	todoid, _ := strconv.ParseInt(mux.Vars(r)["todo_id"], 10, 64)
	if err := h.DS.Delete(todoid); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
