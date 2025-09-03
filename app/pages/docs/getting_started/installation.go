package getting_started

import (
	codeblock "github.com/gofred-io/gofred-website/app/components/code_block"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/button"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	"github.com/gofred-io/gofred/foundation/link"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/widget"
)

func InstallationContent() widget.BaseWidget {
	return container.New(
		column.New(
			[]widget.BaseWidget{
				installationPageHeader(),
				spacer.New(spacer.Height(24)),
				installationPageContent(),
			},
			column.Gap(16),
			column.Flex(1),
		),
		container.Flex(1),
		container.Padding(breakpoint.All(spacing.All(32))),
	)
}

func installationPageHeader() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			text.New(
				"Installation",
				text.FontSize(32),
				text.FontColor("#1F2937"),
				text.FontWeight("700"),
			),
			text.New(
				"Get started with gofred by installing the framework and setting up your development environment.",
				text.FontSize(18),
				text.FontColor("#6B7280"),
				text.FontWeight("400"),
			),
		},
		column.Gap(8),
	)
}

func installationPageContent() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			contentSection("Prerequisites", "Before you begin, make sure you have the following installed on your system:"),
			prerequisitesList(),
			spacer.New(spacer.Height(24)),
			contentSection("Installation", "Install gofred using Go modules:"),
			codeblock.New("go get github.com/gofred-io/gofred"),
			spacer.New(spacer.Height(24)),
			contentSection("Next Steps", "Now that you have gofred installed, you can:"),
			installationNextSteps(),
			spacer.New(spacer.Height(32)),
			navigationButtons("/docs", "/docs/quick-start"),
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
			),
			text.New(
				description,
				text.FontSize(16),
				text.FontColor("#6B7280"),
				text.FontWeight("400"),
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
				icondata.Check,
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
			),
		},
		row.Gap(8),
		row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
	)
}

func installationNextSteps() widget.BaseWidget {
	steps := []struct {
		title       string
		description string
		href        string
	}{
		{
			title:       "Quick Start",
			description: "Create your first gofred application",
			href:        "/docs/quick-start",
		},
		{
			title:       "Build Your First App",
			description: "Create a simple application step by step",
			href:        "/docs/first-app",
		},
		{
			title:       "Project Structure",
			description: "Explore the project structure",
			href:        "/docs/project-structure",
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
							),
							text.New(
								description,
								text.FontSize(14),
								text.FontColor("#6B7280"),
								text.FontWeight("400"),
							),
						},
						column.Gap(4),
						column.Flex(1),
					),
					icon.New(
						icondata.ChevronRight,
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

func navigationButtons(previousHref, nextHref string) widget.BaseWidget {
	return row.New(
		[]widget.BaseWidget{
			link.New(
				button.New(
					row.New(
						[]widget.BaseWidget{
							icon.New(
								icondata.ChevronLeft,
								icon.Width(breakpoint.All(16)),
								icon.Height(breakpoint.All(16)),
								icon.Fill("#FFFFFF"),
							),
							text.New(
								"Previous",
								text.FontSize(14),
								text.FontColor("#FFFFFF"),
								text.FontWeight("500"),
							),
						},
						row.Gap(8),
						row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
					),
					button.BackgroundColor("#6B7280"),
					button.Width(breakpoint.All(120)),
				),
				link.Href(previousHref),
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
							),
							icon.New(
								icondata.ChevronRight,
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
				link.Href(nextHref),
			),
		},
		row.Flex(1),
	)
}
