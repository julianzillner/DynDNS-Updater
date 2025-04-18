# syntax=docker/dockerfile:1.3

# --- STAGE 1: Build ---
FROM golang:1.21 AS builder

WORKDIR /app

# Nur go.mod und go.sum für Caching
COPY go.mod go.sum ./
RUN go mod download

# Rest des Codes kopieren und bauen
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# --- STAGE 2: Run ---
FROM alpine:latest

WORKDIR /app

# Zertifikate für HTTPS-Verbindungen
RUN apk --no-cache add ca-certificates

# Binary vom Builder übernehmen
COPY --from=builder /app/main .

# Standard-Ausführung
CMD ["./main"]
