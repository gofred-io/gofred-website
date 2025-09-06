package docs

import (
	comingsoon "github.com/gofred-io/gofred-website/app/components/coming_soon"
	"github.com/gofred-io/gofred-website/app/components/footer"
	"github.com/gofred-io/gofred-website/app/components/header"
	. "github.com/gofred-io/gofred-website/app/pages/404"
	"github.com/gofred-io/gofred-website/app/pages/docs/core_concepts"
	"github.com/gofred-io/gofred-website/app/pages/docs/getting_started"

	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/foundation/button"
	. "github.com/gofred-io/gofred/foundation/column"
	. "github.com/gofred-io/gofred/foundation/container"
	. "github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	. "github.com/gofred-io/gofred/foundation/link"
	. "github.com/gofred-io/gofred/foundation/router"
	. "github.com/gofred-io/gofred/foundation/row"
	. "github.com/gofred-io/gofred/foundation/spacer"
	. "github.com/gofred-io/gofred/foundation/text"
	. "github.com/gofred-io/gofred/options"
	. "github.com/gofred-io/gofred/options/spacing"
	. "github.com/gofred-io/gofred/widget"
)

func Docs(params RouteParams) Widget {
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
		return NotFoundPage(params)
	}
}

func docsPageTemplate(content Widget) Widget {
	return Column(
		[]Widget{
			header.Get(),
			docsMainContent(content),
			footer.Get(),
		},
	).Flex(1)
}

func docsMainContent(content Widget) Widget {
	return Container(
		Row(
			[]Widget{
				docsSidebar(),
				contentArea(content),
			},
		).Flex(1),
	).Flex(1).
		BackgroundColor("#F8F9FA")
}

func contentArea(content Widget) Widget {
	return Container(
		Column(
			[]Widget{
				docsMobileMenuButton(),
				content,
			},
		).Gap(16).
			Flex(1),
	).Flex(1).
		Padding(AllBP(All(32)))
}

func docsMobileMenuButton() Widget {
	return Button(
		Icon(icondata.Menu).
			Fill("#6B7280").
			Width(AllBP(42)).
			Height(AllBP(42)),
	).Tooltip("Open documentation menu").
		Visible(
			XS(true),
			SM(true),
			MD(true),
		).
		OnClick(func(this Widget, e Event) {
			DocsDrawer().Show()
		})
}

func docsPageContent() Widget {
	return Column(
		[]Widget{
			//heroSection(),
			featuresTitle(),
			featuresSection(),
		},
	).Gap(48).
		CrossAxisAlignment(AxisAlignmentTypeCenter)
}

func heroSection() Widget {
	return Column(
		[]Widget{
			Text("Documentation").
				FontSize(48).
				FontColor("#1F2937").
				FontWeight("700").
				UserSelect(UserSelectTypeNone),
			Spacer(),
			Text("Learn how to build modern web applications with gofred").
				FontSize(20).
				FontColor("#6B7280").
				FontWeight("400").
				UserSelect(UserSelectTypeNone),
			Spacer(),
			Text("Write Go code, get WebAssembly apps. No JavaScript required.").
				FontSize(16).
				FontColor("#9CA3AF").
				FontWeight("400").
				UserSelect(UserSelectTypeNone),
		},
	).Gap(8).
		CrossAxisAlignment(AxisAlignmentTypeCenter)
}

func featuresSection() Widget {
	return Container(
		Column(
			[]Widget{
				Row(
					[]Widget{
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
					},
				).Gap(24),
				Row(
					[]Widget{
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
				).Gap(24),
			},
		).Gap(24),
	).Width(
		AllBP(960),
		XS(240),
		SM(480),
		MD(640),
	)
}

func featureCard(iconData icondata.IconData, title, description, color string, href string) Widget {
	return Container(
		Link(
			Container(
				Column(
					[]Widget{
						Container(
							Icon(iconData).
								Width(AllBP(32)).
								Height(AllBP(32)).
								Fill(color),
						).BackgroundColor("#FFFFFF").
							BorderRadius(12).
							Padding(AllBP(All(16))).
							BorderColor("#E5E7EB").
							BorderWidth(1, 1, 1, 1).
							BorderStyle(BorderStyleTypeSolid),
						Spacer(),
						Text(title).
							FontSize(18).
							FontColor("#1F2937").
							FontWeight("600").
							UserSelect(UserSelectTypeNone),
						Spacer(),
						Text(description).
							FontSize(14).
							FontColor("#6B7280").
							FontWeight("400").
							UserSelect(UserSelectTypeNone).
							LineHeight(1.5),
					},
				).Gap(8).
					CrossAxisAlignment(AxisAlignmentTypeCenter),
			).Padding(AllBP(Axis(24, 48))),
		).Href(href).
			OnClick(func(this Widget, e Event) {
				DocsDrawer().Hide()
			}),
	).Flex(1).
		BackgroundColor("#FFFFFF").
		BorderRadius(12).
		BorderColor("#E5E7EB").
		BorderWidth(1, 1, 1, 1).
		BorderStyle(BorderStyleTypeSolid)
}

func featuresTitle() Widget {
	return Column(
		[]Widget{
			Text("Ready to get started?").
				FontSize(24).
				FontColor("#1F2937").
				FontWeight("600").
				UserSelect(UserSelectTypeNone).
				Align(TextAlignTypeCenter),
			Spacer(),
			Text("Follow our step-by-step installation guide to set up your development environment.").
				FontSize(16).
				FontColor("#6B7280").
				FontWeight("400").
				UserSelect(UserSelectTypeNone),
		},
	).Gap(8).
		CrossAxisAlignment(AxisAlignmentTypeCenter)
}
