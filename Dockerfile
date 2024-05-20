# Builder stage
FROM golang:1.19 as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

# Final stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main /app/main
RUN pwd && ls -l /app

RUN chmod +x /app/main

CMD ["sh", "-c", "./main"]
