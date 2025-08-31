package home

import (
	"fmt"

	"github.com/gofred-io/gofred-website/app/theme"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/container"
	"github.com/gofred-io/gofred/icon"
	iconbutton "github.com/gofred-io/gofred/icon_button"
	icondata "github.com/gofred-io/gofred/icon_data"
	"github.com/gofred-io/gofred/image"
	"github.com/gofred-io/gofred/link"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/row"
	"github.com/gofred-io/gofred/spacer"
	"github.com/gofred-io/gofred/text"
	"github.com/gofred-io/gofred/widget"
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
		options.Height(breakpoint.All(44)),
		options.BackgroundColor("#FFFFFF"),
		options.PaddingV(breakpoint.All(8)),
		options.PaddingH(breakpoint.All(16)),
	)
}

func menuButton() widget.BaseWidget {
	return iconbutton.New(
		icondata.Menu,
		options.OnClick(func(this widget.BaseWidget) {
			fmt.Println("hamburger menu clicked")
		}),
		options.Fill("#003B73"),
		options.Tooltip("Menu"),
		options.Visible(
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
		options.Href("/"),
	)
}

func logo() widget.BaseWidget {
	return image.New(
		"img/gofred.png",
		options.Width(breakpoint.All(32)),
		options.Height(breakpoint.All(32)),
	)
}

func title() widget.BaseWidget {
	return text.New(
		"gofred",
		options.FontSize(20),
		options.FontColor("#003B73"),
		options.FontWeight("500"),
		options.UserSelect(options.UserSelectTypeNone),
	)
}

func documentationLink() widget.BaseWidget {
	return container.New(
		link.New(
			text.New(
				"Docs",
				options.FontSize(16),
				options.FontColor("#2B799B"),
				options.FontWeight("500"),
				options.UserSelect(options.UserSelectTypeNone),
			),
			options.Href("/docs"),
		),
		options.Visible(
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
						options.FontSize(16),
						options.FontColor("#2B799B"),
						options.FontWeight("500"),
						options.UserSelect(options.UserSelectTypeNone),
					),
					icon.New(
						icondata.OpenInNew,
						options.Width(breakpoint.All(18)),
						options.Height(breakpoint.All(18)),
						options.Fill("#2B799B"),
					),
				},
			),
			options.Href("https://github.com/gofred-io/gofred/pulls"),
			options.NewTab(true),
		),
		options.Visible(
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
						options.FontSize(16),
						options.FontColor("#2B799B"),
						options.FontWeight("500"),
						options.UserSelect(options.UserSelectTypeNone),
					),
					icon.New(
						icondata.OpenInNew,
						options.Width(breakpoint.All(18)),
						options.Height(breakpoint.All(18)),
						options.Fill("#2B799B"),
					),
				},
			),
			options.Href("https://github.com/gofred-io/gofred"),
			options.NewTab(true),
		),
		options.Visible(
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
			options.Fill("#2B799B"),
			options.OnClick(func(this widget.BaseWidget) {
				if theme.Get() == theme.ThemeLight {
					theme.Set(theme.ThemeDark)
				} else {
					theme.Set(theme.ThemeLight)
				}
			}),
			options.Tooltip("Toggle theme"),
		)
	})
}
