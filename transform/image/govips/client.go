package govips

import (
	"runtime"

	"github.com/Arimeka/image-transformer/configuration"
	"github.com/Arimeka/image-transformer/transform/options"
	"github.com/Arimeka/image-transformer/transform/result"
	"github.com/Arimeka/mediaprobe"
	"github.com/davidbyttow/govips/pkg/vips"
)

func init() {
	vips.Startup(&vips.Config{ReportLeaks: configuration.IsDevelopment(), ConcurrencyLevel: runtime.NumCPU()})
}

func New() (*Client, error) {
	return &Client{}, nil
}

type Client struct{}

func (client *Client) Process(opts *options.Options, info mediaprobe.Info) (result.Result, error) {
	image, err := NewImage(opts, info)
	if err != nil {
		return result.Result{}, err
	}
	defer image.Close()

	if image.TransformOpts.Crop {
		err = image.Crop()
	} else {
		err = image.Resize()
	}
	if err != nil {
		return result.Result{}, err
	}

	err = image.WriteToFile(opts.OutFilename)
	if err != nil {
		return result.Result{}, err
	}

	outInfo, err := mediaprobe.New(opts.OutFilename)
	if err != nil {
		return result.Result{}, err
	}
	outInfo.Width = image.Width()
	outInfo.Height = image.Height()

	return result.Result{Filename: opts.OutFilename, Info: *outInfo}, nil
}
