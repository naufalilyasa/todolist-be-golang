# Stage 1: build
FROM golang:1.25.1-alpine3.22 AS builder

WORKDIR /app

# Copy go.mod dan go.sum dari root
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build binary dari cmd/app/main.go
RUN go build -o /todo-app ./cmd/app/main.go

# Stage 2: minimal image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /todo-app .

EXPOSE 8080

CMD ["./todo-app"]
