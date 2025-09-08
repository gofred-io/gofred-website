package home

import (
	"github.com/gofred-io/gofred-website/app/components/codeblock"
	"github.com/gofred-io/gofred-website/app/components/footer"
	"github.com/gofred-io/gofred-website/app/components/header"
	"github.com/gofred-io/gofred/basic/pre"
	"github.com/gofred-io/gofred/breakpoint"
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
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/widget"
)

func New(params router.RouteParams) widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
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
func heroSection() widget.BaseWidget {
	return container.New(
		center.New(
			container.New(
				column.New(
					[]widget.BaseWidget{
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
					column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
					column.Gap(0),
				),
				container.MaxWidth(breakpoint.All(1200)),
				container.Padding(
					breakpoint.All(spacing.All(32)),
					breakpoint.XS(spacing.All(16)),
					breakpoint.SM(spacing.All(24)),
					breakpoint.MD(spacing.All(32)),
				),
			),
		),

		container.BackgroundColor("#FAFBFC"),
		container.Flex(1),
	)
}

func heroBadge() widget.BaseWidget {
	return container.New(
		row.New(
			[]widget.BaseWidget{
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
			row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
		),
		container.Padding(breakpoint.All(spacing.All(12))),
		container.BackgroundColor("#EFF6FF"),
		container.BorderColor("#2B799B"),
		container.BorderWidth(1, 1, 1, 1),
		container.BorderRadius(20),
		container.BorderStyle(options.BorderStyleTypeSolid),
	)
}

func heroHeadline() widget.BaseWidget {
	return text.New(
		"Build responsive web apps in Go – no JavaScript required",
		text.FontSize(48),
		text.FontColor("#1F2937"),
		text.FontWeight("700"),
		text.LineHeight(1.2),
		text.Align(options.TextAlignTypeCenter),
	)
}

func heroDescription() widget.BaseWidget {
	return container.New(
		text.New(
			"gofred is a modern web framework that lets you build interactive, responsive web applications using only Go. Create beautiful UIs with a widget-based architecture that compiles to WebAssembly.",
			text.FontSize(18),
			text.FontColor("#6B7280"),
			text.FontWeight("400"),
			text.LineHeight(1.6),
		),
		container.MaxWidth(breakpoint.All(700)),
	)
}

func heroCTAButtons() widget.BaseWidget {
	return grid.New(
		[]widget.BaseWidget{
			// Primary CTA
			link.New(
				container.New(
					center.New(
						row.New(
							[]widget.BaseWidget{
								text.New(
									"Get Started",
									text.FontSize(16),
									text.FontColor("#FFFFFF"),
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
							row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
						),
					),
					container.Padding(breakpoint.All(spacing.All(20))),
					container.BackgroundColor("#2B799B"),
					container.BorderRadius(8),
					container.Width(breakpoint.All(182)),
				),
				link.Href("/docs"),
			),

			// Secondary CTA
			link.New(
				container.New(
					center.New(
						row.New(
							[]widget.BaseWidget{
								icon.New(
									icondata.Github,
									icon.Width(breakpoint.All(16)),
									icon.Height(breakpoint.All(16)),
									icon.Fill("#374151"),
								),
								text.New(
									"View on GitHub",
									text.FontSize(16),
									text.FontColor("#374151"),
									text.FontWeight("500"),
								),
							},
							row.Gap(8),
							row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
						),
					),
					container.Width(breakpoint.All(182)),
					container.Padding(breakpoint.All(spacing.All(20))),
					container.BackgroundColor("#FFFFFF"),
					container.BorderColor("#E5E7EB"),
					container.BorderWidth(1, 1, 1, 1),
					container.BorderRadius(8),
					container.BorderStyle(options.BorderStyleTypeSolid),
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

func heroCodePreview() widget.BaseWidget {
	return container.New(
		column.New(
			[]widget.BaseWidget{
				// Code header
				container.New(
					row.New(
						[]widget.BaseWidget{
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
								text.FontSize(14),
								text.FontColor("#6B7280"),
								text.FontWeight("500"),
							),
						},
						row.Gap(8),
						row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
					),
					container.Padding(breakpoint.All(spacing.All(16))),
					container.BackgroundColor("#F9FAFB"),
					container.BorderColor("#E5E7EB"),
					container.BorderWidth(0, 0, 1, 0),
					container.BorderStyle(options.BorderStyleTypeSolid),
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
		container.BorderColor("#E5E7EB"),
		container.BackgroundColor("#1F2937"),
		container.BorderWidth(1, 1, 1, 1),
		container.BorderStyle(options.BorderStyleTypeSolid),
		container.MaxWidth(breakpoint.All(500)),
		container.Overflow(options.OverflowTypeHidden),
	)
}

func codeLine(content, color string) widget.BaseWidget {
	return pre.New(
		text.New(
			content,
			text.FontSize(14),
			text.FontColor(color),
			text.FontWeight("400"),
		),
	)
}

func codeLineIndented(content, color string) widget.BaseWidget {
	return pre.New(
		row.New(
			[]widget.BaseWidget{
				text.New("  ", text.FontSize(14), text.FontColor("#E5E7EB")),
				text.New(
					content,
					text.FontSize(14),
					text.FontColor(color),
					text.FontWeight("400"),
				),
			},
			row.Gap(0),
		),
	)
}

func codeLineEmpty() widget.BaseWidget {
	return spacer.New(spacer.Height(4))
}

// Modern Features Section
func modernFeaturesSection() widget.BaseWidget {
	return container.New(
		center.New(
			column.New(
				[]widget.BaseWidget{
					// Section header
					center.New(
						column.New(
							[]widget.BaseWidget{
								text.New(
									"Why Choose gofred?",
									text.FontSize(36),
									text.FontColor("#1F2937"),
									text.FontWeight("700"),
									text.Align(options.TextAlignTypeCenter),
								),
								spacer.New(spacer.Height(16)),
								container.New(
									text.New(
										"Build modern web applications with the power of Go. No JavaScript, no complex build tools – just pure Go code that runs everywhere.",
										text.FontSize(18),
										text.FontColor("#6B7280"),
										text.FontWeight("400"),
										text.LineHeight(1.6),
									),
									container.MaxWidth(breakpoint.All(600)),
								),
							},
							column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
						),
					),

					spacer.New(spacer.Height(64)),

					// Features grid
					center.New(
						container.New(
							grid.New(
								[]widget.BaseWidget{
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
				column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
			),
		),
		container.BackgroundColor("#FFFFFF"),
		container.Padding(
			breakpoint.All(spacing.All(64)),
			breakpoint.XS(spacing.Axis(16, 24)),
			breakpoint.SM(spacing.Axis(24, 32)),
			breakpoint.MD(spacing.Axis(32, 48)),
			breakpoint.LG(spacing.Axis(48, 64)),
		),
	)
}

func modernFeatureCard(iconData icondata.IconData, title, description, color string) widget.BaseWidget {
	return container.New(
		column.New(
			[]widget.BaseWidget{
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
					container.BackgroundColor("#F9FAFB"),
					container.BorderRadius(12),
					container.Width(breakpoint.All(64)),
					container.Height(breakpoint.All(64)),
				),

				spacer.New(spacer.Height(24)),

				// Title
				text.New(
					title,
					text.FontSize(20),
					text.FontColor("#1F2937"),
					text.FontWeight("600"),
				),

				spacer.New(spacer.Height(12)),

				// Description
				text.New(
					description,
					text.FontSize(16),
					text.FontColor("#6B7280"),
					text.FontWeight("400"),
					text.LineHeight(1.6),
				),
			},
			column.Gap(0),
		),
		container.Padding(breakpoint.All(spacing.All(24))),
		container.BackgroundColor("#FFFFFF"),
		container.BorderColor("#E5E7EB"),
		container.BorderWidth(1, 1, 1, 1),
		container.BorderRadius(12),
		container.BorderStyle(options.BorderStyleTypeSolid),
	)
}

// Getting Started Section
func gettingStartedSection() widget.BaseWidget {
	return container.New(
		center.New(
			container.New(
				column.New(
					[]widget.BaseWidget{
						// Section header
						column.New(
							[]widget.BaseWidget{
								text.New(
									"Get Started in Minutes",
									text.FontSize(36),
									text.FontColor("#1F2937"),
									text.FontWeight("700"),
									text.Align(options.TextAlignTypeCenter),
								),
								spacer.New(spacer.Height(16)),
								container.New(
									text.New(
										"Follow these simple steps to create your first gofred application.",
										text.FontSize(18),
										text.FontColor("#6B7280"),
										text.FontWeight("400"),
									),
									container.MaxWidth(breakpoint.All(600)),
								),
							},
							column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
						),

						spacer.New(spacer.Height(48)),

						// Steps
						column.New(
							[]widget.BaseWidget{
								gettingStartedStep("1", "Install", "Add gofred to your Go project", "go get github.com/gofred-io/gofred"),
								gettingStartedStep("2", "Create", "Write your first application", "func main() {\n   app := text.New(\"Hello!\") \n   application.Run(app) \n}"),
								gettingStartedStep("3", "Run", "Start the development server", "go run server/server.go"),
							},
							column.Gap(32),
						),

						spacer.New(spacer.Height(48)),

						// CTA
						link.New(
							container.New(
								row.New(
									[]widget.BaseWidget{
										text.New(
											"View Full Tutorial",
											text.FontSize(16),
											text.FontColor("#FFFFFF"),
											text.FontWeight("600"),
											text.Align(options.TextAlignTypeCenter),
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
								container.Padding(breakpoint.All(spacing.All(20))),
								container.BackgroundColor("#2B799B"),
								container.BorderRadius(8),
							),
							link.Href("/docs/first-app"),
						),
					},
					column.Gap(0),
					column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
				),
				container.MaxWidth(breakpoint.All(800)),
				container.Padding(
					breakpoint.XS(spacing.All(16)),
					breakpoint.SM(spacing.All(24)),
					breakpoint.MD(spacing.All(32)),
				),
			),
		),
		container.BackgroundColor("#F9FAFB"),
		container.Padding(
			breakpoint.All(spacing.All(64)),
			breakpoint.XS(spacing.All(12)),
			breakpoint.SM(spacing.All(16)),
			breakpoint.MD(spacing.All(32)),
		),
	)
}

func gettingStartedStep(number, title, description, code string) widget.BaseWidget {
	return row.New(
		[]widget.BaseWidget{
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
				container.BackgroundColor("#2B799B"),
				container.BorderRadius(20),
				container.Visible(
					breakpoint.All(true),
					breakpoint.XS(false),
					breakpoint.SM(false),
				),
			),

			// Content
			column.New(
				[]widget.BaseWidget{
					text.New(
						title,
						text.FontSize(20),
						text.FontColor("#1F2937"),
						text.FontWeight("600"),
					),
					spacer.New(spacer.Height(8)),
					text.New(
						description,
						text.FontSize(16),
						text.FontColor("#6B7280"),
						text.FontWeight("400"),
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
		row.CrossAxisAlignment(options.AxisAlignmentTypeStart),
	)
}

// Community Section
func communitySection() widget.BaseWidget {
	return container.New(
		center.New(
			container.New(
				column.New(
					[]widget.BaseWidget{
						// Section header
						column.New(
							[]widget.BaseWidget{
								text.New(
									"Join the Community",
									text.FontSize(36),
									text.FontColor("#1F2937"),
									text.FontWeight("700"),
									text.Align(options.TextAlignTypeCenter),
								),
								spacer.New(spacer.Height(16)),
								container.New(
									text.New(
										"Connect with other gofred developers, get help, and contribute to the future of Go web development.",
										text.FontSize(18),
										text.FontColor("#6B7280"),
										text.FontWeight("400"),
									),
									container.MaxWidth(breakpoint.All(600)),
								),
							},
							column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
						),

						spacer.New(spacer.Height(48)),

						// Community links
						grid.New(
							[]widget.BaseWidget{
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
					column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
				),
				container.MaxWidth(breakpoint.All(1000)),
				container.Padding(
					breakpoint.XS(spacing.All(16)),
					breakpoint.SM(spacing.All(24)),
					breakpoint.MD(spacing.All(32)),
				),
			),
		),
		container.BackgroundColor("#FFFFFF"),
		container.Padding(
			breakpoint.All(spacing.All(64)),
			breakpoint.XS(spacing.All(8)),
			breakpoint.SM(spacing.All(16)),
			breakpoint.MD(spacing.All(24)),
			breakpoint.LG(spacing.All(32)),
		),
	)
}

func communityLink(iconData icondata.IconData, title, description, href string, newTab bool) widget.BaseWidget {
	return container.New(
		center.New(
			link.New(
				center.New(
					container.New(
						column.New(
							[]widget.BaseWidget{
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
									container.BackgroundColor("#EFF6FF"),
									container.BorderRadius(8),
									container.Width(breakpoint.All(48)),
									container.Height(breakpoint.All(48)),
								),

								spacer.New(spacer.Height(16)),

								// Title
								text.New(
									title,
									text.FontSize(18),
									text.FontColor("#1F2937"),
									text.FontWeight("600"),
								),

								spacer.New(spacer.Height(8)),

								// Description
								text.New(
									description,
									text.FontSize(14),
									text.FontColor("#6B7280"),
									text.FontWeight("400"),
									text.LineHeight(1.5),
								),
							},
							column.Gap(0),
							column.Flex(1),
							column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
						),
						container.Padding(breakpoint.All(spacing.All(24))),
						container.Flex(1),
					),
				),
				link.Href(href),
				link.NewTab(newTab),
			),
		),
		container.BackgroundColor("#FFFFFF"),
		container.BorderColor("#E5E7EB"),
		container.BorderWidth(1, 1, 1, 1),
		container.BorderRadius(8),
		container.BorderStyle(options.BorderStyleTypeSolid),
		container.Flex(1),
	)
}
