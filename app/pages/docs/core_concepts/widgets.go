package core_concepts

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/button"
	codeblock "github.com/gofred-io/gofred/foundation/code_block"
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

func WidgetsContent() widget.BaseWidget {
	return container.New(
		column.New(
			[]widget.BaseWidget{
				widgetsPageHeader(),
				spacer.New(spacer.Height(24)),
				widgetsPageContent(),
			},
			column.Gap(16),
			column.Flex(1),
		),
		container.Flex(1),
		container.Padding(breakpoint.All(spacing.All(32))),
	)
}

func widgetsPageHeader() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			text.New(
				"Foundation Widgets",
				text.FontSize(32),
				text.FontColor("#1F2937"),
				text.FontWeight("700"),
			),
			text.New(
				"Learn about the core building blocks of gofred applications and how to use them effectively.",
				text.FontSize(18),
				text.FontColor("#6B7280"),
				text.FontWeight("400"),
			),
		},
		column.Gap(8),
	)
}

func widgetsPageContent() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			contentSection("What are Widgets?", "Widgets are the fundamental building blocks of gofred applications. They are composable UI components that can be combined to create complex user interfaces. Every element you see in a gofred app is a widget."),
			spacer.New(spacer.Height(24)),

			// Layout Widgets
			contentSection("Layout Widgets", "Layout widgets help you organize and position other widgets in your application."),
			spacer.New(spacer.Height(16)),

			widgetSubsection("Container", "A flexible container widget that can hold other widgets and apply styling properties like padding, background color, borders, and sizing."),
			codeblock.New(`container.New(
    text.New("Hello, World!"),
    container.Padding(breakpoint.All(spacing.All(16))),
    container.BackgroundColor("#F3F4F6"),
    container.BorderRadius(8),
    container.BorderColor("#E5E7EB"),
    container.BorderWidth(1, 1, 1, 1),
)`),
			spacer.New(spacer.Height(16)),

			widgetSubsection("Column", "Arranges child widgets vertically. Perfect for creating vertical layouts and forms."),
			codeblock.New(`column.New(
    []widget.BaseWidget{
        text.New("First Item"),
        text.New("Second Item"),
        text.New("Third Item"),
    },
    column.Gap(16),
    column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
)`),
			spacer.New(spacer.Height(16)),

			widgetSubsection("Row", "Arranges child widgets horizontally. Ideal for creating horizontal layouts and toolbars."),
			codeblock.New(`row.New(
    []widget.BaseWidget{
        button.New(text.New("Cancel")),
        spacer.New(), // Pushes buttons apart
        button.New(text.New("Submit")),
    },
    row.Gap(12),
    row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
)`),
			spacer.New(spacer.Height(16)),

			widgetSubsection("Grid", "Creates a responsive grid layout for organizing widgets in rows and columns."),
			codeblock.New(`grid.New(
    []widget.BaseWidget{
        cardWidget("Card 1"),
        cardWidget("Card 2"),
        cardWidget("Card 3"),
        cardWidget("Card 4"),
    },
    grid.ColumnCount(
        breakpoint.XS(1),
        breakpoint.MD(2),
        breakpoint.LG(4),
    ),
    grid.ColumnGap(16),
    grid.RowGap(16),
)`),
			spacer.New(spacer.Height(16)),

			widgetSubsection("Center", "Centers its child widget both horizontally and vertically."),
			codeblock.New(`center.New(
    text.New(
        "Centered Content",
        text.FontSize(24),
        text.FontWeight("600"),
    ),
)`),
			spacer.New(spacer.Height(16)),

			widgetSubsection("Spacer", "Creates flexible or fixed spacing between widgets."),
			codeblock.New(`// Fixed spacing
spacer.New(spacer.Height(24))
spacer.New(spacer.Width(16))

// Flexible spacing (takes available space)
spacer.New()`),
			spacer.New(spacer.Height(24)),

			// Content Widgets
			contentSection("Content Widgets", "Content widgets display information and media to users."),
			spacer.New(spacer.Height(16)),

			widgetSubsection("Text", "Displays text with customizable styling options like font size, color, weight, and alignment."),
			codeblock.New(`text.New(
    "Hello, gofred!",
    text.FontSize(18),
    text.FontColor("#1F2937"),
    text.FontWeight("600"),
    text.TextAlign(options.TextAlignTypeCenter),
    text.LineHeight(1.5),
)`),
			spacer.New(spacer.Height(16)),

			widgetSubsection("Icon", "Displays scalable vector icons from the built-in icon library."),
			codeblock.New(`icon.New(
    icondata.Home,
    icon.Width(breakpoint.All(24)),
    icon.Height(breakpoint.All(24)),
    icon.Fill("#2B799B"),
)`),
			spacer.New(spacer.Height(16)),

			widgetSubsection("Image", "Displays images with support for various formats and responsive sizing."),
			codeblock.New(`image.New(
    image.Src("/assets/logo.png"),
    image.Alt("Company Logo"),
    image.Width(breakpoint.All(200)),
    image.Height(breakpoint.All(100)),
    image.ObjectFit(options.ObjectFitTypeCover),
)`),
			spacer.New(spacer.Height(24)),

			// Interactive Widgets
			contentSection("Interactive Widgets", "Interactive widgets respond to user input and enable user interactions."),
			spacer.New(spacer.Height(16)),

			widgetSubsection("Button", "A clickable button widget that can trigger actions and navigate between screens."),
			codeblock.New(`button.New(
    text.New("Click Me", text.FontColor("#FFFFFF")),
    button.BackgroundColor("#2B799B"),
    button.BorderRadius(8),
    button.Padding(breakpoint.All(spacing.Symmetric(16, 12))),
    button.OnClick(handleButtonClick),
)

func handleButtonClick(this widget.BaseWidget, e widget.Event) {
    // Handle button click
    fmt.Println("Button clicked!")
}`),
			spacer.New(spacer.Height(16)),

			widgetSubsection("Icon Button", "A button that contains only an icon, perfect for toolbars and action menus."),
			codeblock.New(`iconbutton.New(
    icondata.Settings,
    iconbutton.IconWidth(breakpoint.All(20)),
    iconbutton.IconHeight(breakpoint.All(20)),
    iconbutton.IconFill("#6B7280"),
    iconbutton.BackgroundColor("#F9FAFB"),
    iconbutton.BorderRadius(6),
    iconbutton.OnClick(openSettings),
)`),
			spacer.New(spacer.Height(16)),

			widgetSubsection("Link", "Creates navigational links that can route to different pages or external URLs."),
			codeblock.New(`link.New(
    text.New(
        "Go to Documentation",
        text.FontColor("#2B799B"),
        text.TextDecoration(options.TextDecorationTypeUnderline),
    ),
    link.Href("/docs"),
)`),
			spacer.New(spacer.Height(24)),

			// Navigation Widgets
			contentSection("Navigation Widgets", "Navigation widgets help users move through your application."),
			spacer.New(spacer.Height(16)),

			widgetSubsection("Drawer", "A slide-out panel that can contain navigation menus or additional content."),
			codeblock.New(`drawer.New(
    drawerContent(), // Your drawer content
    drawer.Width(breakpoint.All(300)),
    drawer.BackgroundColor("#FFFFFF"),
    drawer.BorderColor("#E5E7EB"),
    drawer.BorderWidth(0, 1, 0, 0),
)`),
			spacer.New(spacer.Height(16)),

			widgetSubsection("Router", "Manages navigation and routing in your application."),
			codeblock.New(`router.New(
    router.Routes([]router.Route{
        {Path: "/", Handler: homePage},
        {Path: "/about", Handler: aboutPage},
        {Path: "/contact", Handler: contactPage},
    }),
)`),
			spacer.New(spacer.Height(24)),

			// Widget Composition
			contentSection("Widget Composition", "Widgets can be composed together to create complex UI components. Here's an example of building a card component:"),
			codeblock.New(`func cardWidget(title, content string) widget.BaseWidget {
    return container.New(
        column.New(
            []widget.BaseWidget{
                text.New(
                    title,
                    text.FontSize(18),
                    text.FontWeight("600"),
                    text.FontColor("#1F2937"),
                ),
                spacer.New(spacer.Height(8)),
                text.New(
                    content,
                    text.FontSize(14),
                    text.FontColor("#6B7280"),
                    text.LineHeight(1.5),
                ),
                spacer.New(spacer.Height(16)),
                row.New(
                    []widget.BaseWidget{
                        spacer.New(),
                        button.New(
                            text.New("Learn More", text.FontColor("#2B799B")),
                            button.BackgroundColor("transparent"),
                            button.BorderColor("#2B799B"),
                            button.BorderWidth(1, 1, 1, 1),
                        ),
                    },
                ),
            },
            column.Gap(0),
        ),
        container.Padding(breakpoint.All(spacing.All(16))),
        container.BackgroundColor("#FFFFFF"),
        container.BorderRadius(8),
        container.BorderColor("#E5E7EB"),
        container.BorderWidth(1, 1, 1, 1),
        container.BorderStyle(options.BorderStyleTypeSolid),
    )
}`),
			spacer.New(spacer.Height(24)),

			// Best Practices
			contentSection("Best Practices", "Follow these guidelines when working with widgets:"),
			bestPracticesList(),
			spacer.New(spacer.Height(24)),

			contentSection("What's Next?", "Now that you understand widgets, explore these related topics:"),
			widgetsNextStepsList(),
			spacer.New(spacer.Height(32)),
			navigationButtons("/docs/project-structure", "/docs/layouts"),
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

func widgetSubsection(title, description string) widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			text.New(
				title,
				text.FontSize(20),
				text.FontColor("#1F2937"),
				text.FontWeight("600"),
			),
			text.New(
				description,
				text.FontSize(14),
				text.FontColor("#6B7280"),
				text.FontWeight("400"),
			),
		},
		column.Gap(4),
	)
}

func bestPracticesList() widget.BaseWidget {
	practices := []string{
		"Use semantic widget names that clearly describe their purpose",
		"Keep widget composition simple and avoid deep nesting when possible",
		"Use spacer widgets for consistent spacing instead of margins",
		"Leverage breakpoints for responsive design across different screen sizes",
		"Create reusable widget functions for components used multiple times",
		"Use appropriate layout widgets (column, row, grid) based on your design needs",
		"Consider accessibility when choosing colors and font sizes",
		"Test your widgets across different screen sizes and browsers",
	}

	var practiceItems []widget.BaseWidget
	for _, practice := range practices {
		practiceItems = append(practiceItems, listItem(practice))
	}

	return column.New(
		practiceItems,
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

func widgetsNextStepsList() widget.BaseWidget {
	steps := []struct {
		title       string
		description string
		href        string
	}{
		{
			title:       "Layouts",
			description: "Learn advanced layout techniques and responsive design",
			href:        "/docs/layouts",
		},
		{
			title:       "Styling",
			description: "Discover how to style your widgets with colors, fonts, and spacing",
			href:        "/docs/styling",
		},
		{
			title:       "State Management",
			description: "Handle dynamic data and user interactions with hooks",
			href:        "/docs/state",
		},
		{
			title:       "Events",
			description: "Learn how to handle user interactions and events",
			href:        "/docs/events",
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
