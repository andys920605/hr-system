package employee

import "fmt"

type Position int

const (
	PositionEngineer Position = iota + 1 // Engineer
	PositionManager                      // Manager
	PositionAdmin                        // Admin
)

var posMap = map[string]Position{
	"engineer": PositionEngineer,
	"manager":  PositionManager,
	"admin":    PositionAdmin,
}

var posReverseMap = map[Position]string{
	PositionEngineer: "engineer",
	PositionManager:  "manager",
	PositionAdmin:    "admin",
}

func ParsePosition(s string) (Position, error) {
	if pos, ok := posMap[s]; ok {
		return pos, nil
	}
	return 0, fmt.Errorf("invalid position: %s", s)
}

func (p Position) String() string {
	if str, ok := posReverseMap[p]; ok {
		return str
	}
	return "unknown"
}
