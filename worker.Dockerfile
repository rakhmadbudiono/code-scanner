FROM golang:1.19-alpine

ADD . /app
WORKDIR /app

RUN apk add --update alpine-sdk
RUN apk add --no-cache librdkafka

RUN go mod download
RUN go build -tags musl -o /app/worker /app/cmd/worker/main.go

CMD ["/app/worker"]