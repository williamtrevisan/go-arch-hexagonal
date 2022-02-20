package server

import (
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "github.com/williamtrevisan/go-arch-hexagonal/adapters/web/handlers"
    "github.com/williamtrevisan/go-arch-hexagonal/app"
    "log"
    "net/http"
    "os"
    "time"
)

type Webserver struct {
    Service app.ProductServiceInterface
}

func NewWebserver() *Webserver {
    return &Webserver{}
}

func (w Webserver) Serve() {
    r := mux.NewRouter()
    n := negroni.New(
        negroni.NewLogger(),
    )

    handlers.MakeProductHandlers(r, n, w.Service)

    http.Handle("/", r)

    server := &http.Server{
        ReadHeaderTimeout: 10 * time.Second,
        WriteTimeout:      10 * time.Second,
        Addr:              ":9000",
        Handler:           http.DefaultServeMux,
        ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
    }

    err := server.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }
}
