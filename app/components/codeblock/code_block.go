package codeblock

import (
	"github.com/gofred-io/gofred-website/app/components/snackbar"
	"github.com/gofred-io/gofred-website/app/constant"
	"github.com/gofred-io/gofred/application"
	codeblock "github.com/gofred-io/gofred/foundation/code_block"
	"github.com/gofred-io/gofred/hooks"
)

func New(code string) application.BaseWidget {
	themeHook, _ := hooks.UseTheme()

	return codeblock.New(
		code,
		codeblock.OnCopied(func(code string) {
			snackbar.Show("Successfully copied to clipboard", constant.SnackbarTypeSuccess)
		}),
		codeblock.TextStyle(themeHook.ThemeData().TextTheme.CodeBlockStyle.Primary),
	)
}
