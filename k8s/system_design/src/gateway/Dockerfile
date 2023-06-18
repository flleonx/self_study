FROM golang:1.20.5-alpine3.18

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /gateway-service

EXPOSE 8080

CMD ["/gateway-service"]
