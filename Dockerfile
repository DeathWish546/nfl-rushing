FROM golang:1.15-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY . .

RUN go mod download

RUN go build

EXPOSE 8080

CMD ["./nfl-rushing"]
