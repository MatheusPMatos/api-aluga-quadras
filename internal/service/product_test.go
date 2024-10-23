package service

import (
	"testing"
)

func TestCreateSchedule(t *testing.T) {
	schedule := createSchedule()
	for _, i := range schedule {
		if i.Weekday == 1 {
			println(i.FinalTime.String())

		}
	}

}
