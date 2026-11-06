package main

import (
	"fmt"
	"net/http"
	"os"

	"ski-verleih/internal/handler"
	"ski-verleih/internal/store"
)

func main() {
	os.MkdirAll("./uploads", os.ModePerm)

	memStore := store.NewMemoryStore()
	h := handler.NewHandler(memStore)

	// API Routing
	http.HandleFunc("/api/persons", h.HandlePersons)
	http.HandleFunc("/api/rentals", h.HandleRentals)
	http.HandleFunc("/api/stats", h.HandleStats)

	// Static Frontend
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	fmt.Println("Server gestartet auf http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Serverfehler: %v\n", err)
	}
}
