FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

COPY .env .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"] 