FROM golang:1.10.3-alpine

ADD . /myapp

WORKDIR /myapp

RUN CGO_ENABLED=0 GOOS=linux go build -o httpserver httpserver.go

CMD ["./httpserver"]
