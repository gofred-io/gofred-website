package codeblock

import (
	gfcode "github.com/gofred-io/gofred/basic/code"
	"github.com/gofred-io/gofred/basic/pre"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/container"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	iconbutton "github.com/gofred-io/gofred/foundation/icon_button"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/utils"
	"github.com/gofred-io/gofred/widget"
)

func New(code string) widget.BaseWidget {
	return container.New(
		row.New(
			[]widget.BaseWidget{
				pre.New(
					gfcode.New(
						code,
						gfcode.FontSize(14),
						gfcode.FontColor("#F3F4F6"),
						gfcode.FontWeight("400"),
					),
				),
				spacer.New(),
				iconbutton.New(
					icondata.ContentCopy,
					iconbutton.Fill("#F3F4F6"),
					iconbutton.OnClick(func(this widget.BaseWidget, e widget.Event) {
						utils.CopyToClipboard(code)
					}),
				),
			},
			row.Flex(1),
			row.CrossAxisAlignment(options.AxisAlignmentTypeStart),
		),
		container.BackgroundColor("#1F2937"),
		container.Padding(breakpoint.All(spacing.All(16))),
		container.BorderRadius(8),
		container.MaxWidth(720),
	)
}
