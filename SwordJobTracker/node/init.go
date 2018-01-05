package node

import (
	"errors"
	"strconv"

	"sword/SwordJobTracker/config"
	"sword/SwordJobTracker/log"
)

func init() {
	err := config.LoadNodeConfig()
	if err != nil {
		log.Root.Error("Load node config error: %v", err.Error())
		panic("Load node config error: " + err.Error())
	}
	log.Root.Info("Load node config succeed")

	err = config.LoadTaskConfig()
	if err != nil {
		log.Root.Error("Load task config error: %v", err.Error())
		panic("Load task config error: " + err.Error())
	}
	log.Root.Info("Load task config succeed")

	err = initNodeAndTask()
	if err != nil {
		log.Root.Error("Init node and task error: %v", err.Error())
		panic("Init node and task error: " + err.Error())
	}
	log.Root.Info("Init node and task succeed")

	//Print node info
	log.Root.Info("Print nodes: %v", len(nodeMap))
	PrintNodes()

	//Probe nodes
	ProbeNodes()
}

func initNodeAndTask() (err error) {
	nodeTaskMap := config.GetTaskConfigMap()
	nodeLimitMap := config.GetNodeConfigMap()

	nodeMapMutex.Lock()
	defer nodeMapMutex.Unlock()
	for nodeHost, taskLimit := range nodeLimitMap {
		if _, ok := nodeMap[nodeHost]; ok {
			//repeated node, ignore
			continue
		}

		limit, err := strconv.Atoi(taskLimit)
		if err != nil {
			//invalid task limit, error
			log.Root.Error("Invalid node task limit: %v", limit)
			return errors.New("Invalid node task limit")
		}

		node := NewNode(nodeHost, NodeStatusOffline, limit)

		if taskIDList, ok := nodeTaskMap[nodeHost]; ok {
			for _, taskID := range taskIDList {
				taskPara, err := config.ReadTaskPara(taskID)
				if err != nil {
					log.Root.Error("Invalid task param: %v, error: %v", taskID, err.Error())
					continue
				}
				task := NewTask(taskID, TaskStatusStopped, taskPara)
				node.AddTask(task)
			}
		}

		nodeMap[nodeHost] = node
	}

	return nil
}
