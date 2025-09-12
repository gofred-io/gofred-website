package theme

import "github.com/gofred-io/gofred/hooks"

type Theme string

const (
	ThemeLight Theme = "light"
	ThemeDark  Theme = "dark"
)

func init() {
	_, setThemeData := hooks.UseTheme()
	setThemeData(&lightTheme)
}
