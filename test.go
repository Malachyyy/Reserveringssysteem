package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template

func main() {
	tmpl, err := template.ParseFiles("templates/loginform.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", BeginScherm)
	http.HandleFunc("/Account", Account)
	http.ListenAndServe(":8080", nil)
}

func BeginScherm(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "loginform.html", nil)
	if err != nil {
		panic(err)
	}
}

func Account(w http.ResponseWriter, r *http.Request) {
	// Get the username and password from the request
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "admin" && password == "password" {
		// If authentication is successful, return a success message
		fmt.Fprintf(w, "Login successful!")
	} else {
		// If authentication fails, return an error message
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}
}
