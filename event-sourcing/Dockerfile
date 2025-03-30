FROM golang:1.24.0-alpine3.21 AS build-base

WORKDIR /src/

RUN apk add --no-cache postgresql-client bash

RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . .

RUN go mod download
