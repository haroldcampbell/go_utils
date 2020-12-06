package envutils

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
