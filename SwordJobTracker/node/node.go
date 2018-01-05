package node

import (
	"strconv"
	"strings"
	"sync"

	"sword/SwordJobTracker/log"
)

type Node struct {
	nodeHost     string //host : "[ip]:[port]" or domain name
	nodeStatus   int
	taskLimit    int
	taskMap      map[string]*Task
	taskMapMutex *sync.RWMutex
}

func NewNode(host string, status int, limit int) *Node {
	return &Node{
		nodeHost:     host,
		nodeStatus:   status,
		taskLimit:    limit,
		taskMap:      make(map[string]*Task),
		taskMapMutex: new(sync.RWMutex),
	}
}

func (n *Node) Print() {
	log.Root.Info("----------Node Info----------")
	log.Root.Info("Host: %v, Status: %v", n.nodeHost, n.nodeStatus)

	n.taskMapMutex.RLock()
	defer n.taskMapMutex.RUnlock()
	for taskID, task := range n.taskMap {
		log.Root.Info("Task ID: %v, Status: %v, Param: %v", taskID, task.status, task.para)
	}
}

func (n *Node) GetTaskCapacity() (capacity int) {
	n.taskMapMutex.RLock()
	defer n.taskMapMutex.RUnlock()

	runningTaskNum := 0
	for _, v := range n.taskMap {
		if TaskStatusRunning == v.status {
			runningTaskNum++
		}
	}

	if n.taskLimit-runningTaskNum < 0 {
		return 0
	}
	return n.taskLimit - runningTaskNum
}

func (n *Node) HasTask(taskID string) (has bool, t *Task) {
	n.taskMapMutex.RLock()
	defer n.taskMapMutex.RUnlock()

	if elem, ok := n.taskMap[taskID]; ok {
		return true, elem
	}
	return false, nil
}

func (n *Node) AddTask(t *Task) {
	n.taskMapMutex.Lock()
	defer n.taskMapMutex.Unlock()

	n.taskMap[t.id] = t
}

func (n *Node) RemoveTask(taskID string) {
	n.taskMapMutex.Lock()
	defer n.taskMapMutex.Unlock()

	if _, ok := n.taskMap[taskID]; ok {
		delete(n.taskMap, taskID)
	}
}

func (n *Node) UpdateStatus(status int) {
	n.nodeStatus = status
}

func (n *Node) UpdateTaskLimit(limit int) {
	n.taskLimit = limit
}

func (n *Node) UpdateRunningTaskListStatus(onlineTaskIDList []string) {
	n.taskMapMutex.Lock()
	defer n.taskMapMutex.Unlock()
	for k, v := range n.taskMap {
		var status = TaskStatusStopped
		for _, id := range onlineTaskIDList {
			if k == id {
				status = TaskStatusRunning
				break
			}
		}
		v.status = status
	}
}

func (n *Node) SetAllTaskStatusStopped() {
	n.taskMapMutex.Lock()
	defer n.taskMapMutex.Unlock()

	for _, v := range n.taskMap {
		v.status = TaskStatusStopped
	}
}

func (n *Node) UpdateTaskStatus(taskID string, taskStatus int) {
	n.taskMapMutex.Lock()
	defer n.taskMapMutex.Unlock()
	if elem, ok := n.taskMap[taskID]; ok {
		elem.status = taskStatus
	}
}

func (n *Node) GetAllTaskStatus() (taskStatusMap map[string]int) {
	taskStatusMap = make(map[string]int)

	n.taskMapMutex.Lock()
	defer n.taskMapMutex.Unlock()

	for _, v := range n.taskMap {
		taskStatusMap[v.id] = v.status
	}

	return taskStatusMap
}

func GetIpPort(host string) (ip string, port int) {
	keyValueList := strings.Split(host, ":")
	if len(keyValueList) != 2 {
		return "", 0
	}

	key := strings.TrimSpace(keyValueList[0])
	value := strings.TrimSpace(keyValueList[1])

	ip = key
	port, err := strconv.Atoi(value)
	if err != nil {
		return "", 0
	}

	return ip, port
}

func GetHost(ip string, port int) (host string) {
	return ip + strconv.Itoa(port)
}
