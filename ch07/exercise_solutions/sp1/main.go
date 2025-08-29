package main

import (
	"fmt"
)

type Handler interface {
	Handle(event string) string
}

type App struct {
	handlers []Handler
}

func (a *App) Register(h Handler) {
	a.handlers = append(a.handlers, h)
}

func (a *App) Run() {
	fmt.Println("Run start in LogHandler")
	s := "start"
	for _, h := range a.handlers {
		s = h.Handle(s)
	}
	fmt.Println("Run end in LogHandler")
}

type LogHandler struct{}

func (l *LogHandler) Handle(event string) string {
	s := fmt.Sprintf("[LogHandler] Event: %s", event)
	fmt.Println(s, "in LogHandler")
	return s
}

func main() {
	app := App{}
	app.Register(&LogHandler{})
	app.Register(&LogHandler{})
	app.Run()
}
