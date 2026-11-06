package domain

import "time"

type Person struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	AusweisTyp  string `json:"ausweis_typ"`
	AusweisNr   string `json:"ausweis_nr"`
	AusweisFile string `json:"ausweis_file"`
}

type RentalItem struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"` // Ski или Bike
	HourlyRate  float64 `json:"hourly_rate"`
	IsAvailable bool    `json:"is_available"`
}

type Rental struct {
	ID            string    `json:"id"`
	PersonName    string    `json:"person_name"`
	ItemName      string    `json:"item_name"`
	StartTime     time.Time `json:"start_time"`
	DurationHours int       `json:"duration_hours"`
	TotalPrice    float64   `json:"total_price"`
}
