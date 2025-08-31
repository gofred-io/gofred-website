package home

import (
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/center"
	"github.com/gofred-io/gofred/column"
	"github.com/gofred-io/gofred/container"
	"github.com/gofred-io/gofred/grid"
	"github.com/gofred-io/gofred/icon"
	icondata "github.com/gofred-io/gofred/icon_data"
	"github.com/gofred-io/gofred/options"
	"github.com/gofred-io/gofred/row"
	"github.com/gofred-io/gofred/text"
	"github.com/gofred-io/gofred/widget"
)

func features() widget.BaseWidget {
	return container.New(
		column.New(
			[]widget.BaseWidget{
				featureTitle(),
				featureList(),
			},
			column.Gap(32),
			column.Flex(1),
		),
		options.Padding(breakpoint.All(64)),
	)
}

func featureTitle() widget.BaseWidget {
	return center.New(
		text.New(
			"Why gofred?",
			options.FontSize(32),
			options.FontColor("#000000"),
			options.FontWeight("500"),
			options.UserSelect(options.UserSelectTypeNone),
		),
	)
}

func featureList() widget.BaseWidget {
	return center.New(
		container.New(
			grid.New(
				[]widget.BaseWidget{
					featureItem1(),
					featureItem2(),
					featureItem3(),
					featureItem4(),
					featureItem5(),
					featureItem6(),
				},
				options.ColumnCount(
					breakpoint.XS(1),
					breakpoint.SM(1),
					breakpoint.MD(2),
					breakpoint.LG(3),
					breakpoint.XL(3),
					breakpoint.XXL(6),
				),
				options.ColumnGap(48),
				options.RowGap(48),
			),
			options.MaxWidth(960),
			options.WidthP(breakpoint.All(0.75)),
		),
	)
}

func featureItem1() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			row.New(
				[]widget.BaseWidget{
					icon.New(icondata.RocketLaunchOutline, options.Fill("#000000"), options.Width(breakpoint.All(28)), options.Height(breakpoint.All(28))),
					text.New("No JavaScript Required", options.FontSize(20), options.FontColor("#000000"), options.FontWeight("500"), options.LineHeight(1.5)),
				},
				row.Gap(8),
				row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
			),
			text.New(
				"Write modern, reactive web applications entirely in Go. Gofred handles the WebAssembly layer so you can focus on building features — not managing JS interop.",
				options.FontSize(16),
				options.FontColor("#000000"),
				options.FontWeight("400"),
				options.LineHeight(1.5),
			),
		},
		column.Gap(8),
	)
}

func featureItem2() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			row.New(
				[]widget.BaseWidget{
					icon.New(icondata.LightningBoltOutline, options.Fill("#000000"), options.Width(breakpoint.All(28)), options.Height(breakpoint.All(28))),
					text.New("High Performance", options.FontSize(20), options.FontColor("#000000"), options.FontWeight("500"), options.LineHeight(1.5)),
				},
				row.Gap(8),
				row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
			),
			text.New(
				"Leverages Go’s speed and efficiency, compiled to WebAssembly for near-native browser performance.",
				options.FontSize(16),
				options.FontColor("#000000"),
				options.FontWeight("400"),
				options.LineHeight(1.5),
			),
		},
		column.Gap(8),
	)
}

func featureItem3() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			row.New(
				[]widget.BaseWidget{
					icon.New(icondata.PaletteOutline, options.Fill("#000000"), options.Width(breakpoint.All(28)), options.Height(breakpoint.All(28))),
					text.New("Declarative & Reactive", options.FontSize(20), options.FontColor("#000000"), options.FontWeight("500"), options.LineHeight(1.5)),
				},
				row.Gap(8),
				row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
			),
			text.New(
				"Inspired by React and Flutter, define your UI with simple Go interfaces. State changes automatically update the DOM — no manual wiring needed.",
				options.FontSize(16),
				options.FontColor("#000000"),
				options.FontWeight("400"),
				options.LineHeight(1.5),
			),
		},
		column.Gap(8),
	)
}

func featureItem4() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			row.New(
				[]widget.BaseWidget{
					icon.New(icondata.Cellphone, options.Fill("#000000"), options.Width(breakpoint.All(28)), options.Height(breakpoint.All(28))),
					text.New("Responsive by Design", options.FontSize(20), options.FontColor("#000000"), options.FontWeight("500"), options.LineHeight(1.5)),
				},
				row.Gap(8),
				row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
			),
			text.New(
				"Gofred makes it simple to build layouts that adapt to any device size, ensuring great user experiences everywhere.",
				options.FontSize(16),
				options.FontColor("#000000"),
				options.FontWeight("400"),
				options.LineHeight(1.5),
			),
		},
		column.Gap(8),
	)
}

func featureItem5() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			row.New(
				[]widget.BaseWidget{
					icon.New(icondata.PowerPlugOutline, options.Fill("#000000"), options.Width(breakpoint.All(28)), options.Height(breakpoint.All(28))),
					text.New("Extensible & Flexible", options.FontSize(20), options.FontColor("#000000"), options.FontWeight("500"), options.LineHeight(1.5)),
				},
				row.Gap(8),
				row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
			),
			text.New(
				"Use pure Go for core logic, but still extend with HTML, CSS, and JS when you need fine-grained control.",
				options.FontSize(16),
				options.FontColor("#000000"),
				options.FontWeight("400"),
				options.LineHeight(1.5),
			),
		},
		column.Gap(8),
	)
}

func featureItem6() widget.BaseWidget {
	return column.New(
		[]widget.BaseWidget{
			row.New(
				[]widget.BaseWidget{
					icon.New(icondata.Tools, options.Fill("#000000"), options.Width(breakpoint.All(28)), options.Height(breakpoint.All(28))),
					text.New("Built for Go Developers", options.FontSize(20), options.FontColor("#000000"), options.FontWeight("500"), options.LineHeight(1.5)),
				},
				row.Gap(8),
				row.CrossAxisAlignment(options.AxisAlignmentTypeCenter),
			),
			text.New(
				"No need to switch languages — build full web applications using the same patterns and tooling you already know in Go.",
				options.FontSize(16),
				options.FontColor("#000000"),
				options.FontWeight("400"),
				options.LineHeight(1.5),
			),
		},
		column.Gap(8),
	)
}
