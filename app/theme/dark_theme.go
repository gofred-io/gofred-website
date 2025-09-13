package theme

import (
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
	"github.com/gofred-io/gofred/theme/color"
	"github.com/gofred-io/gofred/theme/theme_data"
	td "github.com/gofred-io/gofred/theme/theme_data"
	style "github.com/gofred-io/gofred/theme/theme_style"
)

var (
	darkTheme = &theme_data.ThemeData{
		Name: string(ThemeDark),
		BoxTheme: td.BoxTheme{
			CodeBlockStyle: style.ContainerStyleCollection{
				Primary: style.ContainerStyle{
					BackgroundColor: style.ThemeValue(color.From(0x141c27ff)),
				},
			},
			ContainerStyle: style.ContainerStyleCollection{
				Primary: style.ContainerStyle{
					BackgroundColor: style.ThemeValue(color.From(0x212a32ff)),
					BorderColor:     style.ThemeValue(color.From(0x343f49ff)),
				},
				Secondary: style.ContainerStyle{
					BackgroundColor: style.ThemeValue(color.From(0x26313aff)),
					BorderColor:     style.ThemeValue(color.From(0x343f49ff)),
				},
				Tertiary: style.ContainerStyle{
					BackgroundColor: style.ThemeValue(color.From(0x0f141bff)),
					BorderColor:     style.ThemeValue(color.From(0x4d4f54ff)),
				},
			},
		},
		ButtonTheme: td.ButtonTheme{
			ButtonStyle: style.ButtonStyleCollection{
				Primary: style.ButtonStyle{
					BackgroundColor: style.ThemeValue(color.From(0x1976d2ff)),
					BorderColor:     style.ThemeValue(color.From(0x00000000)),
					BorderRadius:    style.ThemeValue(8),
					BorderWidth:     style.ThemeValue(spacing.All(1)),
					BorderStyle:     style.ThemeValue(theme.BorderStyleTypeSolid),
					TextStyle: style.TextStyle{
						FontSize:   style.ThemeValue(16),
						FontWeight: style.ThemeValue("400"),
						Color:      style.ThemeValue(color.From(0xFFFFFFFF)),
						FontFamily: style.ThemeValue("Ubuntu"),
					},
				},
				Secondary: style.ButtonStyle{
					BackgroundColor: style.ThemeValue(color.From(0xe0e4e4ff)),
					BorderColor:     style.ThemeValue(color.From(0xE5E7EBFF)),
					BorderWidth:     style.ThemeValue(spacing.All(1)),
					BorderStyle:     style.ThemeValue(theme.BorderStyleTypeSolid),
					BorderRadius:    style.ThemeValue(8),
					TextStyle: style.TextStyle{
						FontSize:   style.ThemeValue(16),
						FontWeight: style.ThemeValue("400"),
						Color:      style.ThemeValue(color.From(0x1F2937ff)),
						FontFamily: style.ThemeValue("Ubuntu"),
					},
				},
			},
			IconButtonStyle: style.ButtonStyleCollection{
				Primary: style.ButtonStyle{
					BackgroundColor: style.ThemeValue(color.From(0x00000000)),
					BorderColor:     style.ThemeValue(color.From(0x00000000)),
					BorderRadius:    style.ThemeValue(8),
					Fill:            style.ThemeValue(color.From(0xe0e4e4ff)),
				},
				Secondary: style.ButtonStyle{
					BackgroundColor: style.ThemeValue(color.From(0x00000000)),
					BorderColor:     style.ThemeValue(color.From(0x00000000)),
					BorderRadius:    style.ThemeValue(8),
					Fill:            style.ThemeValue(color.From(0xe0e4e4ff)),
				},
			},
		},
		TextTheme: td.TextTheme{
			CodeBlockStyle: style.TextStyleCollection{
				Primary: style.TextStyle{
					FontSize:   style.ThemeValue(14),
					FontWeight: style.ThemeValue("400"),
					Color:      style.ThemeValue(color.From(0xe0e4e4ff)),
					FontFamily: style.ThemeValue("ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, \"Liberation Mono\", \"Courier New\", monospace"),
				},
			},
			TextStyle: style.TextStyleCollection{
				Primary: style.TextStyle{
					FontSize:   style.ThemeValue(16),
					FontWeight: style.ThemeValue("400"),
					Color:      style.ThemeValue(color.From(0xe0e4e4ff)),
					FontFamily: style.ThemeValue("Ubuntu"),
				},
				Secondary: style.TextStyle{
					FontSize:   style.ThemeValue(16),
					FontWeight: style.ThemeValue("400"),
					Color:      style.ThemeValue(color.From(0x9e9e9eff)),
					FontFamily: style.ThemeValue("Ubuntu"),
				},
				Tertiary: style.TextStyle{
					FontSize:   style.ThemeValue(16),
					FontWeight: style.ThemeValue("400"),
					Color:      style.ThemeValue(color.From(0xecececff)),
					FontFamily: style.ThemeValue("Ubuntu"),
				},
			},
		},
	}
)

func DarkTheme() *theme_data.ThemeData {
	return darkTheme
}
