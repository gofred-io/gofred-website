package notfound

import (
	"github.com/gofred-io/gofred/text"
	"github.com/gofred-io/gofred/widget"
)

func New() widget.BaseWidget {
	return text.New("Page not found")
}
