# 🌤️ Go Weather API

A lightweight **Weather API wrapper** built in Go, using:

- [Echo](https://echo.labstack.com/) for HTTP routing  
- [Zerolog](https://github.com/rs/zerolog) for structured logging  
- [Viper](https://github.com/spf13/viper) for configuration  
- [go-redis/v9](https://github.com/redis/go-redis) for caching
- Docker & Docker Compose for containerization  

> Inspired by the [roadmap.sh Weather API project](https://roadmap.sh/projects/weather-api-wrapper-service). Used [Visual Crossing’s API](https://www.visualcrossing.com/weather-api/) as 3rd party weather service.

---

## 📂 Project Structure

- [cmd/server](cmd/server) → entry point (main.go)
- [config](config) → configuration loading with Viper
- [internal/api](internal/api) → HTTP handlers
- [internal/cache](internal/cache) → Redis cache logic
- [internal/logger](internal/logger) → Zerolog setup
- [internal/client](internal/client) → 3rd party API call
- [internal/service](internal/service) → business logic
- [docker-compose.yml](docker-compose.yml) → Docker services (API + Redis)
- [Dockerfile](Dockerfile) → multi-stage Docker build

## 🚀 Getting Started

### Prerequisites
- [Go](https://go.dev/) 1.25+
- [Docker](https://www.docker.com/), Docker Compose & [Colima](https://github.com/abiosoft/colima) (if on macOS)
- Redis (if want to run locally, pulled automatically via Docker Compose)

### Clone & Run
```bash
git clone https://github.com/gonzalocarreteroh/weather-api.git
cd weather-api
```
Configure environment variables with **Viper**. Create a .env file inside the root directory and populate it as so:
```bash
SERVER_PORT=8080
REDIS_URL=redis:6379
WEATHER_API_KEY=your_api_key_here # create one at https://www.visualcrossing.com/weather-api/
LOG_TYPE=dev   # or "prod" for JSON logs
```

Run locally:
```bash
go run ./cmd/server
```
Run with Docker:
```bash
docker-compose up --build
```

## Usage
**Endpoint**
```
GET /weather/:city
```
**Example**

Type in your terminal:
```bash
curl http://localhost:8080/weather/Toronto
```
**Response**
```json
{
  "description": "Similar temperatures continuing with no rain expected.",
  "currentConditions": {
    "temp": 14
  }
}
```
