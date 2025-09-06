package getting_started

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/foundation/code_block"
	. "github.com/gofred-io/gofred/foundation/column"
	. "github.com/gofred-io/gofred/foundation/container"
	. "github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	. "github.com/gofred-io/gofred/foundation/row"
	. "github.com/gofred-io/gofred/foundation/spacer"
	. "github.com/gofred-io/gofred/foundation/text"
	. "github.com/gofred-io/gofred/options"
	. "github.com/gofred-io/gofred/options/spacing"
	. "github.com/gofred-io/gofred/widget"
)

func ProjectStructureContent() Widget {
	return Container(
		Column(
			[]Widget{
				projectStructurePageHeader(),
				Spacer(),
				projectStructurePageContent(),
			},
		).Gap(16).
			Flex(1),
	).Flex(1).
		Padding(AllBP(All(32)))
}

func projectStructurePageHeader() Widget {
	return Column(
		[]Widget{
			Text("Project Structure").
				FontSize(32).
				FontColor("#1F2937").
				FontWeight("700"),
			Text("Understand the recommended project structure for gofred applications.").
				FontSize(18).
				FontColor("#6B7280").
				FontWeight("400"),
		},
	).Gap(8)
}

func projectStructurePageContent() Widget {
	return Column(
		[]Widget{
			contentSection("Directory Structure", "A typical gofred project follows this structure:"),
			CodeBlock(`my-gofred-app/
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
			Spacer(),
			contentSection("Key Files", "Here's what each important file does:"),
			projectStructureFileList(),
			Spacer(),
			contentSection("Best Practices", "Follow these guidelines for better organization:"),
			projectStructureBestPractices(),
			Spacer(),
			contentSection("Next Steps", "Now that you have learned about the project structure, you can:"),
			projectStructureNextStepsList(),
			Spacer(),
			navigationButtons("/docs/first-app", "/docs/widgets"),
		},
	).Gap(16)
}

func projectStructureFileList() Widget {
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

	var fileItems []Widget
	for _, file := range files {
		fileItems = append(fileItems, projectStructureFileItem(file.path, file.description))
	}

	return Column(
		fileItems,
	).Gap(12)
}

func projectStructureFileItem(path, description string) Widget {
	return Container(
		Row(
			[]Widget{
				Icon(icondata.FileDocument).
					Width(AllBP(20)).
					Height(AllBP(20)).
					Fill("#6B7280"),
				Spacer(),
				Column(
					[]Widget{
						Text(path).
							FontSize(16).
							FontColor("#1F2937").
							FontWeight("600"),
						Text(description).
							FontSize(14).
							FontColor("#6B7280").
							FontWeight("400").
							LineHeight(1.5),
					},
				).Gap(4).
					Flex(1),
			},
		).Gap(12).
			CrossAxisAlignment(AxisAlignmentTypeStart),
	).Padding(AllBP(All(16))).
		BackgroundColor("#FFFFFF").
		BorderRadius(8).
		BorderColor("#E5E7EB").
		BorderWidth(1, 1, 1, 1).
		BorderStyle(BorderStyleTypeSolid)
}

func projectStructureBestPractices() Widget {
	practices := []string{
		"Keep pages in separate directories under app/pages/",
		"Create reusable components in app/components/",
		"Use consistent naming conventions (snake_case for files)",
		"Store HTML templates and CSS files in the web/ directory",
		"Organize static assets (images, fonts, icons) in web/assets/",
		"Keep theme configuration centralized in app/theme/",
		"Use meaningful directory and file names",
	}

	var practiceItems []Widget
	for _, practice := range practices {
		practiceItems = append(practiceItems, listItem(practice))
	}

	return Column(
		practiceItems,
	).Gap(8)
}

func projectStructureNextStepsList() Widget {
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

	var stepItems []Widget
	for _, step := range steps {
		stepItems = append(stepItems, nextStepItem(step.title, step.description, step.href))
	}

	return Column(
		stepItems,
	).Gap(12)
}
