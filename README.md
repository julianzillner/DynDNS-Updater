# Simple DynDNS Updater

This is a minimalistic DynDNS updater service.

## 📦 Setup

1. Clone the repository.
2. Set the required environment variables:
   - `URL` – The DynDNS update URL.
   - `INTERVAL` – Update interval in seconds (e.g., `300` for 5 minutes).

## 🚀 Deploy with Docker Compose

```bash
docker-compose up -d
````

The service will send an update request to the specified URL at the defined interval.


# New Functions