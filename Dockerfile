# -------- Stage 1: Build the Go binary --------
    FROM golang:1.24-alpine AS builder

    # Enable Go modules and disable CGO for static binary
    ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
    
    WORKDIR /app
    
    # Copy go.mod and go.sum to leverage Docker layer caching
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy all project files, including .env file
    COPY . .
    
    # Build the binary
    RUN go build -o app cmd/collect/main.go
    
    # -------- Stage 2: Lightweight Alpine runtime --------
    FROM alpine:latest
    
    # Set the working directory
    WORKDIR /app
    
    # Copy the compiled Go binary and .env file from the builder stage
    COPY --from=builder /app/app .
    COPY --from=builder /app/.env .
    
    # Expose ports for gRPC and HTTP
    EXPOSE 50051
    EXPOSE 8080
    
    # Run the binary
    ENTRYPOINT ["/app/app"]
    