package domain

import "time"

type Person struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	AusweisTyp  string `json:"ausweis_typ"`
	AusweisNr   string `json:"ausweis_nr"`
	AusweisFile string `json:"ausweis_file"`
}

type RentalItem interface {
	GetID() string
	GetName() string
	GetDetails() string
	GetItemType() string
	GetHourlyRate() float64
}

type Bike struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Details    string  `json:"details"`
	ItemType   string  `json:"type"`
	HourlyRate float64 `json:"hourly_rate"`
}

func (b Bike) GetID() string         { return b.ID }
func (b Bike) GetName() string       { return b.Name }
func (b Bike) GetDetails() string    { return b.Details }
func (b Bike) GetItemType() string   { return b.ItemType }
func (b Bike) GetHourlyRate() float64 { return b.HourlyRate }

type Ski struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Details    string  `json:"details"`
	ItemType   string  `json:"type"`
	HourlyRate float64 `json:"hourly_rate"`
}

func (s Ski) GetID() string         { return s.ID }
func (s Ski) GetName() string       { return s.Name }
func (s Ski) GetDetails() string    { return s.Details }
func (s Ski) GetItemType() string   { return s.ItemType }
func (s Ski) GetHourlyRate() float64 { return s.HourlyRate }

type Rental struct {
	ID            string    `json:"id"`
	PersonName    string    `json:"person_name"`
	ItemName      string    `json:"item_name"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	DurationHours int       `json:"duration_hours"`
	Status        string    `json:"status"`
	TotalPrice    float64   `json:"total_price"`
}