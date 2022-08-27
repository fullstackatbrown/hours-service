# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

ENV PORT 8080

WORKDIR /app

COPY . ./

EXPOSE ${PORT}

ENTRYPOINT go run . -port=${PORT}