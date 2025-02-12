package employee

type Status int

const (
	Resigned Status = iota
	Active
)

func (t Status) Int64() int64 {
	return int64(t)
}

func (t Status) String() string {
	switch t {
	case Resigned:
		return "Resigned"
	case Active:
		return "Active"
	default:
		return "Unknown"
	}
}
