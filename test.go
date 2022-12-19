package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Global variable
var tmpl *template.Template
var db *sql.DB

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

	r.ParseForm()
	// Get the username and password from the request
	username := r.FormValue("username")
	password := r.FormValue("password")

	if err := CheckLogin(db, username, password); err == nil {
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

func CheckLogin(db *sql.DB, username, password string) error {
	var id int
	var storedUsername string
	var storedPassword string

	err := db.QueryRow("SELECT id, username, password FROM users WHERE username=?", username).Scan(&id, &storedUsername, &storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("invalid username or password")
		}
		return err
	}
	if password == storedPassword {
		return nil
	}

	return fmt.Errorf("invalid login credentials")
}
