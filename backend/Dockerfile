# Etapa de build
FROM golang:1.25 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o server .

# Etapa de execução
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
