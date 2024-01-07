# Stage 1: Building the application
FROM golang:1.21-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/server

# Stage 2: Building a small image
FROM alpine:latest

# Add CA Certificates for SSL
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose the necessary port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
