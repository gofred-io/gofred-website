package home

import (
	"github.com/gofred-io/gofred-website/app/components/codeblock"
	"github.com/gofred-io/gofred-website/app/components/footer"
	"github.com/gofred-io/gofred-website/app/components/header"
	appTheme "github.com/gofred-io/gofred-website/app/theme"

	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/button"
	"github.com/gofred-io/gofred/foundation/center"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/grid"
	"github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	"github.com/gofred-io/gofred/foundation/link"
	"github.com/gofred-io/gofred/foundation/router"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
)

func New(params router.RouteParams) application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			// Header - fixed at top
			header.Get(),

			// Hero Section - enhanced
			heroSection(),

			// Features Section - redesigned
			modernFeaturesSection(),

			// Getting Started Section - new
			gettingStartedSection(),

			// Community Section - new
			communitySection(),

			// Footer
			footer.Get(),
		},
	)
}

// Enhanced Hero Section
func heroSection() application.BaseWidget {

	return container.New(
		center.New(
			container.New(
				column.New(
					[]application.BaseWidget{
						// Badge
						heroBadge(),
						spacer.New(spacer.Height(24)),

						// Main headline
						heroHeadline(),
						spacer.New(spacer.Height(16)),

						// Subtitle
						heroDescription(),
						spacer.New(spacer.Height(32)),

						// CTA buttons
						heroCTAButtons(),
						spacer.New(spacer.Height(48)),

						// Code preview
						heroCodePreview(),
					},
					column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
					column.Gap(0),
				),
				container.MaxWidth(breakpoint.All(1200)),
				container.Padding(
					breakpoint.All(spacing.All(32)),
					breakpoint.XS(spacing.All(16)),
					breakpoint.SM(spacing.All(24)),
					breakpoint.MD(spacing.All(32)),
				),
				container.ContainerStyle(appTheme.Data().BoxTheme.ContainerStyle.Secondary),
			),
		),
		container.ContainerStyle(appTheme.Data().BoxTheme.ContainerStyle.Secondary),
		container.Flex(1),
	)
}

func heroBadge() application.BaseWidget {
	return container.New(
		row.New(
			[]application.BaseWidget{
				icon.New(
					icondata.RocketLaunchOutline,
					icon.Width(breakpoint.All(16)),
					icon.Height(breakpoint.All(16)),
					icon.Fill("#2B799B"),
				),
				text.New(
					"Build web apps with Go",
					text.FontSize(14),
					text.FontColor("#2B799B"),
					text.FontWeight("500"),
				),
			},
			row.Gap(8),
			row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
		),
		container.Padding(breakpoint.All(spacing.All(12))),
		container.BackgroundColor("#EFF6FF"),
		container.BorderColor("#2B799B"),
		container.BorderWidth(spacing.All(1)),
		container.BorderRadius(20),
		container.BorderStyle(theme.BorderStyleTypeSolid),
	)
}

func heroHeadline() application.BaseWidget {
	return text.New(
		"Build responsive web apps in Go – no JavaScript required",
		text.FontSize(48),
		text.FontWeight("700"),
		text.LineHeight(1.2),
		text.Align(theme.TextAlignTypeCenter),
	)
}

func heroDescription() application.BaseWidget {

	return container.New(
		text.New(
			"gofred is a modern web framework that lets you build interactive, responsive web applications using only Go. Create beautiful UIs with a widget-based architecture that compiles to WebAssembly.",
			text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
			text.FontSize(18),
			text.LineHeight(1.6),
		),
		container.MaxWidth(breakpoint.All(700)),
		container.BackgroundColor("#00000000"),
	)
}

func heroCTAButtons() application.BaseWidget {

	return grid.New(
		[]application.BaseWidget{
			// Primary CTA
			link.New(
				button.New(
					center.New(
						row.New(
							[]application.BaseWidget{
								text.New(
									"Get Started",
									text.TextStyle(appTheme.Data().TextTheme.TextStyle.Tertiary),
									text.FontSize(16),
									text.FontWeight("600"),
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
					),
					button.BorderRadius(8),
					button.Width(breakpoint.All(182)),
					button.Padding(breakpoint.All(spacing.All(20))),
				),
				link.Href("/docs"),
			),

			// Secondary CTA
			link.New(
				button.New(
					center.New(
						row.New(
							[]application.BaseWidget{
								icon.New(
									icondata.Github,
									icon.Width(breakpoint.All(16)),
									icon.Height(breakpoint.All(16)),
									icon.Fill("#374151"),
								),
								text.New(
									"View on GitHub",
									text.TextStyle(appTheme.Data().ButtonTheme.ButtonStyle.Secondary.TextStyle),
									text.FontSize(16),
									text.FontWeight("500"),
								),
							},
							row.Gap(8),
							row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
						),
					),
					button.ButtonStyle(appTheme.Data().ButtonTheme.ButtonStyle.Secondary),
					button.Width(breakpoint.All(182)),
					button.Padding(breakpoint.All(spacing.All(20))),
					button.BorderRadius(8),
				),
				link.Href("https://github.com/gofred-io/gofred"),
				link.NewTab(true),
			),
		},
		grid.RowGap(16),
		grid.ColumnGap(16),
		grid.ColumnCount(
			breakpoint.All(2),
			breakpoint.XS(1),
		),
	)
}

func heroCodePreview() application.BaseWidget {

	return container.New(
		column.New(
			[]application.BaseWidget{
				// Code header
				container.New(
					row.New(
						[]application.BaseWidget{
							container.New(
								spacer.New(),
								container.Width(breakpoint.All(12)),
								container.Height(breakpoint.All(12)),
								container.BackgroundColor("#FF5F58"),
								container.BorderRadius(6),
							),
							container.New(
								spacer.New(),
								container.Width(breakpoint.All(12)),
								container.Height(breakpoint.All(12)),
								container.BackgroundColor("#FFBD2E"),
								container.BorderRadius(6),
							),
							container.New(
								spacer.New(),
								container.Width(breakpoint.All(12)),
								container.Height(breakpoint.All(12)),
								container.BackgroundColor("#28C941"),
								container.BorderRadius(6),
							),
							spacer.New(),
							text.New(
								"main.go",
								text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
								text.FontSize(14),
								text.FontWeight("500"),
							),
						},
						row.Gap(8),
						row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
					),
					container.Padding(breakpoint.All(spacing.All(16))),
					container.ContainerStyle(appTheme.Data().BoxTheme.ContainerStyle.Primary),
					container.BorderWidth(spacing.Bottom(1)),
					container.BorderStyle(theme.BorderStyleTypeSolid),
				),

				// Code content
				codeblock.New(
					`package main

import (
    "github.com/gofred-io/gofred/application"
    "github.com/gofred-io/gofred/foundation/text"
)

func main() {
    app := text.New("Hello, gofred!")
    application.Run(app)
}`,
				),
			},
			column.Gap(0),
		),
		container.BorderRadius(12),
		container.ContainerStyle(appTheme.Data().BoxTheme.ContainerStyle.Primary),
		container.BorderWidth(spacing.All(1)),
		container.BorderStyle(theme.BorderStyleTypeSolid),
		container.MaxWidth(breakpoint.All(500)),
		container.Overflow(theme.OverflowTypeHidden),
	)
}

// Modern Features Section
func modernFeaturesSection() application.BaseWidget {

	return container.New(
		center.New(
			column.New(
				[]application.BaseWidget{
					// Section header
					center.New(
						column.New(
							[]application.BaseWidget{
								text.New(
									"Why Choose gofred?",
									text.FontSize(36),
									text.FontWeight("700"),
									text.Align(theme.TextAlignTypeCenter),
								),
								spacer.New(spacer.Height(16)),
								container.New(
									text.New(
										"Build modern web applications with the power of Go. No JavaScript, no complex build tools – just pure Go code that runs everywhere.",
										text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
										text.FontSize(18),
										text.LineHeight(1.6),
									),
									container.MaxWidth(breakpoint.All(600)),
								),
							},
							column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
						),
					),

					spacer.New(spacer.Height(64)),

					// Features grid
					center.New(
						container.New(
							grid.New(
								[]application.BaseWidget{
									modernFeatureCard(
										icondata.RocketLaunchOutline,
										"Fast Development",
										"Write web applications in Go with hot reload and instant feedback. No build tools or complex setups required.",
										"#2B799B",
									),
									modernFeatureCard(
										icondata.LightningBoltOutline,
										"High Performance",
										"Compiled to WebAssembly for near-native performance. Your apps run fast in any modern browser.",
										"#059669",
									),
									modernFeatureCard(
										icondata.PaletteOutline,
										"Beautiful UIs",
										"Create responsive, beautiful interfaces with a widget-based architecture and powerful styling system.",
										"#7C3AED",
									),
									modernFeatureCard(
										icondata.Cellphone,
										"Mobile Ready",
										"Built-in responsive design system ensures your apps work perfectly on desktop, tablet, and mobile.",
										"#DC2626",
									),
									modernFeatureCard(
										icondata.PowerPlugOutline,
										"Easy Integration",
										"Seamlessly integrate with existing Go backends and APIs. Leverage your existing Go knowledge and libraries.",
										"#EA580C",
									),
									modernFeatureCard(
										icondata.Tools,
										"Developer Friendly",
										"Rich developer experience with comprehensive documentation, examples, and active community support.",
										"#0891B2",
									),
								},
								grid.ColumnCount(
									breakpoint.All(3),
									breakpoint.XS(1),
									breakpoint.SM(1),
									breakpoint.MD(2),
								),
								grid.ColumnGap(24),
								grid.RowGap(24),
							),
							container.MaxWidth(breakpoint.All(1200)),
						),
					),
				},
				column.Gap(0),
				column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),
		),
		container.Padding(
			breakpoint.All(spacing.All(64)),
			breakpoint.XS(spacing.Axis(16, 24)),
			breakpoint.SM(spacing.Axis(24, 32)),
			breakpoint.MD(spacing.Axis(32, 48)),
			breakpoint.LG(spacing.Axis(48, 64)),
		),
	)
}

func modernFeatureCard(iconData icondata.IconData, title, description, color string) application.BaseWidget {

	return container.New(
		column.New(
			[]application.BaseWidget{
				// Icon
				container.New(
					center.New(
						icon.New(
							iconData,
							icon.Width(breakpoint.All(32)),
							icon.Height(breakpoint.All(32)),
							icon.Fill(color),
						),
					),
					container.Padding(breakpoint.All(spacing.All(16))),
					container.BorderRadius(12),
					container.BorderWidth(spacing.All(1)),
					container.BorderStyle(theme.BorderStyleTypeSolid),
					container.Width(breakpoint.All(64)),
					container.Height(breakpoint.All(64)),
					container.ContainerStyle(appTheme.Data().BoxTheme.ContainerStyle.Secondary),
				),

				spacer.New(spacer.Height(24)),

				// Title
				text.New(
					title,
					text.FontSize(20),
					text.FontWeight("600"),
				),

				spacer.New(spacer.Height(12)),

				// Description
				text.New(
					description,
					text.FontSize(16),
					text.LineHeight(1.6),
					text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
				),
			},
			column.Gap(0),
		),
		container.Padding(breakpoint.All(spacing.All(24))),
		container.BorderWidth(spacing.All(1)),
		container.BorderRadius(12),
		container.BorderStyle(theme.BorderStyleTypeSolid),
	)
}

// Getting Started Section
func gettingStartedSection() application.BaseWidget {

	return container.New(
		center.New(
			container.New(
				column.New(
					[]application.BaseWidget{
						// Section header
						column.New(
							[]application.BaseWidget{
								text.New(
									"Get Started in Minutes",
									text.FontSize(36),
									text.FontWeight("700"),
									text.Align(theme.TextAlignTypeCenter),
								),
								spacer.New(spacer.Height(16)),
								container.New(
									text.New(
										"Follow these simple steps to create your first gofred application.",
										text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
										text.FontSize(18),
									),
									container.MaxWidth(breakpoint.All(600)),
									container.BackgroundColor("#00000000"),
								),
							},
							column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
						),

						spacer.New(spacer.Height(48)),

						// Steps
						column.New(
							[]application.BaseWidget{
								gettingStartedStep("1", "Install gofred CLI", "Install the gofred CLI tool", "curl -fsSL https://raw.githubusercontent.com/gofred-io/gofred-cli/refs/heads/master/install.sh | bash"),
								gettingStartedStep("2", "Create your first app", "Create a new Go WebAssembly application", "gofred app create my-app --package my-app"),
								gettingStartedStep("3", "Write your first application", "Write your first application", "func main() {\n   app := text.New(\"Hello!\") \n   application.Run(app) \n}"),
								gettingStartedStep("4", "Run your application", "Start the development server", "gofred app run"),
							},
							column.Gap(32),
						),

						spacer.New(spacer.Height(48)),

						// CTA
						link.New(
							button.New(
								row.New(
									[]application.BaseWidget{
										text.New(
											"View Full Tutorial",
											text.FontSize(16),
											text.FontColor("#FFFFFF"),
											text.FontWeight("600"),
											text.Align(theme.TextAlignTypeCenter),
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
								button.Padding(breakpoint.All(spacing.All(20))),
							),
							link.Href("/docs/first-app"),
						),
					},
					column.Gap(0),
					column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
				),
				container.ContainerStyle(appTheme.Data().BoxTheme.ContainerStyle.Secondary),
				container.MaxWidth(breakpoint.All(800)),
				container.Padding(
					breakpoint.XS(spacing.All(16)),
					breakpoint.SM(spacing.All(24)),
					breakpoint.MD(spacing.All(32)),
				),
			),
		),
		container.ContainerStyle(appTheme.Data().BoxTheme.ContainerStyle.Secondary),
		container.Padding(
			breakpoint.All(spacing.All(64)),
			breakpoint.XS(spacing.All(12)),
			breakpoint.SM(spacing.All(16)),
			breakpoint.MD(spacing.All(32)),
		),
	)
}

func gettingStartedStep(number, title, description, code string) application.BaseWidget {

	return row.New(
		[]application.BaseWidget{
			// Step number
			container.New(
				center.New(
					text.New(
						number,
						text.FontSize(18),
						text.FontColor("#FFFFFF"),
						text.FontWeight("700"),
					),
				),
				container.Width(breakpoint.All(40)),
				container.Height(breakpoint.All(40)),
				container.BackgroundColor("#1976d2"),
				container.BorderRadius(20),
				container.Visible(
					breakpoint.All(true),
					breakpoint.XS(false),
					breakpoint.SM(false),
				),
			),

			// Content
			column.New(
				[]application.BaseWidget{
					text.New(
						title,
						text.FontSize(20),
						text.FontWeight("600"),
					),
					spacer.New(spacer.Height(8)),
					text.New(
						description,
						text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
						text.FontSize(16),
					),
					spacer.New(spacer.Height(12)),
					codeblock.New(
						code,
					),
				},
				column.Gap(0),
				column.Flex(1),
			),
		},
		row.Gap(24),
		row.CrossAxisAlignment(theme.AxisAlignmentTypeStart),
	)
}

// Community Section
func communitySection() application.BaseWidget {

	return container.New(
		center.New(
			container.New(
				column.New(
					[]application.BaseWidget{
						// Section header
						column.New(
							[]application.BaseWidget{
								text.New(
									"Join the Community",
									text.FontSize(36),
									text.FontWeight("700"),
									text.Align(theme.TextAlignTypeCenter),
								),
								spacer.New(spacer.Height(16)),
								container.New(
									text.New(
										"Connect with other gofred developers, get help, and contribute to the future of Go web development.",
										text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
										text.FontSize(18),
									),
									container.MaxWidth(breakpoint.All(600)),
								),
							},
							column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
						),

						spacer.New(spacer.Height(48)),

						// Community links
						grid.New(
							[]application.BaseWidget{
								communityLink(
									icondata.Github,
									"GitHub",
									"Contribute code and report issues",
									"https://github.com/gofred-io/gofred",
									true,
								),
								communityLink(
									icondata.AccountGroup,
									"Discussions",
									"Ask questions and share ideas",
									"https://github.com/orgs/gofred-io/discussions",
									true,
								),
								communityLink(
									icondata.Book,
									"Documentation",
									"Learn with comprehensive guides",
									"/docs",
									false,
								),
							},
							grid.ColumnGap(24),
							grid.RowGap(24),
							grid.ColumnCount(
								breakpoint.All(3),
								breakpoint.XS(1),
								breakpoint.SM(1),
							),
						),
					},
					column.Gap(0),
					column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
				),
				container.MaxWidth(breakpoint.All(1000)),
				container.Padding(
					breakpoint.XS(spacing.All(16)),
					breakpoint.SM(spacing.All(24)),
					breakpoint.MD(spacing.All(32)),
				),
			),
		),
		container.Padding(
			breakpoint.All(spacing.All(64)),
			breakpoint.XS(spacing.All(8)),
			breakpoint.SM(spacing.All(16)),
			breakpoint.MD(spacing.All(24)),
			breakpoint.LG(spacing.All(32)),
		),
	)
}

func communityLink(iconData icondata.IconData, title, description, href string, newTab bool) application.BaseWidget {

	return container.New(
		center.New(
			link.New(
				center.New(
					container.New(
						column.New(
							[]application.BaseWidget{
								// Icon
								container.New(
									center.New(
										icon.New(
											iconData,
											icon.Width(breakpoint.All(24)),
											icon.Height(breakpoint.All(24)),
											icon.Fill("#2B799B"),
										),
									),
									container.Padding(breakpoint.All(spacing.All(12))),
									container.ContainerStyle(appTheme.Data().BoxTheme.ContainerStyle.Secondary),
									container.BorderRadius(8),
									container.BorderWidth(spacing.All(1)),
									container.BorderStyle(theme.BorderStyleTypeSolid),
									container.Width(breakpoint.All(48)),
									container.Height(breakpoint.All(48)),
								),

								spacer.New(spacer.Height(16)),

								// Title
								text.New(
									title,
									text.FontSize(18),
									text.FontWeight("600"),
								),

								spacer.New(spacer.Height(8)),

								// Description
								text.New(
									description,
									text.TextStyle(appTheme.Data().TextTheme.TextStyle.Secondary),
									text.FontSize(14),
									text.LineHeight(1.5),
								),
							},
							column.Gap(0),
							column.Flex(1),
							column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
						),
						container.Flex(1),
					),
				),
				link.Href(href),
				link.NewTab(newTab),
			),
		),
		container.BorderWidth(spacing.All(1)),
		container.BorderRadius(8),
		container.BorderStyle(theme.BorderStyleTypeSolid),
		container.Flex(1),
		container.Padding(breakpoint.All(spacing.All(24))),
	)
}
