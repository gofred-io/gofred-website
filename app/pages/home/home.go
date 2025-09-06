package home

import (
	"github.com/gofred-io/gofred-website/app/components/footer"
	"github.com/gofred-io/gofred-website/app/components/header"
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/foundation/center"
	. "github.com/gofred-io/gofred/foundation/code_block"
	. "github.com/gofred-io/gofred/foundation/column"
	. "github.com/gofred-io/gofred/foundation/container"
	. "github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	. "github.com/gofred-io/gofred/foundation/link"
	. "github.com/gofred-io/gofred/foundation/router"
	. "github.com/gofred-io/gofred/foundation/row"
	. "github.com/gofred-io/gofred/foundation/spacer"
	. "github.com/gofred-io/gofred/foundation/text"
	. "github.com/gofred-io/gofred/options"
	. "github.com/gofred-io/gofred/options/spacing"
	. "github.com/gofred-io/gofred/widget"
)

func Home(params RouteParams) Widget {
	return Column(
		[]Widget{
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
func heroSection() Widget {
	return Container(
		Center(
			Container(
				Column(
					[]Widget{
						// Badge
						heroBadge(),
						Spacer(),

						// Main headline
						heroHeadline(),
						Spacer(),

						// Subtitle
						heroDescription(),
						Spacer(),

						// CTA buttons
						heroCTAButtons(),
						Spacer(),

						// Code preview
						heroCodePreview(),
					},
				).CrossAxisAlignment(AxisAlignmentTypeCenter).
					Gap(0),
			).MaxWidth(AllBP(1200)).
				Padding(
					AllBP(All(32)),
					XS(All(16)),
					SM(All(24)),
					MD(All(32)),
				),
		),
	).BackgroundColor("#FAFBFC").
		Flex(1)
}

func heroBadge() Widget {
	return Container(
		Row(
			[]Widget{
				Icon(icondata.RocketLaunchOutline).
					Width(AllBP(16)).
					Height(AllBP(16)).
					Fill("#2B799B"),
				Text("Build web apps with Go").
					FontSize(14).
					FontColor("#2B799B").
					FontWeight("500"),
			},
		).Gap(8).
			CrossAxisAlignment(AxisAlignmentTypeCenter),
	).Padding(AllBP(All(12))).
		BackgroundColor("#EFF6FF").
		BorderColor("#2B799B").
		BorderWidth(1, 1, 1, 1).
		BorderRadius(20).
		BorderStyle(BorderStyleTypeSolid)
}

func heroHeadline() Widget {
	return Text("Build responsive web apps in Go – no JavaScript required").
		FontSize(48).
		FontColor("#1F2937").
		FontWeight("700").
		LineHeight(1.2).
		Align(TextAlignTypeCenter)
}

func heroDescription() Widget {
	return Container(
		Text("gofred is a modern web framework that lets you build interactive, responsive web applications using only Go. Create beautiful UIs with a widget-based architecture that compiles to WebAssembly.").
			FontSize(18).
			FontColor("#6B7280").
			FontWeight("400").
			LineHeight(1.6),
	).MaxWidth(AllBP(700))
}

func heroCTAButtons() Widget {
	return Row(
		[]Widget{
			// Primary CTA
			Link(
				Container(
					Center(
						Row(
							[]Widget{
								Text("Get Started").
									FontSize(16).
									FontColor("#FFFFFF").
									FontWeight("600"),
								Icon(icondata.ChevronRight).
									Width(AllBP(16)).
									Height(AllBP(16)).
									Fill("#FFFFFF"),
							},
						).Gap(8).
							CrossAxisAlignment(AxisAlignmentTypeCenter),
					),
				).Padding(AllBP(All(20))).
					BackgroundColor("#2B799B").
					BorderRadius(8).
					Width(AllBP(182)),
			).Href("/docs"),

			// Secondary CTA
			Link(
				Container(
					Center(
						Row(
							[]Widget{
								Icon(icondata.Github).
									Width(AllBP(16)).
									Height(AllBP(16)).
									Fill("#374151"),
								Text("View on GitHub").
									FontSize(16).
									FontColor("#374151").
									FontWeight("500"),
							},
						).Gap(8).
							CrossAxisAlignment(AxisAlignmentTypeCenter),
					),
				).Width(AllBP(182)).
					Padding(AllBP(All(20))).
					BackgroundColor("#FFFFFF").
					BorderColor("#E5E7EB").
					BorderWidth(1, 1, 1, 1).
					BorderRadius(8).
					BorderStyle(BorderStyleTypeSolid),
			).Href("https://github.com/gofred-io/gofred").
				NewTab(true),
		},
	).Gap(16)
}

func heroCodePreview() Widget {
	return Container(
		Column(
			[]Widget{
				// Code header
				Container(
					Row(
						[]Widget{
							Container(Spacer()).
								Width(AllBP(12)).
								Height(AllBP(12)).
								BackgroundColor("#FF5F58").
								BorderRadius(6),
							Container(Spacer()).
								Width(AllBP(12)).
								Height(AllBP(12)).
								BackgroundColor("#FFBD2E").
								BorderRadius(6),
							Container(Spacer()).
								Width(AllBP(12)).
								Height(AllBP(12)).
								BackgroundColor("#28C941").
								BorderRadius(6),
							Spacer(),
							Text("main.go").
								FontSize(14).
								FontColor("#6B7280").
								FontWeight("500"),
						},
					).Gap(8).
						CrossAxisAlignment(AxisAlignmentTypeCenter),
				).Padding(AllBP(All(16))).
					BackgroundColor("#F9FAFB").
					BorderColor("#E5E7EB").
					BorderWidth(0, 0, 1, 0).
					BorderStyle(BorderStyleTypeSolid),

				// Code content
				CodeBlock(`package main

import (
    "github.com/gofred-io/gofred/application"
    "github.com/gofred-io/gofred/foundation/text"
)

func main() {
    app := text.New("Hello, gofred!")
    application.Run(app)
}`),
			},
		).Gap(0),
	).BorderRadius(12).
		BorderColor("#E5E7EB").
		BackgroundColor("#1F2937").
		BorderWidth(1, 1, 1, 1).
		BorderStyle(BorderStyleTypeSolid).
		MaxWidth(AllBP(500)).
		Overflow(OverflowTypeHidden)
}

// Modern Features Section
func modernFeaturesSection() Widget {
	return Container(
		Center(
			Column(
				[]Widget{
					// Section header
					Center(
						Column(
							[]Widget{
								Text("Why Choose gofred?").
									FontSize(36).
									FontColor("#1F2937").
									FontWeight("700").
									Align(TextAlignTypeCenter),
								Spacer(),
								Container(
									Text("Build modern web applications with the power of Go. No JavaScript, no complex build tools – just pure Go code that runs everywhere.").
										FontSize(18).
										FontColor("#6B7280").
										FontWeight("400").
										LineHeight(1.6),
								).MaxWidth(AllBP(600)),
							},
						).CrossAxisAlignment(AxisAlignmentTypeCenter),
					),

					Spacer(),

					// Features grid
					Center(
						Container(
							Column(
								[]Widget{
									Row(
										[]Widget{
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
										},
									).Gap(24),
									Row(
										[]Widget{
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
									).Gap(24),
								},
							).Gap(24),
						).MaxWidth(AllBP(1200)),
					),
				},
			).Gap(0).
				CrossAxisAlignment(AxisAlignmentTypeCenter),
		),
	).BackgroundColor("#FFFFFF").
		Padding(
			AllBP(All(64)),
			XS(Axis(16, 24)),
			SM(Axis(24, 32)),
			MD(Axis(32, 48)),
			LG(Axis(48, 64)),
		)
}

func modernFeatureCard(iconData icondata.IconData, title, description, color string) Widget {
	return Container(
		Column(
			[]Widget{
				// Icon
				Container(
					Center(
						Icon(iconData).
							Width(AllBP(32)).
							Height(AllBP(32)).
							Fill(color),
					),
				).Padding(AllBP(All(16))).
					BackgroundColor("#F9FAFB").
					BorderRadius(12).
					Width(AllBP(64)).
					Height(AllBP(64)),

				Spacer(),

				// Title
				Text(title).
					FontSize(20).
					FontColor("#1F2937").
					FontWeight("600"),

				Spacer(),

				// Description
				Text(description).
					FontSize(16).
					FontColor("#6B7280").
					FontWeight("400").
					LineHeight(1.6),
			},
		).Gap(0),
	).Padding(AllBP(All(24))).
		BackgroundColor("#FFFFFF").
		BorderColor("#E5E7EB").
		BorderWidth(1, 1, 1, 1).
		BorderRadius(12).
		BorderStyle(BorderStyleTypeSolid)
}

// Getting Started Section
func gettingStartedSection() Widget {
	return Container(
		Center(
			Container(
				Column(
					[]Widget{
						// Section header
						Column(
							[]Widget{
								Text("Get Started in Minutes").
									FontSize(36).
									FontColor("#1F2937").
									FontWeight("700").
									Align(TextAlignTypeCenter),
								Spacer(),
								Container(
									Text("Follow these simple steps to create your first gofred application.").
										FontSize(18).
										FontColor("#6B7280").
										FontWeight("400"),
								).MaxWidth(AllBP(600)),
							},
						).CrossAxisAlignment(AxisAlignmentTypeCenter),

						Spacer(),

						// Steps
						Column(
							[]Widget{
								gettingStartedStep("1", "Install", "Add gofred to your Go project", "go get github.com/gofred-io/gofred"),
								gettingStartedStep("2", "Create", "Write your first application", "func main() {\n   app := text.New(\"Hello!\") \n   application.Run(app) \n}"),
								gettingStartedStep("3", "Run", "Start the development server", "go run server/server.go"),
							},
						).Gap(32),

						Spacer(),

						// CTA
						Link(
							Container(
								Row(
									[]Widget{
										Text("View Full Tutorial").
											FontSize(16).
											FontColor("#FFFFFF").
											FontWeight("600").
											Align(TextAlignTypeCenter),
										Icon(icondata.ChevronRight).
											Width(AllBP(16)).
											Height(AllBP(16)).
											Fill("#FFFFFF"),
									},
								).Gap(8).
									CrossAxisAlignment(AxisAlignmentTypeCenter),
							).Padding(AllBP(All(20))).
								BackgroundColor("#2B799B").
								BorderRadius(8),
						).Href("/docs/first-app"),
					},
				).Gap(0).
					CrossAxisAlignment(AxisAlignmentTypeCenter),
			).MaxWidth(AllBP(800)).
				Padding(
					XS(All(16)),
					SM(All(24)),
					MD(All(32)),
				),
		),
	).BackgroundColor("#F9FAFB").
		Padding(
			AllBP(All(64)),
			XS(All(12)),
			SM(All(16)),
			MD(All(32)),
		)
}

func gettingStartedStep(number, title, description, code string) Widget {
	return Row(
		[]Widget{
			// Step number
			Container(
				Center(
					Text(number).
						FontSize(18).
						FontColor("#FFFFFF").
						FontWeight("700"),
				),
			).Width(AllBP(40)).
				Height(AllBP(40)).
				BackgroundColor("#2B799B").
				BorderRadius(20).
				Visible(
					AllBP(true),
					XS(false),
					SM(false),
				),

			// Content
			Column(
				[]Widget{
					Text(title).
						FontSize(20).
						FontColor("#1F2937").
						FontWeight("600"),
					Spacer(),
					Text(description).
						FontSize(16).
						FontColor("#6B7280").
						FontWeight("400"),
					Spacer(),
					CodeBlock(code),
				},
			).Gap(0).
				Flex(1),
		},
	).Gap(24).
		CrossAxisAlignment(AxisAlignmentTypeStart)
}

// Community Section
func communitySection() Widget {
	return Container(
		Center(
			Container(
				Column(
					[]Widget{
						// Section header
						Column(
							[]Widget{
								Text("Join the Community").
									FontSize(36).
									FontColor("#1F2937").
									FontWeight("700").
									Align(TextAlignTypeCenter),
								Spacer(),
								Container(
									Text("Connect with other gofred developers, get help, and contribute to the future of Go web development.").
										FontSize(18).
										FontColor("#6B7280").
										FontWeight("400"),
								).MaxWidth(AllBP(600)),
							},
						).CrossAxisAlignment(AxisAlignmentTypeCenter),

						Spacer(),

						// Community links
						Row(
							[]Widget{
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
						).Gap(24),
					},
				).Gap(0).
					CrossAxisAlignment(AxisAlignmentTypeCenter),
			).MaxWidth(AllBP(1000)).
				Padding(
					XS(All(16)),
					SM(All(24)),
					MD(All(32)),
				),
		),
	).BackgroundColor("#FFFFFF").
		Padding(
			AllBP(All(64)),
			XS(All(8)),
			SM(All(16)),
			MD(All(24)),
			LG(All(32)),
		)
}

func communityLink(iconData icondata.IconData, title, description, href string, newTab bool) Widget {
	return Container(
		Center(
			Link(
				Center(
					Container(
						Column(
							[]Widget{
								// Icon
								Container(
									Center(
										Icon(iconData).
											Width(AllBP(24)).
											Height(AllBP(24)).
											Fill("#2B799B"),
									),
								).Padding(AllBP(All(12))).
									BackgroundColor("#EFF6FF").
									BorderRadius(8).
									Width(AllBP(48)).
									Height(AllBP(48)),

								Spacer(),

								// Title
								Text(title).
									FontSize(18).
									FontColor("#1F2937").
									FontWeight("600"),

								Spacer(),

								// Description
								Text(description).
									FontSize(14).
									FontColor("#6B7280").
									FontWeight("400").
									LineHeight(1.5),
							},
						).Gap(0).
							Flex(1).
							CrossAxisAlignment(AxisAlignmentTypeCenter),
					).Padding(AllBP(All(24))).
						Flex(1),
				),
			).Href(href).
				NewTab(newTab),
		),
	).BackgroundColor("#FFFFFF").
		BorderColor("#E5E7EB").
		BorderWidth(1, 1, 1, 1).
		BorderRadius(8).
		BorderStyle(BorderStyleTypeSolid).
		Flex(1)
}
