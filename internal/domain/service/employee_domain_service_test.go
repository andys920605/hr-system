package service

import (
	"context"
	"errors"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"github.com/andys920605/hr-system/internal/domain/model/employee"
	"github.com/andys920605/hr-system/internal/mock"
	"github.com/andys920605/hr-system/internal/north/message"
	"github.com/andys920605/hr-system/pkg/logging"
	"github.com/andys920605/hr-system/pkg/snowflake"
)

func TestEmployeeDomainService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EmployeeDomainService Suite")
}

var _ = Describe("EmployeeDomainService", func() {
	var (
		ctrl             *gomock.Controller
		employeeRepoMock *mock.MockEmployeeRepository
		service          *EmployeeDomainService
		ctx              context.Context
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		employeeRepoMock = mock.NewMockEmployeeRepository(ctrl)
		service = NewEmployeeDomainService(logging.New(), employeeRepoMock)
		ctx = context.Background()
		snowflake.Init(logging.New())
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Create", func() {
		It("should create and save a new employee", func() {
			cmd := message.CreateEmployeeCommand{
				Name:     "John Doe",
				Email:    "johndoe@example.com",
				Phone:    "123456789",
				Address:  "123 Street",
				Position: "engineer",
				Level:    "L3",
			}

			employeeRepoMock.EXPECT().Save(ctx, gomock.Any()).Return(nil)

			err := service.Create(ctx, cmd)
			Expect(err).To(BeNil())
		})

		It("should return error when repository fails to save", func() {
			cmd := message.CreateEmployeeCommand{
				Name:     "John Doe",
				Email:    "johndoe@example.com",
				Phone:    "123456789",
				Address:  "123 Street",
				Position: "engineer",
				Level:    "L3",
			}

			employeeRepoMock.EXPECT().Save(ctx, gomock.Any()).Return(errors.New("db error"))

			err := service.Create(ctx, cmd)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("save"))
		})
	})

	Describe("GetActiveEmployeeByID", func() {
		It("should return employee when employee is active", func() {
			emp := &employee.Employee{Status: employee.Active}
			employeeRepoMock.EXPECT().GetByID(ctx, gomock.Any()).Return(emp, nil)

			query := message.GetActiveEmployeeByIDQuery{ID: 1}
			result, err := service.GetActiveEmployeeByID(ctx, query)

			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
		})

		It("should return error when employee has resigned", func() {
			emp := &employee.Employee{Status: employee.Resigned}
			employeeRepoMock.EXPECT().GetByID(ctx, gomock.Any()).Return(emp, nil)

			query := message.GetActiveEmployeeByIDQuery{ID: 1}
			result, err := service.GetActiveEmployeeByID(ctx, query)

			Expect(err).ToNot(BeNil())
			Expect(result).To(BeNil())
		})

		It("should return error when repository fails to get employee", func() {
			employeeRepoMock.EXPECT().GetByID(ctx, gomock.Any()).Return(nil, errors.New("db error"))

			query := message.GetActiveEmployeeByIDQuery{ID: 1}
			result, err := service.GetActiveEmployeeByID(ctx, query)

			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("get by id"))
			Expect(result).To(BeNil())
		})
	})
})
