package service

import (
	"context"

	"github.com/andys920605/hr-system/internal/domain/model/employee"
	"github.com/andys920605/hr-system/internal/north/message"
	"github.com/andys920605/hr-system/internal/south/port/repository"
	"github.com/andys920605/hr-system/pkg/errors"
	"github.com/andys920605/hr-system/pkg/logging"
)

type EmployeeDomainService struct {
	logging            *logging.Logging
	employeeRepository repository.EmployeeRepository
}

func NewEmployeeDomainService(
	logging *logging.Logging,
	employeeRepository repository.EmployeeRepository,
) *EmployeeDomainService {
	return &EmployeeDomainService{
		logging:            logging,
		employeeRepository: employeeRepository,
	}
}

func (s *EmployeeDomainService) Create(ctx context.Context, cmd message.CreateEmployeeCommand) error {
	info := employee.NewPersonalInfo(cmd.Name, cmd.Email, cmd.Phone, cmd.Address)
	position, err := employee.ParsePosition(cmd.Position)
	if err != nil {
		return errors.Wrap(err, "parse position")
	}
	level, err := employee.ParseLevel(cmd.Level)
	if err != nil {
		return errors.Wrap(err, "parse level")
	}
	employee := employee.NewEmployee(employee.NewID(), info, position, level)

	if err := s.employeeRepository.Save(ctx, employee); err != nil {
		return errors.Wrap(err, "save")
	}

	return nil
}
