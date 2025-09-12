package home

import (
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/center"
	"github.com/gofred-io/gofred/foundation/column"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/grid"
	"github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	"github.com/gofred-io/gofred/foundation/row"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
)

func features() application.BaseWidget {
	return container.New(
		column.New(
			[]application.BaseWidget{
				featureTitle(),
				featureList(),
			},
			column.Gap(32),
			column.Flex(1),
		),
		container.Padding(breakpoint.All(spacing.All(64))),
	)
}

func featureTitle() application.BaseWidget {
	return center.New(
		text.New(
			"Why gofred?",
			text.FontSize(32),
			text.FontColor("#000000"),
			text.FontWeight("500"),
			text.UserSelect(theme.UserSelectTypeNone),
		),
	)
}

func featureList() application.BaseWidget {
	return center.New(
		container.New(
			grid.New(
				[]application.BaseWidget{
					featureItem1(),
					featureItem2(),
					featureItem3(),
					featureItem4(),
					featureItem5(),
					featureItem6(),
				},
				grid.ColumnCount(
					breakpoint.XS(1),
					breakpoint.SM(1),
					breakpoint.MD(2),
					breakpoint.LG(3),
					breakpoint.XL(3),
					breakpoint.XXL(6),
				),
				grid.ColumnGap(48),
				grid.RowGap(48),
			),
			container.MaxWidth(breakpoint.All(960)),
			container.WidthP(breakpoint.All(0.75)),
		),
	)
}

func featureItem1() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			row.New(
				[]application.BaseWidget{
					icon.New(icondata.RocketLaunchOutline, icon.Fill("#000000"), icon.Width(breakpoint.All(28)), icon.Height(breakpoint.All(28))),
					text.New("No JavaScript Required", text.FontSize(20), text.FontColor("#000000"), text.FontWeight("500"), text.LineHeight(1.5)),
				},
				row.Gap(8),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),
			text.New(
				"Write modern, reactive web applications entirely in Go. Gofred handles the WebAssembly layer so you can focus on building features — not managing JS interop.",
				text.FontSize(16),
				text.FontColor("#000000"),
				text.LineHeight(1.5),
			),
		},
		column.Gap(8),
	)
}

func featureItem2() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			row.New(
				[]application.BaseWidget{
					icon.New(icondata.LightningBoltOutline, icon.Fill("#000000"), icon.Width(breakpoint.All(28)), icon.Height(breakpoint.All(28))),
					text.New("High Performance", text.FontSize(20), text.FontColor("#000000"), text.FontWeight("500"), text.LineHeight(1.5)),
				},
				row.Gap(8),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),
			text.New(
				"Leverages Go’s speed and efficiency, compiled to WebAssembly for near-native browser performance.",
				text.FontSize(16),
				text.FontColor("#000000"),
				text.LineHeight(1.5),
			),
		},
		column.Gap(8),
	)
}

func featureItem3() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			row.New(
				[]application.BaseWidget{
					icon.New(icondata.PaletteOutline, icon.Fill("#000000"), icon.Width(breakpoint.All(28)), icon.Height(breakpoint.All(28))),
					text.New("Declarative & Reactive", text.FontSize(20), text.FontColor("#000000"), text.FontWeight("500"), text.LineHeight(1.5)),
				},
				row.Gap(8),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),
			text.New(
				"Inspired by React and Flutter, define your UI with simple Go interfaces. State changes automatically update the DOM — no manual wiring needed.",
				text.FontSize(16),
				text.FontColor("#000000"),
				text.LineHeight(1.5),
			),
		},
		column.Gap(8),
	)
}

func featureItem4() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			row.New(
				[]application.BaseWidget{
					icon.New(icondata.Cellphone, icon.Fill("#000000"), icon.Width(breakpoint.All(28)), icon.Height(breakpoint.All(28))),
					text.New("Responsive by Design", text.FontSize(20), text.FontColor("#000000"), text.FontWeight("500"), text.LineHeight(1.5)),
				},
				row.Gap(8),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),
			text.New(
				"Gofred makes it simple to build layouts that adapt to any device size, ensuring great user experiences everywhere.",
				text.FontSize(16),
				text.FontColor("#000000"),
				text.LineHeight(1.5),
			),
		},
		column.Gap(8),
	)
}

func featureItem5() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			row.New(
				[]application.BaseWidget{
					icon.New(icondata.PowerPlugOutline, icon.Fill("#000000"), icon.Width(breakpoint.All(28)), icon.Height(breakpoint.All(28))),
					text.New("Extensible & Flexible", text.FontSize(20), text.FontColor("#000000"), text.FontWeight("500"), text.LineHeight(1.5)),
				},
				row.Gap(8),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),
			text.New(
				"Use pure Go for core logic, but still extend with HTML, CSS, and JS when you need fine-grained control.",
				text.FontSize(16),
				text.FontColor("#000000"),
				text.LineHeight(1.5),
			),
		},
		column.Gap(8),
	)
}

func featureItem6() application.BaseWidget {
	return column.New(
		[]application.BaseWidget{
			row.New(
				[]application.BaseWidget{
					icon.New(icondata.Tools, icon.Fill("#000000"), icon.Width(breakpoint.All(28)), icon.Height(breakpoint.All(28))),
					text.New("Built for Go Developers", text.FontSize(20), text.FontColor("#000000"), text.FontWeight("500"), text.LineHeight(1.5)),
				},
				row.Gap(8),
				row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
			),
			text.New(
				"No need to switch languages — build full web applications using the same patterns and tooling you already know in Go.",
				text.FontSize(16),
				text.FontColor("#000000"),
				text.LineHeight(1.5),
			),
		},
		column.Gap(8),
	)
}
