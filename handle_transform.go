package transformer

import (
	"net/http"

	"github.com/Arimeka/image-transformer/transform"
	"github.com/Arimeka/image-transformer/transform/options"
	"github.com/valyala/fasthttp"
)

func TransformHandle(ctx *fasthttp.RequestCtx) {
	if ctx.IsOptions() {
		ctx.SetStatusCode(http.StatusNoContent)

		return
	}

	fHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)

		return
	}

	transformOpts, err := options.NewOptions(fHeader)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)

		return
	}
	defer transformOpts.Free()

	err = transformOpts.Parse(ctx)
	if err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)

		return
	}

	processor, err := transform.New(transformOpts)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)

		return
	}

	result, err := processor.Process()
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)

		return
	}

	defer result.Clear()

	ctx.SendFile(result.Filename)
}
