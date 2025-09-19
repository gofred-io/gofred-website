package codeblock

import (
	"github.com/gofred-io/gofred-website/app/components/snackbar"
	"github.com/gofred-io/gofred-website/app/constant"
	"github.com/gofred-io/gofred-website/app/theme"
	"github.com/gofred-io/gofred/application"
	codeblock "github.com/gofred-io/gofred/foundation/code_block"
)

func New(code string) application.BaseWidget {
	return codeblock.New(
		code,
		codeblock.OnCopied(func(code string) {
			snackbar.Show("Successfully copied to clipboard", constant.SnackbarTypeSuccess)
		}),
		codeblock.TextStyle(theme.Data().TextTheme.CodeBlockStyle.Primary),
	)
}
