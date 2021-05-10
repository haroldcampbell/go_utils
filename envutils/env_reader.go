package envutils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/haroldcampbell/go_utils/utils"
)

var envHash = make(map[string]string)

func ReadEnvFile(envFilename string) error {
	if !isEnvFilePresent(envFilename) {
		utils.Log("ReadEnvFile", "Creating %s file\n", envFilename)
		err := createEnvFile(envFilename)
		if err != nil {
			utils.Log("ReadEnvFile", "Failed to initialize %s file: '%v'\n", envFilename, err)
			return err
		}
	}

	rawBytes, err := ioutil.ReadFile(envFilename)
	if err != nil {
		utils.Log("ReadEnvFile", "Failed to read %s file: '%v'\n", envFilename, err)
		return err
	}

	contents := fmt.Sprintf("%s", rawBytes)
	lines := strings.Split(contents, "\n")

	for _, line := range lines {
		line := strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") {
			// ignore comments in the file
			continue
		}

		index := strings.Index(line, "=")
		if index == -1 {
			continue
		}
		key := strings.TrimSpace(line[:index])
		val := strings.TrimSpace(line[index+1:])

		envHash[key] = val
	}

	return nil
}

// GetEnv returns the value associated with the key
func GetEnv(key string) string {
	return envHash[key]
}

// UpdateEnvFile ...
func UpdateEnvFile(envFilename string, key string, val string) {
	envHash[key] = val

	file, err := os.OpenFile(envFilename, os.O_APPEND|os.O_WRONLY|os.O_TRUNC, 0644)
	defer file.Close()

	if err != nil {
		log.Fatalf("[UpdateEnvFile] Failed creating file: %s", err)
	}

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for key := range envHash {
		writer.WriteString(fmt.Sprintf("%s=%s\n", key, envHash[key]))
	}
}

func isEnvFilePresent(envFilename string) bool {
	info, err := os.Stat(envFilename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func createEnvFile(envFilename string) error {
	file, err := os.Create(envFilename)
	defer file.Close()

	if err != nil {
		log.Fatalf("[createEnvFile] Failed while trying to create %s file", envFilename)
		return err
	}
	return nil
}
