package getting_started

import (
	"github.com/gofred-io/gofred/breakpoint"
	codeblock "github.com/gofred-io/gofred/foundation/code_block"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/widget"
)

func QuickStartContent() widget.BaseWidget {
	return container.New(
		column.New(
			[]widget.BaseWidget{
				quickStartPageHeader(),
				spacer.New(spacer.Height(24)),
				quickStartPageContent(),
			},
			column.Gap(16),
			column.Flex(1),
		),
		container.Flex(1),
		container.Padding(breakpoint.All(spacing.All(32))),
	)
}

func quickStartPageHeader() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			text.New(
				"Quick Start",
				text.FontSize(32),
				text.FontColor("#1F2937"),
				text.FontWeight("700"),
			),
			text.New(
				"Get started with gofred by creating your first application.",
				text.FontSize(18),
				text.FontColor("#6B7280"),
				text.FontWeight("400"),
			),
		},
		column.Gap(8),
	)
}

func quickStartPageContent() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			contentSection("Hello, gofred!", "Let's start with a simple hello world application:"),
			codeblock.New(`package main

import (
    "github.com/gofred-io/gofred/application"
    "github.com/gofred-io/gofred/foundation/text"
)

func main() {
    app := text.New("Hello, gofred!")
    application.Run(app)
}`),
			spacer.New(spacer.Height(24)),
			contentSection("Next Steps", "Now that you have gofred installed, you can:"),
			quickStartNextStepsList(),
			spacer.New(spacer.Height(32)),
			navigationButtons("/docs/installation", "/docs/first-app"),
		},
		column.Gap(16),
	)
}

func quickStartNextStepsList() widget.BaseWidget {
	steps := []struct {
		title       string
		description string
		href        string
	}{
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
