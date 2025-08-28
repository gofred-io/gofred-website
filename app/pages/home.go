package pages

import (
	"fmt"

	"github.com/gofred-io/gofred/column"
	"github.com/gofred-io/gofred/container"
	iconbutton "github.com/gofred-io/gofred/icon_button"
	icondata "github.com/gofred-io/gofred/icon_data"
	"github.com/gofred-io/gofred/image"
	"github.com/gofred-io/gofred/row"
	"github.com/gofred-io/gofred/style"
	"github.com/gofred-io/gofred/text"
	"github.com/gofred-io/gofred/widget"
)

func Home() widget.Widget {
	return column.New(
		[]widget.Widget{
			header(),
		},
	)
}

func header() widget.Widget {
	return container.New(
		row.New(
			[]widget.Widget{
				iconbutton.New(
					icondata.HamburgerMenu,
					iconbutton.OnClick(func(this widget.Widget) {
						fmt.Println("hamburger menu clicked")
					}),
				),
				image.New(
					"img/gofred.jpg",
					image.Width(48),
					image.Height(48),
				),
				text.New(
					"gofred",
					text.Style(
						text.Font(
							text.Size(20),
							text.Color("#373F51"),
							text.Weight("600"),
						),
						text.UserSelect(style.UserSelectTypeNone),
					),
				),
			},
			row.Gap(8),
			row.CrossAxisAlignment(style.AlignItemsTypeCenter),
		),
		container.Style(
			container.Height(44),
			container.Background(style.Background{
				Color: "#FFFFFF",
			}),
			container.Padding(style.Padding{
				Top:    8,
				Bottom: 8,
				Left:   16,
				Right:  16,
			}),
		),
	)
}
