package footer

import (
	. "github.com/gofred-io/gofred/breakpoint"
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

func Get() Widget {
	return Container(
		Column(
			[]Widget{
				// Desktop footer (MD and larger)
				Container(
					Center(
						Container(
							Column(
								[]Widget{
									// Footer sections
									footerContent(),

									Spacer().Height(48),

									// Social links
									socialLinks(),

									Spacer().Height(8),

									// Footer bottom with copyright and social links
									footerBottom(),
								},
							).Gap(0),
						).MaxWidth(AllBP(1200)).Padding(
							AllBP(All(32)),
							MD(All(32)),
							LG(All(32)),
						),
					),
				).Visible(
					AllBP(true),
					XS(false),
					SM(false),
				),

				// Mobile footer (XS and SM)
				Container(
					mobileFooter(),
				).Visible(
					AllBP(false),
					XS(true),
					SM(true),
				),
			},
		).Gap(0).Flex(1),
	).BackgroundColor("#1F2937")
}

// Main footer content with different sections
func footerContent() Widget {
	return Column(
		[]Widget{
			// Brand section
			footerBrandSection(),

			Row(
				[]Widget{
					// Quick links
					footerLinksSection("Quick Links", []FooterLink{
						{title: "Documentation", href: "/docs"},
						{title: "Getting Started", href: "/docs/installation"},
						{title: "Examples", href: "/docs/examples"},
						{title: "API Reference", href: "/docs/api"},
					}),

					// Community section
					footerLinksSection("Community", []FooterLink{
						{title: "GitHub", href: "https://github.com/gofred-io/gofred", newTab: true},
						{title: "Discussions", href: "https://github.com/orgs/gofred-io/discussions", newTab: true},
						{title: "Issues", href: "https://github.com/gofred-io/gofred/issues", newTab: true},
						{title: "Contributions", href: "https://github.com/gofred-io/gofred/blob/main/CONTRIBUTING.md", newTab: true},
					}),

					// Resources section
					footerLinksSection("Resources", []FooterLink{
						{title: "Blog", href: "/blog"},
						{title: "Tutorials", href: "/docs/tutorials"},
						{title: "Best Practices", href: "/docs/best-practices"},
						{title: "Support", href: "/docs/support"},
					}),
				},
			).Gap(32),
		},
	)
}

type FooterLink struct {
	title  string
	href   string
	newTab bool
}

// Brand section with logo and description
func footerBrandSection() Widget {
	return Column(
		[]Widget{
			// Logo and brand name
			Row(
				[]Widget{
					Text("🚀").FontSize(32),
					Text("gofred").
						FontSize(24).
						FontColor("#FFFFFF").
						FontWeight("700"),
				},
			).Gap(12).CrossAxisAlignment(AxisAlignmentTypeCenter),

			Spacer().Height(16),

			// Description
			Text(
				"Build responsive web applications using only Go. No JavaScript required - just pure Go code that compiles to WebAssembly.",
			).FontSize(14).
				FontColor("#9CA3AF").
				FontWeight("400").
				LineHeight(1.6),

			Spacer().Height(20),
		},
	).Gap(0).CrossAxisAlignment(AxisAlignmentTypeStart)
}

// Links section with title and list of links
func footerLinksSection(title string, links []FooterLink) Widget {
	var linkWidgets []Widget

	for _, link := range links {
		linkWidgets = append(linkWidgets, footerLink(link.title, link.href, link.newTab))
	}

	return Column(
		[]Widget{
			// Section title
			Text(title).
				FontSize(16).
				FontColor("#FFFFFF").
				FontWeight("600"),

			Spacer().Height(16),

			// Links list
			Column(
				linkWidgets,
			).Gap(12),
		},
	).Gap(0).CrossAxisAlignment(AxisAlignmentTypeCenter)
}

// Individual footer link
func footerLink(title, href string, newTab bool) Widget {
	return Link(
		Text(title).
			FontSize(14).
			FontColor("#9CA3AF").
			FontWeight("400"),
	).Href(href).NewTab(newTab)
}

// Social media links
func socialLinks() Widget {
	return Row(
		[]Widget{
			Spacer(),
			socialIcon(icondata.Github, "GitHub", "https://github.com/gofred-io/gofred"),
			socialIcon(icondata.Twitter, "Twitter", "https://twitter.com/gofred_io"),
			socialIcon(icondata.Youtube, "YouTube", "https://youtube.com/@gofred"),
			socialIcon(icondata.Instagram, "Instagram", "https://instagram.com/gofred_io"),
		},
	).Gap(16).CrossAxisAlignment(AxisAlignmentTypeCenter)
}

// Individual social icon
func socialIcon(iconData icondata.IconData, tooltip, href string) Widget {
	return Link(
		Container(
			Icon(iconData).
				Width(AllBP(20)).
				Height(AllBP(20)).
				Fill("#9CA3AF"),
		).Padding(AllBP(All(8))),
	).Href(href).NewTab(true)
}

// Footer bottom section with copyright and additional info
func footerBottom() Widget {
	return Container(
		Row(
			[]Widget{
				// Copyright and legal
				Column(
					[]Widget{
						Text("© 2025 gofred. All rights reserved.").
							FontSize(14).
							FontColor("#6B7280").
							FontWeight("400"),

						Spacer().Height(8),

						// Legal links
						Row(
							[]Widget{
								footerLink("Privacy Policy", "/privacy", false),
								Text("•").FontSize(14).FontColor("#6B7280"),
								footerLink("Terms of Service", "/terms", false),
								Text("•").FontSize(14).FontColor("#6B7280"),
								footerLink("License", "/license", false),
							},
						).Gap(8).CrossAxisAlignment(AxisAlignmentTypeCenter),
					},
				).Gap(0).CrossAxisAlignment(AxisAlignmentTypeStart),

				Spacer(),

				// Built with love
				Row(
					[]Widget{
						Text("Built with").
							FontSize(14).
							FontColor("#6B7280").
							FontWeight("400"),
						Icon(icondata.Heart).
							Width(AllBP(16)).
							Height(AllBP(16)).
							Fill("#EF4444"),
						Text("using gofred").
							FontSize(14).
							FontColor("#6B7280").
							FontWeight("400"),
					},
				).Gap(8).CrossAxisAlignment(AxisAlignmentTypeCenter),
			},
		).Flex(1).CrossAxisAlignment(AxisAlignmentTypeStart),
	).Padding(
		AllBP(All(24)),
	).BorderColor("#374151").BorderWidth(1, 0, 0, 0).BorderStyle(BorderStyleTypeSolid)
}

// Mobile-optimized footer for smaller screens
func mobileFooter() Widget {
	return Container(
		Center(
			Column(
				[]Widget{
					// Brand section
					Center(
						Column(
							[]Widget{
								// Logo and brand
								Row(
									[]Widget{
										Text("🚀").FontSize(28),
										Text("gofred").
											FontSize(20).
											FontColor("#FFFFFF").
											FontWeight("700"),
									},
								).Gap(10).CrossAxisAlignment(AxisAlignmentTypeCenter),

								Spacer().Height(12),

								// Description
								Container(
									Text("Build responsive web applications using only Go.").
										FontSize(14).
										FontColor("#9CA3AF").
										FontWeight("400"),
								).MaxWidth(AllBP(300)),

								Spacer().Height(20),

								// Social links
								socialLinks(),
							},
						).Gap(0).CrossAxisAlignment(AxisAlignmentTypeCenter),
					),

					Spacer().Height(32),

					// Quick navigation grid
					Row(
						[]Widget{
							mobileFooterSection("Docs", []string{"Getting Started", "Examples", "API"}),
							mobileFooterSection("Community", []string{"GitHub", "Discussions", "Support"}),
						},
					).Gap(24),

					Spacer().Height(24),

					// Copyright
					Center(
						Column(
							[]Widget{
								Text("© 2025 gofred").
									FontSize(14).
									FontColor("#6B7280").
									FontWeight("400"),

								Spacer().Height(8),

								Row(
									[]Widget{
										Text("Built with").
											FontSize(12).
											FontColor("#6B7280").
											FontWeight("400"),
										Icon(icondata.Heart).
											Width(AllBP(14)).
											Height(AllBP(14)).
											Fill("#EF4444"),
										Text("using gofred").
											FontSize(12).
											FontColor("#6B7280").
											FontWeight("400"),
									},
								).Gap(6).CrossAxisAlignment(AxisAlignmentTypeCenter),
							},
						).Gap(0).CrossAxisAlignment(AxisAlignmentTypeCenter),
					),
				},
			).Gap(0).CrossAxisAlignment(AxisAlignmentTypeCenter),
		),
	).Flex(1).Padding(AllBP(All(24)))
}

// Mobile footer section
func mobileFooterSection(title string, links []string) Widget {
	var linkWidgets []Widget
	for _, link := range links {
		linkWidgets = append(linkWidgets, Text(link).
			FontSize(14).
			FontColor("#9CA3AF").
			FontWeight("400"),
		)
	}

	return Column(
		[]Widget{
			Text(title).
				FontSize(16).
				FontColor("#FFFFFF").
				FontWeight("600"),

			Spacer().Height(12),

			Column(
				linkWidgets,
			).Gap(8),
		},
	).Gap(0).CrossAxisAlignment(AxisAlignmentTypeCenter)
}
