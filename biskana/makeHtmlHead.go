package biskana

import (
  "strings"
  "github.com/nasciiboy/regexp3"
)

var htmlHead string
var docInfo  DocInfo

func MakeHtmlHead( str string ) (string, DocInfo) {
  htmlHead = ""
  docInfo  = DocInfo{}

  setupHunter( str )

  return htmlHead, parseOptions( docInfo )
}

func parseOptions( docInfo DocInfo ) (result DocInfo) {
  result = docInfo

  var re regexp3.RE
  if re.Match( docInfo.options, "highlight" ) > 0 {
    result.optionsData.highlight = true
  }
  if re.Match( docInfo.options, "toc" ) > 0 {
    result.optionsData.toc = true
  }
  if re.Match( docInfo.options, "pygments" ) > 0 {
    result.optionsData.pygments = true
  }

  return result
}

func setupHunter( str string ){
  for init, width, line := 0, 0, ""; init < len(str); {
    line, width = getLine( str[init:] )

    switch whoIsThere( line ) {
    case COMMAND               : init += genSetup( str[init:] )
    case HEADLINE, TEXT , LIST : return
    case COMMENT , EMPTY       : init += width
    default                    : return
    }
  }
}

func genSetup( str string ) int {
  line, width := getLine( str )
  init        := width

  var re regexp3.RE
  if  re.Match( line, "#^:.:.:b*<[:w:-_]+>:b*<[^:>]*>:>:b*<.*>" ) == 0 {
    return init
  }

  command := strings.ToLower( re.GetCatch( 1 ) )
  options := re.GetCatch( 2 )
  arg     := re.GetCatch( 3 )

  switch command {
  case "title", "subtitle", "author", "translator", "lang", "language", "licence",
       "date", "tags", "mail", "description", "id", "style", "options":
    var head string
    head, width = dragTextByIndent( str[init:], 2 )
    arg         = linelize( spaceSwap( arg + head, " ") )
    init       += width

    makeSetup( command, options, arg )
  }

  return  init
}

func makeSetup( command, options, arg string ){
  switch command {
  case "title"      : docInfo.title       = arg; makeSetupTitle( options, arg )
  case "style"      : docInfo.style       = arg; makeSetupStyle( options, arg )
  case "subtitle"   : docInfo.subtitle    = arg; makeSetupMeta ( command, arg )
  case "author"     : docInfo.author      = arg; makeSetupMeta ( command, arg )
  case "translator" : docInfo.translator  = arg; makeSetupMeta ( command, arg )
  case "licence"    : docInfo.licence     = arg; makeSetupMeta ( command, arg )
  case "id"         : docInfo.id          = arg; makeSetupMeta ( command, arg )
  case "date"       : docInfo.date        = arg; makeSetupMeta ( command, arg )
  case "tags"       : docInfo.tags        = arg; makeSetupMeta ( command, arg )
  case "description": docInfo.description = arg; makeSetupMeta ( command, arg )
  case "mail"       : docInfo.mail        = arg; makeSetupMeta ( command, arg )
  case "options"    : docInfo.options     = arg
  case "lang", "language":
                      docInfo.lang        = arg
  }
}

func makeSetupTitle( options, args string ){
  htmlHead += "<title>" + ToHtml( args ) + "</title>\n"
}

func makeSetupMeta( name, args string ){
  htmlHead += "<meta name=\"" + name + "\" content=\"" + ToText( args ) + "\" />\n"
}

func makeSetupStyle( name, args string ){
  htmlHead += "<link rel=\"stylesheet\" type=\"text/css\" href=\"" + ToText( args ) + "\" />\n"
}
