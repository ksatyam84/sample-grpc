FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server main.go

# Final minimal image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 5001

CMD ["./server"]