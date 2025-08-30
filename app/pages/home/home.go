package home

import (
	"github.com/gofred-io/gofred/column"
	"github.com/gofred-io/gofred/container"
	"github.com/gofred-io/gofred/widget"
)

func New() widget.Widget {
	return column.New(
		[]widget.Widget{
			container.New(
				column.New(
					[]widget.Widget{
						header(),
						hero(),
					},
				),
				container.Style(
					container.Height(widget.Context().ClientHeight()),
				),
			),
			features(),
			footer(),
		},
	)
}
