package govips

import (
	"fmt"
	"math"
	"sync"

	"github.com/Arimeka/image-transformer/transform/options"
	"github.com/Arimeka/mediaprobe"
	"github.com/davidbyttow/govips/pkg/vips"
)

const (
	lqShrink = 10
)

func NewImage(opts *options.Options, info mediaprobe.Info) (*Image, error) {
	if opts.LowQuality {
		return newLQImage(opts, info)
	}

	image := &Image{
		OriginalInfo:  info,
		TransformOpts: opts,
		mutex:         sync.RWMutex{},
	}

	shrink := image.CalcShrink()
	if shrink < 1 {
		shrink = 1
	}

	var err error
	switch info.MediaSubtype {
	case "jpeg":
		err = image.LoadJPEG(shrink)
	case "png":
		err = image.LoadPNG()
	case "webp":
		err = image.LoadWEBP(shrink)
	default:
		return nil, fmt.Errorf("NewImage: unsupported format %s", info.MediaSubtype)
	}

	return image, err
}

func newLQImage(opts *options.Options, info mediaprobe.Info) (*Image, error) {
	image := &Image{
		OriginalInfo:  info,
		TransformOpts: opts,
		mutex:         sync.RWMutex{},
	}

	shrink := image.CalcShrink()
	if shrink < 1 {
		shrink = 1
	}

	var (
		err         error
		extraShrink float64
		loadShrink  float64
		sigma       float64 = 1
	)

	loadShrink = lqShrink

	if info.Height > 1000 && info.Width > 1000 {
		var m float64 = 1

		if info.Height > info.Width {
			log := math.Log(float64(info.Width))
			sigma = log / 5
			m = float64(info.Width) / log
		} else {
			log := math.Log(float64(info.Height))
			sigma = log / 5
			m = float64(info.Height) / log
		}

		loadShrink = m / 5
	}

	switch info.MediaSubtype {
	case "jpeg":
		err = image.LoadJPEG(int(loadShrink))
		extraShrink = loadShrink / 8
	case "png":
		err = image.LoadPNG()
		extraShrink = loadShrink
	case "webp":
		err = image.LoadWEBP(int(loadShrink))
		extraShrink = 1
	default:
		return nil, fmt.Errorf("NewImage: unsupported format %s", info.MediaSubtype)
	}

	if err = image.Ref.Resize(
		1/extraShrink,
		vips.InputInt("kernel", int(vips.KernelNearest)),
	); err != nil {
		image.Close()
		return nil, err
	}

	if err = image.Ref.Gaussblur(sigma); err != nil {
		image.Close()
		return nil, err
	}

	if err = image.Ref.Resize(
		loadShrink/float64(shrink),
		vips.InputInt("kernel", int(vips.KernelNearest)),
	); err != nil {
		image.Close()
		return nil, err
	}

	return image, err
}

type Image struct {
	Ref *vips.ImageRef

	OriginalInfo  mediaprobe.Info
	TransformOpts *options.Options

	mutex sync.RWMutex
}

func (image *Image) LoadJPEG(shrink int) error {
	image.mutex.Lock()
	defer image.mutex.Unlock()

	switch {
	case shrink > 8:
		shrink = 8
	case shrink > 4:
		shrink = 4
	case shrink > 2:
		shrink = 2
	}

	in, err := vips.Jpegload(
		image.TransformOpts.InFilename,
		vips.InputInt("shrink", shrink),
		vips.InputBool("autorotate", true),
	)
	if err != nil {
		return err
	}
	if image.Ref != nil {
		image.Ref.Close()
	}

	image.Ref = vips.NewImageRef(in, vips.ImageTypeJPEG)

	return nil
}

func (image *Image) LoadPNG() error {
	image.mutex.Lock()
	defer image.mutex.Unlock()

	in, err := vips.Pngload(image.TransformOpts.InFilename)
	if err != nil {
		return err
	}
	if image.Ref != nil {
		image.Ref.Close()
	}

	image.Ref = vips.NewImageRef(in, vips.ImageTypePNG)

	return nil
}

func (image *Image) LoadWEBP(shrink int) error {
	image.mutex.Lock()
	defer image.mutex.Unlock()

	in, err := vips.Webpload(
		image.TransformOpts.InFilename,
		vips.InputInt("shrink", shrink),
	)
	if err != nil {
		return err
	}
	if image.Ref != nil {
		image.Ref.Close()
	}

	image.Ref = vips.NewImageRef(in, vips.ImageTypeWEBP)

	return nil
}

func (image *Image) WriteToFile(filename string) error {
	image.mutex.RLock()
	defer image.mutex.RUnlock()

	var opts []*vips.Option
	if image.TransformOpts.LowQuality {
		opts = append(opts, vips.InputInt("Q", 30))
		opts = append(opts, vips.InputBool("strip", true))
	}

	switch image.TransformOpts.Format {
	case options.FormatPNG:
		return vips.Pngsave(image.Ref.Image(), filename, opts...)
	case options.FormatWEBP:
		return vips.Webpsave(image.Ref.Image(), filename, opts...)
	}

	return vips.Jpegsave(image.Ref.Image(), filename, opts...)
}

func (image *Image) Width() int {
	image.mutex.RLock()
	defer image.mutex.RUnlock()

	return image.Ref.Width()
}

func (image *Image) Height() int {
	image.mutex.RLock()
	defer image.mutex.RUnlock()

	return image.Ref.Height()
}

func (image *Image) CalcShrink() int {
	image.mutex.RLock()
	defer image.mutex.RUnlock()

	var (
		shrinkH = 1
		shrinkW = 1
	)

	if image.TransformOpts.Height != 0 {
		shrinkH = image.OriginalInfo.Height / image.TransformOpts.Height
	}
	if image.TransformOpts.Width != 0 {
		shrinkW = image.OriginalInfo.Width / image.TransformOpts.Width
	}

	if shrinkH > shrinkW {
		return shrinkW
	}

	return shrinkH
}

func (image *Image) calculateHScale() float64 {
	if image.TransformOpts.Width == 0 {
		return 1
	}

	return float64(image.TransformOpts.Width) / float64(image.Ref.Width())
}

func (image *Image) calculateVScale() float64 {
	if image.TransformOpts.Height == 0 {
		return 1
	}

	return float64(image.TransformOpts.Height) / float64(image.Ref.Height())
}

func (image *Image) Close() {
	image.mutex.Lock()
	defer image.mutex.Unlock()

	image.Ref.Close()
}
