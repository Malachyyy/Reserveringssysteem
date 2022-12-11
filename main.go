package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Display the login form
		fmt.Fprintf(w, `
        <html>
             <body>
                <h1>Login</h1>
                <form action="/login" method="post">
                    <label for="username">Username:</label><br>
                    <input type="text" id="username" name="username"><br>
                    <label for="password">Password:</label><br>
                    <input type="password" id="password" name="password"><br><br>
                    <input type="submit" value="Submit">
                </form> 
            </body>
        </html>
     `)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
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
	})

	http.ListenAndServe(":8080", nil)
}
