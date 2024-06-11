package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
)


func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(lrw, r)

		end := time.Now()
		duration := end.Sub(start)

		var statusColor func(format string, a ...interface{}) string

		switch {
		case lrw.statusCode >= 500:
			statusColor = color.New(color.FgRed).SprintfFunc()
		case lrw.statusCode >= 400:
			statusColor = color.New(color.FgYellow).SprintfFunc()
		case lrw.statusCode >= 300:
			statusColor = color.New(color.FgCyan).SprintfFunc()
		default:
			statusColor = color.New(color.FgGreen).SprintfFunc()
		}

		logInfo := fmt.Sprintf("%s %s : %-30s %s : %v",
			end.Format(time.RFC3339),             // TIMESTAMP
			color.BlueString(r.Method),           // REQUEST_HEADER (METHOD)
			color.GreenString(r.URL.Path),        // ENDPOINT
			statusColor("%d", lrw.statusCode),    // RESPONSE_CODE
			duration,                             // TIME_TAKEN
		)

		log.Println(logInfo)
	})
}



type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

