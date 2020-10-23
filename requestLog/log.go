package requestLog

import (
	"log"
	"net/http"
)

type PrefixFn func(req *http.Request) string

var prefixFn PrefixFn = func(_ *http.Request) string {
	return ""
}

func SetPrefixFn(fn PrefixFn) {
	prefixFn = fn
}

type LogLevel int

func (l LogLevel) ToString() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	FATAL
)

var level = INFO

func SetLogLevel(newLevel LogLevel) {
	level = newLevel
}

func Debug(req *http.Request, txt string) {
	normal(DEBUG, req, txt)
}

func Info(req *http.Request, txt string) {
	normal(INFO, req, txt)
}

func Warning(req *http.Request, txt string) {
	normal(WARNING, req, txt)
}

func Fatal(req *http.Request, txt string) {
	fatal(FATAL, req, txt)
}

func normal(cutoff LogLevel, req *http.Request, txt string) {
	if cutoff < level {
		return
	}
	prefix := prefixFn(req)
	log.Printf("%s | %s - %v\n", prefix, cutoff.ToString(), txt)
}

func fatal(cutoff LogLevel, req *http.Request, txt string) {
	if cutoff < level {
		return
	}
	prefix := prefixFn(req)
	log.Fatalf("%s - %v\n", prefix, txt)
}
