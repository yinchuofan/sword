package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"sword/SwordJobTracker/log"
	"sword/SwordJobTracker/node"
)

// HandleRegister : TaskTracker register
// url : /register
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	log.Root.Info("HandleRegister BEGIN")

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
		log.Root.Error("HandleRegister Parse HTTP request body error")
		HttpResponseError(w, ErrForm)
		return
	}

	nodeHost := ""
	if elem, ok := data["nodeHost"]; ok {
		nodeHost = elem.(string)
	} else {
		log.Root.Error("HandleRegister HTTP form data error")
		HttpResponseError(w, ErrForm)
		return
	}

	taskLimit := 0
	if elem, ok := data["taskLimit"]; ok {
		taskLimit = int(elem.(float64))
	} else {
		log.Root.Error("HandleRegister HTTP form data error")
		HttpResponseError(w, ErrForm)
		return
	}

	err = node.RegisterNode(nodeHost, taskLimit)
	if err != nil {
		log.Root.Error("HandleRegister Register node error")
		HttpResponseError(w, ErrServer)
		return
	}

	log.Root.Info("HandleRegister END")
	HttpResponseOk(w)
	return
}

// HandleTaskFinished : TaskTracker task finished
// url : /taskFinished
func HandleTaskFinished(w http.ResponseWriter, r *http.Request) {
	log.Root.Info("HandleTaskFinished BEGIN")

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
		log.Root.Error("HandleTaskFinished Parse HTTP request body error")
		HttpResponseError(w, ErrForm)
		return
	}

	nodeHost := ""
	if elem, ok := data["nodeHost"]; ok {
		nodeHost = elem.(string)
	} else {
		log.Root.Error("HandleTaskFinished HTTP form data error")
		HttpResponseError(w, ErrForm)
		return
	}

	taskID := ""
	if elem, ok := data["taskID"]; ok {
		taskID = elem.(string)
	} else {
		log.Root.Error("HandleTaskFinished HTTP form data error")
		HttpResponseError(w, ErrForm)
		return
	}

	err = node.NodeTaskFinished(nodeHost, taskID)
	if err != nil {
		log.Root.Error("HandleTaskFinished Node task finished process error")
		HttpResponseError(w, ErrServer)
		return
	}

	log.Root.Info("HandleTaskFinished END")
	HttpResponseOk(w)
	return
}

// HandleTaskException : TaskTracker task exception
// url : /taskException
func HandleTaskException(w http.ResponseWriter, r *http.Request) {
	log.Root.Info("HandleTaskException BEGIN")

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
		log.Root.Error("HandleTaskException Parse HTTP request body error")
		HttpResponseError(w, ErrForm)
		return
	}

	nodeHost := ""
	if elem, ok := data["nodeHost"]; ok {
		nodeHost = elem.(string)
	} else {
		log.Root.Error("HandleTaskException HTTP form data error")
		HttpResponseError(w, ErrForm)
		return
	}

	taskID := ""
	if elem, ok := data["taskID"]; ok {
		taskID = elem.(string)
	} else {
		log.Root.Error("HandleTaskException HTTP form data error")
		HttpResponseError(w, ErrForm)
		return
	}

	err = node.NodeTaskException(nodeHost, taskID)
	if err != nil {
		log.Root.Error("HandleTaskFinished Node task exception process error")
		HttpResponseError(w, ErrServer)
		return
	}

	log.Root.Info("HandleTaskException END")
	HttpResponseOk(w)
	return
}
