# Gunakan golang:latest sebagai base image
FROM golang:latest

RUN apk update && apk add --no-cache git

# Setel direktori kerja ke /app di dalam container
WORKDIR /app

RUN mkdir -p ./assets/image/product

COPY .env.example .env

# Salin isi direktori saat ini ke direktori /app di dalam container
COPY . .

# Install dependencies (jika diperlukan)
RUN go get -d -v ./...

# Compile aplikasi Golang
RUN go build -o main .

# Command yang akan dijalankan ketika container dimulai
CMD ["./main"]
