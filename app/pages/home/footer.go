package home

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/center"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/grid"
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

func Footer() widget.BaseWidget {
	return container.New(
		column.New(
			[]widget.BaseWidget{
				// Desktop footer (MD and larger)
				container.New(
					center.New(
						container.New(
							column.New(
								[]widget.BaseWidget{
									// Footer sections
									footerContent(),

									spacer.New(spacer.Height(48)),

									// Footer bottom with copyright and social links
									footerBottom(),
								},
								column.Gap(0),
							),
							container.MaxWidth(breakpoint.All(1200)),
							container.Padding(
								breakpoint.All(spacing.All(32)),
								breakpoint.MD(spacing.All(32)),
								breakpoint.LG(spacing.All(32)),
							),
						),
					),
					container.Visible(
						breakpoint.MD(true),
						breakpoint.LG(true),
						breakpoint.XL(true),
						breakpoint.XXL(true),
					),
				),

				// Mobile footer (XS and SM)
				container.New(
					mobileFooter(),
					container.Visible(
						breakpoint.XS(true),
						breakpoint.SM(true),
						breakpoint.MD(false),
					),
				),
			},
			column.Gap(0),
			column.Flex(1),
		),
		container.BackgroundColor("#1F2937"),
	)
}

// Main footer content with different sections
func footerContent() widget.BaseWidget {
	return grid.New(
		[]widget.BaseWidget{
			// Brand section
			footerBrandSection(),

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
		grid.ColumnCount(
			breakpoint.All(4),
			breakpoint.XS(1),
			breakpoint.SM(2),
			breakpoint.MD(4),
		),
		grid.ColumnGap(32),
		grid.RowGap(32),
	)
}

type FooterLink struct {
	title  string
	href   string
	newTab bool
}

// Brand section with logo and description
func footerBrandSection() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			// Logo and brand name
			row.New(
				[]widget.BaseWidget{
					text.New(
						"🚀",
						text.FontSize(32),
					),
					text.New(
						"gofred",
						text.FontSize(24),
						text.FontColor("#FFFFFF"),
						text.FontWeight("700"),
					),
				},
				row.Gap(12),
				row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
			),

			spacer.New(spacer.Height(16)),

			// Description
			text.New(
				"Build responsive web applications using only Go. No JavaScript required - just pure Go code that compiles to WebAssembly.",
				text.FontSize(14),
				text.FontColor("#9CA3AF"),
				text.FontWeight("400"),
				text.LineHeight(1.6),
			),

			spacer.New(spacer.Height(20)),

			// Social links
			socialLinks(),
		},
		column.Gap(0),
		column.CrossAxisAlignment(options.AxisAlignmentTypeStart),
	)
}

// Links section with title and list of links
func footerLinksSection(title string, links []FooterLink) widget.BaseWidget {
	var linkWidgets []widget.BaseWidget

	for _, link := range links {
		linkWidgets = append(linkWidgets, footerLink(link.title, link.href, link.newTab))
	}

	return column.New(
		[]widget.BaseWidget{
			// Section title
			text.New(
				title,
				text.FontSize(16),
				text.FontColor("#FFFFFF"),
				text.FontWeight("600"),
			),

			spacer.New(spacer.Height(16)),

			// Links list
			column.New(
				linkWidgets,
				column.Gap(12),
			),
		},
		column.Gap(0),
		column.CrossAxisAlignment(options.AxisAlignmentTypeStart),
	)
}

// Individual footer link
func footerLink(title, href string, newTab bool) widget.BaseWidget {
	return link.New(
		text.New(
			title,
			text.FontSize(14),
			text.FontColor("#9CA3AF"),
			text.FontWeight("400"),
		),
		link.Href(href),
		link.NewTab(newTab),
	)
}

// Social media links
func socialLinks() widget.BaseWidget {
	return row.New(
		[]widget.BaseWidget{
			socialIcon(icondata.Github, "GitHub", "https://github.com/gofred-io/gofred"),
			socialIcon(icondata.Twitter, "Twitter", "https://twitter.com/gofred_io"),
			socialIcon(icondata.Youtube, "YouTube", "https://youtube.com/@gofred"),
			socialIcon(icondata.Instagram, "Instagram", "https://instagram.com/gofred_io"),
		},
		row.Gap(16),
		row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
	)
}

// Individual social icon
func socialIcon(iconData icondata.IconData, tooltip, href string) widget.BaseWidget {
	return link.New(
		container.New(
			icon.New(
				iconData,
				icon.Width(breakpoint.All(20)),
				icon.Height(breakpoint.All(20)),
				icon.Fill("#9CA3AF"),
			),
			container.Padding(breakpoint.All(spacing.All(8))),
		),
		link.Href(href),
		link.NewTab(true),
	)
}

// Footer bottom section with copyright and additional info
func footerBottom() widget.BaseWidget {
	return container.New(
		row.New(
			[]widget.BaseWidget{
				// Copyright and legal
				column.New(
					[]widget.BaseWidget{
						text.New(
							"© 2025 gofred. All rights reserved.",
							text.FontSize(14),
							text.FontColor("#6B7280"),
							text.FontWeight("400"),
						),

						spacer.New(spacer.Height(8)),

						// Legal links
						row.New(
							[]widget.BaseWidget{
								footerLink("Privacy Policy", "/privacy", false),
								text.New("•", text.FontSize(14), text.FontColor("#6B7280")),
								footerLink("Terms of Service", "/terms", false),
								text.New("•", text.FontSize(14), text.FontColor("#6B7280")),
								footerLink("License", "/license", false),
							},
							row.Gap(8),
							row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
						),
					},
					column.Gap(0),
					column.CrossAxisAlignment(options.AxisAlignmentTypeStart),
				),

				spacer.New(),

				// Built with love
				row.New(
					[]widget.BaseWidget{
						text.New(
							"Built with",
							text.FontSize(14),
							text.FontColor("#6B7280"),
							text.FontWeight("400"),
						),
						icon.New(
							icondata.Heart,
							icon.Width(breakpoint.All(16)),
							icon.Height(breakpoint.All(16)),
							icon.Fill("#EF4444"),
						),
						text.New(
							"using gofred",
							text.FontSize(14),
							text.FontColor("#6B7280"),
							text.FontWeight("400"),
						),
					},
					row.Gap(8),
					row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
				),
			},
			row.Flex(1),
			row.CrossAxisAlignment(options.AxisAlignmentTypeStart),
		),
		container.Padding(
			breakpoint.All(spacing.All(24)),
		),
		container.BorderColor("#374151"),
		container.BorderWidth(1, 0, 0, 0),
		container.BorderStyle(options.BorderStyleTypeSolid),
	)
}

// Mobile-optimized footer for smaller screens
func mobileFooter() widget.BaseWidget {
	return container.New(
		column.New(
			[]widget.BaseWidget{
				// Brand section
				center.New(
					column.New(
						[]widget.BaseWidget{
							// Logo and brand
							row.New(
								[]widget.BaseWidget{
									text.New(
										"🚀",
										text.FontSize(28),
									),
									text.New(
										"gofred",
										text.FontSize(20),
										text.FontColor("#FFFFFF"),
										text.FontWeight("700"),
									),
								},
								row.Gap(10),
								row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
							),

							spacer.New(spacer.Height(12)),

							// Description
							container.New(
								text.New(
									"Build responsive web applications using only Go.",
									text.FontSize(14),
									text.FontColor("#9CA3AF"),
									text.FontWeight("400"),
								),
								container.MaxWidth(breakpoint.All(300)),
							),

							spacer.New(spacer.Height(20)),

							// Social links
							socialLinks(),
						},
						column.Gap(0),
						column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
					),
				),

				spacer.New(spacer.Height(32)),

				// Quick navigation grid
				grid.New(
					[]widget.BaseWidget{
						mobileFooterSection("Docs", []string{"Getting Started", "Examples", "API"}),
						mobileFooterSection("Community", []string{"GitHub", "Discussions", "Support"}),
					},
					grid.ColumnCount(breakpoint.All(2)),
					grid.ColumnGap(24),
				),

				spacer.New(spacer.Height(24)),

				// Copyright
				center.New(
					column.New(
						[]widget.BaseWidget{
							text.New(
								"© 2025 gofred",
								text.FontSize(14),
								text.FontColor("#6B7280"),
								text.FontWeight("400"),
							),

							spacer.New(spacer.Height(8)),

							row.New(
								[]widget.BaseWidget{
									text.New(
										"Built with",
										text.FontSize(12),
										text.FontColor("#6B7280"),
										text.FontWeight("400"),
									),
									icon.New(
										icondata.Heart,
										icon.Width(breakpoint.All(14)),
										icon.Height(breakpoint.All(14)),
										icon.Fill("#EF4444"),
									),
									text.New(
										"using gofred",
										text.FontSize(12),
										text.FontColor("#6B7280"),
										text.FontWeight("400"),
									),
								},
								row.Gap(6),
								row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
							),
						},
						column.Gap(0),
						column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
					),
				),
			},
			column.Gap(0),
		),
		container.Padding(breakpoint.All(spacing.All(24))),
	)
}

// Mobile footer section
func mobileFooterSection(title string, links []string) widget.BaseWidget {
	var linkWidgets []widget.BaseWidget
	for _, link := range links {
		linkWidgets = append(linkWidgets, text.New(
			link,
			text.FontSize(14),
			text.FontColor("#9CA3AF"),
			text.FontWeight("400"),
		))
	}

	return column.New(
		[]widget.BaseWidget{
			text.New(
				title,
				text.FontSize(16),
				text.FontColor("#FFFFFF"),
				text.FontWeight("600"),
			),

			spacer.New(spacer.Height(12)),

			column.New(
				linkWidgets,
				column.Gap(8),
			),
		},
		column.Gap(0),
		column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
	)
}
