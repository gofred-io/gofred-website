package getting_started

import (
	"github.com/gofred-io/gofred-website/app/components/codeblock"
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
)

func ProjectStructureContent() application.BaseWidget {
	return container.New(
		column.New(
			[]application.BaseWidget{
				projectStructurePageHeader(),
				spacer.New(spacer.Height(24)),
				projectStructurePageContent(),
			},
			column.Gap(16),
			column.Flex(1),
		),
		container.Flex(1),
		container.Padding(breakpoint.All(spacing.All(32))),
	)
}

func projectStructurePageHeader() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			text.New(
				"Project Structure",
				text.FontSize(32),
				text.FontWeight("700"),
			),
			text.New(
				"Understand the recommended project structure for gofred applications.",
				text.FontSize(18),
				text.FontColor("#6B7280"),
			),
		},
		column.Gap(8),
	)
}

func projectStructurePageContent() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			contentSection("Directory Structure", "A typical gofred project follows this structure:"),
			codeblock.New(`my-gofred-app/
├── app/                   # Application code
│   ├── components/        # Reusable components
│   │   └── code_block/    # Code block component
│   ├── pages/             # Page components
│   │   ├── 404/           # 404 error page
│   │   ├── docs/          # Documentation pages
│   │   └── home/          # Home page
│   ├── theme/             # Theme and styling
│   │   └── theme.go       # Theme configuration
│   └── app.go             # Main application setup
├── web/                   # Web assets
│   ├── assets/            # Static assets
│   │   ├── fonts/         # Font files
│   │   ├── icons/         # Icon files
│   │   └── images/        # Image files
│   ├── index.css          # CSS styles
│   └── index.html         # HTML template
├── go.mod                 # Go module file
├── go.sum                 # Go module checksums
└── main.go                # Application entry point`),
			spacer.New(spacer.Height(24)),
			contentSection("Key Files", "Here's what each important file does:"),
			projectStructureFileList(),
			spacer.New(spacer.Height(24)),
			contentSection("Best Practices", "Follow these guidelines for better organization:"),
			projectStructureBestPractices(),
			spacer.New(spacer.Height(24)),
			contentSection("Next Steps", "Now that you have learned about the project structure, you can:"),
			projectStructureNextStepsList(),
			spacer.New(spacer.Height(32)),
			navigationButtons("/docs/first-app", "/docs/widgets"),
		},
		column.Gap(16),
	)
}

func projectStructureFileList() application.BaseWidget {
	files := []struct {
		path        string
		description string
	}{
		{
			path:        "main.go",
			description: "The entry point of your application. Contains the main function and application initialization.",
		},
		{
			path:        "app/app.go",
			description: "Main application setup, routing configuration, and global state management.",
		},
		{
			path:        "app/pages/",
			description: "Directory containing all page components. Each page should be in its own subdirectory (home, docs, 404, etc.).",
		},
		{
			path:        "app/components/",
			description: "Reusable UI components that can be used across multiple pages (code_block, etc.).",
		},
		{
			path:        "app/theme/",
			description: "Theme configuration, colors, typography, and global styling options.",
		},
		{
			path:        "web/index.html",
			description: "HTML template file that serves as the base structure for the web application.",
		},
		{
			path:        "web/index.css",
			description: "Main CSS file containing custom styles and overrides for the application.",
		},
		{
			path:        "web/assets/",
			description: "Directory containing static assets like images, fonts, and icons used throughout the application.",
		},
	}

	var fileItems []application.BaseWidget
	for _, file := range files {
		fileItems = append(fileItems, projectStructureFileItem(file.path, file.description))
	}

	return column.New(
		fileItems,
		column.Gap(12),
	)
}

func projectStructureFileItem(path, description string) application.BaseWidget {
	return container.New(
		row.New(
			[]application.BaseWidget{
				icon.New(
					icondata.FileDocument,
					icon.Width(breakpoint.All(20)),
					icon.Height(breakpoint.All(20)),
					icon.Fill("#6B7280"),
				),
				spacer.New(spacer.Width(12)),
				column.New(
					[]application.BaseWidget{
						text.New(
							path,
							text.FontSize(16),
							text.FontWeight("600"),
						),
						text.New(
							description,
							text.FontSize(14),
							text.FontColor("#6B7280"),
							text.LineHeight(1.5),
						),
					},
					column.Gap(4),
					column.Flex(1),
				),
			},
			row.Gap(12),
			row.CrossAxisAlignment(theme.AxisAlignmentTypeStart),
		),
		container.Padding(breakpoint.All(spacing.All(16))),
		container.BackgroundColor("#FFFFFF"),
		container.BorderRadius(8),
		container.BorderColor("#E5E7EB"),
		container.BorderWidth(1, 1, 1, 1),
		container.BorderStyle(theme.BorderStyleTypeSolid),
	)
}

func projectStructureBestPractices() application.BaseWidget {
	practices := []string{
		"Keep pages in separate directories under app/pages/",
		"Create reusable components in app/components/",
		"Use consistent naming conventions (snake_case for files)",
		"Store HTML templates and CSS files in the web/ directory",
		"Organize static assets (images, fonts, icons) in web/assets/",
		"Keep theme configuration centralized in app/theme/",
		"Use meaningful directory and file names",
	}

	var practiceItems []application.BaseWidget
	for _, practice := range practices {
		practiceItems = append(practiceItems, listItem(practice))
	}

	return column.New(
		practiceItems,
		column.Gap(8),
	)
}

func projectStructureNextStepsList() application.BaseWidget {
	steps := []struct {
		title       string
		description string
		href        string
	}{
		{
			title:       "Learn About Widgets",
			description: "Understand the building blocks of gofred applications",
			href:        "/docs/widgets",
		},
		{
			title:       "Explore Layouts",
			description: "Learn how to arrange widgets with columns, rows, and grids",
			href:        "/docs/layouts",
		},
		{
			title:       "Style Your App",
			description: "Make your application beautiful with colors, fonts, and spacing",
			href:        "/docs/styling",
		},
		{
			title:       "Learn State Management",
			description: "Handle dynamic data and user interactions properly",
			href:        "/docs/state",
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
