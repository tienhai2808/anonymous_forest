FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o main ./cmd/api

FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

RUN addgroup -S -g 1001 go && adduser -S -D -u 1001 -G go fiber

WORKDIR /app

COPY --from=builder --chown=fiber:go /app/main .

COPY --from=builder --chown=fiber:go /app/configs ./configs

USER fiber

EXPOSE 5000

CMD ["./main"]