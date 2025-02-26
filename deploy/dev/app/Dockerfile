FROM golang:1.23 AS builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /build/server .
CMD [ "./server" ]