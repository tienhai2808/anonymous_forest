FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o main ./cmd/api

FROM alpine:latest

RUN apk --no-cache add curl ca-certificates tzdata

RUN adduser -D -u 1001 fiber

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app/configs ./configs

USER fiber

EXPOSE 5000

CMD ["./main"]