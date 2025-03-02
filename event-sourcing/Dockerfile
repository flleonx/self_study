FROM golang:1.24.0-alpine3.21 AS build-base

WORKDIR /src/

RUN go install github.com/air-verse/air@latest

run apk add --no-cache postgresql-client bash

COPY . .

RUN go mod download
