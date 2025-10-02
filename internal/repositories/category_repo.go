package repositories

import (
	"github.com/naufalilyasa/todolist-be-golang/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll() ([]models.Category, error)
	Create(category models.Category) (*models.Category, error)
	Update(category models.Category) (*models.Category, error)
	Delete(id int) error
}

type categoryRepositoryDB struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepositoryDB{db}
}

func (r *categoryRepositoryDB) FindAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepositoryDB) Create(category models.Category) (*models.Category, error) {
	err := r.db.Create(&category).Error
	return &category, err
}

func (r *categoryRepositoryDB) Update(category models.Category) (*models.Category, error) {
	err := r.db.Save(&category).Error
	return &category, err
}

func (r *categoryRepositoryDB) Delete(id int) error {
	return r.db.Delete(&models.Category{}, id).Error
}
