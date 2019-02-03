package usecase

import (
	"go-todo/models"
	"go-todo/todo"
)

type TodoUseCase struct {
	todorepository todo.Repository
}

func NewTodoUseCase(r todo.Repository) todo.UseCase {
	return &TodoUseCase{r}
}

func (tu *TodoUseCase) FindById(id int) (*models.Todo, error) {
	return tu.todorepository.FindById(id)
}

func (tu *TodoUseCase) Save(t *models.Todo) {
	tu.todorepository.Save(t)
}

func (tu *TodoUseCase) Update(id int, t *models.Todo) error {
	return tu.todorepository.Update(id, t)
}

func (tu *TodoUseCase) DeleteById(id int) {
	tu.todorepository.DeleteById(id)
}
