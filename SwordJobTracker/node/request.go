package node

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"sword/SwordJobTracker/log"
)

func StartNodeTask(taskID, taskPara, nodeHost string) (err error) {
	log.Root.Info("StartNodeTask BEGIN. NodeHost: %v, TaskID: %v", nodeHost, taskID)

	mapTaskInfo := map[string]interface{}{
		"taskID":   taskID,
		"taskPara": taskPara,
	}

	taskJSON, err := json.Marshal(mapTaskInfo)
	if err != nil {
		log.Root.Error("Marshal StartTask JSON error.")
		return err
	}

	url := "http://" + nodeHost + "/startTask"
	resp, err := http.Post(url, "text/plain", strings.NewReader(string(taskJSON)))
	if err != nil {
		log.Root.Error("Post StartTask HTTP request error: %v, url: %v", err.Error(), url)
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Root.Error("Read StartTask HTTP response error: %v, url: %v", err.Error(), url)
		return err
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Root.Error("Parse StartTask HTTP response error")
		return err
	}

	errCode := int(result["err"].(float64))
	errMsg := result["errMsg"].(string)
	if errCode != 0 {
		log.Root.Error("Node start task error: " + errMsg)
		return errors.New("node start task error: " + errMsg)
	}

	log.Root.Info("StartNodeTask succeed. NodeHost: %v, TaskID: %v", nodeHost, taskID)
	return nil
}

func StopNodeTask(taskID, nodeHost string) (err error) {
	log.Root.Info("StopNodeTask BEGIN. NodeHost: %v, TaskID: %v", nodeHost, taskID)

	mapTaskInfo := map[string]interface{}{
		"taskID": taskID,
	}

	taskJSON, err := json.Marshal(mapTaskInfo)
	if err != nil {
		log.Root.Error("Marshal StopTask JSON error.")
		return err
	}

	url := "http://" + nodeHost + "/stopTask"
	resp, err := http.Post(url, "text/plain", strings.NewReader(string(taskJSON)))
	if err != nil {
		log.Root.Error("Post StopTask HTTP request error: %v, url: %v", err.Error(), url)
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Root.Error("Read StopTask HTTP response error: %v, url: %v", err.Error(), url)
		return err
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Root.Error("Parse StopTask HTTP response error")
		return err
	}

	errCode := int(result["err"].(float64))
	errMsg := result["errMsg"].(string)
	if errCode != 0 {
		log.Root.Error("Node stop task error: " + errMsg)
		return errors.New("node stop task error: " + errMsg)
	}

	log.Root.Info("StopNodeTask succeed. NodeHost: %v, TaskID: %v", nodeHost, taskID)
	return nil
}

func ProbeNode(nodeHost string) (err error, taskIDList []string) {
	log.Root.Info("ProbeNode BEGIN. NodeHost: %v", nodeHost)

	url := "http://" + nodeHost + "/probe"
	resp, err := http.Post(url, "text/plain", nil)
	if err != nil {
		log.Root.Error("Post ProbeNode HTTP request error: %v, url: %v", err.Error(), url)
		return err, nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Root.Error("Read ProbeNode HTTP response error: %v, url: %v", err.Error(), url)
		return err, nil
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Root.Error("Parse ProbeNode HTTP response error")
		return err, nil
	}

	errCode := int(result["err"].(float64))
	errMsg := result["errMsg"].(string)
	if errCode != 0 {
		log.Root.Error("Node probe error: " + errMsg)
		return errors.New("node probe error: " + errMsg), nil
	}

	taskIDList = []string{}

	data := result["data"].(map[string]interface{})
	if _, ok := data["taskIDList"]; ok {
		taskIDListTmp := data["taskIDList"].([]interface{})
		for _, v := range taskIDListTmp {
			taskIDList = append(taskIDList, v.(string))
		}
	}

	log.Root.Info("ProbeNode succedd. NodeHost: %v, running tasks: %v", nodeHost, taskIDList)
	return nil, taskIDList
}
