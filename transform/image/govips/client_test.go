package govips_test

import (
	"testing"

	"github.com/Arimeka/image-transformer/transform/image/govips"
	"github.com/Arimeka/image-transformer/transform/options"
	"github.com/Arimeka/mediaprobe"
)

func initClientForTest(t *testing.T, filename string) (*govips.Client, mediaprobe.Info) {
	client, err := govips.New()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	info, err := mediaprobe.Parse(filename)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	return client, info
}

func TestNew(t *testing.T) {
	_, err := govips.New()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
}

func TestClient_ProcessCrop(t *testing.T) {
	jpegSrc := "./fixtures/image.jpg"
	jpegOut := "./test/image.jpg"
	opts := &options.Options{
		InFilename:  jpegSrc,
		OutFilename: jpegOut,
		Width:       600,
		Height:      600,
		Enlarge:     true,
		Crop:        true,
		Gravity:     options.GravitySouth,
	}

	client, info := initClientForTest(t, jpegSrc)

	result, err := client.Process(opts, info)
	if err != nil {
		t.Errorf("Filename: %s, Unexpected error %v", jpegSrc, err)
	}
	if result.Info.Height != opts.Height {
		t.Errorf("Filename: %s, Unexpected file height. Expected %d. Got %d", opts.InFilename, opts.Height, result.Info.Height)
	}
	if result.Info.Width != opts.Width {
		t.Errorf("Filename: %s, Unexpected file width. Expected %d. Got %d", opts.InFilename, opts.Width, result.Info.Width)
	}
}

func TestClient_ProcessResize(t *testing.T) {
	jpegSrc := "./fixtures/image.jpg"
	jpegOut := "./test/image.jpg"
	opts := &options.Options{
		InFilename:  jpegSrc,
		OutFilename: jpegOut,
		Width:       600,
		Height:      700,
		Enlarge:     true,
	}

	client, info := initClientForTest(t, jpegSrc)

	result, err := client.Process(opts, info)
	if err != nil {
		t.Errorf("Filename: %s, Unexpected error %v", jpegSrc, err)
	}
	if result.Info.Height != 600 {
		t.Errorf("Filename: %s, Unexpected file height. Expected %d. Got %d", opts.InFilename, 600, result.Info.Height)
	}
	if result.Info.Width != opts.Width {
		t.Errorf("Filename: %s, Unexpected file width. Expected %d. Got %d", opts.InFilename, opts.Width, result.Info.Width)
	}
}
