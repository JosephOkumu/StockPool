package main

import (
    "html/template"
    "log"
    "net/http"
)

var tmpl *template.Template

func init() {
    // Parse all HTML templates
    tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
    // Serve static files
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Route handlers
    http.HandleFunc("/", handleHome)
    http.HandleFunc("/login", handleLogin)
    http.HandleFunc("/dashboard", handleDashboard)

    // Start server
    log.Println("Server starting on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    err := tmpl.ExecuteTemplate(w, "index.html", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
    err := tmpl.ExecuteTemplate(w, "login.html", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func handleDashboard(w http.ResponseWriter, r *http.Request) {
    // TODO: Add authentication check here
    err := tmpl.ExecuteTemplate(w, "dashboard_base.html", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}