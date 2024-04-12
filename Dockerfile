FROM golang:latest

COPY . /go/src/app

WORKDIR /go/src/app
RUN go mod download

RUN go build -o bot .

EXPOSE 8080

CMD ["./bot"]
