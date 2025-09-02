package docs

import (
	codeblock "github.com/gofred-io/gofred-website/app/components/code_block"
	"github.com/gofred-io/gofred-website/app/pages/home"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/widget"
)

func firstAppPage() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			home.Header(),
			firstAppContent(),
			home.Footer(),
		},
		column.Flex(1),
	)
}

func firstAppContent() widget.BaseWidget {
	return container.New(
		row.New(
			[]widget.BaseWidget{
				docsSidebar("/docs/first-app"),
				firstAppContentArea(),
			},
			row.Flex(1),
		),
		container.Flex(1),
		container.BackgroundColor("#F8F9FA"),
	)
}

func firstAppContentArea() widget.BaseWidget {
	return container.New(
		column.New(
			[]widget.BaseWidget{
				docsMobileMenuButton(),
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

func firstAppPageHeader() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			text.New(
				"Your First App",
				text.FontSize(32),
				text.FontColor("#1F2937"),
				text.FontWeight("700"),
			),
			text.New(
				"Build a complete gofred application from scratch with step-by-step instructions.",
				text.FontSize(18),
				text.FontColor("#6B7280"),
				text.FontWeight("400"),
			),
		},
		column.Gap(8),
	)
}

func firstAppPageContent() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
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
    "github.com/gofred-io/gofred/widget"
)

func main() {
    app := createApp()
    application.Run(app)
}

func createApp() widget.BaseWidget {
    return container.New(
        column.New(
            []widget.BaseWidget{
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
    "github.com/gofred-io/gofred/widget"
)

var (
    count, setCount = hooks.UseState(0)
)

func main() {
    app := createApp()
    application.Run(app)
}

func createApp() widget.BaseWidget {
    return container.New(
        column.New(
            []widget.BaseWidget{
                text.New(
                    "Counter App",
                    text.FontSize(24),
                    text.FontColor("#1F2937"),
                    text.FontWeight("700"),
                ),
                listenable.Builder(count, func() widget.BaseWidget {
                    return text.New(
                        fmt.Sprintf("Count: %d", count.Value()),
                        text.FontSize(18),
                        text.FontColor("#2B799B"),
                        text.FontWeight("600"),
                    )
                }),
                spacer.New(spacer.Height(16)),
                row.New(
                    []widget.BaseWidget{
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

func increaseCount(this widget.BaseWidget, e widget.Event) {
    setCount(count.Value() + 1)
}

func decreaseCount(this widget.BaseWidget, e widget.Event) {
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
			navigationButtons("/docs/quick-start", "/docs/widgets"),
		},
		column.Gap(16),
	)
}

func firstAppNextStepsList() widget.BaseWidget {
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
			title:       "State Management",
			description: "Handle dynamic data and user interactions properly",
			href:        "/docs/state",
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
