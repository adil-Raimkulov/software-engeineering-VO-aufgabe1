package store

import (
	"sync"
	"ski-verleih/internal/domain"
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
	s.items["1"] = domain.RentalItem{ID: "1", Name: "Atomic Redster Ski", Type: "Ski", HourlyRate: 12.0, IsAvailable: true}
	s.items["2"] = domain.RentalItem{ID: "2", Name: "KTM Mountainbike", Type: "Bike", HourlyRate: 15.0, IsAvailable: true}
	return s
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

func (s *MemoryStore) GetItems() []domain.RentalItem {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]domain.RentalItem, 0, len(s.items))
	for _, item := range s.items {
		list = append(list, item)
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
