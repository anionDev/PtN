FROM golang:1.22-alpine as builder

WORKDIR /app
COPY . .
RUN go build -o proxy .

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/proxy .

ENV PORT=8080
EXPOSE 8080

CMD ["./proxy"]
