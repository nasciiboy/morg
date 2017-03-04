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
  if re.Match( docInfo.Options, "highlight" ) > 0 {
    result.OptionsData.Highlight = true
  }
  if re.Match( docInfo.Options, "toc" ) > 0 {
    result.OptionsData.Toc       = true
  }
  if re.Match( docInfo.Options, "pygments" ) > 0 {
    result.OptionsData.Pygments  = true
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
  case "title"      : docInfo.Title       = arg; makeSetupTitle( options, arg )
  case "style"      : docInfo.Style       = arg; makeSetupStyle( options, arg )
  case "subtitle"   : docInfo.Subtitle    = arg; makeSetupMeta ( command, arg )
  case "author"     : docInfo.Author      = arg; makeSetupMeta ( command, arg )
  case "translator" : docInfo.Translator  = arg; makeSetupMeta ( command, arg )
  case "licence"    : docInfo.Licence     = arg; makeSetupMeta ( command, arg )
  case "id"         : docInfo.Id          = arg; makeSetupMeta ( command, arg )
  case "date"       : docInfo.Date        = arg; makeSetupMeta ( command, arg )
  case "tags"       : docInfo.Tags        = arg; makeSetupMeta ( command, arg )
  case "description": docInfo.Description = arg; makeSetupMeta ( command, arg )
  case "mail"       : docInfo.Mail        = arg; makeSetupMeta ( command, arg )
  case "options"    : docInfo.Options     = arg
  case "lang", "language":
                      docInfo.Lang        = arg
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
