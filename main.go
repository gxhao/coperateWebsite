package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// PageData creates a structure to hold data for parsing templates
type PageData struct {
	Title string
	Active string
}

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Route handlers
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/about", handleAbout)
	http.HandleFunc("/services", handleServices)
	http.HandleFunc("/contact", handleContact)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	// Parse layout and the specific page template
	// Note: We are parsing files on every request for easier development (hot reload of templates).
	// For production, templates should be cached.
    
    // Check if files exist before parsing to avoid panic
    layoutPath := filepath.Join("templates", "layout.html")
    pagePath := filepath.Join("templates", tmpl + ".html")

	t, err := template.ParseFiles(layoutPath, pagePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
	renderTemplate(w, "index", PageData{Title: "Home", Active: "home"})
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about", PageData{Title: "About Us", Active: "about"})
}

func handleServices(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "services", PageData{Title: "Services", Active: "services"})
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact", PageData{Title: "Contact Us", Active: "contact"})
}
