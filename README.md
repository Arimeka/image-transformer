# Image Transformer

MVP. Crop and resize images. Designed to work under proxy.

## Requirements

libmagic, ffmpeg libraries version 4.x, libvips.

## Usage

Send the file and options, get the converted file.

```bash
 curl -F "file=@image.png" -F "width=320" -F "height=240" -F "crop=true" -F "lq=false" -F "format=jpeg" http://localhost:8080/transform -o test.jpg
```

### Request options

* `file` - string, original file
* `width` - integer, resize to width
* `height` - integer, resize to height
* `enlarge` - boolean, enlarge image if smaller than requested size
* `crop` - boolean, crop image
* `lq` - boolean, genreate low-quality image. It has a much smaller size, but it takes longer to convert.
* `force` - boolean, ignore original aspect ratio.
* `gravity` - enum string, choose area to crop. By default uses smart crop. Expected values:
  * north
  * center
  * south
  * west
  * east
* `format` - enum string, choose output image format. By default uses input file format. Expected values:
  * jpeg
  * webp
  * png

Some options do not work without other options. For example, it is impossible to crop image if only width 
was specified.

## Limitations

Only jpeg, webp or png expected. Other formats not implemented.
