package store

import (
	"sync"

	"fahrrad-verleih/internal/domain"
)

type MemoryStore struct {
	mu      sync.RWMutex
	persons map[string]domain.Person
	items   map[string]domain.RentalItem
	rentals map[string]domain.Rental
}

func NewMemoryStore() *MemoryStore {
	s := &MemoryStore{
		persons: make(map[string]domain.Person),
		items:   make(map[string]domain.RentalItem),
		rentals: make(map[string]domain.Rental),
	}

	// Pre-populated items matching professor's UI
	s.items["1"] = domain.Bike{ID: "1", Name: "KTM Mountainbike", Details: "MTB, L", ItemType: "Bike", HourlyRate: 15.0}
	s.items["2"] = domain.Bike{ID: "2", Name: "Cube Citybike", Details: "City, M", ItemType: "Bike", HourlyRate: 12.0}
	s.items["3"] = domain.Ski{ID: "3", Name: "Atomic Redster", Details: "Slalom, 170 cm", ItemType: "Ski", HourlyRate: 12.0}

	return s
}

func (s *MemoryStore) AddItem(item domain.RentalItem) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[item.GetID()] = item
}

func (s *MemoryStore) GetItems() []domain.RentalItem {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]domain.RentalItem, 0, len(s.items))
	for _, item := range s.items {
		list = append(list, item)
	}
	return list
}

func (s *MemoryStore) AddPerson(p domain.Person) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.persons[p.ID] = p
}

func (s *MemoryStore) GetPersons() []domain.Person {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]domain.Person, 0, len(s.persons))
	for _, p := range s.persons {
		list = append(list, p)
	}
	return list
}

func (s *MemoryStore) AddRental(r domain.Rental) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.rentals[r.ID] = r
}

func (s *MemoryStore) GetRentals() []domain.Rental {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]domain.Rental, 0, len(s.rentals))
	for _, r := range s.rentals {
		list = append(list, r)
	}
	return list
}