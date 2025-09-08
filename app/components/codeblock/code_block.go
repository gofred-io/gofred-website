package codeblock

import (
	"github.com/gofred-io/gofred-website/app/components/snackbar"
	"github.com/gofred-io/gofred-website/app/constant"
	codeblock "github.com/gofred-io/gofred/foundation/code_block"
	"github.com/gofred-io/gofred/widget"
)

func New(code string) widget.BaseWidget {
	return codeblock.New(code, codeblock.OnCopied(func(code string) {
		snackbar.Show("Successfully copied to clipboard", constant.SnackbarTypeSuccess)
	}))
}
