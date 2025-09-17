# Gunakan image Go resmi sebagai base image
FROM golang:1.25.1 AS builder

# Atur working directory di dalam container
WORKDIR /app

# Salin file go.mod dan go.sum untuk mengunduh dependency
COPY go.mod ./
COPY go.sum ./

# Unduh semua dependency Go
RUN go mod download

# Salin semua file source code
COPY . .

# Bangun aplikasi Go Anda
# -ldflags="-s -w" digunakan untuk mengurangi ukuran binary
# -o main menamai executable sebagai 'main'
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o main .

# Gunakan base image yang lebih kecil untuk production
FROM alpine:3.18

# Atur working directory di dalam container
WORKDIR /app

# Salin executable 'main' dari build stage sebelumnya
COPY --from=builder /app/main .

# Expose port yang digunakan aplikasi Go
EXPOSE 8085

# Jalankan executable saat container dijalankan
CMD ["./main"]