package docs

import (
	"github.com/gofred-io/gofred-website/app/pages/home"

	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/button"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
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

func New() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			home.Header(),
			mainContent(),
			home.Footer(),
		},
		column.Flex(1),
	)
}

func mainContent() widget.BaseWidget {
	return container.New(
		row.New(
			[]widget.BaseWidget{
				sidebar(),
				contentArea(),
			},
			row.Flex(1),
		),
		container.Flex(1),
		container.BackgroundColor("#F8F9FA"),
	)
}

func sidebar() widget.BaseWidget {
	return container.New(
		column.New(
			[]widget.BaseWidget{
				sidebarHeader(),
				spacer.New(spacer.Height(16)),
				navigationMenu(),
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
			breakpoint.MD(true),
			breakpoint.LG(true),
			breakpoint.XL(true),
			breakpoint.XXL(true),
		),
	)
}

func sidebarHeader() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			text.New(
				"Documentation",
				text.FontSize(20),
				text.FontColor("#1F2937"),
				text.FontWeight("600"),
				text.UserSelect(options.UserSelectTypeNone),
			),
			text.New(
				"Learn how to build with gofred",
				text.FontSize(14),
				text.FontColor("#6B7280"),
				text.FontWeight("400"),
				text.UserSelect(options.UserSelectTypeNone),
			),
		},
		column.Gap(4),
	)
}

func navigationMenu() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			navSection("Getting Started", []navItem{
				{title: "Installation", href: "/docs/installation", active: true},
				{title: "Quick Start", href: "/docs/quick-start"},
				{title: "Your First App", href: "/docs/first-app"},
			}),
			spacer.New(spacer.Height(16)),
			navSection("Core Concepts", []navItem{
				{title: "Widgets", href: "/docs/widgets"},
				{title: "Layouts", href: "/docs/layouts"},
				{title: "Styling", href: "/docs/styling"},
				{title: "State Management", href: "/docs/state"},
			}),
			spacer.New(spacer.Height(16)),
			navSection("Components", []navItem{
				{title: "Buttons", href: "/docs/components/buttons"},
				{title: "Forms", href: "/docs/components/forms"},
				{title: "Navigation", href: "/docs/components/navigation"},
				{title: "Icons", href: "/docs/components/icons"},
			}),
			spacer.New(spacer.Height(16)),
			navSection("Advanced", []navItem{
				{title: "Routing", href: "/docs/routing"},
				{title: "API Reference", href: "/docs/api"},
				{title: "Best Practices", href: "/docs/best-practices"},
				{title: "Performance", href: "/docs/performance"},
			}),
		},
		column.Gap(8),
	)
}

type navItem struct {
	title  string
	href   string
	active bool
}

func navSection(title string, items []navItem) widget.BaseWidget {
	var sectionItems []widget.BaseWidget

	// Section title
	sectionItems = append(sectionItems, text.New(
		title,
		text.FontSize(14),
		text.FontColor("#374151"),
		text.FontWeight("600"),
		text.UserSelect(options.UserSelectTypeNone),
	))

	// Section items
	for _, item := range items {
		sectionItems = append(sectionItems, navItemWidget(item))
	}

	return column.New(
		sectionItems,
		column.Gap(4),
	)
}

func navItemWidget(item navItem) widget.BaseWidget {
	var textColor string
	var backgroundColor string

	if item.active {
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

func contentArea() widget.BaseWidget {
	return container.New(
		column.New(
			[]widget.BaseWidget{
				mobileMenuButton(),
				pageHeader(),
				spacer.New(spacer.Height(24)),
				pageContent(),
			},
			column.Gap(16),
			column.Flex(1),
		),
		container.Flex(1),
		container.Padding(breakpoint.All(spacing.All(32))),
		container.MaxWidth(800),
	)
}

func mobileMenuButton() widget.BaseWidget {
	return iconbutton.New(
		icondata.Menu,
		iconbutton.Fill("#6B7280"),
		iconbutton.Tooltip("Open documentation menu"),
		iconbutton.Visible(
			breakpoint.XS(true),
			breakpoint.SM(true),
		),
	)
}

func pageHeader() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			text.New(
				"Installation",
				text.FontSize(32),
				text.FontColor("#1F2937"),
				text.FontWeight("700"),
				text.UserSelect(options.UserSelectTypeNone),
			),
			text.New(
				"Get started with gofred by installing the framework and setting up your development environment.",
				text.FontSize(18),
				text.FontColor("#6B7280"),
				text.FontWeight("400"),
				text.UserSelect(options.UserSelectTypeNone),
			),
		},
		column.Gap(8),
	)
}

func pageContent() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			contentSection("Prerequisites", "Before you begin, make sure you have the following installed on your system:"),
			prerequisitesList(),
			spacer.New(spacer.Height(24)),
			contentSection("Installation", "Install gofred using Go modules:"),
			codeBlock("go mod init your-project\ngo get github.com/gofred-io/gofred"),
			spacer.New(spacer.Height(24)),
			contentSection("Quick Start", "Create your first gofred application:"),
			codeBlock(`package main

import (
    "github.com/gofred-io/gofred/foundation/text"
    "github.com/gofred-io/gofred/widget"
)

func main() {
    app := text.New("Hello, gofred!")
    widget.Run(app)
}`),
			spacer.New(spacer.Height(24)),
			contentSection("Next Steps", "Now that you have gofred installed, you can:"),
			nextStepsList(),
			spacer.New(spacer.Height(32)),
			navigationButtons(),
		},
		column.Gap(16),
	)
}

func contentSection(title, description string) widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			text.New(
				title,
				text.FontSize(24),
				text.FontColor("#1F2937"),
				text.FontWeight("600"),
				text.UserSelect(options.UserSelectTypeNone),
			),
			text.New(
				description,
				text.FontSize(16),
				text.FontColor("#6B7280"),
				text.FontWeight("400"),
				text.UserSelect(options.UserSelectTypeNone),
			),
		},
		column.Gap(8),
	)
}

func prerequisitesList() widget.BaseWidget {
	items := []string{
		"Go 1.21 or later",
		"A modern web browser with WebAssembly support",
		"Basic knowledge of Go programming",
	}

	var listItems []widget.BaseWidget
	for _, item := range items {
		listItems = append(listItems, listItem(item))
	}

	return column.New(
		listItems,
		column.Gap(8),
	)
}

func listItem(itemText string) widget.BaseWidget {
	return row.New(
		[]widget.BaseWidget{
			icon.New(
				icondata.Tools,
				icon.Width(breakpoint.All(16)),
				icon.Height(breakpoint.All(16)),
				icon.Fill("#10B981"),
			),
			spacer.New(spacer.Width(8)),
			text.New(
				itemText,
				text.FontSize(16),
				text.FontColor("#374151"),
				text.FontWeight("400"),
				text.UserSelect(options.UserSelectTypeNone),
			),
		},
		row.Gap(8),
		row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
	)
}

func codeBlock(code string) widget.BaseWidget {
	return container.New(
		text.New(
			code,
			text.FontSize(14),
			text.FontColor("#F3F4F6"),
			text.FontWeight("400"),
			text.UserSelect(options.UserSelectTypeNone),
		),
		container.BackgroundColor("#1F2937"),
		container.Padding(breakpoint.All(spacing.All(16))),
		container.BorderRadius(8),
	)
}

func nextStepsList() widget.BaseWidget {
	steps := []struct {
		title       string
		description string
		href        string
	}{
		{
			title:       "Learn the Basics",
			description: "Understand widgets, layouts, and styling",
			href:        "/docs/quick-start",
		},
		{
			title:       "Build Your First App",
			description: "Create a simple application step by step",
			href:        "/docs/first-app",
		},
		{
			title:       "Explore Components",
			description: "Discover available UI components",
			href:        "/docs/components",
		},
	}

	var stepItems []widget.BaseWidget
	for _, step := range steps {
		stepItems = append(stepItems, nextStepItem(step.title, step.description, step.href))
	}

	return column.New(
		stepItems,
		column.Gap(12),
	)
}

func nextStepItem(title, description, href string) widget.BaseWidget {
	return link.New(
		container.New(
			row.New(
				[]widget.BaseWidget{
					column.New(
						[]widget.BaseWidget{
							text.New(
								title,
								text.FontSize(16),
								text.FontColor("#2B799B"),
								text.FontWeight("500"),
								text.UserSelect(options.UserSelectTypeNone),
							),
							text.New(
								description,
								text.FontSize(14),
								text.FontColor("#6B7280"),
								text.FontWeight("400"),
								text.UserSelect(options.UserSelectTypeNone),
							),
						},
						column.Gap(4),
						column.Flex(1),
					),
					icon.New(
						icondata.OpenInNew,
						icon.Width(breakpoint.All(20)),
						icon.Height(breakpoint.All(20)),
						icon.Fill("#9CA3AF"),
					),
				},
				row.Gap(12),
				row.Flex(1),
				row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
			),
			container.Padding(breakpoint.All(spacing.All(16))),
			container.BackgroundColor("#FFFFFF"),
			container.BorderRadius(8),
			container.BorderColor("#E5E7EB"),
			container.BorderWidth(1, 1, 1, 1),
			container.BorderStyle(options.BorderStyleTypeSolid),
		),
		link.Href(href),
	)
}

func navigationButtons() widget.BaseWidget {
	return row.New(
		[]widget.BaseWidget{
			link.New(
				button.New(
					row.New(
						[]widget.BaseWidget{
							icon.New(
								icondata.Menu,
								icon.Width(breakpoint.All(16)),
								icon.Height(breakpoint.All(16)),
								icon.Fill("#FFFFFF"),
							),
							text.New(
								"Previous",
								text.FontSize(14),
								text.FontColor("#FFFFFF"),
								text.FontWeight("500"),
								text.UserSelect(options.UserSelectTypeNone),
							),
						},
						row.Gap(8),
						row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
					),
					button.BackgroundColor("#6B7280"),
					button.Width(breakpoint.All(120)),
				),
				link.Href("/docs"),
			),
			spacer.New(),
			link.New(
				button.New(
					row.New(
						[]widget.BaseWidget{
							text.New(
								"Next",
								text.FontSize(14),
								text.FontColor("#FFFFFF"),
								text.FontWeight("500"),
								text.UserSelect(options.UserSelectTypeNone),
							),
							icon.New(
								icondata.OpenInNew,
								icon.Width(breakpoint.All(16)),
								icon.Height(breakpoint.All(16)),
								icon.Fill("#FFFFFF"),
							),
						},
						row.Gap(8),
						row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
					),
					button.BackgroundColor("#2B799B"),
					button.Width(breakpoint.All(120)),
				),
				link.Href("/docs/quick-start"),
			),
		},
		row.Flex(1),
	)
}
