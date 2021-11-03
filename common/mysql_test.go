package common

import (
	"testing"
	"time"
	"todo/datamodels"
)

func TestNewMysqlConn(t *testing.T) {
	db, err := NewMysqlConn()
	if err != nil {
		panic("无法连接数据库 mys")
	}
	t1 := datamodels.Todo{
		Value:       "今天的第一件事",
		IsCompleted: false,
		CreatedAT:   time.Now(),
		IfRemoved:   false,
		FounderID:   0,
		OwnerID:     0,
	}
	db.Create(&t1)
}
