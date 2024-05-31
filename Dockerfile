FROM golang:1.22-alpine3.20 AS builder

COPY . /github.com/xndrg/scheduly
WORKDIR /github.com/xndrg/scheduly

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/xndrg/scheduly/bin/bot .
COPY --from=0 /github.com/xndrg/scheduly/configs configs/

CMD ["./bot"]
