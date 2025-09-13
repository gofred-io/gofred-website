package theme

import (
	"github.com/gofred-io/gofred/theme/color"
	"github.com/gofred-io/gofred/theme/theme_data"
	theme "github.com/gofred-io/gofred/theme/theme_data"
	style "github.com/gofred-io/gofred/theme/theme_style"
)

var (
	darkTheme = &theme_data.ThemeData{
		Name: string(ThemeDark),
		TextTheme: theme.TextTheme{
			CodeBlockStyle: style.TextStyleCollection{
				Primary: style.TextStyle{
					FontSize:   style.ThemeValue(14),
					FontWeight: style.ThemeValue("400"),
					Color:      style.ThemeValue(color.From(0xe0e4e4)),
					FontFamily: style.ThemeValue("ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, \"Liberation Mono\", \"Courier New\", monospace"),
				},
			},
			TextStyle: style.TextStyleCollection{
				Primary: style.TextStyle{
					FontSize:   style.ThemeValue(16),
					FontWeight: style.ThemeValue("400"),
					Color:      style.ThemeValue(color.From(0xe0e4e4)),
					FontFamily: style.ThemeValue("Ubuntu"),
				},
				Secondary: style.TextStyle{
					FontSize:   style.ThemeValue(16),
					FontWeight: style.ThemeValue("400"),
					Color:      style.ThemeValue(color.From(0x9e9e9e)),
					FontFamily: style.ThemeValue("Ubuntu"),
				},
				Tertiary: style.TextStyle{
					FontSize:   style.ThemeValue(16),
					FontWeight: style.ThemeValue("400"),
					Color:      style.ThemeValue(color.From(0xececec)),
					FontFamily: style.ThemeValue("Ubuntu"),
				},
			},
		},
		BoxTheme: theme.BoxTheme{
			CodeBlockStyle: style.ContainerStyleCollection{
				Primary: style.ContainerStyle{
					BackgroundColor: style.ThemeValue(color.From(0x141c27)),
				},
			},
			ContainerStyle: style.ContainerStyleCollection{
				Primary: style.ContainerStyle{
					BackgroundColor: style.ThemeValue(color.From(0x212a32)),
					BorderColor:     style.ThemeValue(color.From(0x343f49)),
				},
				Secondary: style.ContainerStyle{
					BackgroundColor: style.ThemeValue(color.From(0x26313a)),
					BorderColor:     style.ThemeValue(color.From(0x343f49)),
				},
				Tertiary: style.ContainerStyle{
					BackgroundColor: style.ThemeValue(color.From(0x0f141b)),
					BorderColor:     style.ThemeValue(color.From(0x4d4f54)),
				},
			},
		},
	}
)

func DarkTheme() *theme_data.ThemeData {
	return darkTheme
}
