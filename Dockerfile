FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o backend

FROM alpine

WORKDIR /app

COPY --from=builder /app/backend .

EXPOSE 8080

CMD ["./backend"]