# Gunakan Go sebagai base image
FROM golang:1.21

# Set working directory
WORKDIR /app

# Copy semua file ke container
COPY . .

# Copy file private.pem ke container
COPY private.pem /app/private.pem

# Install dependencies
RUN go mod tidy

# Jalankan aplikasi
CMD ["go", "run", "main.go"]
