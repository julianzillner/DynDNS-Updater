FROM golang:1.21-alpine

WORKDIR /app

# Kopiere alle notwendigen Dateien
COPY . .

# Baue die Anwendung
RUN go build -o main .

# Führe die Anwendung aus
CMD ["./main"]