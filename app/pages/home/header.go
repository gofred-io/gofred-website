package home

import (
	"github.com/gofred-io/gofred-website/app/theme"

	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/drawer"
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
	"github.com/gofred-io/gofred/widget"
)

var (
	leftDrawer = drawer.New(
		drawer.Width(breakpoint.All(250)),
	)
)

func header() widget.BaseWidget {
	return container.New(
		row.New(
			[]widget.BaseWidget{
				menuButton(),
				logoTitle(),
				spacer.New(),
				documentationLink(),
				spacer.New(spacer.Width(8)),
				discussionsLink(),
				spacer.New(spacer.Width(8)),
				githubLink(),
				spacer.New(spacer.Width(8)),
				themeToggleButton(),
			},
			row.Gap(8),
			row.Flex(1),
			row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
		),
		container.Height(breakpoint.All(44)),
		container.BackgroundColor("#FFFFFF"),
		container.PaddingV(breakpoint.All(8)),
		container.PaddingH(breakpoint.All(16)),
	)
}

func menuButton() widget.BaseWidget {
	return iconbutton.New(
		icondata.Menu,
		iconbutton.OnClick(func(this widget.BaseWidget, e widget.Event) {
			leftDrawer.Show()
		}),
		iconbutton.Fill("#003B73"),
		iconbutton.Tooltip("Menu"),
		iconbutton.Visible(
			breakpoint.XS(true),
			breakpoint.SM(true),
		),
	)
}

func logoTitle() widget.BaseWidget {
	return link.New(
		row.New(
			[]widget.BaseWidget{
				logo(),
				title(),
			},
			row.Gap(8),
			row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
		),
		link.Href("/"),
	)
}

func logo() widget.BaseWidget {
	return image.New(
		"img/gofred.png",
		image.Width(breakpoint.All(32)),
		image.Height(breakpoint.All(32)),
	)
}

func title() widget.BaseWidget {
	return text.New(
		"gofred",
		text.FontSize(20),
		text.FontColor("#003B73"),
		text.FontWeight("500"),
		text.UserSelect(options.UserSelectTypeNone),
	)
}

func documentationLink() widget.BaseWidget {
	return container.New(
		link.New(
			text.New(
				"Docs",
				text.FontSize(16),
				text.FontColor("#2B799B"),
				text.FontWeight("500"),
				text.UserSelect(options.UserSelectTypeNone),
			),
			link.Href("/docs"),
		),
		container.Visible(
			breakpoint.All(true),
			breakpoint.XS(false),
			breakpoint.SM(false),
		),
	)
}

func discussionsLink() widget.BaseWidget {
	return container.New(
		link.New(
			row.New(
				[]widget.BaseWidget{
					text.New(
						"Discussions",
						text.FontSize(16),
						text.FontColor("#2B799B"),
						text.FontWeight("500"),
						text.UserSelect(options.UserSelectTypeNone),
					),
					icon.New(
						icondata.OpenInNew,
						icon.Width(breakpoint.All(18)),
						icon.Height(breakpoint.All(18)),
						icon.Fill("#2B799B"),
					),
				},
			),
			link.Href("https://github.com/gofred-io/gofred/pulls"),
			link.NewTab(true),
		),
		container.Visible(
			breakpoint.All(true),
			breakpoint.XS(false),
			breakpoint.SM(false),
		),
	)
}

func githubLink() widget.BaseWidget {
	return container.New(
		link.New(
			row.New(
				[]widget.BaseWidget{
					text.New(
						"GitHub",
						text.FontSize(16),
						text.FontColor("#2B799B"),
						text.FontWeight("500"),
						text.UserSelect(options.UserSelectTypeNone),
					),
					icon.New(
						icondata.OpenInNew,
						icon.Width(breakpoint.All(18)),
						icon.Height(breakpoint.All(18)),
						icon.Fill("#2B799B"),
					),
				},
			),
			link.Href("https://github.com/gofred-io/gofred"),
			link.NewTab(true),
		),
		container.Visible(
			breakpoint.All(true),
			breakpoint.XS(false),
			breakpoint.SM(false),
		),
	)
}

func themeToggleButton() widget.BaseWidget {
	currentTheme := theme.Listenable()
	return listenable.Builder(currentTheme, func() widget.BaseWidget {
		themeIcon := icondata.WhiteBalanceSunny
		if theme.Get() == theme.ThemeDark {
			themeIcon = icondata.MoonWaningCrescent
		}

		return iconbutton.New(
			themeIcon,
			iconbutton.Fill("#2B799B"),
			iconbutton.OnClick(func(this widget.BaseWidget, e widget.Event) {
				if theme.Get() == theme.ThemeLight {
					theme.Set(theme.ThemeDark)
				} else {
					theme.Set(theme.ThemeLight)
				}
			}),
			iconbutton.Tooltip("Toggle theme"),
		)
	})
}
