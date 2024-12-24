FROM golang:1.23-alpine3.21 AS builder

WORKDIR /app

COPY . .

RUN go build -ldflags="-s -w" -o main .


FROM alpine:3.21

COPY --from=builder /app/main /app/main

ENTRYPOINT ["/app/main"]
