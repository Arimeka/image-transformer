package options

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/Arimeka/image-transformer/secure"
	"github.com/valyala/fasthttp"
)

type ImageFormat int

const (
	FormatAuto ImageFormat = iota
	FormatJPEG
	FormatPNG
	FormatWEBP
)

type CropGravity int

const (
	GravitySmart CropGravity = iota
	GravityNorth
	GravityCenter
	GravitySouth
	GravityWest
	GravityEast
)

func NewOptions(fHeader *multipart.FileHeader) (*Options, error) {
	filename := fmt.Sprintf("/tmp/%s", secure.RandString(32))
	outFilename := fmt.Sprintf("/tmp/%s", secure.RandString(32))

	if err := fasthttp.SaveMultipartFile(fHeader, filename); err != nil {
		return nil, err
	}

	return &Options{InFilename: filename, OutFilename: outFilename}, nil
}

type Options struct {
	InFilename  string
	OutFilename string
	Width       int
	Height      int
	Enlarge     bool
	Crop        bool
	LowQuality  bool
	Force       bool
	Gravity     CropGravity
	Format      ImageFormat
}

func (opts *Options) Parse(ctx *fasthttp.RequestCtx) error {
	if width, err := fasthttp.ParseUint(ctx.FormValue("width")); err == nil && width > 0 {
		opts.Width = width
	}
	if height, err := fasthttp.ParseUint(ctx.FormValue("height")); err == nil && height > 0 {
		opts.Height = height
	}

	enlarge := string(ctx.FormValue("enlarge"))
	if enlarge == "true" {
		opts.Enlarge = true
	}
	crop := string(ctx.FormValue("crop"))
	if crop == "true" {
		opts.Crop = true
	}
	lq := string(ctx.FormValue("lq"))
	if lq == "true" {
		opts.LowQuality = true
	}
	force := string(ctx.FormValue("force"))
	if force == "force" {
		opts.Force = true
	}

	switch string(ctx.FormValue("gravity")) {
	case "north":
		opts.Gravity = GravityNorth
	case "center":
		opts.Gravity = GravityCenter
	case "south":
		opts.Gravity = GravitySouth
	case "west":
		opts.Gravity = GravityWest
	case "east":
		opts.Gravity = GravityEast
	}

	opts.SetFormat(string(ctx.FormValue("format")))

	return nil
}

func (opts *Options) SetFormat(format string) {
	switch format {
	case "jpeg":
		opts.Format = FormatJPEG
	case "webp":
		opts.Format = FormatWEBP
	case "png":
		opts.Format = FormatPNG
	default:
		opts.Format = FormatAuto
	}
}

func (opts *Options) Free() error {
	os.Remove(opts.InFilename)
	os.Remove(opts.OutFilename)

	return nil
}
