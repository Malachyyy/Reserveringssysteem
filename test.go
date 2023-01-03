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
var db *sql.DB

func main() {
	// Connect to database
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		fmt.Println("Failed to connecto database", err)
	}
	defer db.Close()

	// Parse all templates in folder templates
	tmpl, _ = template.ParseGlob("templates/*.html")

	http.HandleFunc("/", BeginScherm)
	http.HandleFunc("/login", login)
	http.HandleFunc("/Register", RegisterHandler)
	http.ListenAndServe(":8080", nil)
}

func BeginScherm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Render the login form template
		err := tmpl.ExecuteTemplate(w, "loginform.html", nil)
		if err != nil {
			fmt.Println("Failed to execute loginform", err)
		}
	} else if r.Method == "POST" {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Get the username and password from the request
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		// Retrieve the user's credentials from the database
		var dbUsername, dbPassword string
		err = db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&dbUsername, &dbPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		// Compare the retrieved credentials with the provided credentials
		if dbUsername == username && dbPassword == password {
			// Credentials are correct, redirect to the login page
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		} else {
			// Credentials are incorrect, render an error message
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		}
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
