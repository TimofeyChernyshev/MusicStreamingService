FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY service-streaming/ ./
COPY db/ ./db/
COPY models/ ./models/
COPY storage/song ./storage/song

RUN go build -o /app/streaming .

FROM alpine
WORKDIR /app
COPY --from=builder /app/streaming .
COPY --from=builder /app/storage/song ./storage/song

EXPOSE 3001
CMD ["./streaming"]