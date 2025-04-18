# Build-Stage mit notwendigen Tools
FROM golang:1.21-alpine AS builder

# Git und SSL-Certificates für Module installieren
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Kopiere Go Mod-Dateien (vor dem restlichen Code für besseres Caching)
COPY go.mod go.sum ./

# Downloade Abhängigkeiten (mit Proxy-Fallback)
RUN go mod download -x || go mod download -x

# Kopiere den gesamten Code
COPY . .

# Statisches Binary bauen
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o /main ./main.go

# Final-Stage mit notwendigen Runtime-Dependencies
FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY --from=builder /main .

EXPOSE 8080
CMD ["./main"]