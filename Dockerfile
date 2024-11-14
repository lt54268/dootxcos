FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o dootxcos-app .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/dootxcos-app .

COPY --from=builder /app/.env . 

EXPOSE 6060

CMD ["./dootxcos-app"]
