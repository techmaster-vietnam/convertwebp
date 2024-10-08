# Convert image to webp

## Install

Cần cài đặt những thư viện sau đây:
* webp
* vips
* graphicsmagick

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