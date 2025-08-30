package home

import (
	"github.com/gofred-io/gofred/container"
	iconbutton "github.com/gofred-io/gofred/icon_button"
	icondata "github.com/gofred-io/gofred/icon_data"
	"github.com/gofred-io/gofred/row"
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/text"
	"github.com/gofred-io/gofred/widget"
)

func footer() widget.Widget {
	return container.New(
		row.New(
			[]widget.Widget{
				socials(),
				copyright(),
				githubButton(),
			},
		),
		container.Style(
			container.Background(style.Background{
				Color: "#23395B",
			}),
			container.Padding(style.Padding{
				Top:    16,
				Bottom: 16,
				Left:   16,
				Right:  16,
			}),
		),
	)
}

func socials() widget.Widget {
	twitterButton := iconbutton.New(
		icondata.Twitter,
		iconbutton.Style(
			iconbutton.Fill("#FFFFFF"),
		),
		iconbutton.Tooltip("Follow us on Twitter"),
	)

	youtubeButton := iconbutton.New(
		icondata.Youtube,
		iconbutton.Style(
			iconbutton.Fill("#FFFFFF"),
		),
		iconbutton.Tooltip("Subscribe to our YouTube channel"),
	)

	instagramButton := iconbutton.New(
		icondata.Instagram,
		iconbutton.Style(
			iconbutton.Fill("#FFFFFF"),
		),
		iconbutton.Tooltip("Follow us on Instagram"),
	)

	return row.New(
		[]widget.Widget{
			twitterButton,
			youtubeButton,
			instagramButton,
		},
		row.MainAxisAlignment(style.JustifyContentTypeStart),
		row.CrossAxisAlignment(style.AlignItemsTypeCenter),
		row.Gap(8),
	)
}

func copyright() widget.Widget {
	return row.New(
		[]widget.Widget{
			text.New("Copyright © 2025 Gofred", text.Style(text.Font(text.Size(14), text.Color("#FFFFFF"), text.Weight("400")))),
		},
		row.MainAxisAlignment(style.JustifyContentTypeCenter),
		row.CrossAxisAlignment(style.AlignItemsTypeCenter),
	)
}

func githubButton() widget.Widget {
	return row.New(
		[]widget.Widget{
			iconbutton.New(
				icondata.GitHub,
				iconbutton.Style(iconbutton.Fill("#FFFFFF")),
				iconbutton.Tooltip("Check out our GitHub repository"),
			),
		},
		row.MainAxisAlignment(style.JustifyContentTypeEnd),
		row.CrossAxisAlignment(style.AlignItemsTypeCenter),
		row.Gap(8),
	)
}
