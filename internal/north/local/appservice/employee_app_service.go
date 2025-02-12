package appservice

import (
	"context"

	"github.com/andys920605/hr-system/internal/domain/service"
	"github.com/andys920605/hr-system/internal/north/message"
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
