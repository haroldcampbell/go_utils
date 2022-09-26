package envutils

import (
	"os"
	"path/filepath"

	"github.com/haroldcampbell/go_utils/utils"
)

const isDev = "isDev"
const isProd = "isProd"

// Default to production environment
var isRunningOnLocal = isProd

// SetIsProduction sets the environment to running in local
func SetIsProduction(status bool) {
	if status {
		isRunningOnLocal = isProd
	} else {
		isRunningOnLocal = isDev
	}
}

// IsDevEnv ..
func IsDevEnv() bool {
	return isRunningOnLocal == isDev
}

// IsProdEnv ..
func IsProdEnv() bool {
	return isRunningOnLocal == isProd
}

func ReportAppPath(stem string) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	utils.Log(stem, "Running in folder: %s", exPath)
}
