package utils

import (
	"hmc/utils"
	"time"
)

type TimedBlock struct {
	guid    string
	Name    string
	Started time.Time
	Elapsed time.Duration
}

type timeKeeper map[string]*TimedBlock

var timers = make(timeKeeper, 0)

func StartTimmer(name string) *TimedBlock {
	timer := &TimedBlock{
		guid:    utils.GenerateUUID(),
		Name:    name,
		Started: time.Now(),
	}

	timers[timer.guid] = timer

	return timer
}

func (timer *TimedBlock) StopTimmer() {
	timer.Elapsed = time.Since(timer.Started)
}

func (timer *TimedBlock) ReportElapsed() {
	Tag("%v : %v\n", timer.Name, YellowText(timer.Elapsed))

	t := timers[timer.guid]
	if t == nil {
		return
	}
	delete(timers, t.guid)
}
