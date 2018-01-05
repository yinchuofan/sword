package task

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"sword/SwordTaskTracker/log"
)

func RegisterNode() (err error) {
	log.Root.Info("RegisterNode BEGIN. MasterHost: %v, NodeHost: %v, TaskLimit: %v", MasterHost, NodeHost, TaskLimit)

	mapNodeInfo := map[string]interface{}{
		"nodeHost":  NodeHost,
		"taskLimit": TaskLimit,
	}

	nodeJSON, err := json.Marshal(mapNodeInfo)
	if err != nil {
		log.Root.Error("Marshal Register JSON error.")
		return err
	}

	url := "http://" + MasterHost + "/register"
	resp, err := http.Post(url, "text/plain", strings.NewReader(string(nodeJSON)))
	if err != nil {
		log.Root.Error("Post Register HTTP request error: %v, url: %v", err.Error(), url)
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Root.Error("Read Register HTTP response error: %v, url: %v", err.Error(), url)
		return err
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Root.Error("Parse Register HTTP response error")
		return err
	}

	errCode := int(result["err"].(float64))
	errMsg := result["errMsg"].(string)
	if errCode != 0 {
		log.Root.Error("Node register error: " + errMsg)
		return errors.New("node register error: " + errMsg)
	}

	log.Root.Info("RegisterNode succeed. NodeHost: %v", NodeHost)
	return nil
}

func ReportTaskFinished(taskID string) (err error) {
	log.Root.Info("ReportTaskFinished BEGIN. TaskID: %v", taskID)

	mapTaskInfo := map[string]interface{}{
		"nodeHost": NodeHost,
		"taskID":   taskID,
	}

	taskJSON, err := json.Marshal(mapTaskInfo)
	if err != nil {
		log.Root.Error("Marshal TaskFinished JSON error.")
		return err
	}

	url := "http://" + MasterHost + "/taskFinished"
	resp, err := http.Post(url, "text/plain", strings.NewReader(string(taskJSON)))
	if err != nil {
		log.Root.Error("Post TaskFinished HTTP request error: %v, url: %v", err.Error(), url)
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Root.Error("Read TaskFinished HTTP response error: %v, url: %v", err.Error(), url)
		return err
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Root.Error("Parse TaskFinished HTTP response error")
		return err
	}

	errCode := int(result["err"].(float64))
	errMsg := result["errMsg"].(string)
	if errCode != 0 {
		log.Root.Error("Node report task finished error: " + errMsg)
		return errors.New("node report task finished error: " + errMsg)
	}

	log.Root.Info("ReportTaskFinished succeed. TaskID: %v", taskID)
	return nil
}

func ReportTaskException(taskID string) (err error) {
	log.Root.Info("ReportTaskException BEGIN. TaskID: %v", taskID)

	mapTaskInfo := map[string]interface{}{
		"nodeHost": NodeHost,
		"taskID":   taskID,
	}

	taskJSON, err := json.Marshal(mapTaskInfo)
	if err != nil {
		log.Root.Error("Marshal TaskException JSON error.")
		return err
	}

	url := "http://" + MasterHost + "/taskException"
	resp, err := http.Post(url, "text/plain", strings.NewReader(string(taskJSON)))
	if err != nil {
		log.Root.Error("Post TaskException HTTP request error: %v, url: %v", err.Error(), url)
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Root.Error("Read TaskException HTTP response error: %v, url: %v", err.Error(), url)
		return err
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Root.Error("Parse TaskException HTTP response error")
		return err
	}

	errCode := int(result["err"].(float64))
	errMsg := result["errMsg"].(string)
	if errCode != 0 {
		log.Root.Error("Node report task exception error: " + errMsg)
		return errors.New("node report task exception error: " + errMsg)
	}

	log.Root.Info("ReportTaskException succeed. TaskID: %v", taskID)
	return nil
}
