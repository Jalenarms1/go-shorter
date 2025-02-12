FROM golang:1.23 AS builder

WORKDIR /app

COPY go.sum go.mod ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/go-shorter

FROM debian:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080


CMD [ "./main" ]