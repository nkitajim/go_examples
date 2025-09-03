package main

import (
	"context"
	"fmt"
)

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
)

type logKey struct{}

func ContextWithLevel(ctx context.Context, level Level) context.Context {
	return context.WithValue(ctx, logKey{}, level)
}

func LevelFromContext(ctx context.Context) (Level, bool) {
	level, ok := ctx.Value(logKey{}).(Level)
	return level, ok
}

func Log(ctx context.Context, level Level, message string) {
	var inLevel Level
	inLevel, ok := LevelFromContext(ctx)
	if !ok {
		return
	}
	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}
	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println(message)
	}
}

func main() {
	ctx := context.Background()

	fmt.Println("Set Log level", Info)
	ctx = ContextWithLevel(ctx, Info)
	Log(ctx, Debug, "This is a debug message")
	Log(ctx, Info, "This is a info message")

	fmt.Println("Set Log level", Debug)
	ctx = ContextWithLevel(ctx, Debug)
	Log(ctx, Debug, "This is a debug message")
	Log(ctx, Info, "This is a info message")
}
