package snackbar

import (
	"github.com/gofred-io/gofred-website/app/constant"
	"github.com/gofred-io/gofred/application"
	"github.com/gofred-io/gofred/breakpoint"
	"github.com/gofred-io/gofred/foundation/container"
	"github.com/gofred-io/gofred/foundation/icon"
	icondata "github.com/gofred-io/gofred/foundation/icon/icon_data"
	"github.com/gofred-io/gofred/foundation/row"
	foundationSnackbar "github.com/gofred-io/gofred/foundation/snackbar"
	"github.com/gofred-io/gofred/foundation/text"
	"github.com/gofred-io/gofred/options/spacing"
	"github.com/gofred-io/gofred/theme"
)

var (
	snackbar foundationSnackbar.ISnackbar
)

func Get() foundationSnackbar.ISnackbar {
	if snackbar == nil {
		snackbar = buildSnackbar()
	}
	return snackbar
}

func buildSnackbar() foundationSnackbar.ISnackbar {
	return foundationSnackbar.New(
		foundationSnackbar.Transition(0.3),
		foundationSnackbar.Position(theme.PositionTypeBottomLeft),
	)
}

// Show creates and displays a snackbar with the given message and type
func Show(message string, snackbarType constant.SnackbarType) {
	Get().Show(snackbarContainer(message, snackbarType))
}

func snackbarContainer(message string, snackbarType constant.SnackbarType) application.BaseWidget {
	return container.New(
		row.New(
			[]application.BaseWidget{
				snackbarIcon(snackbarType),
				snackbarContent(message),
			},
			row.Gap(12),
			row.CrossAxisAlignment(theme.AxisAlignmentTypeCenter),
		),
		container.BackgroundColor(getBackgroundColor(snackbarType)),
		container.BorderRadius(8),
		container.Padding(breakpoint.All(spacing.All(16))),
		container.Margin(breakpoint.All(spacing.All(16))),
		container.BoxShadow("0 4px 12px rgba(0, 0, 0, 0.15)"),
	)
}

func snackbarContent(message string) application.BaseWidget {
	return container.New(
		text.New(
			message,
			text.FontSize(14),
			text.FontColor("#FFFFFF"),
			text.FontWeight("500"),
		),
		container.BackgroundColor("#00000000"),
		container.Flex(1),
	)
}

func snackbarIcon(snackbarType constant.SnackbarType) application.BaseWidget {
	iconColor := "#FFFFFF"
	iconData := getIconData(snackbarType)

	return icon.New(
		iconData,
		icon.Width(breakpoint.All(20)),
		icon.Height(breakpoint.All(20)),
		icon.Fill(iconColor),
	)
}

func getBackgroundColor(snackbarType constant.SnackbarType) string {
	switch snackbarType {
	case constant.SnackbarTypeSuccess:
		return "#10B981" // Green
	case constant.SnackbarTypeError:
		return "#EF4444" // Red
	case constant.SnackbarTypeWarning:
		return "#F59E0B" // Amber
	case constant.SnackbarTypeInfo:
		return "#3B82F6" // Blue
	default:
		return "#424242" // Gray
	}
}

func getIconData(snackbarType constant.SnackbarType) icondata.IconData {
	switch snackbarType {
	case constant.SnackbarTypeSuccess:
		return icondata.Check
	case constant.SnackbarTypeError:
		return icondata.Close
	case constant.SnackbarTypeWarning:
		return icondata.Alert
	case constant.SnackbarTypeInfo:
		return icondata.Information
	default:
		return icondata.Information
	}
}
