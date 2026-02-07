package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http" // Standard library for web servers
)

type USER struct {
	ID       int    `json:"Id"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
	IsActive bool   `json:"-"`
}
type Profile struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

// It acts like the receptionist for a specific route.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// We write a simple string back to the client
	if r.Method == "GET" {
		fmt.Fprintf(w, "Welcome to the API! You are at the home page.")
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprintf(w, "Message Sent!")
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func createProfileHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Safety Check: Only allow POST methods
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var p Profile

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	fmt.Printf("Created Profile:%+v", p)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the About Page created by Sahil kumar")
	fmt.Println("Endpoint Hit: home page")
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	// getting the name from the query
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintf(w, "Hello, Stranger!")
	} else {
		fmt.Fprintf(w, "Hello, %s!", name)
	}

}

func main() {
	user1 := USER{
		ID:       22344,
		Username: "SAHIL KUMAR",
		Email:    "sahilk1012@gmail.com",
		IsActive: false,
	}

	jsonData, err := json.MarshalIndent(user1, "", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("API RESPONSE")
	fmt.Println(string(jsonData))

	// 1. REGISTER ROUTES
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/create-Profile", createProfileHandler)
	http.HandleFunc("/greet", greetHandler)
	// 2. START THE SERVER
	fmt.Println("Server starting on port 8080...")

	// ListenAndServe starts the web server.
	// It blocks the program (keeps it running) forever, until you stop it.
	// We check for errors just in case the port is busy.
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
