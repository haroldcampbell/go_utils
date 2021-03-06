package serverutils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/haroldcampbell/go_utils/utils"
)

// ActionStatus return status for handlers
// @action is the name of the action e.g. signin
// @successStatus true is successful, false otherwise
// @message provide details (particularly) when @successStatus is false
type ActionStatus struct {
	Action        string      `json:"action"`
	SuccessStatus bool        `json:"successStatus"`
	Message       string      `json:"message"`
	JSONBody      interface{} `json:"jsonBody"`
	SessionKey    string      `json:"sessionKey"`
	ErrorCode     int         `json:"errorCode"`

	Writer http.ResponseWriter `json:"-"`
}

func (as *ActionStatus) writeActionStatus() ([]byte, error) {
	json, err := json.Marshal(as)
	if err != nil {
		return nil, err
	}
	as.Writer.Header().Set("Content-Type", "application/json")
	as.Writer.Write(json)

	return json, nil
}

func (as *ActionStatus) Write(success bool, message string) ([]byte, error) {
	as.SuccessStatus = success
	as.Message = message

	return as.writeActionStatus()
}

func (as *ActionStatus) Error(error string, code int) ([]byte, error) {
	as.SuccessStatus = false
	as.Message = error
	as.ErrorCode = code

	return as.writeActionStatus()
}

// RespondWithError logs and responds with a generic error message
func RespondWithError(as *ActionStatus, logger *utils.RoutineLogger, errorMessage string, errorCode int) {
	data, err := as.Error(errorMessage, errorCode)
	logger.LogActionStatus(data, err)
}

// ErrorResponse ...
type ErrorResponse struct {
	DidFail      bool
	ErrorMessage string
}

// WriteErrorMessage responds with specified error message and flag
func WriteErrorMessage(w http.ResponseWriter, errMessage string) {
	b, _ := json.Marshal(ErrorResponse{DidFail: true, ErrorMessage: errMessage})
	w.Write(b)
}

// https://gist.github.com/swdunlop/9629168
func identifyPanic() string {
	var name, file string
	var line int
	var pc [16]uintptr

	n := runtime.Callers(3, pc[:])
	for _, pc := range pc[:n] {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line = fn.FileLine(pc)
		name = fn.Name()
		if !strings.HasPrefix(name, "runtime.") {
			break
		}
	}

	switch {
	case name != "":
		return fmt.Sprintf("%v:%v", name, line)
	case file != "":
		return fmt.Sprintf("%v:%v", file, line)
	}

	return fmt.Sprintf("pc:%x", pc)
}

func RecoverPanic(src string) {
	r := recover()
	if r == nil {
		return
	}
	utils.Error(src, "PANIC:[%v] at %s", r, identifyPanic())
}
