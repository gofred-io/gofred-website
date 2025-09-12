package drawer

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/center"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/drawer"
	"github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	iconbutton "github.com/gofred-io/gofred/foundation/icon_button"
	"github.com/gofred-io/gofred/foundation/image"
	"github.com/gofred-io/gofred/foundation/link"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
)

var leftDrawer drawer.IDrawer

func Get() drawer.IDrawer {
	if leftDrawer == nil {
		leftDrawer = buildLeftDrawer()
	}
	return leftDrawer
}

func buildLeftDrawer() drawer.IDrawer {
	return drawer.New(
		container.New(
			column.New(
				[]application.BaseWidget{
					drawerHeader(),
					drawerContent(),
				},
				column.Gap(0),
				column.Flex(1),
			),
			container.Flex(1),
		),
		drawer.ID("root-left-drawer"),
		drawer.Width(breakpoint.All(320)),
		drawer.Transition(0.3),
	)
}

func drawerHeader() application.BaseWidget {
	return container.New(
		row.New(
			[]application.BaseWidget{
				drawerLogo(),
				spacer.New(),
				iconbutton.New(
					icondata.Close,
					iconbutton.Fill("#6B7280"),
					iconbutton.OnClick(func(this application.BaseWidget, e application.Event) {
						leftDrawer.Hide()
					}),
				),
			},
			row.Gap(12),
			row.Flex(1),
			row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
		),
		container.Padding(breakpoint.All(spacing.LRTB(24, 12, 18, 14))),
		container.BorderColor("#E5E7EB"),
		container.BorderWidth(0, 0, 1, 0),
		container.BorderStyle(theme.BorderStyleTypeSolid),
	)
}

func drawerLogo() application.BaseWidget {
	return row.New(
		[]application.BaseWidget{
			image.New(
				"img/gofred.png",
				image.Width(breakpoint.All(32)),
				image.Height(breakpoint.All(32)),
			),
			text.New(
				"gofred",
				text.FontSize(20),
				text.FontWeight("700"),
				text.UserSelect(theme.UserSelectTypeNone),
			),
		},
		row.Gap(16),
		row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
	)
}

func drawerContent() application.BaseWidget {
	return container.New(
		column.New(
			[]application.BaseWidget{
				navigationSection(),
				externalLinksSection(),
				spacer.New(),
				drawerFooter(),
			},
			column.Gap(24),
			column.Flex(1),
		),
		container.Flex(1),
		container.Padding(breakpoint.All(spacing.LRTB(24, 24, 24, 24))),
	)
}

func navigationSection() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			sectionTitle("Navigation"),
			spacer.New(spacer.Height(12)),
			navItem("Home", "/", icondata.Home, false),
			navItem("Documentation", "/docs", icondata.FileDocument, false),
			navItem("Getting Started", "/docs/getting-started", icondata.Play, false),
			navItem("Core Concepts", "/docs/core-concepts", icondata.Lightbulb, false),
			navItem("Components", "/docs/components", icondata.Package, false),
			navItem("API Reference", "/docs/api", icondata.FileDocument, false),
		},
		column.Gap(4),
	)
}

func externalLinksSection() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			sectionTitle("Resources"),
			spacer.New(spacer.Height(12)),
			externalNavItem("GitHub", "https://github.com/gofred-io/gofred", icondata.Github),
			externalNavItem("Discussions", "https://github.com/gofred-io/gofred/discussions", icondata.Comment),
			externalNavItem("Examples", "https://github.com/gofred-io/examples", icondata.Lightbulb),
			externalNavItem("Community", "https://github.com/orgs/gofred-io/discussions", icondata.AccountGroup),
		},
		column.Gap(4),
	)
}

func drawerFooter() application.BaseWidget {
	return container.New(
		center.New(
			column.New(
				[]application.BaseWidget{
					text.New(
						"Built with gofred",
						text.FontSize(12),
						text.FontColor("#9CA3AF"),
					),
					spacer.New(spacer.Height(4)),
					text.New(
						"v1.0.0",
						text.FontSize(11),
						text.FontColor("#D1D5DB"),
					),
				},
				column.Gap(0),
				column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),
		),
		container.Padding(breakpoint.All(spacing.LRTB(0, 0, 0, 0))),
	)
}

func sectionTitle(title string) application.BaseWidget {
	return text.New(
		title,
		text.FontSize(12),
		text.FontColor("#6B7280"),
		text.FontWeight("600"),
		text.UserSelect(theme.UserSelectTypeNone),
	)
}

func navItem(title, href string, iconData icondata.IconData, isExternal bool) application.BaseWidget {
	return container.New(
		link.New(
			row.New(
				[]application.BaseWidget{
					icon.New(
						iconData,
						icon.Width(breakpoint.All(20)),
						icon.Height(breakpoint.All(20)),
						icon.Fill("#6B7280"),
					),
					text.New(
						title,
						text.FontSize(14),
						text.FontColor("#374151"),
						text.FontWeight("500"),
						text.UserSelect(theme.UserSelectTypeNone),
					),
					spacer.New(),
					func() application.BaseWidget {
						if isExternal {
							return icon.New(
								icondata.OpenInNew,
								icon.Width(breakpoint.All(16)),
								icon.Height(breakpoint.All(16)),
								icon.Fill("#9CA3AF"),
							)
						}
						return spacer.New()
					}(),
				},
				row.Gap(12),
				row.Flex(1),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),
			link.Href(href),
			link.OnClick(func(this application.BaseWidget, e application.Event) {
				Get().Hide()
			}),
		),
		container.Padding(breakpoint.All(spacing.LRTB(12, 12, 12, 12))),
		container.BorderRadius(8),
		container.BackgroundColor("transparent"),
		container.BorderWidth(1, 1, 1, 1),
		container.BorderColor("transparent"),
		container.BorderStyle(theme.BorderStyleTypeSolid),
	)
}

func externalNavItem(title, href string, iconData icondata.IconData) application.BaseWidget {
	return container.New(
		link.New(
			row.New(
				[]application.BaseWidget{
					icon.New(
						iconData,
						icon.Width(breakpoint.All(20)),
						icon.Height(breakpoint.All(20)),
						icon.Fill("#6B7280"),
					),
					text.New(
						title,
						text.FontSize(14),
						text.FontColor("#374151"),
						text.FontWeight("500"),
						text.UserSelect(theme.UserSelectTypeNone),
					),
					spacer.New(),
					icon.New(
						icondata.OpenInNew,
						icon.Width(breakpoint.All(16)),
						icon.Height(breakpoint.All(16)),
						icon.Fill("#9CA3AF"),
					),
				},
				row.Gap(12),
				row.Flex(1),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),
			link.Flex(1),
			link.Href(href),
			link.NewTab(true),
		),
		container.Padding(breakpoint.All(spacing.LRTB(12, 12, 12, 12))),
		container.BorderRadius(8),
		container.BackgroundColor("transparent"),
		container.BorderWidth(1, 1, 1, 1),
		container.BorderColor("transparent"),
		container.BorderStyle(theme.BorderStyleTypeSolid),
	)
}
