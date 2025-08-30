package home

import (
	"github.com/gofred-io/gofred/center"
	"github.com/gofred-io/gofred/column"
	"github.com/gofred-io/gofred/container"
	"github.com/gofred-io/gofred/grid"
	"github.com/gofred-io/gofred/icon"
	icondata "github.com/gofred-io/gofred/icon_data"
	"github.com/gofred-io/gofred/row"
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/style/breakpoint"
	"github.com/gofred-io/gofred/text"
	"github.com/gofred-io/gofred/widget"
)

func features() widget.Widget {
	return container.New(
		column.New(
			[]widget.Widget{
				featureTitle(),
				featureList(),
			},
			column.Gap(32),
		),
		container.Style(
			container.Padding(style.Padding{
				Top:    64,
				Bottom: 64,
				Left:   64,
				Right:  64,
			}),
		),
	)
}

func featureTitle() widget.Widget {
	return center.New(
		text.New(
			"Why gofred?",
			text.Style(
				text.Font(text.Size(32), text.Color("#000000"), text.Weight("500")),
			),
		),
	)
}

func featureList() widget.Widget {
	return center.New(
		container.New(
			grid.New(
				[]widget.Widget{
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
			container.Style(
				container.WidthP(0.75),
				container.MaxWidth(960),
			),
		),
	)
}

func featureItem1() widget.Widget {
	return column.New(
		[]widget.Widget{
			row.New(
				[]widget.Widget{
					icon.New(icondata.RocketLaunchOutline, icon.Style(icon.Fill("#000000"), icon.Width(28), icon.Height(28))),
					text.New("No JavaScript Required", text.Style(text.Font(text.Size(20), text.Color("#000000"), text.Weight("500")), text.LineHeight(1.5))),
				},
				row.Gap(8),
				row.Flex(0),
			),
			text.New(
				"Write modern, reactive web applications entirely in Go. Gofred handles the WebAssembly layer so you can focus on building features — not managing JS interop.",
				text.Style(text.Font(text.Size(16), text.Color("#000000"), text.Weight("400")), text.LineHeight(1.5)),
			),
		},
		column.Gap(8),
	)
}

func featureItem2() widget.Widget {
	return column.New(
		[]widget.Widget{
			row.New(
				[]widget.Widget{
					icon.New(icondata.LightningBoltOutline, icon.Style(icon.Fill("#000000"), icon.Width(28), icon.Height(28))),
					text.New("High Performance", text.Style(text.Font(text.Size(20), text.Color("#000000"), text.Weight("500")), text.LineHeight(1.5))),
				},
				row.Gap(8),
				row.Flex(0),
			),
			text.New(
				"Leverages Go’s speed and efficiency, compiled to WebAssembly for near-native browser performance.",
				text.Style(text.Font(text.Size(16), text.Color("#000000"), text.Weight("400")), text.LineHeight(1.5)),
			),
		},
		column.Gap(8),
	)
}

func featureItem3() widget.Widget {
	return column.New(
		[]widget.Widget{
			row.New(
				[]widget.Widget{
					icon.New(icondata.PaletteOutline, icon.Style(icon.Fill("#000000"), icon.Width(28), icon.Height(28))),
					text.New("Declarative & Reactive", text.Style(text.Font(text.Size(20), text.Color("#000000"), text.Weight("500")), text.LineHeight(1.5))),
				},
				row.Gap(8),
				row.Flex(0),
			),
			text.New(
				"Inspired by React and Flutter, define your UI with simple Go interfaces. State changes automatically update the DOM — no manual wiring needed.",
				text.Style(text.Font(text.Size(16), text.Color("#000000"), text.Weight("400")), text.LineHeight(1.5)),
			),
		},
		column.Gap(8),
	)
}

func featureItem4() widget.Widget {
	return column.New(
		[]widget.Widget{
			row.New(
				[]widget.Widget{
					icon.New(icondata.Cellphone, icon.Style(icon.Fill("#000000"), icon.Width(28), icon.Height(28))),
					text.New("Responsive by Design", text.Style(text.Font(text.Size(20), text.Color("#000000"), text.Weight("500")), text.LineHeight(1.5))),
				},
				row.Gap(8),
				row.Flex(0),
			),
			text.New(
				"Gofred makes it simple to build layouts that adapt to any device size, ensuring great user experiences everywhere.",
				text.Style(text.Font(text.Size(16), text.Color("#000000"), text.Weight("400")), text.LineHeight(1.5)),
			),
		},
		column.Gap(8),
	)
}

func featureItem5() widget.Widget {
	return column.New(
		[]widget.Widget{
			row.New(
				[]widget.Widget{
					icon.New(icondata.PowerPlugOutline, icon.Style(icon.Fill("#000000"), icon.Width(28), icon.Height(28))),
					text.New("Extensible & Flexible", text.Style(text.Font(text.Size(20), text.Color("#000000"), text.Weight("500")), text.LineHeight(1.5))),
				},
				row.Gap(8),
				row.Flex(0),
			),
			text.New(
				"Use pure Go for core logic, but still extend with HTML, CSS, and JS when you need fine-grained control.",
				text.Style(text.Font(text.Size(16), text.Color("#000000"), text.Weight("400")), text.LineHeight(1.5)),
			),
		},
		column.Gap(8),
	)
}

func featureItem6() widget.Widget {
	return column.New(
		[]widget.Widget{
			row.New(
				[]widget.Widget{
					icon.New(icondata.Tools, icon.Style(icon.Fill("#000000"), icon.Width(28), icon.Height(28))),
					text.New("Built for Go Developers", text.Style(text.Font(text.Size(20), text.Color("#000000"), text.Weight("500")), text.LineHeight(1.5))),
				},
				row.Gap(8),
				row.Flex(0),
			),
			text.New(
				"No need to switch languages — build full web applications using the same patterns and tooling you already know in Go.",
				text.Style(text.Font(text.Size(16), text.Color("#000000"), text.Weight("400")), text.LineHeight(1.5)),
			),
		},
		column.Gap(8),
	)
}
