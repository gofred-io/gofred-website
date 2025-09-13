package footer

import (
	"github.com/gofred-io/gofred/application"
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
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
)

func Get() application.BaseWidget {
	themeHook, _ := hooks.UseTheme()

	return container.New(
		column.New(
			[]application.BaseWidget{
				// Desktop footer (MD and larger)
				container.New(
					center.New(
						container.New(
							column.New(
								[]application.BaseWidget{
									// Footer sections
									footerContent(),

									spacer.New(spacer.Height(48)),

									// Social links
									socialLinks(),

									spacer.New(spacer.Height(8)),

									// Footer bottom with copyright and social links
									footerBottom(),
								},
								column.Gap(0),
							),
							container.ContainerStyle(themeHook.ThemeData().BoxTheme.ContainerStyle.Tertiary),
							container.MaxWidth(breakpoint.All(1200)),
							container.Padding(
								breakpoint.All(spacing.All(32)),
								breakpoint.MD(spacing.All(32)),
								breakpoint.LG(spacing.All(32)),
							),
						),
					),
					container.ContainerStyle(themeHook.ThemeData().BoxTheme.ContainerStyle.Tertiary),
					container.Visible(
						breakpoint.All(true),
						breakpoint.XS(false),
						breakpoint.SM(false),
					),
				),

				// Mobile footer (XS and SM)
				container.New(
					mobileFooter(),
					container.Visible(
						breakpoint.All(false),
						breakpoint.XS(true),
						breakpoint.SM(true),
					),
					container.ContainerStyle(themeHook.ThemeData().BoxTheme.ContainerStyle.Tertiary),
				),
			},
			column.Gap(0),
			column.Flex(1),
		),
		container.ContainerStyle(themeHook.ThemeData().BoxTheme.ContainerStyle.Tertiary),
	)
}

// Main footer content with different sections
func footerContent() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			// Brand section
			footerBrandSection(),

			grid.New(
				[]application.BaseWidget{
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
					breakpoint.All(3),
				),
				grid.ColumnGap(32),
				grid.RowGap(32),
			),
		},
	)
}

type FooterLink struct {
	title  string
	href   string
	newTab bool
}

// Brand section with logo and description
func footerBrandSection() application.BaseWidget {
	themeHook, _ := hooks.UseTheme()

	return column.New(
		[]application.BaseWidget{
			// Logo and brand name
			row.New(
				[]application.BaseWidget{
					text.New(
						"🚀",
						text.FontSize(32),
					),
					text.New(
						"gofred",
						text.TextStyle(themeHook.ThemeData().TextTheme.TextStyle.Tertiary),
						text.FontSize(24),
						text.FontWeight("700"),
					),
				},
				row.Gap(12),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),

			spacer.New(spacer.Height(16)),

			// Description
			text.New(
				"Build responsive web applications using only Go. No JavaScript required - just pure Go code that compiles to WebAssembly.",
				text.TextStyle(themeHook.ThemeData().TextTheme.TextStyle.Tertiary),
				text.Opacity(0.6),
				text.FontSize(14),
				text.LineHeight(1.6),
			),

			spacer.New(spacer.Height(20)),
		},
		column.Gap(0),
		column.CrossAxisAlignment(theme.AxisAlignmentTypeStart),
	)
}

// Links section with title and list of links
func footerLinksSection(title string, links []FooterLink) application.BaseWidget {
	themeHook, _ := hooks.UseTheme()

	var linkWidgets []application.BaseWidget

	for _, link := range links {
		linkWidgets = append(linkWidgets, footerLink(link.title, link.href, link.newTab))
	}

	return column.New(
		[]application.BaseWidget{
			// Section title
			text.New(
				title,
				text.TextStyle(themeHook.ThemeData().TextTheme.TextStyle.Tertiary),
				text.FontSize(16),
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
		column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
	)
}

// Individual footer link
func footerLink(title, href string, newTab bool) application.BaseWidget {
	themeHook, _ := hooks.UseTheme()

	return link.New(
		text.New(
			title,
			text.TextStyle(themeHook.ThemeData().TextTheme.TextStyle.Tertiary),
			text.FontSize(14),
			text.Opacity(0.6),
		),
		link.Href(href),
		link.NewTab(newTab),
	)
}

// Social media links
func socialLinks() application.BaseWidget {
	return row.New(
		[]application.BaseWidget{
			spacer.New(),
			socialIcon(icondata.Github, "GitHub", "https://github.com/gofred-io/gofred"),
			socialIcon(icondata.Twitter, "Twitter", "https://twitter.com/gofred_io"),
			socialIcon(icondata.Youtube, "YouTube", "https://youtube.com/@gofred"),
			socialIcon(icondata.Instagram, "Instagram", "https://instagram.com/gofred_io"),
		},
		row.Gap(16),
		row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
	)
}

// Individual social icon
func socialIcon(iconData icondata.IconData, tooltip, href string) application.BaseWidget {
	themeHook, _ := hooks.UseTheme()

	return link.New(
		container.New(
			icon.New(
				iconData,
				icon.Width(breakpoint.All(20)),
				icon.Height(breakpoint.All(20)),
				icon.Fill("#9CA3AF"),
			),
			container.Padding(breakpoint.All(spacing.All(8))),
			container.ContainerStyle(themeHook.ThemeData().BoxTheme.ContainerStyle.Tertiary),
		),
		link.Href(href),
		link.NewTab(true),
	)
}

// Footer bottom section with copyright and additional info
func footerBottom() application.BaseWidget {
	themeHook, _ := hooks.UseTheme()

	return container.New(
		row.New(
			[]application.BaseWidget{
				// Copyright and legal
				column.New(
					[]application.BaseWidget{
						text.New(
							"© 2025 gofred. All rights reserved.",
							text.TextStyle(themeHook.ThemeData().TextTheme.TextStyle.Tertiary),
							text.Opacity(0.4),
							text.FontSize(14),
						),

						spacer.New(spacer.Height(8)),

						// Legal links
						row.New(
							[]application.BaseWidget{
								footerLink("Privacy Policy", "/privacy", false),
								text.New("•", text.TextStyle(themeHook.ThemeData().TextTheme.TextStyle.Tertiary), text.FontSize(14)),
								footerLink("Terms of Service", "/terms", false),
								text.New("•", text.TextStyle(themeHook.ThemeData().TextTheme.TextStyle.Tertiary), text.FontSize(14)),
								footerLink("License", "/license", false),
							},
							row.Gap(8),
							row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
						),
					},
					column.Gap(0),
					column.CrossAxisAlignment(theme.AxisAlignmentTypeStart),
				),

				spacer.New(),

				// Built with love
				row.New(
					[]application.BaseWidget{
						text.New(
							"Built with",
							text.TextStyle(themeHook.ThemeData().TextTheme.TextStyle.Tertiary),
							text.FontSize(14),
							text.Opacity(0.4),
						),
						icon.New(
							icondata.Heart,
							icon.Width(breakpoint.All(16)),
							icon.Height(breakpoint.All(16)),
							icon.Fill("#EF4444"),
						),
						text.New(
							"using gofred",
							text.TextStyle(themeHook.ThemeData().TextTheme.TextStyle.Tertiary),
							text.FontSize(14),
							text.Opacity(0.4),
						),
					},
					row.Gap(8),
					row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
				),
			},
			row.Flex(1),
			row.CrossAxisAlignment(theme.AxisAlignmentTypeStart),
		),
		container.Padding(
			breakpoint.All(spacing.All(24)),
		),
		container.ContainerStyle(themeHook.ThemeData().BoxTheme.ContainerStyle.Tertiary),
		container.BorderWidth(spacing.Top(1)),
		container.BorderStyle(theme.BorderStyleTypeSolid),
	)
}

// Mobile-optimized footer for smaller screens
func mobileFooter() application.BaseWidget {
	themeHook, _ := hooks.UseTheme()

	return container.New(
		center.New(
			column.New(
				[]application.BaseWidget{
					// Brand section
					center.New(
						column.New(
							[]application.BaseWidget{
								// Logo and brand
								row.New(
									[]application.BaseWidget{
										text.New(
											"🚀",
											text.FontSize(28),
										),
										text.New(
											"gofred",
											text.TextStyle(themeHook.ThemeData().TextTheme.TextStyle.Tertiary),
											text.FontSize(20),
											text.FontWeight("700"),
										),
									},
									row.Gap(10),
									row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
								),

								spacer.New(spacer.Height(12)),

								// Description
								container.New(
									text.New(
										"Build responsive web applications using only Go.",
										text.TextStyle(themeHook.ThemeData().TextTheme.TextStyle.Tertiary),
										text.FontSize(14),
									),
									container.ContainerStyle(themeHook.ThemeData().BoxTheme.ContainerStyle.Tertiary),
									container.MaxWidth(breakpoint.All(300)),
								),

								spacer.New(spacer.Height(20)),

								// Social links
								socialLinks(),
							},
							column.Gap(0),
							column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
						),
					),

					spacer.New(spacer.Height(32)),

					// Quick navigation grid
					grid.New(
						[]application.BaseWidget{
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
							[]application.BaseWidget{
								text.New(
									"© 2025 gofred",
									text.TextStyle(themeHook.ThemeData().TextTheme.TextStyle.Tertiary),
									text.FontSize(14),
								),

								spacer.New(spacer.Height(8)),

								row.New(
									[]application.BaseWidget{
										text.New(
											"Built with",
											text.TextStyle(themeHook.ThemeData().TextTheme.TextStyle.Tertiary),
											text.FontSize(12),
										),
										icon.New(
											icondata.Heart,
											icon.Width(breakpoint.All(14)),
											icon.Height(breakpoint.All(14)),
											icon.Fill("#EF4444"),
										),
										text.New(
											"using gofred",
											text.TextStyle(themeHook.ThemeData().TextTheme.TextStyle.Tertiary),
											text.FontSize(12),
										),
									},
									row.Gap(6),
									row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
								),
							},
							column.Gap(0),
							column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
						),
					),
				},
				column.Gap(0),
				column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),
		),
		container.Flex(1),
		container.Padding(breakpoint.All(spacing.All(24))),
		container.ContainerStyle(themeHook.ThemeData().BoxTheme.ContainerStyle.Tertiary),
	)
}

// Mobile footer section
func mobileFooterSection(title string, links []string) application.BaseWidget {
	themeHook, _ := hooks.UseTheme()

	var linkWidgets []application.BaseWidget
	for _, link := range links {
		linkWidgets = append(linkWidgets, text.New(
			link,
			text.TextStyle(themeHook.ThemeData().TextTheme.TextStyle.Tertiary),
			text.FontSize(14),
		))
	}

	return column.New(
		[]application.BaseWidget{
			text.New(
				title,
				text.TextStyle(themeHook.ThemeData().TextTheme.TextStyle.Tertiary),
				text.FontSize(16),
				text.FontWeight("600"),
			),

			spacer.New(spacer.Height(12)),

			column.New(
				linkWidgets,
				column.Gap(8),
			),
		},
		column.Gap(0),
		column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
	)
}
