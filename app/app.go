package app

import (
	"github.com/gofred-io/gofred-website/app/pages/home"
	"github.com/gofred-io/gofred/widget"
)

func New() widget.Widget {
	return home.New()
}
