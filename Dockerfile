# Menggunakan image golang sebagai base image
FROM golang:latest

# Setel working directory di dalam container
WORKDIR /app

# Copy the .env.example file to the container
COPY .env.example .env

# Rename the file from .env.example to .env
RUN mv .env.example .env

# Menyalin file go.mod dan go.sum ke working directory
COPY go.mod go.sum ./

# Menjalankan perintah go mod untuk mendownload dependensi
RUN go mod download

# Menyalin seluruh kode sumber ke working directory
COPY . .

# Mengompilasi aplikasi Golang
RUN go build -o main .

# Menjalankan aplikasi saat container dijalankan
CMD ["./main"]