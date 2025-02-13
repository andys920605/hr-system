package mysql

import (
	"context"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

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

	if err := d.client.Table("employee").Create(dto).Error; err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return errors.EmailAlreadyExists.New("email already exists")
		}
		return errors.Wrap(err, "create")
	}

	return nil
}

func (d *EmployeeDao) GetByID(ctx context.Context, id int64) (*employee.Employee, error) {
	var dto Employee
	if err := d.client.Table("employee").Where("id = ?", id).First(&dto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound.New("employee not found")
		}
		return nil, errors.Wrap(err, "failed to get employee by ID")
	}

	emp := &employee.Employee{
		ID: employee.ID(dto.ID),
		Info: employee.PersonalInfo{
			Name:    dto.Info.Name,
			Email:   dto.Info.Email,
			Phone:   dto.Info.Phone,
			Address: dto.Info.Address,
		},
		Position:  employee.Position(dto.Position),
		JobLevel:  employee.JobLevel(dto.JobLevel),
		Status:    employee.Status(dto.Status),
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}

	return emp, nil
}
