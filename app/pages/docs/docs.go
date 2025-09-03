package docs

import (
	comingsoon "github.com/gofred-io/gofred-website/app/components/coming_soon"
	notfound "github.com/gofred-io/gofred-website/app/pages/404"
	"github.com/gofred-io/gofred-website/app/pages/docs/core_concepts"
	"github.com/gofred-io/gofred-website/app/pages/docs/getting_started"
	"github.com/gofred-io/gofred-website/app/pages/home"

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
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/widget"
)

func New(params router.RouteParams) widget.BaseWidget {
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
	case "buttons", "navigation", "icons", "images", "containers", "routing", "api", "best-practices", "performance", "deployment":
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

func docsPageTemplate(content widget.BaseWidget) widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			home.Header(),
			docsMainContent(content),
			home.Footer(),
		},
		column.Flex(1),
	)
}

func docsMainContent(content widget.BaseWidget) widget.BaseWidget {
	return container.New(
		row.New(
			[]widget.BaseWidget{
				docsSidebar(),
				contentArea(content),
			},
			row.Flex(1),
		),
		container.Flex(1),
		container.BackgroundColor("#F8F9FA"),
	)
}

func contentArea(content widget.BaseWidget) widget.BaseWidget {
	return container.New(
		column.New(
			[]widget.BaseWidget{
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

func docsMobileMenuButton() widget.BaseWidget {
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
		iconbutton.OnClick(func(this widget.BaseWidget, e widget.Event) {
			docsDrawer.Show()
		}),
	)
}

func docsPageContent() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			//heroSection(),
			featuresTitle(),
			featuresSection(),
		},
		column.Gap(48),
		column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
	)
}

func heroSection() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			text.New(
				"Documentation",
				text.FontSize(48),
				text.FontColor("#1F2937"),
				text.FontWeight("700"),
				text.UserSelect(options.UserSelectTypeNone),
			),
			spacer.New(spacer.Height(16)),
			text.New(
				"Learn how to build modern web applications with gofred",
				text.FontSize(20),
				text.FontColor("#6B7280"),
				text.FontWeight("400"),
				text.UserSelect(options.UserSelectTypeNone),
			),
			spacer.New(spacer.Height(8)),
			text.New(
				"Write Go code, get WebAssembly apps. No JavaScript required.",
				text.FontSize(16),
				text.FontColor("#9CA3AF"),
				text.FontWeight("400"),
				text.UserSelect(options.UserSelectTypeNone),
			),
		},
		column.Gap(8),
		column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
	)
}

func featuresSection() widget.BaseWidget {
	return container.New(
		grid.New(
			[]widget.BaseWidget{
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

func featureCard(iconData icondata.IconData, title, description, color string, href string) widget.BaseWidget {
	return container.New(
		link.New(
			container.New(
				column.New(
					[]widget.BaseWidget{
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
							container.BorderWidth(1, 1, 1, 1),
							container.BorderStyle(options.BorderStyleTypeSolid),
						),
						spacer.New(spacer.Height(16)),
						text.New(
							title,
							text.FontSize(18),
							text.FontColor("#1F2937"),
							text.FontWeight("600"),
							text.UserSelect(options.UserSelectTypeNone),
						),
						spacer.New(spacer.Height(8)),
						text.New(
							description,
							text.FontSize(14),
							text.FontColor("#6B7280"),
							text.FontWeight("400"),
							text.UserSelect(options.UserSelectTypeNone),
							text.LineHeight(1.5),
						),
					},
					column.Gap(8),
					column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
				),
				container.Padding(breakpoint.All(spacing.Axis(24, 48))),
			),
			link.Href(href),
			link.OnClick(func(this widget.BaseWidget, e widget.Event) {
				docsDrawer.Hide()
			}),
		),
		container.Flex(1),
		container.BackgroundColor("#FFFFFF"),
		container.BorderRadius(12),
		container.BorderColor("#E5E7EB"),
		container.BorderWidth(1, 1, 1, 1),
		container.BorderStyle(options.BorderStyleTypeSolid),
	)
}

func featuresTitle() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			text.New(
				"Ready to get started?",
				text.FontSize(24),
				text.FontColor("#1F2937"),
				text.FontWeight("600"),
				text.UserSelect(options.UserSelectTypeNone),
			),
			spacer.New(spacer.Height(16)),
			text.New(
				"Follow our step-by-step installation guide to set up your development environment.",
				text.FontSize(16),
				text.FontColor("#6B7280"),
				text.FontWeight("400"),
				text.UserSelect(options.UserSelectTypeNone),
			),
		},
		column.Gap(8),
		column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
	)
}
