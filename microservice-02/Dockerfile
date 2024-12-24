FROM golang:1.23-alpine AS builder

WORKDIR /bw01

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bw01 ./cmd/main.go

FROM alpine:latest

WORKDIR /bw01

COPY --from=builder /bw01/bw01 .

RUN adduser -D rishad
USER rishad

EXPOSE 8080

CMD ["./bw01"]
