package datamodels

import "time"

type Todo struct {
	ID          int64     `gorm:"column:ID;PRIMARY_KEY" json:"id"`
	Value       string    `gorm:"column:value" json:"value"`
	IsCompleted bool      `gorm:"column:is_completed" json:"is_completed"`
	CreatedAT   time.Time `gorm:"column:created_at" json:"created_at"`
	IfRemoved   bool      `gorm:"column:if_removed" json:"if_removed"`
	FounderID   int64     `gorm:"column:founder_id" json:"founder_id"`
	OwnerID     int64     `gorm:"column:owner_id" json:"owner_id"`
	Sort        int64     `gorm:"sort" json:"sort"`
}
