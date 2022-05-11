# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

ADD . /app

RUN go mod download

RUN go build -o /docker-go-template

EXPOSE 8000

CMD [ "/docker-go-template" ]