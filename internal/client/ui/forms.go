package ui

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/LobovVit/GophKeeper/internal/client/client"
	"github.com/LobovVit/GophKeeper/internal/client/constants"
	"github.com/LobovVit/GophKeeper/pkg/logger"
	"go.uber.org/zap"
)

func (u App) getWelcomeForm(ctx context.Context) fyne.CanvasObject {

	labelAlert := widget.NewLabel("")
	labelAlert.Hide()
	var buttonLogin *widget.Button
	var buttonReg *widget.Button
	var buttonBack *widget.Button
	var buttonTestConnection *widget.Button
	var buttonChangeServer *widget.Button
	usernameLoginEntry := widget.NewEntry()
	passwordLoginEntry := widget.NewPasswordEntry()
	usernameRegistrationEntry := widget.NewEntry()
	passwordRegistrationEntry := widget.NewPasswordEntry()
	passwordConfirmationRegistrationEntry := widget.NewPasswordEntry()
	separator := widget.NewSeparator()
	var containerFormLogin *fyne.Container
	var containerFormRegistration *fyne.Container
	var exist bool
	var err error

	buttonBack = widget.NewButton(constants.BtnBack, func() {
		u.window.SetContent(u.getWelcomeForm(ctx))
		u.window.Show()
	})

	buttonLogin = widget.NewButton(constants.BtnSubmit, func() {
		labelAlert.Show()
		errMsg, valid := validateLoginForm(usernameLoginEntry, passwordLoginEntry)
		if valid {
			u.accessToken, err = u.client.Authentication(ctx, usernameLoginEntry.Text, passwordLoginEntry.Text)
			if err != nil {
				labelAlert.SetText(constants.ErrLogin)
				logger.Log.Error("authentication", zap.Error(err))
			} else {
				u.password = passwordLoginEntry.Text
				u.data["dataTblText"], u.data["dataTblCard"], u.data["dataTblLoginPassword"], u.data["dataTblBinary"], err = u.client.Synchronization(ctx, u.password, u.accessToken)
				if err != nil {
					labelAlert.SetText(err.Error())
					logger.Log.Error("synchronization", zap.Error(err))
					return
				}
				u.window.Resize(fyne.NewSize(constants.WindowMainWidth, constants.WindowMainHeight))
				u.window.SetContent(u.getMainForm(ctx))
				u.window.Show()

			}
		} else {
			labelAlert.SetText(errMsg)
			logger.Log.Error(errMsg)
		}

	})

	buttonReg = widget.NewButton(constants.BtnSubmit, func() {
		labelAlert.Show()
		errMsg, valid := validateRegistrationForm(usernameRegistrationEntry, passwordRegistrationEntry, passwordConfirmationRegistrationEntry)
		if valid {
			exist, err = u.client.UserExist(ctx, usernameRegistrationEntry.Text)
			if err != nil {
				labelAlert.SetText(constants.ErrRegistration)
				logger.Log.Error("user exist", zap.Error(err))
			}
			if exist {
				labelAlert.SetText(constants.ErrUserExist)
				logger.Log.Error(constants.ErrUserExist)
			} else {
				u.accessToken, err = u.client.Registration(ctx, usernameRegistrationEntry.Text, passwordRegistrationEntry.Text)
				if err != nil {
					labelAlert.SetText(constants.ErrRegistration)
					logger.Log.Error("registration", zap.Error(err))
				} else {
					u.password = passwordRegistrationEntry.Text
					u.data["dataTblText"], u.data["dataTblCard"], u.data["dataTblLoginPassword"], u.data["dataTblBinary"], err = u.client.Synchronization(ctx, u.password, u.accessToken)
					if err != nil {
						labelAlert.SetText(err.Error())
						logger.Log.Error("synchronization", zap.Error(err))
						return
					}
					u.window.Resize(fyne.NewSize(constants.WindowMainWidth, constants.WindowMainHeight))
					u.window.SetContent(u.getMainForm(ctx))
					u.window.Show()
				}
			}
		} else {
			labelAlert.SetText(errMsg)
			logger.Log.Error(errMsg)
		}
	})

	buttonTestConnection = widget.NewButton(constants.BtnTestConnection, func() {
		labelAlert.Show()
		msg, err := u.client.Ping(ctx)
		if err != nil {
			labelAlert.SetText(constants.ErrConnect + " (" + u.client.Cfg.HostGRPC + ")")
			logger.Log.Info(msg)
			logger.Log.Error(constants.ErrConnect, zap.Error(err))
		} else {
			logger.Log.Info(msg)
			labelAlert.SetText(constants.SuccessConnect + " (" + u.client.Cfg.HostGRPC + ")")
		}
	})

	buttonChangeServer = widget.NewButton(constants.BtnSaveConfig, func() {
		hostGRPCEntry := widget.NewEntry()
		hostGRPCEntry.SetText(u.client.Cfg.HostGRPC)
		items := []*widget.FormItem{
			widget.NewFormItem(constants.HostGRPC, hostGRPCEntry)}
		dialog.ShowForm(constants.BtnSaveConfig, constants.Save, constants.Cancel, items, func(b bool) {
			if !b {
				return
			}
			u.client.Cfg.HostGRPC = hostGRPCEntry.Text
			u.client, err = client.New(u.client.Cfg)
		}, u.window)
	})

	formConfigButtons := container.NewVBox(
		buttonTestConnection,
		buttonChangeServer)

	formLogin := u.getFormLogin(usernameLoginEntry, passwordLoginEntry)
	formRegistration := u.getFormRegistration(usernameRegistrationEntry, passwordRegistrationEntry, passwordConfirmationRegistrationEntry)
	containerFormLogin = container.NewVBox(formLogin, buttonLogin, labelAlert, separator, buttonBack)
	containerFormRegistration = container.NewVBox(formRegistration, buttonReg, labelAlert, separator, buttonBack)
	formAuthButtons := container.NewGridWithColumns(2,
		widget.NewButton(constants.BtnLogin, func() {
			msg, err := u.client.Ping(ctx)
			if err != nil {
				dialog.ShowInformation(msg, constants.ErrConnect+" ("+u.client.Cfg.HostGRPC+")", u.window)
				return
			}
			logger.Log.Info("Sign in")
			labelAlert.SetText("")
			labelAlert.Hide()
			u.window.SetContent(containerFormLogin)
			u.window.Show()
		}),
		widget.NewButton(constants.BtnRegistration, func() {
			msg, err := u.client.Ping(ctx)
			if err != nil {
				dialog.ShowInformation(msg, constants.ErrConnect+" ("+u.client.Cfg.HostGRPC+")", u.window)
				return
			}
			logger.Log.Info("Sign up")
			labelAlert.SetText("")
			labelAlert.Hide()
			u.window.SetContent(containerFormRegistration)
			u.window.Show()
		}),
	)

	formConfigContainer := container.NewVBox(formConfigButtons, labelAlert)
	formAuthContainer := container.NewVBox(formConfigContainer, separator, formAuthButtons)
	return formAuthContainer
}

func (u App) getFormLogin(username *widget.Entry, password *widget.Entry) *widget.Form {
	formLogin := widget.NewForm(
		widget.NewFormItem(constants.UsernameItem, username),
		widget.NewFormItem(constants.PasswordItem, password),
	)
	return formLogin
}

func (u App) getFormRegistration(UsernameRegistration *widget.Entry, PasswordRegistration *widget.Entry, NewPasswordEntryRegistration *widget.Entry) *widget.Form {
	formRegistration := widget.NewForm(
		widget.NewFormItem(constants.UsernameItem, UsernameRegistration),
		widget.NewFormItem(constants.PasswordItem, PasswordRegistration),
		widget.NewFormItem(constants.ConfirmPasswordItem, NewPasswordEntryRegistration),
	)
	return formRegistration
}
