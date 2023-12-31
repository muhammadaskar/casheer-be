# Gunakan golang:latest sebagai base image
FROM golang:latest

# RUN chmod 777 ./assets/image/store

# Setel direktori kerja ke /app di dalam container
WORKDIR /app

# RUN mkdir -p ./assets/image/store

COPY .env.example .env

# Salin isi direktori saat ini ke direktori /app di dalam container
COPY . .

# Install dependencies (jika diperlukan)
RUN go get -d -v ./...

# Compile aplikasi Golang
RUN go build -o main .

# Command yang akan dijalankan ketika container dimulai
CMD ["./main"]
