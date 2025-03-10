FROM golang:alpine3.21 AS builder

WORKDIR /app

# Download and install the dependencies
COPY ./backend/go.mod ./backend/go.sum* ./
RUN go mod download

# Copy the source code
COPY ./backend .
COPY ../proto ./proto
RUN go mod tidy

# Build the Go app
RUN go build -o main .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Install certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Expose port
EXPOSE 8080

# Command to run
CMD ["./main"]