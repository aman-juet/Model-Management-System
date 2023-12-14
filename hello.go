package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the About Page.")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You can contact us at contact@example.com.")
}

func main() {
	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Define routes and their respective handlers
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/about", aboutHandler).Methods("GET")
	router.HandleFunc("/contact", contactHandler).Methods("GET")

	// Set up the server
	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
