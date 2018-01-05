package config

import (
	"sync"
)

const (
	NodeConfigPath string = "./config/node.conf"
)

var (
	// Node config : key is nodeHost, value is taskLimit
	nodeConfigMap      map[string]string
	nodeConfigMapMutex = new(sync.RWMutex)
)

func GetNodeConfigMap() map[string]string {
	if nodeConfigMap == nil {
		return map[string]string{}
	}
	return nodeConfigMap
}

func LoadNodeConfig() (err error) {
	nodeConfigMapMutex.Lock()
	defer nodeConfigMapMutex.Unlock()

	nodeConfigMap, err = loadConfig(NodeConfigPath)
	if err != nil {
		return err
	}

	return nil
}

func SaveNodeConfig() (err error) {
	return saveConfig(NodeConfigPath, nodeConfigMap)
}

func SetNodeConfig(nodeHost string, taskLimit string) (err error) {
	nodeConfigMapMutex.Lock()
	defer nodeConfigMapMutex.Unlock()

	nodeConfigMap[nodeHost] = taskLimit

	return SaveNodeConfig()
}

func GetNodeConfig(nodeHost string) (taskLimit string) {
	nodeConfigMapMutex.RLock()
	defer nodeConfigMapMutex.RUnlock()

	if elem, ok := nodeConfigMap[nodeHost]; ok {
		taskLimit = elem
		return taskLimit
	}

	return ""
}

func AddNodeConfig(nodeHost string, taskLimit string) (err error) {
	nodeConfigMapMutex.Lock()
	defer nodeConfigMapMutex.Unlock()

	nodeConfigMap[nodeHost] = taskLimit

	return SaveNodeConfig()
}

func RemoveNodeConfig(nodeHost string) (err error) {
	nodeConfigMapMutex.Lock()
	defer nodeConfigMapMutex.Unlock()

	if _, ok := nodeConfigMap[nodeHost]; ok {
		delete(nodeConfigMap, nodeHost)
	}

	return SaveNodeConfig()
}
