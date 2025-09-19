package theme

import (
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
	"github.com/gofred-io/gofred/theme/color"
	td "github.com/gofred-io/gofred/theme/theme_data"
	style "github.com/gofred-io/gofred/theme/theme_style"
)

var (
	lightTheme = &td.ThemeData{
		Name: string(ThemeLight),
		BoxTheme: td.BoxTheme{
			CodeBlockStyle: style.ContainerStyleCollection{
				Primary: style.ContainerStyle{
					BackgroundColor: style.ThemeValue(color.From(0x1F2937ff)),
				},
			},
			ContainerStyle: style.ContainerStyleCollection{
				Primary: style.ContainerStyle{
					BackgroundColor: style.ThemeValue(color.From(0xfdfdfdff)),
					BorderColor:     style.ThemeValue(color.From(0xecececff)),
				},
				Secondary: style.ContainerStyle{
					BackgroundColor: style.ThemeValue(color.From(0xf8f8f8ff)),
					BorderColor:     style.ThemeValue(color.From(0xE5E7EBFF)),
				},
				Tertiary: style.ContainerStyle{
					BackgroundColor: style.ThemeValue(color.From(0x1F2937ff)),
					BorderColor:     style.ThemeValue(color.From(0x374151ff)),
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
					BackgroundColor: style.ThemeValue(color.From(0xFFFFFFFF)),
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
					Fill:            style.ThemeValue(color.From(0x4c5e78ff)),
				},
				Secondary: style.ButtonStyle{
					BackgroundColor: style.ThemeValue(color.From(0x00000000)),
					BorderColor:     style.ThemeValue(color.From(0x00000000)),
					BorderRadius:    style.ThemeValue(8),
					Fill:            style.ThemeValue(color.From(0x4c5e78ff)),
				},
			},
		},
		TextTheme: td.TextTheme{
			CodeBlockStyle: style.TextStyleCollection{
				Primary: style.TextStyle{
					FontSize:   style.ThemeValue(14),
					FontWeight: style.ThemeValue("400"),
					Color:      style.ThemeValue(color.From(0xF3F4F6ff)),
					FontFamily: style.ThemeValue("ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, \"Liberation Mono\", \"Courier New\", monospace"),
				},
			},
			TextStyle: style.TextStyleCollection{
				Primary: style.TextStyle{
					FontSize:   style.ThemeValue(16),
					FontWeight: style.ThemeValue("400"),
					Color:      style.ThemeValue(color.From(0x1F2937ff)),
					FontFamily: style.ThemeValue("Ubuntu"),
				},
				Secondary: style.TextStyle{
					FontSize:   style.ThemeValue(16),
					FontWeight: style.ThemeValue("400"),
					Color:      style.ThemeValue(color.From(0x6a6b6cff)),
					FontFamily: style.ThemeValue("Ubuntu"),
				},
				Tertiary: style.TextStyle{
					FontSize:   style.ThemeValue(16),
					FontWeight: style.ThemeValue("400"),
					Color:      style.ThemeValue(color.From(0xfafafaff)),
					FontFamily: style.ThemeValue("Ubuntu"),
				},
			},
		},
	}
)

func LightTheme() *td.ThemeData {
	return lightTheme
}
