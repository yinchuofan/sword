package task

import (
	"errors"
	"sync"
	"time"

	"sword/SwordJobTracker/log"
)

var (
	MasterHost = ""
	NodeHost   = ""
	TaskLimit  = 0
	TaskPath   = "./"
)

var (
	//Key is task id, value is task which in running status
	taskMap      = make(map[string]*Task)
	taskMapMutex = new(sync.RWMutex)
)

func Setup(master, node string, limit int, taskPath string) {
	MasterHost = master
	NodeHost = node
	TaskLimit = limit
	TaskPath = taskPath

	go register()
}

func register() {
	for {
		if err := RegisterNode(); err != nil {
			log.Root.Error("Register node failed, continue to register after 1 minitue")
			time.Sleep(time.Minute * 1)
		} else {
			break
		}
	}
}

func GetTaskCapacity() (capacity int) {
	taskMapMutex.RLock()
	defer taskMapMutex.RUnlock()

	return TaskLimit - len(taskMap)
}

func StartTask(taskID, taskPara string) (err error) {
	processID := getTaskProcessID(taskID)
	if processID > 0 {
		log.Root.Info("Task is running. TaskID: %v", taskID)
		return nil
	}

	if capacity := GetTaskCapacity(); capacity <= 0 {
		log.Root.Error("Lack of computing resources. TaskID: %v", taskID)
		return errors.New("lack of computing resources")
	}

	processID, err = StartTaskProcess(taskID, taskPara)
	if err != nil {
		log.Root.Error("Start task process error. TaskID: %v", taskID)
		return err
	}

	log.Root.Info("Start task process succeed. TaskID: %v, ProcessID: %v", taskID, processID)
	task := NewTask(taskID, taskPara, processID)
	addTask(task)

	return nil
}

func StopTask(taskID string) (err error) {
	processID := getTaskProcessID(taskID)
	if processID < 0 {
		log.Root.Error("No task process. TaskID: %v", taskID)
		return errors.New("no task process")
	}

	err = StopTaskProcess(processID)
	if err != nil {
		log.Root.Error("Stop task process error. ProcessID: %v", taskID, processID)
		return err
	}

	log.Root.Info("Stop task process succeed. TaskID: %v, ProcessID: %v", taskID, processID)
	return nil
}

func getTaskProcessID(taskID string) (processID int) {
	taskMapMutex.RLock()
	defer taskMapMutex.RUnlock()

	if elem, ok := taskMap[taskID]; ok {
		return elem.processID
	}

	return -1
}

func addTask(t *Task) {
	taskMapMutex.RLock()
	defer taskMapMutex.RUnlock()

	taskMap[t.taskID] = t
}

func removeTask(taskID string) {
	taskMapMutex.RLock()
	defer taskMapMutex.RUnlock()

	if _, ok := taskMap[taskID]; ok {
		delete(taskMap, taskID)
	}
}

func GetAllRunningTasks() (taskIDList []string, err error) {
	taskMapMutex.RLock()
	defer taskMapMutex.RUnlock()

	taskIDList = []string{}

	for k, _ := range taskMap {
		taskIDList = append(taskIDList, k)
	}

	return taskIDList, nil
}
