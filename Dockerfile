FROM golang:latest
WORKDIR /app
ADD . .
RUN go mod download
CMD go run ./cmd/api/main.go
