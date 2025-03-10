package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/vviveksharma/taskTracker-CLI/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type IDataBaseService interface {
	InitDB() (*gorm.DB, error)
}

type DataBaseService struct {
	Db *gorm.DB
}

func NewDbRequest() (IDataBaseService, error) {
	return &DataBaseService{}, nil
}

func (db *DataBaseService) InitDB() (*gorm.DB, error) {
	var err error
	dsn := "host=db user=postgres password=password dbname=taskTracker-CLI port=5432 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println("the error while creating the database connection: ", err.Error())
		return nil, err
	}
	db.Db = conn

	//Migrating the databases if not done
	db.Db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	err = db.Db.AutoMigrate(&models.DBTask{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}
	fmt.Println("Database migration completed successfully. Database connection successfull.")

	return conn, nil
}
