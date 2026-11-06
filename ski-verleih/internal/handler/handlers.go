package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"ski-verleih/internal/domain"
	"ski-verleih/internal/store"
)

type Handler struct {
	store *store.MemoryStore
}

func NewHandler(s *store.MemoryStore) *Handler {
	return &Handler{store: s}
}

func (h *Handler) HandlePersons(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(h.store.GetPersons())
		return
	}

	if r.Method == http.MethodPost {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			http.Error(w, fmt.Sprintf("Form error: %v", err), http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		ausweisTyp := r.FormValue("ausweis_typ")
		ausweisNr := r.FormValue("ausweis_nr")

		filename := ""
		file, handlerHeader, err := r.FormFile("ausweis_file")
		if err == nil {
			defer file.Close()
			filename = fmt.Sprintf("%d_%s", time.Now().Unix(), handlerHeader.Filename)
			dst, err := os.Create(filepath.Join("./uploads", filename))
			if err == nil {
				defer dst.Close()
				io.Copy(dst, file)
			}
		}

		person := domain.Person{
			ID:          fmt.Sprintf("%d", time.Now().UnixNano()),
			Name:        name,
			AusweisTyp:  ausweisTyp,
			AusweisNr:   ausweisNr,
			AusweisFile: filename,
		}

		h.store.AddPerson(person)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
		return
	}
}

func (h *Handler) HandleRentals(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(h.store.GetRentals())
		return
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		personName := r.FormValue("person_name")
		itemName := r.FormValue("item_name")
		hours, _ := strconv.Atoi(r.FormValue("duration_hours"))
		
		totalPrice := float64(hours) * 15.0

		rental := domain.Rental{
			ID:            fmt.Sprintf("%d", time.Now().UnixNano()),
			PersonName:    personName,
			ItemName:      itemName,
			StartTime:     time.Now(),
			DurationHours: hours,
			TotalPrice:    totalPrice,
		}

		h.store.AddRental(rental)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
		return
	}
}

func (h *Handler) HandleStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stats := map[string]int{
		"available_ski": len(h.store.GetItems()),
		"rented_ski":    len(h.store.GetRentals()),
		"total_ski":     len(h.store.GetItems()) + len(h.store.GetRentals()),
		"total_persons": len(h.store.GetPersons()),
	}
	json.NewEncoder(w).Encode(stats)
}
