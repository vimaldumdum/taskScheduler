package models

import (
	"fmt"
	"time"
)

type Task struct {
	id        string
	recur     bool
	performAt time.Time
	interval  int
	channel   chan bool
}

var taskCount = 1

func NewTask(recur bool, perform time.Time, interval int, c chan bool) Task {
	t := Task{
		recur:     recur,
		performAt: perform,
		interval:  interval,
		channel:   c,
	}
	t.id = fmt.Sprint("T", taskCount)
	taskCount++
	return t
}

func (task Task) GetId() string {
	return task.id
}

func (task Task) GetRecur() bool {
	return task.recur
}

func (task Task) GetPerform() time.Time {
	return task.performAt
}

func (task Task) GetInterval() int {
	return task.interval
}

func (task Task) GetChannel() chan bool {
	return task.channel
}
