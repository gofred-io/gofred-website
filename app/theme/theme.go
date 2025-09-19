package theme

import (
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/theme/theme_data"
)

type Theme string

const (
	ThemeLight Theme = "light"
	ThemeDark  Theme = "dark"
)

var (
	themeHook, setThemeData = hooks.UseTheme()
)

func init() {
	setThemeData(lightTheme)
}

func Data() *theme_data.ThemeData {
	return themeHook.ThemeData()
}
