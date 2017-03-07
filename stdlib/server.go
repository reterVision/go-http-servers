package stdlib

import (
	"context"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/reterVision/go-http-servers/common"
)

const serverShutdownTimeout = 5 * time.Second

// Serve starts a HTTP server written with stdlib
func Serve(addr string) {
	interruptHandler := common.NewInterruptHandler()

	server := &http.Server{
		Addr:           addr,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	go func() {
		log.Print(server.ListenAndServe())
	}()

	interruptHandler.Wait()

	ctx, cancelFunc := context.WithTimeout(context.Background(), serverShutdownTimeout)
	defer cancelFunc()

	server.Shutdown(ctx)
}
