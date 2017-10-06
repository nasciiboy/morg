package nirvana

import (
  term "github.com/nsf/termbox-go"
)

type Window struct {
	Height,  Width int   // maximums of x and y, NOT window size
	CurY,    CurX  int   // current cursor position
	InitY,   InitX int   // screen coords of upper-left-hand corner

  Color   uint8
  Attrs   uint8
	BGChar  rune         // current background char

  Buffer [][]Cell

	// bool	_notimeout;	// no time out on function-key entry?

  Looper bool
  Scroll bool
  Touch  bool
  Echo   bool
  Curs   bool
  Delay  bool
  Resize bool
  // Beep   bool
  // Flash  bool
}

func NewWindow( height, width, yPos, xPos int ) *Window {
  if height <= 0 || width <= 0 {
    height, width = Size()
  }

  w := &Window{
    Height: height,
    Width : width,
    InitY: yPos,
    InitX: xPos,
    BGChar: ' ',
    Looper: true,
    Echo  : true,
    Curs  : true,
  }

  w.resize( height, width )

  return w
}

func (w *Window) Size() (height, width int) {
  return w.Height, w.Width
}

func (w *Window) AddCh( chtype uint64 ) {
  w.Touch = true

  cell := &w.Buffer[w.CurY][w.CurX]
  chAttrs, chColor, _, r := extractData( chtype )

  cell.Ch, cell.Touch = r, true

  if chAttrs == 0 { cell.Attrs = w.Attrs
  }  else         { cell.Attrs = chAttrs | w.Attrs }

  if chColor == 0 { cell.Color = w.Color
  }  else         { cell.Color = chColor }

  if w.Echo {
    printCell( cell, w.InitY + w.CurY, w.InitX + w.CurX )
  }

  w.mvCurs( r == '\n' )

  if w.Curs {
    term.SetCursor( w.InitX + w.CurX, w.InitY + w.CurY )
    term.Flush()
  } else {
    term.HideCursor()
    term.Flush()
  }
}

func (w *Window) AddCell( cell Cell ) {
  w.Touch = true

  cell.Touch = true
  w.Buffer[w.CurY][w.CurX] = cell

  if w.Echo {
    printCell( &cell, w.InitY + w.CurY, w.InitX + w.CurX )
  }

  w.mvCurs( cell.Ch == '\n' )

  if w.Curs {
    term.SetCursor( w.InitX + w.CurX, w.InitY + w.CurY )
    term.Flush()
  } else {
    term.HideCursor()
    term.Flush()
  }
}

func (w *Window) Mv( y, x int ) (err bool) {
  if y < 0 || x < 0 || y >= w.Height || x >= w.Width {
    return true
  }

  w.CurY, w.CurX = y, x
  return false
}

func (w *Window) MvAddCh( y, x int, chtype uint64 ) (err bool) {
  if w.Mv( y, x ) { return true }

  w.AddCh( chtype )
  return false
}

func (w *Window) AddChs( chs []uint64 ){
  for _, chtype := range chs {
    w.AddCh( chtype )
  }
}

func (w *Window) MvAddChs( y, x int, chs []uint64 ) (err bool) {
  if w.Mv( y, x ) { return true }

  w.AddChs( chs )
  return false
}

func (w *Window) AddCells( cells []Cell ){
  for _, cell := range cells {
    w.AddCell( cell )
  }
}

func (w *Window) MvAddCells( y, x int, cells []Cell ) (err bool) {
  if w.Mv( y, x ) { return true }

  w.AddCells( cells )
  return false
}

func (w *Window) AddStr( str string ) {
  w.AddCells( w.StrToCells( str ) )
}

func (w *Window) mvCurs( nl bool ) {
  if w.Looper { w.mvLooper( nl ); return }

  if w.Scroll { w.mvScroll( nl ); return }

  if w.CurY + 1 < w.Height && (nl || w.CurX + 1 >= w.Width) {
    w.CurX = 0
    w.CurY++
    return
  }

  if w.CurX + 1 < w.Width {
    w.CurX++
  }
}

func (w *Window) mvLooper( nl bool ) {
  if nl || w.CurX + 1 >= w.Width {
    w.CurX = 0
  } else {
    w.CurX++
    return
  }

  if w.CurY + 1 < w.Height { w.CurY++
  } else                   { w.CurY = 0 }
}

func (w *Window) mvScroll( nl bool ) {
  if nl || w.CurX + 1 == w.Width {
    w.CurX = 0
  } else {
    w.CurX++
    return
  }

  if w.CurY + 1 < w.Height {
    w.CurY++
  } else {
    for y := 1; y < w.Height; y++ {
      for x := 0; x < w.Width; x++ {
        w.Buffer[y-1][x] = w.Buffer[y][x]
        w.Buffer[y-1][x].Touch = true
      }
    }

    cleanCell := Cell{ w.BGChar, w.Color, w.Attrs, true }
    for y := w.Height - 1; y < w.Height; y++ {
      for x := 0; x < w.Width; x++ {
        w.Buffer[y][x] = cleanCell
      }
    }

    w.Touch = true
    w.Refresh()
  }
}

func (w *Window) resize( height, width int ) {
  if cap( w.Buffer ) > height {
    w.Buffer = w.Buffer[:height]
  } else if len( w.Buffer ) < height {
    for len( w.Buffer ) < height {
      row := make( []Cell, width )
      for i := 0; i < len( row ); i++ {
        row[ i ] = Cell{ Ch: w.BGChar, Color: w.Color, Attrs: w.Attrs, Touch: true }
      }

      w.Buffer = append( w.Buffer, row )
    }
  }

  for y := 0; y < len( w.Buffer ); y++ {
    if cap( w.Buffer[ y ] ) > width {
      w.Buffer[ y ] = w.Buffer[y][:width]
    } else if len( w.Buffer[ y ] ) < width {
      for x := len( w.Buffer[ y ] ); x < width; x++ {
        w.Buffer[ y ] = append( w.Buffer[ y ], Cell{ } )
      }
    }
  }

  buff := make([][]Cell, height)
  for i := 0; i < height; i++ {
    buff[i] = make([]Cell, width)
  }

  w.Buffer = buff

  w.Height, w.Width = height, width
}

func (w *Window) Refresh() {
  if w.Touch {
    for y := 0; y < w.Height; y++ {
      for x := 0; x < w.Width; x++ {
        if w.Buffer[y][x].Touch {
          w.Buffer[y][x].Touch = false
          termBg, termFg := w.Buffer[y][x].getTermColors()
          term.SetCell( w.InitX + x, w.InitY + y, w.Buffer[y][x].Ch, termFg, termBg )
        }
      }
    }

    w.Touch = false
    term.Flush()
  }
}

func (w *Window) Draw() {
  for y := 0; y < w.Height; y++ {
    for x := 0; x < w.Width; x++ {
      termBg, termFg := w.Buffer[y][x].getTermColors()
      term.SetCell( w.InitX + x, w.InitY + y, w.Buffer[y][x].Ch, termFg, termBg )
    }
  }

  w.Touch = false
  term.Flush()
}

func (w *Window) Clear(){
  defaultCell := Cell{ Ch: w.BGChar, Color: w.Color, Touch: true }

  for y := 0; y < len( w.Buffer ); y++ {
    for x := 0; x < len( w.Buffer[y] ); x++ {
      w.Buffer[y][x] = defaultCell
    }
  }

  w.Touch = true
}

// int clrtobot(void);
// int clrtoeol(void);

func (w *Window) MvChFace( y, x int, face uint64 ) (err bool) {
  if y < 0 || x < 0 || y >= w.Height || x >= w.Width {
    return true
  }

  w.Buffer[y][x].Attrs, w.Buffer[y][x].Color, _, _ = extractData( face )
  w.Buffer[y][x].Touch, w.Touch = true, true

  return false
}

func (w *Window) ChFace( face uint64 ){
  w.Attrs, w.Color, _, _ = extractData( face )

  for y := 0; y < w.Height; y++ {
    for x := 0; x < w.Width; x++ {
      w.Buffer[y][x].Attrs = w.Attrs
      w.Buffer[y][x].Color = w.Color
      w.Buffer[y][x].Touch = true
    }
  }

  w.Touch = true
}

func (w *Window) SetFace( face uint64 ){
  w.Attrs, w.Color, _, _ = extractData( face )
}

func (w *Window) SetColor( face uint64 ){
  _, w.Color, _, _ = extractData( face )
}

func (w *Window) SetAttrs( attrs uint64 ){
  w.Attrs = extractAttrs( attrs )
}

func (w *Window) Attron( attrs uint64 ){
  w.Attrs |= extractAttrs( attrs )
}


func (w *Window) Attroff( attrs uint64 ){
}


// int baudrate(void);
// int beep(void);

// int border(chtype ls, chtype rs, chtype ts, chtype bs, chtype tl, chtype tr, chtype bl, chtype br);

// int box(WINDOW *win, chtype verch, chtype horch);
// bool can_change_color(void);

// int cbreak(void);
// int nocbreak(void);

func (w *Window) CursSet( visibility bool ){
  if visibility {
    term.SetCursor( w.InitX + w.CurX, w.InitY + w.CurY )
  } else {
    term.HideCursor()
  }

  w.Curs = visibility
  term.Flush()
}

// int delch(void);
// int mvdelch(int y, int x);
// int deleteln(void);
// int flash(void);
// int flushinp(void); // clear imput buffer

func (w *Window) Getch() uint64 {
  for {
    event := term.PollEvent()
    if event.Type == term.EventKey {
      if event.Ch != 0 {
        if w.Echo { w.AddCh( uint64(event.Ch) ) }
        return uint64(event.Ch)
      }

      switch event.Key {
      case term.KeyF1:            return KeyF1
      case term.KeyF2:            return KeyF2
      case term.KeyF3:            return KeyF3
      case term.KeyF4:            return KeyF4
      case term.KeyF5:            return KeyF5
      case term.KeyF6:            return KeyF6
      case term.KeyF7:            return KeyF7
      case term.KeyF8:            return KeyF8
      case term.KeyF9:            return KeyF9
      case term.KeyF10:           return KeyF10
      case term.KeyF11:           return KeyF11
      case term.KeyF12:           return KeyF12
      case term.KeyInsert:        return KeyInsert
      case term.KeyDelete:        return KeyDelete
      case term.KeyHome:          return KeyHome
      case term.KeyEnd:           return KeyEnd
      case term.KeyPgup:          return KeyPgup
      case term.KeyPgdn:          return KeyPgdn
      case term.KeyArrowUp:       return KeyArrowUp
      case term.KeyArrowDown:     return KeyArrowDown
      case term.KeyArrowLeft:     return KeyArrowLeft
      case term.KeyArrowRight:    return KeyArrowRight
      case term.MouseLeft:        return Mouse
      case term.MouseMiddle:      return Mouse
      case term.MouseRight:       return Mouse
      case term.MouseRelease:     return Mouse
      case term.MouseWheelUp:     return Mouse
      case term.MouseWheelDown:   return Mouse
      case term.KeyCtrlSpace:     return ' ' | Ctrl
      case term.KeyCtrlA:         return 'a' | Ctrl
      case term.KeyCtrlB:         return 'b' | Ctrl
      case term.KeyCtrlC:         return 'c' | Ctrl
      case term.KeyCtrlD:         return 'd' | Ctrl
      case term.KeyCtrlE:         return 'e' | Ctrl
      case term.KeyCtrlF:         return 'f' | Ctrl
      case term.KeyCtrlG:         return 'g' | Ctrl
      // case term.KeyCtrlH:         return 'h' | Ctrl
      // case term.KeyCtrlI:         return 'i' | Ctrl
      case term.KeyCtrlJ:         return 'j' | Ctrl
      case term.KeyCtrlK:         return 'k' | Ctrl
      case term.KeyCtrlL:         return 'l' | Ctrl
      //case term.KeyCtrlM:         return 'm' | Ctrl
      case term.KeyCtrlN:         return 'n' | Ctrl
      case term.KeyCtrlO:         return 'o' | Ctrl
      case term.KeyCtrlP:         return 'p' | Ctrl
      case term.KeyCtrlQ:         return 'q' | Ctrl
      case term.KeyCtrlR:         return 'r' | Ctrl
      case term.KeyCtrlS:         return 's' | Ctrl
      case term.KeyCtrlT:         return 't' | Ctrl
      case term.KeyCtrlU:         return 'u' | Ctrl
      case term.KeyCtrlV:         return 'v' | Ctrl
      case term.KeyCtrlW:         return 'w' | Ctrl
      case term.KeyCtrlX:         return 'x' | Ctrl
      case term.KeyCtrlY:         return 'y' | Ctrl
      case term.KeyCtrlZ:         return 'z' | Ctrl
      case term.KeyCtrl5:         return '5' | Ctrl
      case term.KeyCtrl6:         return '6' | Ctrl
      case term.KeyCtrl7:         return '7' | Ctrl
      // case term.KeyCtrl8:         return '7' | Ctrl
      case term.KeyEsc:           return KeyEsc
      case term.KeyCtrlBackslash: return '\\' | Ctrl
      case term.KeyEnter:
        if w.Echo { w.AddCh( '\n' ) }
        return '\n'
      case term.KeySpace:
        if w.Echo { w.AddCh( ' ' ) }
        return ' '
      case term.KeyTab:
        if w.Echo { w.AddStr( "    "  ) }
        return '\t'
      case term.KeyBackspace :    return KeyBackspace
      case term.KeyBackspace2:    return KeyBackspace
      }
    }
  }
}

// int getmouse(MEVENT *event);
// void getparyx(WINDOW *win, int y, int x);
// int getstr(char *str);
// int getnstr(char *str, int n);
// int halfdelay(int tenths);
// bool has_colors(void);
// int hline(chtype ch, int n);
// void idcok(WINDOW *win, bool bf);
// void immedok(WINDOW *win, bool bf);
// chtype inch(void);
// chtype winch(WINDOW *win);
// chtype mvinch(int y, int x);
// chtype mvwinch(WINDOW *win, int y, int x);
// bool mouse_trafo(int* pY, int* pX, bool to_screen);
// bool wmouse_trafo(const WINDOW* win, int* pY, int* pX, bool to_screen);
// int notimeout(WINDOW *win, bool bf);
// int overlay(const WINDOW *srcwin, WINDOW *dstwin);
// int raw(void);
// int noraw(void);
// int ungetch(int ch);
// int untouchwin(WINDOW *win);
// int vline(chtype ch, int n);
func (w *Window) ValidPos( pos Gps ) bool {
  if pos.Y < 0 || pos.X < 0 || pos.Y >= w.Height || pos.X >= w.Width {
    return false
  }

  return true
}

func (w *Window) Hline( y, x, width int ){
  if w.ValidPos( Gps{ Y: y, X: x } ) &&
     w.ValidPos( Gps{ Y: y, X: x + width } ) {
    Hline( y, x, width, w.GetFace() )
  }
}

func (w *Window) Vline( y, x, width int ){
  if w.ValidPos( Gps{ Y: y, X: x } ) &&
     w.ValidPos( Gps{ Y: y + width, X: x } ) {
    Vline( y, x, width, w.GetFace() )
  }
}

func (w *Window) Box(){
  face := w.GetFace()
  Hline( w.InitY - 1, w.InitX - 1, w.Width, face )
  Vline( w.InitY - 1, w.InitX - 1, w.Height, face )

  Hline( w.InitY + w.Height, w.InitX - 1, w.Width, face )
  Vline( w.InitY - 1, w.InitX + w.Width, w.Height, face )

  tbg, tfg := w.getTermColors()
  term.SetCell( w.InitX - 1, w.InitY - 1, '┌', tfg, tbg )
  term.SetCell( w.InitX + w.Width, w.InitY - 1, '┐', tfg, tbg )
  term.SetCell( w.InitX - 1, w.InitY + w.Height, '└', tfg, tbg )
  term.SetCell( w.InitX + w.Width, w.InitY + w.Height, '┘', tfg, tbg )
}

func (w *Window) StrToCells( str string ) []Cell {
  c := make([]Cell, len( str ) )
  i := 0

  for _, ch := range str {
    c[i].Color, c[i].Attrs, c[i].Ch = w.Color, w.Attrs, ch
    i++
  }

  return c[:i]
}

func (w *Window) GetFace() (ff FontFace) {
  ff       = Colors[ w.Color ]
  ff.Attrs = w.Attrs

  return
}

func (w *Window) GetDefaultCell() Cell {
  return Cell{ Color: w.Color, Attrs: w.Attrs, Touch: true, Ch: w.BGChar }
}

func (w *Window) getTermColors() (bg, fg term.Attribute) {
  color := Colors[ w.Color ]
  fg = term.Attribute( color.Fg ) | term.Attribute( w.Attrs ) << 8
  bg = term.Attribute( color.Bg )

  return
}

func (c *Cell) getTermColors() (bg, fg term.Attribute) {
  color := Colors[ c.Color ]
  fg = term.Attribute( color.Fg ) | term.Attribute( c.Attrs ) << 8
  bg = term.Attribute( color.Bg )

  return
}
