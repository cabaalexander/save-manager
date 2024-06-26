package backend

import (
	"context"
	"runtime"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

type AppMenu struct {
	ctx  context.Context
	Menu *menu.Menu
}

func (m *AppMenu) Startup(ctx context.Context) {
	m.ctx = ctx
}

func (m *AppMenu) addSaveManagerMenu() {
	fileMenu := m.Menu.AddSubmenu("Save Manager")
	fileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		rt.Quit(m.ctx)
	})
}

func (m *AppMenu) addSavesMenu() {
	fileMenu := m.Menu.AddSubmenu("Saves")
	fileMenu.AddText("Quicksave", keys.CmdOrCtrl("s"), func(_ *menu.CallbackData) {
		rt.EventsEmit(m.ctx, "quickSave")
	})
	fileMenu.AddText("Quickload", keys.CmdOrCtrl("l"), func(_ *menu.CallbackData) {
		rt.EventsEmit(m.ctx, "quickLoad")
	})
}

func (m *AppMenu) addMacMenu() {
	if runtime.GOOS == "darwin" {
		m.Menu.Append(menu.EditMenu())
	}
}

func NewMenu() *AppMenu {
	newMenu := menu.NewMenu()

	appMenu := AppMenu{Menu: newMenu}
	appMenu.addSaveManagerMenu()
	appMenu.addSavesMenu()
	appMenu.addMacMenu()

	return &appMenu
}
