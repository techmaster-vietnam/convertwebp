# Convert image to webp

## Chức năng chính
- Convert tất cả các file ảnh (gif, jpeg, jpg, png) trong folder `in` sang định dạng WebP rồi lưu vào folder `out`
- Nếu file đã ở định dạng WebP hay PDF thì copy luôn, không resize, không convert
- Nếu file có kích thước rộng > 1024px hoặc cao > 768px thì resize về để nhỏ hơn 1024x768px, giữ nguyên tỷ lệ rồi mới convert
- Những file định dạng khác thì copy nguyên từ `in` sang `out`

## Install
Chỉ cần cài đặt duy nhất thư viện`vips`

### In Alpine
```
apk add vips --repository=https://dl-cdn.alpinelinux.org/alpine/edge/community
```

## Run
```
go run .
```

## Build
```
go build -o convertwebp .
./convertwebp -in ./in -out ./out -lossyQuality 80
```


## Dockerize then convert using docker
```
docker build -t convertwebp .
docker run --rm -it -v /absolute/localpath/in:/var/in -v /absolute/localpath/out:/var/out convertwebp:latest
```

Ví dụ với đường dẫn cụ thể, tuyệt đối
```
docker run --rm -it -v /Users/cuong/CODE/techmasterweb/convertwebp/in:/var/in -v /Users/cuong/CODE/techmasterweb/convertwebp/out:/var/out convertwebp:latest
```