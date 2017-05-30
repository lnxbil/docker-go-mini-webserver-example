// Example taken from
//  https://tutorialedge.net/post/golang/creating-simple-web-server-with-golang/
package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
    fmt.Println("Starting Webserver on Port 8081")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))
}
