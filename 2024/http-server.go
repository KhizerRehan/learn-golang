// golang http server
// To the equivalent code in Go:
package main


import (
    "fmt"
    "net/http"
)

// *http.Request is a struct that holds all the information about the request
// w http.ResponseWriter is a struct that holds all the information about the response


func dashboard(w http.ResponseWriter, req *http.Request) {
    fmt.Fprint(w, "Dashboard")
}

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
// req.Header is a map of key-value pairs representing the header fields of the request
    for name, headers := range req.Header {
        for _, h := range headers {
        // "%v: %v\n" is a format string that prints the name of the header and its value
        // w is the http.ResponseWriter that we are writing to and we are using the Fprintf method to write to it
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func fullName(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Khizer Rehan")
}

func main() {
    http.HandleFunc("/dashboard", dashboard)
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/fullname", fullName)
    http.HandleFunc("/headers", headers)
    http.ListenAndServe(":9090", nil)
}