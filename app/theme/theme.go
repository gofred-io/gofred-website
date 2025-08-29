package theme

import (
	"github.com/gofred-io/gofred/hooks"
	"github.com/gofred-io/gofred/listenable"
)

type Theme string

const (
	ThemeLight Theme = "light"
	ThemeDark  Theme = "dark"
)

var (
	currentTheme, setCurrentTheme = hooks.UseState(ThemeLight)
)

func Get() Theme {
	return currentTheme.Value()
}

func Listenable() listenable.Listenable[Theme] {
	return currentTheme
}

func Set(theme Theme) {
	setCurrentTheme(theme)
}
