# Build stage
FROM golang:1.25.1 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o weather-api ./cmd/server

# Final stage
FROM debian:bookworm-slim
WORKDIR /app
# Install CA certificates so HTTPS works
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/weather-api .
COPY .env .
CMD ["./weather-api"]