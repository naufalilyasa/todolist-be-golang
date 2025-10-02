package services

import (
	"github.com/naufalilyasa/todolist-be-golang/internal/models"
	"github.com/naufalilyasa/todolist-be-golang/internal/repositories"
)

type CategoryService interface {
	GetCategories() ([]models.Category, error)
	CreateCategory(category models.Category) (*models.Category, error)
	UpdateCategory(category models.Category) (*models.Category, error)
	DeleteCategory(id int) error
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) GetCategories() ([]models.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) CreateCategory(todo models.Category) (*models.Category, error) {
	return s.repo.Create(todo)
}

func (s *categoryService) UpdateCategory(todo models.Category) (*models.Category, error) {
	return s.repo.Update(todo)
}

func (s *categoryService) DeleteCategory(id int) error {
	return s.repo.Delete(id)
}
