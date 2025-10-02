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

type CategoryHandler struct {
	service services.CategoryService
}

func NewCategoryHandler(service services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetCategories()
	if err != nil {
		pkg.JSONError(w, http.StatusInternalServerError, "Failed to get data categories")
		return
	}

	pkg.JSONSuccess(w, http.StatusOK, "Success get categories", categories, nil)
}

// POST /api/todos # Create new todo
func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		pkg.JSONError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	created, err := h.service.CreateCategory(category)
	if err != nil {
		pkg.JSONError(w, http.StatusInternalServerError, "Failed create category")
	}

	pkg.JSONSuccess(w, http.StatusCreated, "Success created todo", created, nil)
}

// PUT /api/todos/:id # Update todo
func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		pkg.JSONError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	category.ID = id
	updated, err := h.service.UpdateCategory(category)
	if err != nil {
		pkg.JSONError(w, http.StatusInternalServerError, "Failed update category")
		return
	}

	pkg.JSONSuccess(w, http.StatusOK, "Success updated todo", updated, nil)
}

// DELETE /api/todos/:id # Delete todo
func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.service.DeleteCategory(id); err != nil {
		pkg.JSONError(w, http.StatusInternalServerError, "Failed delete category")
		return
	}

	pkg.JSONSuccess(w, http.StatusOK, "Success deleted todo", nil, nil)
}
