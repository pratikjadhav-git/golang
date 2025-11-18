
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // This is the code that runs when a request hits your app
    fmt.Fprintf(w, "Hello, World! (Now running as a service)")
}

func main() {
    // 1. Register the handler function for the root path
    http.HandleFunc("/", handler)

    // 2. Read the port from the environment variable (standard in OpenShift/S2I)
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if the environment variable is not set
    }

    log.Printf("Starting web server on port %s...", port)
    
    // 3. Start the server. This line blocks execution and keeps the process running indefinitely.
    // If ListenAndServe fails (e.g., port already in use), it returns an error, which we log and exit on.
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
