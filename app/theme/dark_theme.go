package theme

import "github.com/gofred-io/gofred/theme/theme_data"

var (
	darkTheme = theme_data.NewThemeData(
		theme_data.WithTextTheme(theme_data.TextTheme(
			theme_data.WithTextStyle(theme_data.TextStyleCollection{
				Primary: theme_data.TextStyle{
					FontSize:   theme_data.ThemeValue(16),
					FontWeight: theme_data.ThemeValue("400"),
					Color:      theme_data.ThemeValue("#FFFFFF"),
					FontFamily: theme_data.ThemeValue("Ubuntu"),
				},
			}),
		)),
		theme_data.WithThemeName(string(ThemeDark)),
	)
)

func DarkTheme() *theme_data.ThemeData {
	return &darkTheme
}
