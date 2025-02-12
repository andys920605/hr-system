package repository

import (
	"context"

	"github.com/andys920605/hr-system/internal/domain/model/employee"
)

type EmployeeRepository interface {
	Save(ctx context.Context, employee *employee.Employee) error
}
