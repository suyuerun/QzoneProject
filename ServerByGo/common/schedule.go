package common

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

// 定时任务日程 结构体
type Schedule struct {
	ID          int
	Name        string
	CronStr     string //eg: "*/1 * * * * *"
	Description string //eg: "每隔一秒执行一次"
}

var ScheduleList = make([]*Schedule, 0)

func init() {
	fmt.Println("schedule init")
	location, _ := time.LoadLocation("Asia/Shanghai")
	c := cron.New(
		//打印出 定时任务内部的日志到控制台
		cron.WithLogger(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))),
		//按秒级别 解析 字符串
		cron.WithSeconds(),
		//设置定时任务时区
		cron.WithLocation(location),
	)
	for _, scheduleItem := range ScheduleList {
		//添加定时任务
		entryId, err := c.AddFunc(scheduleItem.CronStr, dispatchCronTask(scheduleItem))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(entryId)
		//协程启动定时任务
		go c.Start()
		//defer c.Stop()
	}

	//select {
	//case <-time.After(time.Second * 3):
	//	fmt.Println("结束定时任务")
	//}
}

/**
根据schedule的id去选择 各种 switch 的case
*/
func dispatchCronTask(schedule *Schedule) func() {
	switch schedule.ID {
	case 1: // 设置logger的配置
		return func() {
			TimerSetLogProperty()
		}
	}
	return func() {

	}
}
