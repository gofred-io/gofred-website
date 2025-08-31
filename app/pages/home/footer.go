package home

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/container"
	iconbutton "github.com/gofred-io/gofred/icon_button"
	icondata "github.com/gofred-io/gofred/icon_data"
	"github.com/gofred-io/gofred/link"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/row"
	"github.com/gofred-io/gofred/text"
	"github.com/gofred-io/gofred/widget"
)

func footer() widget.BaseWidget {
	return container.New(
		row.New(
			[]widget.BaseWidget{
				socials(),
				copyright(),
				githubButton(),
			},
		),
		options.BackgroundColor("#23395B"),
		options.Padding(breakpoint.All(16)),
	)
}

func socials() widget.BaseWidget {
	twitterButton := link.New(
		iconbutton.New(
			icondata.Twitter,
			options.Fill("#FFFFFF"),
			options.Tooltip("Follow us on Twitter"),
			options.Class("footer-icon-button"),
		),
		options.Href("#"),
		options.NewTab(true),
	)

	youtubeButton := link.New(
		iconbutton.New(
			icondata.Youtube,
			options.Fill("#FFFFFF"),
			options.Tooltip("Subscribe to our YouTube channel"),
			options.Class("footer-icon-button"),
		),
		options.Href("#"),
		options.NewTab(true),
	)

	instagramButton := link.New(
		iconbutton.New(
			icondata.Instagram,
			options.Fill("#FFFFFF"),
			options.Tooltip("Follow us on Instagram"),
			options.Class("footer-icon-button"),
		),
		options.Href("#"),
		options.NewTab(true),
	)

	return row.New(
		[]widget.BaseWidget{
			twitterButton,
			youtubeButton,
			instagramButton,
		},
		options.MainAxisAlignment(options.AxisAlignmentTypeStart),
		options.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
		options.ColumnGap(8),
	)
}

func copyright() widget.BaseWidget {
	return row.New(
		[]widget.BaseWidget{
			text.New("Copyright © 2025 Gofred", options.FontSize(14), options.FontColor("#FFFFFF"), options.FontWeight("400")),
		},
		options.MainAxisAlignment(options.AxisAlignmentTypeCenter),
		options.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
		options.ColumnGap(8),
	)
}

func githubButton() widget.BaseWidget {
	githubButton := link.New(
		iconbutton.New(
			icondata.GitHub,
			options.Fill("#FFFFFF"),
			options.Tooltip("Check out our GitHub repository"),
			options.Class("footer-icon-button"),
		),
		options.Href("https://github.com/gofred-io/gofred-website"),
		options.NewTab(true),
	)

	return row.New(
		[]widget.BaseWidget{
			githubButton,
		},
		options.MainAxisAlignment(options.AxisAlignmentTypeEnd),
		options.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
		options.ColumnGap(8),
	)
}
