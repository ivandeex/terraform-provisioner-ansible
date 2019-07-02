package dbg

import (
	"fmt"
	"os"
)

const logPath = "/home/vanko/devel/debug.log"
var logFile *os.File

func StartLog() {
	var err error
	logFile, err = os.OpenFile(logPath, os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
}

func StopLog() {
	_ = logFile.Close()
}

// Log ...
func Log(format string, args ...interface{}) {
	str := "\n>>>> " + fmt.Sprintf(format, args...) + "\n"
	logFile.WriteString(str)
}
