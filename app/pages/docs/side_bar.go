package docs

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/foundation/column"
	. "github.com/gofred-io/gofred/foundation/container"
	. "github.com/gofred-io/gofred/foundation/link"
	. "github.com/gofred-io/gofred/foundation/spacer"
	. "github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/hooks"
	. "github.com/gofred-io/gofred/options"
	. "github.com/gofred-io/gofred/options/spacing"
	. "github.com/gofred-io/gofred/widget"
)

func docsSidebar() Widget {
	return Container(
		Column(
			[]Widget{
				docsSidebarHeader(),
				Spacer(),
				navigationMenu(),
			},
		).Gap(8).
			Flex(1),
	).Width(AllBP(280)).
		BackgroundColor("#FFFFFF").
		Padding(AllBP(All(24))).
		BorderColor("#E5E7EB").
		BorderWidth(0, 1, 0, 0).
		BorderStyle(BorderStyleTypeSolid).
		Visible(
			LG(true),
			XL(true),
			XXL(true),
		)
}

func docsSidebarHeader() Widget {
	return Column(
		[]Widget{
			Text("Documentation").
				FontSize(20).
				FontColor("#1F2937").
				FontWeight("600"),
			Text("Learn how to build with gofred").
				FontSize(14).
				FontColor("#6B7280").
				FontWeight("400"),
		},
	).Gap(4)
}

func navigationMenu() Widget {
	navigate := hooks.UseNavigate()
	activeHref := navigate.Path()

	return Column(
		[]Widget{
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
			Spacer(),
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
			Spacer(),
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
			Spacer(),
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
			Spacer(),
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
	).Gap(8)
}

type navItem struct {
	title string
	href  string
}

func navSection(title string, items []navItem, activeHref string) Widget {
	var sectionItems []Widget

	// Section title
	sectionItems = append(sectionItems, Text(title).
		FontSize(14).
		FontColor("#374151").
		FontWeight("600"))

	// Section items
	for _, item := range items {
		sectionItems = append(sectionItems, navItemWidget(item, activeHref))
	}

	return Column(
		sectionItems,
	).Gap(4)
}

func navItemWidget(item navItem, activeHref string) Widget {
	var textColor string
	var backgroundColor string

	if item.href == activeHref {
		textColor = "#2B799B"
		backgroundColor = "#EBF8FF"
	} else {
		textColor = "#6B7280"
		backgroundColor = "transparent"
	}

	return Link(
		Container(
			Text(item.title).
				FontSize(14).
				FontColor(textColor).
				FontWeight("400").
				UserSelect(UserSelectTypeNone),
		).Padding(AllBP(Axis(8, 12))).
			BackgroundColor(backgroundColor).
			BorderRadius(6),
	).Href(item.href)
}
