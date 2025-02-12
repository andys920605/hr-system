package employee

import (
	"context"

	"golang.org/x/sync/singleflight"

	"github.com/andys920605/hr-system/internal/domain/model/employee"
	"github.com/andys920605/hr-system/internal/south/adapter/repository/dao/employee/mysql"
	"github.com/andys920605/hr-system/internal/south/port/repository"
	"github.com/andys920605/hr-system/pkg/errors"
)

var _ repository.EmployeeRepository = (*EmployeeRepository)(nil)

type EmployeeRepository struct {
	sg          singleflight.Group
	employeeDao *mysql.EmployeeDao
}

func NewEmployeeRepository(
	employeeDao *mysql.EmployeeDao,
) *EmployeeRepository {
	return &EmployeeRepository{
		sg:          singleflight.Group{},
		employeeDao: employeeDao,
	}
}

func (r *EmployeeRepository) Save(ctx context.Context, employee *employee.Employee) error {
	if employee.AggregateRoot.IsNew() {
		if err := r.employeeDao.Create(ctx, employee); err != nil {
			return errors.Wrap(err, "create")
		}
	}

	// todo: employee.IsDomainEventsNotEmpty can trigger events

	return nil
}
