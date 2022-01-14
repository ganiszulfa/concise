package trace

import (
	"github.com/ganiszulfa/concise/pkg/inspect"
	log "github.com/sirupsen/logrus"
)

func Func() {
	name, line := inspect.GetParentFuncProps()
	log.Trace(name, ":", line)
}
