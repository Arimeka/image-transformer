package govips

import (
	"github.com/Arimeka/image-transformer/transform/options"
	"github.com/davidbyttow/govips/pkg/vips"
)

func (image *Image) Crop() error {
	image.mutex.Lock()
	defer image.mutex.Unlock()

	scaleH := image.calculateHScale()
	scaleV := image.calculateVScale()

	scale := scaleH
	if scaleV > scale {
		scale = scaleV
	}

	if scaleH == 1 && scaleV == 1 {
		return nil
	}

	if err := image.Ref.Resize(
		scale,
		vips.InputInt("kernel", int(vips.KernelCubic)),
	); err != nil {
		return err
	}

	if image.TransformOpts.Width == 0 || image.TransformOpts.Height == 0 {
		return nil
	}

	if image.TransformOpts.Gravity == options.GravitySmart {
		return image.Ref.Smartcrop(
			image.TransformOpts.Width,
			image.TransformOpts.Height,
		)
	}

	leftEdge, topEdge := image.cropEdges()

	return image.Ref.ExtractArea(
		leftEdge, topEdge,
		image.TransformOpts.Width, image.TransformOpts.Height,
	)
}

func (image *Image) cropEdges() (leftEdge, topEdge int) {
	switch image.TransformOpts.Gravity {
	case options.GravityEast:
		leftEdge = image.Ref.Width() - image.TransformOpts.Width
		topEdge = (image.Ref.Height() - image.TransformOpts.Height) / 2
	case options.GravitySouth:
		topEdge = image.Ref.Height() - image.TransformOpts.Height
		leftEdge = (image.Ref.Width() - image.TransformOpts.Width) / 2
	case options.GravityCenter:
		topEdge = (image.Ref.Height() - image.TransformOpts.Height) / 2
		leftEdge = (image.Ref.Width() - image.TransformOpts.Width) / 2
	case options.GravityNorth:
		topEdge = 0
		leftEdge = (image.Ref.Width() - image.TransformOpts.Width) / 2
	case options.GravityWest:
		leftEdge = 0
		topEdge = (image.Ref.Height() - image.TransformOpts.Height) / 2
	}

	return
}
