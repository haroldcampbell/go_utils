package keyvals

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

func ReadStore(store map[string]string, targetDirPath, envFilename string) error {
	store = make(map[string]string)

	envFilePath := fmt.Sprintf("%s/%s", targetDirPath, envFilename)

	_, exists := isStoreFilePresent(envFilePath)
	if !exists {
		utils.Log("ReadStore", "Creating file: %s", envFilePath)
		err := createStoreFile(envFilePath)
		if err != nil {
			utils.Log("ReadStore", "Failed to create file: %s err: '%v'", envFilePath, err)
			return err
		}
	}

	rawBytes, err := ioutil.ReadFile(envFilePath)
	if err != nil {
		utils.Log("ReadStore", "Failed to read %s file: '%v'", envFilePath, err)
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

		store[key] = val
	}

	return nil
}

// UpdateStore ...
func UpdateStore(store map[string]string, targetDirPath, envFilename string, key string, val string) {
	store[key] = val

	envFilePath := fmt.Sprintf("%s/%s", targetDirPath, envFilename)
	file, err := os.OpenFile(envFilePath, os.O_APPEND|os.O_WRONLY|os.O_TRUNC, 0644)
	defer file.Close()

	if err != nil {
		log.Fatalf("[UpdateStore] Failed creating file: %s", err)
	}

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for key := range store {
		writer.WriteString(fmt.Sprintf("%s=%s\n", key, store[key]))
	}
}

func isStoreDirPresent(envDirname string) (fs.FileInfo, bool) {
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

		return isStoreDirPresent(targetFile)
	}

	if info != nil && err == nil && info.IsDir() {
		return info, true
	}

	return nil, false
}

func isStoreFilePresent(envFilename string) (fs.FileInfo, bool) {
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

func createStoreFile(envFilename string) error {
	file, err := os.Create(envFilename)
	defer file.Close()

	if err != nil {
		log.Fatalf("[createStoreFile] Failed while trying to create file: '%s' err: %v", envFilename, err)
		return err
	}
	return nil
}
