package models

import "github.com/google/uuid"

type CreateTaskResponse struct {
	Message string `json:"message"`
}

type ListTaskResponse struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
}

type UpdateStatusTaskResponse struct {
	Message string `json:"message"`
}

type DeleteTaskResponse struct {
	Message string `json:"message"`
}
