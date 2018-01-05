package http

import (
	"net/http"
)

func Run(addr string) error {
	// Internal API
	http.HandleFunc("/startTask", HandleStartTask)
	http.HandleFunc("/stopTask", HandleStopTask)
	http.HandleFunc("/probe", HandleProbe)

	return http.ListenAndServe(addr, nil)
}
