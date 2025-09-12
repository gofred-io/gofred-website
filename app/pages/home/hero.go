package home

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/button"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/grid"
	"github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	"github.com/gofred-io/gofred/foundation/link"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/spacer"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
)

func hero() application.BaseWidget {
	return container.New(
		column.New(
			[]application.BaseWidget{
				heroTitle(),
				heroSubtitle(),
				spacer.New(spacer.Height(8)),
				sampleCode(),
				spacer.New(spacer.Height(8)),
				cta(),
			},
			column.Gap(16),
			column.Flex(1),
			column.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			column.MainAxisAlignment(theme.AxisAlignmentTypeCenter),
		),
		container.BackgroundColor("#2B799B"),
		container.Flex(1),
		container.Padding(breakpoint.All(spacing.All(32))),
	)
}

func heroTitle() application.BaseWidget {
	return text.New(
		"Build responsive web apps in Go – no JavaScript required",
		text.FontSize(24),
		text.FontColor("#FFFFFF"),
		text.FontWeight("600"),
		text.UserSelect(theme.UserSelectTypeNone),
	)
}

func heroSubtitle() application.BaseWidget {
	return text.New(
		"Write modern, reactive web applications with WebAssembly using only Go.",
		text.FontSize(18),
		text.FontColor("#FFFFFF"),
		text.UserSelect(theme.UserSelectTypeNone),
	)
}

func sampleCode() application.BaseWidget {
	closeButton := container.New(
		application.Nil,
		container.BackgroundColor("#FF5F58"),
		container.Height(breakpoint.All(16)),
		container.Width(breakpoint.All(16)),
		container.BorderRadius(8),
	)

	fullScreenButton := container.New(
		application.Nil,
		container.BackgroundColor("#FFBD2E"),
		container.Height(breakpoint.All(16)),
		container.Width(breakpoint.All(16)),
		container.BorderRadius(8),
	)

	minimizeButton := container.New(
		application.Nil,
		container.BackgroundColor("#28C941"),
		container.Height(breakpoint.All(16)),
		container.Width(breakpoint.All(16)),
		container.BorderRadius(8),
	)

	windowTitle := text.New(
		"main.go",
		text.FontSize(16),
		text.FontColor("#FFFFFF"),
		text.FontWeight("500"),
		text.UserSelect(theme.UserSelectTypeNone),
	)

	titleBar := row.New(
		[]application.BaseWidget{
			row.New(
				[]application.BaseWidget{
					closeButton,
					fullScreenButton,
					minimizeButton,
				},
				row.MainAxisAlignment(theme.AxisAlignmentTypeStart),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeStart),
				row.Gap(8),
				row.Flex(1),
			),
			windowTitle,
			spacer.New(),
		},
		row.MainAxisSize(theme.AxisSizeTypeMax),
		row.MainAxisAlignment(theme.AxisAlignmentTypeStart),
		row.CrossAxisAlignment(theme.AxisAlignmentTypeStart),
		row.Gap(8),
	)

	packageName := row.New(
		[]application.BaseWidget{
			text.New(
				"package",
				text.FontSize(14),
				text.FontColor("#4F8CBF"),
				text.UserSelect(theme.UserSelectTypeNone),
			),
			text.New(
				"main",
				text.FontSize(14),
				text.FontColor("#49B8A2"),
				text.UserSelect(theme.UserSelectTypeNone),
			),
		},
		row.MainAxisAlignment(theme.AxisAlignmentTypeStart),
		row.CrossAxisAlignment(theme.AxisAlignmentTypeStart),
		row.Gap(6),
	)

	imports := column.New(
		[]application.BaseWidget{
			row.New(
				[]application.BaseWidget{
					text.New(
						"import",
						text.FontSize(14),
						text.FontColor("#AD77A8"),
						text.UserSelect(theme.UserSelectTypeNone),
					),
					text.New(
						"(",
						text.FontSize(14),
						text.FontColor("#E0BD04"),
						text.UserSelect(theme.UserSelectTypeNone),
					),
				},
				row.MainAxisAlignment(theme.AxisAlignmentTypeStart),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeStart),
				row.Gap(6),
			),
			row.New(
				[]application.BaseWidget{
					spacer.New(spacer.Width(6)),
					text.New(
						"\"github.com/gofred-io/gofred-website/app\"",
						text.FontSize(14),
						text.FontColor("#B9836D"),
						text.UserSelect(theme.UserSelectTypeNone),
					),
				},
				row.MainAxisAlignment(theme.AxisAlignmentTypeStart),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeStart),
				row.Gap(6),
			),
			row.New(
				[]application.BaseWidget{
					spacer.New(spacer.Width(6)),
					text.New(
						"\"github.com/gofred-io/gofred/application\"",
						text.FontSize(14),
						text.FontColor("#B9836D"),
						text.UserSelect(theme.UserSelectTypeNone),
					),
				},
				row.MainAxisAlignment(theme.AxisAlignmentTypeStart),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeStart),
				row.Gap(6),
			),
			row.New(
				[]application.BaseWidget{
					text.New(
						")",
						text.FontSize(14),
						text.FontColor("#E0BD04"),
						text.UserSelect(theme.UserSelectTypeNone),
					),
				},
			),
		},
		column.MainAxisAlignment(theme.AxisAlignmentTypeStart),
		column.CrossAxisAlignment(theme.AxisAlignmentTypeStart),
		column.Gap(6),
	)

	mainFunc := column.New(
		[]application.BaseWidget{
			row.New(
				[]application.BaseWidget{
					text.New(
						"func",
						text.FontSize(14),
						text.FontColor("#4F8CBF"),
						text.UserSelect(theme.UserSelectTypeNone),
					),
					row.New(
						[]application.BaseWidget{
							text.New(
								"main",
								text.FontSize(14),
								text.FontColor("#DCDCAA"),
								text.UserSelect(theme.UserSelectTypeNone),
							),
							text.New(
								"()",
								text.FontSize(14),
								text.FontColor("#E0BD04"),
								text.UserSelect(theme.UserSelectTypeNone),
							),
						},
						row.MainAxisAlignment(theme.AxisAlignmentTypeStart),
						row.CrossAxisAlignment(theme.AxisAlignmentTypeStart),
					),
					text.New(
						"{",
						text.FontSize(14),
						text.FontColor("#E0BD04"),
						text.UserSelect(theme.UserSelectTypeNone),
					),
				},
				row.MainAxisAlignment(theme.AxisAlignmentTypeStart),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeStart),
				row.Gap(6),
			),
			row.New(
				[]application.BaseWidget{
					spacer.New(spacer.Width(6)),
					row.New(
						[]application.BaseWidget{
							text.New(
								"application",
								text.FontSize(14),
								text.FontColor("#8AD1DE"),
								text.UserSelect(theme.UserSelectTypeNone),
							),
							text.New(
								".",
								text.FontSize(14),
								text.FontColor("#BCBCBC"),
								text.UserSelect(theme.UserSelectTypeNone),
							),
							text.New(
								"Run",
								text.FontSize(14),
								text.FontColor("#DCDCAA"),
								text.UserSelect(theme.UserSelectTypeNone),
							),
							text.New(
								"(",
								text.FontSize(14),
								text.FontColor("#AD77A8"),
								text.UserSelect(theme.UserSelectTypeNone),
							),
							text.New(
								"app",
								text.FontSize(14),
								text.FontColor("#8AD1DE"),
								text.UserSelect(theme.UserSelectTypeNone),
							),
							text.New(
								".",
								text.FontSize(14),
								text.FontColor("#BCBCBC"),
								text.UserSelect(theme.UserSelectTypeNone),
							),
							text.New(
								"New",
								text.FontSize(14),
								text.FontColor("#DCDCAA"),
								text.UserSelect(theme.UserSelectTypeNone),
							),
							text.New(
								"()",
								text.FontSize(14),
								text.FontColor("#4F8CBF"),
								text.UserSelect(theme.UserSelectTypeNone),
							),
							text.New(
								")",
								text.FontSize(14),
								text.FontColor("#AD77A8"),
								text.UserSelect(theme.UserSelectTypeNone),
							),
						},
					),
				},
				row.MainAxisAlignment(theme.AxisAlignmentTypeStart),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeStart),
				row.Gap(6),
			),
			text.New(
				"}",
				text.FontSize(14),
				text.FontColor("#E0BD04"),
				text.UserSelect(theme.UserSelectTypeNone),
			),
		},
		column.MainAxisAlignment(theme.AxisAlignmentTypeStart),
		column.CrossAxisAlignment(theme.AxisAlignmentTypeStart),
		column.Gap(6),
	)

	return container.New(
		column.New(
			[]application.BaseWidget{
				titleBar,
				packageName,
				spacer.New(spacer.Height(8)),
				imports,
				spacer.New(spacer.Height(8)),
				mainFunc,
			},
			column.MainAxisAlignment(theme.AxisAlignmentTypeStart),
			column.Gap(6),
			column.Flex(1),
		),
		container.BackgroundColor("#282A36"),
		container.Height(breakpoint.All(240)),
		container.BorderRadius(8),
		container.Padding(breakpoint.All(spacing.All(16))),
		container.MaxWidth(breakpoint.All(540)),
		container.WidthP(breakpoint.All(0.5), breakpoint.XS(1.0), breakpoint.SM(1.0), breakpoint.MD(0.75)),
	)
}

func cta() application.BaseWidget {
	return grid.New(
		[]application.BaseWidget{
			link.New(
				button.New(
					text.New(
						"Get Started",
						text.FontSize(14),
						text.FontColor("#FFFFFF"),
						text.FontWeight("500"),
						text.UserSelect(theme.UserSelectTypeNone),
					),
					button.Width(breakpoint.All(100)),
				),
				link.Href("/docs/installation"),
			),
			link.New(
				button.New(
					row.New(
						[]application.BaseWidget{
							icon.New(
								icondata.Github,
								icon.Fill("#FFFFFF"),
								icon.Height(breakpoint.All(16)),
								icon.Width(breakpoint.All(16)),
							),
							text.New(
								"GitHub",
								text.FontSize(14),
								text.FontColor("#FFFFFF"),
								text.FontWeight("500"),
								text.UserSelect(theme.UserSelectTypeNone),
							),
						},
						row.MainAxisAlignment(theme.AxisAlignmentTypeCenter),
						row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
						row.Gap(6),
					),
					button.BackgroundColor("#151B23"),
					button.Width(breakpoint.All(100)),
				),
				link.Href("https://github.com/gofred-io/gofred"),
				link.NewTab(true),
			),
		},
		grid.ColumnGap(8),
		grid.RowGap(8),
		grid.ColumnCount(
			breakpoint.All(2),
			breakpoint.XS(1),
		),
	)
}
