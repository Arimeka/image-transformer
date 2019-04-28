package govips_test

import (
	"testing"

	"github.com/Arimeka/image-transformer/transform/image/govips"
	"github.com/Arimeka/image-transformer/transform/options"
	"github.com/Arimeka/mediaprobe"
)

const (
	benchmarkImageJPEG = "./fixtures/image.jpg"
	benchmarkImagePNG  = "./fixtures/image.png"
	benchmarkImageWEBP = "./fixtures/image.webp"

	outFilename = "/dev/null"

	toWidth     = 300
	toHeight    = 250
	needEnlarge = false
	format      = "webp"
)

func initClient(b *testing.B, filename string) (*govips.Client, mediaprobe.Info, *options.Options) {
	client, err := govips.New()
	if err != nil {
		b.Fatal(err)
	}
	info, err := mediaprobe.Parse(filename)
	if err != nil {
		b.Fatal(err)
	}
	opts := transformOptions(filename)

	return client, info, opts
}

func transformOptions(filename string) *options.Options {
	opts := &options.Options{
		InFilename:  filename,
		OutFilename: outFilename,
		Width:       toWidth,
		Height:      toHeight,
		Enlarge:     needEnlarge,
	}

	switch format {
	case "webp":
		opts.Format = options.FormatWEBP
	case "png":
		opts.Format = options.FormatPNG
	default:
		opts.Format = options.FormatJPEG
	}

	return opts
}

func BenchmarkClient_ProcessJPEGResize(b *testing.B) {
	client, info, opts := initClient(b, benchmarkImageJPEG)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Process(opts, info)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkClient_ProcessPNGResize(b *testing.B) {
	client, info, opts := initClient(b, benchmarkImagePNG)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Process(opts, info)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkClient_ProcessWEBPResize(b *testing.B) {
	client, info, opts := initClient(b, benchmarkImageWEBP)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Process(opts, info)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkClient_ProcessJPEGCrop(b *testing.B) {
	client, info, opts := initClient(b, benchmarkImageJPEG)
	opts.Crop = true
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Process(opts, info)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkClient_ProcessPNGCrop(b *testing.B) {
	client, info, opts := initClient(b, benchmarkImagePNG)
	opts.Crop = true
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Process(opts, info)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkClient_ProcessWEBPCrop(b *testing.B) {
	client, info, opts := initClient(b, benchmarkImageWEBP)
	opts.Crop = true
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Process(opts, info)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkClient_ProcessJPEGLQCrop(b *testing.B) {
	client, info, opts := initClient(b, benchmarkImageJPEG)
	opts.Crop = true
	opts.LowQuality = true
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Process(opts, info)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkClient_ProcessPNGLQCrop(b *testing.B) {
	client, info, opts := initClient(b, benchmarkImagePNG)
	opts.Crop = true
	opts.LowQuality = true
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Process(opts, info)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkClient_ProcessWEBPLQCrop(b *testing.B) {
	client, info, opts := initClient(b, benchmarkImageWEBP)
	opts.Crop = true
	opts.LowQuality = true
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Process(opts, info)
		if err != nil {
			b.Fatal(err)
		}
	}
}
