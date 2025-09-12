package theme

import "github.com/gofred-io/gofred/theme/theme_data"

var (
	lightTheme = theme_data.NewThemeData(
		theme_data.WithTextTheme(theme_data.TextTheme(
			theme_data.WithTextStyle(theme_data.TextStyleCollection{
				Primary: theme_data.TextStyle{
					FontSize:   theme_data.ThemeValue(16),
					FontWeight: theme_data.ThemeValue("400"),
					Color:      theme_data.ThemeValue("#1F2937"),
					FontFamily: theme_data.ThemeValue("Ubuntu"),
				},
			}),
		)),
		theme_data.WithThemeName(string(ThemeLight)),
	)
)

func LightTheme() *theme_data.ThemeData {
	return &lightTheme
}
