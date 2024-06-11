package request

import (
	"errors"
	"unicode/utf8"
)

// UpdateEmployeeByEmployeeIDRequest represents the request to update an employee by employee_id
type UpdateEmployeeByEmployeeIDRequest struct {
	EmployeeID string `json:"employee_id"`
	Name       string `json:"name"`
	Position   string `json:"position"`
	Email      string `json:"email"`
	Salary     int    `json:"salary"`
}

// Validate validates the UpdateEmployeeByEmployeeIDRequest
func (req UpdateEmployeeByEmployeeIDRequest) Validate() error {
	if req.EmployeeID == "" {
		return errors.New("employee_id cannot be empty")
	}
	if req.Name == "" {
		return errors.New("name cannot be empty")
	}
	if utf8.RuneCountInString(req.Name) > maxNameLength {
		return errors.New("name cannot exceed 50 characters")
	}
	if req.Email == "" {
		return errors.New("email cannot be empty")
	}
	if !emailRegex.MatchString(req.Email) {
		return errors.New("email is not valid")
	}
	if utf8.RuneCountInString(req.Email) > maxEmailLength {
		return errors.New("email cannot exceed 100 characters")
	}
	if req.Position == "" {
		return errors.New("position cannot be empty")
	}
	if utf8.RuneCountInString(req.Position) > maxPositionLength {
		return errors.New("position cannot exceed 50 characters")
	}
	if req.Salary <= 0 {
		return errors.New("salary must be greater than zero")
	}
	return nil
}
