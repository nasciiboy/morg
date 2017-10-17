package katana

import (
  "fmt"
  "os"
  "unicode"
  "unicode/utf8"
)

type Position struct {
  Name   string // filename, if any
  Offset int    // byte offset, starting at 0
  Line   int    // line number, starting at 1
  Column int    // column number, starting at 1 (character count per line)
}

func (pos *Position) IsValid() bool { return pos.Line > 0 }

func (pos Position) String() string {
  s := pos.Name
  if s == "" { s = "<input>" }

  if pos.Line > 0 { // IsValid()
    s += fmt.Sprintf(":%d:%d", pos.Line, pos.Column)
  }
  return s
}

type Token struct {
  Type, Rune, PrevRune rune
  Text string
  Position
}

const (
  ScanIdents     = 1 << -ScanIdent
  ScanBools      = 1 << -ScanBool  // includes Indents
  ScanInts       = 1 << -ScanInt   // includes Octals and Hexadecimals
  ScanFloats     = 1 << -ScanFloat // includes Ints
  ScanChars      = 1 << -ScanChar
  ScanStrings    = 1 << -ScanString
  ScanRawStrings = 1 << -ScanRawString
  ScanComments   = 1 << -ScanComment
  SkipComments   = 1 << -SkipComment // if set with ScanComments, comments become white space
  GoTokens       = ScanIdents | ScanBools | ScanFloats | ScanChars | ScanStrings | ScanRawStrings | ScanComments | SkipComments
)

const (
  EOF = -(iota + 1)
  ScanIdent
  ScanBool
  ScanInt
  ScanOctal
  ScanHexadecimal
  ScanFloat
  ScanChar
  ScanString
  ScanRawString
  ScanComment
  SkipComment
)

var tokenStrings = map[rune]string{
  EOF            : "EOF",
  ScanIdent      : "Ident",
  ScanBool       : "Bool",
  ScanInt        : "Int",
  ScanOctal      : "Octal",
  ScanHexadecimal: "Hexadecimal",
  ScanFloat      : "Float",
  ScanChar       : "Char",
  ScanString     : "String",
  ScanRawString  : "RawString",
  ScanComment    : "Comment",
}

func TokenString(tok rune) string {
  if s, found := tokenStrings[tok]; found { return s }

  return fmt.Sprintf("%q", string(tok))
}

const GoWhitespace       = 1<<'\t' | 1<<'\n' | 1<<'\r' | 1<<' '
const catchNewLineInScan = 1<<'\t' | 1<<'\r' | 1<<' '

func quietSplash( s *Scanner, msg string ) {}

type Scanner struct {
  Name         string
  Src          string
  SrcPos       int
  Mode         uint
  Whitespace   uint64

  line         int
  column       int
  lastLineLen  int

  Line         string
  getLine      bool
  ninja        bool

  Rune         rune
  PrevRune     rune
  RunePos      int
  PrevRunePos  int

  LastToken    Token

  CustomError  func(s *Scanner, msg string)
  ErrorCount   int
}

func NewScanner( src string ) *Scanner {
  s := new(Scanner)
  s.Src    = src
  s.SrcPos = 0

  s.line        = 1
  s.column      = 0
  s.lastLineLen = 0
  s.RunePos     = 0
  s.getLine     = true

  s.ErrorCount = 0
  s.Mode       = GoTokens
  s.Whitespace = GoWhitespace

  return s
}

func (s *Scanner) QuietSplash() *Scanner {
  s.CustomError = quietSplash
  return s
}

func (s *Scanner) Init() *Scanner {
  if s.SrcPos == 0 { s.next() }
  return s
}

func (s *Scanner) next() rune {
  if s.SrcPos >= len( s.Src ) {
    // previous character was not EOF
    if s.RunePos != s.SrcPos { s.column++ }

    s.PrevRune, s.PrevRunePos = s.Rune, s.RunePos
    s.Rune, s.RunePos = EOF, s.SrcPos
    s.Line = ""
    s.getLine = false
    return s.Rune
  }

  if s.getLine {
    s.Line, s.getLine = getRawLine( s.Src[ s.SrcPos: ] ), false
  }

  ch, width := rune( s.Src[ s.SrcPos ] ), 1

  if ch >= utf8.RuneSelf {
    ch, width = utf8.DecodeRuneInString( s.Src[ s.SrcPos: ] )
  }

  // advance
  s.PrevRune, s.PrevRunePos = s.Rune, s.RunePos
  s.Rune, s.RunePos = ch, s.SrcPos
  s.SrcPos += width
  s.column++

  switch ch {
  case 0:
    if !s.ninja { s.Error("illegal character NUL")  }
  case utf8.RuneError:
    if !s.ninja { s.Error("illegal UTF-8 encoding") }
  case '\n':
    s.getLine = true
    s.line++
    s.lastLineLen = s.column
    s.column = 0
  }

  return ch
}

func (s *Scanner) Next() rune {
  s.PrevRune, s.PrevRunePos = s.Rune, s.RunePos

  if s.Rune != EOF { s.next() }

  return s.PrevRune
}

func (s *Scanner) RestOfLine() string {
  if s.Rune == EOF { return "" }

  return getRawLine( s.Src[ s.RunePos:] )
}

func (s *Scanner) NextLine() bool {
  if  s.Rune == EOF { return false }

  for s.Rune != '\n' && s.Rune != EOF { s.Next() }
  s.Next()

  return true
}

func (s *Scanner) Text() string {
  if s.RunePos >= len( s.Src ) { return "" }

  return s.Src[s.RunePos:]
}

func (s *Scanner) PrevText() string {
  if s.Rune == EOF { return s.Src }

  return s.Src[:s.RunePos]
}

func (s *Scanner) NinjaMoves( fwd int ){
  s.ninja = true
  for x := 0; s.Rune != EOF && x < fwd; x++ { s.next() }
  s.ninja = false
}

func (s *Scanner) NinjaLenMoves( fwd int ){
  s.ninja = true
  init := s.RunePos
  for s.Rune != EOF && (s.RunePos - init < fwd) { s.next() }
  s.ninja = false
}

func (s *Scanner) Error(msg string) {
  s.ErrorCount++
  if s.CustomError != nil {
    s.CustomError(s, msg)
    return
  }

  fmt.Fprintf(os.Stderr, "%s: %s\n", s.Pos(), msg)
}

func (s *Scanner) Pos() (pos Position) {
  pos.Name   = s.Name
  pos.Offset = s.RunePos

  switch {
  case s.column > 0:
    // common case: last character was not a '\n'
    pos.Line = s.line
    pos.Column = s.column
  case s.lastLineLen > 0:
    // last character was a '\n'
    pos.Line = s.line - 1
    pos.Column = s.lastLineLen
  default:
    // at the beginning of the source
    pos.Line = 1
    pos.Column = 1
  }

  return
}

func (s *Scanner) Peek() Token {
  if s.Rune == EOF { return Token{ Type: EOF, Position: s.Pos() } }

  return Token{ Type: s.Rune, Text: s.Src[s.RunePos:s.SrcPos], Position: s.Pos() }
}

func (s *Scanner) Set2TokenPos(t Token) *Scanner {
  s.LastToken = t

  if t.Type == EOF || t.Position.Offset >= len( s.Src ) {
    s.SrcPos = len(s.Src)
    s.Rune, s.PrevRune, s.RunePos, s.PrevRunePos = EOF, EOF, s.SrcPos, s.SrcPos
    s.Line = ""
    s.column, s.line = t.Column, t.Line
    return s
  }

  s.RunePos = t.Position.Offset
  rune, w  := utf8.DecodeRuneInString( s.Src[s.RunePos:] )
  s.Rune, s.SrcPos = rune, s.RunePos + w
  rune, w   = utf8.DecodeLastRuneInString( s.Src[:s.RunePos] )
  s.PrevRune, s.PrevRunePos = rune, s.RunePos - w
  s.getLine = false

  if s.PrevRune == '\n' {
    if s.Rune == '\n' { s.getLine = true }

    s.Line           = getRawLine( s.Src[ s.RunePos: ] )
    s.lastLineLen    = utf8.RuneCountInString( getLastLine( s.PrevRunePos, s.Src ) )
    s.column, s.line = t.Column, t.Line
  } else if s.Rune == '\n' {
    s.getLine = true
    s.line, s.column = t.Line + 1, 0
    s.Line           = getLastLine( s.RunePos, s.Src )
    s.lastLineLen    = utf8.RuneCountInString( s.Line )
  } else {
    s.line, s.column = t.Line, t.Column
    s.lastLineLen    = t.Column - 1
    s.Line           = getLastLine( s.RunePos, s.Src )
  }

  return s
}

func getLastLine( i int, str string ) (string) {
  if i >= len( str ) || i < 0 { return "" }

  for i--; i >= 0; i-- {
    if str[i] == '\n' { return getRawLine( str[i+1:] ) }
  }

  return getRawLine( str )
}

func (s *Scanner) Copy() *Scanner {
  ns := new(Scanner)
  ns.Src          = s.Src
  ns.SrcPos       = s.SrcPos
  ns.Name         = s.Name
  ns.line         = s.line
  ns.column       = s.column
  ns.lastLineLen  = s.lastLineLen
  ns.Line         = s.Line
  ns.getLine      = s.getLine
  ns.Rune         = s.Rune
  ns.RunePos      = s.RunePos
  ns.PrevRune     = s.PrevRune
  ns.PrevRunePos  = s.PrevRunePos
  ns.CustomError  = s.CustomError
  ns.ErrorCount   = s.ErrorCount
  ns.Mode         = s.Mode
  ns.Whitespace   = s.Whitespace
  ns.LastToken    = s.LastToken

  return ns
}

func (s *Scanner) CopyStats( cs *Scanner ) *Scanner {
  s.Src          = cs.Src
  s.SrcPos       = cs.SrcPos
  s.Name         = cs.Name
  s.line         = cs.line
  s.column       = cs.column
  s.lastLineLen  = cs.lastLineLen
  s.Line         = cs.Line
  s.getLine      = cs.getLine
  s.Rune         = cs.Rune
  s.RunePos      = cs.RunePos
  s.PrevRune     = cs.PrevRune
  s.PrevRunePos  = cs.PrevRunePos
  s.CustomError  = cs.CustomError
  s.ErrorCount   = cs.ErrorCount
  s.Mode         = cs.Mode
  s.Whitespace   = cs.Whitespace
  s.LastToken    = cs.LastToken

  return s
}

func getRawLine( str string ) string {
  for i, c := range str {
    if c == '\n' {
      return str[:i + 1]
    }
  }

  return str
}

func (s *Scanner) isIdentRune(ch rune, i int) bool {
  return ch == '_' || unicode.IsLetter(ch) || unicode.IsDigit(ch) && i > 0
}

func (s *Scanner) ScanIdentifier() rune {
  // we know the zero'th rune is OK; start scanning at the next one
  ch := s.next()
  for i := 1; s.isIdentRune(ch, i); i++ {
    ch = s.next()
  }
  return ch
}

func digitVal(ch rune) int {
  switch {
  case '0' <= ch && ch <= '9':
    return int(ch - '0')
  case 'a' <= ch && ch <= 'f':
    return int(ch - 'a' + 10)
  case 'A' <= ch && ch <= 'F':
    return int(ch - 'A' + 10)
  }
  return 16 // larger than any legal digit val
}

func isDecimal(ch rune) bool { return '0' <= ch && ch <= '9' }

func (s *Scanner) scanMantissa(ch rune) rune {
  for isDecimal(ch) {
    ch = s.next()
  }
  return ch
}

func (s *Scanner) scanFraction(ch rune) rune {
  if ch == '.' {
    ch = s.scanMantissa(s.next())
  }
  return ch
}

func (s *Scanner) scanExponent(ch rune) rune {
  if ch == 'e' || ch == 'E' {
    ch = s.next()
    if ch == '-' || ch == '+' {
      ch = s.next()
    }
    ch = s.scanMantissa(ch)
  }
  return ch
}

func (s *Scanner) scanNumber(ch rune) rune {
  // isDecimal(ch)
  if ch == '0' {
    // int or float
    ch = s.next()
    if ch == 'x' || ch == 'X' {
      // hexadecimal int
      ch = s.next()
      hasMantissa := false
      for digitVal(ch) < 16 {
        ch = s.next()
        hasMantissa = true
      }
      if !hasMantissa {
        s.Error("illegal hexadecimal number")
      }
      return ScanHexadecimal
    } else {
      // octal int or float
      has8or9, theOne := false, true
      for isDecimal(ch) {
        if ch > '7' {
          has8or9 = true
        }
        ch = s.next()
        theOne = false
      }
      if s.Mode&ScanFloats != 0 && (ch == '.' || ch == 'e' || ch == 'E') {
        // float
        ch = s.scanFraction(ch)
        ch = s.scanExponent(ch)
        return ScanFloat
      } else if theOne {
        return ScanInt // lonley 0
      }

      // octal int
      if has8or9 {
        s.Error("illegal octal number")
        return ScanInt
      }

      return ScanOctal
    }
  }
  // decimal int or float
  ch = s.scanMantissa(ch)
  if s.Mode&ScanFloats != 0 && (ch == '.' || ch == 'e' || ch == 'E') {
    // float
    ch = s.scanFraction(ch)
    ch = s.scanExponent(ch)
    return ScanFloat
  }
  return ScanInt
}

func (s *Scanner) scanDigits(ch rune, base, n int) rune {
  for n > 0 && digitVal(ch) < base {
    ch = s.next()
    n--
  }
  if n > 0 {
    s.Error("illegal char escape")
  }
  return ch
}

func (s *Scanner) scanEscape(quote rune) rune {
  ch := s.next() // read character after '/'
  switch ch {
  case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', quote:
    // nothing to do
    ch = s.next()
  case '0', '1', '2', '3', '4', '5', '6', '7':
    ch = s.scanDigits(ch, 8, 3)
  case 'x':
    ch = s.scanDigits(s.next(), 16, 2)
  case 'u':
    ch = s.scanDigits(s.next(), 16, 4)
  case 'U':
    ch = s.scanDigits(s.next(), 16, 8)
  default:
    s.Error("illegal char escape")
  }
  return ch
}

func (s *Scanner) ScanString(quote rune) (n int) {
  ch := s.next() // read character after quote
  for ch != quote {
    if ch == '\n' || ch < 0 {
      s.Error("literal not terminated")
      return
    }
    if ch == '\\' {
      ch = s.scanEscape(quote)
    } else {
      ch = s.next()
    }
    n++
  }
  return
}

func (s *Scanner) ScanRawString() {
  ch := s.next() // read character after '`'
  for ch != '`' {
    if ch < 0 {
      s.Error("literal not terminated")
      return
    }
    ch = s.next()
  }
}

func (s *Scanner) ScanChar() {
  if s.ScanString('\'') != 1 {
    s.Error("illegal char literal")
  }
}

func (s *Scanner) ScanComment(ch rune) rune {
  // ch == '/' || ch == '*'
  if ch == '/' {
    // line comment
    ch = s.next() // read character after "//"
    for ch != '\n' && ch >= 0 {
      ch = s.next()
    }
    return ch
  }

  // general comment
  ch = s.next() // read character after "/*"
  for {
    if ch < 0 {
      s.Error("comment not terminated")
      break
    }
    ch0 := ch
    ch = s.next()
    if ch0 == '*' && ch == '/' {
      ch = s.next()
      break
    }
  }
  return ch
}

func (s *Scanner) Scan() Token {
redo:
  if s.Rune == EOF {
    s.LastToken = Token{ Type: EOF, Position: s.Pos() }
    return s.LastToken
  }

  // skip white space
  for s.Whitespace & (1 << uint(s.Rune) ) != 0 { s.next() }

  // start collecting token text
  s.LastToken.Text, s.LastToken.Type, s.LastToken.Position = "", s.Rune, s.Pos()
  tokPos    := s.RunePos

  switch {
  case s.isIdentRune(s.Rune, 0):
    if s.Mode&(ScanIdents|ScanBools) != 0 {
      s.LastToken.Type = ScanIdent
      s.ScanIdentifier()

      if s.Mode&ScanBools != 0 && (s.Src[tokPos] == 't' || s.Src[tokPos] == 'f') {
        switch s.Src[tokPos:s.RunePos] {
        case "true", "false":  s.LastToken.Type = ScanBool
        }
      }
    } else {
      s.next()
    }
  case isDecimal(s.Rune):
    if s.Mode&(ScanInts|ScanFloats) != 0 {
      s.LastToken.Type = s.scanNumber(s.Rune)
    } else {
      s.next()
    }
  default:
    switch s.Rune {
    case EOF:
    case '"':
      if s.Mode&ScanStrings != 0 {
        s.ScanString('"')
        s.LastToken.Type = ScanString
      }
      s.next()
    case '\'':
      if s.Mode&ScanChars != 0 {
        s.ScanChar()
        s.LastToken.Type = ScanChar
      }
      s.next()
    case '.':
      s.next()
      if isDecimal(s.Rune) && s.Mode&ScanFloats != 0 {
        s.LastToken.Type = ScanFloat
        s.scanMantissa(s.Rune)
        s.scanExponent(s.Rune)
      }
    case '-':
      s.next()
      if isDecimal( s.Rune ) && s.Mode&(ScanInts|ScanFloats) != 0 {
        s.LastToken.Type = s.scanNumber(s.Rune)
      }
    case '/':
      s.next()
      if (s.Rune == '/' || s.Rune == '*') && s.Mode&ScanComments != 0 {
        if s.Mode&SkipComments != 0 {
          s.Rune = s.ScanComment(s.Rune)
          goto redo
        }
        s.LastToken.Type = ScanComment
        s.ScanComment(s.Rune)
      }
    case '`':
      if s.Mode&ScanRawStrings != 0 {
        s.ScanRawString()
        s.LastToken.Type = ScanRawString
      }
      s.next()
    default:
      s.next()
    }
  }

  // token text
  switch s.LastToken.Type {
  case EOF:
  case ScanString:
    if tokPos != s.PrevRunePos && s.PrevRune == '"' {
      s.LastToken.Text = s.Src[tokPos + 1:s.PrevRunePos]
    } else {
      s.LastToken.Text = s.Src[tokPos + 1:s.RunePos]
    }
  case ScanRawString:
    if tokPos != s.PrevRunePos && s.PrevRune == '`' {
      s.LastToken.Text = s.Src[tokPos + 1:s.PrevRunePos]
    } else {
      s.LastToken.Text = s.Src[tokPos + 1:s.RunePos]
    }
  default: s.LastToken.Text = s.Src[tokPos:s.RunePos]
  }

  return s.LastToken
}
