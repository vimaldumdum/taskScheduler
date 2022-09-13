package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/vimaldumdum/taskScheduler/services"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	scheduler := services.NewTaskScheduler()
	scheduler.ShowCurrentTime()

	for true {
		inp, _ := reader.ReadString('\n')
		inp = inp[:len(inp)-1]

		command := strings.Split(inp, " ")

		switch command[0] {
		case "recur":
			interval, _ := strconv.Atoi(command[1])
			scheduler.ScheduleRecurringTask(interval)
			break
		case "at":
			scheduler.ScheduleTaskAt(services.GetTimeFromString(command[1] + " " + command[2]))
			break
		case "stop":
			scheduler.StopRecurringTask(command[1])
			break
		case "exit":
			return
		}
	}
}
