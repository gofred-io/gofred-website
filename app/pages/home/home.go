package home

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/router"
	"github.com/gofred-io/gofred/widget"
)

func New(params router.RouteParams) widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			container.New(
				column.New(
					[]widget.BaseWidget{
						Header(),
						hero(),
					},
					column.Flex(1),
				),
				container.Height(breakpoint.All(widget.Context().ClientHeight())),
				container.Flex(1),
			),
			features(),
			Footer(),
		},
	)
}
