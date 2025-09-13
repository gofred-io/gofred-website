package core_concepts

import (
	"github.com/gofred-io/gofred-website/app/components/codeblock"
	"github.com/gofred-io/gofred/application"
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
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
)

func StylingContent() application.BaseWidget {
	return container.New(
		column.New(
			[]application.BaseWidget{
				stylingPageHeader(),
				spacer.New(spacer.Height(24)),
				stylingPageContent(),
			},
			column.Gap(16),
			column.Flex(1),
		),
		container.Flex(1),
		container.Padding(breakpoint.All(spacing.All(32))),
	)
}

func stylingPageHeader() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			text.New(
				"Styling & Visual Design",
				text.FontSize(32),
				text.FontWeight("700"),
			),
			text.New(
				"Master the art of styling gofred widgets with colors, typography, spacing, borders, and responsive design principles.",
				text.FontSize(18),
				text.FontColor("#6B7280"),
			),
		},
		column.Gap(8),
	)
}

func stylingPageContent() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			stylingContentSection("Styling Philosophy", "gofred follows a design system approach where styling is applied directly to widgets using properties and options. This ensures consistent, predictable, and maintainable styling throughout your application."),
			spacer.New(spacer.Height(24)),

			// Color System
			stylingContentSection("Color System", "Colors in gofred can be applied to text, backgrounds, borders, and icons using hex values, RGB, or named colors."),
			spacer.New(spacer.Height(16)),

			stylingSubsection("Text Colors", "Apply colors to text for emphasis, hierarchy, and branding."),
			codeblock.New(`// Basic text colors
text.New("Primary Text", text.FontColor("#1F2937"))    // Dark gray
text.New("Secondary Text", text.FontColor("#6B7280"))  // Medium gray
text.New("Muted Text", text.FontColor("#9CA3AF"))      // Light gray

// Semantic colors
text.New("Success Message", text.FontColor("#10B981")) // Green
text.New("Error Message", text.FontColor("#EF4444"))   // Red
text.New("Warning Message", text.FontColor("#F59E0B")) // Orange
text.New("Info Message", text.FontColor("#3B82F6"))    // Blue

// Brand colors
text.New("Brand Text", text.FontColor("#2B799B"))      // Primary brand
text.New("Accent Text", text.FontColor("#7C3AED"))     // Purple accent`),
			spacer.New(spacer.Height(16)),

			stylingSubsection("Background Colors", "Set background colors for containers, buttons, and other widgets."),
			codeblock.New(`// Container backgrounds
container.New(
    content,
    container.BackgroundColor("#FFFFFF"),    // White
    container.BackgroundColor("#F9FAFB"),    // Light gray
    container.BackgroundColor("#1F2937"),    // Dark
)

// Button backgrounds
button.New(
    text.New("Primary", text.FontColor("#FFFFFF")),
    button.BackgroundColor("#2B799B"),       // Primary
)
button.New(
    text.New("Success", text.FontColor("#FFFFFF")),
    button.BackgroundColor("#10B981"),       // Success
)
button.New(
    text.New("Danger", text.FontColor("#FFFFFF")),
    button.BackgroundColor("#EF4444"),       // Danger
)`),
			spacer.New(spacer.Height(16)),

			stylingSubsection("Icon Colors", "Colorize icons to match your design system."),
			codeblock.New(`// Icon colors with fill
icon.New(
    icondata.Heart,
    icon.Fill("#EF4444"),              // Red heart
    icon.Width(breakpoint.All(24)),
    icon.Height(breakpoint.All(24)),
)

// Status icons
icon.New(icondata.Check, icon.Fill("#10B981"))    // Green check
icon.New(icondata.X, icon.Fill("#EF4444"))        // Red X
icon.New(icondata.AlertTriangle, icon.Fill("#F59E0B")) // Orange warning`),
			spacer.New(spacer.Height(24)),

			// Typography
			stylingContentSection("Typography", "Typography creates hierarchy, improves readability, and establishes your brand voice."),
			spacer.New(spacer.Height(16)),

			stylingSubsection("Font Sizes", "Use a consistent scale for font sizes throughout your application."),
			codeblock.New(`// Heading sizes
text.New("Large Heading", text.FontSize(48))      // 48px
text.New("Page Title", text.FontSize(32))         // 32px
text.New("Section Title", text.FontSize(24))      // 24px
text.New("Subsection", text.FontSize(20))         // 20px
text.New("Heading", text.FontSize(18))            // 18px

// Body text sizes
text.New("Large Body", text.FontSize(16))         // 16px
text.New("Body Text", text.FontSize(14))          // 14px
text.New("Small Text", text.FontSize(12))         // 12px
text.New("Caption", text.FontSize(10))            // 10px`),
			spacer.New(spacer.Height(16)),

			stylingSubsection("Font Weights", "Control text emphasis with different font weights."),
			codeblock.New(`text.New("Thin Text", text.FontWeight("100"))
text.New("Light Text", text.FontWeight("300"))
text.New("Normal Text", text.FontWeight("400"))     // Default
text.New("Medium Text", text.FontWeight("500"))
text.New("Semibold Text", text.FontWeight("600"))
text.New("Bold Text", text.FontWeight("700"))
text.New("Extra Bold", text.FontWeight("800"))
text.New("Black Text", text.FontWeight("900"))`),
			spacer.New(spacer.Height(16)),

			stylingSubsection("Text Alignment & Layout", "Control text alignment and layout properties."),
			codeblock.New(`// Text alignment
text.New("Left Aligned", text.TextAlign(options.TextAlignTypeLeft))
text.New("Center Aligned", text.TextAlign(theme.TextAlignTypeCenter))
text.New("Right Aligned", text.TextAlign(options.TextAlignTypeRight))
text.New("Justified", text.TextAlign(options.TextAlignTypeJustify))

// Line height for readability
text.New(
    "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
    text.LineHeight(1.2),    // Tight
    text.LineHeight(1.5),    // Normal (recommended for body text)
    text.LineHeight(1.8),    // Loose
)

// Text decorations
text.New("Underlined Link", text.TextDecoration(options.TextDecorationTypeUnderline))
text.New("Strike Through", text.TextDecoration(options.TextDecorationTypeLineThrough))`),
			spacer.New(spacer.Height(24)),

			// Spacing System
			stylingContentSection("Spacing System", "Consistent spacing creates visual rhythm and improves user experience."),
			spacer.New(spacer.Height(16)),

			stylingSubsection("Padding", "Add internal spacing to widgets with padding."),
			codeblock.New(`// Uniform padding
container.New(
    content,
    container.Padding(breakpoint.All(spacing.All(16))),    // 16px all sides
    container.Padding(breakpoint.All(spacing.All(24))),    // 24px all sides
)

// Symmetric padding (horizontal, vertical)
container.New(
    content,
    container.Padding(breakpoint.All(spacing.Symmetric(24, 16))), // 24px h, 16px v
)

// Individual sides
container.New(
    content,
    container.Padding(breakpoint.All(spacing.Only(
        16,  // top
        24,  // right
        16,  // bottom
        12,  // left
    ))),
)

// Responsive padding
container.New(
    content,
    container.Padding(
        breakpoint.XS(spacing.All(8)),   // Small on mobile
        breakpoint.MD(spacing.All(16)),  // Medium on tablet
        breakpoint.LG(spacing.All(24)),  // Large on desktop
    ),
)`),
			spacer.New(spacer.Height(16)),

			stylingSubsection("Margins & Gaps", "Control spacing between widgets."),
			codeblock.New(`// Column gaps
column.New(
    []application.BaseWidget{item1, item2, item3},
    column.Gap(8),     // Small gap
    column.Gap(16),    // Medium gap
    column.Gap(24),    // Large gap
)

// Row gaps
row.New(
    []application.BaseWidget{item1, item2, item3},
    row.Gap(12),       // Consistent spacing
)

// Grid gaps
grid.New(
    items,
    grid.ColumnGap(16),  // Horizontal spacing
    grid.RowGap(20),     // Vertical spacing
)

// Manual spacing with spacers
column.New(
    []application.BaseWidget{
        section1(),
        spacer.New(spacer.Height(32)),  // Fixed spacing
        section2(),
        spacer.New(),                   // Flexible spacing
        footer(),
    },
)`),
			spacer.New(spacer.Height(24)),

			// Borders & Shadows
			stylingContentSection("Borders & Visual Effects", "Add definition and depth to your interfaces with borders and visual effects."),
			spacer.New(spacer.Height(16)),

			stylingSubsection("Border Styles", "Define borders for containers and interactive elements."),
			codeblock.New(`// Basic border
container.New(
    content,
    container.BorderColor("#E5E7EB"),
    container.BorderWidth(spacing.All(1)),      // top, right, bottom, left
    container.BorderStyle(theme.BorderStyleTypeSolid),
)

// Colored borders
container.New(
    content,
    container.BorderColor("#2B799B"),       // Blue border
    container.BorderWidth(2, 2, 2, 2),      // Thicker border
)

// Partial borders
container.New(
    content,
    container.BorderColor("#E5E7EB"),
    container.BorderWidth(0, 0, 1, 0),      // Bottom border only
)

// Focus states for interactive elements
button.New(
    text.New("Click me"),
    button.BorderColor("#2B799B"),
    button.BorderWidth(2, 2, 2, 2),
    button.BorderStyle(theme.BorderStyleTypeSolid),
)`),
			spacer.New(spacer.Height(16)),

			stylingSubsection("Border Radius", "Round corners for modern, friendly interfaces."),
			codeblock.New(`// Border radius values
container.New(content, container.BorderRadius(4))   // Subtle rounding
container.New(content, container.BorderRadius(8))   // Standard rounding
container.New(content, container.BorderRadius(12))  // More rounded
container.New(content, container.BorderRadius(16))  // Rounded
container.New(content, container.BorderRadius(24))  // Very rounded

// Circular elements
container.New(
    icon.New(icondata.User),
    container.BorderRadius(24),          // Full circle
    container.Width(breakpoint.All(48)),
    container.Height(breakpoint.All(48)),
)

// Pill-shaped buttons
button.New(
    text.New("Pill Button"),
    button.BorderRadius(9999),
    button.Padding(breakpoint.All(spacing.Symmetric(20, 12))),
)`),
			spacer.New(spacer.Height(24)),

			// Responsive Styling
			stylingContentSection("Responsive Styling", "Adapt your styling to different screen sizes and devices."),
			spacer.New(spacer.Height(16)),

			stylingSubsection("Breakpoint-Based Styling", "Apply different styles at different screen sizes."),
			codeblock.New(`// Responsive font sizes
text.New(
    "Responsive Heading",
    text.FontSize(
        breakpoint.XS(24),    // Small on mobile
        breakpoint.MD(32),    // Medium on tablet
        breakpoint.LG(48),    // Large on desktop
    ),
)

// Responsive padding
container.New(
    content,
    container.Padding(
        breakpoint.XS(spacing.All(8)),
        breakpoint.SM(spacing.All(16)),
        breakpoint.MD(spacing.All(24)),
        breakpoint.LG(spacing.All(32)),
    ),
)

// Responsive visibility
container.New(
    mobileOnlyContent,
    container.Visible(
        breakpoint.XS(true),
        breakpoint.MD(false),
    ),
)

// Responsive colors (for themes)
text.New(
    "Theme-aware text",
    text.FontColor(
        breakpoint.All("#1F2937"),    // Dark text for light theme
        // Could add dark theme support here
    ),
)`),
			spacer.New(spacer.Height(16)),

			stylingSubsection("Mobile-First Styling", "Start with mobile styles and enhance for larger screens."),
			codeblock.New(`// Mobile-first approach
func responsiveCard() application.BaseWidget {
    return container.New(
        cardContent(),
        // Base styles (mobile)
        container.Padding(breakpoint.All(spacing.All(12))),
        container.BackgroundColor("#FFFFFF"),
        container.BorderRadius(8),
        container.BorderColor("#E5E7EB"),
        container.BorderWidth(spacing.All(1)),
        
        // Enhanced styles for larger screens
        container.Padding(
            breakpoint.MD(spacing.All(16)),
            breakpoint.LG(spacing.All(24)),
        ),
        container.BorderRadius(
            breakpoint.MD(12),
        ),
    )
}`),
			spacer.New(spacer.Height(24)),

			// Component Styling Patterns
			stylingContentSection("Component Styling Patterns", "Learn common patterns for styling different types of components."),
			spacer.New(spacer.Height(16)),

			stylingSubsection("Button Styles", "Create consistent button styling across your application."),
			codeblock.New(`// Primary button
func primaryButton(label string) application.BaseWidget {
    return button.New(
        text.New(label, text.FontColor("#FFFFFF"), text.FontWeight("500")),
        button.BackgroundColor("#2B799B"),
        button.BorderRadius(8),
        button.Padding(breakpoint.All(spacing.Symmetric(16, 12))),
    )
}

// Secondary button
func secondaryButton(label string) application.BaseWidget {
    return button.New(
        text.New(label, text.FontColor("#2B799B"), text.FontWeight("500")),
        button.BackgroundColor("transparent"),
        button.BorderColor("#2B799B"),
        button.BorderWidth(1, 1, 1, 1),
        button.BorderRadius(8),
        button.Padding(breakpoint.All(spacing.Symmetric(16, 12))),
    )
}

// Danger button
func dangerButton(label string) application.BaseWidget {
    return button.New(
        text.New(label, text.FontColor("#FFFFFF"), text.FontWeight("500")),
        button.BackgroundColor("#EF4444"),
        button.BorderRadius(8),
        button.Padding(breakpoint.All(spacing.Symmetric(16, 12))),
    )
}`),
			spacer.New(spacer.Height(16)),

			stylingSubsection("Card Styles", "Design consistent card components."),
			codeblock.New(`// Basic card
func basicCard(content application.BaseWidget) application.BaseWidget {
    return container.New(
        content,
        container.BackgroundColor("#FFFFFF"),
        container.BorderRadius(8),
        container.BorderColor("#E5E7EB"),
        container.BorderWidth(spacing.All(1)),
        container.Padding(breakpoint.All(spacing.All(16))),
    )
}

// Elevated card (with shadow effect using border)
func elevatedCard(content application.BaseWidget) application.BaseWidget {
    return container.New(
        content,
        container.BackgroundColor("#FFFFFF"),
        container.BorderRadius(12),
        container.BorderColor("#E5E7EB"),
        container.BorderWidth(spacing.All(1)),
        container.Padding(breakpoint.All(spacing.All(20))),
    )
}

// Status cards with colored borders
func statusCard(content application.BaseWidget, status string) application.BaseWidget {
    var borderColor string
    switch status {
    case "success":
        borderColor = "#10B981"
    case "warning":
        borderColor = "#F59E0B"
    case "error":
        borderColor = "#EF4444"
    default:
        borderColor = "#E5E7EB"
    }
    
    return container.New(
        content,
        container.BackgroundColor("#FFFFFF"),
        container.BorderRadius(8),
        container.BorderColor(borderColor),
        container.BorderWidth(1, 1, 1, 3), // Thicker left border
        container.Padding(breakpoint.All(spacing.All(16))),
    )
}`),
			spacer.New(spacer.Height(16)),

			stylingSubsection("Form Styling", "Style form elements for better user experience."),
			codeblock.New(`// Form container
func formContainer(children []application.BaseWidget) application.BaseWidget {
    return container.New(
        column.New(children, column.Gap(16)),
        container.BackgroundColor("#FFFFFF"),
        container.BorderRadius(8),
        container.Padding(breakpoint.All(spacing.All(24))),
        container.BorderColor("#E5E7EB"),
        container.BorderWidth(spacing.All(1)),
    )
}

// Form field with label
func formField(label, placeholder string) application.BaseWidget {
    return column.New(
        []application.BaseWidget{
            text.New(
                label,
                text.FontSize(14),
                text.FontWeight("500"),
                text.FontColor("#374151"),
            ),
            // Input field would go here
            // This is where you'd add your input widget
        },
        column.Gap(6),
    )
}

// Error message styling
func errorMessage(message string) application.BaseWidget {
    return text.New(
        message,
        text.FontSize(12),
        text.FontColor("#EF4444"),
            )
}`),
			spacer.New(spacer.Height(24)),

			// Design System
			stylingContentSection("Building a Design System", "Create a consistent design system for your application."),
			spacer.New(spacer.Height(16)),

			stylingSubsection("Color Palette", "Define a consistent color palette."),
			codeblock.New(`// Define your color palette
const (
    // Primary colors
    PrimaryBlue    = "#2B799B"
    PrimaryDark    = "#1F5A73"
    PrimaryLight   = "#4A9BC7"
    
    // Neutral colors
    Gray900        = "#1F2937"
    Gray800        = "#374151"
    Gray700        = "#4B5563"
    Gray600        = "#6B7280"
    Gray500        = "#9CA3AF"
    Gray400        = "#D1D5DB"
    Gray300        = "#E5E7EB"
    Gray200        = "#F3F4F6"
    Gray100        = "#F9FAFB"
    Gray50         = "#FAFAFA"
    
    // Semantic colors
    Success        = "#10B981"
    Warning        = "#F59E0B"
    Error          = "#EF4444"
    Info           = "#3B82F6"
    
    // Background colors
    BackgroundWhite = "#FFFFFF"
    BackgroundGray  = "#F9FAFB"
    BackgroundDark  = "#1F2937"
)

// Use consistent colors throughout your app
text.New("Primary text", text.FontColor(Gray900))
text.New("Secondary text", text.FontColor(Gray600))
container.New(content, container.BackgroundColor(BackgroundWhite))`),
			spacer.New(spacer.Height(16)),

			stylingSubsection("Typography Scale", "Establish a typographic hierarchy."),
			codeblock.New(`// Typography scale
const (
    // Font sizes
    TextXS   = 10
    TextSM   = 12
    TextBase = 14
    TextLG   = 16
    TextXL   = 18
    Text2XL  = 20
    Text3XL  = 24
    Text4XL  = 32
    Text5XL  = 48
    
    // Font weights
    FontThin      = "100"
    FontLight     = "300"
    FontNormal    = "400"
    FontMedium    = "500"
    FontSemibold  = "600"
    FontBold      = "700"
)

// Consistent text components
func headingXL(content string) application.BaseWidget {
    return text.New(
        content,
        text.FontSize(Text4XL),
        text.FontWeight(FontBold),
        text.FontColor(Gray900),
    )
}

func bodyText(content string) application.BaseWidget {
    return text.New(
        content,
        text.FontSize(TextBase),
        text.FontWeight(FontNormal),
        text.FontColor(Gray700),
        text.LineHeight(1.5),
    )
}`),
			spacer.New(spacer.Height(16)),

			stylingSubsection("Spacing Scale", "Use consistent spacing throughout your application."),
			codeblock.New(`// Spacing scale (in pixels)
const (
    Space1  = 4
    Space2  = 8
    Space3  = 12
    Space4  = 16
    Space5  = 20
    Space6  = 24
    Space8  = 32
    Space10 = 40
    Space12 = 48
    Space16 = 64
    Space20 = 80
    Space24 = 96
)

// Helper functions for consistent spacing
func cardPadding() spacing.Spacing {
    return spacing.All(Space4) // 16px
}

func sectionMargin() application.BaseWidget {
    return spacer.New(spacer.Height(Space8)) // 32px
}

// Use the scale consistently
container.New(
    content,
    container.Padding(breakpoint.All(cardPadding())),
)`),
			spacer.New(spacer.Height(24)),

			// Best Practices
			stylingContentSection("Styling Best Practices", "Follow these guidelines for effective styling."),
			stylingBestPracticesList(),
			spacer.New(spacer.Height(24)),

			stylingContentSection("What's Next?", "Now that you understand styling, explore these related topics:"),
			stylingNextStepsList(),
			spacer.New(spacer.Height(32)),
			stylingNavigationButtons("/docs/layouts", "/docs/state"),
		},
		column.Gap(16),
	)
}

func stylingContentSection(title, description string) application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			text.New(
				title,
				text.FontSize(24),
				text.FontWeight("600"),
			),
			text.New(
				description,
				text.FontSize(16),
				text.FontColor("#6B7280"),
			),
		},
		column.Gap(8),
	)
}

func stylingSubsection(title, description string) application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			text.New(
				title,
				text.FontSize(20),
				text.FontWeight("600"),
			),
			text.New(
				description,
				text.FontSize(14),
				text.FontColor("#6B7280"),
			),
		},
		column.Gap(4),
	)
}

func stylingBestPracticesList() application.BaseWidget {
	practices := []string{
		"Establish a consistent color palette and stick to it throughout your application",
		"Use a typographic scale with defined font sizes and weights for hierarchy",
		"Create a spacing system and use it consistently for padding, margins, and gaps",
		"Start with mobile-first styling and enhance for larger screens",
		"Group related styling properties into reusable component functions",
		"Use semantic color names (success, warning, error) rather than specific colors",
		"Maintain sufficient color contrast for accessibility (4.5:1 for normal text)",
		"Test your styling across different screen sizes and browsers",
		"Keep styling simple and avoid over-decorating interfaces",
		"Use consistent border radius values throughout your design system",
		"Apply hover and focus states to interactive elements for better UX",
		"Document your design system and share it with your team",
	}

	var practiceItems []application.BaseWidget
	for _, practice := range practices {
		practiceItems = append(practiceItems, stylingListItem(practice))
	}

	return column.New(
		practiceItems,
		column.Gap(8),
	)
}

func stylingListItem(itemText string) application.BaseWidget {
	return row.New(
		[]application.BaseWidget{
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
			),
		},
		row.Gap(8),
		row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
	)
}

func stylingNextStepsList() application.BaseWidget {
	steps := []struct {
		title       string
		description string
		href        string
	}{
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

	var stepItems []application.BaseWidget
	for _, step := range steps {
		stepItems = append(stepItems, stylingNextStepItem(step.title, step.description, step.href))
	}

	return column.New(
		stepItems,
		column.Gap(12),
	)
}

func stylingNextStepItem(title, description, href string) application.BaseWidget {
	return link.New(
		container.New(
			row.New(
				[]application.BaseWidget{
					column.New(
						[]application.BaseWidget{
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
				row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),
			container.Padding(breakpoint.All(spacing.All(16))),
			container.BackgroundColor("#FFFFFF"),
			container.BorderRadius(8),
			container.BorderColor("#E5E7EB"),
			container.BorderWidth(spacing.All(1)),
			container.BorderStyle(theme.BorderStyleTypeSolid),
		),
		link.Href(href),
	)
}

func stylingNavigationButtons(previousHref, nextHref string) application.BaseWidget {
	return row.New(
		[]application.BaseWidget{
			link.New(
				button.New(
					row.New(
						[]application.BaseWidget{
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
						row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
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
						[]application.BaseWidget{
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
						row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
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
