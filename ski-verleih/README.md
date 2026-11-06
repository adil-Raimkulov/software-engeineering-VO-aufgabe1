# Ski-Verleih System


Ein modularer Web-Prototyp zur Verwaltung und Reservierung von Skiausrüstung, entwickelt im Rahmen des Universitätsstudiums (Paris Lodron Universität Salzburg).

## Tech Stack

- **Backend:** Go (`net/http`, Structs, HTTP-Handler)
- **Datenbank:** In-Memory Store (Thread-safe mit `sync.RWMutex`)
- **Frontend:** HTML5 / CSS3 / JavaScript (Responsive Layout, Anker-Navigation)

## Projektstruktur

```text
ski-verleih/
├── internal/
│   ├── domain/      # Kern-Structs (Ski, Person, Rental)
│   ├── handler/     # HTTP-Handler & Routen
│   └── store/       # In-Memory-Datenbank
├── uploads/         # Speicherort für Ausweis-Scans
├── main.go          # Server-Einstiegspunkt
└── index.html       # Interaktives Dashboard und Formulare
