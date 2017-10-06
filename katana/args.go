package katana

import "fmt"

const (
  Empty = iota
	Ident
  Bool
	Int
  Octal
  Hexadecimal
	Float
	Char
	String
	RawString
	Comment
)

const comma = Comment + 1

var typeArgToString = map[ byte ]string{
  Empty       : "Empty",
	Ident       : "Ident",
  Bool        : "Bool",
	Int         : "Int",
  Octal       : "Octal",
  Hexadecimal : "Hexadecimal",
	Float       : "Float",
	Char        : "Char",
	String      : "String",
	RawString   : "RawString",
	Comment     : "Comment",
}

func TypeToString( at byte  ) string {
  if s, found := typeArgToString[ at ]; found { return s }

  return "unknow type"
}

type Arg struct {
  Data string
  Type byte
}

type ArgType struct {
  Name string
  Args []Arg
}

type ArgMap map[string][]Arg

func (scan *Scanner) GetArgs() (args []ArgType) {
  for scan.Rune != EOF {
    if arg, _ := scan.GetArgType(); arg != nil {
      args = append( args, *arg )
    }
  }

  return
}

func (s *Scanner) GetArgType() (*ArgType, string) {
  name, closing := "", rune( 0 )
  switch s.Scan(); s.LastToken.Type {
  case EOF : return nil, "GetArg: EOF"
  case '\n': return nil, `GetArg: \n`
  default:
    return nil, `GetArg: "` + s.LastToken.Text + `" no is a valid Argument Name`
  case ScanIdent:
    name = s.LastToken.Text

    switch s.Scan(); s.LastToken.Type {
    case '(': closing = ')'
    case '[': closing = ']'
    case '{': closing = '}'
    case '<': closing = '>'
    case EOF, '\n': return &ArgType{ Name: name }, ""
    default:
      s.Set2TokenPos( s.LastToken )
      return &ArgType{ Name: name }, ""
    }
  }

  vargs, ferr := s.getFarg( name, s.LastToken.Text, closing )
  if ferr != "" {
    return &ArgType{ Name: name, Args: vargs }, ""
  }

  return &ArgType{ Name: name, Args: vargs }, "GetArg:" + ferr
}

func (s *Scanner) getFarg( name, opening string, closing rune ) (farg []Arg, err string) {
  s.Scan()
  for lastType := byte(Empty); ; s.Scan() {
    switch s.LastToken.Type {
    case closing:
      if lastType == comma { farg = append( farg, Arg{ Data: "", Type: Empty } ) }
      return
    case ScanIdent, ScanBool, ScanInt, ScanOctal, ScanHexadecimal, ScanFloat, ScanChar, ScanString, ScanRawString, ScanComment:
      lastType = typeConverter( s.LastToken.Type )
      farg = append( farg, Arg{ Data: s.LastToken.Text, Type: lastType } )
    case ',':
      if lastType == comma || lastType == Empty {
        farg = append( farg, Arg{ Data: "", Type: Empty } )
      }
      lastType = comma
    case EOF : return farg, "getFargs: EOF"
    case '\n': return farg, "getFargs: '\\n'"
    default:
      switch s.LastToken.Type {
      case ')', ']', '}', '>': err = `getFargs: close "` + name + opening + `" with "` + s.LastToken.Text + `"`
      default: err = "getFargs: unknow element in argument"
      }

      return
    }
  }
}

func (d *Scanner) syncOpt( opt ArgType, m map[string][]Arg ) (r ArgType, cannon bool) {
  def, ok := m[ opt.Name ]
  if !ok { return opt, false }

  if len( opt.Args ) > len( def ) {
    d.Error( `syncOpt(): Arg "` + opt.Name + `" exceeds the number of elements` )
    opt.Args = opt.Args[:len( def )]
  } else if len( opt.Args ) < len( def ) {
    tmp := make( []Arg, len( def ) )
    copy( tmp, opt.Args )
    opt.Args = tmp
  }

  for i, arg := range opt.Args {
    if arg.Type != def[i].Type {
      if arg.Type != Empty {
        d.Error( fmt.Sprintf( "syncOpt(): Arg \"%s.Args[%d]\" type %q, expected %q", opt.Name, i, TypeToString( arg.Type ), TypeToString( def[i].Type ) ) )
      }

      opt.Args[i] = def[i]
    }
  }

  return opt, true
}

func typeConverter( r rune ) byte {
  switch r {
	case ScanIdent        : return Ident
  case ScanBool         : return Bool
	case ScanInt          : return Int
  case ScanOctal        : return Octal
  case ScanHexadecimal  : return Hexadecimal
	case ScanFloat        : return Float
	case ScanChar         : return Char
	case ScanString       : return String
	case ScanRawString    : return RawString
	case ScanComment      : return Comment
  }

  return Empty
}
