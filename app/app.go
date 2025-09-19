package app

import (
	"github.com/gofred-io/gofred-website/app/components/drawer"
	notfound "github.com/gofred-io/gofred-website/app/pages/404"
	"github.com/gofred-io/gofred-website/app/pages/docs"
	docsDrawer "github.com/gofred-io/gofred-website/app/pages/docs/drawer"
	"github.com/gofred-io/gofred-website/app/pages/home"
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/foundation/router"
	"github.com/gofred-io/gofred/foundation/scaffold"
	"github.com/gofred-io/gofred/theme/theme_provider"
)

func New() application.BaseWidget {
	return scaffold.New(
		theme_provider.New(
			router.New(
				router.Route("/", home.New),
				router.Route("/docs/:section", docs.New),
				router.NotFound(notfound.New),
			),
		),
		scaffold.Drawer(drawer.New()),
		scaffold.Drawer(docsDrawer.New()),
	)
}
