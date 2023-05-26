package crontab

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"woods/pkg/util"
)

var crontab *cron.Cron

type Task1 struct {
	Name string
}

func (t1 *Task1) Run() {
	fmt.Println("task1", t1.Name)
}

type Task2 struct {
	Name string
}

func (t2 *Task2) Run() {
	fmt.Println("task2", t2.Name)
}
func BeginCrontab() {
	crontab = cron.New(cron.WithSeconds())
	task := func() {
		currentTime := util.Now()
		fmt.Println("hello world", currentTime)
		today := util.GetDateTimestamp(currentTime)
		yesterday := util.GetDateTimestamp(currentTime.AddDate(0, 0, -1))
		fmt.Println("今天：", today, "昨天：", yesterday)
	}
	//spec := "0 0 1 * * ?"
	spec := "*/5 * * * * ?"
	crontab.AddFunc(spec, task)
	crontab.Start()
}
