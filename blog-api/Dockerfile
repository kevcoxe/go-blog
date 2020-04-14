FROM golang

WORKDIR /app

ADD . /app

RUN go build -o app .

ENTRYPOINT ["./app"]

