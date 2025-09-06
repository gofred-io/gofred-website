package notfound

import (
	"github.com/gofred-io/gofred-website/app/components/footer"
	"github.com/gofred-io/gofred-website/app/components/header"
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/foundation/button"
	. "github.com/gofred-io/gofred/foundation/column"
	. "github.com/gofred-io/gofred/foundation/container"
	. "github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	. "github.com/gofred-io/gofred/foundation/link"
	. "github.com/gofred-io/gofred/foundation/router"
	. "github.com/gofred-io/gofred/foundation/row"
	. "github.com/gofred-io/gofred/foundation/spacer"
	. "github.com/gofred-io/gofred/foundation/text"
	. "github.com/gofred-io/gofred/options"
	. "github.com/gofred-io/gofred/options/spacing"
	. "github.com/gofred-io/gofred/widget"
)

func NotFoundPage(params RouteParams) Widget {
	return Column(
		[]Widget{
			header.Get(),
			mainContent(),
			footer.Get(),
		},
	).Flex(1)
}

func mainContent() Widget {
	return Container(
		Column(
			[]Widget{
				errorIcon(),
				Spacer().Height(24),
				errorTitle(),
				Spacer().Height(16),
				errorDescription(),
				Spacer().Height(32),
				actionButtons(),
			},
		).Gap(16).Flex(1).CrossAxisAlignment(AxisAlignmentTypeCenter).MainAxisAlignment(AxisAlignmentTypeCenter),
	).BackgroundColor("#F8F9FA").Flex(1).Padding(AllBP(All(32)))
}

func errorIcon() Widget {
	return Container(
		Icon(icondata.Alert).
			Width(AllBP(120)).
			Height(AllBP(120)).
			Fill("#E74C3C"),
	).BackgroundColor("#FFFFFF").BorderRadius(48).Padding(AllBP(All(20))).BorderColor("#E74C3C").BorderWidth(3, 3, 3, 3).BorderStyle(BorderStyleTypeSolid)
}

func errorTitle() Widget {
	return Text("404").
		FontSize(72).
		FontColor("#2C3E50").
		FontWeight("700").
		UserSelect(UserSelectTypeNone)
}

func errorDescription() Widget {
	return Column(
		[]Widget{
			Text("Oops! The page you're looking for doesn't exist.").
				FontSize(24).
				FontColor("#34495E").
				FontWeight("500").
				UserSelect(UserSelectTypeNone),
			Spacer().Height(8),
			Text("It might have been moved, deleted, or you entered the wrong URL.").
				FontSize(16).
				FontColor("#7F8C8D").
				FontWeight("400").
				UserSelect(UserSelectTypeNone),
		},
	).Gap(8).CrossAxisAlignment(AxisAlignmentTypeCenter)
}

func actionButtons() Widget {
	return Row(
		[]Widget{
			backButton(),
			homeButton(),
		},
	).Gap(16).CrossAxisAlignment(AxisAlignmentTypeCenter)
}

func backButton() Widget {
	return Button(
		Container(
			Row(
				[]Widget{
					Icon(icondata.ChevronLeft).
						Width(AllBP(20)).
						Height(AllBP(20)).
						Fill("#FFFFFF"),
					Text("Go Back").
						FontSize(16).
						FontColor("#FFFFFF").
						FontWeight("500").
						UserSelect(UserSelectTypeNone),
				},
			).Gap(8).CrossAxisAlignment(AxisAlignmentTypeCenter),
		).Padding(AllBP(Right(4))),
	).OnClick(func(this Widget, e Event) {
		Context().GoBack()
	})
}

func homeButton() Widget {
	return Link(
		Button(
			Container(
				Row(
					[]Widget{
						Icon(icondata.Home).
							Width(AllBP(20)).
							Height(AllBP(20)).
							Fill("#FFFFFF"),
						Text("Go Home").
							FontSize(16).
							FontColor("#FFFFFF").
							FontWeight("500").
							UserSelect(UserSelectTypeNone),
					},
				).Gap(8).CrossAxisAlignment(AxisAlignmentTypeCenter),
			).Padding(AllBP(Right(4))),
		),
	).Href("/")
}
