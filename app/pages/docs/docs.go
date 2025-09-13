package docs

import (
	comingsoon "github.com/gofred-io/gofred-website/app/components/coming_soon"
	"github.com/gofred-io/gofred-website/app/components/footer"
	"github.com/gofred-io/gofred-website/app/components/header"
	notfound "github.com/gofred-io/gofred-website/app/pages/404"
	"github.com/gofred-io/gofred-website/app/pages/docs/core_concepts"
	"github.com/gofred-io/gofred-website/app/pages/docs/getting_started"

	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/grid"
	"github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	iconbutton "github.com/gofred-io/gofred/foundation/icon_button"
	"github.com/gofred-io/gofred/foundation/link"
	"github.com/gofred-io/gofred/foundation/router"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
)

func New(params router.RouteParams) application.BaseWidget {
	section := params.Get("section")

	switch section {
	case "":
		return docsPageTemplate(docsPageContent())
	case "installation":
		return docsPageTemplate(getting_started.InstallationContent())
	case "quick-start":
		return docsPageTemplate(getting_started.QuickStartContent())
	case "first-app":
		return docsPageTemplate(getting_started.FirstAppContent())
	case "project-structure":
		return docsPageTemplate(getting_started.ProjectStructureContent())
	case "widgets":
		return docsPageTemplate(core_concepts.WidgetsContent())
	case "layouts":
		return docsPageTemplate(core_concepts.LayoutsContent())
	case "styling":
		return docsPageTemplate(core_concepts.StylingContent())
	case "state":
		return docsPageTemplate(core_concepts.StateManagementContent())
	case "events":
		return docsPageTemplate(core_concepts.EventHandlingContent())
	case "buttons", "navigation", "icons", "images", "containers",
		"routing", "api", "best-practices", "performance", "deployment",
		"examples", "tutorials", "community", "support":
		return docsPageTemplate(comingsoon.ComingSoonContent("Coming Soon", []comingsoon.Suggestion{
			{Title: "Buttons", Description: "Learn about buttons and how to use them", Href: "/docs/buttons"},
			{Title: "Navigation", Description: "Learn about navigation and how to use it", Href: "/docs/navigation"},
			{Title: "Icons", Description: "Learn about icons and how to use them", Href: "/docs/icons"},
			{Title: "Images", Description: "Learn about images and how to use them", Href: "/docs/images"},
			{Title: "Containers", Description: "Learn about containers and how to use them", Href: "/docs/containers"},
		}))
	default:
		return notfound.New(params)
	}
}

func docsPageTemplate(content application.BaseWidget) application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			header.Get(),
			docsMainContent(content),
			footer.Get(),
		},
		column.Flex(1),
	)
}

func docsMainContent(content application.BaseWidget) application.BaseWidget {
	return container.New(
		row.New(
			[]application.BaseWidget{
				docsSidebar(),
				contentArea(content),
			},
			row.Flex(1),
		),
		container.Flex(1),
		container.BackgroundColor("#F8F9FA"),
	)
}

func contentArea(content application.BaseWidget) application.BaseWidget {
	return container.New(
		column.New(
			[]application.BaseWidget{
				docsMobileMenuButton(),
				content,
			},
			column.Gap(16),
			column.Flex(1),
		),
		container.Flex(1),
		container.Padding(breakpoint.All(spacing.All(32))),
	)
}

func docsMobileMenuButton() application.BaseWidget {
	return iconbutton.New(
		icondata.Menu,
		iconbutton.Fill("#6B7280"),
		iconbutton.Tooltip("Open documentation menu"),
		iconbutton.Visible(
			breakpoint.XS(true),
			breakpoint.SM(true),
			breakpoint.MD(true),
		),
		iconbutton.Width(breakpoint.All(42)),
		iconbutton.Height(breakpoint.All(42)),
		iconbutton.OnClick(func(this application.BaseWidget, e application.Event) {
			docsDrawer.Show()
		}),
	)
}

func docsPageContent() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			//heroSection(),
			featuresTitle(),
			featuresSection(),
		},
		column.Gap(48),
		column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
	)
}

func heroSection() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			text.New(
				"Documentation",
				text.FontSize(48),
				text.FontWeight("700"),
				text.UserSelect(theme.UserSelectTypeNone),
			),
			spacer.New(spacer.Height(16)),
			text.New(
				"Learn how to build modern web applications with gofred",
				text.FontSize(20),
				text.FontColor("#6B7280"),
				text.UserSelect(theme.UserSelectTypeNone),
			),
			spacer.New(spacer.Height(8)),
			text.New(
				"Write Go code, get WebAssembly apps. No JavaScript required.",
				text.FontSize(16),
				text.FontColor("#9CA3AF"),
				text.UserSelect(theme.UserSelectTypeNone),
			),
		},
		column.Gap(8),
		column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
	)
}

func featuresSection() application.BaseWidget {
	return container.New(
		grid.New(
			[]application.BaseWidget{
				featureCard(
					icondata.RocketLaunchOutline,
					"Get Started",
					"Get up and running in minutes with our simple installation guide and first app tutorial.",
					"#2B799B",
					"/docs/installation",
				),
				featureCard(
					icondata.Tools,
					"Core Concepts",
					"Learn about widgets, layouts, styling, and state management in gofred applications.",
					"#10B981",
					"/docs/core-concepts",
				),
				featureCard(
					icondata.PaletteOutline,
					"Components",
					"Explore the rich set of UI components available for building beautiful interfaces.",
					"#F59E0B",
					"/docs/components",
				),
				featureCard(
					icondata.Cog,
					"Advanced",
					"Master advanced concepts like routing, performance optimization, and deployment strategies.",
					"#8B5CF6",
					"/docs/advanced",
				),
			},
			grid.RowGap(24),
			grid.ColumnGap(24),
			grid.ColumnCount(
				breakpoint.All(4),
				breakpoint.XS(1),
				breakpoint.SM(2),
				breakpoint.MD(2),
			),
		),
		container.Width(
			breakpoint.All(960),
			breakpoint.XS(240),
			breakpoint.SM(480),
			breakpoint.MD(640),
		),
	)
}

func featureCard(iconData icondata.IconData, title, description, color string, href string) application.BaseWidget {
	return container.New(
		link.New(
			container.New(
				column.New(
					[]application.BaseWidget{
						container.New(
							icon.New(
								iconData,
								icon.Width(breakpoint.All(32)),
								icon.Height(breakpoint.All(32)),
								icon.Fill(color),
							),
							container.BackgroundColor("#FFFFFF"),
							container.BorderRadius(12),
							container.Padding(breakpoint.All(spacing.All(16))),
							container.BorderColor("#E5E7EB"),
							container.BorderWidth(spacing.All(1)),
							container.BorderStyle(theme.BorderStyleTypeSolid),
						),
						spacer.New(spacer.Height(16)),
						text.New(
							title,
							text.FontSize(18),
							text.FontWeight("600"),
							text.UserSelect(theme.UserSelectTypeNone),
						),
						spacer.New(spacer.Height(8)),
						text.New(
							description,
							text.FontSize(14),
							text.FontColor("#6B7280"),
							text.UserSelect(theme.UserSelectTypeNone),
							text.LineHeight(1.5),
						),
					},
					column.Gap(8),
					column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
				),
				container.Padding(breakpoint.All(spacing.Axis(24, 48))),
			),
			link.Href(href),
			link.OnClick(func(this application.BaseWidget, e application.Event) {
				docsDrawer.Hide()
			}),
		),
		container.Flex(1),
		container.BackgroundColor("#FFFFFF"),
		container.BorderRadius(12),
		container.BorderColor("#E5E7EB"),
		container.BorderWidth(spacing.All(1)),
		container.BorderStyle(theme.BorderStyleTypeSolid),
	)
}

func featuresTitle() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			text.New(
				"Ready to get started?",
				text.FontSize(24),
				text.FontWeight("600"),
				text.UserSelect(theme.UserSelectTypeNone),
				text.Align(theme.TextAlignTypeCenter),
			),
			spacer.New(spacer.Height(16)),
			text.New(
				"Follow our step-by-step installation guide to set up your development environment.",
				text.FontSize(16),
				text.FontColor("#6B7280"),
				text.UserSelect(theme.UserSelectTypeNone),
			),
		},
		column.Gap(8),
		column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
	)
}
