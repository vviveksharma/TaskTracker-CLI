package services

import (
	"log"

	"github.com/google/uuid"
	"github.com/vviveksharma/taskTracker-CLI/common"
	"github.com/vviveksharma/taskTracker-CLI/dal"
	"github.com/vviveksharma/taskTracker-CLI/models"
)

type ITodo interface {
	CreateTask(req *models.CreateTaskRequest) (resp *models.CreateTaskResponse, err error)
	UpdateStatusTask(id uuid.UUID) (resp *models.UpdateStatusTaskResponse, err error)
	ListAllTask() (resp []*models.ListTaskResponse, err error)
	DeleteTask(id uuid.UUID) (resp *models.DeleteTaskResponse, err error)
}

type TodoService struct {
	TaskRepo dal.ITask
}

func NewTodoServiceRequest() (ITodo, error) {
	ser := &TodoService{}
	err := ser.SetRepo()
	if err != nil {
		log.Println("error while injecting the repo layer: ", err)
		return nil, err
	}
	return ser, nil
}

func (t *TodoService) SetRepo() error {
	var err error
	todo, err := dal.NewTaskDalRequest()
	if err != nil {
		return err
	}
	t.TaskRepo = todo
	return nil
}

// Task Creatation Service Layer
func (t *TodoService) CreateTask(req *models.CreateTaskRequest) (resp *models.CreateTaskResponse, err error) {
	err = t.TaskRepo.Create(&models.DBTask{
		Name:        req.Name,
		Description: req.Description,
		Status:      common.PENDING,
	})
	if err != nil {
		return nil, &common.ServiceResponse{
			Code:    500,
			Message: "error while creating the new task: " + err.Error(),
		}
	}
	return &models.CreateTaskResponse{
		Message: "Task with Name" + req.Name + " created sucessfully",
	}, nil
}

// Status Update of the task
func (t *TodoService) UpdateStatusTask(id uuid.UUID) (resp *models.UpdateStatusTaskResponse, err error) {
	dbresp, err := t.TaskRepo.FindById(id)
	if err != nil {
		return nil, &common.ServiceResponse{
			Code:    500,
			Message: "error while updating the status of the task with id: " + err.Error(),
		}
	}
	if dbresp.Status == common.COMPLETED {
		return nil, &common.ServiceResponse{
			Code:    409,
			Message: "Task is already completed.",
		}
	}
	err = t.TaskRepo.UpdateTask(id)
	if err != nil {
		return nil, &common.ServiceResponse{
			Code:    500,
			Message: "error while updating the task with id " + id.String() + " " + err.Error(),
		}
	}
	return &models.UpdateStatusTaskResponse{
		Message: "Updated the task with id " + id.String(),
	}, nil
}

// List all Task present
func (t *TodoService) ListAllTask() (resp []*models.ListTaskResponse, err error) {
	dbresp, err := t.TaskRepo.FindAll()
	if err != nil {
		return nil, &common.ServiceResponse{
			Code:    500,
			Message: "error while listing all the task: " + err.Error(),
		}
	}
	for _, task := range dbresp {
		taskDetails := &models.ListTaskResponse{}
		taskDetails.Id = task.Id
		taskDetails.Name = task.Name
		taskDetails.Description = task.Description
		taskDetails.Status = task.Status
		resp = append(resp, taskDetails)
	}

	return resp, nil
}

// Delete the task
func (t *TodoService) DeleteTask(id uuid.UUID) (resp *models.DeleteTaskResponse, err error) {
	dbresp, err := t.TaskRepo.FindById(id)
	log.Println("******************** the error : ", err)
	log.Println("the response : *******************: ", dbresp)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, &common.ServiceResponse{
				Code:    400,
				Message: "record with this id: " + id.String() + " not found",
			}
		}
		return nil, &common.ServiceResponse{
			Code:    500,
			Message: "error while fetching details with the given id: " + id.String(),
		}
	}
	log.Println("Deleting the task with name " + dbresp.Name)
	err = t.TaskRepo.Delete(id)
	if err != nil {
		return nil, &common.ServiceResponse{
			Code:    500,
			Message: "error while deleting the task with id: " + id.String(),
		}
	}
	return &models.DeleteTaskResponse{
		Message: "Task deleted successfully with name "+dbresp.Name,
	}, nil
}
