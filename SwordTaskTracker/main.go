package main

import (
	"flag"

	"sword/SwordTaskTracker/http"
	"sword/SwordTaskTracker/log"
	"sword/SwordTaskTracker/task"
)

var (
	//master para
	mPtr      = flag.String("m", "", "master host")
	masterPtr = flag.String("master", "", "master host")

	//node ip para
	iPtr  = flag.String("i", "", "node ip")
	ipPtr = flag.String("ip", "", "node ip")

	//node port para
	pPtr    = flag.String("p", "", "listen port")
	portPtr = flag.String("port", "", "listen port")

	//task limit para
	lPtr     = flag.Int("l", 0, "task limit")
	limitPtr = flag.Int("limit", 0, "task limit")

	//task process special directory
	tPtr        = flag.String("t", "", "task process special directory")
	taskPathPtr = flag.String("taskpath", "", "task process special directory")
)

func main() {
	masterHost, nodeIP, nodePort, taskLimit, taskPath := parseFlags()
	log.Root.Info("MasterHost: %v, NodeIP: %v, NodePort: %v, TaskLimit: %v, TaskPath: %v", masterHost, nodeIP, nodePort, taskLimit, taskPath)

	nodeHost := nodeIP + ":" + nodePort
	task.Setup(masterHost, nodeHost, taskLimit, taskPath)

	log.Root.Info("Listening and Serving HTTP on :%v", nodePort)
	err := http.Run("0.0.0.0:" + nodePort)
	if err != nil {
		log.Root.Error("Listening and Serving HTTP error")
		panic("ListenAndServe: " + err.Error())
	}
}

func parseFlags() (masterHost, nodeIP, nodePort string, taskLimit int, taskPath string) {
	flag.Parse()

	switch true {
	case *mPtr != "":
		masterHost = *mPtr
	case *masterPtr != "":
		masterHost = *masterPtr
	default:
		masterHost = "127.0.0.1:6000"
	}

	switch true {
	case *iPtr != "":
		nodeIP = *iPtr
	case *ipPtr != "":
		nodeIP = *ipPtr
	default:
		nodeIP = "127.0.0.1"
	}

	switch true {
	case *pPtr != "":
		nodePort = *pPtr
	case *portPtr != "":
		nodePort = *portPtr
	default:
		nodePort = "6100"
	}

	switch true {
	case *lPtr != 0:
		taskLimit = *lPtr
	case *limitPtr != 0:
		taskLimit = *limitPtr
	default:
		taskLimit = 4
	}

	switch true {
	case *tPtr != "":
		taskPath = *tPtr
	case *taskPathPtr != "":
		taskPath = *taskPathPtr
	default:
		taskPath = "./task"
	}

	return masterHost, nodeIP, nodePort, taskLimit, taskPath
}
