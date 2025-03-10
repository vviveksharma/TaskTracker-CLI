package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DBTask struct {
	Id          uuid.UUID `gorm:"primaryKey,column:id"`
	Name        string    `gorm:"column:name;not null"`
	Description string    `gorm:"column:description;not null"`
	Status      string    `gorm:"column:status;not null"`
}

func (DBTask) TableName() string {
	return "task_tbl"
}

func (*DBTask) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New().String()
	tx.Statement.SetColumn("Id", uuid)
	return nil
}
