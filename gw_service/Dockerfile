# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app
ADD . .
RUN go mod download
RUN go get github.com/cosmtrek/air
RUN go install github.com/cosmtrek/air

ENTRYPOINT ["air"]