package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/naufalilyasa/todolist-be-golang/internal/models"
	"github.com/naufalilyasa/todolist-be-golang/internal/services"
	"github.com/naufalilyasa/todolist-be-golang/pkg"
)

type TodoHandler struct {
	service services.TodoService
}

func NewTodoHandler(service services.TodoService) *TodoHandler {
	return &TodoHandler{service}
}

/*
Todos GET /api/todos

	# List todos with pagination and optional filters
	# Query params: page, limit, search, sort_by, sort_order
*/
func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	// Pagination
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	page := 1
	limit := 10
	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	// Filters
	search := r.URL.Query().Get("search")
	status := r.URL.Query().Get("status")
	priority := r.URL.Query().Get("priority")
	categoryId := r.URL.Query().Get("category_id")

	// Optional: set empty string to nil
	var statusPtr, priorityPtr, categoryPtr *string
	if status != "" {
		statusPtr = &status
	}
	if priority != "" {
		priorityPtr = &priority
	}
	if categoryId != "" {
		categoryPtr = &categoryId
	}

	todos, total, err := h.service.FindAllWithFilters(page, limit, search, statusPtr, priorityPtr, categoryPtr)
	if err != nil {
		pkg.JSONError(w, http.StatusInternalServerError, "Failed to fetch todos")
		return
	}

	pagination := &pkg.Pagination{
		CurrentPage: page,
		PerPage:     limit,
		Total:       total,
		Total_pages: int((total + int64(limit) - 1) / int64(limit)),
	}

	pkg.JSONSuccess(w, http.StatusOK, "Todos fetched successfully", todos, pagination)
}

// GET /api/todos/:id # Get specific todo
func (h *TodoHandler) GetTodoById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	todo, err := h.service.GetTodoById(id)
	if err != nil {
		pkg.JSONError(w, http.StatusNotFound, "Failed to get todo")
		return
	}

	pkg.JSONSuccess(w, http.StatusOK, "Success get todo", todo, nil)
}

// POST /api/todos # Create new todo
func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		pkg.JSONError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	created, err := h.service.CreateTodo(todo)
	if err != nil {
		pkg.JSONError(w, http.StatusInternalServerError, "Failed create todo")
	}

	pkg.JSONSuccess(w, http.StatusCreated, "Success created todo", created, nil)
}

// PUT /api/todos/:id # Update todo
func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		pkg.JSONError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	todo.ID = id
	updated, err := h.service.UpdateTodo(todo)
	if err != nil {
		pkg.JSONError(w, http.StatusInternalServerError, "Failed update todo")
		return
	}

	pkg.JSONSuccess(w, http.StatusOK, "Success update todo", updated, nil)
}

// DELETE /api/todos/:id # Delete todo
func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.service.DeleteTodo(id); err != nil {
		pkg.JSONError(w, http.StatusInternalServerError, "Failed delete todo")
		return
	}

	pkg.JSONSuccess(w, http.StatusOK, "Success delete todo", nil, nil)
}

// PATCH /api/todos/:id/complete # Toggle completion status
func (h *TodoHandler) UpdateComplete(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		pkg.JSONError(w, http.StatusBadRequest, "Invalid todo ID")
		return
	}

	var input struct {
		IsCompleted bool `json:"is_completed"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		pkg.JSONError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	todo = models.Todo{
		ID:          id,
		IsCompleted: input.IsCompleted,
	}

	updated, err := h.service.UpdateTodoComplete(todo)
	if err != nil {
		pkg.JSONError(w, http.StatusInternalServerError, "Failed update todo")
		return
	}

	pkg.JSONSuccess(w, http.StatusOK, "Success toogle completion todo", updated, nil)
}
