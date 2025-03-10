package models

type CreateTaskRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ListTaskRequest struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}
