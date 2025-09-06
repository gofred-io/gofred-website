package drawer

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/foundation/button"
	. "github.com/gofred-io/gofred/foundation/center"
	. "github.com/gofred-io/gofred/foundation/column"
	. "github.com/gofred-io/gofred/foundation/container"
	. "github.com/gofred-io/gofred/foundation/drawer"
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

var leftDrawer *FDrawer

func RootDrawer() *FDrawer {
	if leftDrawer == nil {
		leftDrawer = buildLeftDrawer()
	}
	return leftDrawer
}

func buildLeftDrawer() *FDrawer {
	return Drawer(
		Container(
			Column(
				[]Widget{
					drawerHeader(),
					drawerContent(),
				},
			).Gap(0).Flex(1),
		).Flex(1),
	).ID("root-left-drawer").Width(AllBP(320)).Transition(0.3)
}

func drawerHeader() Widget {
	return Container(
		Row(
			[]Widget{
				drawerLogo(),
				Spacer(),
				Button(
					Icon(icondata.Close).
						Width(AllBP(20)).
						Height(AllBP(20)).
						Fill("#6B7280"),
				).OnClick(func(this Widget, e Event) {
					RootDrawer().Hide()
				}),
			},
		).Gap(12).Flex(1).CrossAxisAlignment(AxisAlignmentTypeCenter),
	).Padding(AllBP(LRTB(24, 12, 18, 14))).BorderColor("#E5E7EB").BorderWidth(0, 0, 1, 0).BorderStyle(BorderStyleTypeSolid)
}

func drawerLogo() Widget {
	return Row(
		[]Widget{
			Container(
				Text("🚀").FontSize(24),
			).Width(AllBP(32)).Height(AllBP(32)),
			Text("gofred").
				FontSize(20).
				FontColor("#1F2937").
				FontWeight("700").
				UserSelect(UserSelectTypeNone),
		},
	).Gap(16).CrossAxisAlignment(AxisAlignmentTypeCenter)
}

func drawerContent() Widget {
	return Container(
		Column(
			[]Widget{
				navigationSection(),
				externalLinksSection(),
				Spacer(),
				drawerFooter(),
			},
		).Gap(24).Flex(1),
	).Flex(1).Padding(AllBP(LRTB(24, 24, 24, 24)))
}

func navigationSection() Widget {
	return Column(
		[]Widget{
			sectionTitle("Navigation"),
			Spacer().Height(12),
			navItem("Home", "/", icondata.Home, false),
			navItem("Documentation", "/docs", icondata.FileDocument, false),
			navItem("Getting Started", "/docs/getting-started", icondata.Play, false),
			navItem("Core Concepts", "/docs/core-concepts", icondata.Lightbulb, false),
			navItem("Components", "/docs/components", icondata.Package, false),
			navItem("API Reference", "/docs/api", icondata.FileDocument, false),
		},
	).Gap(4)
}

func externalLinksSection() Widget {
	return Column(
		[]Widget{
			sectionTitle("Resources"),
			Spacer().Height(12),
			externalNavItem("GitHub", "https://github.com/gofred-io/gofred", icondata.Github),
			externalNavItem("Discussions", "https://github.com/gofred-io/gofred/discussions", icondata.Comment),
			externalNavItem("Examples", "https://github.com/gofred-io/examples", icondata.Lightbulb),
			externalNavItem("Community", "https://github.com/orgs/gofred-io/discussions", icondata.AccountGroup),
		},
	).Gap(4)
}

func drawerFooter() Widget {
	return Container(
		Center(
			Column(
				[]Widget{
					Text("Built with gofred").
						FontSize(12).
						FontColor("#9CA3AF").
						FontWeight("400"),
					Spacer().Height(4),
					Text("v1.0.0").
						FontSize(11).
						FontColor("#D1D5DB").
						FontWeight("400"),
				},
			).Gap(0).CrossAxisAlignment(AxisAlignmentTypeCenter),
		),
	).Padding(AllBP(LRTB(0, 0, 0, 0)))
}

func sectionTitle(title string) Widget {
	return Text(title).
		FontSize(12).
		FontColor("#6B7280").
		FontWeight("600").
		UserSelect(UserSelectTypeNone)
}

func navItem(title, href string, iconData icondata.IconData, isExternal bool) Widget {
	return Container(
		Link(
			Row(
				[]Widget{
					Icon(iconData).
						Width(AllBP(20)).
						Height(AllBP(20)).
						Fill("#6B7280"),
					Text(title).
						FontSize(14).
						FontColor("#374151").
						FontWeight("500").
						UserSelect(UserSelectTypeNone),
					Spacer(),
					func() Widget {
						if isExternal {
							return Icon(icondata.OpenInNew).
								Width(AllBP(16)).
								Height(AllBP(16)).
								Fill("#9CA3AF")
						}
						return Spacer()
					}(),
				},
			).Gap(12).Flex(1).CrossAxisAlignment(AxisAlignmentTypeCenter),
		).Href(href).OnClick(func(this Widget, e Event) {
			RootDrawer().Hide()
		}),
	).Padding(AllBP(LRTB(12, 12, 12, 12))).BorderRadius(8).BackgroundColor("transparent").BorderWidth(1, 1, 1, 1).BorderColor("transparent").BorderStyle(BorderStyleTypeSolid)
}

func externalNavItem(title, href string, iconData icondata.IconData) Widget {
	return Container(
		Link(
			Row(
				[]Widget{
					Icon(iconData).
						Width(AllBP(20)).
						Height(AllBP(20)).
						Fill("#6B7280"),
					Text(title).
						FontSize(14).
						FontColor("#374151").
						FontWeight("500").
						UserSelect(UserSelectTypeNone),
					Spacer(),
					Icon(icondata.OpenInNew).
						Width(AllBP(16)).
						Height(AllBP(16)).
						Fill("#9CA3AF"),
				},
			).Gap(12).Flex(1).CrossAxisAlignment(AxisAlignmentTypeCenter),
		).Flex(1).Href(href).NewTab(true),
	).Padding(AllBP(LRTB(12, 12, 12, 12))).BorderRadius(8).BackgroundColor("transparent").BorderWidth(1, 1, 1, 1).BorderColor("transparent").BorderStyle(BorderStyleTypeSolid)
}
