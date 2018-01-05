package config

import (
	"sync"
)

const (
	TaskConfigPath string = "./config/task.conf"
)

var (
	// Task config : key is taskID, value is nodeHost
	taskConfigMap      map[string]string
	taskConfigMapMutex = new(sync.RWMutex)
)

func LoadTaskConfig() (err error) {
	taskConfigMapMutex.Lock()
	defer taskConfigMapMutex.Unlock()

	taskConfigMap, err = loadConfig(TaskConfigPath)
	if err != nil {
		return err
	}

	return nil
}

func SaveTaskConfig() (err error) {
	return saveConfig(TaskConfigPath, taskConfigMap)
}

func GetTaskConfigMap() (nodeTaskMap map[string][]string) {
	//key is nodeHost, value is taskID list
	nodeTaskMap = make(map[string][]string)
	for taskID, nodeHost := range taskConfigMap {
		if _, ok := nodeTaskMap[nodeHost]; ok {
			nodeTaskMap[nodeHost] = append(nodeTaskMap[nodeHost], taskID)
		} else {
			taskIDList := []string{taskID}
			nodeTaskMap[nodeHost] = taskIDList
		}
	}
	return nodeTaskMap
}

func SetTaskConfig(taskID string, nodeHost string) (err error) {
	taskConfigMapMutex.Lock()
	defer taskConfigMapMutex.Unlock()

	taskConfigMap[taskID] = nodeHost

	return SaveTaskConfig()
}

func GetTaskConfig(taskID string) (nodeHost string) {
	taskConfigMapMutex.RLock()
	defer taskConfigMapMutex.RUnlock()

	if elem, ok := taskConfigMap[taskID]; ok {
		nodeHost = elem
		return nodeHost
	}

	return ""
}

func AddTaskConfig(taskID string, nodeHost string) (err error) {
	taskConfigMapMutex.Lock()
	defer taskConfigMapMutex.Unlock()

	taskConfigMap[taskID] = nodeHost

	return SaveTaskConfig()
}

func RemoveTaskConfig(taskID string) (err error) {
	taskConfigMapMutex.Lock()
	defer taskConfigMapMutex.Unlock()

	if _, ok := taskConfigMap[taskID]; ok {
		delete(taskConfigMap, taskID)
	}

	return SaveTaskConfig()
}
