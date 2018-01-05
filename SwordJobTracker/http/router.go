package http

import (
	"net/http"
)

func Run(addr string) error {

	// External API
	http.HandleFunc("/createTask", HandleCreateTask)
	http.HandleFunc("/startTask", HandleStartTask)
	http.HandleFunc("/stopTask", HandleStopTask)
	http.HandleFunc("/deleteTask", HandleDeleteTask)
	http.HandleFunc("/queryTaskStatus", HandleQueryTaskStatus)
	http.HandleFunc("/queryAllTaskStatus", HandleQueryAllTaskStatus)
	http.HandleFunc("/queryTaskCapacity", HandleQueryTaskCapacity)

	// Internal API
	http.HandleFunc("/register", HandleRegister)
	http.HandleFunc("/taskFinished", HandleTaskFinished)
	http.HandleFunc("/taskException", HandleTaskException)

	return http.ListenAndServe(addr, nil)
}
