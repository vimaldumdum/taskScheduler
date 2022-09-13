package services

import (
	"fmt"
	"time"

	"github.com/vimaldumdum/taskScheduler/models"
)

var taskSchedulerObj TaskScheduler

type TaskScheduler interface {
	ScheduleTaskAt(at time.Time)
	ScheduleRecurringTask(interval int)
	ShowCurrentTime()
	StopRecurringTask(id string)
}

type TaskSchedulerImpl struct {
	tasks map[string]models.Task
}

func NewTaskScheduler() TaskScheduler {
	if taskSchedulerObj == nil {
		taskSchedulerObj = &TaskSchedulerImpl{
			tasks: make(map[string]models.Task),
		}
	}
	return taskSchedulerObj
}

func (scheduler *TaskSchedulerImpl) ScheduleTaskAt(at time.Time) {
	after := at.Sub(time.Now())
	channel := make(chan bool)
	t := models.NewTask(false, at, 0, channel)
	scheduler.tasks[t.GetId()] = t
	fmt.Println("inside scheduler at =", at.Format("2006-01-02 15:04"), ", now", time.Now().Format("2006-01-02 15:04"))
	go func() {
		fmt.Println("Before sleep, time left", after.Seconds())
		time.Sleep(after)
		printTask(t)
	}()
}

func (scheduler *TaskSchedulerImpl) StopRecurringTask(id string) {
	close(scheduler.tasks[id].GetChannel())
	fmt.Println(id, "Stopped")
}

func (scheduler *TaskSchedulerImpl) ScheduleRecurringTask(interval int) {
	channel := make(chan bool)
	t := models.NewTask(true, time.Now(), interval, channel)
	scheduler.tasks[t.GetId()] = t
	go scheduleInterval(t.GetInterval(), t.GetChannel(), t)
}

func (scheduler *TaskSchedulerImpl) ShowCurrentTime() {
	fmt.Println("Current time is: ", time.Now())
}

func scheduleInterval(interval int, c chan bool, t models.Task) {

	for true {
		select {
		case _, open := <-c:
			if !open {
				return
			}
		default:
			printTask(t)
			time.Sleep(time.Second * time.Duration(interval))
		}
	}
}

func printTask(t models.Task) {
	fmt.Println("Task ", t.GetId(), "executed at: ", time.Now().Local())
}

func GetTimeFromString(t string) time.Time {
	ti, _ := time.ParseInLocation("2006-01-02 15:04", t, time.Now().Local().Location())
	return ti
}
