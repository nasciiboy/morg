package nirvana

import (
  "time"

  term "github.com/nsf/termbox-go"
)

var Colors = [256]FontFace {
  { Fg: termColorDefault, Bg: termColorDefault  },
  { Fg: termColorBlack,   Bg: termColorBlack    }, { Fg: termColorRed, Bg: termColorBlack    }, { Fg: termColorGreen, Bg: termColorBlack   }, { Fg: termColorYellow, Bg: termColorBlack    }, { Fg: termColorBlue, Bg: termColorBlack    }, { Fg: termColorMagenta, Bg: termColorBlack    }, { Fg: termColorCyan, Bg: termColorBlack    }, { Fg: termColorWhite, Bg: termColorBlack    },
  { Fg: termColorBlack,   Bg: termColorRed      }, { Fg: termColorRed, Bg: termColorRed      }, { Fg: termColorGreen, Bg: termColorRed     }, { Fg: termColorYellow, Bg: termColorRed      }, { Fg: termColorBlue, Bg: termColorRed      }, { Fg: termColorMagenta, Bg: termColorRed      }, { Fg: termColorCyan, Bg: termColorRed      }, { Fg: termColorWhite, Bg: termColorRed      },
  { Fg: termColorBlack,   Bg: termColorGreen    }, { Fg: termColorRed, Bg: termColorGreen    }, { Fg: termColorGreen, Bg: termColorGreen   }, { Fg: termColorYellow, Bg: termColorGreen    }, { Fg: termColorBlue, Bg: termColorGreen    }, { Fg: termColorMagenta, Bg: termColorGreen    }, { Fg: termColorCyan, Bg: termColorGreen    }, { Fg: termColorWhite, Bg: termColorGreen    },
  { Fg: termColorBlack,   Bg: termColorYellow   }, { Fg: termColorRed, Bg: termColorYellow   }, { Fg: termColorGreen, Bg: termColorYellow  }, { Fg: termColorYellow, Bg: termColorYellow   }, { Fg: termColorBlue, Bg: termColorYellow   }, { Fg: termColorMagenta, Bg: termColorYellow   }, { Fg: termColorCyan, Bg: termColorYellow   }, { Fg: termColorWhite, Bg: termColorYellow   },
  { Fg: termColorBlack,   Bg: termColorBlue     }, { Fg: termColorRed, Bg: termColorBlue     }, { Fg: termColorGreen, Bg: termColorBlue    }, { Fg: termColorYellow, Bg: termColorBlue     }, { Fg: termColorBlue, Bg: termColorBlue     }, { Fg: termColorMagenta, Bg: termColorBlue     }, { Fg: termColorCyan, Bg: termColorBlue     }, { Fg: termColorWhite, Bg: termColorBlue     },
  { Fg: termColorBlack,   Bg: termColorMagenta  }, { Fg: termColorRed, Bg: termColorMagenta  }, { Fg: termColorGreen, Bg: termColorMagenta }, { Fg: termColorYellow, Bg: termColorMagenta  }, { Fg: termColorBlue, Bg: termColorMagenta  }, { Fg: termColorMagenta, Bg: termColorMagenta  }, { Fg: termColorCyan, Bg: termColorMagenta  }, { Fg: termColorWhite, Bg: termColorMagenta  },
  { Fg: termColorBlack,   Bg: termColorCyan     }, { Fg: termColorRed, Bg: termColorCyan     }, { Fg: termColorGreen, Bg: termColorCyan    }, { Fg: termColorYellow, Bg: termColorCyan     }, { Fg: termColorBlue, Bg: termColorCyan     }, { Fg: termColorMagenta, Bg: termColorCyan     }, { Fg: termColorCyan, Bg: termColorCyan     }, { Fg: termColorWhite, Bg: termColorCyan     },
  { Fg: termColorBlack,   Bg: termColorWhite    }, { Fg: termColorRed, Bg: termColorWhite    }, { Fg: termColorGreen, Bg: termColorWhite   }, { Fg: termColorYellow, Bg: termColorWhite    }, { Fg: termColorBlue, Bg: termColorWhite    }, { Fg: termColorMagenta, Bg: termColorWhite    }, { Fg: termColorCyan, Bg: termColorWhite    }, { Fg: termColorWhite, Bg: termColorWhite    },
}

func Init() (*Window, error) {
  err := term.Init()
  if err != nil { return nil, err }

  height, width := Size()

  stdscr := Window {
    Height: height,
    Width : width,
    BGChar: ' ',
    Looper: true,
    Echo  : true,
    Curs  : true,
    Delay : true,
    Resize: true,
  }

  stdscr.resize( height, width )

  return &stdscr, nil
}

func Close() {
  term.Close()
}

func Size() (height, width int ) {
  width, height = term.Size()
  return
}

func printCell( cell *Cell, y, x int ){
  tbg, tfg := cell.getTermColors()
  term.SetCell( x, y, cell.Ch, tfg, tbg )
  term.Flush()
  cell.Touch  = false
}

func SetFontFace( cPair uint8, Bg, Fg uint64 ){
  attrs, color, _, _ := extractData( Bg )
  Colors[ cPair ].Bg  = color
  attrs, color, _, _  = extractData( Fg )
  Colors[ cPair ].Fg  = color
  Colors[ cPair ].Attrs = attrs
}

func extractData( chtype uint64 ) (attrs, color uint8, keyMod uint16, r rune) {
  color  = uint8((chtype & hasColor) >> 48)
  attrs  = uint8(chtype >> 56)
  keyMod = uint16((chtype & hasMod) >> 32)

  if (chtype & hasKey) == 0 {
    r = rune(chtype & hasRune)
  }

  return
}

func extractCell( chtype uint64 ) (c Cell) {
  c.Color  = uint8((chtype & hasColor) >> 48)
  c.Attrs  = uint8(chtype >> 56)
  c.Touch  = true

  if (chtype & hasKey) == 0 {
    c.Ch = rune(chtype & hasRune)
  }

  return
}

func extractRune( chtype uint64 ) rune {
  if (chtype & hasKey) != 0 || (chtype & hasMod) != 0 {
    return 0
  }

  return rune(chtype & hasRune)
}

func extractAttrs( chtype uint64 ) uint8 {
  return uint8(chtype >> 56)
}

func getAttrs( chtype uint64 ) uint64 {
  return chtype & hasAttr
}

func Napms( ms uint ){
  var t time.Duration
  t = time.Duration( ms ) * time.Millisecond
  time.Sleep( t )
}

// int notimeout(WINDOW *win, bool bf);
// int overlay(const WINDOW *srcwin, WINDOW *dstwin);
// int raw(void);
// int noraw(void);
// int ungetch(int ch);
// int untouchwin(WINDOW *win);

func Hline( y, x, width int, fontFace FontFace ){
  tbg, tfg := fontFace.getTermColors()
  sum := -1

  if width < 0 { sum = 1 }

  for i := width; i != 0; i += sum {
    term.SetCell( x + i, y, '─', tfg, tbg )
  }
}

func Vline( y, x, width int, fontFace FontFace ){
  tbg, tfg := fontFace.getTermColors()
  sum := -1

  if width < 0 { sum = 1 }

  for i := width; i != 0; i += sum {
    term.SetCell( x, y + i, '│', tfg, tbg )
  }
}

func StrToCells( str string, color, attrs uint8  ) []Cell {
  c := make([]Cell, len( str ) )
  i := 0

  for _, ch := range str {
    c[i].Color, c[i].Attrs, c[i].Ch = color, attrs, ch
    i++
  }

  return c[:i]
}

func (c *FontFace) getTermColors() (bg, fg term.Attribute) {
  fg = term.Attribute( c.Fg ) | term.Attribute( c.Attrs ) << 8
  bg = term.Attribute( c.Bg )

  return
}
