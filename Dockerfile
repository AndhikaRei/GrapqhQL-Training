FROM golang:1.16.4-alpine3.13

RUN mkdir -p /app

COPY . /app

WORKDIR /app

# RUN go mod tidy
RUN go mod vendor
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN sleep 10
CMD ["go","run", "./cmd/main/server.go"]


