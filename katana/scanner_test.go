 package katana

import (
  "bytes"
  "fmt"
  "testing"
  "unicode/utf8"
)

type token struct {
  tok  rune
  text string
}

var f100 = "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"

var tokenList = []token{
  {ScanComment, "// line comments"},
  {ScanComment, "//"},
  {ScanComment, "////"},
  {ScanComment, "// comment"},
  {ScanComment, "// /* comment */"},
  {ScanComment, "// // comment //"},
  {ScanComment, "//" + f100},

  {ScanComment, "// general comments"},
  {ScanComment, "/**/"},
  {ScanComment, "/***/"},
  {ScanComment, "/* comment */"},
  {ScanComment, "/* // comment */"},
  {ScanComment, "/* /* comment */"},
  {ScanComment, "/*\n comment\n*/"},
  {ScanComment, "/*" + f100 + "*/"},

  {ScanComment, "// identifiers"},
  {ScanIdent, "a"},
  {ScanIdent, "a0"},
  {ScanIdent, "foobar"},
  {ScanIdent, "abc123"},
  {ScanIdent, "LGTM"},
  {ScanIdent, "_"},
  {ScanIdent, "_abc123"},
  {ScanIdent, "abc123_"},
  {ScanIdent, "_abc_123_"},
  {ScanIdent, "_äöü"},
  {ScanIdent, "_本"},
  {ScanIdent, "äöü"},
  {ScanIdent, "本"},
  {ScanIdent, "a۰۱۸"},
  {ScanIdent, "foo६४"},
  {ScanIdent, "bar９８７６"},
  {ScanIdent, f100},

  {ScanComment, "// bools"},
  {ScanBool, "true"},
  {ScanBool, "false"},

  {ScanComment, "// decimal ints"},
  {ScanInt, "0"},
  {ScanInt, "1"},
  {ScanInt, "9"},
  {ScanInt, "42"},
  {ScanInt, "1234567890"},
  {ScanInt, "-0"},
  {ScanInt, "-1"},
  {ScanInt, "-9"},
  {ScanInt, "-42"},
  {ScanInt, "-1234567890"},

  {ScanComment, "// octal ints"},
  {ScanOctal, "00"},
  {ScanOctal, "01"},
  {ScanOctal, "07"},
  {ScanOctal, "042"},
  {ScanOctal, "01234567"},

  {ScanComment, "// hexadecimal ints"},
  {ScanHexadecimal, "0x0"},
  {ScanHexadecimal, "0x1"},
  {ScanHexadecimal, "0xf"},
  {ScanHexadecimal, "0x42"},
  {ScanHexadecimal, "0x123456789abcDEF"},
  {ScanHexadecimal, "0x" + f100},
  {ScanHexadecimal, "0X0"},
  {ScanHexadecimal, "0X1"},
  {ScanHexadecimal, "0XF"},
  {ScanHexadecimal, "0X42"},
  {ScanHexadecimal, "0X123456789abcDEF"},
  {ScanHexadecimal, "0X" + f100},

  {ScanComment, "// floats"},
  {ScanFloat, "0."},
  {ScanFloat, "1."},
  {ScanFloat, "42."},
  {ScanFloat, "01234567890."},
  {ScanFloat, ".0"},
  {ScanFloat, ".1"},
  {ScanFloat, ".42"},
  {ScanFloat, ".0123456789"},
  {ScanFloat, "0.0"},
  {ScanFloat, "1.0"},
  {ScanFloat, "42.0"},
  {ScanFloat, "01234567890.0"},
  {ScanFloat, "0e0"},
  {ScanFloat, "1e0"},
  {ScanFloat, "42e0"},
  {ScanFloat, "01234567890e0"},
  {ScanFloat, "0E0"},
  {ScanFloat, "1E0"},
  {ScanFloat, "42E0"},
  {ScanFloat, "01234567890E0"},
  {ScanFloat, "0e+10"},
  {ScanFloat, "1e-10"},
  {ScanFloat, "42e+10"},
  {ScanFloat, "01234567890e-10"},
  {ScanFloat, "0E+10"},
  {ScanFloat, "1E-10"},
  {ScanFloat, "42E+10"},
  {ScanFloat, "01234567890E-10"},
  // {Float, "-0."},
  {ScanFloat, "-1."},
  {ScanFloat, "-42."},
  {ScanFloat, "-01234567890."},
  // {Float, "-.0"},
  // {Float, "-.1"},
  // {Float, "-.42"},
  // {Float, "-.0123456789"},
  // {Float, "-0.0"},
  // {Float, "-1.0"},
  // {Float, "-42.0"},
  // {Float, "-01234567890.0"},
  // {Float, "-0e0"},
  // {Float, "-1e0"},
  // {Float, "-42e0"},
  // {Float, "-01234567890e0"},
  // {Float, "-0E0"},
  // {Float, "-1E0"},
  // {Float, "-42E0"},
  // {Float, "-01234567890E0"},
  // {Float, "-0e+10"},
  // {Float, "-1e-10"},
  // {Float, "-42e+10"},
  // {Float, "-01234567890e-10"},
  // {Float, "-0E+10"},
  // {Float, "-1E-10"},
  // {Float, "-42E+10"},
  // {Float, "-01234567890E-10"},

  {ScanComment, "// chars"},
  {ScanChar, `' '`},
  {ScanChar, `'a'`},
  {ScanChar, `'本'`},
  {ScanChar, `'\a'`},
  {ScanChar, `'\b'`},
  {ScanChar, `'\f'`},
  {ScanChar, `'\n'`},
  {ScanChar, `'\r'`},
  {ScanChar, `'\t'`},
  {ScanChar, `'\v'`},
  {ScanChar, `'\''`},
  {ScanChar, `'\000'`},
  {ScanChar, `'\777'`},
  {ScanChar, `'\x00'`},
  {ScanChar, `'\xff'`},
  {ScanChar, `'\u0000'`},
  {ScanChar, `'\ufA16'`},
  {ScanChar, `'\U00000000'`},
  {ScanChar, `'\U0000ffAB'`},

  {ScanComment, "// strings"},
  {ScanString, `" "`},
  {ScanString, `"a"`},
  {ScanString, `"本"`},
  {ScanString, `"\a"`},
  {ScanString, `"\b"`},
  {ScanString, `"\f"`},
  {ScanString, `"\n"`},
  {ScanString, `"\r"`},
  {ScanString, `"\t"`},
  {ScanString, `"\v"`},
  {ScanString, `"\""`},
  {ScanString, `"\000"`},
  {ScanString, `"\777"`},
  {ScanString, `"\x00"`},
  {ScanString, `"\xff"`},
  {ScanString, `"\u0000"`},
  {ScanString, `"\ufA16"`},
  {ScanString, `"\U00000000"`},
  {ScanString, `"\U0000ffAB"`},
  {ScanString, `"` + f100 + `"`},

  {ScanComment, "// raw strings"},
  {ScanRawString, "``"},
  {ScanRawString, "`\\`"},
  {ScanRawString, "`" + "\n\n/* foobar */\n\n" + "`"},
  {ScanRawString, "`" + f100 + "`"},

  {ScanComment, "// individual characters"},
  // NUL character is not allowed
  {'\x01', "\x01"},
  {' ' - 1, string(' ' - 1)},
  {'+', "+"},
  {'/', "/"},
  {'.', "."},
  {'~', "~"},
  {'(', "("},
}

func makeSource(pattern string) string {
  var buf bytes.Buffer
  for _, k := range tokenList {
    fmt.Fprintf(&buf, pattern, k.text)
  }
  return buf.String()
}

func checkTok(t *testing.T, s *Scanner, tok Token, line int, got, want rune, text string) {
  if got != want {
    t.Fatalf("tok = %s, want %s for %q", TokenString(got), TokenString(want), text)
  }
  if tok.Line != line {
    t.Errorf("line = %d, want %d for %q", tok.Line, line, text)
  }
  if tok.Type == ScanString || tok.Type == ScanRawString {
    text = rmQuotes( text, tok.Type )
  }

  if tok.Text != text {
    t.Errorf("text = %q, want %q", tok.Text, text)
  } else {
    // check idempotency of TokenText() call
    if tok.Text != text {
      t.Errorf("text = %q, want %q (idempotency check)", tok.Text, text)
    }
  }
}

func rmQuotes( text string, class rune ) string {
  key := byte( '"'  )
  if class == ScanRawString { key = '`' }

  switch len( text  ) {
  case 0: return ""
  case 1:
    if text[0] == key { return "" }
    return text
  default:
    end := len( text ) - 1
    if text[0] == key {
      if text[end] == key {
        return text[1:end]
      }
    }
    return text[1:]
  }

  return text
}


func countNewlines(s string) int {
  n := 0
  for _, ch := range s {
    if ch == '\n' {
      n++
    }
  }
  return n
}

func testScan(t *testing.T, mode uint) {
  s := NewScanner(makeSource(" \t%s\n")).Init()
  s.Mode = mode
  s.Scan()
  line := 1
  for _, k := range tokenList {
    if mode&SkipComments == 0 || k.tok != ScanComment {
      checkTok(t, s, s.LastToken, line, s.LastToken.Type, k.tok, k.text)
      s.Scan()
    }
    line += countNewlines(k.text) + 1 // each token is on a new line
  }

  checkTok(t, s, s.LastToken, line, s.LastToken.Type, EOF, "")
}

func TestScan(t *testing.T) {
  testScan(t, GoTokens)
  testScan(t, GoTokens&^SkipComments)
}

func TestPosition(t *testing.T) {
  src := makeSource("\t\t\t\t%s\n")
  s := NewScanner(src).Init()
  s.Mode = GoTokens &^ SkipComments
  s.Scan()
  pos := Position{"", 4, 1, 5}
  for _, k := range tokenList {
    if s.LastToken.Offset != pos.Offset {
      t.Errorf("offset = %d, want %d for %q", s.LastToken.Offset, pos.Offset, k.text)
    }
    if s.LastToken.Line != pos.Line {
      t.Errorf("line = %d, want %d for %q", s.LastToken.Line, pos.Line, k.text)
    }
    if s.LastToken.Column != pos.Column {
      t.Errorf("column = %d, want %d for %q", s.LastToken.Column, pos.Column, k.text)
    }
    pos.Offset += 4 + len(k.text) + 1     // 4 tabs + token bytes + newline
    pos.Line += countNewlines(k.text) + 1 // each token is on a new line
    s.Scan()
  }
  // make sure there were no token-internal errors reported by scanner
  if s.ErrorCount != 0 {
    t.Errorf("%d errors", s.ErrorCount)
  }
}

func TestScanZeroMode(t *testing.T) {
  src := makeSource("%s\n")
  str := src
  s := NewScanner(src).Init()
  s.Mode = 0       // don't recognize any token classes
  s.Whitespace = 0 // don't skip any whitespace
  s.Scan()
  for i, ch := range str {
    if s.LastToken.Type != ch {
      t.Fatalf("%d. tok = %s, want %s", i, TokenString(s.LastToken.Type), TokenString(ch))
    }
    s.Scan()
  }
  if s.LastToken.Type != EOF {
    t.Fatalf("tok = %s, want EOF", TokenString(s.LastToken.Type))
  }
  if s.ErrorCount != 0 {
    t.Errorf("%d errors", s.ErrorCount)
  }
}

func testScanSelectedMode(t *testing.T, mode uint, class rune) {
  src := makeSource("%s\n")
  s := NewScanner(src).Init()
  s.Mode = mode
  s.Scan()
  for s.LastToken.Type != EOF {
    if s.LastToken.Type < 0 && s.LastToken.Type != class {
      t.Fatalf("tok = %s, want %s", TokenString(s.LastToken.Type), TokenString(class))
    }
    s.Scan()
  }
  if s.ErrorCount != 0 {
    t.Errorf("%d errors", s.ErrorCount)
  }
}

func TestScanSelectedMask(t *testing.T) {
  testScanSelectedMode(t, 0, 0)
  testScanSelectedMode(t, ScanIdents, ScanIdent)
  // Don't test ScanInts and ScanNumbers since some parts of
  // the floats in the source look like (illegal) octal ints
  // and ScanNumbers may return either Int or Float.
  testScanSelectedMode(t, ScanChars, ScanChar)
  testScanSelectedMode(t, ScanStrings, ScanString)
  testScanSelectedMode(t, SkipComments, 0)
  testScanSelectedMode(t, ScanComments, ScanComment)
}

func TestScanNext(t *testing.T) {
  const BOM = '\uFEFF'
  BOMs := string(BOM)
  s := NewScanner( "if a == bcd /* com" + BOMs + "ment */ {\n\ta += c\n}" + BOMs + "// line comment ending in eof" ).Init()
  s.Scan(); checkTok(t, s, s.LastToken, 1, s.LastToken.Type, ScanIdent, "if") // the first BOM is ignored
  s.Scan(); checkTok(t, s, s.LastToken, 1, s.LastToken.Type, ScanIdent, "a")
  s.Scan(); checkTok(t, s, s.LastToken, 1, s.LastToken.Type, '=', "=")
  if ch := s.Next(); ch != '=' { t.Errorf( "%q != %q", ch, '=' ) }
  if ch := s.Next(); ch != ' ' { t.Errorf( "%q != %q", ch, ' ' ) }
  if ch := s.Next(); ch != 'b' { t.Errorf( "%q != %q", ch, 'b' ) }
  s.Scan(); checkTok(t, s, s.LastToken, 1, s.LastToken.Type, ScanIdent, "cd")
  s.Scan(); checkTok(t, s, s.LastToken, 1, s.LastToken.Type, '{', "{")
  s.Scan(); checkTok(t, s, s.LastToken, 2, s.LastToken.Type, ScanIdent, "a")
  s.Scan(); checkTok(t, s, s.LastToken, 2, s.LastToken.Type, '+', "+")
  if ch := s.Next(); ch != '=' { t.Errorf( "%q != %q", ch, '=' ) }
  s.Scan(); checkTok(t, s, s.LastToken, 2, s.LastToken.Type, ScanIdent, "c")
  s.Scan(); checkTok(t, s, s.LastToken, 3, s.LastToken.Type, '}', "}")
  s.Scan(); checkTok(t, s, s.LastToken, 3, s.LastToken.Type, BOM, BOMs)
  s.Scan(); checkTok(t, s, s.LastToken, 3, s.LastToken.Type, -1, "")
  if s.ErrorCount != 0 {
    t.Errorf("%d errors", s.ErrorCount)
  }
}

func TestScanWhitespace(t *testing.T) {
  var buf bytes.Buffer
  var ws uint64
  // start at 1, NUL character is not allowed
  for ch := byte(1); ch < ' '; ch++ {
    buf.WriteByte(ch)
    ws |= 1 << ch
  }
  const orig = 'x'
  buf.WriteByte(orig)

  s := NewScanner( buf.String() ).Init()
  s.Mode = 0
  s.Whitespace = ws
  s.Scan()
  if s.LastToken.Type != orig {
    t.Errorf("tok = %s, want %s", TokenString(s.LastToken.Type), TokenString(orig))
  }
}

func testError(t *testing.T, src, pos, msg string, tok rune) {
  s := NewScanner( src )
  errorCalled := false
  s.CustomError = func(s *Scanner, m string) {
    if !errorCalled {
      // only look at first error
      if p := s.Pos().String(); p != pos {
        t.Errorf("pos = %q, want %q for %q", p, pos, src)
      }
      if m != msg {
        t.Errorf("msg = %q, want %q for %q", m, msg, src)
      }
      errorCalled = true
    }
  }

  s.Init().Scan()
  if s.LastToken.Type != tok {
    t.Errorf("tok = %s, want %s for %q", TokenString(s.LastToken.Type), TokenString(tok), src)
  }
  if !errorCalled {
    t.Errorf("error handler not called for %q", src)
  }
  if s.ErrorCount == 0 {
    t.Errorf("count = %d, want > 0 for %q", s.ErrorCount, src)
  }
}

func TestError(t *testing.T) {
  testError(t, "\x00", "<input>:1:1", "illegal character NUL", 0)
  testError(t, "\x80", "<input>:1:1", "illegal UTF-8 encoding", utf8.RuneError)
  testError(t, "\xff", "<input>:1:1", "illegal UTF-8 encoding", utf8.RuneError)

  testError(t, "a\x00", "<input>:1:2", "illegal character NUL", ScanIdent)
  testError(t, "ab\x80", "<input>:1:3", "illegal UTF-8 encoding", ScanIdent)
  testError(t, "abc\xff", "<input>:1:4", "illegal UTF-8 encoding", ScanIdent)

  testError(t, `"a`+"\x00", "<input>:1:3", "illegal character NUL", ScanString)
  testError(t, `"ab`+"\x80", "<input>:1:4", "illegal UTF-8 encoding", ScanString)
  testError(t, `"abc`+"\xff", "<input>:1:5", "illegal UTF-8 encoding", ScanString)

  testError(t, "`a"+"\x00", "<input>:1:3", "illegal character NUL", ScanRawString)
  testError(t, "`ab"+"\x80", "<input>:1:4", "illegal UTF-8 encoding", ScanRawString)
  testError(t, "`abc"+"\xff", "<input>:1:5", "illegal UTF-8 encoding", ScanRawString)

  testError(t, `'\"'`, "<input>:1:3", "illegal char escape", ScanChar)
  testError(t, `"\'"`, "<input>:1:3", "illegal char escape", ScanString)

  testError(t, `01238`, "<input>:1:6", "illegal octal number", ScanInt)
  testError(t, `01238123`, "<input>:1:9", "illegal octal number", ScanInt)
  testError(t, `0x`, "<input>:1:3", "illegal hexadecimal number", ScanHexadecimal)
  testError(t, `0xg`, "<input>:1:3", "illegal hexadecimal number", ScanHexadecimal)
  testError(t, `'aa'`, "<input>:1:4", "illegal char literal", ScanChar)

  testError(t, `'`, "<input>:1:2", "literal not terminated", ScanChar)
  testError(t, `'`+"\n", "<input>:1:2", "literal not terminated", ScanChar)
  testError(t, `"abc`, "<input>:1:5", "literal not terminated", ScanString)
  testError(t, `"abc`+"\n", "<input>:1:5", "literal not terminated", ScanString)
  testError(t, "`abc\n", "<input>:2:1", "literal not terminated", ScanRawString)
  testError(t, `/*/`, "<input>:1:4", "comment not terminated", EOF)
}

func checkPos(t *testing.T, got, want Position) {
  if got.Offset != want.Offset || got.Line != want.Line || got.Column != want.Column {
    t.Errorf("got offset, line, column = %d, %d, %d; want %d, %d, %d",
      got.Offset, got.Line, got.Column, want.Offset, want.Line, want.Column)
  }
}

func checkNextPos(t *testing.T, s *Scanner, offset, line, column int, char rune) {
  if ch := s.Next(); ch != char {
    t.Errorf("ch = %s, want %s", TokenString(ch), TokenString(char))
  }
  want := Position{Offset: offset, Line: line, Column: column}
  checkPos(t, s.Pos(), want)
}

func checkScanPos(t *testing.T, s *Scanner, offset, line, column int, char rune) {
  want := Position{Offset: offset, Line: line, Column: column}
  checkPos(t, s.Pos(), want)
  s.Scan()
  if s.LastToken.Type != char {
    t.Errorf("ch = %s, want %s", TokenString(s.LastToken.Type), TokenString(char))
  }
  checkPos(t, s.LastToken.Position, want)
}

func TestPos(t *testing.T) {
  // corner case: empty source
  s := NewScanner("").Init()
  checkPos(t, s.Pos(), Position{Offset: 0, Line: 1, Column: 1})
  checkPos(t, s.Pos(), Position{Offset: 0, Line: 1, Column: 1})

  // corner case: source with only a newline
  s = NewScanner("\n").Init()
  checkPos(t, s.Pos(), Position{Offset: 0, Line: 1, Column: 1})
  checkNextPos(t, s, 1, 2, 1, '\n')
  // after EOF position doesn't change
  for i := 10; i > 0; i-- {
    checkScanPos(t, s, 1, 2, 1, EOF)
  }
  if s.ErrorCount != 0 {
    t.Errorf("%d errors", s.ErrorCount)
  }

  // corner case: source with only a single character
  s = NewScanner("本").Init()
  checkPos(t, s.Pos(), Position{Offset: 0, Line: 1, Column: 1})
  checkNextPos(t, s, 3, 1, 2, '本')
  // after EOF position doesn't change
  for i := 10; i > 0; i-- {
    checkScanPos(t, s, 3, 1, 2, EOF)
  }
  if s.ErrorCount != 0 {
    t.Errorf("%d errors", s.ErrorCount)
  }

  // positions after calling Next
  s = NewScanner("  foo६४  \n\n本語\n").Init()
  checkNextPos(t, s, 1, 1, 2, ' ')
  checkNextPos(t, s, 2, 1, 3, ' ')
  checkNextPos(t, s, 3, 1, 4, 'f')
  checkNextPos(t, s, 4, 1, 5, 'o')
  checkNextPos(t, s, 5, 1, 6, 'o')
  checkNextPos(t, s, 8, 1, 7, '६')
  checkNextPos(t, s, 11, 1, 8, '४')
  checkNextPos(t, s, 12, 1, 9, ' ')
  checkNextPos(t, s, 13, 1, 10, ' ')
  checkNextPos(t, s, 14, 2, 1, '\n')
  checkNextPos(t, s, 15, 3, 1, '\n')
  checkNextPos(t, s, 18, 3, 2, '本')
  checkNextPos(t, s, 21, 3, 3, '語')
  checkNextPos(t, s, 22, 4, 1, '\n')
  // after EOF position doesn't change
  for i := 10; i > 0; i-- {
    checkScanPos(t, s, 22, 4, 1, EOF)
  }
  if s.ErrorCount != 0 {
    t.Errorf("%d errors", s.ErrorCount)
  }

  // positions after calling Scan
  s = NewScanner("abc\n本語\n\nx").Init()
  s.Mode = 0
  s.Whitespace = 0
  checkScanPos(t, s, 0, 1, 1, 'a')
  checkScanPos(t, s, 1, 1, 2, 'b')
  checkScanPos(t, s, 2, 1, 3, 'c')
  checkScanPos(t, s, 3, 1, 4, '\n')
  checkScanPos(t, s, 4, 2, 1, '本')
  checkScanPos(t, s, 7, 2, 2, '語')
  checkScanPos(t, s, 10, 2, 3, '\n')
  checkScanPos(t, s, 11, 3, 1, '\n')
  checkScanPos(t, s, 12, 4, 1, 'x')
  // after EOF position doesn't change
  for i := 10; i > 0; i-- {
    checkScanPos(t, s, 13, 4, 2, EOF)
  }
  if s.ErrorCount != 0 {
    t.Errorf("%d errors", s.ErrorCount)
  }
}

func TestGetLine(t *testing.T){
  data := []struct{
    input    string
    output   string
  } {
    { "", "" },
    { "\n\n\n", "\n" },
    { "line", "line" },
    { "line\n", "line\n" },
    { "line\t\v\rline", "line\t\v\rline" },
    { "line\t\v\nline\n", "line\t\v\n" },
    { "line\t\v\n\nline\n\n", "line\t\v\n" },
    { "line\t\v\n\n\tline\n\n\n", "line\t\v\n" },
    { "\n1\n2\n3\n", "\n" },
    { "1\n2\n3\n4\n5\n", "1\n" },
  }

  for _, d := range data {
    scan := NewScanner( d.input ).Init()

    if scan.Line != d.output {
      t.Errorf( "TestGetLine( %q ) \nreturn   %q\nexpected %q", d.input, scan.Line, d.output )
    }
  }

  data2 := []struct{
    input    string
    output   string
  } {
    { "", "" },
    { "               \n\n\n", "               \n" },
    { "               line", "               line" },
    { "               line\n", "               line\n" },
    { "               line\t\v\rline", "               line\t\v\rline" },
    { "               line\t\v\nline\n", "               line\t\v\n" },
    { "               line\t\v\n\nline\n\n", "               line\t\v\n" },
    { "               line\t\v\n\n\tline\n\n\n", "               line\t\v\n" },
    { "               \n1\n2\n3\n", "               \n" },
    { "               1\n2\n3\n4\n5\n", "               1\n" },
  }

  for i, d := range data2 {
    scan := NewScanner( d.input ).Init()

    for x := 0; x < i; x++ { scan.Next() }

    if scan.Line != d.output {
      t.Errorf( "TestGetLine( %q ) \nreturn   %q\nexpected %q", d.input, scan.Line, d.output )
    }
  }

  data3 := []struct{
    input    string
    output   string
  } {
    { "hela aleh", "hela aleh" },
    { "\n\nhela", "hela" },
    { "\n\n\n\n\n", "\n" },
    { "\n\nline", "line" },
    { "\n\nline\n", "line\n" },
    { "\n\nline\t\v\rline", "line\t\v\rline" },
    { "\n\nline\t\v\nline\n", "line\t\v\n" },
    { "\n\nline\t\v\n\nline\n\n", "line\t\v\n" },
    { "\n\nline\t\v\n\n\tline\n\n\n", "line\t\v\n" },
    { "\n\n\n1\n2\n3\n", "\n" },
    { "\n\n1\n2\n3\n4\n5\n", "1\n" },
    { "", "" },
  }

  for _, d := range data3 {
    scan := NewScanner( d.input ).Init()

    for x := 0; x < 2; x++ { scan.Next() }

    if scan.Line != d.output {
      t.Errorf( "TestGetLine( %q ) \nreturn   %q\nexpected %q\n    peek %q [%d]", d.input, scan.Line, d.output, scan.Rune, scan.SrcPos )
    }
  }
}

func TestNextLine(t *testing.T){
  data := []struct{
    input    string
    output   []string
  } {
    { "", []string{} },
    { "\n", []string{ "\n" } },
    { "\n\n\n", []string{ "\n", "\n", "\n" } },
    { "\n\ta", []string{ "\n", "\ta" } },
    { "line", []string{ "line" } },
    { "line\t\v\rlinr", []string{ "line\t\v\rlinr" } },
    { "line\t\v\r\nlinr\n", []string{ "line\t\v\r\n", "linr\n" } },
    { "line\t\v\r\nlinr\n\n", []string{ "line\t\v\r\n", "linr\n", "\n", } },
    { "line\t\v\r\nlinr\n\n\n2.6", []string{ "line\t\v\r\n", "linr\n", "\n", "\n", "2.6" } },
    { "\n1\n2\n3\n", []string{ "\n", "1\n", "2\n", "3\n" } },
    { "1\n2\n3\n4\n5\n", []string{ "1\n", "2\n", "3\n", "4\n", "5\n" } },
  }

  for _, d := range data {
    scan := NewScanner( d.input ).Init()

    for i, line := range d.output {
      if scan.Line !=  line {
        t.Errorf( "TestNextLine( %q ) [%d]\nreturn   %q\nexpected %q\n    peek %q [%d]", d.input, i, scan.Line, line, scan.Rune, scan.SrcPos )
      }

      scan.NextLine()
    }
  }
}

func TestScanAndNextLine(t *testing.T){
  scan := NewScanner( "\nhey line\nnext liner\n\nbig opai\nexit" ).Init()

  scan.Scan()
  expected := "hey"
  if scan.LastToken.Text != expected { t.Errorf( "TestScanAndNextLine()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}

  expected = " line\n"
  if scan.RestOfLine() != expected { t.Errorf( "TestScanAndNextLine()\nreturn   %q\nexpected %q\n", scan.RestOfLine(), expected )}


  scan.NextLine()
  expected = "next liner\n"
  if scan.Line != expected { t.Errorf( "TestScanAndNextLine()\nreturn   %q\nexpected %q\n", scan.Line, expected )}

  expected = "next liner\n"
  if scan.RestOfLine() != expected { t.Errorf( "TestScanAndNextLine()\nreturn   %q\nexpected %q\n", scan.RestOfLine(), expected )}

  scan.Scan()
  expected = "next"
  if scan.LastToken.Text != expected { t.Errorf( "TestScanAndNextLine()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}

  scan.Scan()
  expected = "liner"
  if scan.LastToken.Text != expected { t.Errorf( "TestScanAndNextLine()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}

  expected = "next liner\n"
  if scan.Line != expected { t.Errorf( "TestScanAndNextLine()\nreturn   %q\nexpected %q\n", scan.Line, expected )}

  expected = "\n"
  if scan.RestOfLine() != expected { t.Errorf( "TestScanAndNextLine()\nreturn   %q\nexpected %q\n", scan.RestOfLine(), expected )}

  scan.Next()
  if scan.PrevRune != '\n' || scan.Rune != '\n' { t.Errorf( "TestScanAndNextLine: fail Runes\n" )}

  expected = "\n"
  if scan.RestOfLine() != expected { t.Errorf( "TestScanAndNextLine()\nreturn   %q\nexpected %q\n", scan.RestOfLine(), expected )}

  scan.NextLine()
  expected = "big opai\n"
  if scan.Line != expected { t.Errorf( "TestScanAndNextLine()\nreturn   %q\nexpected %q\n", scan.Line, expected )}

  expected = "big opai\n"
  if scan.RestOfLine() != expected { t.Errorf( "TestScanAndNextLine()\nreturn   %q\nexpected %q\n", scan.RestOfLine(), expected )}

  scan.NextLine()
  expected = "exit"
  if scan.Line != expected { t.Errorf( "TestScanAndNextLine()\nreturn   %q\nexpected %q\n", scan.Line, expected )}

  scan.Scan()
  expected = "exit"
  if scan.LastToken.Text != expected { t.Errorf( "TestScanAndNextLine()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}

  if scan.PrevRune != 't' || scan.Rune != EOF { t.Errorf( "TestScanAndNextLine: fail Runes %q %d\n", scan.PrevRune, scan.Rune )}

  scan.Scan()
  expected = ""
  if scan.LastToken.Text != expected { t.Errorf( "TestScanAndNextLine()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}

  if scan.PrevRune != 't' || scan.Rune != EOF { t.Errorf( "TestScanAndNextLine: fail Runes %q %d\n", scan.PrevRune, scan.Rune )}

  expected = ""
  if scan.Line != expected { t.Errorf( "TestScanAndNextLine()\nreturn   %q\nexpected %q\n", scan.Line, expected )}

  if scan.PrevRune != 't' || scan.Rune != EOF { t.Errorf( "TestScanAndNextLine: fail Runes %q %d\n", scan.PrevRune, scan.Rune )}

  scan.Next()
  expected = ""
  if scan.Line != expected { t.Errorf( "TestScanAndNextLine()\nreturn   %q\nexpected %q\n", scan.Line, expected )}

  if scan.PrevRune != EOF || scan.Rune != EOF { t.Errorf( "TestScanAndNextLine: fail Runes %q %d\n", scan.PrevRune, scan.Rune )}

  scan.NextLine()
  expected = ""
  if scan.Line != expected { t.Errorf( "TestScanAndNextLine()\nreturn   %q\nexpected %q\n", scan.Line, expected )}
}

func TestNinja(t *testing.T){
  scan := NewScanner( "\n--y \"text\" line\nnext liner\nuooo\nbig opai\nexit" ).Init()

  // scan.NinjaLenMoves( 2 )

  scan.Scan()
  expected := "-"
  if scan.LastToken.Text != expected { t.Errorf( "TestNinja()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}
  scan.Scan()
  expected = "-"
  if scan.LastToken.Text != expected { t.Errorf( "TestNinja()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}
  scan.Scan()
  expected = "y"
  if scan.LastToken.Text != expected { t.Errorf( "TestNinja()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}

  scan.Scan()
  expected = "text"
  if scan.LastToken.Text != expected { t.Errorf( "TestNinja()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}

  expected = " line\n"
  if scan.RestOfLine() != expected { t.Errorf( "TestNinja()\nreturn   %q\nexpected %q\n", scan.RestOfLine(), expected )}

  scan.NextLine()
  expected = "next liner\n"
  if scan.Line != expected { t.Errorf( "TestNinja()\nreturn   %q\nexpected %q\n", scan.Line, expected )}

  expected = "next liner\n"
  if scan.RestOfLine() != expected { t.Errorf( "TestNinja()\nreturn   %q\nexpected %q\n", scan.RestOfLine(), expected )}

  scan.Scan()
  expected = "next"
  if scan.LastToken.Text != expected { t.Errorf( "TestNinja()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}

  scan.NinjaLenMoves( len( scan.RestOfLine() ) )
  expected = "uooo\n"
  if scan.Line != expected { t.Errorf( "TestNinja()\nreturn   %q\nexpected %q\n", scan.Line, expected )}

  scan.Next()
  if scan.PrevRune != 'u' || scan.Rune != 'o' { t.Errorf( "TestNinja: fail Runes\n" )}

  scan.Next()
  if scan.PrevRune != 'o' || scan.Rune != 'o' { t.Errorf( "TestNinja: fail Runes\n" )}

  scan.NinjaLenMoves( len( scan.RestOfLine() ) )
  expected = "big opai\n"
  if scan.Line != expected { t.Errorf( "TestNinja()\nreturn   %q\nexpected %q\n", scan.Line, expected )}

  scan.Scan()
  expected = "big"
  if scan.LastToken.Text != expected { t.Errorf( "TestNinja()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}

  expected = " opai\n"
  if scan.RestOfLine() != expected { t.Errorf( "TestNinja()\nreturn   %q\nexpected %q\n", scan.RestOfLine(), expected )}

  scan.NinjaLenMoves( len( scan.RestOfLine() ) )
  expected = "exit"
  if scan.Line != expected { t.Errorf( "TestNinja()\nreturn   %q\nexpected %q\n", scan.Line, expected )}

  scan.NinjaLenMoves( 1000 )
  expected = ""
  if scan.Line != expected { t.Errorf( "TestScanAndNextLine()\nreturn   %q\nexpected %q\n", scan.Line, expected )}
}

func TestSet2TokenPos(t *testing.T){
  scan := NewScanner( "\nhey line\nnext liner\nuooo\nbig opai\nexit" ).Init()

  scan.NinjaLenMoves( 2 )

  copy := scan.Copy()
  tok  := scan.Scan()
  expected := "ey"
  if scan.LastToken.Text != expected { t.Errorf( "TestSet2TokenPos()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}

  scan.NextLine()
  expected = "next liner\n"
  if scan.Line != expected { t.Errorf( "TestSet2TokenPos()\nreturn   %q\nexpected %q\n", scan.Line, expected )}

  scan.Set2TokenPos( tok )
  checkScanners( t, scan, copy )

  scan.NextLine()
  expected = "next liner\n"
  if scan.Line != expected { t.Errorf( "TestSet2TokenPos()\nreturn   %q\nexpected %q\n", scan.Line, expected )}

  expected = "next liner\n"
  if scan.RestOfLine() != expected { t.Errorf( "TestSet2TokenPos()\nreturn   %q\nexpected %q\n", scan.RestOfLine(), expected )}

  copy = scan.Copy()
  tok  = scan.Scan()
  expected = "next"
  if scan.LastToken.Text != expected { t.Errorf( "TestSet2TokenPos()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}

  scan.NinjaLenMoves( len( scan.RestOfLine() ) )
  expected = "uooo\n"
  if scan.Line != expected { t.Errorf( "TestSet2TokenPos()\nreturn   %q\nexpected %q\n", scan.Line, expected )}

  scan.Set2TokenPos( tok )
  checkScanners( t, scan, copy )

  scan.Scan()
  expected = "next"
  if scan.LastToken.Text != expected { t.Errorf( "TestSet2TokenPos()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}

  scan.Mode = 0
  scan.Scan()
  expected = "l"
  if scan.LastToken.Text != expected { t.Errorf( "TestSet2TokenPos()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}
  scan.Scan()
  expected = "i"
  if scan.LastToken.Text != expected { t.Errorf( "TestSet2TokenPos()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}
  scan.Scan()
  expected = "n"
  if scan.LastToken.Text != expected { t.Errorf( "TestSet2TokenPos()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}
  scan.Scan()
  expected = "e"
  if scan.LastToken.Text != expected { t.Errorf( "TestSet2TokenPos()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}
  scan.Next()

  copyN := scan.Copy()
  tokN := scan.Peek()
  expected = "\n"
  if tokN.Text != expected { t.Errorf( "TestSet2TokenPos()\nreturn   %q\nexpected %q\n", tokN.Text, expected )}

  scan.Mode = GoTokens
  scan.Scan()
  expected = "uooo"
  if scan.LastToken.Text != expected { t.Errorf( "TestSet2TokenPos()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}

  scan.Set2TokenPos( tok )
  checkScanners( t, scan, copy )

  scan.Set2TokenPos( tokN )
  checkScanners( t, scan, copyN )

  scan.NextLine()
  expected = "uooo\n"
  if scan.RestOfLine() != expected { t.Errorf( "TestSet2TokenPos()\nreturn   %q\nexpected %q\n", scan.RestOfLine(), expected )}

  scan.NextLine()
  expected = "big opai\n"
  if scan.RestOfLine() != expected { t.Errorf( "TestSet2TokenPos()\nreturn   %q\nexpected %q\n", scan.RestOfLine(), expected )}

  scan.NextLine()
  expected = "exit"
  if scan.RestOfLine() != expected { t.Errorf( "TestSet2TokenPos()\nreturn   %q\nexpected %q\n", scan.RestOfLine(), expected )}

  scan.NextLine()
  expected = ""
  if scan.RestOfLine() != expected { t.Errorf( "TestSet2TokenPos()\nreturn   %q\nexpected %q\n", scan.RestOfLine(), expected )}

  // EOF

  copy = scan.Copy()
  tok  = scan.Scan()

  expected = ""
  if scan.LastToken.Text != expected { t.Errorf( "TestSet2TokenPos()\nreturn   %q\nexpected %q\n", scan.LastToken.Text, expected )}

  scan.NinjaLenMoves( len( scan.RestOfLine() ) )
  expected = ""
  if scan.Line != expected { t.Errorf( "TestSet2TokenPos()\nreturn   %q\nexpected %q\n", scan.Line, expected )}

  scan.Set2TokenPos( tok )
  checkScanners( t, scan, copy )
}

func checkScanners( t *testing.T, a, b *Scanner ){
  e := new( bytes.Buffer )
  if a.Name != b.Name {
    fmt.Fprintf( e, "Name    ouput %q\n     expected %q\n", a.Name, b.Name )
  }
  if a.Src != b.Src {
    fmt.Fprintf( e, "Src     ouput %q\n     expected %q\n", a.Src, b.Src )
  }
  if a.SrcPos != b.SrcPos {
    fmt.Fprintf( e, "SrcPos  ouput %d\n     expected %d\n", a.SrcPos, b.SrcPos )
  }
  if a.line != b.line {
    fmt.Fprintf( e, "line    ouput %d\n     expected %d\n", a.line, b.line )
  }
  if a.column != b.column {
    fmt.Fprintf( e, "column  ouput %d\n     expected %d\n", a.column, b.column )
  }
  if a.lastLineLen != b.lastLineLen {
    fmt.Fprintf( e, "lstLLen ouput %d\n     expected %d\n", a.lastLineLen, b.lastLineLen )
  }
  if a.Line != b.Line {
    fmt.Fprintf( e, "Line    ouput %q\n     expected %q\n", a.Line, b.Line )
  }
  if a.getLine != b.getLine {
    fmt.Fprintf( e, "getLine ouput %t\n     expected %t\n", a.getLine, b.getLine )
  }
  if a.Rune != b.Rune {
    fmt.Fprintf( e, "Rune    ouput %d\n     expected %d\n", a.Rune, b.Rune )
  }
  if a.RunePos != b.RunePos {
    fmt.Fprintf( e, "RunePos ouput %d\n     expected %d\n", a.RunePos, b.RunePos )
  }
  if a.PrevRune != b.PrevRune {
    fmt.Fprintf( e, "PRune   ouput %d\n     expected %d\n", a.PrevRune, b.PrevRune )
  }
  if a.PrevRunePos != b.PrevRunePos {
    fmt.Fprintf( e, "PRPos   ouput %d\n     expected %d\n", a.PrevRunePos, b.PrevRunePos )
  }

  if e.Len() > 0 {
    t.Errorf( "checkScanners()\n%s", e )
  }
}

func TestLimits( t *testing.T ){
  s := NewScanner( "hey\nlisten!" ).Init()
  s.NextLine()

  expected := "listen!"
  if s.Line != expected {
    t.Errorf( "Line    ouput %q\n     expected %q\n", s.Line, expected )
  }

  expected = "hey\n"
  if s.Src[:s.RunePos] != expected {
    t.Errorf( "Line    ouput %q\n     expected %q\n", s.Src[:s.RunePos], expected )
  }

  expected = "listen!"
  if s.Text() != expected {
    t.Errorf( "Line    ouput %q\n     expected %q\n", s.Text(), expected )
  }

  s.NextLine()
  expected = ""
  if s.Line != expected {
    t.Errorf( "Line    ouput %q\n     expected %q\n", s.Line, expected )
  }

  expected = "hey\nlisten!"
  if s.Src[:s.RunePos] != expected {
    t.Errorf( "Line    ouput %q\n     expected %q\n", s.Src[:s.RunePos], expected )
  }

  expected = ""
  if s.Text() != expected {
    t.Errorf( "Line    ouput %q\n     expected %q\n", s.Text(), expected )
  }

  s.NextLine()
  expected = ""
  if s.Line != expected {
    t.Errorf( "Line    ouput %q\n     expected %q\n", s.Line, expected )
  }

  expected = "hey\nlisten!"
  if s.Src[:s.RunePos] != expected {
    t.Errorf( "Line    ouput %q\n     expected %q\n", s.Src[:s.RunePos], expected )
  }

  expected = ""
  if s.Text() != expected {
    t.Errorf( "Line    ouput %q\n     expected %q\n", s.Text(), expected )
  }

  c := s.Copy()

  c.NextLine()
  expected = ""
  if c.Line != expected {
    t.Errorf( "Line    ouput %q\n     expected %q\n", c.Line, expected )
  }

  expected = "hey\nlisten!"
  if c.Src[:c.RunePos] != expected {
    t.Errorf( "Line    ouput %q\n     expected %q\n", c.Src[:c.RunePos], expected )
  }

  expected = ""
  if c.Text() != expected {
    t.Errorf( "Line    ouput %q\n     expected %q\n", c.Text(), expected )
  }


}
