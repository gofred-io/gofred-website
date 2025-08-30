package home

import (
	"fmt"

	"github.com/gofred-io/gofred-website/app/theme"
	"github.com/gofred-io/gofred/container"
	"github.com/gofred-io/gofred/icon"
	iconbutton "github.com/gofred-io/gofred/icon_button"
	icondata "github.com/gofred-io/gofred/icon_data"
	"github.com/gofred-io/gofred/image"
	"github.com/gofred-io/gofred/link"
	"github.com/gofred-io/gofred/listenable"
	"github.com/gofred-io/gofred/row"
	"github.com/gofred-io/gofred/spacer"
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/style/breakpoint"
	"github.com/gofred-io/gofred/text"
	"github.com/gofred-io/gofred/widget"
)

func header() widget.Widget {
	return container.New(
		row.New(
			[]widget.Widget{
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
			row.CrossAxisAlignment(style.AlignItemsTypeCenter),
		),
		container.Style(
			container.Height(44),
			container.Background(style.Background{
				Color: "#FFFFFF",
			}),
			container.Padding(style.Padding{
				Top:    8,
				Bottom: 8,
				Left:   16,
				Right:  16,
			}),
		),
	)
}

func menuButton() widget.Widget {
	return iconbutton.New(
		icondata.Menu,
		iconbutton.OnClick(func(this widget.Widget) {
			fmt.Println("hamburger menu clicked")
		}),
		iconbutton.Style(
			iconbutton.Fill("#003B73"),
		),
		iconbutton.Visible(
			breakpoint.XS(true),
			breakpoint.SM(true),
		),
	)
}

func logoTitle() widget.Widget {
	return link.New(
		row.New(
			[]widget.Widget{
				logo(),
				title(),
			},
			row.Gap(8),
		),
		link.Href("/"),
	)
}

func logo() widget.Widget {
	return image.New(
		"img/gofred.png",
		image.Width(32),
		image.Height(32),
	)
}

func title() widget.Widget {
	return text.New(
		"gofred",
		text.Style(
			text.Font(
				text.Size(20),
				text.Color("#003B73"),
				text.Weight("500"),
			),
			text.UserSelect(style.UserSelectTypeNone),
		),
	)
}

func documentationLink() widget.Widget {
	return container.New(
		link.New(
			text.New(
				"Docs",
				text.Style(
					text.Font(text.Size(16), text.Color("#2B799B"), text.Weight("500")),
				),
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

func discussionsLink() widget.Widget {
	return container.New(
		link.New(
			row.New(
				[]widget.Widget{
					text.New(
						"Discussions",
						text.Style(
							text.Font(text.Size(16), text.Color("#2B799B"), text.Weight("500")),
						),
					),
					icon.New(
						icondata.OpenInNew,
						icon.Style(
							icon.Width(18),
							icon.Height(18),
							icon.Fill("#2B799B"),
						),
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

func githubLink() widget.Widget {
	return container.New(
		link.New(
			row.New(
				[]widget.Widget{
					text.New(
						"GitHub",
						text.Style(
							text.Font(text.Size(16), text.Color("#2B799B"), text.Weight("500")),
						),
					),
					icon.New(
						icondata.OpenInNew,
						icon.Style(
							icon.Width(18),
							icon.Height(18),
							icon.Fill("#2B799B"),
						),
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

func themeToggleButton() widget.Widget {
	currentTheme := theme.Listenable()
	return listenable.Builder(currentTheme, func() widget.Widget {
		themeIcon := icondata.WhiteBalanceSunny
		if theme.Get() == theme.ThemeDark {
			themeIcon = icondata.MoonWaningCrescent
		}

		return iconbutton.New(
			themeIcon,
			iconbutton.Style(
				iconbutton.Fill("#2B799B"),
			),
			iconbutton.OnClick(func(this widget.Widget) {
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
