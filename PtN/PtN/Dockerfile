# Build stage
FROM golang:1.24.2-alpine AS builder

WORKDIR /app
COPY PtN .
RUN go build -o server main.go

# Minimal final image with shell
FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8080
ENTRYPOINT ["./server"]
