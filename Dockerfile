# syntax=docker/dockerfile:1

FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o shoffice ./cmd/app/main.go

FROM alpine

WORKDIR /build

COPY web/ web/
COPY .env .

COPY --from=builder /build/shoffice /build/shoffice

EXPOSE 8080

CMD [ "./shoffice" ]