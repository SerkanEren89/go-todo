package todo

import "go-todo/models"

type Repository interface {
	FindById(id int) (*models.Todo, error)
	Save(t *models.Todo)
	Update(id int, t *models.Todo) error
	DeleteById(id int)
}
