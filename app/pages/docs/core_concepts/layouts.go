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

func LayoutsContent() Widget {
	return Container(
		Column(
			[]Widget{
				layoutsPageHeader(),
				Spacer().Height(24),
				layoutsPageContent(),
			},
		).Gap(16).Flex(1),
	).Flex(1).Padding(AllBP(All(32)))
}

func layoutsPageHeader() Widget {
	return Column(
		[]Widget{
			Text("Layouts & Responsive Design").
				FontSize(32).
				FontColor("#1F2937").
				FontWeight("700"),
			Text("Master the art of creating beautiful, responsive layouts that work perfectly across all devices and screen sizes.").
				FontSize(18).
				FontColor("#6B7280").
				FontWeight("400"),
		},
	).Gap(8)
}

func layoutsPageContent() Widget {
	return Column(
		[]Widget{
			layoutContentSection("Layout Fundamentals", "Layouts in gofred are built using a flexbox-inspired system that makes it easy to create responsive, flexible designs. The main layout widgets work together to help you arrange content exactly how you want it."),
			Spacer().Height(24),

			// Column Layouts
			layoutContentSection("Column Layouts", "Columns arrange widgets vertically, perfect for forms, lists, and content that flows from top to bottom."),
			Spacer().Height(16),

			layoutSubsection("Basic Column", "The foundation of vertical layouts."),
			CodeBlock(`Column(
    []Widget{
        Text("Header", .FontSize(24), .FontWeight("600")),
        Text("Subtitle", .FontSize(16), .FontColor("#6B7280")),
        Text("Content goes here..."),
        Button(Text("Action Button")),
    },
    .Gap(16), // Space between items
)`),
			Spacer().Height(16),

			layoutSubsection("Column Alignment", "Control how items are aligned within the column."),
			CodeBlock(`// Center-aligned column
Column(
    []Widget{
        Text("Centered Title"),
        Text("Centered content"),
        Button(Text("Centered Button")),
    },
    .Gap(12),
    column.CrossAxisAlignment(AxisAlignmentTypeCenter),
)

// Right-aligned column
Column(
    items,
    column.CrossAxisAlignment(AxisAlignmentTypeEnd),
)`),
			Spacer().Height(16),

			layoutSubsection("Flexible Columns", "Make columns expand to fill available space."),
			CodeBlock(`Column(
    []Widget{
        header(),
        Column(
            contentItems,
            .Flex(1), // This column takes remaining space
        ),
        footer(),
    },
    .Gap(0),
)`),
			Spacer().Height(24),

			// Row Layouts
			layoutContentSection("Row Layouts", "Rows arrange widgets horizontally, ideal for navigation bars, button groups, and side-by-side content."),
			Spacer().Height(16),

			layoutSubsection("Basic Row", "The foundation of horizontal layouts."),
			CodeBlock(`Row(
    []Widget{
        Icon(icondata.User, .Width(AllBP(20))),
        Text("John Doe", .FontWeight("500")),
        Spacer(), // Pushes next items to the right
        Text("Online", .FontColor("#10B981")),
        Button(Text("Contact")),
    },
    .Gap(12),
    .CrossAxisAlignment(AxisAlignmentTypeCenter),
)`),
			Spacer().Height(16),

			layoutSubsection("Row Alignment", "Control vertical alignment of items in a row."),
			CodeBlock(`// Top-aligned row
Row(
    []Widget{
        largeImage(),
        Column(smallTextItems),
    },
    .CrossAxisAlignment(AxisAlignmentTypeStart),
)

// Bottom-aligned row
Row(
    items,
    .CrossAxisAlignment(AxisAlignmentTypeEnd),
)

// Stretch items to same height
Row(
    items,
    .CrossAxisAlignment(AxisAlignmentTypeStretch),
)`),
			Spacer().Height(16),

			layoutSubsection("Flexible Rows", "Distribute space between row items."),
			CodeBlock(`Row(
    []Widget{
        Container(
            Text("Left Panel"),
            .Flex(1), // Takes 1/3 of space
        ),
        Container(
            Text("Main Content"),
            .Flex(2), // Takes 2/3 of space
        ),
        Container(
            Text("Right Panel"),
            container.Width(AllBP(200)), // Fixed width
        ),
    },
    .Gap(16),
)`),
			Spacer().Height(24),

			// Grid Layouts
			layoutContentSection("Grid Layouts", "Grids create two-dimensional layouts perfect for cards, galleries, and dashboard-style interfaces."),
			Spacer().Height(16),

			layoutSubsection("Responsive Grid", "Create grids that adapt to screen size."),
			CodeBlock(`grid.New(
    []Widget{
        productCard("Product 1"),
        productCard("Product 2"),
        productCard("Product 3"),
        productCard("Product 4"),
        productCard("Product 5"),
        productCard("Product 6"),
    },
    grid.ColumnCount(
        breakpoint.XS(1),  // 1 column on extra small screens
        breakpoint.SM(1),  // 1 column on small screens
        breakpoint.MD(2),  // 2 columns on medium screens
        breakpoint.LG(3),  // 3 columns on large screens
        breakpoint.XL(4),  // 4 columns on extra large screens
    ),
    grid.ColumnGap(16),
    grid.RowGap(16),
)`),
			Spacer().Height(16),

			layoutSubsection("Fixed Grid", "Create grids with consistent column counts."),
			CodeBlock(`// 3-column grid for desktop dashboard
grid.New(
    []Widget{
        dashboardCard("Sales", "$12,345"),
        dashboardCard("Orders", "156"),
        dashboardCard("Users", "1,234"),
        dashboardCard("Revenue", "$45,678"),
        dashboardCard("Growth", "+12%"),
        dashboardCard("Conversion", "3.2%"),
    },
    grid.ColumnCount(AllBP(3)),
    grid.ColumnGap(24),
    grid.RowGap(24),
)`),
			Spacer().Height(24),

			// Container Layouts
			layoutContentSection("Container Layouts", "Containers provide structure, spacing, and styling to your layouts."),
			Spacer().Height(16),

			layoutSubsection("Basic Container", "Wrap content with padding and styling."),
			CodeBlock(`Container(
    Column(
        []Widget{
            Text("Card Title", .FontWeight("600")),
            Text("Card content goes here..."),
        },
        .Gap(8),
    ),
    .Padding(AllBP(All(16))),
    container.BackgroundColor("#FFFFFF"),
    container.BorderRadius(8),
    container.BorderColor("#E5E7EB"),
    container.BorderWidth(1, 1, 1, 1),
)`),
			Spacer().Height(16),

			layoutSubsection("Responsive Container", "Containers that adapt to screen size."),
			CodeBlock(`Container(
    content,
    .Padding(
        breakpoint.XS(All(8)),
        breakpoint.SM(All(16)),
        breakpoint.MD(All(24)),
        breakpoint.LG(All(32)),
    ),
    container.MaxWidth(
        breakpoint.SM(AllBP(640)),
        breakpoint.MD(AllBP(768)),
        breakpoint.LG(AllBP(1024)),
        breakpoint.XL(AllBP(1280)),
    ),
)`),
			Spacer().Height(24),

			// Responsive Design
			layoutContentSection("Responsive Design", "gofred's breakpoint system makes it easy to create layouts that work on any device."),
			Spacer().Height(16),

			layoutSubsection("Breakpoint System", "Define different behaviors for different screen sizes."),
			CodeBlock(`// Available breakpoints:
// breakpoint.XS   - Extra small (< 640px)
// breakpoint.SM   - Small (≥ 640px)
// breakpoint.MD   - Medium (≥ 768px)
// breakpoint.LG   - Large (≥ 1024px)
// breakpoint.XL   - Extra large (≥ 1280px)
// breakpoint.XXL  - 2X large (≥ 1536px)

// Example: Different layouts for different screen sizes
func responsiveLayout() Widget {
    return Container(
        grid.New(
            contentCards(),
            grid.ColumnCount(
                breakpoint.XS(1),  // Stack on mobile
                breakpoint.MD(2),  // Side-by-side on tablet
                breakpoint.LG(4),  // Four columns on desktop
            ),
            grid.ColumnGap(16),
        ),
        .Padding(
            breakpoint.XS(All(16)),
            breakpoint.LG(All(32)),
        ),
    )
}`),
			Spacer().Height(16),

			layoutSubsection("Mobile-First Design", "Start with mobile layouts and enhance for larger screens."),
			CodeBlock(`// Mobile-first navigation
func navigationBar() Widget {
    return Container(
        Row(
            []Widget{
                // Logo
                logo(),
                Spacer(),
                // Mobile: Hamburger menu, Desktop: Full navigation
                Container(
                    Row(
                        []Widget{
                            // Shown on mobile, hidden on desktop
                            Container(
                                hamburgerMenu(),
                                container.Visible(
                                    breakpoint.XS(true),
                                    breakpoint.MD(false),
                                ),
                            ),
                            // Hidden on mobile, shown on desktop
                            Container(
                                Row(navItems(), .Gap(24)),
                                container.Visible(
                                    breakpoint.XS(false),
                                    breakpoint.MD(true),
                                ),
                            ),
                        },
                    ),
                ),
            },
            .CrossAxisAlignment(AxisAlignmentTypeCenter),
        ),
        .Padding(AllBP(All(16))),
    )
}`),
			Spacer().Height(24),

			// Common Layout Patterns
			layoutContentSection("Common Layout Patterns", "Learn proven patterns for building common UI layouts."),
			Spacer().Height(16),

			layoutSubsection("Header-Content-Footer", "The classic three-section layout."),
			CodeBlock(`func appLayout() Widget {
    return Column(
        []Widget{
            // Header
            Container(
                navigationBar(),
                container.Height(AllBP(64)),
                container.BackgroundColor("#FFFFFF"),
                container.BorderColor("#E5E7EB"),
                container.BorderWidth(0, 0, 1, 0),
            ),
            // Main Content
            Container(
                mainContent(),
                .Flex(1), // Takes remaining height
                container.BackgroundColor("#F9FAFB"),
            ),
            // Footer
            Container(
                footer(),
                container.BackgroundColor("#1F2937"),
            ),
        },
        column.Height(AllBP(widget.Context().ClientHeight())),
    )
}`),
			Spacer().Height(16),

			layoutSubsection("Sidebar Layout", "Content with a side navigation panel."),
			CodeBlock(`func dashboardLayout() Widget {
    return Row(
        []Widget{
            // Sidebar
            Container(
                sidebar(),
                container.Width(
                    breakpoint.MD(AllBP(240)),
                    breakpoint.LG(AllBP(280)),
                ),
                container.Visible(
                    breakpoint.XS(false),
                    breakpoint.MD(true),
                ),
                container.BackgroundColor("#FFFFFF"),
                container.BorderColor("#E5E7EB"),
                container.BorderWidth(0, 1, 0, 0),
            ),
            // Main Content
            Container(
                dashboardContent(),
                .Flex(1),
                .Padding(AllBP(All(24))),
            ),
        },
        row.Height(AllBP(widget.Context().ClientHeight())),
    )
}`),
			Spacer().Height(16),

			layoutSubsection("Card Layout", "Reusable card components for content display."),
			CodeBlock(`func cardLayout(title, content string, actions []Widget) Widget {
    return Container(
        Column(
            []Widget{
                // Card Header
                Row(
                    []Widget{
                        Text(
                            title,
                            .FontSize(18),
                            .FontWeight("600"),
                            .FontColor("#1F2937"),
                        ),
                        Spacer(),
                        // Optional header actions
                        Row(actions, .Gap(8)),
                    },
                    .CrossAxisAlignment(AxisAlignmentTypeCenter),
                ),
                // Card Content
                Container(
                    Text(
                        content,
                        .FontSize(14),
                        .FontColor("#6B7280"),
                        text.LineHeight(1.5),
                    ),
                    .Padding(AllBP(Symmetric(0, 16))),
                ),
                // Card Footer
                Row(
                    []Widget{
                        Spacer(),
                        Button(
                            Text("Learn More", .FontColor("#2B799B")),
                            .BackgroundColor("transparent"),
                        ),
                    },
                ),
            },
            .Gap(16),
        ),
        .Padding(AllBP(All(16))),
        container.BackgroundColor("#FFFFFF"),
        container.BorderRadius(8),
        container.BorderColor("#E5E7EB"),
        container.BorderWidth(1, 1, 1, 1),
        container.BorderStyle(options.BorderStyleTypeSolid),
    )
}`),
			Spacer().Height(24),

			// Spacing and Alignment
			layoutContentSection("Spacing & Alignment", "Master the art of spacing and alignment for polished layouts."),
			Spacer().Height(16),

			layoutSubsection("Consistent Spacing", "Use a consistent spacing scale throughout your app."),
			CodeBlock(`// Define your spacing scale
const (
    SpaceXS  = 4
    SpaceSM  = 8
    SpaceMD  = 16
    SpaceLG  = 24
    SpaceXL  = 32
    SpaceXXL = 48
)

// Use spacing consistently
Column(
    []Widget{
        sectionHeader(),
        Spacer(.Height(SpaceLG)),
        sectionContent(),
        Spacer(.Height(SpaceXL)),
        nextSection(),
    },
)`),
			Spacer().Height(16),

			layoutSubsection("Spacer Widget", "Use spacers for flexible and fixed spacing."),
			CodeBlock(`// Fixed spacing
Spacer().Height(24)
Spacer(.Width(16))

// Flexible spacing - takes all available space
Spacer()

// Example: Pushing items apart
Row(
    []Widget{
        leftContent(),
        Spacer(), // Pushes rightContent to the far right
        rightContent(),
    },
)`),
			Spacer().Height(24),

			// Best Practices
			layoutContentSection("Layout Best Practices", "Follow these guidelines for creating effective layouts."),
			layoutBestPracticesList(),
			Spacer().Height(24),

			layoutContentSection("What's Next?", "Now that you understand layouts, explore these related topics:"),
			layoutsNextStepsList(),
			Spacer().Height(32),
			layoutNavigationButtons("/docs/widgets", "/docs/styling"),
		},
	).Gap(16)
}

func layoutContentSection(title, description string) Widget {
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

func layoutSubsection(title, description string) Widget {
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

func layoutBestPracticesList() Widget {
	practices := []string{
		"Start with mobile-first design and progressively enhance for larger screens",
		"Use consistent spacing throughout your application with a defined scale",
		"Choose the right layout widget for your content (column for vertical, row for horizontal, grid for two-dimensional)",
		"Leverage flexbox properties (Flex, CrossAxisAlignment) for flexible layouts",
		"Use containers to group related content and provide consistent styling",
		"Test your layouts across different screen sizes and orientations",
		"Keep layout hierarchy simple - avoid deeply nested layout widgets when possible",
		"Use spacer widgets instead of manual margins for better layout control",
		"Consider content flow and reading patterns when designing layouts",
		"Make interactive elements easily accessible with proper spacing and sizing",
	}

	var practiceItems []Widget
	for _, practice := range practices {
		practiceItems = append(practiceItems, layoutListItem(practice))
	}

	return Column(
		practiceItems,
	).Gap(8)
}

func layoutListItem(itemText string) Widget {
	return Row(
		[]Widget{
			Icon(icondata.Check).
				Width(AllBP(16)).
				Height(AllBP(16)).
				Fill("#10B981"),
			Spacer().Width(8),
			Text(itemText).
				FontSize(16).
				FontColor("#374151").
				FontWeight("400"),
		},
	).Gap(8).CrossAxisAlignment(AxisAlignmentTypeCenter)
}

func layoutsNextStepsList() Widget {
	steps := []struct {
		title       string
		description string
		href        string
	}{
		{
			title:       "Styling",
			description: "Learn how to style your layouts with colors, fonts, and visual effects",
			href:        "/docs/styling",
		},
		{
			title:       "State Management",
			description: "Handle dynamic content and user interactions in your layouts",
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
		stepItems = append(stepItems, layoutNextStepItem(step.title, step.description, step.href))
	}

	return Column(
		stepItems,
	).Gap(12)
}

func layoutNextStepItem(title, description, href string) Widget {
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
					).Gap(4).Flex(1),
					Icon(icondata.ChevronRight).
						Width(AllBP(20)).
						Height(AllBP(20)).
						Fill("#9CA3AF"),
				},
			).Gap(12).Flex(1).CrossAxisAlignment(AxisAlignmentTypeCenter),
		).Padding(AllBP(All(16))).
			BackgroundColor("#FFFFFF").
			BorderRadius(8).
			BorderColor("#E5E7EB").
			BorderWidth(1, 1, 1, 1).
			BorderStyle(BorderStyleTypeSolid),
	).Href(href)
}

func layoutNavigationButtons(previousHref, nextHref string) Widget {
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
					).Gap(8).CrossAxisAlignment(AxisAlignmentTypeCenter),
				).BackgroundColor("#6B7280").Width(AllBP(120)),
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
					).Gap(8).CrossAxisAlignment(AxisAlignmentTypeCenter),
				).BackgroundColor("#2B799B").Width(AllBP(120)),
			).Href(nextHref),
		},
	).Flex(1)
}
