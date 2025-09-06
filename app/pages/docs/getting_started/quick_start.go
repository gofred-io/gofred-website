package getting_started

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/foundation/code_block"
	. "github.com/gofred-io/gofred/foundation/column"
	. "github.com/gofred-io/gofred/foundation/container"
	. "github.com/gofred-io/gofred/foundation/spacer"
	. "github.com/gofred-io/gofred/foundation/text"
	. "github.com/gofred-io/gofred/options/spacing"
	. "github.com/gofred-io/gofred/widget"
)

func QuickStartContent() Widget {
	return Container(
		Column(
			[]Widget{
				quickStartPageHeader(),
				Spacer(),
				quickStartPageContent(),
			},
		).Gap(16).
			Flex(1),
	).Flex(1).
		Padding(AllBP(All(32)))
}

func quickStartPageHeader() Widget {
	return Column(
		[]Widget{
			Text("Quick Start").
				FontSize(32).
				FontColor("#1F2937").
				FontWeight("700"),
			Text("Get started with gofred by creating your first application.").
				FontSize(18).
				FontColor("#6B7280").
				FontWeight("400"),
		},
	).Gap(8)
}

func quickStartPageContent() Widget {
	return Column(
		[]Widget{
			contentSection("Hello, gofred!", "Let's start with a simple hello world application:"),
			CodeBlock(`package main

import (
    "github.com/gofred-io/gofred/application"
    "github.com/gofred-io/gofred/foundation/text"
)

func main() {
    app := text.New("Hello, gofred!")
    application.Run(app)
}`),
			Spacer(),
			contentSection("Next Steps", "Now that you have gofred installed, you can:"),
			quickStartNextStepsList(),
			Spacer(),
			navigationButtons("/docs/installation", "/docs/first-app"),
		},
	).Gap(16)
}

func quickStartNextStepsList() Widget {
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

	var stepItems []Widget
	for _, step := range steps {
		stepItems = append(stepItems, nextStepItem(step.title, step.description, step.href))
	}

	return Column(
		stepItems,
	).Gap(12)
}
