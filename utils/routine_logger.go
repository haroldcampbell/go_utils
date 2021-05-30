package utils

import (
	"fmt"
	"net/http"
	"sync"
)

var goRoutineCounter = 0
var goRoutineCounterMutex sync.Mutex

type MuteLevel uint

const (
	None MuteLevel = iota
	TestMute
	FullMute
)

// var isRunningInTest = false
var muteLevel MuteLevel

func init() {
	// isRunningInTest = flag.Lookup("test.v") != nil
	// fmt.Printf("test.v: %v\n", flag.Lookup("test.v") == nil)
}

func SetMuteLevel(l MuteLevel) {
	muteLevel = l
}

func getNextGoRoutineCounter() int {
	goRoutineCounterMutex.Lock()
	defer goRoutineCounterMutex.Unlock()

	goRoutineCounter++

	return goRoutineCounter
}

type RoutineLogger struct {
	id          int
	handlerName string // The name of the handler
	m           sync.Mutex
	Mute        bool
	HideHeading bool
	indentCount int
	identStr    string
}

type loggerType = func(format string, v ...interface{})

var routineLoggerMutex sync.Mutex

func NewGoRoutineLogger(handlerName string) *RoutineLogger {
	routineLoggerMutex.Lock()
	defer routineLoggerMutex.Unlock()

	logger := &RoutineLogger{handlerName: handlerName}
	logger.id = getNextGoRoutineCounter()

	return logger
}

// NilLogger used in tests
func NilLogger() *RoutineLogger {
	return NewGoRoutineLogger("")
}

func (l *RoutineLogger) Indent() {
	str := ""

	l.indentCount++
	for i := 0; i < l.indentCount; i++ {
		str += "\t"
	}

	l.identStr = str
}

func (l *RoutineLogger) Unindent() {
	str := ""

	l.indentCount--
	for i := 0; i < l.indentCount; i++ {
		str += "\t"
	}

	l.identStr = str
}

func (l *RoutineLogger) Header() string {
	if l.HideHeading {
		return ""
	}
	return fmt.Sprintf("[%s(%d)]", YellowText(l.handlerName), l.id)
}

func (l *RoutineLogger) Print(format string, v ...interface{}) {
	if muteLevel == FullMute {
		//TODO: Refactor this
		return
	}

	if l.HideHeading {
		fmt.Printf(format, v...)
	} else {
		buf := formattedDateTime()
		str := fmt.Sprintf(format, v...)

		fmt.Printf("%s %s", ColoredText(DarkGrayTextFG, string(buf)), str)
	}
}

func (l *RoutineLogger) Err(err error) {
	if muteLevel == FullMute {
		//TODO: Refactor this
		return
	}

	l.m.Lock()
	defer l.m.Unlock()

	header := RedText("ERROR ") + l.Header()
	message := fmt.Sprintf("%v", err)
	l.Print("%s%s %s\n", l.identStr, header, message)
}

func (l *RoutineLogger) Error(format string, v ...interface{}) {
	if muteLevel == FullMute {
		//TODO: Refactor this
		return
	}

	l.m.Lock()
	defer l.m.Unlock()

	header := RedText("ERROR ") + l.Header()
	message := fmt.Sprintf(format, v...)
	l.Print("%s%s %s\n", l.identStr, header, message)
}

func (l *RoutineLogger) LogRequest(r *http.Request, v ...interface{}) {
	if muteLevel == FullMute || muteLevel == TestMute {
		//TODO: Refactor this
		return
	}

	if v == nil {
		l.Log("%v\n", r.URL.String())
		return
	}
	l.Log("%v => %v\n", r.URL.String(), ToString(v[0]))
}

func (l *RoutineLogger) Log(format string, v ...interface{}) {
	if muteLevel == FullMute || muteLevel == TestMute {
		//TODO: Refactor this
		return
	}

	l.m.Lock()
	defer l.m.Unlock()

	if l.Mute {
		return
	}

	header := l.Header()
	message := fmt.Sprintf(format, v...)

	l.Print("%s%s %s\n", l.identStr, header, message)
}

func (l *RoutineLogger) Tag(format string, v ...interface{}) {
	if muteLevel == FullMute {
		//TODO: Refactor this
		return
	}

	l.m.Lock()
	defer l.m.Unlock()

	if l.Mute {
		return
	}
	header := l.Header()
	message := fmt.Sprintf(format, v...)

	l.Print("%s%s %s %s", l.identStr, header, CyanText("==>"), message)
}

const TruncatedMessageLimitTiny = 100
const TruncatedMessageLimitSmall = 500

func (l *RoutineLogger) LogActionStatus(data []byte, err error, limitMessage ...int) {
	if muteLevel == FullMute {
		//TODO: Refactor this
		return
	}

	header := l.Header()

	if err != nil {
		l.Print("%s%s %s: %s\n", l.identStr, header, RedText("ERROR"), RedText(err))
		return
	}

	if muteLevel == TestMute {
		testLimit := TruncatedMessageLimitTiny
		l.Print("%s%s sent(%s): %s...\n", l.identStr, header, CyanText("test-truncated"), data[0:testLimit])
		return
	}

	if len(limitMessage) == 0 || len(data) < limitMessage[0] {
		l.Print("%s%s sent: %s\n", l.identStr, header, data)
		return
	}

	l.Print("%s%s sent(%s): %s...\n", l.identStr, header, CyanText("truncated"), data[0:limitMessage[0]])

}
