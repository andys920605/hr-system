package employee

import "github.com/andys920605/hr-system/pkg/dddcore"

const (
	CreatedEmployeeEventName dddcore.DomainEventName = "employee_created"
)

func NewCreatedEmployeeEvent() dddcore.DomainEvent {
	return dddcore.NewDomainEvent(CreatedEmployeeEventName)
}
