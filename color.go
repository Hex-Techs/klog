package klog

import (
	"fmt"
	"time"
)

var (
	Black            = string([]byte{27, 91, 48, 59, 51, 48, 109})
	Red              = string([]byte{27, 91, 48, 59, 51, 49, 109})
	Green            = string([]byte{27, 91, 48, 59, 51, 50, 109})
	Yellow           = string([]byte{27, 91, 48, 59, 51, 51, 109})
	Bule             = string([]byte{27, 91, 48, 59, 51, 52, 109})
	Magenta          = string([]byte{27, 91, 48, 59, 51, 53, 109})
	Cyan             = string([]byte{27, 91, 48, 59, 51, 54, 109})
	White            = string([]byte{27, 91, 48, 59, 51, 55, 109})
	HighlightBlack   = string([]byte{27, 91, 49, 59, 51, 48, 109})
	HighlightRed     = string([]byte{27, 91, 49, 59, 51, 49, 109})
	HighlightGreen   = string([]byte{27, 91, 49, 59, 51, 50, 109})
	HighlightYellow  = string([]byte{27, 91, 49, 59, 51, 51, 109})
	HighlightBlue    = string([]byte{27, 91, 49, 59, 51, 52, 109})
	HighlightMagenta = string([]byte{27, 91, 49, 59, 51, 53, 109})
	HighlightCyan    = string([]byte{27, 91, 49, 59, 51, 54, 109})
	HighlightWhite   = string([]byte{27, 91, 49, 59, 51, 55, 109})
	UnderlineBlack   = string([]byte{27, 91, 52, 59, 51, 48, 109})
	UnderlineRed     = string([]byte{27, 91, 52, 59, 51, 49, 109})
	UnderlineGreen   = string([]byte{27, 91, 52, 59, 51, 50, 109})
	UnderlineYellow  = string([]byte{27, 91, 52, 59, 51, 51, 109})
	UnderlineBlue    = string([]byte{27, 91, 52, 59, 51, 52, 109})
	UnderlineMagenta = string([]byte{27, 91, 52, 59, 51, 53, 109})
	UnderlineCyan    = string([]byte{27, 91, 52, 59, 51, 54, 109})
	UnderlineWhite   = string([]byte{27, 91, 52, 59, 51, 55, 109})
	BlinkingBlack    = string([]byte{27, 91, 53, 59, 51, 48, 109})
	BlinkingRed      = string([]byte{27, 91, 53, 59, 51, 49, 109})
	BlinkingGreen    = string([]byte{27, 91, 53, 59, 51, 50, 109})
	BlinkingYellow   = string([]byte{27, 91, 53, 59, 51, 51, 109})
	BlinkingBlue     = string([]byte{27, 91, 53, 59, 51, 52, 109})
	BlinkingMagenta  = string([]byte{27, 91, 53, 59, 51, 53, 109})
	BlinkingCyan     = string([]byte{27, 91, 53, 59, 51, 54, 109})
	BlinkingWhite    = string([]byte{27, 91, 53, 59, 51, 55, 109})
	ReverseBlack     = string([]byte{27, 91, 55, 59, 51, 48, 109})
	ReverseRed       = string([]byte{27, 91, 55, 59, 51, 49, 109})
	ReverseGreen     = string([]byte{27, 91, 55, 59, 51, 50, 109})
	ReverseYellow    = string([]byte{27, 91, 55, 59, 51, 51, 109})
	ReverseBlue      = string([]byte{27, 91, 55, 59, 51, 52, 109})
	ReverseMagenta   = string([]byte{27, 91, 55, 59, 51, 53, 109})
	ReverseCyan      = string([]byte{27, 91, 55, 59, 51, 54, 109})
	ReverseWhite     = string([]byte{27, 91, 55, 59, 51, 55, 109})
	InvisibleBlack   = string([]byte{27, 91, 56, 59, 51, 48, 109})
	InvisibleRed     = string([]byte{27, 91, 56, 59, 51, 49, 109})
	InvisibleGreen   = string([]byte{27, 91, 56, 59, 51, 50, 109})
	InvisibleYellow  = string([]byte{27, 91, 56, 59, 51, 51, 109})
	InvisibleBlue    = string([]byte{27, 91, 56, 59, 51, 52, 109})
	InvisibleMagenta = string([]byte{27, 91, 56, 59, 51, 53, 109})
	InvisibleCyan    = string([]byte{27, 91, 56, 59, 51, 54, 109})
	InvisibleWhite   = string([]byte{27, 91, 56, 59, 51, 55, 109})

	Reset = string([]byte{27, 91, 48, 109})
)

func (l *loggingT) colorHeader(buf *buffer, t time.Time, s severity, file string, line int) {
	buf.WriteString(l.time(t))
	buf.tmp[0] = ' '
	buf.Write(buf.tmp[:1])
	buf.WriteString(l.color.pidColor)
	buf.nDigits(7, 1, pid, ' ')
	buf.Write(buf.tmp[:8])
	buf.WriteString(Reset)
	buf.tmp[0] = ' '
	buf.tmp[1] = '['
	buf.Write(buf.tmp[:2])
	buf.WriteString(l.color.level(s))
	buf.tmp[0] = ']'
	buf.tmp[1] = ' '
	buf.Write(buf.tmp[:2])
	buf.WriteString(l.color.colorFile(file, line))
	buf.tmp[0] = ']'
	buf.tmp[1] = ' '
	buf.Write(buf.tmp[:2])
}

func (l *loggingT) time(t time.Time) string {
	return l.color.time(t)
}

type color struct {
	dataColor  string
	timeColor  string
	zoneColor  string
	pidColor   string
	levleColor levelColor
	fileColor  string
	msgColor   string
}

func (c color) time(t time.Time) string {
	zone, _ := t.Zone()
	return fmt.Sprintf("%s%04d-%02d-%02d%s %s%02d:%02d:%02d.%03d%s %s%s%s", c.dataColor, t.Year(), int(t.Month()), t.Day(), Reset, c.timeColor, t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1000/1000, Reset, c.zoneColor, zone, Reset)
}

func (c color) colorFile(file string, line int) string {
	return fmt.Sprintf("%s%s:%d%s", c.fileColor, file, line, Reset)
}

func (c color) level(s severity) string {
	switch s {
	case infoLog:
		return c.levleColor.info
	case warningLog:
		return c.levleColor.waring
	case errorLog:
		return c.levleColor.err
	case fatalLog:
		return c.levleColor.fatal
	default:
		return ""
	}
}

type levelColor struct {
	info   string
	waring string
	err    string
	fatal  string
}
