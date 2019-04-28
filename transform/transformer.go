package transform

import (
	"fmt"

	"github.com/Arimeka/image-transformer/configuration"
	"github.com/Arimeka/image-transformer/transform/image/govips"
	"github.com/Arimeka/image-transformer/transform/options"
	"github.com/Arimeka/image-transformer/transform/result"
	"github.com/Arimeka/mediaprobe"
)

type Adapter struct {
	imageClient Client

	Options *options.Options
}

func New(opts *options.Options) (*Adapter, error) {
	var (
		client Client
		err    error
	)

	switch configuration.ImageProcessor() {
	case "govips":
		client, err = govips.New()
	default:
		return nil, fmt.Errorf("unknown image processor %s", configuration.ImageProcessor())
	}
	if err != nil {
		return nil, err
	}

	return &Adapter{
		imageClient: client,
		Options:     opts,
	}, nil
}

func (adapter *Adapter) Process() (result.Result, error) {
	info, _ := mediaprobe.Parse(adapter.Options.InFilename)

	if adapter.Options.Format == options.FormatAuto {
		adapter.Options.SetFormat(info.MediaSubtype)
	}

	if info.MediaType == "image" {
		return adapter.imageClient.Process(adapter.Options, info)
	}

	return result.Result{}, fmt.Errorf("not implemented transformation for %s", info.MediaType)
}
