package katana

import (
  "fmt"
  "strconv"
  "strings"

  "github.com/nasciiboy/txt"
  "github.com/nasciiboy/regexp4"
)

func (d *doc) GetToc(){
  if d.Rune == EOF { return }
  d.Toc = make( []DocNode, 0, 16 )

  headline := DocNode{ Node: Headline{ Level: 0 } }
  d.walkMorg( &headline )
  d.Toc = append( d.Toc, headline )

  for d.Rune != EOF {
    headline := d.getHeadline()
    d.walkMorg( &headline )
    d.Toc = append( d.Toc, headline )
  }
}

func (d *doc) walkMorg( node *DocNode ){
  for d.Rune != EOF {
    switch whoIsThere( d.Line ) {
    case NodeHeadline  : return
    case NodeTable     : d.getTable  ( node )
    case NodeBlock     : d.getBlock  ( node )
    case NodeText      : d.getText   ( node )
    case NodeList      : d.getList   ( node )
    case NodeAbout     : d.getAbout  ( node )
    case NodeSeparator : d.NextLine(); node.AddNode( Separator{} )
    case NodeComment   : d.NextLine()
    case NodeEmpty     : d.NextLine()
    default            : d.NextLine()
    }
  }
}

func (d *doc) getText( node *DocNode ){
  mScan := d.Scanner.Copy()
  d.NextLine()

  for run := true; run; {
    switch whoIsThere( d.Line ) {
    case NodeBlock, NodeText, NodeList, NodeSeparator:
      d.NextLine()
    case NodeHeadline, NodeComment, NodeEmpty:
      run = false
    default:
      run = false
    }
  }

  mScan.Src = d.PrevText()
  node.AddNode( mScan.GetFancyMarkup() )
}

func (d *doc) getList( node *DocNode ){
  indentBase  := txt.CountIndentSpaces( d.Line )
  indentLevel := indentBase + 2
  listType    := listHat( d.Text(), indentLevel )
  list        := DocNode{ Node: List{ ListType: listType } }

  for {
    d.getListElement( &list, listType, indentLevel )

    if whoIsThere( d.Line ) != NodeList || listType != listHat( d.Text(), indentLevel ) || txt.CountIndentSpaces( d.Line ) < indentBase {
      break
    }
  }

  node.Add( list )
}

func listHat( str string, indentLevel int ) int {
  list, _ := txt.DragLineAndTextByIndent( str, indentLevel )
  return whatListIsThere( list )
}

var relip = regexp4.Compile( "#^<:b*<-|:+|:>|(:d+|:a+)[.)]>>" )
var rede  = regexp4.Compile( "#?:b<::{2}><.?>" )

func (d *doc) getListElement( node *DocNode, listType, indent int ){
  list := d.Copy()
  d.NextLine()
  _, w := txt.DragAllTextByIndent( d.Text(), indent )
  d.NinjaLenMoves( w )
  list.Src = d.Src[ :d.RunePos ]

  re := relip.Copy()
  re.FindString( list.Text() )
  prefix := re.GetCatch( 2 )
  spaces := fmt.Sprintf( "%*s", re.LenCatch( 1 ), "" )
  list.Src, list.RunePos, list.PrevRunePos, list.SrcPos = re.RplCatch( spaces, 1 ), 0,0,0
  list.Line, list.getLine = "", true
  list.Init()

  listElement := DocNode{}
  switch listType {
  case NodeListMDef, NodeListPDef:
    head, _ := txt.DragTextByIndent( list.Text(), indent )

    if red := rede.Copy(); red.FindString( head ) {
      mark := list.Copy()
      mark.Src = list.Src[ :list.RunePos + red.GpsCatch(1) ]

      list.NinjaLenMoves( red.GpsCatch(2) )
      listElement.Node = ListElement{ Mark: mark.GetFancyMarkup(), Prefix: prefix }
    } else {
      listElement.Node = ListElement{ Prefix: prefix }
    }
    list.Line = txt.GetRawLine( list.Text() )
  default: listElement.Node = ListElement{ Prefix: prefix }
  }

  d.cloneStats().swapScanner( list ).walkMorg( &listElement )
  node.Add( listElement )
}

var rebout = regexp4.Compile( "#?:b<::{2}>" )

func (d *doc) getAbout( node *DocNode ) {
  indent := txt.CountInitSpaces( d.Line ) + 2
  d.NinjaLenMoves( indent )
  head := d.Copy()
  d.NextLine()
  _, w := txt.DragAllTextByIndent( d.Text(), indent )
  d.NinjaLenMoves( w )
  head.Src = d.Src[ :d.RunePos ]

  about := DocNode{}
  body := head.Copy()
  if re := rebout.Copy(); re.FindString( body.Text() ) {
    body.NinjaLenMoves( re.GpsCatch( 1 ) )
    head.Src = body.Src[ :body.RunePos ]
    body.NinjaMoves( 2 )
    body.Line = txt.GetRawLine( body.Text() )
    about.Node = About{ Mark: head.GetFancyMarkup() }
  } else {
    d.Error( "getAbout: empty body" )
    return
  }

  d.cloneStats().swapScanner( body ).walkMorg( &about )
  node.Add( about )
}

func (d *doc) getHeadline() (node DocNode) {
  hLevel := txt.CountInitChars( d.Line )
  d.NinjaMoves( hLevel )
  mScan := d.Copy()
  d.NextLine()
  _, wBody := txt.DragTextByIndent( d.Src[d.RunePos:], hLevel + 1 )
  d.NinjaLenMoves( wBody )
  mScan.Src = d.Src[ : d.RunePos ]
  node.Node = Headline{ Level: hLevel, Mark: mScan.GetFancyCustomMarkup( MarkupNil, 0 ) }
  return
}

func (d *doc) getTable( doc *DocNode ){
  init, indentLevel := d.RunePos, txt.CountIndentSpaces( d.Line )

  for d.NextLine(); whoIsThere( d.Line ) == NodeTable && indentLevel == txt.CountIndentSpaces( d.Line ); d.NextLine() {
  }

  strTable := txt.RmSpacesAtEnd( txt.RmIndent( d.Src[init:d.RunePos], indentLevel ) )

  table := DocNode{ Node: Table{} }
  makeTable( &table, strTable )
  doc.Add( table )
}

func makeTable( table *DocNode, str string ){
  headerTable, width := getTableHeader( str )

  if width > 0 {
    tableRow := DocNode{ Node: TableRow{ Type: TableHead } }
    makeTableRow( &tableRow, headerTable )
    table.Add( tableRow )
  }

  bodyTable := str[width:]

  if len(bodyTable) > 0 {
    makeTableBody( table, bodyTable )
  }
}

var rehe = regexp4.Compile( "#?\n<:b*:|(=+:|)+:b*\n*>" )

func getTableHeader( str string ) (string, int) {
  if re := rehe.Copy(); re.FindString( str ) {
    return str[:re.GpsCatch( 1 )], re.GpsCatch( 1 ) + re.LenCatch( 1 )
  }

  return "", 0
}

var rero = regexp4.Compile( "#?\n<:b*:|(-+:|)+:b*\n*>" )

func getTableRow( str string ) (string, int) {
  if re := rero.Copy(); re.FindString( str ) {
    return str[:re.GpsCatch( 1 )], re.GpsCatch( 1 ) + re.LenCatch( 1 )
  }

  return str, len(str)
}

func makeTableBody( table *DocNode, str string ){
  row, init := getTableRow( str )
  for init < len(str) {
    tableRow := DocNode{ Node: TableRow{ Type: TableBody } }
    makeTableRow( &tableRow, row )
    table.Add( tableRow )

    irow, width := getTableRow( str[init:] )
    row,  init   = irow, init + width
  }

  tableRow := DocNode{ Node: TableRow{ Type: TableBody } }
  makeTableRow( &tableRow, row )
  table.Add( tableRow )
}

func makeTableRow( doc *DocNode, str string ){
  s := txt.GetLines( str )
  var cells []string

  for _, line := range( s ) {
    var re regexp4.RE
    re.Match( line, ":|:b<:b:|>+#!" )

    for i := 1; i <= re.TotCatch(); i++ {
      if i <= len( cells ) {
        cells[i-1] += " " + re.GetCatch( i )
      } else {
        cells = append( cells, re.GetCatch( i ) )
      }
    }
  }

  for _, c := range( cells ) {
    m := new(Scanner).NewSrc( txt.Linelize( txt.SpaceSwap( c, " " ) ) ).QuietSplash().Init().GetMarkup()

    doc.AddNode( TableCell{ Mark: m } )
  }
}

func (d *doc) getBlock( node *DocNode ){
  if block := d.GetBlock(); block != nil {
    d.parseBlock( node, block )
    return
  }

  d.NextLine()
}

var suffix = regexp4.Compile( "#$:.<:w+>" )

func (d *doc) parseBlock( pNode *DocNode, b *Block ) {
  var node DocNode
  switch b.Comm.Text {
  case "figure":
    node.Node = Figure{ Args: b.Args, Title: b.Head.GetFancyMarkup() }
    if b.Body.Text() != "" {
      d.cloneStats().swapScanner( &b.Body ).walkMorg( &node )
    }
  case "src", "code"          : node = d.makeCode( b )
  case "srci"                 : node = d.makeSrci( b )
  case "cols"                 : node = d.makeCols( b )
  case "img", "video", "audio": node = d.makeMedia( b )
  case "quote"                : node = d.makeQuote( b )
  case "example", "pre", "math", "diagram", "art":
    node.Node = Brick{
      Type  : b.Comm.Text,
      Head  : txt.RmSpacesToTheSides( b.Head.Text() ),
      Body  : txt.RmIndent( b.Body.Text(), b.Indent ),
      Args  : b.Args,
    }
  case "center", "bold", "verse", "emph", "tab", "italic":
    node.Node = Wrap{ Type: b.Comm.Text, Head: txt.RmSpacesToTheSides( b.Head.Text() ), Args: b.Args }

    if b.Body.Text() != "" {
      d.cloneStats().swapScanner( &b.Body ).walkMorg( &node )
    }
  case "pret": node.Node = Pret{ IndentMarkup: b.Body.GetMarkup(), Indent: b.Indent, Args: b.Args }
  default: return
  }

  pNode.Add( node )
}

func (d *doc) makeCols( b *Block ) (node DocNode) {
  node.Node = Columns{ Head: txt.RmSpacesToTheSides( b.Head.Text() ), Args: b.Args }

  current := DocNode{}
  d.cloneStats().swapScanner( &b.Body ).walkMorg( &current )
  node.Add( current )

  for _, col := range b.Attach {
    current = DocNode{}
    d.cloneStats().swapScanner( col.Body ).walkMorg( &current )
    node.Add( current )
  }

  return
}

func (d *doc) makeMedia( b *Block ) (node DocNode) {
  media := Media{ Src: txt.RmSpacesToTheSides( b.Head.Text() ), Args: b.Args }
  media.Type = b.Comm.Text
  if re := suffix.Copy(); re.FindString( media.Src ) {
    media.Ext = strings.ToLower( re.GetCatch( 1 ) )
  } else {
    d.Error( `parseBlock: media-src "` + media.Src + `" no have extension` )
  }

  node.Node = media
  if b.Body.Text() != "" {
    d.cloneStats().swapScanner( &b.Body ).walkMorg( &node )
  }

  return
}

var requ = regexp4.Compile( "#^:b*--" )

func (d *doc) makeQuote( b *Block ) (node DocNode) {
  quote, re := Quote{}, requ.Copy()
  for b.Body.Rune != EOF {
    if txt.HasOnlySpaces( b.Body.Line ){
      b.Body.NextLine()
      continue
    } else if re.FindString( b.Body.Line ) {
      b.Body.Scan() // '-'-
      b.Body.Next() //  -'-'x'
      quot := b.Body.Copy()
      b.Body.NextLine()
      _, width := txt.DragTextByIndent( b.Body.Src[ b.Body.RunePos: ], b.Indent + 2 )
      b.Body.NinjaLenMoves( width )
      quot.Src = b.Body.Src[ :b.Body.RunePos ]
      m := quot.GetFancyMarkup()
      m.Type = 'q'
      quote.Quotex = append( quote.Quotex, m )
    } else {
      text := b.Body.Copy()
      _, width := txt.DragTextByIndent( b.Body.Src[ b.Body.RunePos: ], b.Indent )
      b.Body.NinjaLenMoves( width )
      text.Src = b.Body.Src[ :b.Body.RunePos ]
      quote.Quotex = append( quote.Quotex, text.GetFancyMarkup() )
    }
  }

  node.Node = quote
  return
}

func (d *doc) makeCode( b *Block ) (node DocNode) {
  var code Code
  argNum, boolNum := b.Args[ "n" ]
  if code.IndexNum = 1; boolNum {
    code.Indexed   = true
    code.IndexNum, _ = strconv.Atoi( argNum[0].Data )
  }

  argStyle, boolStyle := b.Args[ "style" ]
  if code.Style = d.TextOptions[ "fancyCode" ]; boolStyle {
    code.Style = argStyle[0].Data
  }

  code.Lang  = txt.RmSpacesToTheSides( b.Head.Text() )
  code.Body  = txt.RmIndent( b.Body.Text(), b.Indent )
  code.Args  = b.Args
  code.SBody = b.Body
  code.SBodyIndent = b.Indent

  node.Node = code
  return
}

func (d *doc) makeSrci( b *Block ) (node DocNode) {
  var srci Srci
  argNum, boolNum := b.Args[ "n" ]
  if srci.IndexNum = 1; boolNum {
    srci.Indexed   = true
    srci.IndexNum, _ = strconv.Atoi( argNum[0].Data )
  }

  argStyle, boolStyle := b.Args[ "style" ]
  if srci.Style = d.TextOptions[ "fancyCode" ]; boolStyle {
    srci.Style = argStyle[0].Data
  }

  argPrompt, boolPrompt := b.Args[ "prompt" ]
  if srci.Prompt = "> "; boolPrompt {
    srci.Prompt = argPrompt[0].Data
  }


  srci.Lang  = txt.RmSpacesToTheSides( b.Head.Text() )
  srci.Args  = b.Args

  srci.Body = make( []BinariString, 0, 4 )
  body := txt.RmIndent( b.Body.Text(), b.Indent )
  scann := new(Scanner).NewSrc( body ).QuietSplash().Init()
  for scann.Rune != EOF  {
    if len( scann.Line ) >= 2 && scann.Line[:2] == "> " {
      srci.Body = append( srci.Body, BinariString{  true, srciGetCode( scann ) } )
    } else {
      srci.Body = append( srci.Body, BinariString{ false, srciGetText( scann ) } )
    }
  }

  node.Node = srci
  return
}

func srciGetCode( s *Scanner ) string {
  init := s.RunePos
  for s.NextLine(); reni( s.Line ); s.NextLine() {}

  return txt.RmInitRect( s.Src[init:s.RunePos], 2 )
}

func reni( str string ) bool {
  switch len( str ){
  case 0: return false
  case 1:
    if str[0] == '^' { return true }
  default:
    if str[0] == '^' && txt.CountInitSpaces( str[1:] ) > 0 { return true }
  }

  return false
}

func srciGetText( s *Scanner ) string {
  init := s.RunePos
  for ; s.Rune != EOF; s.NextLine() {
    if len( s.Line ) >= 2 && s.Line[:2] == "> " { break }
  }

  return s.Src[ init:s.RunePos]
}
