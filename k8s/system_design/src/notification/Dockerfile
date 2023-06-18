FROM golang:1.20.5-bullseye

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /notification-service

CMD ["/notification-service"]
