# Fahrrad-Verleih System (Prototyp)

Ein modularer Web-Prototyp zur Verwaltung und Reservierung von Fahrrädern, entwickelt als universitäres Projekt.

## Tech Stack

- **Backend:** Go (`net/http`, Structs, HTTP-Handler)
- **Datenbank:** In-Memory Store (Thread-safe mit `sync.RWMutex`)
- **Frontend:** HTML5 / CSS3 / JavaScript (Responsive Layout, Anker-Navigation)

## Projektstruktur

```text
fahrrad-verleih/
├── internal/
│   ├── domain/      # Kern-Structs (Bike, Person, Rental)
│   ├── handler/     # HTTP-Handler & Routen
│   └── store/       # In-Memory-Datenbank
├── uploads/         # Speicherort für Ausweis-Scans
├── main.go          # Server-Einstiegspunkt
└── index.html       # Interaktives Dashboard und Formulare
