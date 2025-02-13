package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/andys920605/hr-system/internal/north/local/appservice"
	"github.com/andys920605/hr-system/internal/north/message"
	"github.com/andys920605/hr-system/internal/north/remote/source/handler/request"
	"github.com/andys920605/hr-system/pkg/errors"
	"github.com/andys920605/hr-system/pkg/http/template_response"
)

type EmployeeHandler struct {
	employeeAppService *appservice.EmployeeAppService
}

func NewEmployeeHandler(employeeAppService *appservice.EmployeeAppService) *EmployeeHandler {
	return &EmployeeHandler{
		employeeAppService: employeeAppService,
	}
}

func (h *EmployeeHandler) Create(c *gin.Context) {
	var req request.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(errors.InvalidRequest.Wrap(err, "invalid request"))
		return
	}
	if err := req.Valid(); err != nil {
		_ = c.Error(errors.InvalidRequest.Wrap(err, "invalid request"))
		return
	}
	cmd := message.CreateEmployeeCommand{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Address:  req.Address,
		Level:    req.Level,
		Position: req.Position,
	}
	err := h.employeeAppService.Create(c.Request.Context(), cmd)
	if err != nil {
		_ = c.Error(errors.InternalServerError.Wrap(err, "internal server error"))
		return
	}

	template_response.OK(nil).To(c, http.StatusOK)
}

func (h *EmployeeHandler) GetActiveEmployeeByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(errors.InvalidRequest.New("invalid id: must be an integer"))
		return
	}
	cmd := message.GetActiveEmployeeByIDQuery{
		ID: id,
	}
	employee, err := h.employeeAppService.GetActiveEmployeeByID(c.Request.Context(), cmd)
	if err != nil {
		_ = c.Error(errors.InternalServerError.Wrap(err, "internal server error"))
		return
	}

	template_response.OK(employee).To(c, http.StatusOK)
}
