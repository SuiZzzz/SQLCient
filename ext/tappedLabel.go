package ext

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TappableLabel struct {
	widget.Label
	Message  string
	OnTapped func()
}

func NewTappableLabel(text string, tap func()) *TappableLabel {
	label := &TappableLabel{}
	label.ExtendBaseWidget(label)
	label.SetText(text)
	label.OnTapped = tap
	return label
}
func (t *TappableLabel) Tapped(*fyne.PointEvent) {
	t.OnTapped()
}

func (t *TappableLabel) TappedSecondary(*fyne.PointEvent) {
}
