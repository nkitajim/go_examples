package main

import (
	"encoding/json"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/go_examples/ch12/exercise_solutions/sp1/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Example API
// @version 1.0
// @description これはサンプルAPIです
// @host localhost:8080
// @BasePath /

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ip, _, _ := net.SplitHostPort(req.RemoteAddr)
		reqID := middleware.GetReqID(req.Context())
		ww := middleware.NewWrapResponseWriter(rw, req.ProtoMajor)
		t1 := time.Now()

		defer func() {
			slog.Info("log",
				"IP", ip,
				"Method", req.Method,
				"RequestURI", req.RequestURI,
				"RemoteAddr", req.RemoteAddr,
				"RequestID", reqID,
				"Status", ww.Status(),
				"Size", ww.BytesWritten(),
				"durationTime", time.Since(t1)/1000,
			)
		}()
		next.ServeHTTP(ww, req)
	})
}

type TimeOut struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth int    `json:"day_of_month"`
	Month      string `json:"month"`
	Year       int    `json:"year"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Second     int    `json:"second"`
}

// @Summary 時刻を返す
// @Accept  json
// @Produce  json
// @Success 200 {object} TimeOut
// @Router / [get]
func getRoot(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	out := buildJSON(now)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out))
}

func buildJSON(now time.Time) string {
	timeOut := TimeOut{
		DayOfWeek:  now.Weekday().String(),
		DayOfMonth: now.Day(),
		Month:      now.Month().String(),
		Year:       now.Year(),
		Hour:       now.Hour(),
		Minute:     now.Minute(),
		Second:     now.Second(),
	}
	out, _ := json.Marshal(timeOut)
	return string(out)
}

func router() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(loggingMiddleware)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

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
