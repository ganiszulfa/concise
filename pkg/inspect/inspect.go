// Package inspect provide functionality for developers to debug and develop features
package inspect

import (
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Do printout the value passed to it, it reduces work for developers
// to insert the line number, func name when debugging
func Do(vars ...interface{}) {
	name, line := GetParentFuncProps()

	if len(vars) == 0 {
		// when sometime we only want to know if
		// the execution is passing a certain line of code
		log.Printf("%16s %4d %3d %+v\n", name, line, 0, "PASSING HERE")
		return
	}

	for i, v := range vars {
		log.Printf("%16s %4d %3d %+v\n", name, line, i, v)
	}
}

func GetParentFuncProps() (funcName string, lineNo int) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame2, _ := frames.Next()
	lineNo = frame2.Line
	names := strings.Split(frame2.Function, "/")
	return names[len(names)-1], lineNo
}
