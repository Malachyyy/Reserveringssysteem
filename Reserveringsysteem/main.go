package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Global variable for template
var tmpl *template.Template

func main() {
	// Connect to database
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		fmt.Println("Failed to connecto database", err)
	}
	defer db.Close()
	fmt.Println("Connection successfull!")

	// Parse all templates in folder templates
	tmpl, _ = template.ParseGlob("templates/*.html")

	http.HandleFunc("/", BeginScherm)
	http.HandleFunc("/login", login)
	http.HandleFunc("/Register", RegisterHandler)
	http.ListenAndServe(":8080", nil)
}

func BeginScherm(w http.ResponseWriter, r *http.Request) {
	// Show login form
	err := tmpl.ExecuteTemplate(w, "loginform.html", nil)
	if err != nil {
		fmt.Println("Failed to execute loginform", err)
	}

	// Get the username and password from the request
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "admin" && password == "password" {
		// If authentication is successful, return a success message
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	} else {
		// If authentication fails, return an error message
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Successful!")
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Show Regristration form
	err := tmpl.ExecuteTemplate(w, "Regristratieform.html", nil)
	if err != nil {
		fmt.Println("Failed to execute Regristratieform", err)
	}
}
