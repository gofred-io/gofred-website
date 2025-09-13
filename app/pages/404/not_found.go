package notfound

import (
	"github.com/gofred-io/gofred-website/app/components/footer"
	"github.com/gofred-io/gofred-website/app/components/header"
	"github.com/gofred-io/gofred/application"
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
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
)

func New(params router.RouteParams) application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			header.Get(),
			mainContent(),
			footer.Get(),
		},
		column.Flex(1),
	)
}

func mainContent() application.BaseWidget {
	return container.New(
		column.New(
			[]application.BaseWidget{
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
			column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			column.MainAxisAlignment(theme.AxisAlignmentTypeCenter),
		),
		container.BackgroundColor("#F8F9FA"),
		container.Flex(1),
		container.Padding(breakpoint.All(spacing.All(32))),
	)
}

func errorIcon() application.BaseWidget {
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
		container.BorderWidth(spacing.All(3)),
		container.BorderStyle(theme.BorderStyleTypeSolid),
	)
}

func errorTitle() application.BaseWidget {
	return text.New(
		"404",
		text.FontSize(72),
		text.FontColor("#2C3E50"),
		text.FontWeight("700"),
		text.UserSelect(theme.UserSelectTypeNone),
	)
}

func errorDescription() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			text.New(
				"Oops! The page you're looking for doesn't exist.",
				text.FontSize(24),
				text.FontColor("#34495E"),
				text.FontWeight("500"),
				text.UserSelect(theme.UserSelectTypeNone),
			),
			spacer.New(spacer.Height(8)),
			text.New(
				"It might have been moved, deleted, or you entered the wrong URL.",
				text.FontSize(16),
				text.FontColor("#7F8C8D"),
				text.UserSelect(theme.UserSelectTypeNone),
			),
		},
		column.Gap(8),
		column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
	)
}

func actionButtons() application.BaseWidget {
	return row.New(
		[]application.BaseWidget{
			backButton(),
			homeButton(),
		},
		row.Gap(16),
		row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
	)
}

func backButton() application.BaseWidget {
	return button.New(
		container.New(
			row.New(
				[]application.BaseWidget{
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
						text.UserSelect(theme.UserSelectTypeNone),
					),
				},
				row.Gap(8),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),
			container.Padding(breakpoint.All(spacing.Right(4))),
		),
		button.OnClick(func(this application.BaseWidget, e application.Event) {
			application.Context().GoBack()
		}),
	)
}

func homeButton() application.BaseWidget {
	return link.New(
		button.New(
			container.New(
				row.New(
					[]application.BaseWidget{
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
							text.UserSelect(theme.UserSelectTypeNone),
						),
					},
					row.Gap(8),
					row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
				),
				container.Padding(breakpoint.All(spacing.Right(4))),
			),
		),
		link.Href("/"),
	)
}
