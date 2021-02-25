package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type CronJob struct {
	expr *cronexpr.Expression
	nextTime time.Time
}

func main() {
	var (
		expr *cronexpr.Expression
		err error
		cornJob *CronJob
		scheduleTable map[string]*CronJob
	)

	scheduleTable = make(map[string]*CronJob)
	now := time.Now()
	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	cornJob = &CronJob{
		expr: expr,
		nextTime: expr.Next(now),
	}
	scheduleTable["job1"] = cornJob

	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}
	cornJob = &CronJob{
		expr: expr,
		nextTime: expr.Next(now),
	}
	scheduleTable["job2"] = cornJob

	go func() {
		var (
			now time.Time
			jobName string
			cronJob *CronJob
		)

		for {
			now = time.Now()
			for jobName, cronJob = range scheduleTable {
				if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now){
					go func(jobName string) {
						fmt.Println(jobName)
					}(jobName)
					cronJob.nextTime = cronJob.expr.Next(now)
					fmt.Println(jobName, "next time: ", cronJob.nextTime)
				}
			}
		}
		select {
		case <- time.NewTicker(100 * time.Millisecond).C:
		}
	}()

	time.Sleep(100 * time.Second)
}
