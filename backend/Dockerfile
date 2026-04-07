# Build stage
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
# We build from the root, but the main file is in cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

# Run stage
FROM alpine:latest

# Install CA certificates for secure connections (Telegram, Gemini, Turso)
RUN apk --no-go-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Create the session directory
RUN mkdir -p .session

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
