package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
)

func main() {
    m := http.NewServeMux()
    m.HandleFunc("/", handlePage)

    // Read port from environment
    port := os.Getenv("PORT")
    if port == "" {
        port = "8999"
    }

    srv := http.Server{
        Handler:      m,
        Addr:         ":" + port,
        WriteTimeout: 30 * time.Second,
        ReadTimeout:  30 * time.Second,
    }

    fmt.Println("server started on port", port)
    err := srv.ListenAndServe()
    log.Fatal(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    w.WriteHeader(200)
    const page = `<html>
<head></head>
<body>
    <p> Hello from Docker! I'm a Go server. </p>
</body>
</html>`
    w.Write([]byte(page))
}
