package comingsoon

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/foundation/button"
	. "github.com/gofred-io/gofred/foundation/center"
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

// ComingSoonContent creates a coming soon page with optional title and suggestions
func ComingSoonContent(title string, suggestions []Suggestion) Widget {
	return Container(
		Center(
			Container(
				Column(
					[]Widget{
						comingSoonHeader(title),
						Spacer().Height(32),
						comingSoonIcon(),
						Spacer().Height(24),
						comingSoonMessage(),
						Spacer().Height(32),
						comingSoonSuggestions(suggestions),
						Spacer().Height(32),
						comingSoonActions(),
					},
				).Gap(0).CrossAxisAlignment(AxisAlignmentTypeCenter),
			).MaxWidth(AllBP(600)).Padding(AllBP(All(32))),
		),
	).Flex(1)
}

// Suggestion represents a suggested page or action for users
type Suggestion struct {
	Title       string
	Description string
	Href        string
}

func comingSoonHeader(title string) Widget {
	pageTitle := title
	if pageTitle == "" {
		pageTitle = "Coming Soon"
	}

	return Column(
		[]Widget{
			Text(pageTitle).FontSize(32).FontColor("#1F2937").FontWeight("700"),
			Spacer().Height(8),
			Text("This documentation page is currently under development").FontSize(18).FontColor("#6B7280").FontWeight("400"),
		},
	).Gap(0).CrossAxisAlignment(AxisAlignmentTypeCenter)
}

func comingSoonIcon() Widget {
	return Container(
		Icon(icondata.Clock).Width(AllBP(80)).Height(AllBP(80)).Fill("#9CA3AF"),
	).Padding(AllBP(All(20))).BackgroundColor("#F9FAFB").BorderRadius(40).
		BorderColor("#E5E7EB").BorderWidth(1, 1, 1, 1).BorderStyle(BorderStyleTypeSolid)
}

func comingSoonMessage() Widget {
	return Column(
		[]Widget{
			Text("We're working hard to bring you comprehensive documentation for this topic. In the meantime, you can explore the available sections or check back soon for updates.").FontSize(16).FontColor("#6B7280").FontWeight("400").LineHeight(1.6),
			Spacer().Height(16),
			Container(
				Row(
					[]Widget{
						Icon(icondata.Alert).Width(AllBP(16)).Height(AllBP(16)).Fill("#3B82F6"),
						Text("Want to contribute? This documentation is open source and we welcome contributions!").FontSize(14).FontColor("#3B82F6").FontWeight("400"),
					},
				).Gap(8).CrossAxisAlignment(AxisAlignmentTypeCenter),
			).Padding(AllBP(All(12))).BackgroundColor("#EFF6FF").BorderRadius(8).BorderColor("#DBEAFE").BorderWidth(1, 1, 1, 1).BorderStyle(BorderStyleTypeSolid),
		},
	).Gap(0).CrossAxisAlignment(AxisAlignmentTypeCenter)
}

func comingSoonSuggestions(suggestions []Suggestion) Widget {
	if len(suggestions) == 0 {
		// Default suggestions if none provided
		suggestions = []Suggestion{
			{
				Title:       "Foundation Widgets",
				Description: "Learn about the core building blocks of gofred applications",
				Href:        "/docs/widgets",
			},
			{
				Title:       "Layouts & Responsive Design",
				Description: "Master responsive layouts and component organization",
				Href:        "/docs/layouts",
			},
			{
				Title:       "Styling & Visual Design",
				Description: "Create beautiful interfaces with colors, typography, and spacing",
				Href:        "/docs/styling",
			},
			{
				Title:       "State Management",
				Description: "Handle dynamic data and create reactive user interfaces",
				Href:        "/docs/state",
			},
		}
	}

	var suggestionWidgets []Widget
	suggestionWidgets = append(suggestionWidgets,
		Text("While you're here, check out these available topics:").FontSize(18).FontColor("#1F2937").FontWeight("600"),
	)
	suggestionWidgets = append(suggestionWidgets, Spacer().Height(16))

	for _, suggestion := range suggestions {
		suggestionWidgets = append(suggestionWidgets, suggestionCard(suggestion))
	}

	return Column(
		suggestionWidgets,
	).Gap(12).CrossAxisAlignment(AxisAlignmentTypeCenter)
}

func suggestionCard(suggestion Suggestion) Widget {
	return Link(
		Container(
			Row(
				[]Widget{
					Column(
						[]Widget{
							Text(suggestion.Title).FontSize(16).FontColor("#2B799B").FontWeight("500"),
							Text(suggestion.Description).FontSize(14).FontColor("#6B7280").FontWeight("400").LineHeight(1.4),
						},
					).Gap(4).Flex(1),
					Icon(icondata.ChevronRight).Width(AllBP(20)).Height(AllBP(20)).Fill("#9CA3AF"),
				},
			).Gap(12).Flex(1).CrossAxisAlignment(AxisAlignmentTypeCenter),
		).Padding(AllBP(All(16))).BackgroundColor("#FFFFFF").BorderRadius(8).BorderColor("#E5E7EB").BorderWidth(1, 1, 1, 1).BorderStyle(BorderStyleTypeSolid).Width(AllBP(400)),
	).Href(suggestion.Href)
}

func comingSoonActions() Widget {
	return Column(
		[]Widget{
			Text(
				"Need help or have questions?",
			).FontSize(16).FontColor("#1F2937").FontWeight("500"),
			Spacer().Height(16),
			Row(
				[]Widget{
					Link(
						Button(
							Row(
								[]Widget{
									Icon(icondata.Home).Width(AllBP(16)).Height(AllBP(16)).Fill("#FFFFFF"),
									Text("Back to Docs").FontSize(14).FontColor("#FFFFFF").FontWeight("500"),
								},
							).Gap(8).CrossAxisAlignment(AxisAlignmentTypeCenter),
						).BackgroundColor("#2B799B").BorderRadius(6),
					).Href("/docs"),
					Link(
						Button(
							Row(
								[]Widget{
									Icon(icondata.AccountGroup).Width(AllBP(16)).Height(AllBP(16)).Fill("#FFFFFF"),
									Text("GitHub Discussions").FontSize(14).FontColor("#FFFFFF").FontWeight("500"),
								},
							).Gap(8).CrossAxisAlignment(AxisAlignmentTypeCenter),
						).BackgroundColor("#2B799B").BorderRadius(6),
					).Href("https://github.com/orgs/gofred-io/discussions").NewTab(true),
				},
			).Gap(12).CrossAxisAlignment(AxisAlignmentTypeCenter),
		},
	).Gap(0).CrossAxisAlignment(AxisAlignmentTypeCenter)
}

// Simple coming soon page without custom suggestions
func Simple(title string) Widget {
	return ComingSoonContent(title, nil)
}

// Coming soon page with custom suggestions
func WithSuggestions(title string, suggestions []Suggestion) Widget {
	return ComingSoonContent(title, suggestions)
}

// Coming soon page for a specific documentation section
func ForDocsSection(sectionTitle string, relatedSections []Suggestion) Widget {
	return WithSuggestions(sectionTitle, relatedSections)
}
