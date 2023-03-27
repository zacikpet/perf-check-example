FROM golang:alpine AS builder

RUN apk update && apk add --no-cache 'git=~2'

ENV GO111MODULE=on

WORKDIR $GOPATH/src/packages/goginapp/

COPY . .

RUN go get -d -v

RUN go build -o /go/main .

###

FROM alpine:3

WORKDIR /

COPY --from=builder /go/main /go/main

EXPOSE 8080

WORKDIR /go

ENTRYPOINT ["/go/main"]
