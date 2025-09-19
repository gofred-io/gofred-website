package drawer

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/drawer"
	"github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	iconbutton "github.com/gofred-io/gofred/foundation/icon_button"
	"github.com/gofred-io/gofred/foundation/link"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/scaffold"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
)

const (
	Name = "docs"
)

func New() (string, *drawer.Drawer) {
	return Name, drawer.New(
		func() application.BaseWidget {
			return container.New(
				column.New(
					[]application.BaseWidget{
						drawerHeader(),
						drawerNavigation(),
					},
				),
			)
		},
		drawer.ID("docs-left-drawer"),
		drawer.Width(breakpoint.All(320)),
		drawer.Transition(0.3),
	)
}

func drawerHeader() application.BaseWidget {
	return container.New(
		row.New(
			[]application.BaseWidget{
				drawerTitle(),
				spacer.New(),
				iconbutton.New(
					icondata.Close,
					iconbutton.Fill("#6B7280"),
					iconbutton.OnClick(func(this application.BaseWidget, e application.Event) {
						scaffold.Get().Drawer(Name).Hide()
					}),
				),
			},
			row.Gap(8),
			row.Flex(1),
			row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
		),
		container.Padding(breakpoint.All(spacing.LRTB(20, 20, 16, 8))),
		container.BorderColor("#E5E7EB"),
		container.BorderWidth(spacing.Bottom(1)),
		container.BorderStyle(theme.BorderStyleTypeSolid),
	)
}

func drawerTitle() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			text.New(
				"Documentation",
				text.FontSize(18),
				text.FontWeight("600"),
				text.UserSelect(theme.UserSelectTypeNone),
			),
			text.New(
				"Learn how to build with gofred",
				text.FontSize(12),
				text.FontColor("#6B7280"),
				text.UserSelect(theme.UserSelectTypeNone),
			),
		},
		column.Gap(2),
	)
}

func drawerNavigation() application.BaseWidget {
	navigate := hooks.UseNavigate()

	return container.New(
		listenable.Builder(navigate, func() application.BaseWidget {
			activeHref := navigate.Path()

			return column.New(
				[]application.BaseWidget{
					drawerNavSection("Getting Started", []drawerNavItem{
						{title: "Installation", href: "/docs/installation", icon: icondata.Download},
						{title: "Quick Start", href: "/docs/quick-start", icon: icondata.RocketLaunchOutline},
						{title: "Your First App", href: "/docs/first-app", icon: icondata.Application},
						{title: "Project Structure", href: "/docs/project-structure", icon: icondata.Folder},
					}, activeHref),
					spacer.New(spacer.Height(16)),
					drawerNavSection("Core Concepts", []drawerNavItem{
						{title: "Widgets", href: "/docs/widgets", icon: icondata.Widgets},
						{title: "Layouts", href: "/docs/layouts", icon: icondata.Grid},
						{title: "Styling", href: "/docs/styling", icon: icondata.PaletteOutline},
						{title: "State Management", href: "/docs/state", icon: icondata.Database},
						{title: "Event Handling", href: "/docs/events", icon: icondata.Mouse},
					}, activeHref),
					spacer.New(spacer.Height(16)),
					drawerNavSection("Components", []drawerNavItem{
						{title: "Buttons", href: "/docs/components/buttons", icon: icondata.ButtonPointer},
						{title: "Navigation", href: "/docs/components/navigation", icon: icondata.Menu},
						{title: "Icons", href: "/docs/components/icons", icon: icondata.Star},
						{title: "Images", href: "/docs/components/images", icon: icondata.Image},
						{title: "Containers", href: "/docs/components/containers", icon: icondata.Package},
					}, activeHref),
					spacer.New(spacer.Height(16)),
					drawerNavSection("Advanced", []drawerNavItem{
						{title: "Routing", href: "/docs/routing", icon: icondata.SignDirection},
						{title: "API Reference", href: "/docs/api", icon: icondata.FileDocument},
						{title: "Best Practices", href: "/docs/best-practices", icon: icondata.ThumbUp},
						{title: "Performance", href: "/docs/performance", icon: icondata.Speedometer},
						{title: "Deployment", href: "/docs/deployment", icon: icondata.Server},
					}, activeHref),
					spacer.New(spacer.Height(16)),
					drawerNavSection("Resources", []drawerNavItem{
						{title: "Examples", href: "/docs/examples", icon: icondata.Lightbulb},
						{title: "Tutorials", href: "/docs/tutorials", icon: icondata.Book},
						{title: "Community", href: "/docs/community", icon: icondata.AccountGroup},
						{title: "Support", href: "/docs/support", icon: icondata.Help},
					}, activeHref),
				},
				column.Gap(8),
				column.Flex(1),
			)
		},
		),
		container.Padding(breakpoint.All(spacing.LRTB(20, 20, 20, 20))),
	)
}

type drawerNavItem struct {
	title string
	href  string
	icon  icondata.IconData
}

func drawerNavSection(title string, items []drawerNavItem, activeHref string) application.BaseWidget {
	var sectionItems []application.BaseWidget

	// Section title
	sectionItems = append(sectionItems, text.New(
		title,
		text.FontSize(12),
		text.FontColor("#6B7280"),
		text.FontWeight("600"),
		text.UserSelect(theme.UserSelectTypeNone),
	))

	// Section items
	for _, item := range items {
		sectionItems = append(sectionItems, drawerNavItemWidget(item, activeHref))
	}

	return column.New(
		sectionItems,
		column.Gap(4),
	)
}

func drawerNavItemWidget(item drawerNavItem, activeHref string) application.BaseWidget {
	var textColor string
	var backgroundColor string
	var iconColor string

	if item.href == activeHref {
		textColor = "#2B799B"
		backgroundColor = "#EBF8FF"
		iconColor = "#2B799B"
	} else {
		textColor = "#374151"
		backgroundColor = "transparent"
		iconColor = "#9CA3AF"
	}

	return link.New(
		container.New(
			row.New(
				[]application.BaseWidget{
					icon.New(
						item.icon,
						icon.Width(breakpoint.All(20)),
						icon.Height(breakpoint.All(20)),
						icon.Fill(iconColor),
					),
					spacer.New(spacer.Width(6)),
					text.New(
						item.title,
						text.FontSize(14),
						text.FontColor(textColor),
						text.UserSelect(theme.UserSelectTypeNone),
					),
				},
				row.Gap(8),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),
			container.Padding(breakpoint.All(spacing.Axis(8, 12))),
			container.BackgroundColor(backgroundColor),
			container.BorderRadius(6),
		),
		link.Href(item.href),
		link.OnClick(func(this application.BaseWidget, e application.Event) {
			scaffold.Get().Drawer(Name).Hide()
		}),
	)
}
