package main

import (
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC3339)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(t))
}

func router() chi.Router {
	r := chi.NewRouter().With(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			xff := req.Header.Get("X-Forwarded-For")
			xrip := req.Header.Get("X-Real-Ip")
			ip, _, _ := net.SplitHostPort(req.RemoteAddr)
			slog.Info("log",
				"IP", ip,
				"X-Forwarded-For", xff,
				"X-Real-Ip", xrip,
			)
			handler.ServeHTTP(rw, req)
		})
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		getRoot(w, r)
	})

	return r
}

func main() {
	ops := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &ops))
	slog.SetDefault(logger)

	r := router()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
