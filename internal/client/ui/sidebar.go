package ui

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/LobovVit/GophKeeper/internal/client/constants"
)

// sidebar defines the data structure
type sidebar struct {
	title string
	view  func(ctx context.Context, w fyne.Window) fyne.CanvasObject
}

func (u App) makeSidebar(ctx context.Context, setSidebar func(sidebar sidebar)) fyne.CanvasObject {
	var sidebarItems = map[string]sidebar{
		"login":  {"Logins", u.loginPassword},
		"text":   {"Texts", u.text},
		"binary": {"Files", u.binary},
		"card":   {"Bank cards", u.card}}
	var sidebarIndex = map[string][]string{
		"": {"login", "text", "binary", "card"}}

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return sidebarIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := sidebarIndex[uid]
			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := sidebarItems[uid]
			if !ok {
				fyne.LogError("Missing panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.title)
			obj.(*widget.Label).TextStyle = fyne.TextStyle{}
		},
		OnSelected: func(uid string) {
			if t, ok := sidebarItems[uid]; ok {
				setSidebar(t)
			}
		},
	}
	return tree
}

func (u App) loginPassword(ctx context.Context, w fyne.Window) fyne.CanvasObject {
	var tblLoginPassword *widget.Table
	tblLoginPassword = widget.NewTable(
		func() (int, int) {
			return len(u.data["dataTblLoginPassword"]), len(u.data["dataTblLoginPassword"][0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel(constants.TblLabel)
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(u.data["dataTblLoginPassword"][i.Row][i.Col])
		})

	tblLoginPassword.OnSelected = func(id widget.TableCellID) {
		u.currentIndex["dataTblLoginPassword"] = id.Row
		u.currentRow["dataTblLoginPassword"] = u.data["dataTblLoginPassword"][id.Row]
	}

	setColWidthLoginPassword(tblLoginPassword)
	return container.NewBorder(nil, u.getBtnLoginPassword(ctx), nil, nil, tblLoginPassword)
}

func (u App) text(ctx context.Context, w fyne.Window) fyne.CanvasObject {
	var tblText *widget.Table
	tblText = widget.NewTable(
		func() (int, int) {
			return len(u.data["dataTblText"]), len(u.data["dataTblText"][0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel(constants.TblLabel)
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(u.data["dataTblText"][i.Row][i.Col])
		})

	tblText.OnSelected = func(id widget.TableCellID) {
		u.currentIndex["dataTblText"] = id.Row
		u.currentRow["dataTblText"] = u.data["dataTblText"][id.Row]
	}

	setColWidthText(tblText)
	return container.NewBorder(nil, u.getBtnText(ctx), nil, nil, tblText)
}

func (u App) binary(ctx context.Context, w fyne.Window) fyne.CanvasObject {
	var tblBinary *widget.Table
	tblBinary = widget.NewTable(
		func() (int, int) {
			return len(u.data["dataTblBinary"]), len(u.data["dataTblBinary"][0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel(constants.TblLabel)
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(u.data["dataTblBinary"][i.Row][i.Col])
		})

	tblBinary.OnSelected = func(id widget.TableCellID) {
		u.currentIndex["dataTblBinary"] = id.Row
		u.currentRow["dataTblBinary"] = u.data["dataTblBinary"][id.Row]
	}

	setColWidthBinary(tblBinary)
	return container.NewBorder(nil, u.getBtnBinary(ctx), nil, nil, tblBinary)
}

func (u App) card(ctx context.Context, w fyne.Window) fyne.CanvasObject {
	var tblCard *widget.Table
	tblCard = widget.NewTable(
		func() (int, int) {
			return len(u.data["dataTblCard"]), len(u.data["dataTblCard"][0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel(constants.TblLabel)
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(u.data["dataTblCard"][i.Row][i.Col])
		})

	tblCard.OnSelected = func(id widget.TableCellID) {
		u.currentIndex["dataTblCard"] = id.Row
		u.currentRow["dataTblCard"] = u.data["dataTblCard"][id.Row]
	}

	setColWidthCard(tblCard)
	return container.NewBorder(nil, u.getBtnCard(ctx), nil, nil, tblCard)
}

func setColWidthCard(table *widget.Table) {
	colWidths := []float32{150, 150, 150, 200, 150, 50, 150, 150, 150, 150}
	for idx, colWidth := range colWidths {
		table.SetColumnWidth(idx, colWidth)
	}
}

func setColWidthText(table *widget.Table) {
	colWidths := []float32{150, 150, 500, 150, 150, 150}
	for idx, colWidth := range colWidths {
		table.SetColumnWidth(idx, colWidth)
	}
}

func setColWidthLoginPassword(table *widget.Table) {
	colWidths := []float32{150, 150, 150, 150, 150, 150, 150}
	for idx, colWidth := range colWidths {
		table.SetColumnWidth(idx, colWidth)
	}
}

func setColWidthBinary(table *widget.Table) {
	colWidths := []float32{150, 150}
	for idx, colWidth := range colWidths {
		table.SetColumnWidth(idx, colWidth)
	}
}
