package home

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/drawer"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	iconbutton "github.com/gofred-io/gofred/foundation/icon_button"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/widget"
)

var leftDrawer *drawer.Drawer

func init() {
	leftDrawer = buildLeftDrawer()
}

func buildLeftDrawer() *drawer.Drawer {
	return drawer.New(
		container.New(
			column.New(
				[]widget.BaseWidget{
					container.New(
						row.New(
							[]widget.BaseWidget{
								logoTitle(),
								spacer.New(),
								iconbutton.New(
									icondata.Close,
									iconbutton.Fill("#003B73"),
									iconbutton.OnClick(func(this widget.BaseWidget, e widget.Event) {
										leftDrawer.Hide()
									}),
								),
							},
							row.Gap(8),
							row.Flex(1),
							row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
						),
						container.Padding(breakpoint.All(spacing.LRTB(16, 16, 16, 8))),
						container.BorderColor("#CECECE"),
						container.BorderWidth(0, 0, 1, 0),
						container.BorderStyle(options.BorderStyleTypeSolid),
					),
					container.New(
						column.New(
							[]widget.BaseWidget{
								documentationLink(true),
								discussionsLink(true),
								githubLink(true),
							},
							column.Gap(16),
							column.Flex(1),
						),
						container.Padding(breakpoint.All(spacing.All(16))),
					),
				},
				column.Gap(8),
				column.Flex(1),
			),
		),
		drawer.Width(breakpoint.All(250)),
		drawer.Transition(0.3),
	)
}
