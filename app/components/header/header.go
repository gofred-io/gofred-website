package header

import (
	"github.com/gofred-io/gofred-website/app/components/drawer"
	"github.com/gofred-io/gofred-website/app/components/snackbar"
	"github.com/gofred-io/gofred-website/app/constant"
	"github.com/gofred-io/gofred-website/app/theme"

	"github.com/gofred-io/gofred/basic/button"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	iconbutton "github.com/gofred-io/gofred/foundation/icon_button"
	"github.com/gofred-io/gofred/foundation/image"
	"github.com/gofred-io/gofred/foundation/link"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/widget"
)

func Get() widget.BaseWidget {
	return container.New(
		row.New(
			[]widget.BaseWidget{
				// Mobile menu button
				menuButton(),

				// Logo and title
				logoTitle(),

				// Spacer to push navigation to the right
				spacer.New(),

				// Desktop navigation
				desktopNavigation(),

				spacer.New(spacer.Width(24)),

				// Actions section
				headerActions(),
			},
			row.Gap(0),
			row.Flex(1),
			row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
		),
		container.Height(breakpoint.All(72)),
		container.BackgroundColor("#FFFFFF"),
		container.Padding(
			breakpoint.All(spacing.Axis(32, 0)),
			breakpoint.XS(spacing.Axis(0, 0)),
			breakpoint.SM(spacing.Axis(8, 0)),
			breakpoint.MD(spacing.Axis(16, 0)),
			breakpoint.LG(spacing.Axis(24, 0)),
		),
		container.BorderColor("#E5E7EB"),
		container.BorderWidth(0, 0, 1, 0),
		container.BorderStyle(options.BorderStyleTypeSolid),
		// Add subtle shadow for depth
		container.BoxShadow("0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06)"),
	)
}

func menuButton() widget.BaseWidget {
	return container.New(
		iconbutton.New(
			icondata.Menu,
			iconbutton.OnClick(func(this widget.BaseWidget, e widget.Event) {
				drawer.Get().Show()
			}),
			iconbutton.Fill("#1F2937"),
			iconbutton.Tooltip("Open menu"),
		),
		container.Padding(breakpoint.All(spacing.All(8))),
		container.Visible(
			breakpoint.XS(true),
			breakpoint.SM(true),
			breakpoint.MD(false),
			breakpoint.LG(false),
		),
	)
}

func logoTitle() widget.BaseWidget {
	return container.New(
		link.New(
			row.New(
				[]widget.BaseWidget{
					logo(),
					title(),
					button.New(
						text.New("show snackbar"),
						button.OnClick(func(this widget.BaseWidget, e widget.Event) {
							snackbar.Show("Drawer closed", constant.SnackbarTypeSuccess)
						}),
					),
				},
				row.Gap(12),
				row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
			),
			link.Href("/"),
		),
		container.Padding(
			breakpoint.All(spacing.All(8)),
			breakpoint.XS(spacing.All(4)),
		),
	)
}

func logo() widget.BaseWidget {
	return image.New(
		"img/gofred.png",
		image.Width(
			breakpoint.All(40),
			breakpoint.XS(36),
			breakpoint.SM(38),
		),
		image.Height(
			breakpoint.All(40),
			breakpoint.XS(36),
			breakpoint.SM(38),
		),
	)
}

func title() widget.BaseWidget {
	return container.New(
		text.New(
			"gofred",
			text.FontSize(24),
			text.FontColor("#1F2937"),
			text.FontWeight("700"),
			text.UserSelect(options.UserSelectTypeNone),
		),
		container.Visible(
			breakpoint.All(true),
			breakpoint.XS(false),
		),
	)
}

// Desktop navigation - hidden on mobile
func desktopNavigation() widget.BaseWidget {
	return container.New(
		row.New(
			[]widget.BaseWidget{
				navigationLink("Documentation", "/docs", false),
				navigationLink("Examples", "/docs/examples", false),
				navigationLink("Community", "https://github.com/orgs/gofred-io/discussions", true),
			},
			row.Gap(32),
			row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
		),
		container.Visible(
			breakpoint.MD(true),
			breakpoint.LG(true),
			breakpoint.XL(true),
			breakpoint.XXL(true),
		),
	)
}

// Header actions section
func headerActions() widget.BaseWidget {
	return row.New(
		[]widget.BaseWidget{
			githubButton(),
			//spacer.New(spacer.Width(12)),
			//themeToggleButton(),
		},
		row.Gap(0),
		row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
	)
}

// Modern navigation link
func navigationLink(label, href string, external bool) widget.BaseWidget {
	linkWidget := row.New(
		[]widget.BaseWidget{
			text.New(
				label,
				text.FontSize(16),
				text.FontColor("#4B5563"),
				text.FontWeight("500"),
				text.UserSelect(options.UserSelectTypeNone),
			),
		},
		row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
	)

	if external {
		linkWidget = row.New(
			[]widget.BaseWidget{
				text.New(
					label,
					text.FontSize(16),
					text.FontColor("#4B5563"),
					text.FontWeight("500"),
					text.UserSelect(options.UserSelectTypeNone),
				),
				spacer.New(spacer.Width(4)),
				icon.New(
					icondata.OpenInNew,
					icon.Width(breakpoint.All(14)),
					icon.Height(breakpoint.All(14)),
					icon.Fill("#6B7280"),
				),
			},
			row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
		)
	}

	return container.New(
		link.New(
			linkWidget,
			link.Href(href),
			link.NewTab(external),
		),
		container.Padding(breakpoint.All(spacing.All(8))),
	)
}

// GitHub button with icon
func githubButton() widget.BaseWidget {
	return container.New(
		link.New(
			row.New(
				[]widget.BaseWidget{
					icon.New(
						icondata.Github,
						icon.Width(breakpoint.All(20)),
						icon.Height(breakpoint.All(20)),
						icon.Fill("#374151"),
					),
					container.New(
						text.New(
							"GitHub",
							text.FontSize(14),
							text.FontColor("#374151"),
							text.FontWeight("600"),
							text.UserSelect(options.UserSelectTypeNone),
						),
						container.Visible(
							breakpoint.All(true),
							breakpoint.XS(false),
							breakpoint.SM(false),
						),
					),
				},
				row.Gap(8),
				row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
			),
			link.Href("https://github.com/gofred-io/gofred"),
			link.NewTab(true),
		),
		container.Padding(breakpoint.All(spacing.All(12))),
		container.Margin(
			breakpoint.All(spacing.All(0)),
			breakpoint.XS(spacing.Right(8)),
		),
		container.BackgroundColor("#F9FAFB"),
		container.BorderColor("#E5E7EB"),
		container.BorderWidth(1, 1, 1, 1),
		container.BorderRadius(8),
		container.BorderStyle(options.BorderStyleTypeSolid),
	)
}

func themeToggleButton() widget.BaseWidget {
	currentTheme := theme.Listenable()
	return listenable.Builder(currentTheme, func() widget.BaseWidget {
		themeIcon := icondata.WhiteBalanceSunny
		themeTooltip := "Switch to dark mode"
		if theme.Get() == theme.ThemeDark {
			themeIcon = icondata.MoonWaningCrescent
			themeTooltip = "Switch to light mode"
		}

		return container.New(
			iconbutton.New(
				themeIcon,
				iconbutton.Fill("#6B7280"),
				iconbutton.OnClick(func(this widget.BaseWidget, e widget.Event) {
					if theme.Get() == theme.ThemeLight {
						theme.Set(theme.ThemeDark)
					} else {
						theme.Set(theme.ThemeLight)
					}
				}),
				iconbutton.Tooltip(themeTooltip),
			),
			container.Padding(breakpoint.All(spacing.All(8))),
			container.BackgroundColor("#F9FAFB"),
			container.BorderColor("#E5E7EB"),
			container.BorderWidth(1, 1, 1, 1),
			container.BorderRadius(8),
			container.BorderStyle(options.BorderStyleTypeSolid),
		)
	})
}
