FROM golang:1.19.4-alpine3.16 AS builder

COPY . /github.com/POMBNK/restAPI/
WORKDIR /github.com/POMBNK/restAPI/

RUN go mod download
RUN go build -o ./bin/rest-api cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/POMBNK/restAPI/bin/rest-api .
COPY --from=0 /github.com/POMBNK/restAPI/config.yml .
COPY --from=0 /github.com/POMBNK/restAPI/.env .

EXPOSE 8080

CMD ["./rest-api"]

