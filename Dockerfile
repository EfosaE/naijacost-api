# Start from the official Go image
FROM golang:1.24-bullseye AS builder

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to take advantage of Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN go build -o server ./cmd/server

# Use a minimal base image for the final build
FROM debian:bullseye-slim

# Set working directory in the final container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/server .

# Expose the port the server listens on (change if needed)
EXPOSE 8080

# Command to run the binary
CMD ["./server"]
