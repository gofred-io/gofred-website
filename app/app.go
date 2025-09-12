package app

import (
	notfound "github.com/gofred-io/gofred-website/app/pages/404"
	"github.com/gofred-io/gofred-website/app/pages/docs"
	"github.com/gofred-io/gofred-website/app/pages/home"
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/foundation/router"
	"github.com/gofred-io/gofred/theme/theme_provider"
)

func New() application.BaseWidget {
	return theme_provider.New(
		router.New(
			router.Route("/", home.New),
			router.Route("/docs/:section", docs.New),
			router.NotFound(notfound.New),
		),
	)
}
