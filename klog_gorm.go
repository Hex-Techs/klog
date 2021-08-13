package klog

import "github.com/jinzhu/gorm"

var DefaultGormLogger = GormLogger{depth: 9}

type GormLogger struct {
	depth int
}

func (logger GormLogger) Print(values ...interface{}) {
	InfoDepth(logger.depth, gorm.LogFormatter(values...)...)
}
