package builder

import "errors"

var (
	ErrFailRenderTemplate   = errors.New("failed to render the template")
	ErrFailReadChartConfig  = errors.New("failed to read chart config")
	ErrFailProcessChartFile = errors.New("failed to process charts file")
)
