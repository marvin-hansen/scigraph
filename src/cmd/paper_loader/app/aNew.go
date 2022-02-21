package app

type App struct {
	state *State
}

func NewApp() (a *App) {
	a = &App{state: newState()}
	a.init()
	return a
}
