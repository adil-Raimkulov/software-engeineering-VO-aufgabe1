package main

import (
	"fmt"
	"net/http"
	"os"

	"fahrrad-verleih/internal/handler"
	"fahrrad-verleih/internal/store"
)

func main() {
	os.MkdirAll("./uploads", os.ModePerm)

	memStore := store.NewMemoryStore()
	h := handler.NewHandler(memStore)

	http.HandleFunc("/api/items", h.HandleItems)
	http.HandleFunc("/api/persons", h.HandlePersons)
	http.HandleFunc("/api/rentals", h.HandleRentals)
	http.HandleFunc("/api/stats", h.HandleStats)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	fmt.Println("Fahrrad-Verleih Server gestartet auf http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Serverfehler: %v\n", err)
	}
}