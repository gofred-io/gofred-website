package getting_started

import (
	"github.com/gofred-io/gofred-website/app/components/codeblock"
	appTheme "github.com/gofred-io/gofred-website/app/theme"

	"github.com/gofred-io/gofred/application"
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
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
)

func InstallationContent() application.BaseWidget {
	return container.New(
		column.New(
			[]application.BaseWidget{
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

func installationPageHeader() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			text.New(
				"Installation",
				text.FontSize(32),
				text.FontWeight("700"),
			),
			text.New(
				"Install the gofred CLI tool and set up your development environment for Go WebAssembly applications.",
				text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
				text.FontSize(18),
			),
		},
		column.Gap(8),
	)
}

func installationPageContent() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			contentSection("Prerequisites", "Before you begin, make sure you have the following installed on your system:"),
			prerequisitesList(),
			spacer.New(spacer.Height(24)),
			contentSection("Installation", "Install the gofred CLI tool using the installation script:"),
			codeblock.New(`curl -fsSL https://raw.githubusercontent.com/gofred-io/gofred-cli/refs/heads/master/install.sh | bash`),
			text.New(
				"This script will detect your operating system and architecture, download the appropriate binary, and install it to ~/.local/bin (or ~/AppData/Local/bin on Windows).",
				text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
				text.FontSize(14),
			),
			spacer.New(spacer.Height(16)),
			contentSection("Verify Installation", "Check that gofred is installed correctly:"),
			codeblock.New(`gofred version`),
			spacer.New(spacer.Height(24)),
			contentSection("Create Your First App", "Create a new Go WebAssembly application:"),
			codeblock.New(`gofred app create my-app --package my-app`),
			text.New(
				"This will create a complete project structure with main.go, web assets, and VS Code configuration.",
				text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
				text.FontSize(14),
			),
			spacer.New(spacer.Height(16)),
			contentSection("Run Your Application", "Navigate to your app directory and start the development server:"),
			codeblock.New(`cd my-app
gofred app run`),
			text.New(
				"This will compile your Go code to WebAssembly, start a development server, and automatically open your browser with hot reload enabled.",
				text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
				text.FontSize(14),
			),
			spacer.New(spacer.Height(24)),
			contentSection("Next Steps", "Now that you have gofred installed, you can:"),
			installationNextSteps(),
			spacer.New(spacer.Height(32)),
			navigationButtons("/docs", "/docs/quick-start"),
		},
		column.Gap(16),
	)
}

func contentSection(title, description string) application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			text.New(
				title,
				text.FontSize(24),
				text.FontWeight("600"),
			),
			text.New(
				description,
				text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
				text.FontSize(16),
			),
		},
		column.Gap(8),
	)
}

func prerequisitesList() application.BaseWidget {
	items := []string{
		"Go 1.25.1 or later",
		"A modern web browser with WebAssembly support",
		"curl, tar, and jq (for installation script)",
		"Basic knowledge of Go programming",
	}

	var listItems []application.BaseWidget
	for _, item := range items {
		listItems = append(listItems, listItem(item))
	}

	return column.New(
		listItems,
		column.Gap(8),
	)
}

func listItem(itemText string) application.BaseWidget {
	return row.New(
		[]application.BaseWidget{
			icon.New(
				icondata.Check,
				icon.Width(breakpoint.All(16)),
				icon.Height(breakpoint.All(16)),
				icon.Fill("#10B981"),
			),
			spacer.New(spacer.Width(8)),
			text.New(
				itemText,
				text.TextStyle(appTheme.Data().TextTheme.TextStyle.Primary),
				text.FontSize(16),
			),
		},
		row.Gap(8),
		row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
	)
}

func installationNextSteps() application.BaseWidget {
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

	var stepItems []application.BaseWidget
	for _, step := range steps {
		stepItems = append(stepItems, nextStepItem(step.title, step.description, step.href))
	}

	return column.New(
		stepItems,
		column.Gap(12),
	)
}

func nextStepItem(title, description, href string) application.BaseWidget {
	return link.New(
		container.New(
			row.New(
				[]application.BaseWidget{
					column.New(
						[]application.BaseWidget{
							text.New(
								title,
								text.TextStyle(appTheme.Data().TextTheme.TextStyle.Primary),
								text.FontSize(16),
								text.FontWeight("500"),
							),
							text.New(
								description,
								text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
								text.FontSize(14),
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
				row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),
			container.ContainerStyle(appTheme.Data().BoxTheme.ContainerStyle.Primary),
			container.Padding(breakpoint.All(spacing.All(16))),
			container.BorderRadius(8),
			container.BorderWidth(spacing.All(1)),
			container.BorderStyle(theme.BorderStyleTypeSolid),
		),
		link.Href(href),
	)
}

func navigationButtons(previousHref, nextHref string) application.BaseWidget {
	return row.New(
		[]application.BaseWidget{
			link.New(
				button.New(
					row.New(
						[]application.BaseWidget{
							icon.New(
								icondata.ChevronLeft,
								icon.Width(breakpoint.All(16)),
								icon.Height(breakpoint.All(16)),
								icon.Fill("#FFFFFF"),
							),
							text.New(
								"Previous",
								text.TextStyle(appTheme.Data().ButtonTheme.ButtonStyle.Primary.TextStyle),
								text.FontSize(14),
								text.FontWeight("500"),
							),
						},
						row.Gap(8),
						row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
					),
					button.Width(breakpoint.All(120)),
				),
				link.Href(previousHref),
			),
			spacer.New(),
			link.New(
				button.New(
					row.New(
						[]application.BaseWidget{
							text.New(
								"Next",
								text.TextStyle(appTheme.Data().ButtonTheme.ButtonStyle.Primary.TextStyle),
								text.FontSize(14),
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
						row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
					),
					button.Width(breakpoint.All(120)),
				),
				link.Href(nextHref),
			),
		},
		row.Flex(1),
	)
}
