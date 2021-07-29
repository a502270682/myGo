package cron

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"time"
)

var cronCtl Ctl

type Ctl struct {
	C *cron.Cron
}

func init() {
	cronCtl.C = cron.New(cron.WithSeconds(), cron.WithLocation(time.Local))
}

type SchoolJob struct {
	JobName string
}

func (s *SchoolJob) Run() {
	// to something
}

func StartCron() {
	if cronCtl.C != nil {
		ctl := cronCtl.C
		schoolJob := SchoolJob{
			JobName: "schoolJob",
		}
		// every 1:00 am
		_, err := ctl.AddJob("0 0 1 * * *", &schoolJob)
		if err != nil {
			panic(errors.Wrap(err, fmt.Sprintf("failed to start %s", schoolJob.JobName)))
		}
		ctl.Start()
	} else {
		panic("cron not init")
	}
}

func StopCron() {
	if cronCtl.C != nil {
		cronCtl.C.Stop()
	}
}
