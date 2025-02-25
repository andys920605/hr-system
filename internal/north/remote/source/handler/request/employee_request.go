package request

import (
	"errors"
	"regexp"
)

type CreateEmployeeRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Level    string `json:"level" binding:"required"`
	Position string `json:"position" binding:"required"`
}

type EmployeeResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Level    string `json:"level"`
	Position string `json:"position"`
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func (r *CreateEmployeeRequest) Valid() error {
	if !emailRegex.MatchString(r.Email) {
		return errors.New("invalid email format")
	}
	return nil
}
