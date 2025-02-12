package mysql

import (
	"context"

	"github.com/andys920605/hr-system/internal/domain/model/employee"
	"github.com/andys920605/hr-system/pkg/errors"
	"github.com/andys920605/hr-system/pkg/mysqlx"
)

type EmployeeDao struct {
	client *mysqlx.Client
}

func NewEmployeeDao(client *mysqlx.Client) *EmployeeDao {
	return &EmployeeDao{
		client: client,
	}
}

func (d *EmployeeDao) Create(ctx context.Context, emp *employee.Employee) error {
	dto := Employee{
		ID: ID(emp.ID),
		Info: PersonalInfo{
			Name:    emp.Info.Name,
			Email:   emp.Info.Email,
			Phone:   emp.Info.Phone,
			Address: emp.Info.Address,
		},
		Position:  int(emp.Position),
		JobLevel:  int(emp.JobLevel),
		Status:    int(emp.Status),
		CreatedAt: emp.CreatedAt,
		UpdatedAt: emp.UpdatedAt,
	}

	if err := d.client.Create(dto).Error; err != nil {
		return errors.Wrap(err, "create")
	}
	return nil
}
