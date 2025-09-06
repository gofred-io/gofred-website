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

func StylingContent() Widget {
	return Container(
		Column(
			[]Widget{
				stylingPageHeader(),
				Spacer().Height(24),
				stylingPageContent(),
			},
		).Gap(16).Flex(1),
	).Flex(1).Padding(AllBP(All(32)))
}

func stylingPageHeader() Widget {
	return Column(
		[]Widget{
			Text("Styling & Visual Design").
				FontSize(32).
				FontColor("#1F2937").
				FontWeight("700"),
			Text("Master the art of styling gofred widgets with colors, typography, spacing, borders, and responsive design principles.").
				FontSize(18).
				FontColor("#6B7280").
				FontWeight("400"),
		},
	).Gap(8)
}

func stylingPageContent() Widget {
	return Column(
		[]Widget{
			stylingContentSection("Styling Philosophy", "gofred follows a design system approach where styling is applied directly to widgets using properties and options. This ensures consistent, predictable, and maintainable styling throughout your application."),
			Spacer().Height(24),

			// Color System
			stylingContentSection("Color System", "Colors in gofred can be applied to text, backgrounds, borders, and icons using hex values, RGB, or named colors."),
			Spacer().Height(16),

			stylingSubsection("Text Colors", "Apply colors to text for emphasis, hierarchy, and branding."),
			CodeBlock(`// Basic text colors
Text("Primary Text", .FontColor("#1F2937"))    // Dark gray
Text("Secondary Text", .FontColor("#6B7280"))  // Medium gray
Text("Muted Text", .FontColor("#9CA3AF"))      // Light gray

// Semantic colors
Text("Success Message", .FontColor("#10B981")) // Green
Text("Error Message", .FontColor("#EF4444"))   // Red
Text("Warning Message", .FontColor("#F59E0B")) // Orange
Text("Info Message", .FontColor("#3B82F6"))    // Blue

// Brand colors
Text("Brand Text", .FontColor("#2B799B"))      // Primary brand
Text("Accent Text", .FontColor("#7C3AED"))     // Purple accent`),
			Spacer().Height(16),

			stylingSubsection("Background Colors", "Set background colors for containers, buttons, and other widgets."),
			CodeBlock(`// Container backgrounds
Container(
    content,
    container.BackgroundColor("#FFFFFF"),    // White
    container.BackgroundColor("#F9FAFB"),    // Light gray
    container.BackgroundColor("#1F2937"),    // Dark
)

// Button backgrounds
Button(
    Text("Primary", .FontColor("#FFFFFF")),
    button.BackgroundColor("#2B799B"),       // Primary
)
Button(
    Text("Success", .FontColor("#FFFFFF")),
    button.BackgroundColor("#10B981"),       // Success
)
Button(
    Text("Danger", .FontColor("#FFFFFF")),
    button.BackgroundColor("#EF4444"),       // Danger
)`),
			Spacer().Height(16),

			stylingSubsection("Icon Colors", "Colorize icons to match your design system."),
			CodeBlock(`// Icon colors with fill
Icon(
    icondata.Heart,
    icon.Fill("#EF4444"),              // Red heart
    icon.Width(AllBP(24)),
    icon.Height(AllBP(24)),
)

// Status icons
Icon(icondata.Check, icon.Fill("#10B981"))    // Green check
Icon(icondata.X, icon.Fill("#EF4444"))        // Red X
Icon(icondata.AlertTriangle, icon.Fill("#F59E0B")) // Orange warning`),
			Spacer().Height(24),

			// Typography
			stylingContentSection("Typography", "Typography creates hierarchy, improves readability, and establishes your brand voice."),
			Spacer().Height(16),

			stylingSubsection("Font Sizes", "Use a consistent scale for font sizes throughout your application."),
			CodeBlock(`// Heading sizes
Text("Large Heading", .FontSize(48))      // 48px
Text("Page Title", .FontSize(32))         // 32px
Text("Section Title", .FontSize(24))      // 24px
Text("Subsection", .FontSize(20))         // 20px
Text("Heading", .FontSize(18))            // 18px

// Body text sizes
Text("Large Body", .FontSize(16))         // 16px
Text("Body Text", .FontSize(14))          // 14px
Text("Small Text", .FontSize(12))         // 12px
Text("Caption", .FontSize(10))            // 10px`),
			Spacer().Height(16),

			stylingSubsection("Font Weights", "Control text emphasis with different font weights."),
			CodeBlock(`Text("Thin Text", .FontWeight("100"))
Text("Light Text", .FontWeight("300"))
Text("Normal Text", .FontWeight("400"))     // Default
Text("Medium Text", .FontWeight("500"))
Text("Semibold Text", .FontWeight("600"))
Text("Bold Text", .FontWeight("700"))
Text("Extra Bold", .FontWeight("800"))
Text("Black Text", .FontWeight("900"))`),
			Spacer().Height(16),

			stylingSubsection("Text Alignment & Layout", "Control text alignment and layout properties."),
			CodeBlock(`// Text alignment
Text("Left Aligned", text.TextAlign(options.TextAlignTypeLeft))
Text("Center Aligned", text.TextAlign(options.TextAlignTypeCenter))
Text("Right Aligned", text.TextAlign(options.TextAlignTypeRight))
Text("Justified", text.TextAlign(options.TextAlignTypeJustify))

// Line height for readability
Text(
    "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
    text.LineHeight(1.2),    // Tight
    text.LineHeight(1.5),    // Normal (recommended for body text)
    text.LineHeight(1.8),    // Loose
)

// Text decorations
Text("Underlined Link", text.TextDecoration(options.TextDecorationTypeUnderline))
Text("Strike Through", text.TextDecoration(options.TextDecorationTypeLineThrough))`),
			Spacer().Height(24),

			// Spacing System
			stylingContentSection("Spacing System", "Consistent spacing creates visual rhythm and improves user experience."),
			Spacer().Height(16),

			stylingSubsection("Padding", "Add internal spacing to widgets with padding."),
			CodeBlock(`// Uniform padding
Container(
    content,
    .Padding(AllBP(All(16))),    // 16px all sides
    .Padding(AllBP(All(24))),    // 24px all sides
)

// Symmetric padding (horizontal, vertical)
Container(
    content,
    .Padding(AllBP(spacing.Symmetric(24, 16))), // 24px h, 16px v
)

// Individual sides
Container(
    content,
    .Padding(AllBP(spacing.Only(
        16,  // top
        24,  // right
        16,  // bottom
        12,  // left
    ))),
)

// Responsive padding
Container(
    content,
    .Padding(
        breakpoint.XS(All(8)),   // Small on mobile
        breakpoint.MD(All(16)),  // Medium on tablet
        breakpoint.LG(All(24)),  // Large on desktop
    ),
)`),
			Spacer().Height(16),

			stylingSubsection("Margins & Gaps", "Control spacing between widgets."),
			CodeBlock(`// Column gaps
Column(
    []Widget{item1, item2, item3},
    .Gap(8),     // Small gap
    .Gap(16),    // Medium gap
    .Gap(24),    // Large gap
)

// Row gaps
Row(
    []Widget{item1, item2, item3},
    row.Gap(12),       // Consistent spacing
)

// Grid gaps
grid.New(
    items,
    grid.ColumnGap(16),  // Horizontal spacing
    grid.RowGap(20),     // Vertical spacing
)

// Manual spacing with spacers
Column(
    []Widget{
        section1(),
        Spacer().Height(32),  // Fixed spacing
        section2(),
        Spacer(),                   // Flexible spacing
        footer(),
    },
)`),
			Spacer().Height(24),

			// Borders & Shadows
			stylingContentSection("Borders & Visual Effects", "Add definition and depth to your interfaces with borders and visual effects."),
			Spacer().Height(16),

			stylingSubsection("Border Styles", "Define borders for containers and interactive elements."),
			CodeBlock(`// Basic border
Container(
    content,
    container.BorderColor("#E5E7EB"),
    container.BorderWidth(1, 1, 1, 1),      // top, right, bottom, left
    container.BorderStyle(options.BorderStyleTypeSolid),
)

// Colored borders
Container(
    content,
    container.BorderColor("#2B799B"),       // Blue border
    container.BorderWidth(2, 2, 2, 2),      // Thicker border
)

// Partial borders
Container(
    content,
    container.BorderColor("#E5E7EB"),
    container.BorderWidth(0, 0, 1, 0),      // Bottom border only
)

// Focus states for interactive elements
Button(
    Text("Click me"),
    button.BorderColor("#2B799B"),
    button.BorderWidth(2, 2, 2, 2),
    button.BorderStyle(options.BorderStyleTypeSolid),
)`),
			Spacer().Height(16),

			stylingSubsection("Border Radius", "Round corners for modern, friendly interfaces."),
			CodeBlock(`// Border radius values
Container(content, container.BorderRadius(4))   // Subtle rounding
Container(content, container.BorderRadius(8))   // Standard rounding
Container(content, container.BorderRadius(12))  // More rounded
Container(content, container.BorderRadius(16))  // Rounded
Container(content, container.BorderRadius(24))  // Very rounded

// Circular elements
Container(
    Icon(icondata.User),
    container.BorderRadius(24),          // Full circle
    container.Width(AllBP(48)),
    container.Height(AllBP(48)),
)

// Pill-shaped buttons
Button(
    Text("Pill Button"),
    button.BorderRadius(9999),
    button.Padding(AllBP(spacing.Symmetric(20, 12))),
)`),
			Spacer().Height(24),

			// Responsive Styling
			stylingContentSection("Responsive Styling", "Adapt your styling to different screen sizes and devices."),
			Spacer().Height(16),

			stylingSubsection("Breakpoint-Based Styling", "Apply different styles at different screen sizes."),
			CodeBlock(`// Responsive font sizes
Text(
    "Responsive Heading",
    .FontSize(
        breakpoint.XS(24),    // Small on mobile
        breakpoint.MD(32),    // Medium on tablet
        breakpoint.LG(48),    // Large on desktop
    ),
)

// Responsive padding
Container(
    content,
    .Padding(
        breakpoint.XS(All(8)),
        breakpoint.SM(All(16)),
        breakpoint.MD(All(24)),
        breakpoint.LG(All(32)),
    ),
)

// Responsive visibility
Container(
    mobileOnlyContent,
    container.Visible(
        breakpoint.XS(true),
        breakpoint.MD(false),
    ),
)

// Responsive colors (for themes)
Text(
    "Theme-aware text",
    .FontColor(
        AllBP("#1F2937"),    // Dark text for light theme
        // Could add dark theme support here
    ),
)`),
			Spacer().Height(16),

			stylingSubsection("Mobile-First Styling", "Start with mobile styles and enhance for larger screens."),
			CodeBlock(`// Mobile-first approach
func responsiveCard() Widget {
    return Container(
        cardContent(),
        // Base styles (mobile)
        .Padding(AllBP(All(12))),
        container.BackgroundColor("#FFFFFF"),
        container.BorderRadius(8),
        container.BorderColor("#E5E7EB"),
        container.BorderWidth(1, 1, 1, 1),
        
        // Enhanced styles for larger screens
        .Padding(
            breakpoint.MD(All(16)),
            breakpoint.LG(All(24)),
        ),
        container.BorderRadius(
            breakpoint.MD(12),
        ),
    )
}`),
			Spacer().Height(24),

			// Component Styling Patterns
			stylingContentSection("Component Styling Patterns", "Learn common patterns for styling different types of components."),
			Spacer().Height(16),

			stylingSubsection("Button Styles", "Create consistent button styling across your application."),
			CodeBlock(`// Primary button
func primaryButton(label string) Widget {
    return Button(
        Text(label, .FontColor("#FFFFFF"), .FontWeight("500")),
        button.BackgroundColor("#2B799B"),
        button.BorderRadius(8),
        button.Padding(AllBP(spacing.Symmetric(16, 12))),
    )
}

// Secondary button
func secondaryButton(label string) Widget {
    return Button(
        Text(label, .FontColor("#2B799B"), .FontWeight("500")),
        button.BackgroundColor("transparent"),
        button.BorderColor("#2B799B"),
        button.BorderWidth(1, 1, 1, 1),
        button.BorderRadius(8),
        button.Padding(AllBP(spacing.Symmetric(16, 12))),
    )
}

// Danger button
func dangerButton(label string) Widget {
    return Button(
        Text(label, .FontColor("#FFFFFF"), .FontWeight("500")),
        button.BackgroundColor("#EF4444"),
        button.BorderRadius(8),
        button.Padding(AllBP(spacing.Symmetric(16, 12))),
    )
}`),
			Spacer().Height(16),

			stylingSubsection("Card Styles", "Design consistent card components."),
			CodeBlock(`// Basic card
func basicCard(content Widget) Widget {
    return Container(
        content,
        container.BackgroundColor("#FFFFFF"),
        container.BorderRadius(8),
        container.BorderColor("#E5E7EB"),
        container.BorderWidth(1, 1, 1, 1),
        .Padding(AllBP(All(16))),
    )
}

// Elevated card (with shadow effect using border)
func elevatedCard(content Widget) Widget {
    return Container(
        content,
        container.BackgroundColor("#FFFFFF"),
        container.BorderRadius(12),
        container.BorderColor("#E5E7EB"),
        container.BorderWidth(1, 1, 1, 1),
        .Padding(AllBP(All(20))),
    )
}

// Status cards with colored borders
func statusCard(content Widget, status string) Widget {
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
    
    return Container(
        content,
        container.BackgroundColor("#FFFFFF"),
        container.BorderRadius(8),
        container.BorderColor(borderColor),
        container.BorderWidth(1, 1, 1, 3), // Thicker left border
        .Padding(AllBP(All(16))),
    )
}`),
			Spacer().Height(16),

			stylingSubsection("Form Styling", "Style form elements for better user experience."),
			CodeBlock(`// Form container
func formContainer(children []Widget) Widget {
    return Container(
        Column(children, .Gap(16)),
        container.BackgroundColor("#FFFFFF"),
        container.BorderRadius(8),
        .Padding(AllBP(All(24))),
        container.BorderColor("#E5E7EB"),
        container.BorderWidth(1, 1, 1, 1),
    )
}

// Form field with label
func formField(label, placeholder string) Widget {
    return Column(
        []Widget{
            Text(
                label,
                .FontSize(14),
                .FontWeight("500"),
                .FontColor("#374151"),
            ),
            // Input field would go here
            // This is where you'd add your input widget
        },
        .Gap(6),
    )
}

// Error message styling
func errorMessage(message string) Widget {
    return Text(
        message,
        .FontSize(12),
        .FontColor("#EF4444"),
        .FontWeight("400"),
    )
}`),
			Spacer().Height(24),

			// Design System
			stylingContentSection("Building a Design System", "Create a consistent design system for your application."),
			Spacer().Height(16),

			stylingSubsection("Color Palette", "Define a consistent color palette."),
			CodeBlock(`// Define your color palette
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
Text("Primary text", .FontColor(Gray900))
Text("Secondary text", .FontColor(Gray600))
Container(content, container.BackgroundColor(BackgroundWhite))`),
			Spacer().Height(16),

			stylingSubsection("Typography Scale", "Establish a typographic hierarchy."),
			CodeBlock(`// Typography scale
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
func headingXL(content string) Widget {
    return Text(
        content,
        .FontSize(Text4XL),
        .FontWeight(FontBold),
        .FontColor(Gray900),
    )
}

func bodyText(content string) Widget {
    return Text(
        content,
        .FontSize(TextBase),
        .FontWeight(FontNormal),
        .FontColor(Gray700),
        text.LineHeight(1.5),
    )
}`),
			Spacer().Height(16),

			stylingSubsection("Spacing Scale", "Use consistent spacing throughout your application."),
			CodeBlock(`// Spacing scale (in pixels)
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
    return All(Space4) // 16px
}

func sectionMargin() Widget {
    return Spacer(.Height(Space8)) // 32px
}

// Use the scale consistently
Container(
    content,
    .Padding(AllBP(cardPadding())),
)`),
			Spacer().Height(24),

			// Best Practices
			stylingContentSection("Styling Best Practices", "Follow these guidelines for effective styling."),
			stylingBestPracticesList(),
			Spacer().Height(24),

			stylingContentSection("What's Next?", "Now that you understand styling, explore these related topics:"),
			stylingNextStepsList(),
			Spacer().Height(32),
			stylingNavigationButtons("/docs/layouts", "/docs/state"),
		},
	).Gap(16)
}

func stylingContentSection(title, description string) Widget {
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

func stylingSubsection(title, description string) Widget {
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

func stylingBestPracticesList() Widget {
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

	var practiceItems []Widget
	for _, practice := range practices {
		practiceItems = append(practiceItems, stylingListItem(practice))
	}

	return Column(
		practiceItems,
	).Gap(8)
}

func stylingListItem(itemText string) Widget {
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

func stylingNextStepsList() Widget {
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

	var stepItems []Widget
	for _, step := range steps {
		stepItems = append(stepItems, stylingNextStepItem(step.title, step.description, step.href))
	}

	return Column(
		stepItems,
	).Gap(12)
}

func stylingNextStepItem(title, description, href string) Widget {
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

func stylingNavigationButtons(previousHref, nextHref string) Widget {
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
