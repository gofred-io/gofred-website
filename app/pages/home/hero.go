package home

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/foundation/button"
	. "github.com/gofred-io/gofred/foundation/column"
	. "github.com/gofred-io/gofred/foundation/container"
	. "github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	. "github.com/gofred-io/gofred/foundation/link"
	. "github.com/gofred-io/gofred/foundation/row"
	. "github.com/gofred-io/gofred/foundation/spacer"
	. "github.com/gofred-io/gofred/foundation/text"
	. "github.com/gofred-io/gofred/options"
	. "github.com/gofred-io/gofred/options/spacing"
	. "github.com/gofred-io/gofred/widget"
)

func hero() Widget {
	return Container(
		Column(
			[]Widget{
				heroTitle(),
				heroSubtitle(),
				Spacer(),
				sampleCode(),
				Spacer(),
				cta(),
			},
		).Gap(16).
			Flex(1).
			CrossAxisAlignment(AxisAlignmentTypeCenter).
			MainAxisAlignment(AxisAlignmentTypeCenter),
	).BackgroundColor("#2B799B").
		Flex(1).
		Padding(AllBP(All(32)))
}

func heroTitle() Widget {
	return Text("Build responsive web apps in Go – no JavaScript required").
		FontSize(24).
		FontColor("#FFFFFF").
		FontWeight("600").
		UserSelect(UserSelectTypeNone)
}

func heroSubtitle() Widget {
	return Text("Write modern, reactive web applications with WebAssembly using only Go.").
		FontSize(18).
		FontColor("#FFFFFF").
		FontWeight("400").
		UserSelect(UserSelectTypeNone)
}

func sampleCode() Widget {
	closeButton := Container(nil).
		BackgroundColor("#FF5F58").
		Height(AllBP(16)).
		Width(AllBP(16)).
		BorderRadius(8)

	fullScreenButton := Container(nil).
		BackgroundColor("#FFBD2E").
		Height(AllBP(16)).
		Width(AllBP(16)).
		BorderRadius(8)

	minimizeButton := Container(nil).
		BackgroundColor("#28C941").
		Height(AllBP(16)).
		Width(AllBP(16)).
		BorderRadius(8)

	windowTitle := Text("main.go").
		FontSize(16).
		FontColor("#FFFFFF").
		FontWeight("500").
		UserSelect(UserSelectTypeNone)

	titleBar := Row(
		[]Widget{
			Row(
				[]Widget{
					closeButton,
					fullScreenButton,
					minimizeButton,
				},
			).MainAxisAlignment(AxisAlignmentTypeStart).
				CrossAxisAlignment(AxisAlignmentTypeStart).
				Gap(8).
				Flex(1),
			windowTitle,
			Spacer(),
		},
	).MainAxisSize(AxisSizeTypeMax).
		MainAxisAlignment(AxisAlignmentTypeStart).
		CrossAxisAlignment(AxisAlignmentTypeStart).
		Gap(8)

	packageName := Row(
		[]Widget{
			Text("package").
				FontSize(14).
				FontColor("#4F8CBF").
				FontWeight("400").
				UserSelect(UserSelectTypeNone),
			Text("main").
				FontSize(14).
				FontColor("#49B8A2").
				FontWeight("400").
				UserSelect(UserSelectTypeNone),
		},
	).MainAxisAlignment(AxisAlignmentTypeStart).
		CrossAxisAlignment(AxisAlignmentTypeStart).
		Gap(6)

	imports := Column(
		[]Widget{
			Row(
				[]Widget{
					Text("import").
						FontSize(14).
						FontColor("#AD77A8").
						FontWeight("400").
						UserSelect(UserSelectTypeNone),
					Text("(").
						FontSize(14).
						FontColor("#E0BD04").
						FontWeight("400").
						UserSelect(UserSelectTypeNone),
				},
			).MainAxisAlignment(AxisAlignmentTypeStart).
				CrossAxisAlignment(AxisAlignmentTypeStart).
				Gap(6),
			Row(
				[]Widget{
					Spacer(),
					Text("\"github.com/gofred-io/gofred-website/app\"").
						FontSize(14).
						FontColor("#B9836D").
						FontWeight("400").
						UserSelect(UserSelectTypeNone),
				},
			).MainAxisAlignment(AxisAlignmentTypeStart).
				CrossAxisAlignment(AxisAlignmentTypeStart).
				Gap(6),
			Row(
				[]Widget{
					Spacer(),
					Text("\"github.com/gofred-io/gofred/application\"").
						FontSize(14).
						FontColor("#B9836D").
						FontWeight("400").
						UserSelect(UserSelectTypeNone),
				},
			).MainAxisAlignment(AxisAlignmentTypeStart).
				CrossAxisAlignment(AxisAlignmentTypeStart).
				Gap(6),
			Row(
				[]Widget{
					Text(")").
						FontSize(14).
						FontColor("#E0BD04").
						FontWeight("400").
						UserSelect(UserSelectTypeNone),
				},
			),
		},
	).MainAxisAlignment(AxisAlignmentTypeStart).
		CrossAxisAlignment(AxisAlignmentTypeStart).
		Gap(6)

	mainFunc := Column(
		[]Widget{
			Row(
				[]Widget{
					Text("func").
						FontSize(14).
						FontColor("#4F8CBF").
						FontWeight("400").
						UserSelect(UserSelectTypeNone),
					Row(
						[]Widget{
							Text("main").
								FontSize(14).
								FontColor("#DCDCAA").
								FontWeight("400").
								UserSelect(UserSelectTypeNone),
							Text("()").
								FontSize(14).
								FontColor("#E0BD04").
								FontWeight("400").
								UserSelect(UserSelectTypeNone),
						},
					).MainAxisAlignment(AxisAlignmentTypeStart).
						CrossAxisAlignment(AxisAlignmentTypeStart),
					Text("{").
						FontSize(14).
						FontColor("#E0BD04").
						FontWeight("400").
						UserSelect(UserSelectTypeNone),
				},
			).MainAxisAlignment(AxisAlignmentTypeStart).
				CrossAxisAlignment(AxisAlignmentTypeStart).
				Gap(6),
			Row(
				[]Widget{
					Spacer(),
					Row(
						[]Widget{
							Text("application").
								FontSize(14).
								FontColor("#8AD1DE").
								FontWeight("400").
								UserSelect(UserSelectTypeNone),
							Text(".").
								FontSize(14).
								FontColor("#BCBCBC").
								FontWeight("400").
								UserSelect(UserSelectTypeNone),
							Text("Run").
								FontSize(14).
								FontColor("#DCDCAA").
								FontWeight("400").
								UserSelect(UserSelectTypeNone),
							Text("(").
								FontSize(14).
								FontColor("#AD77A8").
								FontWeight("400").
								UserSelect(UserSelectTypeNone),
							Text("app").
								FontSize(14).
								FontColor("#8AD1DE").
								FontWeight("400").
								UserSelect(UserSelectTypeNone),
							Text(".").
								FontSize(14).
								FontColor("#BCBCBC").
								FontWeight("400").
								UserSelect(UserSelectTypeNone),
							Text("New").
								FontSize(14).
								FontColor("#DCDCAA").
								FontWeight("400").
								UserSelect(UserSelectTypeNone),
							Text("()").
								FontSize(14).
								FontColor("#4F8CBF").
								FontWeight("400").
								UserSelect(UserSelectTypeNone),
							Text(")").
								FontSize(14).
								FontColor("#AD77A8").
								FontWeight("400").
								UserSelect(UserSelectTypeNone),
						},
					),
				},
			).MainAxisAlignment(AxisAlignmentTypeStart).
				CrossAxisAlignment(AxisAlignmentTypeStart).
				Gap(6),
			Text("}").
				FontSize(14).
				FontColor("#E0BD04").
				FontWeight("400").
				UserSelect(UserSelectTypeNone),
		},
	).MainAxisAlignment(AxisAlignmentTypeStart).
		CrossAxisAlignment(AxisAlignmentTypeStart).
		Gap(6)

	return Container(
		Column(
			[]Widget{
				titleBar,
				packageName,
				Spacer(),
				imports,
				Spacer(),
				mainFunc,
			},
		).MainAxisAlignment(AxisAlignmentTypeStart).
			Gap(6).
			Flex(1),
	).BackgroundColor("#282A36").
		Height(AllBP(240)).
		BorderRadius(8).
		Padding(AllBP(All(16))).
		MaxWidth(AllBP(540)).
		WidthP(AllBP(0.5), XS(1.0), SM(1.0), MD(0.75))
}

func cta() Widget {
	return Row(
		[]Widget{
			Link(
				Button(
					Text("Get Started").
						FontFamily("Ubuntu").
						FontSize(14).
						FontColor("#FFFFFF").
						FontWeight("500").
						UserSelect(UserSelectTypeNone),
				).Width(AllBP(100)),
			).Href("/docs/installation"),
			Link(
				Button(
					Row(
						[]Widget{
							Icon(icondata.Github).
								Fill("#FFFFFF").
								Height(AllBP(16)).
								Width(AllBP(16)),
							Text("GitHub").
								FontSize(14).
								FontColor("#FFFFFF").
								FontWeight("500").
								FontFamily("Ubuntu").
								UserSelect(UserSelectTypeNone),
						},
					).MainAxisAlignment(AxisAlignmentTypeCenter).
						CrossAxisAlignment(AxisAlignmentTypeCenter).
						Gap(6),
				).BackgroundColor("#151B23").
					Width(AllBP(100)),
			).Href("https://github.com/gofred-io/gofred").
				NewTab(true),
		},
	).Gap(8)
}
