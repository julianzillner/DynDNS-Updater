FROM golang:1.21

WORKDIR /app

# Erst die Mod-Dateien kopieren (für besseres Caching)
COPY go.mod go.sum ./
RUN go mod download

# Dann den Rest kopieren
COPY . .

# Build
RUN go build -o main .

# Ausführung
CMD ["./main"]