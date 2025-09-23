# Build stage
FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o weather-api ./cmd/server

# Final stage
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/weather-api .
CMD ["./weather-api"]