# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

ENV PORT 8080

WORKDIR /app

COPY . ./

RUN go build .

EXPOSE ${PORT}

ENTRYPOINT ./hours -port=${PORT}