FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY server/ ./ 
COPY db/ ./db/
COPY models/ ./models/
COPY frontend/ ./frontend/
COPY storage/cover ./storage/cover

RUN go build -o /app/server

FROM alpine
WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/frontend ./frontend
COPY --from=builder /app/storage/cover ./storage/cover

EXPOSE 8080
CMD ["./server"]