# Gunakan Go sebagai base image
FROM golang:1.21

# Set working directory di dalam container
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Install dependencies (go.mod dan go.sum harus ada)
RUN go mod tidy && go mod verify

# Jalankan aplikasi langsung menggunakan `go run`
CMD ["go", "run", "main.go"]
