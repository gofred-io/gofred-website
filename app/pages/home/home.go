package home

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/column"
	"github.com/gofred-io/gofred/container"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/widget"
)

func New() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			container.New(
				column.New(
					[]widget.BaseWidget{
						header(),
						hero(),
					},
					column.Flex(1),
				),
				options.Height(breakpoint.All(widget.Context().ClientHeight())),
				options.Flex(1),
			),
			features(),
			footer(),
		},
	)
}
