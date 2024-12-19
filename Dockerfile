# syntax=docker/dockerfile:1

FROM golang:1.23.4

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-health-check

EXPOSE 80

CMD ["/go-health-check"]