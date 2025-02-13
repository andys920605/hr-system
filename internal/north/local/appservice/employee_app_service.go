package appservice

import (
	"context"

	"github.com/andys920605/hr-system/internal/domain/service"
	"github.com/andys920605/hr-system/internal/north/message"
	"github.com/andys920605/hr-system/internal/north/remote/source/handler/request"
	"github.com/andys920605/hr-system/pkg/errors"
	"github.com/andys920605/hr-system/pkg/logging"
)

type EmployeeAppService struct {
	logging               *logging.Logging
	employeeDomainService *service.EmployeeDomainService
}

func NewEmployeeAppService(
	logging *logging.Logging,
	employeeDomainService *service.EmployeeDomainService,
) *EmployeeAppService {
	return &EmployeeAppService{
		logging:               logging,
		employeeDomainService: employeeDomainService,
	}
}

func (s *EmployeeAppService) Create(ctx context.Context, cmd message.CreateEmployeeCommand) error {
	if err := s.employeeDomainService.Create(ctx, cmd); err != nil {
		return errors.Wrap(err, "create")
	}
	return nil
}

func (s *EmployeeAppService) GetActiveEmployeeByID(ctx context.Context, query message.GetActiveEmployeeByIDQuery) (*request.EmployeeResponse, error) {
	employee, err := s.employeeDomainService.GetActiveEmployeeByID(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "get active employee by id")
	}
	return &request.EmployeeResponse{
		ID:       employee.ID.String(),
		Name:     employee.Info.Name,
		Email:    employee.Info.Email,
		Phone:    employee.Info.Phone,
		Address:  employee.Info.Address,
		Level:    employee.JobLevel.String(),
		Position: employee.Position.String(),
	}, nil
}
