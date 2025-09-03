package notfound

import (
	"github.com/gofred-io/gofred-website/app/pages/home"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/button"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	"github.com/gofred-io/gofred/foundation/link"
	"github.com/gofred-io/gofred/foundation/router"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/widget"
)

func New(params router.RouteParams) widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			home.Header(),
			mainContent(),
			home.Footer(),
		},
		column.Flex(1),
	)
}

func mainContent() widget.BaseWidget {
	return container.New(
		column.New(
			[]widget.BaseWidget{
				errorIcon(),
				spacer.New(spacer.Height(24)),
				errorTitle(),
				spacer.New(spacer.Height(16)),
				errorDescription(),
				spacer.New(spacer.Height(32)),
				actionButtons(),
			},
			column.Gap(16),
			column.Flex(1),
			column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
			column.MainAxisAlignment(options.AxisAlignmentTypeCenter),
		),
		container.BackgroundColor("#F8F9FA"),
		container.Flex(1),
		container.Padding(breakpoint.All(spacing.All(32))),
	)
}

func errorIcon() widget.BaseWidget {
	return container.New(
		icon.New(
			icondata.Alert,
			icon.Width(breakpoint.All(120)),
			icon.Height(breakpoint.All(120)),
			icon.Fill("#E74C3C"),
		),
		container.BackgroundColor("#FFFFFF"),
		container.BorderRadius(48),
		container.Padding(breakpoint.All(spacing.All(20))),
		container.BorderColor("#E74C3C"),
		container.BorderWidth(3, 3, 3, 3),
		container.BorderStyle(options.BorderStyleTypeSolid),
	)
}

func errorTitle() widget.BaseWidget {
	return text.New(
		"404",
		text.FontSize(72),
		text.FontColor("#2C3E50"),
		text.FontWeight("700"),
		text.UserSelect(options.UserSelectTypeNone),
	)
}

func errorDescription() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			text.New(
				"Oops! The page you're looking for doesn't exist.",
				text.FontSize(24),
				text.FontColor("#34495E"),
				text.FontWeight("500"),
				text.UserSelect(options.UserSelectTypeNone),
			),
			spacer.New(spacer.Height(8)),
			text.New(
				"It might have been moved, deleted, or you entered the wrong URL.",
				text.FontSize(16),
				text.FontColor("#7F8C8D"),
				text.FontWeight("400"),
				text.UserSelect(options.UserSelectTypeNone),
			),
		},
		column.Gap(8),
		column.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
	)
}

func actionButtons() widget.BaseWidget {
	return row.New(
		[]widget.BaseWidget{
			backButton(),
			homeButton(),
		},
		row.Gap(16),
		row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
	)
}

func backButton() widget.BaseWidget {
	return button.New(
		container.New(
			row.New(
				[]widget.BaseWidget{
					icon.New(
						icondata.ChevronLeft,
						icon.Width(breakpoint.All(20)),
						icon.Height(breakpoint.All(20)),
						icon.Fill("#FFFFFF"),
					),
					text.New(
						"Go Back",
						text.FontSize(16),
						text.FontColor("#FFFFFF"),
						text.FontWeight("500"),
						text.UserSelect(options.UserSelectTypeNone),
					),
				},
				row.Gap(8),
				row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
			),
			container.Padding(breakpoint.All(spacing.Right(4))),
		),
		button.OnClick(func(this widget.BaseWidget, e widget.Event) {
			widget.Context().GoBack()
		}),
	)
}

func homeButton() widget.BaseWidget {
	return link.New(
		button.New(
			container.New(
				row.New(
					[]widget.BaseWidget{
						icon.New(
							icondata.Home,
							icon.Width(breakpoint.All(20)),
							icon.Height(breakpoint.All(20)),
							icon.Fill("#FFFFFF"),
						),
						text.New(
							"Go Home",
							text.FontSize(16),
							text.FontColor("#FFFFFF"),
							text.FontWeight("500"),
							text.UserSelect(options.UserSelectTypeNone),
						),
					},
					row.Gap(8),
					row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
				),
				container.Padding(breakpoint.All(spacing.Right(4))),
			),
		),
		link.Href("/"),
	)
}
