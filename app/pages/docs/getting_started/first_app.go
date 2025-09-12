package getting_started

import (
	"github.com/gofred-io/gofred-website/app/components/codeblock"
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/options/spacing"
)

func FirstAppContent() application.BaseWidget {
	return container.New(
		column.New(
			[]application.BaseWidget{
				firstAppPageHeader(),
				spacer.New(spacer.Height(24)),
				firstAppPageContent(),
			},
			column.Gap(16),
			column.Flex(1),
		),
		container.Flex(1),
		container.Padding(breakpoint.All(spacing.All(32))),
	)
}

func firstAppPageHeader() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			text.New(
				"Your First App",
				text.FontSize(32),
				text.FontWeight("700"),
			),
			text.New(
				"Build a complete gofred application from scratch with step-by-step instructions.",
				text.FontSize(18),
				text.FontColor("#6B7280"),
			),
		},
		column.Gap(8),
	)
}

func firstAppPageContent() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			contentSection("Project Setup", "Let's create a new gofred project and build a simple todo application:"),
			codeblock.New(`mkdir my-gofred-app
cd my-gofred-app
go mod init my-gofred-app
go get github.com/gofred-io/gofred`),
			spacer.New(spacer.Height(24)),
			contentSection("Basic Structure", "Create the main application file:"),
			codeblock.New(`package main

import (
    "github.com/gofred-io/gofred/application"
    "github.com/gofred-io/gofred/breakpoint"
    "github.com/gofred-io/gofred/foundation/column"
    "github.com/gofred-io/gofred/foundation/container"
    "github.com/gofred-io/gofred/foundation/text"
    "github.com/gofred-io/gofred/options/spacing"
    "github.com/gofred-io/gofred/application"
)

func main() {
    app := createApp()
    application.Run(app)
}

func createApp() application.BaseWidget {
    return container.New(
        column.New(
            []application.BaseWidget{
                text.New(
                    "My First gofred App",
                    text.FontSize(24),
                                        text.FontWeight("700"),
                ),
                text.New(
                    "Welcome to gofred! This is your first application.",
                    text.FontSize(16),
                    text.FontColor("#6B7280"),
                                    ),
            },
            column.Gap(16),
        ),
        container.Padding(breakpoint.All(spacing.All(32))),
        container.BackgroundColor("#FFFFFF"),
    )
}
`),
			spacer.New(spacer.Height(24)),
			contentSection("Adding Interactivity", "Let's add a simple counter with buttons:"),
			codeblock.New(`package main

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
    "github.com/gofred-io/gofred/application"
)

var (
    count, setCount = hooks.UseState(0)
)

func main() {
    app := createApp()
    application.Run(app)
}

func createApp() application.BaseWidget {
    return container.New(
        column.New(
            []application.BaseWidget{
                text.New(
                    "Counter App",
                    text.FontSize(24),
                                        text.FontWeight("700"),
                ),
                listenable.Builder(count, func() application.BaseWidget {
                    return text.New(
                        fmt.Sprintf("Count: %d", count.Value()),
                        text.FontSize(18),
                        text.FontColor("#2B799B"),
                        text.FontWeight("600"),
                    )
                }),
                spacer.New(spacer.Height(16)),
                row.New(
                    []application.BaseWidget{
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

func increaseCount(this application.BaseWidget, e application.Event) {
    setCount(count.Value() + 1)
}

func decreaseCount(this application.BaseWidget, e application.Event) {
    setCount(count.Value() - 1)
}`),
			spacer.New(spacer.Height(24)),
			contentSection("Running Your App", "To run your application:"),
			codeblock.New(`go run server/server.go`),
			spacer.New(spacer.Height(16)),
			text.New(
				"Your app will compile to WebAssembly and run in your browser automatically!",
				text.FontSize(16),
				text.FontColor("#10B981"),
				text.FontWeight("500"),
			),
			spacer.New(spacer.Height(24)),
			contentSection("What's Next?", "Now that you've built your first app, explore these topics:"),
			firstAppNextStepsList(),
			spacer.New(spacer.Height(32)),
			navigationButtons("/docs/quick-start", "/docs/project-structure"),
		},
		column.Gap(16),
	)
}

func firstAppNextStepsList() application.BaseWidget {
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

	var stepItems []application.BaseWidget
	for _, step := range steps {
		stepItems = append(stepItems, nextStepItem(step.title, step.description, step.href))
	}

	return column.New(
		stepItems,
		column.Gap(12),
	)
}
