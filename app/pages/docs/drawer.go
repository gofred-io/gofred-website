package docs

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/foundation/button"
	. "github.com/gofred-io/gofred/foundation/column"
	. "github.com/gofred-io/gofred/foundation/container"
	. "github.com/gofred-io/gofred/foundation/drawer"
	. "github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	. "github.com/gofred-io/gofred/foundation/link"
	. "github.com/gofred-io/gofred/foundation/row"
	. "github.com/gofred-io/gofred/foundation/spacer"
	. "github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/options"
	. "github.com/gofred-io/gofred/options/spacing"
	. "github.com/gofred-io/gofred/widget"
)

var (
	docsDrawer *FDrawer
)

func DocsDrawer() *FDrawer {
	if docsDrawer == nil {
		docsDrawer = buildDocsDrawer()
	}
	return docsDrawer
}

func buildDocsDrawer() *FDrawer {
	return Drawer(
		Container(
			Column(
				[]Widget{
					drawerHeader(),
					drawerNavigation(),
				},
			).Gap(8).
				Flex(1),
		),
	).ID("docs-left-drawer").Width(AllBP(320)).Transition(0.3)
}

func drawerHeader() Widget {
	return Container(
		Row(
			[]Widget{
				drawerTitle(),
				Spacer(),
				Button(
					Icon(icondata.Close).
						Fill("#6B7280"),
				).OnClick(func(this Widget, e Event) {
					DocsDrawer().Hide()
				}),
			},
		).Gap(8).
			Flex(1).
			CrossAxisAlignment(options.AxisAlignmentTypeCenter),
	).Padding(AllBP(LRTB(20, 20, 16, 8))).
		BorderColor("#E5E7EB").
		BorderWidth(0, 0, 1, 0).
		BorderStyle(options.BorderStyleTypeSolid)
}

func drawerTitle() Widget {
	return Column(
		[]Widget{
			Text("Documentation").
				FontSize(18).
				FontColor("#1F2937").
				FontWeight("600").
				UserSelect(options.UserSelectTypeNone),
			Text("Learn how to build with gofred").
				FontSize(12).
				FontColor("#6B7280").
				FontWeight("400").
				UserSelect(options.UserSelectTypeNone),
		},
	).Gap(2)
}

func drawerNavigation() Widget {
	navigate := hooks.UseNavigate()
	activeHref := navigate.Path()

	return Container(
		Column(
			[]Widget{
				drawerNavSection("Getting Started", []drawerNavItem{
					{title: "Installation", href: "/docs/installation", icon: icondata.Download},
					{title: "Quick Start", href: "/docs/quick-start", icon: icondata.RocketLaunchOutline},
					{title: "Your First App", href: "/docs/first-app", icon: icondata.Application},
					{title: "Project Structure", href: "/docs/project-structure", icon: icondata.Folder},
				}, activeHref),
				Spacer(),
				drawerNavSection("Core Concepts", []drawerNavItem{
					{title: "Widgets", href: "/docs/widgets", icon: icondata.Widgets},
					{title: "Layouts", href: "/docs/layouts", icon: icondata.Grid},
					{title: "Styling", href: "/docs/styling", icon: icondata.PaletteOutline},
					{title: "State Management", href: "/docs/state", icon: icondata.Database},
					{title: "Event Handling", href: "/docs/events", icon: icondata.Mouse},
				}, activeHref),
				Spacer(),
				drawerNavSection("Components", []drawerNavItem{
					{title: "Buttons", href: "/docs/components/buttons", icon: icondata.ButtonPointer},
					{title: "Navigation", href: "/docs/components/navigation", icon: icondata.Menu},
					{title: "Icons", href: "/docs/components/icons", icon: icondata.Star},
					{title: "Images", href: "/docs/components/images", icon: icondata.Image},
					{title: "Containers", href: "/docs/components/containers", icon: icondata.Package},
				}, activeHref),
				Spacer(),
				drawerNavSection("Advanced", []drawerNavItem{
					{title: "Routing", href: "/docs/routing", icon: icondata.SignDirection},
					{title: "API Reference", href: "/docs/api", icon: icondata.FileDocument},
					{title: "Best Practices", href: "/docs/best-practices", icon: icondata.ThumbUp},
					{title: "Performance", href: "/docs/performance", icon: icondata.Speedometer},
					{title: "Deployment", href: "/docs/deployment", icon: icondata.Server},
				}, activeHref),
				Spacer(),
				drawerNavSection("Resources", []drawerNavItem{
					{title: "Examples", href: "/docs/examples", icon: icondata.Lightbulb},
					{title: "Tutorials", href: "/docs/tutorials", icon: icondata.Book},
					{title: "Community", href: "/docs/community", icon: icondata.AccountGroup},
					{title: "Support", href: "/docs/support", icon: icondata.Help},
				}, activeHref),
			},
		).Gap(8).
			Flex(1),
	).Padding(AllBP(LRTB(20, 20, 20, 20)))
}

type drawerNavItem struct {
	title string
	href  string
	icon  icondata.IconData
}

func drawerNavSection(title string, items []drawerNavItem, activeHref string) Widget {
	var sectionItems []Widget

	// Section title
	sectionItems = append(sectionItems, Text(title).
		FontSize(12).
		FontColor("#6B7280").
		FontWeight("600").
		UserSelect(options.UserSelectTypeNone))

	// Section items
	for _, item := range items {
		sectionItems = append(sectionItems, drawerNavItemWidget(item, activeHref))
	}

	return Column(
		sectionItems,
	).Gap(4)
}

func drawerNavItemWidget(item drawerNavItem, activeHref string) Widget {
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

	return Link(
		Container(
			Row(
				[]Widget{
					Icon(item.icon).
						Width(AllBP(20)).
						Height(AllBP(20)).
						Fill(iconColor),
					Spacer(),
					Text(item.title).
						FontSize(14).
						FontColor(textColor).
						FontWeight("400").
						UserSelect(options.UserSelectTypeNone),
				},
			).Gap(8).
				CrossAxisAlignment(options.AxisAlignmentTypeCenter),
		).Padding(AllBP(Axis(8, 12))).
			BackgroundColor(backgroundColor).
			BorderRadius(6),
	).Href(item.href).
		OnClick(func(this Widget, e Event) {
			DocsDrawer().Hide()
		})
}
