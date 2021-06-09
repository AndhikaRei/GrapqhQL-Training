FROM golang:1.16.5-buster

RUN mkdir -p /app

COPY . /app

WORKDIR /app

RUN go mod vendor
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

CMD ["go","run", "./cmd/main/server.go"]


