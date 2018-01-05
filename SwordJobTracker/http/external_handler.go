package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"sword/SwordJobTracker/log"
	"sword/SwordJobTracker/node"
)

// HandleCreateTask : Create task
// url : /createTask
func HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	log.Root.Info("HandleCreateTask BEGIN")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		HttpResponseError(w, ErrNotFound)
		return
	}

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	data := make(map[string]interface{})
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Root.Error("HandleCreateTask Parse HTTP request body error")
		HttpResponseError(w, ErrForm)
		return
	}

	elem, ok := data["taskPara"]
	if !ok {
		log.Root.Error("HandleCreateTask HTTP form data error")
		HttpResponseError(w, ErrForm)
		return
	}

	taskPara := elem.(string)
	taskID, err := node.CreateTask(taskPara)
	if err != nil {
		log.Root.Error("HandleCreateTask Create task error")
		HttpResponseError(w, ErrServer)
		return
	}

	log.Root.Info("HandleCreateTask END")
	HttpResponseData(w, H{
		"taskID": taskID,
	})
	return
}

// HandleStartTask : Start task
// url : /startTask
func HandleStartTask(w http.ResponseWriter, r *http.Request) {
	log.Root.Info("HandleStartTask BEGIN")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		HttpResponseError(w, ErrNotFound)
		return
	}

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	data := make(map[string]interface{})
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Root.Error("HandleStartTask Parse HTTP request body error")
		HttpResponseError(w, ErrForm)
		return
	}

	elem, ok := data["taskID"]
	if !ok {
		log.Root.Error("HandleStartTask HTTP form data error")
		HttpResponseError(w, ErrForm)
		return
	}

	taskID := elem.(string)
	taskCapacity, err := node.StartTask(taskID)
	if err != nil {
		log.Root.Error("HandleStartTask Start task error. TaskID: %v", taskID)
		HttpResponseError(w, ErrServer)
		return
	}

	if taskCapacity < 0 {
		log.Root.Error("HandleStartTask Lack of computing resources")
		HttpResponseError(w, ErrLackResources)
		return
	}

	log.Root.Info("HandleStartTask END")
	HttpResponseData(w, H{
		"taskCapacity": taskCapacity,
	})
	return
}

// HandleStopTask : Stop task
// url : /stopTask
func HandleStopTask(w http.ResponseWriter, r *http.Request) {
	log.Root.Info("HandleStopTask BEGIN")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		HttpResponseError(w, ErrNotFound)
		return
	}

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	data := make(map[string]interface{})
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Root.Error("HandleStopTask Parse HTTP request body error")
		HttpResponseError(w, ErrForm)
		return
	}

	elem, ok := data["taskID"]
	if !ok {
		log.Root.Error("HandleStopTask HTTP form data error")
		HttpResponseError(w, ErrForm)
		return
	}

	taskID := elem.(string)
	taskCapacity, err := node.StopTask(taskID)
	if err != nil {
		log.Root.Error("HandleStopTask Stop task error. TaskID: %v", taskID)
		HttpResponseError(w, ErrServer)
		return
	}

	log.Root.Info("HandleStopTask END")
	HttpResponseData(w, H{
		"taskCapacity": taskCapacity,
	})
	return
}

// HandleDeleteTask : Delete task
// url : /deleteTask
func HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	log.Root.Info("HandleDeleteTask BEGIN")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		HttpResponseError(w, ErrNotFound)
		return
	}

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	data := make(map[string]interface{})
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Root.Error("HandleDeleteTask Parse HTTP request body error")
		HttpResponseError(w, ErrForm)
		return
	}

	elem, ok := data["taskID"]
	if !ok {
		log.Root.Error("HandleDeleteTask HTTP form data error")
		HttpResponseError(w, ErrForm)
		return
	}

	taskID := elem.(string)
	err = node.DeleteTask(taskID)
	if err != nil {
		log.Root.Error("HandleDeleteTask Delete task error. TaskID: %v", taskID)
		HttpResponseError(w, ErrServer)
		return
	}

	log.Root.Info("HandleDeleteTask END")
	HttpResponseOk(w)
	return
}

// HandleQueryTaskStatus : Query tasks status
// url : /queryTaskStatus
func HandleQueryTaskStatus(w http.ResponseWriter, r *http.Request) {
	log.Root.Info("HandleQueryTaskStatus BEGIN")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		HttpResponseError(w, ErrNotFound)
		return
	}

	taskIDList := r.URL.Query()["taskID"]
	if len(taskIDList) == 0 {
		log.Root.Error("HandleQueryTaskStatus Parse HTTP request param error")
		HttpResponseError(w, ErrParams)
		return
	}

	taskStatusMap, err := node.QueryTaskStatus(taskIDList)
	if err != nil {
		log.Root.Error("HandleQueryTaskStatus Query task status error. TaskIDList: %v", taskIDList)
		HttpResponseError(w, ErrServer)
		return
	}

	taskStatusList := []interface{}{}
	for k, v := range taskStatusMap {
		jsonMap := map[string]interface{}{
			"taskID":     k,
			"taskStatus": v,
		}
		taskStatusList = append(taskStatusList, jsonMap)
	}

	log.Root.Info("HandleQueryTaskStatus END")
	HttpResponseData(w, H{
		"taskStatusList": taskStatusList,
	})
	return
}

// HandleQueryAllTaskStatus : Start all task status
// url : /queryAllTaskStatus
func HandleQueryAllTaskStatus(w http.ResponseWriter, r *http.Request) {
	log.Root.Info("HandleQueryAllTaskStatus BEGIN")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		HttpResponseError(w, ErrNotFound)
		return
	}

	taskStatusMap, err := node.QueryAllTaskStatus()
	if err != nil {
		log.Root.Error("HandleQueryAllTaskStatus Query all task status error")
		HttpResponseError(w, ErrServer)
		return
	}

	taskStatusList := []interface{}{}
	for k, v := range taskStatusMap {
		jsonMap := map[string]interface{}{
			"taskID":     k,
			"taskStatus": v,
		}
		taskStatusList = append(taskStatusList, jsonMap)
	}

	log.Root.Info("HandleQueryAllTaskStatus END")
	HttpResponseData(w, H{
		"taskStatusList": taskStatusList,
	})
	return
}

// HandleQueryTaskCapacity : Start all task status
// url : /queryTaskCapacity
func HandleQueryTaskCapacity(w http.ResponseWriter, r *http.Request) {
	log.Root.Info("HandleQueryTaskCapacity BEGIN")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		HttpResponseError(w, ErrNotFound)
		return
	}

	taskCapacity, taskLimit, err := node.QueryTaskCapacity()
	if err != nil {
		log.Root.Error("HandleQueryTaskCapacity Query task capacity error")
		HttpResponseError(w, ErrServer)
		return
	}

	log.Root.Info("HandleQueryTaskCapacity END")
	HttpResponseData(w, H{
		"taskCapacity": taskCapacity,
		"taskLimit":    taskLimit,
	})
	return
}
