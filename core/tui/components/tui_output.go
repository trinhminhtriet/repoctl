package components

import (
	"github.com/rivo/tview"
	"github.com/trinhminhtriet/repoctl/core/tui/misc"
)

func CreateOutputView(title string) (*tview.TextView, *misc.ThreadSafeWriter) {
	streamView := CreateText(title)
	ansiWriter := misc.NewThreadSafeWriter(streamView)

	return streamView, ansiWriter
}
