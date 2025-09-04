package drawer

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/drawer"
	"github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	iconbutton "github.com/gofred-io/gofred/foundation/icon_button"
	"github.com/gofred-io/gofred/foundation/link"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/widget"
)

var leftDrawer *drawer.Drawer

func init() {
	leftDrawer = buildLeftDrawer()
}

func Get() *drawer.Drawer {
	return leftDrawer
}

func buildLeftDrawer() *drawer.Drawer {
	return drawer.New(
		container.New(
			column.New(
				[]widget.BaseWidget{
					container.New(
						row.New(
							[]widget.BaseWidget{
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

func documentationLink(inMenu bool) widget.BaseWidget {
	var (
		visible = []breakpoint.BreakpointOptions[bool]{breakpoint.All(true)}
	)

	if !inMenu {
		visible = append(visible, breakpoint.XS(false), breakpoint.SM(false))
	}

	return container.New(
		link.New(
			text.New(
				"Docs",
				text.FontSize(16),
				text.FontColor("#2B799B"),
				text.FontWeight("500"),
				text.UserSelect(options.UserSelectTypeNone),
			),
			link.Href("/docs"),
			link.OnClick(func(this widget.BaseWidget, e widget.Event) {
				leftDrawer.Hide()
			}),
		),
		container.Visible(visible...),
	)
}

func discussionsLink(inMenu bool) widget.BaseWidget {
	var (
		visible = []breakpoint.BreakpointOptions[bool]{breakpoint.All(true)}
	)

	if !inMenu {
		visible = append(visible, breakpoint.XS(false), breakpoint.SM(false))
	}

	return container.New(
		link.New(
			row.New(
				[]widget.BaseWidget{
					text.New(
						"Discussions",
						text.FontSize(16),
						text.FontColor("#2B799B"),
						text.FontWeight("500"),
						text.UserSelect(options.UserSelectTypeNone),
					),
					icon.New(
						icondata.OpenInNew,
						icon.Width(breakpoint.All(18)),
						icon.Height(breakpoint.All(18)),
						icon.Fill("#2B799B"),
					),
				},
			),
			link.Href("https://github.com/gofred-io/gofred/pulls"),
			link.NewTab(true),
		),
		container.Visible(visible...),
	)
}

func githubLink(inMenu bool) widget.BaseWidget {
	var (
		visible = []breakpoint.BreakpointOptions[bool]{breakpoint.All(true)}
	)

	if !inMenu {
		visible = append(visible, breakpoint.XS(false), breakpoint.SM(false))
	}

	return container.New(
		link.New(
			row.New(
				[]widget.BaseWidget{
					text.New(
						"GitHub",
						text.FontSize(16),
						text.FontColor("#2B799B"),
						text.FontWeight("500"),
						text.UserSelect(options.UserSelectTypeNone),
					),
					icon.New(
						icondata.OpenInNew,
						icon.Width(breakpoint.All(18)),
						icon.Height(breakpoint.All(18)),
						icon.Fill("#2B799B"),
					),
				},
			),
			link.Href("https://github.com/gofred-io/gofred"),
			link.NewTab(true),
		),
		container.Visible(visible...),
	)
}
