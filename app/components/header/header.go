package header

import (
	. "github.com/gofred-io/gofred-website/app/components/drawer"

	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/foundation/container"
	. "github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	. "github.com/gofred-io/gofred/foundation/icon_button"
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
		Row(
			[]Widget{
				// Mobile menu button
				menuButton(),

				// Logo and title
				logoTitle(),

				// Spacer to push navigation to the right
				Spacer(),

				// Desktop navigation
				desktopNavigation(),

				Spacer().Width(24),

				// Actions section
				headerActions(),
			},
		).Gap(0).Flex(1).CrossAxisAlignment(AxisAlignmentTypeCenter),
	).Height(AllBP(72)).BackgroundColor("#FFFFFF").Padding(
		AllBP(Axis(32, 0)),
		XS(Axis(0, 0)),
		SM(Axis(8, 0)),
		MD(Axis(16, 0)),
		LG(Axis(24, 0)),
	).BorderColor("#E5E7EB").BorderWidth(0, 0, 1, 0).BorderStyle(BorderStyleTypeSolid).
		// Add subtle shadow for depth
		BoxShadow("0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06)")
}

func menuButton() Widget {
	return Container(
		IconButton(icondata.Menu).
			Width(AllBP(20)).
			Height(AllBP(20)).
			Fill("#1F2937").
			OnClick(func(this Widget, e Event) {
				RootDrawer().Show()
			}).Tooltip("Open menu"),
	).Padding(AllBP(All(8))).Visible(
		XS(true),
		SM(true),
		MD(false),
		LG(false),
	)
}

func logoTitle() Widget {
	return Container(
		Link(
			Row(
				[]Widget{
					logo(),
					title(),
				},
			).Gap(12).CrossAxisAlignment(AxisAlignmentTypeCenter),
		).Href("/"),
	).Padding(
		AllBP(All(8)),
		XS(All(4)),
	)
}

func logo() Widget {
	return Container(
		Text("🚀").FontSize(32),
	).Width(
		AllBP(40),
		XS(36),
		SM(38),
	).Height(
		AllBP(40),
		XS(36),
		SM(38),
	)
}

func title() Widget {
	return Container(
		Text("gofred").
			FontSize(24).
			FontColor("#1F2937").
			FontWeight("700").
			UserSelect(UserSelectTypeNone),
	).Visible(
		AllBP(true),
		XS(false),
	)
}

// Desktop navigation - hidden on mobile
func desktopNavigation() Widget {
	return Container(
		Row(
			[]Widget{
				navigationLink("Documentation", "/docs", false),
				navigationLink("Examples", "/docs/examples", false),
				navigationLink("Community", "https://github.com/orgs/gofred-io/discussions", true),
			},
		).Gap(32).CrossAxisAlignment(AxisAlignmentTypeCenter),
	).Visible(
		MD(true),
		LG(true),
		XL(true),
		XXL(true),
	)
}

// Header actions section
func headerActions() Widget {
	return Row(
		[]Widget{
			githubButton(),
			//Spacer().Width(12),
			//themeToggleButton(),
		},
	).Gap(0).CrossAxisAlignment(AxisAlignmentTypeCenter)
}

// Modern navigation link
func navigationLink(label, href string, external bool) Widget {
	linkWidget := Row(
		[]Widget{
			Text(label).
				FontSize(16).
				FontColor("#4B5563").
				FontWeight("500").
				UserSelect(UserSelectTypeNone),
		},
	).CrossAxisAlignment(AxisAlignmentTypeCenter)

	if external {
		linkWidget = Row(
			[]Widget{
				Text(label).
					FontSize(16).
					FontColor("#4B5563").
					FontWeight("500").
					UserSelect(UserSelectTypeNone),
				Spacer().Width(4),
				Icon(icondata.OpenInNew).
					Width(AllBP(14)).
					Height(AllBP(14)).
					Fill("#6B7280"),
			},
		).CrossAxisAlignment(AxisAlignmentTypeCenter)
	}

	return Container(
		Link(linkWidget).
			Href(href).
			NewTab(external),
	).Padding(AllBP(All(8)))
}

// GitHub button with icon
func githubButton() Widget {
	return Container(
		Link(
			Row(
				[]Widget{
					Icon(icondata.Github).
						Width(AllBP(20)).
						Height(AllBP(20)).
						Fill("#374151"),
					Container(
						Text("GitHub").
							FontSize(14).
							FontColor("#374151").
							FontWeight("600").
							UserSelect(UserSelectTypeNone),
					).Visible(
						AllBP(true),
						XS(false),
						SM(false),
					),
				},
			).Gap(8).CrossAxisAlignment(AxisAlignmentTypeCenter),
		).Href("https://github.com/gofred-io/gofred").NewTab(true),
	).Padding(AllBP(All(12))).Margin(
		AllBP(All(0)),
		XS(Right(8)),
	).BackgroundColor("#F9FAFB").BorderColor("#E5E7EB").BorderWidth(1, 1, 1, 1).BorderRadius(8).BorderStyle(BorderStyleTypeSolid)
}

// TODO: Fix theme toggle button with correct listenable API
// func themeToggleButton() Widget {
// 	currentTheme := theme.Listenable()
// 	return Builder(currentTheme, func() Widget {
// 		themeIcon := icondata.WhiteBalanceSunny
// 		themeTooltip := "Switch to dark mode"
// 		if theme.Get() == theme.ThemeDark {
// 			themeIcon = icondata.MoonWaningCrescent
// 			themeTooltip = "Switch to light mode"
// 		}

// 		return Container(
// 			Button(
// 				Icon(themeIcon).
// 					Width(AllBP(20)).
// 					Height(AllBP(20)).
// 					Fill("#6B7280"),
// 			).OnClick(func(this Widget, e Event) {
// 				if theme.Get() == theme.ThemeLight {
// 					theme.Set(theme.ThemeDark)
// 				} else {
// 					theme.Set(theme.ThemeLight)
// 				}
// 			}).Tooltip(themeTooltip),
// 		).Padding(AllBP(All(8))).BackgroundColor("#F9FAFB").BorderColor("#E5E7EB").BorderWidth(1, 1, 1, 1).BorderRadius(8).BorderStyle(BorderStyleTypeSolid)
// 	})
// }
