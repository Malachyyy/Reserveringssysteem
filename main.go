package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", BeginScherm)
	http.HandleFunc("/Account", Account)
	http.HandleFunc("/Register", RegisterHandler)

	http.ListenAndServe(":8080", nil)
}

func BeginScherm(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("loginform.html")
	if err != nil {
		panic(err)
	}
	tpl.ExecuteTemplate(w, "loginform.html", nil)

	http.Redirect(w, r, "/Register", http.StatusSeeOther)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("Regristratieform.html")
	if err != nil {
		panic(err)
	}
	tpl.ExecuteTemplate(w, "Regristratieform.html", nil)
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
