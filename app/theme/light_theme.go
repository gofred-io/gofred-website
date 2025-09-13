package theme

import (
	"github.com/gofred-io/gofred/theme/color"
	theme "github.com/gofred-io/gofred/theme/theme_data"
	style "github.com/gofred-io/gofred/theme/theme_style"
)

var (
	lightTheme = &theme.ThemeData{
		Name: string(ThemeLight),
		TextTheme: theme.TextTheme{
			CodeBlockStyle: style.TextStyleCollection{
				Primary: style.TextStyle{
					FontSize:   style.ThemeValue(14),
					FontWeight: style.ThemeValue("400"),
					Color:      style.ThemeValue(color.From(0xF3F4F6)),
					FontFamily: style.ThemeValue("ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, \"Liberation Mono\", \"Courier New\", monospace"),
				},
			},
			TextStyle: style.TextStyleCollection{
				Primary: style.TextStyle{
					FontSize:   style.ThemeValue(16),
					FontWeight: style.ThemeValue("400"),
					Color:      style.ThemeValue(color.From(0x1F2937)),
					FontFamily: style.ThemeValue("Ubuntu"),
				},
				Secondary: style.TextStyle{
					FontSize:   style.ThemeValue(16),
					FontWeight: style.ThemeValue("400"),
					Color:      style.ThemeValue(color.From(0x6a6b6c)),
					FontFamily: style.ThemeValue("Ubuntu"),
				},
				Tertiary: style.TextStyle{
					FontSize:   style.ThemeValue(16),
					FontWeight: style.ThemeValue("400"),
					Color:      style.ThemeValue(color.From(0xfafafa)),
					FontFamily: style.ThemeValue("Ubuntu"),
				},
			},
		},
		BoxTheme: theme.BoxTheme{
			CodeBlockStyle: style.ContainerStyleCollection{
				Primary: style.ContainerStyle{
					BackgroundColor: style.ThemeValue(color.From(0x1F2937)),
				},
			},
			ContainerStyle: style.ContainerStyleCollection{
				Primary: style.ContainerStyle{
					BackgroundColor: style.ThemeValue(color.From(0xfdfdfd)),
					BorderColor:     style.ThemeValue(color.From(0xececec)),
				},
				Secondary: style.ContainerStyle{
					BackgroundColor: style.ThemeValue(color.From(0xf8f8f8)),
					BorderColor:     style.ThemeValue(color.From(0xE5E7EB)),
				},
				Tertiary: style.ContainerStyle{
					BackgroundColor: style.ThemeValue(color.From(0x1F2937)),
					BorderColor:     style.ThemeValue(color.From(0x374151)),
				},
			},
		},
	}
)

func LightTheme() *theme.ThemeData {
	return lightTheme
}
