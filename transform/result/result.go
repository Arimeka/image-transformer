package result

import (
	"os"

	"github.com/Arimeka/mediaprobe"
)

type Result struct {
	Filename string
	Info     mediaprobe.Info
}

func (result Result) Clear() error {
	return os.Remove(result.Filename)
}
