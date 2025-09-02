package home

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/container"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	iconbutton "github.com/gofred-io/gofred/foundation/icon_button"
	"github.com/gofred-io/gofred/foundation/link"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/options/spacing"
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
			row.Flex(1),
		),
		container.BackgroundColor("#23395B"),
		container.Padding(breakpoint.All(spacing.All(16))),
		container.Flex(1),
	)
}

func socials() widget.BaseWidget {
	twitterButton := link.New(
		iconbutton.New(
			icondata.Twitter,
			iconbutton.Fill("#FFFFFF"),
			iconbutton.Tooltip("Follow us on Twitter"),
			iconbutton.Class("footer-icon-button"),
		),
		link.Href("#"),
		link.NewTab(true),
	)

	youtubeButton := link.New(
		iconbutton.New(
			icondata.Youtube,
			iconbutton.Fill("#FFFFFF"),
			iconbutton.Tooltip("Subscribe to our YouTube channel"),
			iconbutton.Class("footer-icon-button"),
		),
		link.Href("#"),
		link.NewTab(true),
	)

	instagramButton := link.New(
		iconbutton.New(
			icondata.Instagram,
			iconbutton.Fill("#FFFFFF"),
			iconbutton.Tooltip("Follow us on Instagram"),
			iconbutton.Class("footer-icon-button"),
		),
		link.Href("#"),
		link.NewTab(true),
	)

	return row.New(
		[]widget.BaseWidget{
			twitterButton,
			youtubeButton,
			instagramButton,
		},
		row.MainAxisAlignment(options.AxisAlignmentTypeStart),
		row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
		row.Gap(8),
	)
}

func copyright() widget.BaseWidget {
	return row.New(
		[]widget.BaseWidget{
			text.New("Copyright © 2025 Gofred", text.FontSize(14), text.FontColor("#FFFFFF"), text.FontWeight("400")),
		},
		row.MainAxisAlignment(options.AxisAlignmentTypeCenter),
		row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
		row.Gap(8),
		row.Flex(1),
	)
}

func githubButton() widget.BaseWidget {
	githubButton := link.New(
		iconbutton.New(
			icondata.GitHub,
			iconbutton.Fill("#FFFFFF"),
			iconbutton.Tooltip("Check out our GitHub repository"),
			iconbutton.Class("footer-icon-button"),
		),
		link.Href("https://github.com/gofred-io/gofred-website"),
		link.NewTab(true),
	)

	return row.New(
		[]widget.BaseWidget{
			githubButton,
		},
		row.MainAxisAlignment(options.AxisAlignmentTypeEnd),
		row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
		row.Gap(8),
	)
}
