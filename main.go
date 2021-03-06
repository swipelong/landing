package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseFiles(
	"templates/index.html",
	"templates/unisexAd.html",
	"templates/maleAd.html",
	"templates/newCallToAction.html",
	"templates/about.html",
	"templates/blog.html",
	"templates/femaleAd.html"))

// swipelong.com/
func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// swipelong.com/ad0
func unisexAdHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "unisexAd.html", nil)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func newCallToActionHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "newCallToAction.html", nil)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// swipelong.com/ad1
func maleAdHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "maleAd.html", nil)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// swipelong.com/ad2
func femaleAdHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "femaleAd.html", nil)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// swipelong.com/about
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// swipelong.com/blog
func blogHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "blog.html", nil)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type EmailSubmission struct {
	EmailAddress string
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request.
	decoder := json.NewDecoder(r.Body)
	var es EmailSubmission
	err := decoder.Decode(&es)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	// Process request.
	ef := new(EmailForm)
	http_code, err := ef.CreateEmailSubmission(es.EmailAddress, r.RemoteAddr)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http_code)
	} else {
		fmt.Fprintf(w, "OK")
	}
}

func main() {
	// normal routes.
	http.HandleFunc("/", unisexAdHandler)
	http.HandleFunc("/ad0", indexHandler)
	http.HandleFunc("/ad1", maleAdHandler)
	http.HandleFunc("/ad2", femaleAdHandler)
	// Don't know if he wanted ad3 or ad4 so make them both work.
	http.HandleFunc("/ad3", newCallToActionHandler)
	http.HandleFunc("/ad4", newCallToActionHandler)
	http.HandleFunc("/api", apiHandler)
	// New blog pages.
	http.HandleFunc("/about", aboutHandler);
	http.HandleFunc("/blog", blogHandler);
	// public routes.
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public", fs))
	// run.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
