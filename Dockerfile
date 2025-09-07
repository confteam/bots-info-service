FROM golang AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bots-info-service ./cmd/server/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/bots-info-service .

CMD ["./bots-info-service"]
