package http

import (
	"encoding/json"
	"net/http"
)

type H map[string]interface{}

func HttpResponseError(w http.ResponseWriter, err Err) {
	data := map[string]interface{}{
		"err":    err.Code,
		"errMsg": err.Message,
	}

	dataJSON, _ := json.Marshal(data)

	w.Write(dataJSON)
}

func HttpResponseOk(w http.ResponseWriter) {
	data := map[string]interface{}{
		"err":    ErrNoError.Code,
		"errMsg": ErrNoError.Message,
	}

	dataJSON, _ := json.Marshal(data)

	w.Write(dataJSON)
}

func HttpResponseData(w http.ResponseWriter, h H) {
	data := map[string]interface{}{
		"err":    ErrNoError.Code,
		"errMsg": ErrNoError.Message,
		"data":   h,
	}

	dataJSON, _ := json.Marshal(data)

	w.Write(dataJSON)
}
