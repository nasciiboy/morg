package nirvana

import (
  "github.com/nasciiboy/tui"
)

func Show( doc string ){
  wout, err := tui.Init()
	if err != nil {
		panic(err)
	}
  defer tui.Close()

  wout.Echo = false

  pad := tui.NewPad( wout )
  pad.ParseMorg( doc )
  pad.Draw()

  for {
		switch wout.Getch() {
    case 'q': return
    case tui.KeyPgup      : pad.Scroll( tui.PgUp )
    case tui.KeyPgdn      : pad.Scroll( tui.PgDown )
    case tui.KeyArrowLeft : pad.Scroll( tui.Left )
    case tui.KeyArrowDown : pad.Scroll( tui.Down )
    case tui.KeyArrowRight: pad.Scroll( tui.Right )
    case tui.KeyArrowUp   : pad.Scroll( tui.Up )
    case tui.KeyHome      : pad.Scroll( tui.Start )
    case tui.KeyEnd       : pad.Scroll( tui.End )
    case '1': pad.Scroll( tui.DownLeft )
    case '2': pad.Scroll( tui.Down )
    case '3': pad.Scroll( tui.DownRight )
    case '4': pad.Scroll( tui.Left )
    case '6': pad.Scroll( tui.Right )
    case '7': pad.Scroll( tui.UpLeft )
    case '8': pad.Scroll( tui.Up )
    case '9': pad.Scroll( tui.UpRight )
		}
	}
}
