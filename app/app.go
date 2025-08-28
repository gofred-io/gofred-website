package app

import (
	"github.com/gofred-io/gofred-website/app/pages"
	"github.com/gofred-io/gofred/widget"
)

func New() widget.Widget {
	return pages.Home()
}
