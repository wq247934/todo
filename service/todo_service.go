package service

import (
	"todo/datamodels"
	"todo/repositories"
)

type ITodoService interface {
	FindByOwnerID(ownerID int64) []datamodels.Todo
}
type todoService struct {
	todoRepository repositories.ITodoRepository
}

func (t todoService) FindByOwnerID(ownerID int64) []datamodels.Todo {
	return t.todoRepository.Select(ownerID)
}

func NewTodoService(r repositories.ITodoRepository) ITodoService {
	return todoService{todoRepository: r}
}
