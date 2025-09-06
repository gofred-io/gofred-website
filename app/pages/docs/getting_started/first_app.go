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

func FirstAppContent() Widget {
	return Container(
		Column(
			[]Widget{
				firstAppPageHeader(),
				Spacer(),
				firstAppPageContent(),
			},
		).Gap(16).
			Flex(1),
	).Flex(1).
		Padding(AllBP(All(32)))
}

func firstAppPageHeader() Widget {
	return Column(
		[]Widget{
			Text("Your First App").
				FontSize(32).
				FontColor("#1F2937").
				FontWeight("700"),
			Text("Build a complete gofred application from scratch with step-by-step instructions.").
				FontSize(18).
				FontColor("#6B7280").
				FontWeight("400"),
		},
	).Gap(8)
}

func firstAppPageContent() Widget {
	return Column(
		[]Widget{
			contentSection("Project Setup", "Let's create a new gofred project and build a simple todo application:"),
			CodeBlock(`mkdir my-gofred-app
cd my-gofred-app
go mod init my-gofred-app
go get github.com/gofred-io/gofred`),
			Spacer(),
			contentSection("Basic Structure", "Create the main application file:"),
			CodeBlock(`package main

import (
    "github.com/gofred-io/gofred/application"
    "github.com/gofred-io/gofred/breakpoint"
    "github.com/gofred-io/gofred/foundation/column"
    "github.com/gofred-io/gofred/foundation/container"
    "github.com/gofred-io/gofred/foundation/text"
    "github.com/gofred-io/gofred/options/spacing"
    "github.com/gofred-io/gofred/widget"
)

func main() {
    app := createApp()
    application.Run(app)
}

func createApp() widget.Widget {
    return container.New(
        column.New(
            []widget.Widget{
                text.New(
                    "My First gofred App",
                    text.FontSize(24),
                    text.FontColor("#1F2937"),
                    text.FontWeight("700"),
                ),
                text.New(
                    "Welcome to gofred! This is your first application.",
                    text.FontSize(16),
                    text.FontColor("#6B7280"),
                    text.FontWeight("400"),
                ),
            },
            column.Gap(16),
        ),
        container.Padding(breakpoint.All(spacing.All(32))),
        container.BackgroundColor("#FFFFFF"),
    )
}
`),
			Spacer(),
			contentSection("Adding Interactivity", "Let's add a simple counter with buttons:"),
			CodeBlock(`package main

import (
    "fmt"

    "github.com/gofred-io/gofred/application"
    "github.com/gofred-io/gofred/breakpoint"
    "github.com/gofred-io/gofred/foundation/button"
    "github.com/gofred-io/gofred/foundation/column"
    "github.com/gofred-io/gofred/foundation/container"
    "github.com/gofred-io/gofred/foundation/row"
    "github.com/gofred-io/gofred/foundation/spacer"
    "github.com/gofred-io/gofred/foundation/text"
    "github.com/gofred-io/gofred/hooks"
    "github.com/gofred-io/gofred/listenable"
    "github.com/gofred-io/gofred/options/spacing"
    "github.com/gofred-io/gofred/widget"
)

var (
    count, setCount = hooks.UseState(0)
)

func main() {
    app := createApp()
    application.Run(app)
}

func createApp() widget.Widget {
    return container.New(
        column.New(
            []widget.Widget{
                text.New(
                    "Counter App",
                    text.FontSize(24),
                    text.FontColor("#1F2937"),
                    text.FontWeight("700"),
                ),
                listenable.Builder(count, func() widget.Widget {
                    return text.New(
                        fmt.Sprintf("Count: %d", count.Value()),
                        text.FontSize(18),
                        text.FontColor("#2B799B"),
                        text.FontWeight("600"),
                    )
                }),
                spacer.New(spacer.Height(16)),
                row.New(
                    []widget.Widget{
                        button.New(
                            text.New("Decrease", text.FontColor("#FFFFFF")),
                            button.BackgroundColor("#EF4444"),
                            button.OnClick(decreaseCount),
                        ),
                        spacer.New(spacer.Width(16)),
                        button.New(
                            text.New("Increase", text.FontColor("#FFFFFF")),
                            button.BackgroundColor("#10B981"),
                            button.OnClick(increaseCount),
                        ),
                    },
                    row.Gap(16),
                ),
            },
            column.Gap(16),
        ),
        container.Padding(breakpoint.All(spacing.All(32))),
        container.BackgroundColor("#FFFFFF"),
    )
}

func increaseCount(this widget.Widget, e widget.Event) {
    setCount(count.Value() + 1)
}

func decreaseCount(this widget.Widget, e widget.Event) {
    setCount(count.Value() - 1)
}`),
			Spacer(),
			contentSection("Running Your App", "To run your application:"),
			CodeBlock(`go run server/server.go`),
			Spacer(),
			Text("Your app will compile to WebAssembly and run in your browser automatically!").
				FontSize(16).
				FontColor("#10B981").
				FontWeight("500"),
			Spacer(),
			contentSection("What's Next?", "Now that you've built your first app, explore these topics:"),
			firstAppNextStepsList(),
			Spacer(),
			navigationButtons("/docs/quick-start", "/docs/project-structure"),
		},
	).Gap(16)
}

func firstAppNextStepsList() Widget {
	steps := []struct {
		title       string
		description string
		href        string
	}{
		{
			title:       "Project Structure",
			description: "Explore the project structure",
			href:        "/docs/project-structure",
		},
		// {
		// 	title:       "Learn About Widgets",
		// 	description: "Understand the building blocks of gofred applications",
		// 	href:        "/docs/widgets",
		// },
		// {
		// 	title:       "Explore Layouts",
		// 	description: "Learn how to arrange widgets with columns, rows, and grids",
		// 	href:        "/docs/layouts",
		// },
		// {
		// 	title:       "Style Your App",
		// 	description: "Make your application beautiful with colors, fonts, and spacing",
		// 	href:        "/docs/styling",
		// },
		// {
		// 	title:       "State Management",
		// 	description: "Handle dynamic data and user interactions properly",
		// 	href:        "/docs/state",
		// },
	}

	var stepItems []Widget
	for _, step := range steps {
		stepItems = append(stepItems, nextStepItem(step.title, step.description, step.href))
	}

	return Column(
		stepItems,
	).Gap(12)
}
