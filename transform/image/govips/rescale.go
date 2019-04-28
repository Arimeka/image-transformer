package govips

import (
	"github.com/davidbyttow/govips/pkg/vips"
)

func (image *Image) Resize() error {
	image.mutex.Lock()
	defer image.mutex.Unlock()

	scaleH := image.calculateHScale()
	scaleV := image.calculateVScale()

	if image.TransformOpts.Force {
		return image.Ref.Resize(
			scaleH,
			vips.InputDouble("vscale", scaleV),
			vips.InputInt("kernel", int(vips.KernelCubic)),
		)
	}

	scale := scaleH
	if scaleV < scale {
		scale = scaleV
	}

	if !image.TransformOpts.Enlarge && scale > 1 {
		return nil
	}

	return image.Ref.Resize(
		scale,
		vips.InputInt("kernel", int(vips.KernelCubic)),
	)
}
