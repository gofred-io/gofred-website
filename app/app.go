package app

import (
	notfound "github.com/gofred-io/gofred-website/app/pages/404"
	"github.com/gofred-io/gofred-website/app/pages/docs"
	"github.com/gofred-io/gofred-website/app/pages/home"
	"github.com/gofred-io/gofred/foundation/router"
	"github.com/gofred-io/gofred/widget"
)

func New() widget.BaseWidget {
	return router.New(
		router.Route("/", home.New),
		router.Route("/docs", docs.New),
		router.NotFound(notfound.New),
	)
}
