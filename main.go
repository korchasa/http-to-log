package main

import (
    "fmt"
    "github.com/fatih/color"
    "io"
    "log"
    "net/http"
    "os"
    "time"
)

func main() {
    listen := os.Getenv("LISTEN")
    if listen == "" {
        listen = ":8080"
    }
    http.HandleFunc("/", handler)
    log.Printf("Starting server on %s", listen)
    log.Fatal(http.ListenAndServe(listen, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    request := color.New(color.FgHiBlue).SprintFunc()
    method := color.New(color.FgHiYellow).SprintFunc()
    url := color.New(color.FgHiMagenta).SprintFunc()
    headerName := color.New(color.FgHiCyan).SprintFunc()
    headerValue := color.New(color.FgWhite).SprintFunc()
    bodyMeta := color.New(color.FgWhite).SprintFunc()
    bodyContent := color.New(color.FgHiGreen).SprintFunc()

    fmt.Println(request(">>> " + time.Now().Format("2006-01-02 15:04:05.000") + " from " + r.RemoteAddr))

    fmt.Println(method(r.Method), url(r.URL.String()))

    for name, values := range r.Header {
        for _, value := range values {
            fmt.Println(headerName(name), ":", headerValue(value))
        }
    }

    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Fatalf("Failed to read request body: %v", err)
    }

    fmt.Println(bodyMeta(fmt.Sprintf("Body[%db]:", len(body))))
    fmt.Println(bodyContent(string(body)))
    fmt.Println()

    w.WriteHeader(http.StatusOK)
    _, err = w.Write([]byte("OK"))
    if err != nil {
        log.Printf("Failed to write response: %v", err)
    }
}
