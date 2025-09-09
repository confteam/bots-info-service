FROM golang AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bots-info-service ./cmd/server

FROM alpine

WORKDIR /app

COPY --from=builder /app/bots-info-service .

CMD ["./bots-info-service"]
