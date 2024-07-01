package ui

import (
	"context"
	"errors"
	"io"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/LobovVit/GophKeeper/internal/client/constants"
)

// ----BtnLoginPassword----
func (u App) getBtnLoginPasswordDelete(ctx context.Context) fyne.CanvasObject {
	var buttonLoginPasswordDelete *widget.Button

	buttonLoginPasswordDelete = widget.NewButton(constants.BtnDeleteLoginPassword, func() {
		if u.currentIndex["dataTblLoginPassword"] == 0 {
			dialog.ShowInformation(constants.Err, constants.ErrLoginPasswordTblIndexDelete, u.window)
			return
		}
		cnf := dialog.NewConfirm(constants.BtnDeleteLoginPassword, u.currentRow["dataTblLoginPassword"][0],
			func(b bool) {
				if !b {
					return
				}
				err := u.client.LoginPasswordDelete(ctx, u.currentRow["dataTblLoginPassword"], u.accessToken)
				if err != nil {
					dialog.ShowError(err, u.window)
					return
				}
				u.data["dataTblLoginPassword"] = removeRow(u.data["dataTblLoginPassword"], u.currentIndex["dataTblLoginPassword"])
				u.window.Content().Refresh()
			}, u.window)
		cnf.SetDismissText(constants.Cancel)
		cnf.SetConfirmText(constants.Save)
		cnf.Show()
	})
	return buttonLoginPasswordDelete
}

func (u App) getBtnLoginPasswordCreate(ctx context.Context) fyne.CanvasObject {
	var buttonLoginPasswordCreate *widget.Button
	buttonLoginPasswordCreate = widget.NewButton(constants.BtnAddLoginPassword, func() {
		loginPasswordName := widget.NewEntry()
		loginPasswordDescription := widget.NewEntry()
		login := widget.NewEntry()
		password := widget.NewEntry()
		items := []*widget.FormItem{
			widget.NewFormItem(constants.NameItem, loginPasswordName),
			widget.NewFormItem(constants.DescriptionItem, loginPasswordDescription),
			widget.NewFormItem(constants.LoginItem, login),
			widget.NewFormItem(constants.PasswordItem, password),
		}
		dialogForm := dialog.NewForm(constants.BtnAddLoginPassword, constants.Save, constants.Cancel, items, func(b bool) {
			if !b {
				return
			}
			if searchByColumn(u.data["dataTblLoginPassword"], 0, loginPasswordName.Text) {
				dialog.ShowError(errors.New(constants.ErrLoginPasswordExist), u.window)
				return
			}
			errMsg, valid := validateLoginPasswordForm(loginPasswordName, loginPasswordDescription, login, password)
			if valid {
				err := u.client.LoginPasswordCreate(ctx, loginPasswordName.Text, loginPasswordDescription.Text, u.password, login.Text, password.Text, u.accessToken)
				if err != nil {
					dialog.ShowError(err, u.window)
					return
				}
				u.data["dataTblLoginPassword"] = append(u.data["dataTblLoginPassword"], []string{loginPasswordName.Text, loginPasswordDescription.Text, login.Text, password.Text, time.Now().Format(constants.LayoutDateAndTime.ToString()), time.Now().Format(constants.LayoutDateAndTime.ToString())})
			} else {
				dialog.ShowError(errors.New(errMsg), u.window)
				return
			}
			u.window.Content().Refresh()
		}, u.window)
		dialogForm.Resize(fyne.NewSize(constants.WindowDialogWidth, constants.WindowDialogHeight))
		dialogForm.Show()
	})
	return buttonLoginPasswordCreate
}

func (u App) getBtnLoginPasswordUpdate(ctx context.Context) fyne.CanvasObject {
	var buttonLoginPasswordUpdate *widget.Button
	buttonLoginPasswordUpdate = widget.NewButton(constants.BtnUpdateLoginPassword, func() {
		if u.currentIndex["dataTblLoginPassword"] == 0 {
			dialog.ShowInformation(constants.Err, constants.ErrLoginPasswordTblIndexUpdate, u.window)
			return
		}
		loginPasswordName := widget.NewEntry()
		loginPasswordName.Text = u.currentRow["dataTblLoginPassword"][0]
		loginPasswordDescription := widget.NewEntry()
		loginPasswordDescription.Text = u.currentRow["dataTblLoginPassword"][1]
		login := widget.NewEntry()
		login.Text = u.currentRow["dataTblLoginPassword"][2]
		password := widget.NewEntry()
		password.Text = u.currentRow["dataTblLoginPassword"][3]
		items := []*widget.FormItem{
			widget.NewFormItem(constants.LoginItem, login),
			widget.NewFormItem(constants.PasswordItem, password),
		}
		dialogForm := dialog.NewForm(constants.BtnUpdateLoginPassword, constants.Save, constants.Cancel, items, func(b bool) {
			if !b {
				return
			}
			errMsg, valid := validateLoginPasswordForm(loginPasswordName, loginPasswordDescription, login, password)
			if valid {
				err := u.client.LoginPasswordUpdate(ctx, loginPasswordName.Text, u.password, login.Text, password.Text, u.accessToken)
				if err != nil {
					dialog.ShowError(err, u.window)
					return
				}
				u.data["dataTblLoginPassword"][u.currentIndex["dataTblLoginPassword"]] = []string{loginPasswordName.Text, loginPasswordDescription.Text, login.Text, password.Text, u.currentRow["dataTblLoginPassword"][4], time.Now().Format(constants.LayoutDateAndTime.ToString())}
			} else {
				dialog.ShowError(errors.New(errMsg), u.window)
				return
			}
			u.window.Content().Refresh()
		}, u.window)
		dialogForm.Resize(fyne.NewSize(constants.WindowDialogWidth, constants.WindowDialogHeight))
		dialogForm.Show()
	})
	return buttonLoginPasswordUpdate
}

func (u App) getBtnLoginPassword(ctx context.Context) fyne.CanvasObject {
	btnLoginPassword := container.NewGridWithColumns(3, u.getBtnLoginPasswordCreate(ctx), u.getBtnLoginPasswordUpdate(ctx), u.getBtnLoginPasswordDelete(ctx))
	return btnLoginPassword
}

// ---------BtnCard----
func (u App) getBtnCardDelete(ctx context.Context) fyne.CanvasObject {
	var buttonCardDelete *widget.Button

	buttonCardDelete = widget.NewButton(constants.BtnDeleteCard, func() {
		if u.currentIndex["dataTblCard"] == 0 {
			dialog.ShowInformation(constants.Err, constants.ErrCardTblIndexDelete, u.window)
			return
		}
		cnf := dialog.NewConfirm(constants.BtnDeleteCard, u.currentRow["dataTblCard"][0],
			func(b bool) {
				if !b {
					return
				}
				err := u.client.CardDelete(ctx, u.currentRow["dataTblCard"], u.accessToken)
				if err != nil {
					dialog.ShowError(err, u.window)
					return
				}
				u.data["dataTblCard"] = removeRow(u.data["dataTblCard"], u.currentIndex["dataTblCard"])
				u.window.Content().Refresh()
			}, u.window)
		cnf.SetDismissText(constants.Cancel)
		cnf.SetConfirmText(constants.Save)
		cnf.Show()
	})
	return buttonCardDelete
}

func (u App) getBtnCardCreate(ctx context.Context) fyne.CanvasObject {
	var buttonCardCreate *widget.Button
	buttonCardCreate = widget.NewButton(constants.BtnAddCard, func() {
		cardName := widget.NewEntry()
		cardDescription := widget.NewEntry()
		paymentSystem := widget.NewEntry()
		number := widget.NewEntry()
		holder := widget.NewEntry()
		endDate := widget.NewEntry()
		endDate.Validator = validation.NewTime(constants.LayoutDate.ToString())
		cvc := widget.NewEntry()
		items := []*widget.FormItem{
			widget.NewFormItem(constants.NameItem, cardName),
			widget.NewFormItem(constants.DescriptionItem, cardDescription),
			widget.NewFormItem(constants.PaymentSystemItem, paymentSystem),
			widget.NewFormItem(constants.NumberItem, number),
			widget.NewFormItem(constants.HolderItem, holder),
			widget.NewFormItem(constants.EndDateItem, endDate),
			widget.NewFormItem(constants.CVCItem, cvc),
		}
		dialogForm := dialog.NewForm(constants.BtnAddCard, constants.Save, constants.Cancel, items, func(b bool) {
			if !b {
				return
			}
			if searchByColumn(u.data["dataTblCard"], 0, cardName.Text) {
				dialog.ShowError(errors.New(constants.ErrCardExist), u.window)
				return
			}
			errMsg, valid := validateCardForm(cardName, cardDescription, paymentSystem, number, holder, cvc, endDate)
			if valid {
				err := u.client.CardCreate(ctx, cardName.Text, cardDescription.Text, u.password, paymentSystem.Text, number.Text, holder.Text, cvc.Text, endDate.Text, u.accessToken)
				if err != nil {
					dialog.ShowError(err, u.window)
					return
				}
				u.data["dataTblCard"] = append(u.data["dataTblCard"], []string{cardName.Text, cardDescription.Text, paymentSystem.Text, number.Text, holder.Text, cvc.Text, endDate.Text, time.Now().Format(constants.LayoutDateAndTime.ToString()), time.Now().Format(constants.LayoutDateAndTime.ToString())})
			} else {
				dialog.ShowError(errors.New(errMsg), u.window)
				return
			}
			u.window.Content().Refresh()
		}, u.window)
		dialogForm.Resize(fyne.NewSize(constants.WindowDialogWidth, constants.WindowDialogHeight))
		dialogForm.Show()
	})
	return buttonCardCreate
}

func (u App) getBtnCardUpdate(ctx context.Context) fyne.CanvasObject {
	var buttonCardUpdate *widget.Button
	buttonCardUpdate = widget.NewButton(constants.BtnUpdateCard, func() {
		if u.currentIndex["dataTblCard"] == 0 {
			dialog.ShowInformation(constants.Err, constants.ErrCardTblIndexUpdate, u.window)
			return
		}
		cardName := widget.NewEntry()
		cardName.Text = u.currentRow["dataTblCard"][0]
		cardDescription := widget.NewEntry()
		cardDescription.Text = u.currentRow["dataTblCard"][1]
		paymentSystem := widget.NewEntry()
		paymentSystem.Text = u.currentRow["dataTblCard"][2]
		number := widget.NewEntry()
		number.Text = u.currentRow["dataTblCard"][3]
		holder := widget.NewEntry()
		holder.Text = u.currentRow["dataTblCard"][4]
		cvc := widget.NewEntry()
		cvc.Text = u.currentRow["dataTblCard"][5]
		endDate := widget.NewEntry()
		endDate.Validator = validation.NewTime(constants.LayoutDate.ToString())
		endDate.Text = u.currentRow["dataTblCard"][6]

		items := []*widget.FormItem{
			widget.NewFormItem(constants.PaymentSystemItem, paymentSystem),
			widget.NewFormItem(constants.NumberItem, number),
			widget.NewFormItem(constants.HolderItem, holder),
			widget.NewFormItem(constants.EndDateItem, endDate),
			widget.NewFormItem(constants.CVCItem, cvc),
		}
		dialogForm := dialog.NewForm(constants.BtnUpdateCard, constants.Save, constants.Cancel, items, func(b bool) {
			if !b {
				return
			}
			errMsg, valid := validateCardForm(cardName, cardDescription, paymentSystem, number, holder, cvc, endDate)
			if valid {
				err := u.client.CardUpdate(ctx, cardName.Text, u.password, paymentSystem.Text, number.Text, holder.Text, cvc.Text, endDate.Text, u.accessToken)
				if err != nil {
					dialog.ShowError(err, u.window)
					return
				}
				u.data["dataTblCard"][u.currentIndex["dataTblCard"]] = []string{cardName.Text, cardDescription.Text, paymentSystem.Text, number.Text, holder.Text, cvc.Text, endDate.Text, u.currentRow["dataTblCard"][7], time.Now().Format(constants.LayoutDateAndTime.ToString())}
			} else {
				dialog.ShowError(errors.New(errMsg), u.window)
				return
			}
			u.window.Content().Refresh()
		}, u.window)
		dialogForm.Resize(fyne.NewSize(constants.WindowDialogWidth, constants.WindowDialogHeight))
		dialogForm.Show()
	})
	return buttonCardUpdate
}

func (u App) getBtnCard(ctx context.Context) fyne.CanvasObject {
	btnCard := container.NewGridWithColumns(3, u.getBtnCardCreate(ctx), u.getBtnCardUpdate(ctx), u.getBtnCardDelete(ctx))
	return btnCard
}

// ---------BtnText----
func (u App) getBtnTextDelete(ctx context.Context) fyne.CanvasObject {
	var buttonTextDelete *widget.Button

	buttonTextDelete = widget.NewButton(constants.BtnDeleteText, func() {
		if u.currentIndex["dataTblText"] == 0 {
			dialog.ShowInformation(constants.Err, constants.ErrTextTblIndexDelete, u.window)
			return
		}
		cnf := dialog.NewConfirm(constants.BtnDeleteText, u.currentRow["dataTblText"][0],
			func(b bool) {
				if !b {
					return
				}
				err := u.client.TextDelete(ctx, u.currentRow["dataTblText"], u.accessToken)
				if err != nil {
					dialog.ShowError(err, u.window)
					return
				}
				u.data["dataTblText"] = removeRow(u.data["dataTblText"], u.currentIndex["dataTblText"])
				u.window.Content().Refresh()
			}, u.window)
		cnf.SetDismissText(constants.Cancel)
		cnf.SetConfirmText(constants.Save)
		cnf.Show()
	})
	return buttonTextDelete
}

func (u App) getBtnTextCreate(ctx context.Context) fyne.CanvasObject {
	var buttonTextCreate *widget.Button
	buttonTextCreate = widget.NewButton(constants.BtnAddText, func() {
		textName := widget.NewEntry()
		textDescription := widget.NewEntry()
		text := widget.NewEntry()
		items := []*widget.FormItem{
			widget.NewFormItem(constants.NameItem, textName),
			widget.NewFormItem(constants.DescriptionItem, textDescription),
			widget.NewFormItem(constants.DataItem, text),
		}
		dialogForm := dialog.NewForm(constants.BtnAddText, constants.Save, constants.Cancel, items, func(b bool) {
			if !b {
				return
			}
			if searchByColumn(u.data["dataTblText"], 0, textName.Text) {
				dialog.ShowError(errors.New(constants.ErrTextExist), u.window)
				return
			}
			errMsg, valid := validateTextForm(textName, textDescription, text)
			if valid {
				err := u.client.TextCreate(ctx, textName.Text, textDescription.Text, u.password, text.Text, u.accessToken)
				if err != nil {
					dialog.ShowError(err, u.window)
					return
				}
				u.data["dataTblText"] = append(u.data["dataTblText"], []string{textName.Text, textDescription.Text, text.Text, time.Now().Format(constants.LayoutDateAndTime.ToString()), time.Now().Format(constants.LayoutDateAndTime.ToString())})
			} else {
				dialog.ShowError(errors.New(errMsg), u.window)
				return
			}
			u.window.Content().Refresh()
		}, u.window)
		dialogForm.Resize(fyne.NewSize(constants.WindowDialogWidth, constants.WindowDialogHeight))
		dialogForm.Show()

	})
	return buttonTextCreate
}

func (u App) getBtnTextUpdate(ctx context.Context) fyne.CanvasObject {
	var buttonTextUpdate *widget.Button
	buttonTextUpdate = widget.NewButton(constants.BtnUpdateText, func() {
		if u.currentIndex["dataTblText"] == 0 {
			dialog.ShowInformation(constants.Err, constants.ErrTextTblIndexUpdate, u.window)
			return
		}
		textName := widget.NewEntry()
		textName.Text = u.currentRow["dataTblText"][0]
		textDescription := widget.NewEntry()
		textDescription.Text = u.currentRow["dataTblText"][1]
		text := widget.NewEntry()
		text.Text = u.currentRow["dataTblText"][2]

		items := []*widget.FormItem{
			widget.NewFormItem(constants.DataItem, text),
		}
		dialogForm := dialog.NewForm(constants.BtnUpdateText, constants.Save, constants.Cancel, items, func(b bool) {
			if !b {
				return
			}
			errMsg, valid := validateTextForm(textName, textDescription, text)
			if valid {
				err := u.client.TextUpdate(ctx, textName.Text, u.password, text.Text, u.accessToken)
				if err != nil {
					dialog.ShowError(err, u.window)
					return
				}
				u.data["dataTblText"][u.currentIndex["dataTblText"]] = []string{textName.Text, textDescription.Text, text.Text, u.currentRow["dataTblText"][3], time.Now().Format(constants.LayoutDateAndTime.ToString())}
			} else {
				dialog.ShowError(errors.New(errMsg), u.window)
				return
			}
			u.window.Content().Refresh()
		}, u.window)
		dialogForm.Resize(fyne.NewSize(constants.WindowDialogWidth, constants.WindowDialogHeight))
		dialogForm.Show()

	})
	return buttonTextUpdate
}

func (u App) getBtnText(ctx context.Context) fyne.CanvasObject {
	btnText := container.NewGridWithColumns(3, u.getBtnTextCreate(ctx), u.getBtnTextUpdate(ctx), u.getBtnTextDelete(ctx))
	return btnText
}

// -----Binary----
func (u App) getBtnBinaryDelete(ctx context.Context) fyne.CanvasObject {
	var buttonBinaryDelete *widget.Button

	buttonBinaryDelete = widget.NewButton(constants.BtnDeleteBinary, func() {
		if u.currentIndex["dataTblBinary"] == 0 {
			dialog.ShowInformation(constants.Err, constants.ErrBinaryTblIndexDelete, u.window)
			return
		}
		cnf := dialog.NewConfirm(constants.BtnDeleteBinary, u.currentRow["dataTblBinary"][0],
			func(b bool) {
				if !b {
					return
				}
				err := u.client.FileRemove(ctx, u.currentRow["dataTblBinary"], u.accessToken)
				if err != nil {
					dialog.ShowError(err, u.window)
					return
				}
				u.data["dataTblBinary"] = removeRow(u.data["dataTblBinary"], u.currentIndex["dataTblBinary"])
				u.window.Content().Refresh()
			}, u.window)
		cnf.SetDismissText(constants.Cancel)
		cnf.SetConfirmText(constants.Save)
		cnf.Show()
	})
	return buttonBinaryDelete
}

func (u App) getBtnBinaryCreate(ctx context.Context) fyne.CanvasObject {
	var buttonBinaryCreate *widget.Button
	buttonBinaryCreate = widget.NewButton(constants.BtnUploadBinary, func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, u.window)
				return
			}
			if reader == nil {
				return
			}
			if searchByColumn(u.data["dataTblBinary"], 0, reader.URI().Name()) {
				dialog.ShowError(errors.New(constants.ErrBinaryExist), u.window)
				return
			}
			data, err := io.ReadAll(reader)
			if err != nil {
				dialog.ShowError(err, u.window)
				return
			}
			if len(data) > u.client.Cfg.FileSize {
				dialog.ShowError(errors.New(constants.ErrFileSize), u.window)
				return
			}
			name, err := u.client.FileUpload(ctx, reader.URI().Name(), u.password, data, u.accessToken)
			if err != nil {
				dialog.ShowError(err, u.window)
				return
			}
			u.data["dataTblBinary"] = append(u.data["dataTblBinary"], []string{name, time.Now().Format(constants.LayoutDateAndTime.ToString())})
		}, u.window)
		fd.Show()
	})
	return buttonBinaryCreate
}

func (u App) getBtnBinaryGet(ctx context.Context) fyne.CanvasObject {
	var buttonBinaryGet *widget.Button
	buttonBinaryGet = widget.NewButton(constants.BtnDownloadBinary, func() {
		if u.currentIndex["dataTblBinary"] == 0 {
			dialog.ShowInformation(constants.Err, constants.ErrBinaryTblIndexDownload, u.window)
			return
		}
		dialogForm := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, u.window)
				return
			}
			if writer == nil {
				log.Println("Cancelled")
				return
			}
			f := writer
			defer f.Close()
			file, err := u.client.FileDownload(ctx, u.currentRow["dataTblBinary"][0], u.password, u.accessToken)
			if err != nil {
				dialog.ShowError(err, u.window)
			}
			_, err = f.Write(file)
			if err != nil {
				dialog.ShowError(err, u.window)
			}
			err = f.Close()
			if err != nil {
				dialog.ShowError(err, u.window)
			}
		}, u.window)
		dialogForm.Resize(fyne.NewSize(constants.WindowDialogWidth, constants.WindowDialogHeight))
		dialogForm.Show()

	})
	return buttonBinaryGet
}

func (u App) getBtnBinary(ctx context.Context) fyne.CanvasObject {
	btnText := container.NewGridWithColumns(3, u.getBtnBinaryCreate(ctx), u.getBtnBinaryGet(ctx), u.getBtnBinaryDelete(ctx))
	return btnText
}
