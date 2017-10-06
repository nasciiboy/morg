package katana

import (
  "fmt"
  "strings"

  "github.com/nasciiboy/txt"
  "github.com/nasciiboy/regexp4"
)

type BlockType int

const (
  BlockDefault BlockType = iota
  BlockConf
  BlockArgs
  BlockArgsBody
)

type MiniBlock struct {
  // Head *Scanner
  Body *Scanner
}

type Block struct {
  Comm   Token
  Args   ArgMap
  Head   Scanner
  Body   Scanner
  Type   BlockType
  Indent int
  Attach []MiniBlock
}

var BlockHat = map[string]BlockType {
  "title"       : BlockConf,
  "subtitle"    : BlockConf,
  "author"      : BlockConf,
  "style"       : BlockConf,
  "translator"  : BlockConf,
  "lang"        : BlockConf,
  "language"    : BlockConf,
  "licence"     : BlockConf,
  "date"        : BlockConf,
  "tags"        : BlockConf,
  "mail"        : BlockConf,
  "description" : BlockConf,
  "id"          : BlockConf,
  "options"     : BlockConf,
  "source"      : BlockConf,
  "cover"       : BlockConf,

  "figure"      : BlockArgsBody,

  "media"       : BlockDefault,
  "img"         : BlockDefault,
  "video"       : BlockDefault,
  "audio"       : BlockDefault,

  "ignore"      : BlockDefault,
  "center"      : BlockDefault,
  "bold"        : BlockDefault,
  "emph"        : BlockDefault,
  "verse"       : BlockDefault,
  "tab"         : BlockDefault,
  "italic"      : BlockDefault,
  "cols"        : BlockDefault,
  "code"        : BlockDefault,
  "src"         : BlockDefault,
  "srci"        : BlockDefault,
  "example"     : BlockDefault,
  "pre"         : BlockDefault,
  "pret"        : BlockDefault,
  "math"        : BlockDefault,
  "diagram"     : BlockDefault,
  "art"         : BlockDefault,
  "quote"       : BlockDefault,
}

func (s *Scanner) GetBlock() *Block {
  s.Scan()
  if s.PrevRune != '.' || s.Rune != '.' {
    s.Error( `GetBlock: "` + s.Line + `" wrong input` )
    return nil
  }
  s.Next()

  block := Block{ Comm: s.Scan(), Indent: txt.CountInitSpaces( s.Line ) }
  if block.Comm.Type != ScanIdent {
    s.Error( `GetBlock: "` + block.Comm.Text + `" is not a valid identifier` )
    return nil
  }
  block.Comm.Text = strings.ToLower( block.Comm.Text )
  block.Type      = BlockHat[ block.Comm.Text ]
  if block.Type == BlockConf && block.Indent > 0 {
    s.Error( "GetBlock: config block, indent > 0" )
  }

  s.Whitespace = catchNewLineInScan
  args, fatal := s.getBlockArgs()
  s.Whitespace = GoWhitespace
  if fatal != "" {
    s.Error( "GetBlock:" + fatal )
    return nil
  }

  block.Args, block.Head = s.toArgMap( args ), *s.Copy()
  block.Indent += 2

  s.NextLine()
  if block.Type == BlockDefault {
    block.Head.Src = s.Src[ :s.RunePos ]
  }

  switch block.Type {
  case BlockArgs, BlockConf:
    _, width := txt.DragTextByIndent( s.Src[ s.RunePos: ], block.Indent )
    s.NinjaLenMoves( width )
    block.Head.Src = s.Src[ :s.RunePos ]
  case BlockArgsBody:
    _, width := txt.DragTextByIndent( s.Src[ s.RunePos: ], block.Indent )
    s.NinjaLenMoves( width )
    block.Head.Src = s.Src[ :s.RunePos ]

    fallthrough
  default:
    rech := regexp4.Compile( fmt.Sprintf( "#^:b{%d}:<:>", block.Indent - 2 ) )
    rebo := regexp4.Compile( fmt.Sprintf( "#^:b{%d}:<:b*%s#*:.:.", block.Indent - 2, block.Comm.Text ) )
    ptrBody := &block.Body
  redo:
    *ptrBody = *s.Copy()
    body, width  := txt.DragAllTextByIndent( s.Src[ s.RunePos: ], block.Indent )
    s.NinjaLenMoves( width )
    ptrBody.Src = s.Src[ :s.RunePos ]

    if rebo.FindString( s.Line ) {
      s.NextLine()
      break
    } else if rech.FindString( s.Line ) {
      var miniblock MiniBlock
      miniblock.Body = new(Scanner)
      block.Attach = append( block.Attach, miniblock )
      ptrBody = miniblock.Body
      s.NextLine()
      goto redo
    } else {
      if txt.HasOnlySpaces( body ) {
        ptrBody.Src = ptrBody.Src[ :ptrBody.RunePos ]
      } else {
        ptrBody.Src = ptrBody.Src[ :ptrBody.RunePos + findEndOfLastTextLine( body ) ]
      }
    }
  }

  return &block
}

func findEndOfLastTextLine( str string ) int {
  if rms := len( txt.RmSpacesAtEnd( str ) ); rms < len( str ) {
    if str[ rms ] == '\n' { rms++ }
    return rms
  }

  return len( str )
}


func (s *Scanner) getBlockArgs() (args []ArgType, err string) {
  s.Scan()

redo:
  name, closing := "", rune( 0 )
  switch s.LastToken.Type {
  case '>' :
    if err != "" { s.Error( "GetBlock:" + err ) }
    return args, ""
  case EOF : return nil, "syntax error: unexpected EOF in block statement"
  case '\n': return nil, "syntax error: unexpected '\\n' in block statement"
  default:
    return nil, `syntax error: "` + s.LastToken.Text + `" no is a valid Argument Name`
  case ScanIdent:
    name = s.LastToken.Text
    switch s.Scan(); s.LastToken.Type {
    case '(': closing = ')'
    case '[': closing = ']'
    case '{': closing = '}'
    case '<': closing = '>'
    default:
      args = append( args, ArgType{ Name: name } )
      goto redo
    }
  }

  vargs, ferr := s.getFarg( name, s.LastToken.Text, closing )
  if err != "" && ferr != "" { err += ":" + ferr }
  args = append( args, ArgType{ Name: name, Args: vargs } )
  s.Scan()
  goto redo
}

var blockArgs = map[string][]Arg {
  "n"      : {{ "0", Int }},
  "style"  : {{ "nascii", String }},
  "prompt" : {{ ">", String }},
}

func (d *Scanner) toArgMap( args []ArgType ) ArgMap {
  am := make( ArgMap )
  for _, opt := range args {
    copt, cannon := d.syncOpt( opt, blockArgs )
    if !cannon {
      am[ opt.Name ] = opt.Args
      continue
    }

    am[ copt.Name ] = copt.Args
  }

  return am
}
