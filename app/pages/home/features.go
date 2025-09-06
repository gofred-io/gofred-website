package home

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/foundation/center"
	. "github.com/gofred-io/gofred/foundation/column"
	. "github.com/gofred-io/gofred/foundation/container"
	. "github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	. "github.com/gofred-io/gofred/foundation/row"
	. "github.com/gofred-io/gofred/foundation/text"
	. "github.com/gofred-io/gofred/options"
	. "github.com/gofred-io/gofred/options/spacing"
	. "github.com/gofred-io/gofred/widget"
)

func features() Widget {
	return Container(
		Column(
			[]Widget{
				featureTitle(),
				featureList(),
			},
		).Gap(32).
			Flex(1),
	).Padding(AllBP(All(64)))
}

func featureTitle() Widget {
	return Center(
		Text("Why gofred?").
			FontSize(32).
			FontColor("#000000").
			FontWeight("500").
			UserSelect(UserSelectTypeNone),
	)
}

func featureList() Widget {
	return Center(
		Container(
			Column(
				[]Widget{
					Row(
						[]Widget{
							featureItem1(),
							featureItem2(),
							featureItem3(),
						},
					).Gap(48),
					Row(
						[]Widget{
							featureItem4(),
							featureItem5(),
							featureItem6(),
						},
					).Gap(48),
				},
			).Gap(48),
		).MaxWidth(AllBP(960)).
			WidthP(AllBP(0.75)),
	)
}

func featureItem1() Widget {
	return Column(
		[]Widget{
			Row(
				[]Widget{
					Icon(icondata.RocketLaunchOutline).
						Fill("#000000").
						Width(AllBP(28)).
						Height(AllBP(28)),
					Text("No JavaScript Required").
						FontSize(20).
						FontColor("#000000").
						FontWeight("500").
						LineHeight(1.5),
				},
			).Gap(8).
				CrossAxisAlignment(AxisAlignmentTypeCenter),
			Text("Write modern, reactive web applications entirely in Go. Gofred handles the WebAssembly layer so you can focus on building features — not managing JS interop.").
				FontSize(16).
				FontColor("#000000").
				FontWeight("400").
				LineHeight(1.5),
		},
	).Gap(8)
}

func featureItem2() Widget {
	return Column(
		[]Widget{
			Row(
				[]Widget{
					Icon(icondata.LightningBoltOutline).
						Fill("#000000").
						Width(AllBP(28)).
						Height(AllBP(28)),
					Text("High Performance").
						FontSize(20).
						FontColor("#000000").
						FontWeight("500").
						LineHeight(1.5),
				},
			).Gap(8).
				CrossAxisAlignment(AxisAlignmentTypeCenter),
			Text("Leverages Go's speed and efficiency, compiled to WebAssembly for near-native browser performance.").
				FontSize(16).
				FontColor("#000000").
				FontWeight("400").
				LineHeight(1.5),
		},
	).Gap(8)
}

func featureItem3() Widget {
	return Column(
		[]Widget{
			Row(
				[]Widget{
					Icon(icondata.PaletteOutline).
						Fill("#000000").
						Width(AllBP(28)).
						Height(AllBP(28)),
					Text("Declarative & Reactive").
						FontSize(20).
						FontColor("#000000").
						FontWeight("500").
						LineHeight(1.5),
				},
			).Gap(8).
				CrossAxisAlignment(AxisAlignmentTypeCenter),
			Text("Inspired by React and Flutter, define your UI with simple Go interfaces. State changes automatically update the DOM — no manual wiring needed.").
				FontSize(16).
				FontColor("#000000").
				FontWeight("400").
				LineHeight(1.5),
		},
	).Gap(8)
}

func featureItem4() Widget {
	return Column(
		[]Widget{
			Row(
				[]Widget{
					Icon(icondata.Cellphone).
						Fill("#000000").
						Width(AllBP(28)).
						Height(AllBP(28)),
					Text("Responsive by Design").
						FontSize(20).
						FontColor("#000000").
						FontWeight("500").
						LineHeight(1.5),
				},
			).Gap(8).
				CrossAxisAlignment(AxisAlignmentTypeCenter),
			Text("Gofred makes it simple to build layouts that adapt to any device size, ensuring great user experiences everywhere.").
				FontSize(16).
				FontColor("#000000").
				FontWeight("400").
				LineHeight(1.5),
		},
	).Gap(8)
}

func featureItem5() Widget {
	return Column(
		[]Widget{
			Row(
				[]Widget{
					Icon(icondata.PowerPlugOutline).
						Fill("#000000").
						Width(AllBP(28)).
						Height(AllBP(28)),
					Text("Extensible & Flexible").
						FontSize(20).
						FontColor("#000000").
						FontWeight("500").
						LineHeight(1.5),
				},
			).Gap(8).
				CrossAxisAlignment(AxisAlignmentTypeCenter),
			Text("Use pure Go for core logic, but still extend with HTML, CSS, and JS when you need fine-grained control.").
				FontSize(16).
				FontColor("#000000").
				FontWeight("400").
				LineHeight(1.5),
		},
	).Gap(8)
}

func featureItem6() Widget {
	return Column(
		[]Widget{
			Row(
				[]Widget{
					Icon(icondata.Tools).
						Fill("#000000").
						Width(AllBP(28)).
						Height(AllBP(28)),
					Text("Built for Go Developers").
						FontSize(20).
						FontColor("#000000").
						FontWeight("500").
						LineHeight(1.5),
				},
			).Gap(8).
				CrossAxisAlignment(AxisAlignmentTypeCenter),
			Text("No need to switch languages — build full web applications using the same patterns and tooling you already know in Go.").
				FontSize(16).
				FontColor("#000000").
				FontWeight("400").
				LineHeight(1.5),
		},
	).Gap(8)
}
