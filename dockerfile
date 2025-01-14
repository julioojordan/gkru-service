# Gunakan Go sebagai base image
FROM golang:1.20

# Set working directory
WORKDIR /app

# Copy semua file ke container
COPY . .

# Install dependencies
RUN go mod tidy

# Build aplikasi
RUN go build -o main .

# Jalankan aplikasi
CMD ["./main"]
