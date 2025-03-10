package dal

import (
	"log"

	"github.com/google/uuid"
	"github.com/vviveksharma/taskTracker-CLI/common"
	"github.com/vviveksharma/taskTracker-CLI/db"

	"github.com/vviveksharma/taskTracker-CLI/models"
)

type Task struct{}

type ITask interface {
	Create(value *models.DBTask) error
	FindAll() (resp []*models.DBTask, err error)
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*models.DBTask, error)
	UpdateTask(id uuid.UUID) error
}

func NewTaskDalRequest() (ITask, error) {
	return &Task{}, nil
}

func (t *Task) Create(value *models.DBTask) error {
	db, err := db.NewDbRequest()
	if err != nil {
		log.Println("error in creating a DB request")
		return err
	}
	dbConn, err := db.InitDB()
	if err != nil {
		return err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return transaction.Error
	}
	defer transaction.Rollback()
	cerr := transaction.Create(&value)
	if cerr.Error != nil {
		log.Println("error while creating the task: " + cerr.Error.Error())
		return cerr.Error
	}
	transaction.Commit()
	return nil
}

func (t *Task) FindAll() (resp []*models.DBTask, err error) {
	db, err := db.NewDbRequest()
	if err != nil {
		log.Println("error in creating a DB request")
		return nil, err
	}
	dbConn, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	defer transaction.Rollback()
	task := transaction.Find(&resp)
	if task.Error != nil {
		log.Println("error while finding the records: " + task.Error.Error())
	}
	return resp, nil
}

func (ls *Task) FindById(id uuid.UUID) (*models.DBTask, error) {
	db, err := db.NewDbRequest()
	if err != nil {
		log.Println("error in creating a DB request")
		return nil, err
	}
	dbConn, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	defer transaction.Rollback()
	var response *models.DBTask
	loginDetails := transaction.First(&response, &models.DBTask{
		Id: id,
	})
	if loginDetails.Error != nil {
		return nil, loginDetails.Error
	}
	return response, nil
}

func (t *Task) Delete(id uuid.UUID) error {
	db, err := db.NewDbRequest()
	if err != nil {
		log.Println("error in creating a DB request")
		return err
	}
	dbConn, err := db.InitDB()
	if err != nil {
		return err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return transaction.Error
	}
	defer transaction.Rollback()
	del := transaction.Unscoped().Where("id =?", id).Delete(&models.DBTask{
		Id: id,
	})
	if del.Error != nil {
		return del.Error
	}
	transaction.Commit()
	return nil
}

func (t *Task) UpdateTask(id uuid.UUID) error {
	db, err := db.NewDbRequest()
	if err != nil {
		log.Println("error in creating a DB request")
		return err
	}
	dbConn, err := db.InitDB()
	if err != nil {
		return err
	}
	transaction := dbConn.Begin()
	if transaction.Error != nil {
		return transaction.Error
	}
	defer transaction.Rollback()
	update := transaction.Model(&models.DBTask{}).Where("id = ?", id).Update("status", common.COMPLETED)
	if update.Error != nil {
		return update.Error
	}
	transaction.Commit()
	return nil
}
