package app

import (
	"github.com/robfig/cron"
)

func Scheduler() {
	c := cron.New()
	c.AddFunc(Config.Schedule, func() {
		NetatmoCh <- "ping"
	})
	c.Start()
}
