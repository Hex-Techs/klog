package klog

import "github.com/jinzhu/gorm"

var DefaultGormLogger = GormLogger{V(0)}

// GormLogWriter log writer interface
type GormLogWriter interface {
	Infoln(v ...interface{})
}

type GormLogger struct {
	GormLogWriter
}

func (logger GormLogger) Print(values ...interface{}) {
	logger.Infoln(gorm.LogFormatter(values...)...)
}
