# ---------- build stage ----------
FROM golang:1.24 AS builder

WORKDIR /app

COPY ../go.mod go.sum ./
RUN go mod download

COPY .. .

RUN go build -o server ./cmd/server

# ---------- run stage ----------
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 50051 8080

CMD ["./server"]
