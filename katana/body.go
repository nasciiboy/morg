package katana

import (
  "strings"
  "fmt"

  "github.com/nasciiboy/regexp3"
  "github.com/nasciiboy/utils/text"
)

func GetToc( str string ) []DocNode {
  var toc DocNode
  toc.AddNode( Headline{ Level: 0 } )
  init := walkMorg( toc.GetLast(), str )

  for width, line := 0, ""; init < len( str ); {
    line, width = text.GetLine( str[init:] )

    if whoIsThere( line ) == HeadlineNode {
      init += getHeadline( &toc, str[init:] )
      init += walkMorg( toc.GetLast(), str[init:] )
      continue
    }

    init += width
  }

  return toc.Cont
}

func walkMorg( doc *DocNode, str string ) int {
  for init, width, line := 0, 0, ""; init < len(str); {
    line, width = text.GetLine( str[init:] )

    switch whoIsThere( line ) {
    case HeadlineNode: return init
    case TableNode   : init += getTable  ( doc, str[init:] )
    case CommandNode : init += getCommand( doc, str[init:] )
    case TextNode    : init += getText   ( doc, str[init:] )
    case ListNode    : init += walkList  ( doc, str[init:] )
    case AboutNode   : init += getAbout  ( doc, str[init:] )
    case CommentNode : init += width
    case EmptyNode   : init += width
    default          : init += width
    }
  }

  return len( str )
}

func getText( doc *DocNode, str string ) int {
  var mark Markup

  for init, width, line := 0, 0, ""; len(str[init:]) > 0; {
    line, width = text.GetLine( str[init:] )

    switch whoIsThere( line ) {
    case CommandNode, TextNode, ListNode:
      init += width
    case HeadlineNode, CommentNode, EmptyNode :
      mark.Parse( text.Linelize(str[:init]) )
      mark.Type = MarkupText
      doc.AddNode( Text{ Mark: mark } )

      return init
    default      : init += width
    }
  }

  mark.Parse( text.Linelize( str ) )
  mark.Type = MarkupText
  doc.AddNode( Text{ Mark: mark } )

  return len(str)
}

func walkList( doc *DocNode, str string ) int {
  indentBase    := text.CountIndentSpaces( str )
  indentLevel   := indentBase + 2
  sHead, wHead  := dragListHeader( str, indentLevel  )
  listType      := whatListIsThere( sHead )
  sBody, wBody  := dragAllTextByIndent( str[wHead:], indentLevel )
  init          := wHead + wBody

  doc.AddNode( List{ ListType: listType } )
  cListType := listType

  for {
    doc.AddToLast( makeNodeList( sHead, sBody, cListType ) )

    sHead, wHead  = dragListHeader( str[init:], indentLevel  )

    cListType = whatListIsThere( sHead )
    if whoIsThere( sHead ) != ListNode || cListType != listType || text.CountIndentSpaces( sHead ) < indentBase {
      break
    }

    sBody, wBody = dragAllTextByIndent( str[init + wHead:], indentLevel )
    init += wHead + wBody
  }

  return init
}

func makeNodeList( sHead, sBody string, listType int ) (node DocNode) {
  listElement := ListElement{}
  sHead, listElement.Prefix = rmListPrefix( sHead, listType )

  switch listType {
  case ListMdefNode, ListPdefNode:
    var re regexp3.RE
    re.Match( sHead, "#^:b*(-|:+):b+<:S>" )
    sHead = sHead[ re.GpsCatch( 1 ): ]

    re.Match( sHead, "#?<:b::{2}><.?>" )
    sBody = sHead[ re.GpsCatch( 2 ): ] + sBody
    sHead = sHead[ :re.GpsCatch( 1 ) ]

    mark := Markup{}
    mark.Parse( text.Linelize( sHead )  )
    listElement.Mark = mark

    walkMorg( &node, sBody )
  case ListMinusNode, ListPlusNode, ListNumNode, ListAlphaNode, ListDialogNode:
    walkMorg( &node, sHead + "\n" + sBody )
  }

  node.Node = listElement
  return
}

func rmListPrefix( list string, listType int ) (text, prefix string) {
  var re regexp3.RE

  switch listType {
  case ListMinusNode, ListPlusNode :
    re.Match( list, "#^<:b*<-|:+>:b+>" )
  case ListNumNode,   ListAlphaNode:
    re.Match( list, "#^<:b*<:d+|:a+>[.)]:b+>" )
  case ListMdefNode, ListPdefNode:
    re.Match( list, "#^<:b*<-|:+>:b+>" )
  case ListDialogNode: re.Match( list, "#^<:b*<:>>:b+>" )
  }

  text  = fmt.Sprintf( "%*s", re.LenCatch( 1 ), "" )
  text += list[ re.LenCatch( 1 ):]
  prefix = re.GetCatch( 2 )

  return
}

func dragListHeader( str string, indentLevel int ) (string, int) {
  _, wHead    := text.GetLine( str )
  _, wBody    := text.DragTextByIndent( str[wHead:], indentLevel )
  width       := wHead + wBody

  return str[:width], width
}

func getAbout( doc *DocNode, str string ) int {
  _, wHead    := text.GetLine( str )
  _, wBody    := dragAllTextByIndent( str[wHead:], text.CountIndentSpaces( str ) + 2 )
  width       := wHead + wBody
  head        := str[:width]

  var re regexp3.RE
  re.Match( str, "#^<:b*::{2}:b+>" )
  head = head[ re.LenCatch( 1 ): ]

  re.Match( head, "#?<:b::{2}><.?>" )
  body := head[ re.GpsCatch( 2 ): ]
  head  = text.Linelize( head[ :re.GpsCatch( 1 ) ] )

  mark := Markup{}
  mark.Parse( head )
  doc.AddNode( About{ Mark: mark } )
  walkMorg( doc.GetLast(), body )

  return width
}

func getHeadline( toc *DocNode, str string ) int {
  sHead, width := text.GetLine( str )

  var re regexp3.RE
  re.Match( sHead, "#^<:*+>:b+<.+>" )

  hLevel       := len( re.GetCatch( 1 ) )
  indentLevel  := hLevel + 1
  sBody, wBody := text.DragTextByIndent( str[width:], indentLevel )
  width        += wBody
  sHead         = text.Linelize( text.SpaceSwap( re.GetCatch( 2 ) + " " +  sBody, " " ) )

  if re.Match( sHead, "#?<:s*:<:>:s*>" ) > 0 {
    sHead = re.RplCatch( "<>", 1 )
  }

  mark, _, _   := MarkupParser( sHead, MarkupHeadline, 0 )
  toc.AddNode( Headline{ Mark: mark, Level: hLevel } )

  return  width
}

func getTable( doc *DocNode, str string ) int {
  line, width := text.GetLine( str )
  init        := 0
  indentLevel := text.CountIndentSpaces( line )
  doc.AddNode( Table{} )

  for whoIsThere( line ) == TableNode && indentLevel == text.CountIndentSpaces( line ) {
    init += width
    line, width = text.GetLine( str[init:] )
  }

  strTable := text.RmSpacesAtEnd( text.RmIndent( str[:init], indentLevel ) )
  makeTable( doc.GetLast(), strTable )

  return init
}

func makeTable( table *DocNode, str string ){
  headerTable, width := getTableHeader( str )

  if width > 0 {
    table.AddNode( TableRow{ Kind: TableHead } )
    makeTableRow( table.GetLast(), headerTable )
  }

  bodyTable := str[width:]

  if len(bodyTable) > 0 {
    makeTableBody( table, bodyTable )
  }
}

func getTableHeader( str string ) (string, int) {
  var re regexp3.RE
  if re.Match( str, "#?\n<:b*:|(=+:|)+:b*\n*>" ) > 0 {
    return str[:re.GpsCatch( 1 )], re.GpsCatch( 1 ) + re.LenCatch( 1 )
  }

  return "", 0
}

func getTableRow( str string ) (string, int) {
  var re regexp3.RE
  if re.Match( str, "#?\n<:b*:|(-+:|)+:b*\n*>" ) > 0 {
    return str[:re.GpsCatch( 1 )], re.GpsCatch( 1 ) + re.LenCatch( 1 )
  }

  return str, len(str)
}

func makeTableBody( table *DocNode, str string ){
  row, init := getTableRow( str )
  for init < len(str) {
    table.AddNode( TableRow{ Kind: TableBody } )
    makeTableRow( table.GetLast(), row )

    irow, width := getTableRow( str[init:] )
    row,  init   = irow, init + width
  }

  table.AddNode( TableRow{ Kind: TableBody } )
  makeTableRow( table.GetLast(), row )
}

func makeTableRow( doc *DocNode, str string ){
  s := text.GetLines( str )
  var cells []string

  for _, line := range( s ) {
    var re regexp3.RE
    re.Match( line, ":|:b<:b:|>+#!" )

    for i := 1; i <= re.TotCatch(); i++ {
      if i <= len( cells ) {
        cells[i-1] += " " + re.GetCatch( i )
      } else {
        cells = append( cells, re.GetCatch( i ) )
      }
    }
  }

  m := Markup{}
  for _, c := range( cells ) {
    m.Parse( text.Linelize( text.SpaceSwap( c, " " ) ) )
    doc.AddNode( TableCell{ Mark: m } )
  }
}


func getCommand( doc *DocNode, str string ) int {
  line, width  := text.GetLine( str )
  init         := width

  var re regexp3.RE
  re.Match( line, "#^<:b*>:.:.:b*<[:w:-:_]+><[^:>]*>:>:b*<.*>" )

  indentLevel  := len(re.GetCatch( 1 )) + 2
  command      := strings.ToLower( re.GetCatch( 2 ) )
  params       := re.GetCatch( 3 )
  args         := re.GetCatch( 4 )
  body         := ""

  closePattern := fmt.Sprintf( "#^%s:< (%s)#*:.:.", re.GetCatch( 1 ), command )

  switch command {
  case "title", "subtitle", "author", "translator", "lang", "language", "licence",
       "date", "tags", "mail", "description", "id", "style", "options":
    _, width    = text.DragTextByIndent( str[init:], indentLevel )
    return init + width
  case "figure", "img", "video", "ignore":
    var head string
    head, width = text.DragTextByIndent( str[init:], indentLevel )
    args        = text.Linelize( text.SpaceSwap( args + head, " ") )
    init       += width

    fallthrough
  case "center", "bold", "emph", "italic", "cols":
    body, width = getBodyCommand( str[init:], closePattern, indentLevel )
    init       += width
  case "src", "example", "pre", "diagram", "art", "quote":
    body, width = getBodyCommand( str[init:], closePattern, indentLevel )
    init       += width

    body = text.RmIndent( body, indentLevel )
  }

  doc.Add( makeCommand( command, params, args, body ) )

  return  init
}

func getBodyCommand( str, closePattern string, indentLevel int ) (body string, w int) {
  var re regexp3.RE

  for init, width, line := 0, 0, ""; len(str[init:]) > 0; {
    line, width = text.GetLine( str[init:] )

    switch whoIsThere( line ) {
    case CommandNode, ListNode:
      if indent := text.CountIndentSpaces( line ); indent < 2 || indent < indentLevel {
        return text.RmSpacesAtEnd( str[:init] ), init
      }

      init += width
    case TextNode:
      if re.Match( line, closePattern ) > 0 {
        return text.RmSpacesAtEnd( str[:init] ), init + width
      } else if indent := text.CountIndentSpaces( line ); indent < 2 || indent < indentLevel {
        return text.RmSpacesAtEnd( str[:init] ), init
      }

      init += width
    case HeadlineNode, CommentNode:
      return text.RmSpacesAtEnd( str[:init] ), init
    case EmptyNode : init += width
    default        : init += width
    }
  }

  return str, len(str)
}

func makeCommand( command, params, arg, body string ) (node DocNode) {
  switch command {
  case "figure":
    mark := Markup{}
    mark.Parse( arg )

    node.Node = Command{ Comm: command, Params: params, Mark: mark }
    walkMorg( &node, body )
  case "cols":
    node = makeCommandCols( command, params, arg, body )
  case "img", "video":
    node.Node = Command{ Comm: command, Params: params, Arg: arg }
    walkMorg( &node, body )
  case "quote":
    node = makeCommandQuote( command, params, body )
  case "src", "example", "pre", "diagram", "art":
    node.Node = Command{ Comm: command, Params: params, Arg: arg, Body: body }
  case "center", "bold", "emph", "italic":
    node.Node = Command{ Comm: command, Params: params }
    walkMorg( &node, body )
  }

  return
}

func makeCommandCols( command, params, arg, body string ) (node DocNode) {
  node.Node = Command{ Comm: command, Params: params }
  cols := getCols( body )

  for _, col := range( cols ) {
    node.AddNode( nil )
    walkMorg( node.GetLast(), col )
  }

  return
}

func getCols( str string ) (result []string) {
  init, width, last, line := 0, 0, 0, "";
  var re regexp3.RE
  for init < len(str) {
    line, width = text.GetLine( str[init:] )

    if re.Match( line, "#^$:b*:::::b*" ) > 0 {
      result = append( result, str[last:init] )
      last = init + width
    }

    init += width
  }

  if last < init {
    result = append( result, str[last:init] )
  }

  return result
}

func makeCommandQuote( command, params, body string ) (node DocNode) {
  node.Node = Command{ Comm: command, Params: params }
  init, width, line := 0, 0, "";
  var re regexp3.RE
  for init < len(body) {
    line, width = text.GetLine( body[init:] )
    if whoIsThere( line ) == EmptyNode {
      init += width
    } else if re.Match( line, "#^--" ) > 0 {
      init += width
      t, w := text.DragTextByIndent( body[init:], 2 )
      mark := Markup{}
      mark.Parse( text.Linelize( line[2:] + " " + t ) )
      node.AddNode( Text{ Mark: mark, TextType: TextQuoteAuthor } )
      init += w
    } else {
      init += getText( &node, body[init:] )
    }
  }

  return
}

func dragAllTextByIndent( str string, indent int ) (string, int) {
  var re regexp3.RE
  strIndent := fmt.Sprintf( "#^:b{%d,}:S", indent )

  for init, width, line := 0, 0, ""; init < len(str); {
    line, width = text.GetLine( str[init:] )

    if re.Match( line, strIndent ) == 0  {
      switch whoIsThere( line ) {
      case EmptyNode, CommentNode:
        init += width
        continue
      default: return str[:init], init
      }
    }

    init += width
  }

  return str, len(str)
}
