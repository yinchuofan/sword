package config

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func loadConfig(configPath string) (kvp map[string]string, err error) {
	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	kvp = make(map[string]string)

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		s := strings.TrimSpace(string(b))
		if strings.Index(s, "#") == 0 {
			continue
		}

		keyValueList := strings.Split(s, "=")
		if len(keyValueList) != 2 {
			continue
		}

		key := strings.TrimSpace(keyValueList[0])
		value := strings.TrimSpace(keyValueList[1])

		kvp[key] = value
	}

	return kvp, nil
}

func saveConfig(configPath string, kvp map[string]string) (err error) {
	f, err := os.OpenFile(configPath, os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for k, v := range kvp {
		keyValueString := k + "=" + v + "\r\n"

		_, err := w.WriteString(keyValueString)
		if err != nil {
			return err
		}
	}

	return w.Flush()
}
