package app

import (
	. "github.com/gofred-io/gofred-website/app/pages/404"
	. "github.com/gofred-io/gofred-website/app/pages/docs"
	. "github.com/gofred-io/gofred-website/app/pages/home"
	. "github.com/gofred-io/gofred/foundation/router"
	"github.com/gofred-io/gofred/widget"
)

func GofredApp() widget.Widget {
	return Router(
		Route("/", Home),
		Route("/docs/:section", Docs),
		NotFound(NotFoundPage),
	)
}
