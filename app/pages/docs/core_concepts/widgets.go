package core_concepts

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/foundation/button"
	. "github.com/gofred-io/gofred/foundation/code_block"
	. "github.com/gofred-io/gofred/foundation/column"
	. "github.com/gofred-io/gofred/foundation/container"
	. "github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	. "github.com/gofred-io/gofred/foundation/link"
	. "github.com/gofred-io/gofred/foundation/row"
	. "github.com/gofred-io/gofred/foundation/spacer"
	. "github.com/gofred-io/gofred/foundation/text"
	. "github.com/gofred-io/gofred/options"
	. "github.com/gofred-io/gofred/options/spacing"
	. "github.com/gofred-io/gofred/widget"
)

func WidgetsContent() Widget {
	return Container(
		Column(
			[]Widget{
				widgetsPageHeader(),
				Spacer(),
				widgetsPageContent(),
			},
		).Gap(16).
			Flex(1),
	).Flex(1).
		Padding(AllBP(All(32)))
}

func widgetsPageHeader() Widget {
	return Column(
		[]Widget{
			Text("Foundation Widgets").
				FontSize(32).
				FontColor("#1F2937").
				FontWeight("700"),
			Text("Learn about the core building blocks of gofred applications and how to use them effectively.").
				FontSize(18).
				FontColor("#6B7280").
				FontWeight("400"),
		},
	).Gap(8)
}

func widgetsPageContent() Widget {
	return Column(
		[]Widget{
			contentSection("What are Widgets?", "Widgets are the fundamental building blocks of gofred applications. They are composable UI components that can be combined to create complex user interfaces. Every element you see in a gofred app is a widget."),
			Spacer(),

			// Layout Widgets
			contentSection("Layout Widgets", "Layout widgets help you organize and position other widgets in your application."),
			Spacer(),

			widgetSubsection("Container", "A flexible container widget that can hold other widgets and apply styling properties like padding, background color, borders, and sizing."),
			CodeBlock(`container.New(
    text.New("Hello, World!"),
    container.Padding(breakpoint.All(spacing.All(16))),
    container.BackgroundColor("#F3F4F6"),
    container.BorderRadius(8),
    container.BorderColor("#E5E7EB"),
    container.BorderWidth(1, 1, 1, 1),
)`),
			Spacer(),

			widgetSubsection("Column", "Arranges child widgets vertically. Perfect for creating vertical layouts and forms."),
			CodeBlock(`column.New(
    []widget.Widget{
        text.New("First Item"),
        text.New("Second Item"),
        text.New("Third Item"),
    },
    column.Gap(16),
    column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
)`),
			Spacer(),

			widgetSubsection("Row", "Arranges child widgets horizontally. Ideal for creating horizontal layouts and toolbars."),
			CodeBlock(`row.New(
    []widget.Widget{
        button.New(text.New("Cancel")),
        spacer.New(), // Pushes buttons apart
        button.New(text.New("Submit")),
    },
    row.Gap(12),
    row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
)`),
			Spacer(),

			widgetSubsection("Grid", "Creates a responsive grid layout for organizing widgets in rows and columns."),
			CodeBlock(`grid.New(
    []widget.Widget{
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
			Spacer(),

			widgetSubsection("Center", "Centers its child widget both horizontally and vertically."),
			CodeBlock(`center.New(
    text.New(
        "Centered Content",
        text.FontSize(24),
        text.FontWeight("600"),
    ),
)`),
			Spacer(),

			widgetSubsection("Spacer", "Creates flexible or fixed spacing between widgets."),
			CodeBlock(`// Fixed spacing
spacer.New(spacer.Height(24))
spacer.New(spacer.Width(16))

// Flexible spacing (takes available space)
spacer.New()`),
			Spacer(),

			// Content Widgets
			contentSection("Content Widgets", "Content widgets display information and media to users."),
			Spacer(),

			widgetSubsection("Text", "Displays text with customizable styling options like font size, color, weight, and alignment."),
			CodeBlock(`text.New(
    "Hello, gofred!",
    text.FontSize(18),
    text.FontColor("#1F2937"),
    text.FontWeight("600"),
    text.TextAlign(options.TextAlignTypeCenter),
    text.LineHeight(1.5),
)`),
			Spacer(),

			widgetSubsection("Icon", "Displays scalable vector icons from the built-in icon library."),
			CodeBlock(`icon.New(
    icondata.Home,
    icon.Width(breakpoint.All(24)),
    icon.Height(breakpoint.All(24)),
    icon.Fill("#2B799B"),
)`),
			Spacer(),

			widgetSubsection("Image", "Displays images with support for various formats and responsive sizing."),
			CodeBlock(`image.New(
    image.Src("/assets/logo.png"),
    image.Alt("Company Logo"),
    image.Width(breakpoint.All(200)),
    image.Height(breakpoint.All(100)),
    image.ObjectFit(options.ObjectFitTypeCover),
)`),
			Spacer(),

			// Interactive Widgets
			contentSection("Interactive Widgets", "Interactive widgets respond to user input and enable user interactions."),
			Spacer(),

			widgetSubsection("Button", "A clickable button widget that can trigger actions and navigate between screens."),
			CodeBlock(`button.New(
    text.New("Click Me", text.FontColor("#FFFFFF")),
    button.BackgroundColor("#2B799B"),
    button.BorderRadius(8),
    button.Padding(breakpoint.All(spacing.Symmetric(16, 12))),
    button.OnClick(handleButtonClick),
)

func handleButtonClick(this widget.Widget, e widget.Event) {
    // Handle button click
    fmt.Println("Button clicked!")
}`),
			Spacer(),

			widgetSubsection("Icon Button", "A button that contains only an icon, perfect for toolbars and action menus."),
			CodeBlock(`iconbutton.New(
    icondata.Settings,
    iconbutton.IconWidth(breakpoint.All(20)),
    iconbutton.IconHeight(breakpoint.All(20)),
    iconbutton.IconFill("#6B7280"),
    iconbutton.BackgroundColor("#F9FAFB"),
    iconbutton.BorderRadius(6),
    iconbutton.OnClick(openSettings),
)`),
			Spacer(),

			widgetSubsection("Link", "Creates navigational links that can route to different pages or external URLs."),
			CodeBlock(`link.New(
    text.New(
        "Go to Documentation",
        text.FontColor("#2B799B"),
        text.TextDecoration(options.TextDecorationTypeUnderline),
    ),
    link.Href("/docs"),
)`),
			Spacer(),

			// Navigation Widgets
			contentSection("Navigation Widgets", "Navigation widgets help users move through your application."),
			Spacer(),

			widgetSubsection("Drawer", "A slide-out panel that can contain navigation menus or additional content."),
			CodeBlock(`drawer.New(
    drawerContent(), // Your drawer content
    drawer.Width(breakpoint.All(300)),
    drawer.BackgroundColor("#FFFFFF"),
    drawer.BorderColor("#E5E7EB"),
    drawer.BorderWidth(0, 1, 0, 0),
)`),
			Spacer(),

			widgetSubsection("Router", "Manages navigation and routing in your application."),
			CodeBlock(`router.New(
    router.Routes([]router.Route{
        {Path: "/", Handler: homePage},
        {Path: "/about", Handler: aboutPage},
        {Path: "/contact", Handler: contactPage},
    }),
)`),
			Spacer(),

			// Widget Composition
			contentSection("Widget Composition", "Widgets can be composed together to create complex UI components. Here's an example of building a card component:"),
			CodeBlock(`func cardWidget(title, content string) widget.Widget {
    return container.New(
        column.New(
            []widget.Widget{
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
                    []widget.Widget{
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
			Spacer(),

			// Best Practices
			contentSection("Best Practices", "Follow these guidelines when working with widgets:"),
			bestPracticesList(),
			Spacer(),

			contentSection("What's Next?", "Now that you understand widgets, explore these related topics:"),
			widgetsNextStepsList(),
			Spacer(),
			navigationButtons("/docs/project-structure", "/docs/layouts"),
		},
	).Gap(16)
}

func contentSection(title, description string) Widget {
	return Column(
		[]Widget{
			Text(title).
				FontSize(24).
				FontColor("#1F2937").
				FontWeight("600"),
			Text(description).
				FontSize(16).
				FontColor("#6B7280").
				FontWeight("400"),
		},
	).Gap(8)
}

func widgetSubsection(title, description string) Widget {
	return Column(
		[]Widget{
			Text(title).
				FontSize(20).
				FontColor("#1F2937").
				FontWeight("600"),
			Text(description).
				FontSize(14).
				FontColor("#6B7280").
				FontWeight("400"),
		},
	).Gap(4)
}

func bestPracticesList() Widget {
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

	var practiceItems []Widget
	for _, practice := range practices {
		practiceItems = append(practiceItems, listItem(practice))
	}

	return Column(
		practiceItems,
	).Gap(8)
}

func listItem(itemText string) Widget {
	return Row(
		[]Widget{
			Icon(icondata.Check).
				Width(AllBP(16)).
				Height(AllBP(16)).
				Fill("#10B981"),
			Spacer(),
			Text(itemText).
				FontSize(16).
				FontColor("#374151").
				FontWeight("400"),
		},
	).Gap(8).
		CrossAxisAlignment(AxisAlignmentTypeCenter)
}

func widgetsNextStepsList() Widget {
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

	var stepItems []Widget
	for _, step := range steps {
		stepItems = append(stepItems, nextStepItem(step.title, step.description, step.href))
	}

	return Column(
		stepItems,
	).Gap(12)
}

func nextStepItem(title, description, href string) Widget {
	return Link(
		Container(
			Row(
				[]Widget{
					Column(
						[]Widget{
							Text(title).
								FontSize(16).
								FontColor("#2B799B").
								FontWeight("500"),
							Text(description).
								FontSize(14).
								FontColor("#6B7280").
								FontWeight("400"),
						},
					).Gap(4).
						Flex(1),
					Icon(icondata.ChevronRight).
						Width(AllBP(20)).
						Height(AllBP(20)).
						Fill("#9CA3AF"),
				},
			).Gap(12).
				Flex(1).
				CrossAxisAlignment(AxisAlignmentTypeCenter),
		).Padding(AllBP(All(16))).
			BackgroundColor("#FFFFFF").
			BorderRadius(8).
			BorderColor("#E5E7EB").
			BorderWidth(1, 1, 1, 1).
			BorderStyle(BorderStyleTypeSolid),
	).Href(href)
}

func navigationButtons(previousHref, nextHref string) Widget {
	return Row(
		[]Widget{
			Link(
				Button(
					Row(
						[]Widget{
							Icon(icondata.ChevronLeft).
								Width(AllBP(16)).
								Height(AllBP(16)).
								Fill("#FFFFFF"),
							Text("Previous").
								FontSize(14).
								FontColor("#FFFFFF").
								FontWeight("500"),
						},
					).Gap(8).
						CrossAxisAlignment(AxisAlignmentTypeCenter),
				).BackgroundColor("#6B7280").
					Width(AllBP(120)),
			).Href(previousHref),
			Spacer(),
			Link(
				Button(
					Row(
						[]Widget{
							Text("Next").
								FontSize(14).
								FontColor("#FFFFFF").
								FontWeight("500"),
							Icon(icondata.ChevronRight).
								Width(AllBP(16)).
								Height(AllBP(16)).
								Fill("#FFFFFF"),
						},
					).Gap(8).
						CrossAxisAlignment(AxisAlignmentTypeCenter),
				).BackgroundColor("#2B799B").
					Width(AllBP(120)),
			).Href(nextHref),
		},
	).Flex(1)
}
