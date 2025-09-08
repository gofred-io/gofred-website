package core_concepts

import (
	"github.com/gofred-io/gofred-website/app/components/codeblock"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/button"
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

func LayoutsContent() widget.BaseWidget {
	return container.New(
		column.New(
			[]widget.BaseWidget{
				layoutsPageHeader(),
				spacer.New(spacer.Height(24)),
				layoutsPageContent(),
			},
			column.Gap(16),
			column.Flex(1),
		),
		container.Flex(1),
		container.Padding(breakpoint.All(spacing.All(32))),
	)
}

func layoutsPageHeader() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			text.New(
				"Layouts & Responsive Design",
				text.FontSize(32),
				text.FontColor("#1F2937"),
				text.FontWeight("700"),
			),
			text.New(
				"Master the art of creating beautiful, responsive layouts that work perfectly across all devices and screen sizes.",
				text.FontSize(18),
				text.FontColor("#6B7280"),
				text.FontWeight("400"),
			),
		},
		column.Gap(8),
	)
}

func layoutsPageContent() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			layoutContentSection("Layout Fundamentals", "Layouts in gofred are built using a flexbox-inspired system that makes it easy to create responsive, flexible designs. The main layout widgets work together to help you arrange content exactly how you want it."),
			spacer.New(spacer.Height(24)),

			// Column Layouts
			layoutContentSection("Column Layouts", "Columns arrange widgets vertically, perfect for forms, lists, and content that flows from top to bottom."),
			spacer.New(spacer.Height(16)),

			layoutSubsection("Basic Column", "The foundation of vertical layouts."),
			codeblock.New(`column.New(
    []widget.BaseWidget{
        text.New("Header", text.FontSize(24), text.FontWeight("600")),
        text.New("Subtitle", text.FontSize(16), text.FontColor("#6B7280")),
        text.New("Content goes here..."),
        button.New(text.New("Action Button")),
    },
    column.Gap(16), // Space between items
)`),
			spacer.New(spacer.Height(16)),

			layoutSubsection("Column Alignment", "Control how items are aligned within the column."),
			codeblock.New(`// Center-aligned column
column.New(
    []widget.BaseWidget{
        text.New("Centered Title"),
        text.New("Centered content"),
        button.New(text.New("Centered Button")),
    },
    column.Gap(12),
    column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
)

// Right-aligned column
column.New(
    items,
    column.CrossAxisAlignment(options.AxisAlignmentTypeEnd),
)`),
			spacer.New(spacer.Height(16)),

			layoutSubsection("Flexible Columns", "Make columns expand to fill available space."),
			codeblock.New(`column.New(
    []widget.BaseWidget{
        header(),
        column.New(
            contentItems,
            column.Flex(1), // This column takes remaining space
        ),
        footer(),
    },
    column.Gap(0),
)`),
			spacer.New(spacer.Height(24)),

			// Row Layouts
			layoutContentSection("Row Layouts", "Rows arrange widgets horizontally, ideal for navigation bars, button groups, and side-by-side content."),
			spacer.New(spacer.Height(16)),

			layoutSubsection("Basic Row", "The foundation of horizontal layouts."),
			codeblock.New(`row.New(
    []widget.BaseWidget{
        icon.New(icondata.User, icon.Width(breakpoint.All(20))),
        text.New("John Doe", text.FontWeight("500")),
        spacer.New(), // Pushes next items to the right
        text.New("Online", text.FontColor("#10B981")),
        button.New(text.New("Contact")),
    },
    row.Gap(12),
    row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
)`),
			spacer.New(spacer.Height(16)),

			layoutSubsection("Row Alignment", "Control vertical alignment of items in a row."),
			codeblock.New(`// Top-aligned row
row.New(
    []widget.BaseWidget{
        largeImage(),
        column.New(smallTextItems),
    },
    row.CrossAxisAlignment(options.AxisAlignmentTypeStart),
)

// Bottom-aligned row
row.New(
    items,
    row.CrossAxisAlignment(options.AxisAlignmentTypeEnd),
)

// Stretch items to same height
row.New(
    items,
    row.CrossAxisAlignment(options.AxisAlignmentTypeStretch),
)`),
			spacer.New(spacer.Height(16)),

			layoutSubsection("Flexible Rows", "Distribute space between row items."),
			codeblock.New(`row.New(
    []widget.BaseWidget{
        container.New(
            text.New("Left Panel"),
            container.Flex(1), // Takes 1/3 of space
        ),
        container.New(
            text.New("Main Content"),
            container.Flex(2), // Takes 2/3 of space
        ),
        container.New(
            text.New("Right Panel"),
            container.Width(breakpoint.All(200)), // Fixed width
        ),
    },
    row.Gap(16),
)`),
			spacer.New(spacer.Height(24)),

			// Grid Layouts
			layoutContentSection("Grid Layouts", "Grids create two-dimensional layouts perfect for cards, galleries, and dashboard-style interfaces."),
			spacer.New(spacer.Height(16)),

			layoutSubsection("Responsive Grid", "Create grids that adapt to screen size."),
			codeblock.New(`grid.New(
    []widget.BaseWidget{
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
			spacer.New(spacer.Height(16)),

			layoutSubsection("Fixed Grid", "Create grids with consistent column counts."),
			codeblock.New(`// 3-column grid for desktop dashboard
grid.New(
    []widget.BaseWidget{
        dashboardCard("Sales", "$12,345"),
        dashboardCard("Orders", "156"),
        dashboardCard("Users", "1,234"),
        dashboardCard("Revenue", "$45,678"),
        dashboardCard("Growth", "+12%"),
        dashboardCard("Conversion", "3.2%"),
    },
    grid.ColumnCount(breakpoint.All(3)),
    grid.ColumnGap(24),
    grid.RowGap(24),
)`),
			spacer.New(spacer.Height(24)),

			// Container Layouts
			layoutContentSection("Container Layouts", "Containers provide structure, spacing, and styling to your layouts."),
			spacer.New(spacer.Height(16)),

			layoutSubsection("Basic Container", "Wrap content with padding and styling."),
			codeblock.New(`container.New(
    column.New(
        []widget.BaseWidget{
            text.New("Card Title", text.FontWeight("600")),
            text.New("Card content goes here..."),
        },
        column.Gap(8),
    ),
    container.Padding(breakpoint.All(spacing.All(16))),
    container.BackgroundColor("#FFFFFF"),
    container.BorderRadius(8),
    container.BorderColor("#E5E7EB"),
    container.BorderWidth(1, 1, 1, 1),
)`),
			spacer.New(spacer.Height(16)),

			layoutSubsection("Responsive Container", "Containers that adapt to screen size."),
			codeblock.New(`container.New(
    content,
    container.Padding(
        breakpoint.XS(spacing.All(8)),
        breakpoint.SM(spacing.All(16)),
        breakpoint.MD(spacing.All(24)),
        breakpoint.LG(spacing.All(32)),
    ),
    container.MaxWidth(
        breakpoint.SM(breakpoint.All(640)),
        breakpoint.MD(breakpoint.All(768)),
        breakpoint.LG(breakpoint.All(1024)),
        breakpoint.XL(breakpoint.All(1280)),
    ),
)`),
			spacer.New(spacer.Height(24)),

			// Responsive Design
			layoutContentSection("Responsive Design", "gofred's breakpoint system makes it easy to create layouts that work on any device."),
			spacer.New(spacer.Height(16)),

			layoutSubsection("Breakpoint System", "Define different behaviors for different screen sizes."),
			codeblock.New(`// Available breakpoints:
// breakpoint.XS   - Extra small (< 640px)
// breakpoint.SM   - Small (≥ 640px)
// breakpoint.MD   - Medium (≥ 768px)
// breakpoint.LG   - Large (≥ 1024px)
// breakpoint.XL   - Extra large (≥ 1280px)
// breakpoint.XXL  - 2X large (≥ 1536px)

// Example: Different layouts for different screen sizes
func responsiveLayout() widget.BaseWidget {
    return container.New(
        grid.New(
            contentCards(),
            grid.ColumnCount(
                breakpoint.XS(1),  // Stack on mobile
                breakpoint.MD(2),  // Side-by-side on tablet
                breakpoint.LG(4),  // Four columns on desktop
            ),
            grid.ColumnGap(16),
        ),
        container.Padding(
            breakpoint.XS(spacing.All(16)),
            breakpoint.LG(spacing.All(32)),
        ),
    )
}`),
			spacer.New(spacer.Height(16)),

			layoutSubsection("Mobile-First Design", "Start with mobile layouts and enhance for larger screens."),
			codeblock.New(`// Mobile-first navigation
func navigationBar() widget.BaseWidget {
    return container.New(
        row.New(
            []widget.BaseWidget{
                // Logo
                logo(),
                spacer.New(),
                // Mobile: Hamburger menu, Desktop: Full navigation
                container.New(
                    row.New(
                        []widget.BaseWidget{
                            // Shown on mobile, hidden on desktop
                            container.New(
                                hamburgerMenu(),
                                container.Visible(
                                    breakpoint.XS(true),
                                    breakpoint.MD(false),
                                ),
                            ),
                            // Hidden on mobile, shown on desktop
                            container.New(
                                row.New(navItems(), row.Gap(24)),
                                container.Visible(
                                    breakpoint.XS(false),
                                    breakpoint.MD(true),
                                ),
                            ),
                        },
                    ),
                ),
            },
            row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
        ),
        container.Padding(breakpoint.All(spacing.All(16))),
    )
}`),
			spacer.New(spacer.Height(24)),

			// Common Layout Patterns
			layoutContentSection("Common Layout Patterns", "Learn proven patterns for building common UI layouts."),
			spacer.New(spacer.Height(16)),

			layoutSubsection("Header-Content-Footer", "The classic three-section layout."),
			codeblock.New(`func appLayout() widget.BaseWidget {
    return column.New(
        []widget.BaseWidget{
            // Header
            container.New(
                navigationBar(),
                container.Height(breakpoint.All(64)),
                container.BackgroundColor("#FFFFFF"),
                container.BorderColor("#E5E7EB"),
                container.BorderWidth(0, 0, 1, 0),
            ),
            // Main Content
            container.New(
                mainContent(),
                container.Flex(1), // Takes remaining height
                container.BackgroundColor("#F9FAFB"),
            ),
            // Footer
            container.New(
                footer(),
                container.BackgroundColor("#1F2937"),
            ),
        },
        column.Height(breakpoint.All(widget.Context().ClientHeight())),
    )
}`),
			spacer.New(spacer.Height(16)),

			layoutSubsection("Sidebar Layout", "Content with a side navigation panel."),
			codeblock.New(`func dashboardLayout() widget.BaseWidget {
    return row.New(
        []widget.BaseWidget{
            // Sidebar
            container.New(
                sidebar(),
                container.Width(
                    breakpoint.MD(breakpoint.All(240)),
                    breakpoint.LG(breakpoint.All(280)),
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
            container.New(
                dashboardContent(),
                container.Flex(1),
                container.Padding(breakpoint.All(spacing.All(24))),
            ),
        },
        row.Height(breakpoint.All(widget.Context().ClientHeight())),
    )
}`),
			spacer.New(spacer.Height(16)),

			layoutSubsection("Card Layout", "Reusable card components for content display."),
			codeblock.New(`func cardLayout(title, content string, actions []widget.BaseWidget) widget.BaseWidget {
    return container.New(
        column.New(
            []widget.BaseWidget{
                // Card Header
                row.New(
                    []widget.BaseWidget{
                        text.New(
                            title,
                            text.FontSize(18),
                            text.FontWeight("600"),
                            text.FontColor("#1F2937"),
                        ),
                        spacer.New(),
                        // Optional header actions
                        row.New(actions, row.Gap(8)),
                    },
                    row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
                ),
                // Card Content
                container.New(
                    text.New(
                        content,
                        text.FontSize(14),
                        text.FontColor("#6B7280"),
                        text.LineHeight(1.5),
                    ),
                    container.Padding(breakpoint.All(spacing.Symmetric(0, 16))),
                ),
                // Card Footer
                row.New(
                    []widget.BaseWidget{
                        spacer.New(),
                        button.New(
                            text.New("Learn More", text.FontColor("#2B799B")),
                            button.BackgroundColor("transparent"),
                        ),
                    },
                ),
            },
            column.Gap(16),
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

			// Spacing and Alignment
			layoutContentSection("Spacing & Alignment", "Master the art of spacing and alignment for polished layouts."),
			spacer.New(spacer.Height(16)),

			layoutSubsection("Consistent Spacing", "Use a consistent spacing scale throughout your app."),
			codeblock.New(`// Define your spacing scale
const (
    SpaceXS  = 4
    SpaceSM  = 8
    SpaceMD  = 16
    SpaceLG  = 24
    SpaceXL  = 32
    SpaceXXL = 48
)

// Use spacing consistently
column.New(
    []widget.BaseWidget{
        sectionHeader(),
        spacer.New(spacer.Height(SpaceLG)),
        sectionContent(),
        spacer.New(spacer.Height(SpaceXL)),
        nextSection(),
    },
)`),
			spacer.New(spacer.Height(16)),

			layoutSubsection("Spacer Widget", "Use spacers for flexible and fixed spacing."),
			codeblock.New(`// Fixed spacing
spacer.New(spacer.Height(24))
spacer.New(spacer.Width(16))

// Flexible spacing - takes all available space
spacer.New()

// Example: Pushing items apart
row.New(
    []widget.BaseWidget{
        leftContent(),
        spacer.New(), // Pushes rightContent to the far right
        rightContent(),
    },
)`),
			spacer.New(spacer.Height(24)),

			// Best Practices
			layoutContentSection("Layout Best Practices", "Follow these guidelines for creating effective layouts."),
			layoutBestPracticesList(),
			spacer.New(spacer.Height(24)),

			layoutContentSection("What's Next?", "Now that you understand layouts, explore these related topics:"),
			layoutsNextStepsList(),
			spacer.New(spacer.Height(32)),
			layoutNavigationButtons("/docs/widgets", "/docs/styling"),
		},
		column.Gap(16),
	)
}

func layoutContentSection(title, description string) widget.BaseWidget {
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

func layoutSubsection(title, description string) widget.BaseWidget {
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

func layoutBestPracticesList() widget.BaseWidget {
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

	var practiceItems []widget.BaseWidget
	for _, practice := range practices {
		practiceItems = append(practiceItems, layoutListItem(practice))
	}

	return column.New(
		practiceItems,
		column.Gap(8),
	)
}

func layoutListItem(itemText string) widget.BaseWidget {
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

func layoutsNextStepsList() widget.BaseWidget {
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

	var stepItems []widget.BaseWidget
	for _, step := range steps {
		stepItems = append(stepItems, layoutNextStepItem(step.title, step.description, step.href))
	}

	return column.New(
		stepItems,
		column.Gap(12),
	)
}

func layoutNextStepItem(title, description, href string) widget.BaseWidget {
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

func layoutNavigationButtons(previousHref, nextHref string) widget.BaseWidget {
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
