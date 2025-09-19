package comingsoon

import (
	appTheme "github.com/gofred-io/gofred-website/app/theme"

	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/button"
	"github.com/gofred-io/gofred/foundation/center"
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

// ComingSoonContent creates a coming soon page with optional title and suggestions
func ComingSoonContent(title string, suggestions []Suggestion) application.BaseWidget {
	return container.New(
		center.New(
			container.New(
				column.New(
					[]application.BaseWidget{
						comingSoonHeader(title),
						spacer.New(spacer.Height(32)),
						comingSoonIcon(),
						spacer.New(spacer.Height(24)),
						comingSoonMessage(),
						spacer.New(spacer.Height(32)),
						comingSoonSuggestions(suggestions),
						spacer.New(spacer.Height(32)),
						comingSoonActions(),
					},
					column.Gap(0),
					column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
				),
				container.MaxWidth(breakpoint.All(600)),
				container.Padding(breakpoint.All(spacing.All(32))),
			),
		),
		container.Flex(1),
	)
}

// Suggestion represents a suggested page or action for users
type Suggestion struct {
	Title       string
	Description string
	Href        string
}

func comingSoonHeader(title string) application.BaseWidget {
	pageTitle := title
	if pageTitle == "" {
		pageTitle = "Coming Soon"
	}

	return column.New(
		[]application.BaseWidget{
			text.New(
				pageTitle,
				text.FontSize(32),
				text.FontWeight("700"),
			),
			spacer.New(spacer.Height(8)),
			text.New(
				"This documentation page is currently under development",
				text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
				text.FontSize(18),
			),
		},
		column.Gap(0),
		column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
	)
}

func comingSoonIcon() application.BaseWidget {
	return container.New(
		icon.New(
			icondata.Clock,
			icon.Width(breakpoint.All(80)),
			icon.Height(breakpoint.All(80)),
			icon.Fill("#9CA3AF"),
		),
		container.Padding(breakpoint.All(spacing.All(20))),
		container.BackgroundColor("#F9FAFB"),
		container.BorderRadius(40),
		container.BorderColor("#E5E7EB"),
		container.BorderWidth(spacing.All(1)),
		container.BorderStyle(theme.BorderStyleTypeSolid),
	)
}

func comingSoonMessage() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			text.New(
				"We're working hard to bring you comprehensive documentation for this topic. In the meantime, you can explore the available sections or check back soon for updates.",
				text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
				text.FontSize(16),
				text.LineHeight(1.6),
			),
			spacer.New(spacer.Height(16)),
			container.New(
				row.New(
					[]application.BaseWidget{
						icon.New(
							icondata.Alert,
							icon.Width(breakpoint.All(16)),
							icon.Height(breakpoint.All(16)),
							icon.Fill("#3B82F6"),
						),
						text.New(
							"Want to contribute? This documentation is open source and we welcome contributions!",
							text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
							text.FontSize(14),
						),
					},
					row.Gap(8),
					row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
				),
				container.Padding(breakpoint.All(spacing.All(12))),
				container.BackgroundColor("#EFF6FF"),
				container.BorderRadius(8),
				container.BorderColor("#DBEAFE"),
				container.BorderWidth(spacing.All(1)),
				container.BorderStyle(theme.BorderStyleTypeSolid),
			),
		},
		column.Gap(0),
		column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
	)
}

func comingSoonSuggestions(suggestions []Suggestion) application.BaseWidget {
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

	var suggestionWidgets []application.BaseWidget
	suggestionWidgets = append(suggestionWidgets,
		text.New(
			"While you're here, check out these available topics:",
			text.FontSize(18),
			text.FontWeight("600"),
		),
	)
	suggestionWidgets = append(suggestionWidgets, spacer.New(spacer.Height(16)))

	for _, suggestion := range suggestions {
		suggestionWidgets = append(suggestionWidgets, suggestionCard(suggestion))
	}

	return column.New(
		suggestionWidgets,
		column.Gap(12),
		column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
	)
}

func suggestionCard(suggestion Suggestion) application.BaseWidget {
	return link.New(
		container.New(
			row.New(
				[]application.BaseWidget{
					column.New(
						[]application.BaseWidget{
							text.New(
								suggestion.Title,
								text.TextStyle(appTheme.Data().TextTheme.TextStyle.Primary),
								text.FontSize(16),
								text.FontWeight("500"),
							),
							text.New(
								suggestion.Description,
								text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
								text.FontSize(14),
								text.LineHeight(1.4),
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
			container.BorderRadius(8),
			container.BorderWidth(spacing.All(1)),
			container.BorderStyle(theme.BorderStyleTypeSolid),
			container.Width(breakpoint.All(400)),
		),
		link.Href(suggestion.Href),
	)
}

func comingSoonActions() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			text.New(
				"Need help or have questions?",
				text.FontSize(16),
				text.FontWeight("500"),
			),
			spacer.New(spacer.Height(16)),
			row.New(
				[]application.BaseWidget{
					link.New(
						button.New(
							row.New(
								[]application.BaseWidget{
									icon.New(
										icondata.Home,
										icon.Width(breakpoint.All(16)),
										icon.Height(breakpoint.All(16)),
										icon.Fill("#FFFFFF"),
									),
									text.New(
										"Back to Docs",
										text.TextStyle(appTheme.Data().ButtonTheme.ButtonStyle.Primary.TextStyle),
										text.FontSize(14),
										text.FontWeight("500"),
									),
								},
								row.Gap(8),
								row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
							),
							button.BorderRadius(6),
						),
						link.Href("/docs"),
					),
					link.New(
						button.New(
							row.New(
								[]application.BaseWidget{
									icon.New(
										icondata.AccountGroup,
										icon.Width(breakpoint.All(16)),
										icon.Height(breakpoint.All(16)),
										icon.Fill("#FFFFFF"),
									),
									text.New(
										"GitHub Discussions",
										text.TextStyle(appTheme.Data().ButtonTheme.ButtonStyle.Primary.TextStyle),
										text.FontSize(14),
										text.FontWeight("500"),
									),
								},
								row.Gap(8),
								row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
							),
							button.BorderRadius(6),
						),
						link.Href("https://github.com/orgs/gofred-io/discussions"),
						link.NewTab(true),
					),
				},
				row.Gap(12),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),
		},
		column.Gap(0),
		column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
	)
}

// Simple coming soon page without custom suggestions
func Simple(title string) application.BaseWidget {
	return ComingSoonContent(title, nil)
}

// Coming soon page with custom suggestions
func WithSuggestions(title string, suggestions []Suggestion) application.BaseWidget {
	return ComingSoonContent(title, suggestions)
}

// Coming soon page for a specific documentation section
func ForDocsSection(sectionTitle string, relatedSections []Suggestion) application.BaseWidget {
	return WithSuggestions(sectionTitle, relatedSections)
}
