package handlers

import "github.com/vviveksharma/taskTracker-CLI/services"

type Handler struct {
	TaskHandler services.ITodo
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Task(sr services.ITodo) *Handler {
	h.TaskHandler = sr
	return h
}
