// Package timer is used to create one-off timers for use in benchmarking and
// testing things.
package timer

import (
	"time"

	"github.com/sirupsen/logrus"
)

// Timer is the representation of an individual timer
type Timer struct {
	name      string
	startTime time.Time
	log       *logrus.Entry
}

// Start begins a new timer with name "name" and returns it
func Start(name string) *Timer {
	now := time.Now()
	r := &Timer{
		name:      name,
		startTime: now,
		log: logrus.WithFields(logrus.Fields{
			"name":  name,
			"start": now,
		}),
	}
	r.log.Infof("timer starting")
	return r
}

// End stops the timer and logs the result
func (t *Timer) End() {
	t.log.WithField(`execution_time`, time.Since(t.startTime)).Infof(`timer ending`)
}
