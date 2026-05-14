# Build Stage
FROM golang:1.26.3-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main main.go

# Production Stage
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .
COPY --from=builder /app/.env.example .env

# Create logs directory
RUN mkdir -p storage/logs

# Expose port
EXPOSE 8000

CMD ["./main"]
