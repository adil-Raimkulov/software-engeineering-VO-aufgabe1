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

	"fahrrad-verleih/internal/domain"
	"fahrrad-verleih/internal/store"
)

type Handler struct {
	store *store.MemoryStore
}

func NewHandler(s *store.MemoryStore) *Handler {
	return &Handler{store: s}
}

func (h *Handler) HandleItems(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(h.store.GetItems())
		return
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.FormValue("name")
		details := r.FormValue("details")
		rate, _ := strconv.ParseFloat(r.FormValue("hourly_rate"), 64)

		item := domain.Bike{
			ID:         fmt.Sprintf("%d", time.Now().UnixNano()),
			Name:       name,
			Details:    details,
			ItemType:   "Bike",
			HourlyRate: rate,
		}
		h.store.AddItem(item)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
		return
	}
}

func (h *Handler) HandlePersons(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(h.store.GetPersons())
		return
	}

	if r.Method == http.MethodPost {
		r.ParseMultipartForm(10 << 20)
		name := r.FormValue("name")
		ausweisTyp := r.FormValue("ausweis_typ")
		ausweisNr := r.FormValue("ausweis_nr")

		filename := ""
		file, handlerHeader, err := r.FormFile("ausweis_file")
		if err == nil {
			defer file.Close()
			filename = fmt.Sprintf("%d_%s", time.Now().Unix(), handlerHeader.Filename)
			dst, _ := os.Create(filepath.Join("./uploads", filename))
			defer dst.Close()
			io.Copy(dst, file)
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
		startStr := r.FormValue("start_time")
		endStr := r.FormValue("end_time")

		startTime, _ := time.Parse("2006-01-02T15:04", startStr)
		endTime, _ := time.Parse("2006-01-02T15:04", endStr)
		if startTime.IsZero() {
			startTime = time.Now()
		}
		if endTime.IsZero() || endTime.Before(startTime) {
			endTime = startTime.Add(1 * time.Hour)
		}

		duration := int(endTime.Sub(startTime).Hours())
		if duration < 1 {
			duration = 1
		}

		var hourlyRate float64 = 15.0
		for _, item := range h.store.GetItems() {
			if item.GetName() == itemName {
				hourlyRate = item.GetHourlyRate()
				break
			}
		}

		rental := domain.Rental{
			ID:            fmt.Sprintf("%d", time.Now().UnixNano()),
			PersonName:    personName,
			ItemName:      itemName,
			StartTime:     startTime,
			EndTime:       endTime,
			DurationHours: duration,
			Status:        "Ausgeliehen",
			TotalPrice:    float64(duration) * hourlyRate,
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
		"available_bikes": len(h.store.GetItems()),
		"rented_bikes":    len(h.store.GetRentals()),
		"total_bikes":     len(h.store.GetItems()) + len(h.store.GetRentals()),
		"total_persons":   len(h.store.GetPersons()),
	}
	json.NewEncoder(w).Encode(stats)
}