package repository

import (
	"github.com/jinzhu/gorm"
	"go-todo/models"
	"go-todo/todo"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) todo.Repository {
	return &TodoRepository{db}
}

func (tr TodoRepository) FindById(id int) (*models.Todo, error) {
	var t models.Todo
	if err := tr.db.Where("id = ?", id).First(&t); err != nil {
		return nil, models.ErrNotFound
	} else {
		return &t, nil
	}
}

func (tr TodoRepository) Save(t *models.Todo) {
	tr.db.Save(t)
}

func (tr TodoRepository) Update(id int, t *models.Todo) error {
	var taskToUpdate models.Todo
	if err := tr.db.Where("id = ?", id).First(&taskToUpdate).Error; err != nil {
		return models.ErrNotFound
	}
	tr.db.Save(t)
	return nil
}

func (tr TodoRepository) DeleteById(id int) {
	var t models.Todo
	tr.db.Where("id = ?", id).Delete(t)
}
