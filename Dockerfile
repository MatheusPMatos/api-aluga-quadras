
FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /api

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /api .

EXPOSE 80

CMD ["./api"]
