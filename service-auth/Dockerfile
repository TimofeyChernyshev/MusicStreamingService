FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY service-auth/ ./ 
COPY db/ ./db/
COPY models/ ./models/
COPY utils/ ./utils/

RUN go build -o /app/auth .

FROM alpine
WORKDIR /app
COPY --from=builder /app/auth .

EXPOSE 3002
CMD ["./auth"]