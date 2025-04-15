FROM golang:1.24 AS builder

WORKDIR /app
COPY ./PtN .
RUN go build -v -o output ./...

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/output .

ENV PORT=8080
EXPOSE 8080

CMD ["./app"]
