package repositories

import (
	"errors"
	"fmt"
	"strings"

	"github.com/naufalilyasa/todolist-be-golang/internal/models"
	"gorm.io/gorm"
)

type TodoRepository interface {
	FindAll() ([]models.Todo, error)
	FindById(id int) (*models.Todo, error)
	Create(todo models.Todo) (*models.Todo, error)
	Update(todo models.Todo) (*models.Todo, error)
	Delete(id int) error
	UpdateComplete(todo models.Todo) (*models.Todo, error)
	FindAllWithFilters(
		page, limit int,
		search string,
		status, priority, category *string,
	) ([]models.Todo, int64, error)
}

type todoRepositoryDB struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepositoryDB{db}
}

func (r *todoRepositoryDB) FindAll() ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Preload("Category").Find(&todos).Error
	return todos, err
}

func (r *todoRepositoryDB) FindById(id int) (*models.Todo, error) {
	var todo models.Todo
	err := r.db.Preload("Category").First(&todo, id).Error
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *todoRepositoryDB) Create(todo models.Todo) (*models.Todo, error) {
	err := r.db.Create(&todo).Error
	return &todo, err
}

func (r *todoRepositoryDB) Update(todo models.Todo) (*models.Todo, error) {
	err := r.db.Save(&todo).Error
	return &todo, err
}

func (r *todoRepositoryDB) Delete(id int) error {
	return r.db.Delete(&models.Todo{}, id).Error
}

func (r *todoRepositoryDB) UpdateComplete(todo models.Todo) (*models.Todo, error) {
	tx := r.db.Model(&models.Todo{}).
		Where("id = ?", todo.ID).
		Update("is_completed", todo.IsCompleted)

	if tx.Error != nil {
		fmt.Println("Update error:", tx.Error)
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		fmt.Println("Update error: no rows affected")
		return nil, errors.New("todo not found")
	}

	var updated models.Todo
	if err := r.db.First(&updated, todo.ID).Error; err != nil {
		fmt.Println("First error:", err)
		return nil, err
	}

	return &updated, nil
}

func (r *todoRepositoryDB) FindAllWithFilters(
	page, limit int,
	search string,
	status, priority, category *string,
) ([]models.Todo, int64, error) {
	var todos []models.Todo
	var total int64

	query := r.db.Model(&models.Todo{}).Preload("Category")

	fmt.Println(search)
	// Search
	if search != "" {
		query = query.Where("title ILIKE ?", "%"+search+"%")
	}

	// Filter status
	if *status != "" {
		var isCompleted bool
		switch strings.ToLower(*status) {
		case "completed":
			isCompleted = true
		case "pending":
			isCompleted = false
		default:

		}
		query = query.Where("is_completed = ?", isCompleted)
	}

	// Filter priority
	if priority != nil {
		query = query.Where("priority = ?", *priority)
	}

	// Filter category
	if category != nil {
		query = query.Where("category_id = ?", *category)
	}

	// Count total filtered data
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (page - 1) * limit
	if err := query.
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&todos).Error; err != nil {
		return nil, 0, err
	}

	return todos, total, nil
}
