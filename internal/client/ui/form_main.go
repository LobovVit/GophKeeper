package ui

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/LobovVit/GophKeeper/internal/client/constants"
)

func (u App) getMainForm(ctx context.Context) fyne.CanvasObject {
	labelAlert := widget.NewLabel("")
	labelAlert.Hide()
	var err error
	content := container.NewStack()
	title := widget.NewLabel("")
	var buttonTopSynchronization *widget.Button

	buttonTopSynchronization = widget.NewButton(constants.BtnUpdateData, func() {
		cnf := dialog.NewConfirm(constants.BtnUpdateData, constants.BtnUpdateData,
			func(b bool) {
				if !b {
					return
				}
				u.data["dataTblText"], u.data["dataTblCard"], u.data["dataTblLoginPassword"], u.data["dataTblBinary"], err = u.client.Synchronization(ctx, u.password, u.accessToken)
				if err != nil {
					dialog.ShowError(err, u.window)
					return
				}
			}, u.window)
		cnf.SetDismissText(constants.Cancel)
		cnf.SetConfirmText(constants.Save)
		cnf.Show()
	})
	setSidebar := func(s sidebar) {
		title.SetText(s.title)
		content.Objects = []fyne.CanvasObject{s.view(ctx, u.window)}
		content.Refresh()
	}
	item := container.NewBorder(container.NewVBox(buttonTopSynchronization, widget.NewSeparator()), nil, nil, nil, content)
	side := container.NewBorder(nil, nil, nil, nil, u.makeSidebar(ctx, setSidebar))
	formMainContainer := container.NewHSplit(side, item)
	formMainContainer.Offset = 0.1
	return formMainContainer
}
