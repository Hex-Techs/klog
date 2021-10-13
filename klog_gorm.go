package klog

import (
	"fmt"

	"github.com/go-logr/logr"
)

type GormLoggerWriter struct {
}

func (l *GormLoggerWriter) Printf(format string, v ...interface{}) {
	printDepth(logging.logr, 2, fmt.Sprintf(format, v...))
}

func printDepth(logr logr.Logger, depth int, args ...interface{}) {
	buf := logging.getBuffer()
	if !logging.skipHeaders {
		now := timeNow()
		if logging.enableColor {
			logging.colorFormatHeaderTimestampSeverity(buf, infoLog, now)
		} else {
			logging.formatHeaderTimestampSeverity(buf, infoLog, now)
		}
	}
	// if logr is set, we clear the generated header as we rely on the backing
	// logr implementation to print headers
	if logr != nil {
		logging.putBuffer(buf)
		buf = logging.getBuffer()
	}
	fmt.Fprint(buf, args...)
	if buf.Bytes()[buf.Len()-1] != '\n' {
		buf.WriteByte('\n')
	}
	logging.output(infoLog, logr, buf, "", 0, false)
}
