FROM golang:1.20.5-bullseye

WORKDIR /app

RUN apt-get update && apt-get install -y ffmpeg

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /converter-service

CMD ["/converter-service"]
