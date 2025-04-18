# syntax=docker/dockerfile:1.3

# --- Build Stage ---
FROM golang:1.21 AS builder

# Setze Arbeitsverzeichnis im Container
WORKDIR /app

# Nur Mod-Dateien zuerst kopieren für Caching
COPY go.mod go.sum ./

# Module herunterladen (nutzt Cache falls nichts geändert wurde)
RUN go mod download

# Restlichen Code kopieren
COPY . .

# Go-Projekt bauen
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# --- Laufzeit Stage (kleines Image) ---
FROM alpine:latest

WORKDIR /app

# Benötigte Zertifikate für HTTPS, etc.
RUN apk --no-cache add ca-certificates

# Binary aus dem Builder-Stage übernehmen
COPY --from=builder /app/main .

# Startbefehl
CMD ["./main"]
