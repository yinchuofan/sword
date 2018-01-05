package node

import (
	"errors"
	"strconv"
	"sync"
	"time"

	"sword/SwordJobTracker/config"
	"sword/SwordJobTracker/log"
)

var (
	nodeMap      = make(map[string]*Node)
	nodeMapMutex = new(sync.RWMutex)
)

func PrintNodes() {
	nodeMapMutex.RLock()
	defer nodeMapMutex.RUnlock()

	for nodeHost, node := range nodeMap {
		log.Root.Info("NodeHost: %v", nodeHost)
		node.Print()
	}
}

func ProbeNodes() {
	go func() {
		for {
			nodeMapMutex.RLock()
			for k, _ := range nodeMap {
				go Probe(k)
			}
			nodeMapMutex.RUnlock()

			time.Sleep(time.Minute * 1)
		}
	}()
}

func Probe(nodeHost string) {
	err, taskIDList := ProbeNode(nodeHost)

	nodeMapMutex.Lock()
	defer nodeMapMutex.Unlock()

	if elem, ok := nodeMap[nodeHost]; ok {
		if err != nil {
			//Node status is offline
			elem.UpdateStatus(NodeStatusOffline)

			//Update all task status to stopped
			elem.SetAllTaskStatusStopped()
		} else {
			//Node status is online
			elem.UpdateStatus(NodeStatusOnline)

			//Update tasks status to running
			elem.UpdateRunningTaskListStatus(taskIDList)
		}
	}
}

func RegisterNode(nodeHost string, taskLimit int) (err error) {
	nodeMapMutex.Lock()
	defer nodeMapMutex.Unlock()

	if elem, ok := nodeMap[nodeHost]; ok {
		//Update node task limit
		elem.UpdateTaskLimit(taskLimit)

		//Update node status
		elem.UpdateStatus(NodeStatusOnline)

		//Update node config
		if err = config.SetNodeConfig(nodeHost, strconv.Itoa(taskLimit)); err != nil {
			log.Root.Error("Set node config error. NodeHost: %v, TaskLimit: %v", nodeHost, taskLimit)
			return err
		}
	} else {
		//Create node
		node := NewNode(nodeHost, NodeStatusOnline, taskLimit)

		//Add node
		nodeMap[nodeHost] = node

		//Add node config
		if err = config.AddNodeConfig(nodeHost, strconv.Itoa(taskLimit)); err != nil {
			log.Root.Error("Add node config error. NodeHost: %v, TaskLimit: %v", nodeHost, taskLimit)
			return err
		}
	}

	log.Root.Info("Register node succeed. NodeHost: %v, TaskLimit: %v", nodeHost, taskLimit)
	return nil
}

func NodeTaskFinished(nodeHost string, taskID string) (err error) {
	nodeMapMutex.Lock()
	defer nodeMapMutex.Unlock()

	if elem, ok := nodeMap[nodeHost]; ok {
		elem.UpdateTaskStatus(taskID, TaskStatusStopped)
		return nil
	} else {
		log.Root.Error("Task is not existed. NodeHost: %v, TaskID: %v", nodeHost, taskID)
		return errors.New("task is not existed")
	}
}

func NodeTaskException(nodeHost string, taskID string) (err error) {
	nodeMapMutex.Lock()
	defer nodeMapMutex.Unlock()

	if elem, ok := nodeMap[nodeHost]; ok {
		elem.UpdateTaskStatus(taskID, TaskStatusInterrupted)
		return nil
	} else {
		log.Root.Error("Task is not existed. NodeHost: %v, TaskID: %v", nodeHost, taskID)
		return errors.New("task is not existed")
	}
}

func CreateTask(taskPara string) (taskID string, err error) {
	//Generate task id
	taskID = NewTaskID()
	if len(taskID) == 0 {
		return "", errors.New("Generate task id error")
	}

	//Save task parameter information
	err = config.WriteTaskPara(taskID, taskPara)
	if err != nil {
		return "", err
	}

	return taskID, nil
}

func StartTask(taskID string) (taskCapacity int, err error) {
	node, taskPara, err := searchAvailableNode(taskID)
	if err != nil {
		log.Root.Error("Search available node error. TaskID: %v", taskID)
		return 0, err
	}

	if node == nil {
		//No available node
		log.Root.Error("Lack of computing resource. TaskID: %v", taskID)
		return -1, nil
	}

	//Start task
	if err = StartNodeTask(taskID, taskPara, node.nodeHost); err != nil {
		log.Root.Error("Start node task error. NodeHost: %v, TaskID: %v", node.nodeHost, taskID)
		return 0, err
	}
	node.UpdateTaskStatus(taskID, TaskStatusRunning)

	//Statistic total task capacity
	taskCapacity = statisticTaskCapacity()

	return taskCapacity, nil
}

func searchAvailableNode(taskID string) (node *Node, taskPara string, err error) {
	nodeMapMutex.Lock()
	defer nodeMapMutex.Unlock()

	//Find node which own the task
	for _, v := range nodeMap {
		if has, task := v.HasTask(taskID); has {
			//Task in node
			if NodeStatusOnline == v.nodeStatus && v.GetTaskCapacity() > 0 {
				return v, task.para, nil
			} else {
				//Remove task from config
				if err = config.RemoveTaskConfig(taskID); err != nil {
					log.Root.Error("Remove task config error. TaskID: %v", taskID)
					return nil, "", err
				}

				//Remove task from node
				v.RemoveTask(taskID)
				break
			}
		}
	}

	//Reselect node for task
	for _, v := range nodeMap {
		if NodeStatusOnline == v.nodeStatus && v.GetTaskCapacity() > 0 {
			//New task
			if taskPara, err = config.ReadTaskPara(taskID); err != nil {
				log.Root.Error("Read task para error. TaskID: %v", taskID)
				return nil, "", err
			}

			//Add task config
			if err = config.AddTaskConfig(taskID, v.nodeHost); err != nil {
				log.Root.Error("Add task config error. TaskID: %v", taskID)
				return nil, "", err
			}

			task := NewTask(taskID, TaskStatusStopped, taskPara)
			v.AddTask(task)

			return v, taskPara, nil
		}
	}

	//No available node
	return nil, "", nil
}

func statisticTaskCapacity() (capacity int) {
	nodeMapMutex.RLock()
	defer nodeMapMutex.RUnlock()

	capacity = 0

	for _, v := range nodeMap {
		if NodeStatusOnline == v.nodeStatus {
			capacity += v.GetTaskCapacity()
		}
	}

	return capacity
}

func statisticTaskLimit() (limit int) {
	nodeMapMutex.RLock()
	defer nodeMapMutex.RUnlock()

	limit = 0

	for _, v := range nodeMap {
		limit += v.taskLimit
	}

	return limit
}

func findTaskNode(taskID string) (node *Node, task *Task) {
	nodeMapMutex.Lock()
	defer nodeMapMutex.Unlock()

	for _, v := range nodeMap {
		if has, task := v.HasTask(taskID); has {
			return v, task
		}
	}

	return nil, nil
}

func StopTask(taskID string) (taskCapacity int, err error) {
	node, _ := findTaskNode(taskID)
	if node == nil {
		log.Root.Error("Task is not existed. TaskID: %v", taskID)
		return 0, errors.New("task is not existed")
	}

	if NodeStatusOffline == node.nodeStatus {
		log.Root.Error("Node is offline. NodeHost: %v, TaskID: %v", node.nodeHost, taskID)
		return 0, errors.New("node is offline")
	}

	//Stop task
	if err = StopNodeTask(taskID, node.nodeHost); err != nil {
		log.Root.Error("Stop node task error. NodeHost: %v, TaskID: %v", node.nodeHost, taskID)
		return 0, err
	}
	node.UpdateTaskStatus(taskID, TaskStatusStopped)

	//Statistic total task capacity
	taskCapacity = statisticTaskCapacity()

	return taskCapacity, nil
}

func DeleteTask(taskID string) (err error) {
	//Judje is task on running, can not delete running task
	node, task := findTaskNode(taskID)
	if node == nil {
		log.Root.Error("Task is not existed. TaskID: %v", taskID)
		return errors.New("task is not existed")
	}

	if TaskStatusRunning == task.status {
		log.Root.Error("Can not delete running task. NodeHost: %v, TaskID: %v", node.nodeHost, taskID)
		return errors.New("task is running")
	}

	//Remove task from config
	if err = config.RemoveTaskConfig(taskID); err != nil {
		log.Root.Error("Remove task config error. TaskID: %v", taskID)
		return err
	}

	//Delete task parameter file
	if err = config.DeleteTaskParaFile(taskID); err != nil {
		log.Root.Error("Delete task para file error. TaskID: %v", taskID)
	}

	//Remove task from node
	node.RemoveTask(taskID)

	return nil
}

func getTaskStatus(taskID string) (status int) {
	nodeMapMutex.Lock()
	defer nodeMapMutex.Unlock()

	for _, v := range nodeMap {
		if has, task := v.HasTask(taskID); has {
			return task.status
		}
	}

	return TaskStatusStopped
}

func QueryTaskStatus(taskIDList []string) (taskStatusMap map[string]int, err error) {
	taskStatusMap = make(map[string]int)

	for _, v := range taskIDList {
		taskStatusMap[v] = getTaskStatus(v)
	}

	return taskStatusMap, nil
}

func QueryAllTaskStatus() (taskStatusMap map[string]int, err error) {
	taskStatusMap = make(map[string]int)

	nodeMapMutex.RLock()
	defer nodeMapMutex.RUnlock()

	for _, node := range nodeMap {
		nodeTaskStatusMap := node.GetAllTaskStatus()
		for k, v := range nodeTaskStatusMap {
			taskStatusMap[k] = v
		}
	}

	return taskStatusMap, nil
}

func QueryTaskCapacity() (taskCapacity, taskLimit int, err error) {
	taskCapacity = statisticTaskCapacity()
	taskLimit = statisticTaskLimit()
	return taskCapacity, taskLimit, nil
}
