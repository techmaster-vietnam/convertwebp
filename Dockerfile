# Sử dụng Golang image chính thức
FROM golang:alpine AS builder

# Cài đặt thư viện vips
RUN apk add musl-dev vips-dev gcc upx

# Tạo thư mục làm việc
WORKDIR /app

# Sao chép các file go.mod và go.sum để cài đặt các dependencies
COPY go.mod go.sum ./

# Tải các dependencies
RUN go mod download

# Sao chép mã nguồn vào container
COPY . .

RUN go mod tidy

# Biên dịch ứng dụng
RUN go build -o myapp && upx -9 myapp

# Tạo một image nhỏ hơn để chạy ứng dụng
FROM alpine:latest

# Cài đặt thư viện vips
RUN apk add --no-cache vips --repository=https://dl-cdn.alpinelinux.org/alpine/edge/community

# Tạo thư mục làm việc
WORKDIR /app

# Sao chép ứng dụng từ builder
COPY --from=builder /app/myapp .

# Chạy ứng dụng
CMD ["./myapp", "-in", "/var/in", "-out", "/var/out"]