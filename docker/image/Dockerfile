FROM go1.10.3-alpine as builder

Add . /myapp

WORKDIR /myapp

RUN CGO_ENABLED=0 GOOS=linux go build -o httpserver http.go

CMD ["./httpserver"]
