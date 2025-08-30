package home

import (
	"github.com/gofred-io/gofred/button"
	"github.com/gofred-io/gofred/column"
	"github.com/gofred-io/gofred/container"
	"github.com/gofred-io/gofred/icon"
	icondata "github.com/gofred-io/gofred/icon_data"
	"github.com/gofred-io/gofred/link"
	"github.com/gofred-io/gofred/row"
	"github.com/gofred-io/gofred/spacer"
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/text"
	"github.com/gofred-io/gofred/widget"
)

func hero() widget.Widget {
	return container.New(
		column.New(
			[]widget.Widget{
				heroTitle(),
				heroSubtitle(),
				spacer.New(spacer.Height(8)),
				sampleCode(),
				spacer.New(spacer.Height(8)),
				cta(),
			},
			column.Gap(16),
			column.CrossAxisAlignment(style.AlignItemsTypeCenter),
			column.MainAxisAlignment(style.JustifyContentTypeCenter),
		),
		container.Style(
			container.Background(
				style.Background{
					Color: "#2B799B",
				},
			),
			container.Flex(1),
			container.Padding(style.Padding{
				Top:    32,
				Bottom: 32,
				Left:   32,
				Right:  32,
			}),
		),
	)
}

func heroTitle() widget.Widget {
	return text.New(
		"Build responsive web apps in Go – no JavaScript required",
		text.Style(
			text.Font(text.Size(24), text.Color("#FFFFFF"), text.Weight("600")),
		),
	)
}

func heroSubtitle() widget.Widget {
	return text.New(
		"Write modern, reactive web applications with WebAssembly using only Go.",
		text.Style(
			text.Font(text.Size(18), text.Color("#FFFFFF"), text.Weight("400")),
		),
	)
}

func sampleCode() widget.Widget {
	closeButton := container.New(
		widget.Nil,
		container.Style(
			container.Background(
				style.Background{
					Color: "#FF5F58",
				},
			),
			container.Height(16),
			container.Width(16),
			container.Border(
				style.Border{
					Radius: 8,
				},
			),
		),
	)

	fullScreenButton := container.New(
		widget.Nil,
		container.Style(
			container.Background(
				style.Background{
					Color: "#FFBD2E",
				},
			),
			container.Height(16),
			container.Width(16),
			container.Border(
				style.Border{
					Radius: 8,
				},
			),
		),
	)

	minimizeButton := container.New(
		widget.Nil,
		container.Style(
			container.Background(
				style.Background{
					Color: "#28C941",
				},
			),
			container.Height(16),
			container.Width(16),
			container.Border(
				style.Border{
					Radius: 8,
				},
			),
		),
	)

	windowTitle := text.New(
		"main.go",
		text.Style(
			text.Font(text.Size(16), text.Color("#FFFFFF"), text.Weight("500")),
			text.UserSelect(style.UserSelectTypeNone),
		),
	)

	titleBar := row.New(
		[]widget.Widget{
			row.New(
				[]widget.Widget{
					closeButton,
					fullScreenButton,
					minimizeButton,
				},
				row.MainAxisAlignment(style.JustifyContentTypeStart),
				row.CrossAxisAlignment(style.AlignItemsTypeStart),
				row.Gap(8),
			),
			windowTitle,
			spacer.New(),
		},
		row.MainAxisAlignment(style.JustifyContentTypeStart),
		row.CrossAxisAlignment(style.AlignItemsTypeStart),
		row.Gap(8),
		row.Flex(0),
	)

	packageName := row.New(
		[]widget.Widget{
			text.New(
				"package",
				text.Style(
					text.Font(text.Size(14), text.Color("#4F8CBF"), text.Weight("400")),
					text.UserSelect(style.UserSelectTypeNone),
				),
			),
			text.New(
				"main",
				text.Style(
					text.Font(text.Size(14), text.Color("#49B8A2"), text.Weight("400")),
					text.UserSelect(style.UserSelectTypeNone),
				),
			),
		},
		row.MainAxisAlignment(style.JustifyContentTypeStart),
		row.CrossAxisAlignment(style.AlignItemsTypeStart),
		row.Gap(6),
		row.Flex(0),
	)

	imports := column.New(
		[]widget.Widget{
			row.New(
				[]widget.Widget{
					text.New(
						"import",
						text.Style(
							text.Font(text.Size(14), text.Color("#AD77A8"), text.Weight("400")),
							text.UserSelect(style.UserSelectTypeNone),
						),
					),
					text.New(
						"(",
						text.Style(
							text.Font(text.Size(14), text.Color("#E0BD04"), text.Weight("400")),
							text.UserSelect(style.UserSelectTypeNone),
						),
					),
				},
				row.MainAxisAlignment(style.JustifyContentTypeStart),
				row.CrossAxisAlignment(style.AlignItemsTypeStart),
				row.Gap(6),
				row.Flex(0),
			),
			row.New(
				[]widget.Widget{
					spacer.New(spacer.Width(6)),
					text.New(
						"\"github.com/gofred-io/gofred-website/app\"",
						text.Style(
							text.Font(text.Size(14), text.Color("#B9836D"), text.Weight("400")),
							text.UserSelect(style.UserSelectTypeNone),
						),
					),
				},
				row.MainAxisAlignment(style.JustifyContentTypeStart),
				row.CrossAxisAlignment(style.AlignItemsTypeStart),
				row.Gap(6),
				row.Flex(0),
			),
			row.New(
				[]widget.Widget{
					spacer.New(spacer.Width(6)),
					text.New(
						"\"github.com/gofred-io/gofred/application\"",
						text.Style(
							text.Font(text.Size(14), text.Color("#B9836D"), text.Weight("400")),
							text.UserSelect(style.UserSelectTypeNone),
						),
					),
				},
				row.MainAxisAlignment(style.JustifyContentTypeStart),
				row.CrossAxisAlignment(style.AlignItemsTypeStart),
				row.Gap(6),
				row.Flex(0),
			),
			row.New(
				[]widget.Widget{
					text.New(
						")",
						text.Style(
							text.Font(text.Size(14), text.Color("#E0BD04"), text.Weight("400")),
							text.UserSelect(style.UserSelectTypeNone),
						),
					),
				},
			),
		},
		column.MainAxisAlignment(style.JustifyContentTypeStart),
		column.CrossAxisAlignment(style.AlignItemsTypeStart),
		column.Gap(6),
	)

	mainFunc := column.New(
		[]widget.Widget{
			row.New(
				[]widget.Widget{
					text.New(
						"func",
						text.Style(
							text.Font(text.Size(14), text.Color("#4F8CBF"), text.Weight("400")),
							text.UserSelect(style.UserSelectTypeNone),
						),
					),
					row.New(
						[]widget.Widget{
							text.New(
								"main",
								text.Style(
									text.Font(text.Size(14), text.Color("#DCDCAA"), text.Weight("400")),
									text.UserSelect(style.UserSelectTypeNone),
								),
							),
							text.New(
								"()",
								text.Style(
									text.Font(text.Size(14), text.Color("#E0BD04"), text.Weight("400")),
									text.UserSelect(style.UserSelectTypeNone),
								),
							),
						},
						row.MainAxisAlignment(style.JustifyContentTypeStart),
						row.CrossAxisAlignment(style.AlignItemsTypeStart),
						row.Flex(0),
					),
					text.New(
						"{",
						text.Style(
							text.Font(text.Size(14), text.Color("#E0BD04"), text.Weight("400")),
							text.UserSelect(style.UserSelectTypeNone),
						),
					),
				},
				row.MainAxisAlignment(style.JustifyContentTypeStart),
				row.CrossAxisAlignment(style.AlignItemsTypeStart),
				row.Gap(6),
				row.Flex(0),
			),
			row.New(
				[]widget.Widget{
					spacer.New(spacer.Width(6)),
					row.New(
						[]widget.Widget{
							text.New(
								"application",
								text.Style(
									text.Font(text.Size(14), text.Color("#8AD1DE"), text.Weight("400")),
									text.UserSelect(style.UserSelectTypeNone),
								),
							),
							text.New(
								".",
								text.Style(
									text.Font(text.Size(14), text.Color("#BCBCBC"), text.Weight("400")),
									text.UserSelect(style.UserSelectTypeNone),
								),
							),
							text.New(
								"Run",
								text.Style(
									text.Font(text.Size(14), text.Color("#DCDCAA"), text.Weight("400")),
									text.UserSelect(style.UserSelectTypeNone),
								),
							),
							text.New(
								"(",
								text.Style(
									text.Font(text.Size(14), text.Color("#AD77A8"), text.Weight("400")),
									text.UserSelect(style.UserSelectTypeNone),
								),
							),
							text.New(
								"app",
								text.Style(
									text.Font(text.Size(14), text.Color("#8AD1DE"), text.Weight("400")),
									text.UserSelect(style.UserSelectTypeNone),
								),
							),
							text.New(
								".",
								text.Style(
									text.Font(text.Size(14), text.Color("#BCBCBC"), text.Weight("400")),
									text.UserSelect(style.UserSelectTypeNone),
								),
							),
							text.New(
								"New",
								text.Style(
									text.Font(text.Size(14), text.Color("#DCDCAA"), text.Weight("400")),
									text.UserSelect(style.UserSelectTypeNone),
								),
							),
							text.New(
								"()",
								text.Style(
									text.Font(text.Size(14), text.Color("#4F8CBF"), text.Weight("400")),
									text.UserSelect(style.UserSelectTypeNone),
								),
							),
							text.New(
								")",
								text.Style(
									text.Font(text.Size(14), text.Color("#AD77A8"), text.Weight("400")),
									text.UserSelect(style.UserSelectTypeNone),
								),
							),
						},
					),
				},
				row.MainAxisAlignment(style.JustifyContentTypeStart),
				row.CrossAxisAlignment(style.AlignItemsTypeStart),
				row.Gap(6),
				row.Flex(0),
			),
			text.New(
				"}",
				text.Style(
					text.Font(text.Size(14), text.Color("#E0BD04"), text.Weight("400")),
					text.UserSelect(style.UserSelectTypeNone),
				),
			),
		},
		column.MainAxisAlignment(style.JustifyContentTypeStart),
		column.CrossAxisAlignment(style.AlignItemsTypeStart),
		column.Gap(6),
	)

	return container.New(
		column.New(
			[]widget.Widget{
				titleBar,
				packageName,
				spacer.New(spacer.Height(8)),
				imports,
				spacer.New(spacer.Height(8)),
				mainFunc,
			},
			column.MainAxisAlignment(style.JustifyContentTypeStart),
			column.Gap(6),
		),
		container.Style(
			container.Background(
				style.Background{
					Color: "#282A36",
				},
			),
			container.WidthP(0.5),
			container.Height(240),
			container.Border(
				style.Border{
					Radius: 8,
				},
			),
			container.Padding(style.Padding{
				Top:    16,
				Bottom: 16,
				Left:   16,
				Right:  16,
			}),
			container.MaxWidth(540),
		),
	)
}

func cta() widget.Widget {
	return row.New(
		[]widget.Widget{
			link.New(
				button.New(
					text.New(
						"Get Started",
						text.Style(
							text.Font(text.Size(14), text.Color("#FFFFFF"), text.Weight("500"), text.Family("Ubuntu")),
						),
					),
					button.Style(
						button.Width(100),
					),
				),
				link.Href("/docs#getting-started"),
			),
			link.New(
				button.New(
					row.New(
						[]widget.Widget{
							icon.New(
								icondata.GitHub,
								icon.Style(
									icon.Fill("#FFFFFF"),
									icon.Height(16),
									icon.Width(16),
								),
							),
							text.New(
								"GitHub",
								text.Style(
									text.Font(text.Size(14), text.Color("#FFFFFF"), text.Weight("500"), text.Family("Ubuntu")),
								),
							),
						},
						row.MainAxisAlignment(style.JustifyContentTypeCenter),
						row.CrossAxisAlignment(style.AlignItemsTypeCenter),
						row.Gap(6),
						row.Flex(0),
					),
					button.Style(
						button.Background(
							style.Background{
								Color: "#151b23",
							},
						),
						button.Width(100),
					),
				),
				link.Href("https://github.com/gofred-io/gofred"),
				link.NewTab(true),
			),
		},
		row.MainAxisAlignment(style.JustifyContentTypeStart),
		row.CrossAxisAlignment(style.AlignItemsTypeStart),
		row.Gap(8),
		row.Flex(0),
	)
}
