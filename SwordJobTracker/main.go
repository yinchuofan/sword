package main

import (
	"flag"

	"sword/SwordJobTracker/http"
	"sword/SwordJobTracker/log"
	_ "sword/SwordJobTracker/node"
)

var (
	pPtr    = flag.String("p", "", "listen port")
	portPtr = flag.String("port", "", "listen port")
)

func main() {
	flag.Parse()

	port := ""

	switch true {
	case *pPtr != "":
		port = *pPtr
	case *portPtr != "":
		port = *portPtr
	default:
		port = "6000"
	}

	log.Root.Info("Listening and Serving HTTP on :%v", port)
	err := http.Run("0.0.0.0:" + port)
	if err != nil {
		log.Root.Error("Listening and Serving HTTP error")
		panic("ListenAndServe: " + err.Error())
	}
}
