package envutils

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/haroldcampbell/go_utils/utils"
)

var envHash = make(map[string]string)

func ReadEnvFile(targetDirPath, envFilename string) error {
	envFilePath := fmt.Sprintf("%s/%s", targetDirPath, envFilename)

	_, exists := isEnvFilePresent(envFilePath)
	if !exists {
		utils.Log("ReadEnvFile", "Creating file: %s", envFilePath)
		err := createEnvFile(envFilePath)
		if err != nil {
			utils.Log("ReadEnvFile", "Failed to create file: %s err: '%v'", envFilePath, err)
			return err
		}
	}

	rawBytes, err := ioutil.ReadFile(envFilePath)
	if err != nil {
		utils.Log("ReadEnvFile", "Failed to read %s file: '%v'", envFilePath, err)
		return err
	}

	contents := string(rawBytes)
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
func UpdateEnvFile(targetDirPath, envFilename string, key string, val string) {
	envHash[key] = val

	envFilePath := fmt.Sprintf("%s/%s", targetDirPath, envFilename)
	file, err := os.OpenFile(envFilePath, os.O_APPEND|os.O_WRONLY|os.O_TRUNC, 0644)
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

// func getPath(info fs.FileInfo, name string) (string, error) {
// 	if info == nil {
// 		return "", fmt.Errorf("missing fileInfo")
// 	}

// 	if info.Mode()&os.ModeSymlink != 0 {
// 		targetFile, err := os.Readlink(info.Name())

// 		utils.Log("getPath", "is symlink. targetFile: '%v'", targetFile)

// 		if err != nil {
// 			return "", fmt.Errorf("readlink failed file: %s  err: %v", info.Name(), err)
// 		}

// 		return targetFile, nil
// 	}

// 	utils.Log("getPath", "not symlink. targetFile: '%v'", name)

// 	return name, nil
// }

func isEnvDirPresent(envDirname string) (fs.FileInfo, bool) {
	info, err := os.Stat(envDirname)
	if info != nil && err == nil && info.IsDir() {
		return info, true
	}

	info, err = os.Lstat(envDirname)
	if info != nil && err == nil && info.Mode()&os.ModeSymlink != 0 {
		targetFile, err := os.Readlink(info.Name())

		if err != nil {
			return nil, false
		}

		return isEnvDirPresent(targetFile)
	}

	if info != nil && err == nil && info.IsDir() {
		return info, true
	}

	return nil, false
}

func isEnvFilePresent(envFilename string) (fs.FileInfo, bool) {
	info, err := os.Stat(envFilename)

	if info != nil && err == nil && !info.IsDir() {
		return info, true
	}

	info, err = os.Lstat(envFilename)
	if info != nil && err == nil && !info.IsDir() {
		return info, true
	}

	return nil, false
}

func createEnvFile(envFilename string) error {
	file, err := os.Create(envFilename)
	defer file.Close()

	if err != nil {
		log.Fatalf("[createEnvFile] Failed while trying to create file: '%s' err: %v", envFilename, err)
		return err
	}
	return nil
}
