package main

import (
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"todo_module/internal/controller"
)

func main() {
	slog.Info("Todo Server Starting...")
	mux := http.NewServeMux()
	path, handler := controller.NewHandler()
	mux.Handle(path, handler)
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	slog.Info("Todo Server Started")

	tls := false // OPTIONAL https hosting
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.key"
		err := http.ListenAndServeTLS(
			":443",
			certFile,
			keyFile,
			h2c.NewHandler(mux, &http2.Server{}),
		)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	} else {
		http.ListenAndServe(
			":3000",
			// Use h2c so we can serve HTTP/2 without TLS.
			h2c.NewHandler(mux, &http2.Server{}),
		)
	}
}
