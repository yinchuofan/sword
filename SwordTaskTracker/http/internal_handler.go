package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"sword/SwordTaskTracker/log"
	"sword/SwordTaskTracker/task"
)

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

	taskID := ""
	if elem, ok := data["taskID"]; ok {
		taskID = elem.(string)
	} else {
		log.Root.Error("HandleStartTask HTTP form data error")
		HttpResponseError(w, ErrForm)
		return
	}

	taskPara := ""
	if elem, ok := data["taskPara"]; ok {
		taskPara = elem.(string)
	} else {
		log.Root.Error("HandleStartTask HTTP form data error")
		HttpResponseError(w, ErrForm)
		return
	}

	err = task.StartTask(taskID, taskPara)
	if err != nil {
		log.Root.Error("HandleStartTask Start task error. TaskID: %v", taskID)
		HttpResponseError(w, ErrServer)
		return
	}

	log.Root.Info("HandleStartTask END")
	HttpResponseOk(w)
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

	taskID := ""
	if elem, ok := data["taskID"]; ok {
		taskID = elem.(string)
	} else {
		log.Root.Error("HandleStopTask HTTP form data error")
		HttpResponseError(w, ErrForm)
		return
	}

	err = task.StopTask(taskID)
	if err != nil {
		log.Root.Error("HandleStopTask Start task error. TaskID: %v", taskID)
		HttpResponseError(w, ErrServer)
		return
	}

	log.Root.Info("HandleStopTask END")
	HttpResponseOk(w)
	return
}

// HandleProbe : Probe node tasks status
// url : /probe
func HandleProbe(w http.ResponseWriter, r *http.Request) {
	log.Root.Info("HandleProbe BEGIN")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		HttpResponseError(w, ErrNotFound)
		return
	}

	taskIDList, err := task.GetAllRunningTasks()
	if err != nil {
		log.Root.Error("HandleProbe Get all running tasks error")
		HttpResponseError(w, ErrServer)
		return
	}

	log.Root.Info("HandleProbe END")
	HttpResponseData(w, H{
		"taskIDList": taskIDList,
	})
	return
}
