FROM golang:alpine

WORKDIR /go/src/app

ADD . .

RUN go build -o ./app

EXPOSE 8080

CMD ["./app"]
