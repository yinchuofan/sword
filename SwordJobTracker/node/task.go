package node

import (
	"strings"

	"github.com/satori/go.uuid"
)

type Task struct {
	id     string
	status int
	para   string
}

func NewTask(id string, status int, para string) *Task {
	return &Task{
		id:     id,
		status: status,
		para:   para,
	}
}

func NewTaskID() (taskID string) {
	taskID = uuid.NewV4().String()
	taskID = strings.Replace(taskID, "-", "", -1)
	return taskID
}
