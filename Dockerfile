# Step 1: Build the Go binary
FROM golang:1.20-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

# Step 2: Minimal runtime image
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/app .
COPY --from=builder /app/views ./views

EXPOSE 3000

CMD ["./app"]

