package main

import (
    "fmt"
    "net/http"
    "github.com/go-chi/chi" // Chi is a lightweight, idiomatic router that is used for HTTP services.
    "github.com/avukadin/goapi/internal/handlers" // Import for custom handlers (not defined here but assumed to exist).
    log "github.com/sirupsen/logrus" // Logrus is a structured logger for Go, used for logging errors and information.
)

func main() {
    log.SetReportCaller(true) // Enables logging of the caller function for better traceability of logs.
    
    var r *chi.Mux = chi.NewRouter() // This creates a new HTTP router using the Chi framework.
    handlers.Handler(r) // Assumes a Handler function exists in the handlers package to register routes and endpoints.

    fmt.Println("Starting GO API service...") // Outputs an indication that the API is starting.

    fmt.Println(`
      ██████╗  ██████╗     █████╗ ██████╗ ██████╗ ██╗
      ██╔══██╗██╔════╝    ██╔══██╗██╔══██╗██╔══██╗██║
      ██████╔╝██║  ███╗   ███████║██████╔╝██║  ██║██║
      ██╔══██╗██║   ██║   ██╔══██║██╔══██╗██║  ██║██║
      ██║  ██║╚██████╔╝██╗██║  ██║██║  ██║██████╔╝███████╗
      ╚═╝  ╚═╝ ╚═════╝ ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═════╝ ╚══════╝
    `) // This ASCII art logo provides visual feedback in the terminal when the API starts.

    err := http.ListenAndServe("localhost:8000", r) // The API will start listening for requests on port 8000.
    
    if err != nil { 
        log.Error(err) // If there is an error in starting the server, it will be logged using Logrus.
    }
}
