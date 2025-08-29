package home

import (
	"github.com/gofred-io/gofred/column"
	"github.com/gofred-io/gofred/widget"
)

func New() widget.Widget {
	return column.New(
		[]widget.Widget{
			header(),
		},
	)
}
