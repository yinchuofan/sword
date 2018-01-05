package config

import (
	"io/ioutil"
	"os"
)

const (
	TaskParaPath string = "./config/"
)

func ReadTaskPara(taskID string) (para string, err error) {
	taskParaFilePath := TaskParaPath + taskID + ".json"
	f, err := os.Open(taskParaFilePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	para = string(data)

	return para, nil
}

func WriteTaskPara(taskID string, para string) (err error) {
	taskParaFilePath := TaskParaPath + taskID + ".json"

	return ioutil.WriteFile(taskParaFilePath, []byte(para), 0666)
}

func DeleteTaskParaFile(taskID string) (err error) {
	taskParaFilePath := TaskParaPath + taskID + ".json"
	return os.Remove(taskParaFilePath)
}
