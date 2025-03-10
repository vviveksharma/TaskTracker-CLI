package common

import "fmt"

const (
	PENDING   = "PENDING"
	COMPLETED = "COMPLETED"
)

type ServiceResponse struct {
	Code    int
	Message string
	Data    interface{}
}

func (e *ServiceResponse) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Data: %+v", e.Code, e.Message, e.Data)
}
