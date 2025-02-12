package employee

import (
	"time"

	"github.com/andys920605/hr-system/pkg/dddcore"
)

type Employee struct {
	*dddcore.AggregateRoot
	ID        ID
	Info      PersonalInfo
	Position  Position
	JobLevel  JobLevel
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewEmployee(id ID, info PersonalInfo, pos Position, level JobLevel) *Employee {
	now := time.Now().UTC()
	root := dddcore.NewAggregateRoot().SetNew()
	// Implementation of event-driven design based on CDC (Change Data Capture)
	root.AppendDomainEvent(
		NewCreatedEmployeeEvent(),
	)
	return &Employee{
		AggregateRoot: root,
		ID:            id,
		Info:          info,
		Position:      pos,
		JobLevel:      level,
		Status:        Active,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}
