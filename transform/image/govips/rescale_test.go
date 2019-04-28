package govips_test

import (
	"testing"

	"github.com/Arimeka/image-transformer/transform/image/govips"
	"github.com/Arimeka/image-transformer/transform/options"
)

func TestImage_Resize(t *testing.T) {
	webpSrc := "./fixtures/image.webp"
	webpOut := "./test/image.webp"
	opts := &options.Options{
		InFilename:  webpSrc,
		OutFilename: webpOut,
		Width:       600,
		Height:      350,
	}

	_, info := initClientForTest(t, webpSrc)
	image, err := govips.NewImage(opts, info)
	if err != nil {
		t.Errorf("Filename: %s, Unexpected error %v", webpSrc, err)
	}

	err = image.Resize()
	if err != nil {
		t.Errorf("Filename: %s, Unexpected error %v", webpSrc, err)
	}
	defer image.Close()

	if image.Height() != opts.Height {
		t.Errorf("Filename: %s, Unexpected file height. Expected %d. Got %d", opts.InFilename, opts.Height, image.Height())
	}
	if image.Width() != 342 {
		t.Errorf("Filename: %s, Unexpected file width. Expected %d. Got %d", opts.InFilename, 342, image.Width())
	}
}

func TestImage_ResizeForce(t *testing.T) {
	webpSrc := "./fixtures/image.webp"
	webpOut := "./test/image.webp"
	opts := &options.Options{
		InFilename:  webpSrc,
		OutFilename: webpOut,
		Width:       600,
		Height:      350,
		Force:       true,
	}

	_, info := initClientForTest(t, webpSrc)
	image, err := govips.NewImage(opts, info)
	if err != nil {
		t.Errorf("Filename: %s, Unexpected error %v", webpSrc, err)
	}

	err = image.Resize()
	if err != nil {
		t.Errorf("Filename: %s, Unexpected error %v", webpSrc, err)
	}
	defer image.Close()

	if image.Height() != opts.Height {
		t.Errorf("Filename: %s, Unexpected file height. Expected %d. Got %d", opts.InFilename, opts.Height, image.Height())
	}
	if image.Width() != opts.Width {
		t.Errorf("Filename: %s, Unexpected file width. Expected %d. Got %d", opts.InFilename, opts.Width, image.Width())
	}
}

func TestImage_ResizeNotEnlarge(t *testing.T) {
	webpSrc := "./fixtures/image.webp"
	webpOut := "./test/image.webp"
	opts := &options.Options{
		InFilename:  webpSrc,
		OutFilename: webpOut,
		Width:       1000,
		Height:      1000,
	}

	_, info := initClientForTest(t, webpSrc)
	image, err := govips.NewImage(opts, info)
	if err != nil {
		t.Errorf("Filename: %s, Unexpected error %v", webpSrc, err)
	}

	err = image.Resize()
	if err != nil {
		t.Errorf("Filename: %s, Unexpected error %v", webpSrc, err)
	}
	defer image.Close()

	if image.Height() != info.Height {
		t.Errorf("Filename: %s, Unexpected file height. Expected %d. Got %d", opts.InFilename, info.Height, image.Height())
	}
	if image.Width() != info.Width {
		t.Errorf("Filename: %s, Unexpected file width. Expected %d. Got %d", opts.InFilename, info.Width, image.Width())
	}
}

func TestImage_ResizeEnlarge(t *testing.T) {
	webpSrc := "./fixtures/image.webp"
	webpOut := "./test/image.webp"
	opts := &options.Options{
		InFilename:  webpSrc,
		OutFilename: webpOut,
		Width:       1000,
		Height:      1000,
		Enlarge:     true,
	}

	_, info := initClientForTest(t, webpSrc)
	image, err := govips.NewImage(opts, info)
	if err != nil {
		t.Errorf("Filename: %s, Unexpected error %v", webpSrc, err)
	}

	err = image.Resize()
	if err != nil {
		t.Errorf("Filename: %s, Unexpected error %v", webpSrc, err)
	}
	defer image.Close()

	if image.Height() != opts.Height {
		t.Errorf("Filename: %s, Unexpected file height. Expected %d. Got %d", opts.InFilename, opts.Height, image.Height())
	}
	if image.Width() != 977 {
		t.Errorf("Filename: %s, Unexpected file width. Expected %d. Got %d", opts.InFilename, 977, image.Width())
	}
}
