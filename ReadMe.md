# Convert image to webp
The WebP format always has a smaller size compared to other image formats. The task of this docker image/source code is to convert all image files in the `in` folder to WebP format and save them in the `out` folder.

![Size after compression](result.jpg)

### Main Features
- Convert all image files (gif, jpeg, jpg, png) in the `in` folder to WebP format and save them in the `out` folder
- If the file is already in WebP or PDF format, just copy it without resizing or converting
- If the file has a width > 1024px or height > 768px, resize it to be smaller than 1024x768px, maintaining the aspect ratio before converting
- Other file formats are copied directly from `in` to `out`

### Use without installation
```
docker run --rm -it -v /absolute/localpath/in:/var/in -v /absolute/localpath/out:/var/out convertwebp:latest
```

### Installation for development
Only the `vips` library needs to be installed. On MacOSX, run:
```
brew install vips
```

### Test run
Ensure there is an `out` directory in convertwebp before running
```
git clone https://github.com/techmaster-vietnam/convertwebp
cd convertwebp
go run .
```

### Build a binary file and run
```
go build -o convertwebp .
./convertwebp -in ./in -out ./out -lossyQuality 80
```

### Package as a Docker image and run
```
docker build -t convertwebp .
docker run --rm -it -v /absolute/localpath/in:/var/in -v /absolute/localpath/out:/var/out convertwebp:latest
```

Example with specific, absolute paths
```
docker run --rm -it -v /Users/cuong/CODE/techmasterweb/convertwebp/in:/var/in -v /Users/cuong/CODE/techmasterweb/convertwebp/out:/var/out convertwebp:latest
```

### Reuse the source code
Just copy the `convertwebp` directory into your project