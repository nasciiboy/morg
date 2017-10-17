package katana

import (
  "fmt"
  "bytes"

  "github.com/nasciiboy/regexp4"
  "github.com/nasciiboy/txt"
)

type Doc struct {
  Toc         []DocNode

  Title         Markup
  Subtitle      Markup
  Lang          string
  Licence       string
  Date          string
  ID            string
  Description   string
  Author      []string
  Translator  []string
  Source      []string
  Style       []string
  Tags        []string
  HShift        int
  BoolOptions   map[string]bool      // as ==> someOption || someOption()
  TextOptions   map[string]string    // as ==> someOption("text")
  NumbOptions   map[string]int       // as ==> someOption( numb )
  MultOptions   map[string][]Arg     // as ==> someOption( numb, "text", float, bool, ident, ... )
  MoreConfigs   map[string][]string  // as ==> ..whatever > whatever
}

type doc struct {
  Doc
  toc DocNode
  *Scanner
}

func Parse( name, src string ) (d *Doc, errs string) {
  doc, buff   := new(doc), new(bytes.Buffer)
  doc.Scanner  = NewScanner( src )
  doc.Scanner.CustomError = func (s *Scanner, msg string){
    fmt.Fprintf( buff, "katana:%s: %s\n", s.Pos(), msg )
  }
  doc.Scanner.Name = name
  doc.BoolOptions  = make( map[string]bool     )
  doc.TextOptions  = make( map[string]string   )
  doc.NumbOptions  = make( map[string]int      )
  doc.MultOptions  = make( map[string][]Arg    )
  doc.MoreConfigs  = make( map[string][]string )
  buff.Grow( 4096 )
  doc.Scanner.Init()
  doc.SetupHunter()
  doc.GetToc()

  return &doc.Doc, buff.String()
}

func (d *doc) cloneStats() *doc {
  nd := new(doc)
  nd.Title        = d.Title
  nd.Subtitle     = d.Subtitle
  nd.Lang         = d.Lang
  nd.Licence      = d.Licence
  nd.Date         = d.Date
  nd.ID           = d.ID
  nd.Description  = d.Description
  nd.Author       = d.Author
  nd.Translator   = d.Translator
  nd.Source       = d.Source
  nd.Style        = d.Style
  nd.Tags         = d.Tags
  nd.HShift       = d.HShift
  nd.BoolOptions  = d.BoolOptions
  nd.TextOptions  = d.TextOptions
  nd.NumbOptions  = d.NumbOptions
  nd.MultOptions  = d.MultOptions
  nd.MoreConfigs  = d.MoreConfigs
  nd.Scanner      = d.Scanner

  return nd
}

func (d *doc) swapScanner( s *Scanner ) *doc {
  d.Scanner = s
  return d
}

var reComment   = regexp4.Compile( "#^$:b*:@(:b+.+|:s*)"               )
var reHeadline  = regexp4.Compile( "#^:*+:b"                           )
var reTable     = regexp4.Compile( "#^$:b*:|([^|]+:|)+:s*"             )
var reList      = regexp4.Compile( "#^:b*(-|:+|:>|(:d+|:a+)[.)]):b+:S" )
var reAbout     = regexp4.Compile( "#^:b*::{2}:b+:S"                   )
var reCommand   = regexp4.Compile( "#^:b*:.:.:b*[:w:-:_:&]+[^:>\n]*:>" )
var reSeparator = regexp4.Compile( "#^$:b*:.{4}:s*"                    )

func whoIsThere( line string ) int {
  if len( line ) == 0 || txt.HasOnlySpaces( line ) { return NodeEmpty
  } else if reComment  .Copy().FindString( line )  { return NodeComment
  } else if reHeadline .Copy().FindString( line )  { return NodeHeadline
  } else if reTable    .Copy().FindString( line )  { return NodeTable
  } else if reList     .Copy().FindString( line )  { return NodeList
  } else if reAbout    .Copy().FindString( line )  { return NodeAbout
  } else if reCommand  .Copy().FindString( line )  { return NodeBlock
  } else if reSeparator.Copy().FindString( line )  { return NodeSeparator
  } else                                           { return NodeText
  }
}

var reListPreMP  = regexp4.Compile( "#^:b*<-|:+>:b+<:S>" )
var reListPosDef = regexp4.Compile( "#?:b::{2}"          )
var reListNum    = regexp4.Compile( "#^:b*:d+[.)]:b+:S"  )
var reListAlpha  = regexp4.Compile( "#^:b*:a+[.)]:b+:S"  )
var reListDialog = regexp4.Compile( "#^:b*:>:b+:S"       )

func whatListIsThere( line string ) int {
  re := reListPreMP.Copy()
  if re.FindString( line ) {
    if reListPosDef.Copy().FindString( line[re.GpsCatch( 2 ):] ) {
      if re.GetCatch( 1 ) == "-" { return NodeListMDef }
      return  NodeListPDef
    }

    if re.GetCatch( 1 ) == "-" { return NodeListMinus }
    return NodeListPlus
  } else if reListNum   .Copy().FindString( line ) { return NodeListNum
  } else if reListAlpha .Copy().FindString( line ) { return NodeListAlpha
  } else if reListDialog.Copy().FindString( line ) { return NodeListDialog
  }

  return NodeEmpty
}
