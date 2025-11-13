# Multi-stage para Go
# Etapa 1: Compilar
FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server

# Etapa 2: Runtime mínimo
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server .
CMD ["./server"]

