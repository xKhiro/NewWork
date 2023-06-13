# NewWork

# Workspace Booking System

ToDo

## Anforderungen

- Go 1.16 oder höher
- Gorilla Mux Paket
- Node v20 oder höher

## Setup

### Backend

Um das Gorilla Mux Paket zu installieren, öffnen Sie die Konsole und geben Sie den folgenden Befehl ein:

```shell
go get -u github.com/gorilla/mux
```

### Frontend

Navigieren Sie in der Konsole zum Projektverzeichnis `Frontend` und geben Sie den folgenden Befehl ein, um alle Abhängigkeiten zu installieren.

```shell
npm install
```

## Ausführen des Projekts

Navigieren Sie in der Konsole zum Projektverzeichnis `Backend` und führen Sie den folgenden Befehl aus, um das Projekt zu starten:

```shell
go run main.go
```

Der Server läuft standardmäßig auf Port 8000.

Um den Entwicklungsserver für das Frontend zu starten, öffnen Sie eine neue Konsole und navigieren in das Verzeichnis `Frontend`. Anschließend geben Sie den folgenden Befehl ein:

```shell
npm run dev
```

Die Webseite ist über <http://localhost:5173> im Webbrowser aufrufbar.

## API Endpunkte

Die API unterstützt die folgenden Endpunkte:

- GET /users/{personId}/bookings: Erhalten Sie alle Buchungen für einen Benutzer
- POST /users/{personId}/bookings: Erstellen Sie eine neue Buchung für einen Benutzer
- DELETE /users/{personId}/bookings/{bookingId}: Stornieren Sie eine Buchung für einen Benutzer
- GET /workspaces: Erhalten Sie alle Arbeitsbereiche
