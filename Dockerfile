# COMPILING
FROM golang:1.16-alpine3.14 AS compiler

RUN mkdir -p /go/src/server/log
ADD . /go/src/server

WORKDIR /go/src/server

ENV CGO_ENABLED=0 GOOS=linux GOPROXY=https://goproxy.cn
RUN go mod tidy
RUN go build -o build/server github.com/AsdGroup8/ASD_QRCodeCheckIn

# RUNNING
FROM alpine:3.14
COPY --from=compiler /go/src/server/build/server .
ENTRYPOINT "./server"
