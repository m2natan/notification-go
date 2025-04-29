# -------- Stage 1: Build the Go binary --------
    FROM golang:1.24-alpine AS builder

    # Disable CGO for pure static binary
    ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
    
    WORKDIR /app
    
    # Copy go.mod and go.sum first for better caching
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy the rest of the application
    COPY . .
    
    # Build the binary
    RUN go build -o app cmd/collect/main.go
    
    # -------- Stage 2: Lightweight runtime --------
    FROM alpine:latest
    
    WORKDIR /app
    
    # Copy just the binary and .env file
    COPY --from=builder /app/app .
    COPY --from=builder /app/.env .
    
    # Expose ports
    EXPOSE 50051 8080
    
    # Run the application
    ENTRYPOINT ["/app/app"]