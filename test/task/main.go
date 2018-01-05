package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	flag.Parse()
	for k, v := range flag.Args() {
		fmt.Printf("arg %v: %v\n", k, v)
	}

	currentPath, err := GetCurrentPath()
	if err != nil {
		fmt.Println("get current path error")
		return
	}
	fmt.Println("current path: " + currentPath)

	i := 0

	for {
		fmt.Printf("process: %v\n", i)
		i++
		time.Sleep(time.Second * 1)
	}
}

func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}
