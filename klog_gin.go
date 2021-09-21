package klog

import (
	"time"

	"github.com/gin-gonic/gin"
)

func setupLogging(duration time.Duration) {
	go func() {
		for range time.Tick(duration) {
			Flush()
		}
	}()
}

// ErrorLogger returns an ErrorLoggerT with parameter gin.ErrorTypeAny
func GinErrorLogger() gin.HandlerFunc {
	return GinErrorLoggerT(gin.ErrorTypeAny)
}

// GinErrorLoggerT returns an GinErrorLoggerT middleware with the given
// type gin.ErrorType.
func GinErrorLoggerT(typ gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if !c.Writer.Written() {
			json := c.Errors.ByType(typ).JSON()
			if json != nil {
				c.JSON(-1, json)
			}
		}
	}
}

// Logger prints a logline for each request and measures the time to
// process for a call. It formats the log entries similar to
// http://godoc.org/github.com/gin-gonic/gin#Logger does.
//
// Example:
//        router := gin.New()
//        router.Use(gin.Recovery())
//        router.Use(klog.Logger(3 * time.Second))
func GinLogger(duration time.Duration) gin.HandlerFunc {
	setupLogging(duration)
	return func(c *gin.Context) {
		t := time.Now()

		// process request
		c.Next()

		latency := time.Since(t)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		statusColor := colorForStatus(statusCode)
		methodColor := colorForMethod(method)
		path := c.Request.URL.Path

		switch {
		case statusCode >= 400 && statusCode <= 499:
			{
				Warningf("[GIN] |%s %3d %s| %12v | %s |%s  %s %-7s %s\n%s",
					statusColor, statusCode, Reset,
					latency,
					clientIP,
					methodColor, method, Reset,
					path,
					c.Errors.String(),
				)
			}
		case statusCode >= 500:
			{
				Errorf("[GIN] |%s %3d %s| %12v | %s |%s  %s %-7s %s\n%s",
					statusColor, statusCode, Reset,
					latency,
					clientIP,
					methodColor, method, Reset,
					path,
					c.Errors.String(),
				)
			}
		default:
			V(0).Infof("[GIN] |%s %3d %s| %12v | %s |%s  %s %-7s %s\n%s",
				statusColor, statusCode, Reset,
				latency,
				clientIP,
				methodColor, method, Reset,
				path,
				c.Errors.String(),
			)
		}

	}
}

func GinLoggerWithOutPaths(duration time.Duration, paths map[string]bool) gin.HandlerFunc {
	setupLogging(duration)
	return func(c *gin.Context) {
		t := time.Now()

		// process request
		c.Next()

		latency := time.Since(t)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		statusColor := colorForStatus(statusCode)
		methodColor := colorForMethod(method)
		path := c.Request.URL.Path

		if nolog, ok := paths[path]; ok {
			if nolog {
				return
			}
		}

		switch {
		case statusCode >= 400 && statusCode <= 499:
			{
				Warningf("[GIN] |%s %3d %s| %12v | %s |%s  %s %-7s %s\n%s",
					statusColor, statusCode, Reset,
					latency,
					clientIP,
					methodColor, method, Reset,
					path,
					c.Errors.String(),
				)
			}
		case statusCode >= 500:
			{
				Errorf("[GIN] |%s %3d %s| %12v | %s |%s  %s %-7s %s\n%s",
					statusColor, statusCode, Reset,
					latency,
					clientIP,
					methodColor, method, Reset,
					path,
					c.Errors.String(),
				)
			}
		default:
			V(0).Infof("[GIN] |%s %3d %s| %12v | %s |%s  %s %-7s %s\n%s",
				statusColor, statusCode, Reset,
				latency,
				clientIP,
				methodColor, method, Reset,
				path,
				c.Errors.String(),
			)
		}

	}
}

func colorForStatus(code int) string {
	switch {
	case code >= 200 && code <= 299:
		return HighlightGreen
	case code >= 300 && code <= 399:
		return HighlightWhite
	case code >= 400 && code <= 499:
		return HighlightYellow
	default:
		return BlinkingRed
	}
}

func colorForMethod(method string) string {
	switch {
	case method == "GET":
		return HighlightBlue
	case method == "POST":
		return HighlightCyan
	case method == "PUT":
		return HighlightYellow
	case method == "DELETE":
		return HighlightRed
	case method == "PATCH":
		return HighlightGreen
	case method == "HEAD":
		return HighlightMagenta
	case method == "OPTIONS":
		return HighlightWhite
	default:
		return Reset
	}
}
