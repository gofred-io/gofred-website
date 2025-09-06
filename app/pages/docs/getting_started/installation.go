package getting_started

import (
	. "github.com/gofred-io/gofred/breakpoint"
	. "github.com/gofred-io/gofred/foundation/button"
	. "github.com/gofred-io/gofred/foundation/code_block"
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

func InstallationContent() Widget {
	return Container(
		Column(
			[]Widget{
				installationPageHeader(),
				Spacer(),
				installationPageContent(),
			},
		).Gap(16).
			Flex(1),
	).Flex(1).
		Padding(AllBP(All(32)))
}

func installationPageHeader() Widget {
	return Column(
		[]Widget{
			Text("Installation").
				FontSize(32).
				FontColor("#1F2937").
				FontWeight("700"),
			Text("Get started with gofred by installing the framework and setting up your development environment.").
				FontSize(18).
				FontColor("#6B7280").
				FontWeight("400"),
		},
	).Gap(8)
}

func installationPageContent() Widget {
	return Column(
		[]Widget{
			contentSection("Prerequisites", "Before you begin, make sure you have the following installed on your system:"),
			prerequisitesList(),
			Spacer(),
			contentSection("Installation", "Install gofred using Go modules:"),
			CodeBlock("go get github.com/gofred-io/gofred"),
			Spacer(),
			contentSection("Next Steps", "Now that you have gofred installed, you can:"),
			installationNextSteps(),
			Spacer(),
			navigationButtons("/docs", "/docs/quick-start"),
		},
	).Gap(16)
}

func contentSection(title, description string) Widget {
	return Column(
		[]Widget{
			Text(title).
				FontSize(24).
				FontColor("#1F2937").
				FontWeight("600"),
			Text(description).
				FontSize(16).
				FontColor("#6B7280").
				FontWeight("400"),
		},
	).Gap(8)
}

func prerequisitesList() Widget {
	items := []string{
		"Go 1.21 or later",
		"A modern web browser with WebAssembly support",
		"Basic knowledge of Go programming",
	}

	var listItems []Widget
	for _, item := range items {
		listItems = append(listItems, listItem(item))
	}

	return Column(
		listItems,
	).Gap(8)
}

func listItem(itemText string) Widget {
	return Row(
		[]Widget{
			Icon(icondata.Check).
				Width(AllBP(16)).
				Height(AllBP(16)).
				Fill("#10B981"),
			Spacer(),
			Text(itemText).
				FontSize(16).
				FontColor("#374151").
				FontWeight("400"),
		},
	).Gap(8).
		CrossAxisAlignment(AxisAlignmentTypeCenter)
}

func installationNextSteps() Widget {
	steps := []struct {
		title       string
		description string
		href        string
	}{
		{
			title:       "Quick Start",
			description: "Create your first gofred application",
			href:        "/docs/quick-start",
		},
		{
			title:       "Build Your First App",
			description: "Create a simple application step by step",
			href:        "/docs/first-app",
		},
		{
			title:       "Project Structure",
			description: "Explore the project structure",
			href:        "/docs/project-structure",
		},
	}

	var stepItems []Widget
	for _, step := range steps {
		stepItems = append(stepItems, nextStepItem(step.title, step.description, step.href))
	}

	return Column(
		stepItems,
	).Gap(12)
}

func nextStepItem(title, description, href string) Widget {
	return Link(
		Container(
			Row(
				[]Widget{
					Column(
						[]Widget{
							Text(title).
								FontSize(16).
								FontColor("#2B799B").
								FontWeight("500"),
							Text(description).
								FontSize(14).
								FontColor("#6B7280").
								FontWeight("400"),
						},
					).Gap(4).
						Flex(1),
					Icon(icondata.ChevronRight).
						Width(AllBP(20)).
						Height(AllBP(20)).
						Fill("#9CA3AF"),
				},
			).Gap(12).
				Flex(1).
				CrossAxisAlignment(AxisAlignmentTypeCenter),
		).Padding(AllBP(All(16))).
			BackgroundColor("#FFFFFF").
			BorderRadius(8).
			BorderColor("#E5E7EB").
			BorderWidth(1, 1, 1, 1).
			BorderStyle(BorderStyleTypeSolid),
	).Href(href)
}

func navigationButtons(previousHref, nextHref string) Widget {
	return Row(
		[]Widget{
			Link(
				Button(
					Row(
						[]Widget{
							Icon(icondata.ChevronLeft).
								Width(AllBP(16)).
								Height(AllBP(16)).
								Fill("#FFFFFF"),
							Text("Previous").
								FontSize(14).
								FontColor("#FFFFFF").
								FontWeight("500"),
						},
					).Gap(8).
						CrossAxisAlignment(AxisAlignmentTypeCenter),
				).BackgroundColor("#6B7280").
					Width(AllBP(120)),
			).Href(previousHref),
			Spacer(),
			Link(
				Button(
					Row(
						[]Widget{
							Text("Next").
								FontSize(14).
								FontColor("#FFFFFF").
								FontWeight("500"),
							Icon(icondata.ChevronRight).
								Width(AllBP(16)).
								Height(AllBP(16)).
								Fill("#FFFFFF"),
						},
					).Gap(8).
						CrossAxisAlignment(AxisAlignmentTypeCenter),
				).BackgroundColor("#2B799B").
					Width(AllBP(120)),
			).Href(nextHref),
		},
	).Flex(1)
}
