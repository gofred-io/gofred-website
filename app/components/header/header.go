package header

import (
	"github.com/gofred-io/gofred-website/app/components/drawer"
	webtheme "github.com/gofred-io/gofred-website/app/theme"

	"github.com/gofred-io/gofred/application"
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
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
)

func Get() application.BaseWidget {
	themeHook, _ := hooks.UseTheme()
	return container.New(
		row.New(
			[]application.BaseWidget{
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
			row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
		),
		container.Height(breakpoint.All(72)),
		container.Padding(
			breakpoint.All(spacing.Axis(32, 0)),
			breakpoint.XS(spacing.Axis(0, 0)),
			breakpoint.SM(spacing.Axis(8, 0)),
			breakpoint.MD(spacing.Axis(16, 0)),
			breakpoint.LG(spacing.Axis(24, 0)),
		),
		container.BorderWidth(spacing.Bottom(1)),
		container.BorderStyle(theme.BorderStyleTypeSolid),
		container.ContainerStyle(themeHook.ThemeData().BoxTheme.ContainerStyle.Primary),
	)
}

func menuButton() application.BaseWidget {
	return container.New(
		iconbutton.New(
			icondata.Menu,
			iconbutton.OnClick(func(this application.BaseWidget, e application.Event) {
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

func logoTitle() application.BaseWidget {
	return container.New(
		link.New(
			row.New(
				[]application.BaseWidget{
					logo(),
					title(),
				},
				row.Gap(12),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),
			link.Href("/"),
		),
		container.Padding(
			breakpoint.All(spacing.All(8)),
			breakpoint.XS(spacing.All(4)),
		),
	)
}

func logo() application.BaseWidget {
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

func title() application.BaseWidget {
	return container.New(
		text.New(
			"gofred",
			text.FontSize(24),
			text.FontWeight("700"),
			text.UserSelect(theme.UserSelectTypeNone),
		),
		container.Visible(
			breakpoint.All(true),
			breakpoint.XS(false),
		),
	)
}

// Desktop navigation - hidden on mobile
func desktopNavigation() application.BaseWidget {
	return container.New(
		row.New(
			[]application.BaseWidget{
				navigationLink("Documentation", "/docs", false),
				navigationLink("Examples", "/docs/examples", false),
				navigationLink("Community", "https://github.com/orgs/gofred-io/discussions", true),
			},
			row.Gap(32),
			row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
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
func headerActions() application.BaseWidget {
	return row.New(
		[]application.BaseWidget{
			//githubButton(),
			//spacer.New(spacer.Width(12)),
			themeToggleButton(),
		},
		row.Gap(0),
		row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
	)
}

// Modern navigation link
func navigationLink(label, href string, external bool) application.BaseWidget {
	themeHook, _ := hooks.UseTheme()

	linkWidget := row.New(
		[]application.BaseWidget{
			text.New(
				label,
				text.FontSize(16),
				text.FontWeight("500"),
				text.UserSelect(theme.UserSelectTypeNone),
				text.TextStyle(themeHook.ThemeData().TextTheme.TextStyle.Primary),
			),
		},
		row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
	)

	if external {
		linkWidget = row.New(
			[]application.BaseWidget{
				text.New(
					label,
					text.FontSize(16),
					text.FontWeight("500"),
					text.UserSelect(theme.UserSelectTypeNone),
					text.TextStyle(themeHook.ThemeData().TextTheme.TextStyle.Primary),
				),
				spacer.New(spacer.Width(4)),
				icon.New(
					icondata.OpenInNew,
					icon.Width(breakpoint.All(14)),
					icon.Height(breakpoint.All(14)),
					icon.Fill("#6B7280"),
				),
			},
			row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
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
func githubButton() application.BaseWidget {
	themeHook, _ := hooks.UseTheme()

	return container.New(
		link.New(
			row.New(
				[]application.BaseWidget{
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
							text.FontWeight("600"),
							text.UserSelect(theme.UserSelectTypeNone),
							text.TextStyle(themeHook.ThemeData().TextTheme.TextStyle.Primary),
						),
						container.BackgroundColor("#00000000"),
						container.Visible(
							breakpoint.All(true),
							breakpoint.XS(false),
							breakpoint.SM(false),
						),
					),
				},
				row.Gap(8),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),
			link.Href("https://github.com/gofred-io/gofred"),
			link.NewTab(true),
		),
		container.Padding(breakpoint.All(spacing.All(12))),
		container.Margin(
			breakpoint.All(spacing.All(0)),
			breakpoint.XS(spacing.Right(8)),
		),
		container.ContainerStyle(themeHook.ThemeData().BoxTheme.ContainerStyle.Secondary),
		container.BorderWidth(spacing.All(1)),
		container.BorderRadius(8),
		container.BorderStyle(theme.BorderStyleTypeSolid),
	)
}

func themeToggleButton() application.BaseWidget {
	themeHook, setThemeData := hooks.UseTheme()
	return listenable.Builder(themeHook, func() application.BaseWidget {
		themeIcon := icondata.WhiteBalanceSunny
		themeTooltip := "Switch to dark mode"
		themeData := themeHook.ThemeData()
		if themeData.Name == string(webtheme.ThemeDark) {
			themeIcon = icondata.MoonWaningCrescent
			themeTooltip = "Switch to light mode"
		}

		return container.New(
			iconbutton.New(
				themeIcon,
				iconbutton.Fill("#6B7280"),
				iconbutton.OnClick(func(this application.BaseWidget, e application.Event) {
					if themeData.Name == string(webtheme.ThemeLight) {
						setThemeData(webtheme.DarkTheme())
					} else {
						setThemeData(webtheme.LightTheme())
					}
				}),
				iconbutton.Tooltip(themeTooltip),
			),
			container.Padding(breakpoint.All(spacing.All(8))),
			container.BorderRadius(8),
		)
	})
}
