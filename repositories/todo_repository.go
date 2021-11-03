package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"todo/datamodels"
)

type ITodoRepository interface {
	Select(ownerID int64) []datamodels.Todo
	Insert(todo datamodels.Todo)
}
type todoRepository struct {
	conn *gorm.DB
}

func NewTodoRepository(conn *gorm.DB) ITodoRepository {
	return &todoRepository{conn: conn}
}

func (t todoRepository) Select(ownerID int64) []datamodels.Todo {
	var todos []datamodels.Todo
	err := t.conn.Debug().Where("owner_id=?", ownerID).Find(&todos)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("todos", todos)
	return todos
}

func (t todoRepository) Insert(todo datamodels.Todo) {

}
