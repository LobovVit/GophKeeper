package ui

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/LobovVit/GophKeeper/internal/client/client"
	"github.com/LobovVit/GophKeeper/internal/client/config"
	"github.com/LobovVit/GophKeeper/internal/client/constants"
	"github.com/LobovVit/GophKeeper/internal/client/model"
)

// App - struct is used to create App.
type App struct {
	ui           fyne.App
	client       *client.Client
	accessToken  model.Token
	password     string
	window       fyne.Window
	data         map[string][][]string
	currentIndex map[string]int
	currentRow   map[string][]string
}

// New - method creates a new App.
func New(cfg *config.Config, ver string) (*App, error) {
	var t App
	t.ui = app.NewWithID("GophKeeper.ya")
	var err error
	t.client, err = client.New(cfg)
	if err != nil {
		return nil, err
	}
	t.accessToken = model.Token{}
	t.password = ""
	t.window = t.ui.NewWindow(ver)
	t.window.Resize(fyne.NewSize(400, 200))
	t.window.SetMaster()
	t.data = make(map[string][][]string)
	t.currentRow = make(map[string][]string)
	t.currentIndex = make(map[string]int)
	return &t, nil
}

// Run - method starts an App instance
func (u App) Run(ctx context.Context) error {
	u.window.Resize(fyne.NewSize(constants.WindowAuthWidth, constants.WindowAuthHeight))
	u.window.SetContent(u.getWelcomeForm(ctx))
	u.window.Show()

	//run
	u.window.ShowAndRun()
	return nil
}
