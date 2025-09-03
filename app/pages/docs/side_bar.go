package docs

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/link"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/widget"
)

func docsSidebar(activeHref string) widget.BaseWidget {
	return container.New(
		column.New(
			[]widget.BaseWidget{
				docsSidebarHeader(),
				spacer.New(spacer.Height(16)),
				navigationMenu(activeHref),
			},
			column.Gap(8),
			column.Flex(1),
		),
		container.Width(breakpoint.All(280)),
		container.BackgroundColor("#FFFFFF"),
		container.Padding(breakpoint.All(spacing.All(24))),
		container.BorderColor("#E5E7EB"),
		container.BorderWidth(0, 1, 0, 0),
		container.BorderStyle(options.BorderStyleTypeSolid),
		container.Visible(
			breakpoint.LG(true),
			breakpoint.XL(true),
			breakpoint.XXL(true),
		),
	)
}

func docsSidebarHeader() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			text.New(
				"Documentation",
				text.FontSize(20),
				text.FontColor("#1F2937"),
				text.FontWeight("600"),
			),
			text.New(
				"Learn how to build with gofred",
				text.FontSize(14),
				text.FontColor("#6B7280"),
				text.FontWeight("400"),
			),
		},
		column.Gap(4),
	)
}

func navigationMenu(activeHref string) widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			navSection(
				"Getting Started",
				[]navItem{
					{title: "Installation", href: "/docs/installation"},
					{title: "Quick Start", href: "/docs/quick-start"},
					{title: "Your First App", href: "/docs/first-app"},
					{title: "Project Structure", href: "/docs/project-structure"},
				},
				activeHref,
			),
			spacer.New(spacer.Height(16)),
			navSection(
				"Core Concepts",
				[]navItem{
					{title: "Widgets", href: "/docs/widgets"},
					{title: "Layouts", href: "/docs/layouts"},
					{title: "Styling", href: "/docs/styling"},
					{title: "State Management", href: "/docs/state"},
					{title: "Event Handling", href: "/docs/events"},
				},
				activeHref,
			),
			spacer.New(spacer.Height(16)),
			navSection(
				"Components",
				[]navItem{
					{title: "Buttons", href: "/docs/components/buttons"},
					{title: "Navigation", href: "/docs/components/navigation"},
					{title: "Icons", href: "/docs/components/icons"},
					{title: "Images", href: "/docs/components/images"},
					{title: "Containers", href: "/docs/components/containers"},
				},
				activeHref,
			),
			spacer.New(spacer.Height(16)),
			navSection(
				"Advanced",
				[]navItem{
					{title: "Routing", href: "/docs/routing"},
					{title: "API Reference", href: "/docs/api"},
					{title: "Best Practices", href: "/docs/best-practices"},
					{title: "Performance", href: "/docs/performance"},
					{title: "Deployment", href: "/docs/deployment"},
				},
				activeHref,
			),
			spacer.New(spacer.Height(16)),
			navSection(
				"Resources",
				[]navItem{
					{title: "Examples", href: "/docs/examples"},
					{title: "Tutorials", href: "/docs/tutorials"},
					{title: "Community", href: "/docs/community"},
					{title: "Support", href: "/docs/support"},
				},
				activeHref,
			),
		},
		column.Gap(8),
	)
}

type navItem struct {
	title string
	href  string
}

func navSection(title string, items []navItem, activeHref string) widget.BaseWidget {
	var sectionItems []widget.BaseWidget

	// Section title
	sectionItems = append(sectionItems, text.New(
		title,
		text.FontSize(14),
		text.FontColor("#374151"),
		text.FontWeight("600"),
	))

	// Section items
	for _, item := range items {
		sectionItems = append(sectionItems, navItemWidget(item, activeHref))
	}

	return column.New(
		sectionItems,
		column.Gap(4),
	)
}

func navItemWidget(item navItem, activeHref string) widget.BaseWidget {
	var textColor string
	var backgroundColor string

	if item.href == activeHref {
		textColor = "#2B799B"
		backgroundColor = "#EBF8FF"
	} else {
		textColor = "#6B7280"
		backgroundColor = "transparent"
	}

	return link.New(
		container.New(
			text.New(
				item.title,
				text.FontSize(14),
				text.FontColor(textColor),
				text.FontWeight("400"),
				text.UserSelect(options.UserSelectTypeNone),
			),
			container.Padding(breakpoint.All(spacing.Axis(8, 12))),
			container.BackgroundColor(backgroundColor),
			container.BorderRadius(6),
		),
		link.Href(item.href),
	)
}
