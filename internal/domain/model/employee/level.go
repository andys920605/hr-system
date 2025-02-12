package employee

import (
	"fmt"
	"strconv"
)

type JobLevel int

const (
	Level1 JobLevel = 1
	Level2 JobLevel = 2
	Level3 JobLevel = 3
	Level4 JobLevel = 4
	Level5 JobLevel = 5
)

func (l JobLevel) String() string {
	return fmt.Sprintf("L%d", l)
}

func ParseLevel(s string) (JobLevel, error) {
	if len(s) < 2 || (s[0] != 'L' && s[0] != 'l') {
		return 0, fmt.Errorf("invalid job level format: %s", s)
	}
	num, err := strconv.Atoi(s[1:])
	if err != nil {
		return 0, fmt.Errorf("invalid number in job level: %s", s)
	}
	level := JobLevel(num)
	if level < Level1 || level > Level5 {
		return 0, fmt.Errorf("job level out of range: %d", num)
	}
	return level, nil
}
