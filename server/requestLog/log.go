package ctxuestLog

import (
	"context"
	"log"
)

type PrefixFn func(ctx context.Context) string

var prefixFn PrefixFn = func(_ context.Context) string {
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

func Debug(ctx context.Context, txt string) {
	normal(DEBUG, ctx, txt)
}

func Info(ctx context.Context, txt string) {
	normal(INFO, ctx, txt)
}

func Warning(ctx context.Context, txt string) {
	normal(WARNING, ctx, txt)
}

func Fatal(ctx context.Context, txt string) {
	fatal(FATAL, ctx, txt)
}

func normal(cutoff LogLevel, ctx context.Context, txt string) {
	if cutoff < level {
		return
	}
	prefix := prefixFn(ctx)
	log.Printf("%s | %s - %v\n", prefix, cutoff.ToString(), txt)
}

func fatal(cutoff LogLevel, ctx context.Context, txt string) {
	if cutoff < level {
		return
	}
	prefix := prefixFn(ctx)
	log.Fatalf("%s - %v\n", prefix, txt)
}
