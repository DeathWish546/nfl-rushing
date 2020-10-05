FROM golang:1.15-alpine

ENV GO111MODULE=on

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY go.mod . 
COPY go.sum .

RUN go mod download

COPY . .

RUN go build

EXPOSE 8080

CMD ["./nfl-rushing"]
