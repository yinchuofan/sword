package task

type Task struct {
	taskID    string
	taskPara  string
	processID int
}

func NewTask(taskID string, taskPara string, processID int) *Task {
	return &Task{
		taskID:    taskID,
		taskPara:  taskPara,
		processID: processID,
	}
}
