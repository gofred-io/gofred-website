package docs

import (
	appTheme "github.com/gofred-io/gofred-website/app/theme"

	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/link"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
	"github.com/gofred-io/gofred/theme/theme_style"
)

func docsSidebar() application.BaseWidget {
	return container.New(
		column.New(
			[]application.BaseWidget{
				docsSidebarHeader(),
				spacer.New(spacer.Height(16)),
				navigationMenu(),
			},
			column.Gap(8),
			column.Flex(1),
		),
		container.Width(breakpoint.All(280)),
		container.Padding(breakpoint.All(spacing.All(24))),
		container.BorderWidth(spacing.Right(1)),
		container.BorderStyle(theme.BorderStyleTypeSolid),
		container.Visible(
			breakpoint.LG(true),
			breakpoint.XL(true),
			breakpoint.XXL(true),
		),
	)
}

func docsSidebarHeader() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			text.New(
				"Documentation",
				text.FontSize(20),
				text.FontWeight("700"),
			),
			text.New(
				"Learn how to build with gofred",
				text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
				text.FontSize(14),
			),
		},
		column.Gap(4),
	)
}

func navigationMenu() application.BaseWidget {
	navigate := hooks.UseNavigate()

	return listenable.Builder(navigate, func() application.BaseWidget {
		activeHref := navigate.Path()

		return column.New(
			[]application.BaseWidget{
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
						{title: "Buttons", href: "/docs/buttons"},
						{title: "Navigation", href: "/docs/navigation"},
						{title: "Icons", href: "/docs/icons"},
						{title: "Images", href: "/docs/images"},
						{title: "Containers", href: "/docs/containers"},
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
	})
}

type navItem struct {
	title string
	href  string
}

func navSection(title string, items []navItem, activeHref string) application.BaseWidget {
	var sectionItems []application.BaseWidget

	// Section title
	sectionItems = append(sectionItems, text.New(
		title,
		text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
		text.FontSize(14),
		text.FontWeight("700"),
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

func navItemWidget(item navItem, activeHref string) application.BaseWidget {
	var containerStyle theme_style.ContainerStyle
	var textStyle theme_style.TextStyle

	if item.href == activeHref {
		containerStyle = appTheme.Data().BoxTheme.ContainerStyle.Tertiary
		textStyle = appTheme.Data().TextTheme.TextStyle.Tertiary
	} else {
		containerStyle = appTheme.Data().BoxTheme.ContainerStyle.Primary
		textStyle = appTheme.Data().TextTheme.TextStyle.Primary
	}

	return link.New(
		container.New(
			text.New(
				item.title,
				text.TextStyle(textStyle),
				text.FontSize(14),
				text.UserSelect(theme.UserSelectTypeNone),
			),
			container.ContainerStyle(containerStyle),
			container.Padding(breakpoint.All(spacing.Axis(8, 12))),
			container.BorderRadius(6),
		),
		link.Href(item.href),
	)
}
