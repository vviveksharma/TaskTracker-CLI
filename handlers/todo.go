package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vviveksharma/taskTracker-CLI/common"
	"github.com/vviveksharma/taskTracker-CLI/models"
)

func (h *Handler) CreateTask(c *fiber.Ctx) error {
	var requestBody *models.CreateTaskRequest
	err := c.BodyParser(&requestBody)
	if err != nil {
		log.Println("Error in parsing the request Body")
		return c.Status(fiber.StatusBadGateway).JSON(&common.ServiceResponse{
			Code:    502,
			Message: "error while parsing the requestBody: " + err.Error(),
		})
	}
	if requestBody.Description == "" || requestBody.Name == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&common.ServiceResponse{
			Code:    400,
			Message: "Name and Description are mandatory fields",
		})
	}
	resp, err := h.TaskHandler.CreateTask(requestBody)
	if err != nil {
		return c.JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *Handler) ListTask(c *fiber.Ctx) error {
	resp, err := h.TaskHandler.ListAllTask()
	if err != nil {
		return c.JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(&common.ServiceResponse{
		Code: 200,
		Data: resp,
	})
}

func (h *Handler) UpdateTask(c *fiber.Ctx) error {
	id := c.Query("taskId")
	if id == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&common.ServiceResponse{
			Code:    400,
			Message: "id in the flags are required",
		})
	}
	resp, err := h.TaskHandler.UpdateStatusTask(uuid.MustParse(id))
	if err != nil {
		return c.JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *Handler) DeleteTask(c *fiber.Ctx) error {
	id := c.Query("taskId")
	if id == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&common.ServiceResponse{
			Code:    400,
			Message: "id in the flags are required",
		})
	}
	resp, err := h.TaskHandler.DeleteTask(uuid.MustParse(id))
	if err != nil {
		return c.JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
