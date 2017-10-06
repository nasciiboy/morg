package nirvana

import "github.com/nasciiboy/morg/katana"

func Show( doc *katana.Doc ){
  wout, err := Init()
	if err != nil {
		panic(err)
	}
  defer Close()

  wout.Echo = false

  pad := NewPad( wout )
  pad.ParseMorg( doc )
  pad.Draw()

  for {
		switch wout.Getch() {
    case 'q': return
    case KeyPgup      : pad.Scroll( PgUp )
    case KeyPgdn      : pad.Scroll( PgDown )
    case KeyArrowLeft : pad.Scroll( Left )
    case KeyArrowDown : pad.Scroll( Down )
    case KeyArrowRight: pad.Scroll( Right )
    case KeyArrowUp   : pad.Scroll( Up )
    case KeyHome      : pad.Scroll( Start )
    case KeyEnd       : pad.Scroll( End )
    case '1': pad.Scroll( DownLeft )
    case '2': pad.Scroll( Down )
    case '3': pad.Scroll( DownRight )
    case '4': pad.Scroll( Left )
    case '6': pad.Scroll( Right )
    case '7': pad.Scroll( UpLeft )
    case '8': pad.Scroll( Up )
    case '9': pad.Scroll( UpRight )
		}
	}
}
