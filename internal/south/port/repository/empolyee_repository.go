package repository

import (
	"context"

	"github.com/andys920605/hr-system/internal/domain/model/employee"
)

//go:generate mockgen -destination=../../../mock/iemployee_mock_repository.go -package=mock github.com/andys920605/hr-system/internal/south/port/repository EmployeeRepository
type EmployeeRepository interface {
	Save(ctx context.Context, employee *employee.Employee) error
	GetByID(ctx context.Context, id int64) (*employee.Employee, error)
}
