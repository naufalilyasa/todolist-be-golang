package services

import (
	"github.com/naufalilyasa/todolist-be-golang/internal/models"
	"github.com/naufalilyasa/todolist-be-golang/internal/repositories"
)

type TodoService interface {
	GetTodos() ([]models.Todo, error)
	GetTodoById(id int) (*models.Todo, error)
	CreateTodo(todo models.Todo) (*models.Todo, error)
	UpdateTodo(todo models.Todo) (*models.Todo, error)
	DeleteTodo(id int) error
	UpdateTodoComplete(todo models.Todo) (*models.Todo, error)
	FindAllWithPaginationAndSearch(page, limit int, search string) ([]models.Todo, int64, error)
}

type todoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
	return &todoService{repo}
}

func (s *todoService) GetTodos() ([]models.Todo, error) {
	return s.repo.FindAll()
}

func (s *todoService) GetTodoById(id int) (*models.Todo, error) {
	return s.repo.FindById(id)
}

func (s *todoService) CreateTodo(todo models.Todo) (*models.Todo, error) {
	return s.repo.Create(todo)
}

func (s *todoService) UpdateTodo(todo models.Todo) (*models.Todo, error) {
	return s.repo.Update(todo)
}

func (s *todoService) DeleteTodo(id int) error {
	return s.repo.Delete(id)
}

func (s *todoService) UpdateTodoComplete(todo models.Todo) (*models.Todo, error) {
	return s.repo.UpdateComplete(todo)
}

func (s *todoService) FindAllWithPaginationAndSearch(page, limit int, search string) ([]models.Todo, int64, error) {
	return s.repo.FindAllWithPaginationAndSearch(page, limit, search)
}
