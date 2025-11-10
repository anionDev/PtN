FROM golang:1.25.1-alpine AS builder

WORKDIR /app
COPY PtN .
RUN go build -o program main.go

FROM alpine:3.22.2
RUN apk --no-cache add ca-certificates bash
WORKDIR /root/
COPY --from=builder /app/program .
EXPOSE 8080
ENTRYPOINT ["./program"]
