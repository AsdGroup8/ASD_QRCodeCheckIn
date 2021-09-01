# COMPILING
FROM golang:1.16-alpine3.14 AS COMPILER

RUN mkdir -p work
WORKDIR /work
COPY . .
ENV CGO_ENABLED=0 GOOS=linux 
RUN go mod tidy
RUN go build . -o /work/server

# RUNNING
FROM alpine:3.14
COPY --from=COMPILER /work/server /
ENTRYPOINT "./server"
