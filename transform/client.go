package transform

import (
	"github.com/Arimeka/image-transformer/transform/options"
	"github.com/Arimeka/image-transformer/transform/result"
	"github.com/Arimeka/mediaprobe"
)

type Client interface {
	Process(opts *options.Options, info mediaprobe.Info) (result.Result, error)
}
