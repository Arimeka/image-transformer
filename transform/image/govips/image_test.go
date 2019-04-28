package govips_test

import (
	"testing"

	"github.com/Arimeka/image-transformer/transform/image/govips"
	"github.com/Arimeka/image-transformer/transform/options"
)

func testNewImageWEBPShrink(t *testing.T, shrink int) {
	src := "./fixtures/image.webp"
	out := "/dev/null"

	_, info := initClientForTest(t, src)

	toH := info.Height / shrink
	toW := info.Width / shrink

	if toH > toW {
		shrink = info.Height / toW
	} else {
		shrink = info.Height / toH
	}

	opts := &options.Options{
		InFilename:  src,
		OutFilename: out,
		Width:       toW,
		Height:      toH,
		Enlarge:     true,
	}

	image, err := govips.NewImage(opts, info)
	if err != nil {
		t.Errorf("Filename: %s, Unexpected error %v", src, err)
	}
	defer image.Close()

	if image.Height() != info.Height/shrink {
		t.Errorf("Filename: %s, Unexpected file height. Expected %d. Got %d", opts.InFilename, info.Height/shrink, image.Height())
	}
	if image.Width() != info.Width/shrink {
		t.Errorf("Filename: %s, Unexpected file width. Expected %d. Got %d", opts.InFilename, info.Width/shrink, image.Width())
	}
}

func testNewImagePNGShrink(t *testing.T, shrink int) {
	src := "./fixtures/image.png"
	out := "/dev/null"

	_, info := initClientForTest(t, src)

	opts := &options.Options{
		InFilename:  src,
		OutFilename: out,
		Width:       info.Width / shrink,
		Height:      info.Height / shrink,
		Enlarge:     true,
	}

	image, err := govips.NewImage(opts, info)
	if err != nil {
		t.Errorf("Filename: %s, Unexpected error %v", src, err)
	}
	defer image.Close()

	if image.Height() != info.Height {
		t.Errorf("Filename: %s, Unexpected file height. Expected %d. Got %d", opts.InFilename, info.Height, image.Height())
	}
	if image.Width() != info.Width {
		t.Errorf("Filename: %s, Unexpected file width. Expected %d. Got %d", opts.InFilename, info.Width, image.Width())
	}
}

func testNewImageJPEGShrink(t *testing.T, shrink int) {
	src := "./fixtures/image.jpg"
	out := "/dev/null"

	_, info := initClientForTest(t, src)

	opts := &options.Options{
		InFilename:  src,
		OutFilename: out,
		Width:       info.Width / shrink,
		Height:      info.Height / shrink,
		Enlarge:     true,
	}

	switch {
	case shrink > 8:
		shrink = 8
	case shrink > 4:
		shrink = 4
	case shrink > 2:
		shrink = 2
	}

	image, err := govips.NewImage(opts, info)
	if err != nil {
		t.Errorf("Filename: %s, Unexpected error %v", src, err)
	}
	defer image.Close()

	if image.Height() != info.Height/shrink {
		t.Errorf("Filename: %s, Unexpected file height. Expected %d. Got %d", opts.InFilename, info.Height/shrink, image.Height())
	}
	if image.Width() != info.Width/shrink {
		t.Errorf("Filename: %s, Unexpected file width. Expected %d. Got %d", opts.InFilename, info.Width/shrink, image.Width())
	}
}

func testNewImageShrink(t *testing.T, shrink int) {
	testNewImageJPEGShrink(t, shrink)
	testNewImageWEBPShrink(t, shrink)
	testNewImagePNGShrink(t, shrink)
}

func TestNewImage(t *testing.T) {
	testNewImageShrink(t, 20)
	testNewImageShrink(t, 10)
	testNewImageShrink(t, 7)
	testNewImageShrink(t, 5)
	testNewImageShrink(t, 4)
	testNewImageShrink(t, 3)
	testNewImageShrink(t, 1)
}

func TestNewImageLQ(t *testing.T) {
	src := "./fixtures/image.png"
	out := "/dev/null"
	shrink := 20

	_, info := initClientForTest(t, src)

	opts := &options.Options{
		InFilename:  src,
		OutFilename: out,
		Width:       info.Width / shrink,
		Height:      info.Height / shrink,
		Enlarge:     true,
		LowQuality:  true,
	}

	image, err := govips.NewImage(opts, info)
	if err != nil {
		t.Errorf("Filename: %s, Unexpected error %v", src, err)
	}
	defer image.Close()

	if image.Height() != info.Height/shrink {
		t.Errorf("Filename: %s, Unexpected file height. Expected %d. Got %d", opts.InFilename, info.Height/shrink, image.Height())
	}
	if image.Width() != info.Width/shrink {
		t.Errorf("Filename: %s, Unexpected file width. Expected %d. Got %d", opts.InFilename, info.Width/shrink, image.Width())
	}
}
