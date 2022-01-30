package main

import (
    "fmt"
    "net/http"
    "html/template"
    "path"
)

func home(w http.ResponseWriter, r *http.Request){
    fmt.Println(w, "Welcome to the Home!")
    fmt.Println("Endpoint Hit: Home")

    fp := path.Join("templates", "index.html")
        
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/html")

    if err := tmpl.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func authorize(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the Authorize!")
    fmt.Println("Endpoint Hit: Authorize")
    fmt.Println("Authorize GET params:", r.URL.Query())
}

func handleRequests() {
    http.HandleFunc("/", home)
    http.HandleFunc("/authorize", authorize)
    http.ListenAndServe(":3000", nil)
}

func main() {
    fmt.Println("Hello, World! Before")
    handleRequests()
    fmt.Println("Hello, World! After")
}
