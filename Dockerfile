# Stage 1: Build binary
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Copy go.mod dan go.sum terlebih dahulu (supaya cache build efisien)
COPY go.mod ./
RUN go mod download

# Copy semua source code
COPY . .

# Build aplikasi (output binary bernama app)
RUN go build -o app main.go


# Stage 2: Runtime image (lebih kecil)
FROM alpine:3.19
WORKDIR /app

# Copy binary dari stage builder
COPY --from=builder /app/app .

# Set port default (ubah sesuai kebutuhan)
EXPOSE 8080

# Jalankan aplikasi
CMD ["./app"]

