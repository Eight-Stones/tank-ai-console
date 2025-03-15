package view

import (
	"context"
	"sync"
	"time"

	"github.com/rivo/tview"

	"go-micro-service-template/entity"
	er "go-micro-service-template/pkg/error"
)

type TankManager interface {
	CreateGame(ctx context.Context) (string, error)
	StartGame(ctx context.Context, id string) error
	StopGame(ctx context.Context, id string) error

	Game(ctx context.Context, id string) (*entity.Game, error)
	Games(ctx context.Context) ([]*entity.Game, error)
	AddBot(ctx context.Context, id string) (string, error)
	View(ctx context.Context, id string) ([][]*entity.Cell, error)
}

const (
	emptyLabel     = `   `
	bulletLabel    = ` * `
	tankUpLabel    = ` ▲ `
	tankDownLabel  = ` ▼ `
	tankLeftLabel  = ` ◀ `
	tankRightLabel = ` ▶ `
)

var keyboardRunes = []rune{
	// Буквы (верхний ряд)
	'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p',

	// Буквы (средний ряд)
	'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l',

	// буквы (нижний ряд)
	'z', 'x', 'c', 'v', 'b', 'n', 'm',
}

type game struct {
	ID string
	mu sync.Mutex
}

func (g *game) setID(in string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.ID = in
}

func (g *game) getID() string {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.ID
}

type infoJobChan struct {
	start chan string
	end   chan struct{}
}

type View struct {
	opt          options
	game         game
	jobs         infoJobChan
	app          *tview.Application
	controls     *tview.List
	games        *tview.List
	gameControls *tview.List
	players      *tview.Table
	field        *tview.Table
	grid         *tview.Grid

	manager TankManager
}

func New(in ...Option) *View {
	cfg := options{
		redrawTimeout: time.Millisecond * 100,
	}

	for _, o := range in {
		o(&cfg)
	}

	app := tview.NewApplication()

	headerLabel := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText("Tanks AI")

	controlsLabel := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText("Controls")

	controls := tview.NewList()

	gamesLabel := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText("Games list")

	games := tview.NewList()

	gameControlsLabel := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText("Manage")

	gameControls := tview.NewList()

	playersLabel := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText("Players")

	players := tview.NewTable().
		SetBorders(true)

	fieldLabel := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText("Field")

	field := tview.NewTable().
		SetBorders(true)

	grid := tview.NewGrid().
		SetRows(1, 1, 35).
		SetColumns(20, 50, 20, 50, 85).
		SetBorders(true).AddItem(headerLabel, 0, 0, 1, 5, 0, 0, false).
		AddItem(controlsLabel, 1, 0, 1, 1, 0, 0, false).
		AddItem(gamesLabel, 1, 1, 1, 1, 0, 0, false).
		AddItem(gameControlsLabel, 1, 2, 1, 1, 0, 0, false).
		AddItem(playersLabel, 1, 3, 1, 1, 0, 0, false).
		AddItem(fieldLabel, 1, 4, 1, 1, 0, 0, false).
		AddItem(controls, 2, 0, 1, 1, 0, 0, false).
		AddItem(games, 2, 1, 1, 1, 0, 0, false).
		AddItem(gameControls, 2, 2, 1, 1, 0, 0, false).
		AddItem(players, 2, 3, 1, 1, 0, 0, false).
		AddItem(field, 2, 4, 1, 1, 0, 0, false)

	return &View{
		opt: cfg,
		game: game{
			ID: "",
			mu: sync.Mutex{},
		},
		jobs: infoJobChan{
			start: make(chan string),
			end:   make(chan struct{}),
		},
		app:          app,
		controls:     controls,
		games:        games,
		gameControls: gameControls,
		players:      players,
		field:        field,
		grid:         grid,
		manager:      cfg.manager,
	}
}

func (v *View) Run(ctx context.Context) error {
	v.drawJob(ctx)
	v.prepareControls()
	v.prepareGames()
	v.prepareGameControls()
	v.prepareGame()
	v.prepareField()
	v.createViewJob(ctx)

	if err := v.app.SetRoot(v.grid, true).SetFocus(v.controls).Run(); err != nil {
		return er.Wrap(err, "Run")
	}
	return nil
}

func (v *View) Stop() {
	v.app.Stop()
}
