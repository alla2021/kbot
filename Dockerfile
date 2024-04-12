FROM quay.io/projectquay/golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o bot .

FROM quay.io/projectquay/golang:1.20

WORKDIR /app

COPY --from=builder /app/bot .

EXPOSE 8080

CMD ["./bot"]
