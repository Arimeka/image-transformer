package govips_test

import (
	"testing"

	"github.com/Arimeka/image-transformer/transform/image/govips"
	"github.com/Arimeka/image-transformer/transform/options"
)

func testImage_CropGravity(t *testing.T, gravity options.CropGravity) {
	pngSrc := "./fixtures/image.png"
	pngOut := "./test/image.png"
	opts := &options.Options{
		InFilename:  pngSrc,
		OutFilename: pngOut,
		Width:       600,
		Height:      350,
		Enlarge:     true,
		Crop:        true,
		Gravity:     gravity,
	}

	_, info := initClientForTest(t, pngSrc)
	image, err := govips.NewImage(opts, info)
	if err != nil {
		t.Errorf("Filename: %s, Unexpected error %v", pngSrc, err)
	}

	err = image.Crop()
	if err != nil {
		t.Errorf("Filename: %s, Unexpected error %v", pngSrc, err)
	}
	defer image.Close()

	if image.Height() != opts.Height {
		t.Errorf("Filename: %s, Unexpected file height. Expected %d. Got %d", opts.InFilename, opts.Height, image.Height())
	}
	if image.Width() != opts.Width {
		t.Errorf("Filename: %s, Unexpected file width. Expected %d. Got %d", opts.InFilename, opts.Width, image.Width())
	}
}

func TestImage_Crop(t *testing.T) {
	testImage_CropGravity(t, options.GravitySmart)
	testImage_CropGravity(t, options.GravityEast)
	testImage_CropGravity(t, options.GravitySouth)
	testImage_CropGravity(t, options.GravityWest)
	testImage_CropGravity(t, options.GravityNorth)
	testImage_CropGravity(t, options.GravityCenter)
}

func TestImage_CropOneSide(t *testing.T) {
	webpSrc := "./fixtures/image.webp"
	webpOut := "./test/image.webp"
	opts := &options.Options{
		InFilename:  webpSrc,
		OutFilename: webpOut,
		Height:      1000,
		Crop:        true,
	}

	_, info := initClientForTest(t, webpSrc)
	image, err := govips.NewImage(opts, info)
	if err != nil {
		t.Errorf("Filename: %s, Unexpected error %v", webpSrc, err)
	}

	err = image.Crop()
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
func TestImage_CropSameSize(t *testing.T) {
	webpSrc := "./fixtures/image.webp"
	webpOut := "./test/image.webp"

	_, info := initClientForTest(t, webpSrc)

	opts := &options.Options{
		InFilename:  webpSrc,
		OutFilename: webpOut,
		Height:      info.Height,
		Width:       info.Width,
		Crop:        true,
	}

	image, err := govips.NewImage(opts, info)
	if err != nil {
		t.Errorf("Filename: %s, Unexpected error %v", webpSrc, err)
	}

	err = image.Crop()
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
