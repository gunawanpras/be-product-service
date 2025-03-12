package timeutil

import (
	"time"
)

type (
	Time interface {
		Now() time.Time
	}

	timeHelper struct{}
)

func (h *timeHelper) Now() time.Time {
	return time.Now()
}

var TimeHelper Time = &timeHelper{}
