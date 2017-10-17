package katana

import (
  "github.com/nasciiboy/txt"
  "github.com/nasciiboy/regexp4"
)

const (
  MarkupNil byte = iota
  MarkupErr
  MarkupEsc
  MarkupEmpty
)

type Markup struct {
  Left   []Markup
  Right  []Markup

  Type   byte
  Brace  byte
  Data   string
}

func (s *Scanner) GetFancyMarkup() (m Markup) {
  fancy := NewScanner( txt.Linelize( s.Text() ) ).QuietSplash().Init()
  m = fancy.GetMarkup()

  if fancy.ErrorCount > 0 { s.GetMarkup() }

  return
}

var ceparator = regexp4.Compile( "#?<:s*:<:>:s*>" )

func (s *Scanner) GetFancyCustomMarkup( label, brace byte ) (m Markup) {
  str := s.Text()
  if re := ceparator.Copy(); re.FindString( str ) {
    str = re.RplCatch( "<>", 1 )
  }

  fancy := NewScanner( txt.Linelize( str ) ).QuietSplash().Init()
  m = fancy.MarkupParser( label, brace )

  if fancy.ErrorCount > 0 { s.MarkupParser( label, brace ) }

  return
}

func (s *Scanner) GetMarkup() (m Markup) {
  last := s.RunePos
  for s.Next(); s.PrevRune != EOF; {
    if s.PrevRune == '@' {
      if last < s.PrevRunePos {
        m.Right = append( m.Right, Markup{ Data: s.Src[ last:s.PrevRunePos ] } )
      }

      m.Right = append( m.Right, s.MarckupTrigger() )
      last    = s.PrevRunePos
      continue
    }

    s.Next()
  }

  if len(m.Right) == 0 { m.Data = s.Src[ last: ]; return }

  if last < s.PrevRunePos {
    m.Right = append( m.Right, Markup{ Type: MarkupNil, Data: s.Src[last:] } )
  }

  return
}

func (s *Scanner) MarckupTrigger() (node Markup) {
  last := s.PrevRunePos // @
  for i := 1; s.Next() != EOF; i++ {
    switch s.PrevRune {
    case '(', '<', '[', '{':
      if i == 1 {
        node.Type, node.Data = MarkupNil, s.Src[s.PrevRunePos:s.RunePos]
        s.Next() // skip: ), }, >, ]
        return
      }

      s.Next() // skip: (, [, <, {
      node = s.MarkupParser( s.Src[s.PrevRunePos - 2], s.Src[s.PrevRunePos - 1] )

      if i == 2 { return node }

      return mm( s.Src[last + 1:last + i], &node )
    case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
      '!', '#', '$', '%', '&', '*', '+', '-', '.', ':', '=', '?', '^', '_', '|', '~', '/', '\\', '"', '`', '\'',
      'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
      'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
    case '@', ')', '>', ']', '}':
      if i == 1 {
        node.Type, node.Data = MarkupEsc, s.Src[s.PrevRunePos:s.RunePos]
        s.Next()
        return
      }

      s.Error( "Markup: un-open markup" )
      node.Type, node.Data = MarkupErr, s.Src[last:s.RunePos]
      s.Next()
      return
    default:
      s.Error( "Markup: unknow markup" )
      // return Markup{ Type: MarkupErr, Data: s.Src[last:s.RunePos] }
      node.Type, node.Data = MarkupErr, s.Src[last:s.RunePos]
      s.Next()
      return
    }
  }

  s.Error( "Markup: init multi-markup" )
  return Markup{ Type: MarkupErr, Data: s.Src[last:] }
}

func mm( str string, last *Markup ) (node Markup) {
  if len( str ) < 2 { return *last }

  node.Type, node.Brace = str[0], last.Brace
  node.Right = append( node.Right, mm( str[1:], last ) )
  return
}

func (s *Scanner) MarkupParser( label, brace byte ) (node Markup) {
  if s.PrevRune == EOF {
    s.Error( "Markup: empty markup" )
    return Markup{ Type: MarkupErr, Brace: brace }
  }

  node.Type, node.Brace = label, brace
  brace, last := comBrace( brace ), s.PrevRunePos

  for s.PrevRune != EOF {
    if s.PrevRune == '<' && s.Rune == '>' {
      if last < s.PrevRunePos {
        node.Right = append( node.Right, Markup{ Data: s.Src[last:s.PrevRunePos] } )
      }

      node.Left, node.Right = node.Right, nil
      s.Next(); // '>'
      last = s.RunePos
    } else {
      if s.Src[s.PrevRunePos] == brace {
        if len(node.Left) == 0 && len(node.Right) == 0 {
          node.Data = s.Src[last:s.PrevRunePos];
        } else if last < s.PrevRunePos {
          node.Right = append( node.Right, Markup{ Data: s.Src[last:s.PrevRunePos] } )
        }

        s.Next() // ), >, ], }
        return  node
      }

      if s.PrevRune == '@' {
        if last < s.PrevRunePos {
          node.Right = append( node.Right, Markup{ Data: s.Src[last:s.PrevRunePos] } )
        }

        node.Right = append( node.Right, s.MarckupTrigger() )
        last = s.PrevRunePos
        continue
      }
    }

    s.Next()
  }

  if len(node.Left) == 0 && len(node.Right) == 0 {
    node.Data = s.Src[last:]
  } else if last < s.PrevRunePos {
    node.Right = append( node.Right, Markup{ Data: s.Src[last:], Left: nil, Right: nil } )
  }

  if brace != 0 { s.Error( "Markup not terminated" ) }
  return node
}

func (m *Markup) HasSomething() bool {
  if len( m.Data ) != 0 || len(m.Right) != 0 || len(m.Left) != 0 {
    return true
  }

  return false
}

func (m *Markup) String() (str string) {
  if len( m.Right ) == 0 { return m.Data }

  for _, r := range m.Right {
    str += r.String()
  }

  return str
}

func (m *Markup) GetLeft() (str string) {
  if len( m.Left ) == 0 { return m.Data }

  for _, l := range m.Left {
    str += l.String()
  }

  return str
}

func (m *Markup) MakeLeft() (str string) {
  if len(m.Data) > 0 { return m.Data }

  if len( m.Left ) > 0 {
    for _, l := range m.Left {
      str += l.String()
    }
    return str
  }

  for _, r := range m.Right {
    str += r.String()
  }

  return
}

func (m *Markup) Rebuild() (str string) {
  return m.rebuild( false )
}

func (m *Markup) rebuild( mm bool ) (str string){
  if m.Type != 0 && len( m.Data ) == 0 && len( m.Right ) == 1 && len( m.Left ) == 0 && m.Brace != 0 {
    if !mm { return "@" + m.rebuild( true ) }

    return string( m.Type ) + m.Right[0].rebuild( true )
  }

  if len( m.Left ) == 0 && len( m.Right ) == 0 {
    if m.Type == MarkupEsc { return "@" + m.Data }

    if m.Type >= ' ' && m.Brace != 0 {
      if !mm { str = "@" }
      return str + string(m.Type) + string(m.Brace) + m.Data + string( comBrace( m.Brace ) )
    }

    return m.Data
  }

  if m.Type >= ' ' && m.Brace != 0 {
    if !mm { str = "@" }
    str +=  string(m.Type) + string(m.Brace)
  }

  for _, c := range m.Left {
    str += c.rebuild( false )
  }

  if len( m.Left ) > 0 { str += "<>" }

  for _, c := range m.Right {
    str += c.rebuild( false )
  }

  if m.Type >= ' ' && m.Brace != 0 { str += string( comBrace( m.Brace ) ) }

  return str
}

func comBrace( op byte ) byte {
  switch op {
  case '(': return ')'
  case '[': return ']'
  case '{': return '}'
  case '<': return '>'
  case ')': return '('
  case ']': return '['
  case '}': return '{'
  case '>': return '<'
  default : return 0
  }
}
