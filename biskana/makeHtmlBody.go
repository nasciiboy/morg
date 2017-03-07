package biskana

import (
  "fmt"
  "strings"
  "github.com/nasciiboy/regexp3"
  "github.com/nasciiboy/pygments"
)

var htmlBody string
var htmlToc  string
var goptions Options

func MakeHtmlBodyWithOptions( str string, opt Options ) (string, string) {
  htmlBody, htmlToc = "", ""
  goptions = opt

  walkMorg( str, 0 )

  return htmlBody, htmlToc
}

func MakeHtmlBody( str string ) (string, string) {
  htmlBody, htmlToc = "", ""
  goptions = Options{}

  walkMorg( str, 0 )

  return htmlBody, htmlToc
}

func walkMorg( str string, level int ) int {
  for init, width, line := 0, 0, ""; init < len(str); {
    line, width = getLine( str[init:] )

    switch whoIsThere( line ) {
    case HEADLINE:
      if level > 0 { return init }
      init += getHeadline( str[init:], level )
    case TABLE   : init += getTable  ( str[init:] )
    case COMMAND : init += getCommand( str[init:] )
    case TEXT    : init += getText   ( str[init:] )
    case LIST    : init += walkList  ( str[init:] )
    case ABOUT   : init += getAbout  ( str[init:] )
    case COMMENT : init += width
    case EMPTY   : init += width
    default      : init += width
    }
  }

  return len( str )
}

func walkList( str string ) int {
  indentBase   := countIndentSpaces( str )
  indentLevel  := indentBase + 2
  sHead, wHead := dragListHeader( str, indentLevel  )
  sBody, wBody := dragAllTextByIndent( str[wHead:], indentLevel )
  listType     := whatListIsThere( sHead  )

  switch listType {
  case LIST_DIALOG           : htmlBody += "<ul class=\"dialog\" >\n"
  case LIST_MINUS, LIST_PLUS : htmlBody += "<ul>\n"
  case LIST_NUM,   LIST_ALPHA: htmlBody += "<ol>\n"
  case LIST_MDEF,  LIST_PDEF : htmlBody += "<dl>\n"
  }

  init, cListType := wHead + wBody, listType;
  for {
    makeList( sHead, sBody, cListType )

    sHead, wHead  = dragListHeader( str[init:], indentLevel  )

    cListType = whatListIsThere( sHead )
    if whoIsThere( sHead ) != LIST || cListType != listType || countIndentSpaces( sHead ) < indentBase {
      break
    }

    sBody, wBody = dragAllTextByIndent( str[init + wHead:], indentLevel )
    init += wHead + wBody
  }

  switch listType {
  case LIST_MINUS, LIST_PLUS, LIST_DIALOG : htmlBody += "</ul>\n"
  case LIST_NUM,   LIST_ALPHA             : htmlBody += "</ol>\n"
  case LIST_MDEF,  LIST_PDEF              : htmlBody += "</dl>\n"
  }

  return init
}

func makeList( head, body string, listType uint ){
  switch listType {
  case LIST_DIALOG           : makeDList( head, body )
  case LIST_MINUS, LIST_PLUS : makeUList( head, body )
  case LIST_NUM,   LIST_ALPHA: makeOList( head, body )
  case LIST_MDEF,  LIST_PDEF : makeDefList( head, body )
  }
}

func makeDList( head, body string ){
  var re regexp3.RE
  re.Match( head, "#^:b*:>:b+<:S>" )
  htmlBody += "<li>"
  walkMorg( head[ re.GpsCatch( 1 ): ], 0 )
  walkMorg( body, 0 )
  htmlBody += "</li>\n"
}

func makeOList( head, body string ){
  var re regexp3.RE
  re.Match( head, "#^:b*(-|:+|(:d+|:a+)[.)]):b+<:S>" )
  htmlBody += "<li>"
  walkMorg( head[ re.GpsCatch( 1 ): ], 0 )
  walkMorg( body, 0 )
  htmlBody += "</li>\n"
}

func makeUList( head, body string ){
  var re regexp3.RE
  re.Match( head, "#^:b*(-|:+|(:d+|:a+)[.)]):b+<:S>" )
  htmlBody += "<li>"
  walkMorg( head[ re.GpsCatch( 1 ): ], 0 )
  walkMorg( body, 0 )
  htmlBody += "</li>\n"
}

func makeDefList( head, body string ){
  var re regexp3.RE
  re.Match( head, "#^:b*(-|:+):b+<:S>" )
  head = head[ re.GpsCatch( 1 ): ]

  re.Match( head, "#?<:b::{2}><.?>" )
  body = head[ re.GpsCatch( 2 ): ] + body
  head = head[ :re.GpsCatch( 1 ) ]

  htmlBody += "<dt>"
  htmlBody += ToHtml( linelize( head ) )
  htmlBody += "</dt>\n"

  htmlBody += "<dd>"
  walkMorg( body, 0 )
  htmlBody += "</dd>\n"
}

func dragListHeader( str string, indentLevel int ) (string, int) {
  _, wHead    := getLine( str )
  _, wBody    := dragTextByIndent( str[wHead:], indentLevel )
  width       := wHead + wBody

  return str[:width], width
}

func getAbout( str string ) int {
  _, wHead    := getLine( str )
  _, wBody    := dragAllTextByIndent( str[wHead:], countIndentSpaces( str ) + 2 )
  width       := wHead + wBody
  head        := str[:width]

  var re regexp3.RE
  re.Match( str, "#^:b*::{2}:b+<:S>" )
  head = head[ re.GpsCatch( 1 ): ]

  re.Match( head, "#?<:b::{2}><.?>" )
  body := head[ re.GpsCatch( 2 ): ]
  head  = head[ :re.GpsCatch( 1 ) ]

  htmlBody += "<div class=\"about\" >\n"
  htmlBody += "<div class=\"about-dt\" >\n"
  walkMorg( head, 0 )
  htmlBody += "</div>\n"

  htmlBody += "<div class=\"about-dd\" >\n"
  walkMorg( body, 0 )
  htmlBody += "</div>\n"
  htmlBody += "</div>\n"

  return width
}

func getText( str string ) int {
  for init, width, line := 0, 0, ""; len(str[init:]) > 0; {
    line, width = getLine( str[init:] )

    switch whoIsThere( line ) {
    case COMMAND, TEXT, LIST:
      init += width
    case HEADLINE, COMMENT, EMPTY :
      htmlBody += "<p>";
      htmlBody += ToHtml( linelize(str[:init]) );
      htmlBody += "</p>\n";
      return init
    default      : init += width
    }
  }

  htmlBody += "<p>";
  htmlBody += ToHtml( linelize(str) );
  htmlBody += "</p>\n";

  return len(str)
}

func getHeadline( str string, level int ) int {
  sHead, width := getLine( str )

  var re regexp3.RE
  re.Match( sHead, "#^<:*+>:b+<.+>" )

  indentLevel  := len( re.GetCatch( 1 ) ) + 1
  hLevel       := len( re.GetCatch( 1 ) ) + goptions.HShift
  sBody, wBody := dragTextByIndent( str[width:], indentLevel )
  width        += wBody
  sHead         = spaceSwap( re.GetCatch( 2 ) + " " +  sBody, " " )
  resultHead, _, customHead, _ := marckupParser( sHead, 0 )
  resultHead, customHead = linelize(resultHead), linelize(customHead)

  htmlBody += fmt.Sprintf( "<h%d id=\"%s\" >", hLevel, ToLink( customHead ) )
  htmlBody += resultHead
  htmlBody += fmt.Sprintf( "</h%d>\n", hLevel )

  if goptions.Toc {
    htmlToc += fmt.Sprintf( "<span class=\"toc\" ><a class=\"h%d\" href=\"#%s\" >", hLevel, ToLink( customHead ) )
    htmlToc += resultHead
    htmlToc += fmt.Sprintf( "</a></span>\n" )
  }

  if len( str[width:] ) > 0 {
    if re.Match( str[width:], "#^(:b*\n)*:*+:b" ) > 0 {
      return width
    } else if re.Match( str[width:], "#?:S" ) > 0 {
      htmlBody += fmt.Sprintf( "<div class=\"hBody-%d\" >\n", hLevel )
      width    +=  walkMorg( str[width:], 1 )
      htmlBody += "</div>\n"
    }
  }

  return  width
}

func getTable( str string ) int {
  line, width := getLine( str )
  init        := 0
  indentLevel := countIndentSpaces( line )

  for whoIsThere( line ) == TABLE && indentLevel == countIndentSpaces( line ) {
    init += width
    line, width = getLine( str[init:] )
  }

  strTable := clearSpacesAtEnd( rmIndent( str[:init], indentLevel ) )
  makeTable( strTable )

  return init
}

func makeTable( str string ){
  htmlBody += "<table><tbody>\n"
  headerTable, width := getTableHeader( str )

  if width > 0 {
    htmlBody += "<thead><tr>"
    makeTableCell( headerTable, "th" )
    htmlBody += "</tr></thead>\n"
  }

  bodyTable := str[width:]

  if len(bodyTable) > 0 {
    htmlBody += "<tbody>\n"
    makeTableBody( bodyTable )
    htmlBody += "</tbody>\n"
  }

  htmlBody += "</table>\n"
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

func makeTableBody( str string ){
  row, init := getTableRow( str )
  for init < len(str) {
    htmlBody += "<tr>"
    makeTableCell( row, "td" )
    htmlBody += "</tr>\n"
    irow, width := getTableRow( str[init:] )
    row,  init   = irow, init + width
  }

  htmlBody += "<tr>"
  makeTableCell( row, "td" )
  htmlBody += "</tr>\n"
}

func makeTableCell( str, kind string ){
  var re regexp3.RE
  re.Match( str, ":|:b+<:b+:|>+#!" )

  for i := uint32(1); i <= re.TotCatch(); i++ {
    htmlBody += "<" + kind + ">"
    htmlBody += ToHtml( linelize( re.GetCatch( i ) ))
    htmlBody += "</" + kind + ">"
  }
}

func getCommand( str string ) int {
  line, width  := getLine( str )
  init         := width

  var re regexp3.RE
  re.Match( line, "#^<:b*>:.:.:b*<[:w:-:_]+><[^:>]*>:>:b*<.*>" )

  indentLevel  := len(re.GetCatch( 1 )) + 2
  command      := strings.ToLower( re.GetCatch( 2 ) )
  options      := re.GetCatch( 3 )
  args         := re.GetCatch( 4 )
  body         := ""

  closePattern := fmt.Sprintf( "#^%s:< (%s)#*:.:.", re.GetCatch( 1 ), command )

  switch command {
  case "title", "subtitle", "author", "translator", "lang", "language", "licence",
       "date", "tags", "mail", "description", "id", "style", "options":
    _, width    = dragTextByIndent( str[init:], indentLevel )
    init       += width
  case "figure", "img", "video", "ignore":
    var head string
    head, width = dragTextByIndent( str[init:], indentLevel )
    args        = linelize( spaceSwap( args + head, " ") )
    init       += width

    fallthrough
  case "src", "center", "bold", "emph", "italic", "quote", "example", "pre", "diagram":
    body, width = getBodyCommand( str[init:], closePattern, indentLevel )
    init       += width

    body = rmIndent( body, indentLevel )
  }

  makeCommand( command, options, args, body )

  return  init
}

func getBodyCommand( str, closePattern string, indentLevel int ) (body string, w int) {
  var re regexp3.RE

  for init, width, line := 0, 0, ""; len(str[init:]) > 0; {
    line, width = getLine( str[init:] )

    switch whoIsThere( line ) {
    case COMMAND, LIST:
      if indent := countIndentSpaces( line ); indent < 2 || indent < indentLevel {
        return clearSpacesAtEnd( str[:init] ), init
      }

      init += width
    case TEXT:
      if re.Match( line, closePattern ) > 0 {
        if len(str) >= init - 1 {
          return str[:init - 1 ], init + width
        }
        return str[:init], init + width
      } else if indent := countIndentSpaces( line ); indent < 2 || indent < indentLevel {
        return clearSpacesAtEnd( str[:init] ), init
      }

      init += width
    case HEADLINE, COMMENT:
      return clearSpacesAtEnd( str[:init] ), init
    case EMPTY : init += width
    default    : init += width
    }
  }

  return str, len(str)
}

func makeCommand( command, options, args, body string ){
  switch command {
  // case "title", "subtitle", "author", "translator", "lang", "language", "licence",
  //      "date", "tags", "mail", "description", "id", "style", "options":
  case "figure" : makeCommandFigure( options, args, body )
  case "img"    : makeCommandImg   ( options, args, body )
  case "video"  : makeCommandVideo ( options, args, body )
  case "quote"  : makeCommandQuote ( command, args, body )
  case "src"    : makeCommandSrc   ( options, args, body )
  case "example", "pre", "diagram":
    makeCommandExample( options, args, body )
  case "center", "bold", "emph", "italic":
    makeCommandFont( command, args, body )
  }
}

func makeCommandSrc( options, args, body string ){
  if goptions.Pygments {
    pygCode, make := pygments.Highlight( body, args, "html", "utf-8" )

    if make == false { goto simple }

    htmlBody += pygCode
    return
  }

simple:
  htmlBody += fmt.Sprintf( "<pre class=\"code\" ><code class=\"%s\">", args )
  htmlBody += ToSafeHtml( body )
  htmlBody += "</code></pre>\n"
}

func makeCommandFigure( options, args, body string ){
  htmlBody += "<div class=\"figure\" >\n"
  htmlBody += "<h1>" + ToHtml( args ) + "</h1>\n"
  walkMorg( body, 0 )
  htmlBody += "</div>\n"
}

func makeCommandImg( options, args, body string ){
  htmlBody += "<div class=\"img\">\n"
  htmlBody += "<img src=\"" + args + "\" />\n"
  walkMorg( body, 0 )
  htmlBody += "</div>\n"
}

func makeCommandVideo( options, args, body string ){
  htmlBody += "<video controls >\n"
  htmlBody += "<source src=\"" + args + "\""

  var re regexp3.RE
  if re.Match( args, "#$:.<ogg|mp4>" ) > 0 {
    htmlBody +=  " type=\"video/" + re.GetCatch( 1 ) + "\""
  }

  htmlBody += " >\n"
  walkMorg( body, 0 )
  htmlBody += "</video>\n"
}

func makeCommandFont( command, args, body string ){
  htmlBody += "<div class=\"" + command + "\" >\n"
  walkMorg( body, 0 )
  htmlBody += "</div>\n"
}

func makeCommandExample( options, args, body string ){
  htmlBody += "<pre><code class=\"example\">"
  htmlBody += ToSafeHtml( body )
  htmlBody += "</code></pre>\n"
}

func makeCommandQuote( options, args, body string ){
  htmlBody += "<blockquote>\n"
  walkMorg( body, 0 )
  htmlBody += "</blockquote>\n"
}
