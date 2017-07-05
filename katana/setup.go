package katana

import (
  "strings"

  "github.com/nasciiboy/regexp3"
  "github.com/nasciiboy/utils/text"

)

func GetSetup( str string ) (Doc, int)  {
  setup, width := setupHunter( str )

  return parseOptions( setup ), width
}

func parseOptions( docInfo Doc ) Doc {
  if searchOption( docInfo.Options, "highlight" ) {
    docInfo.OptionsData.Highlight = true
  }

  if searchOption( docInfo.Options, "toc" ) {
    docInfo.OptionsData.Toc       = true
  }

  if searchOption( docInfo.Options, "pygments" ) {
    docInfo.OptionsData.Pygments  = true
  }

  if searchOption( docInfo.Options, "mathjax" ) {
    docInfo.OptionsData.Mathjax   = true
  }

  return docInfo
}

func searchOption( str []string, keyword string ) bool {
  var re regexp3.RE

  for _, s := range( str ) {
    if re.Match( s, keyword ) > 0 {
      return true
    }
  }
  return false
}

func setupHunter( str string ) (setup Doc, w int) {
  for init, width, line := 0, 0, ""; init < len(str); {
    line, width = text.GetLine( str[init:] )

    switch whoIsThere( line ) {
    case CommandNode:
      if isSetup( line ) == false { return setup, init }

      init += getSetupCommand( str[init:], &setup )
    case HeadlineNode, TextNode , ListNode: return setup, init
    case CommentNode , EmptyNode: init += width
    default: return setup, init
    }
  }

  return setup, len( str )
}

func isSetup( str string ) bool {
  var re regexp3.RE
  if  re.Match( str, "#^:.:.:b*<[:w:-_]+>:b*([^:>]*):>" ) == 0 {
    return false
  }

  switch re.GetCatch( 1 ) {
  case "title", "subtitle", "author", "translator", "lang", "language", "licence",
       "date", "tags", "mail", "description", "id", "style", "options", "copy", "genre", "cover":
    return true
  }

  return false
}

func getSetupCommand( str string, docInfo *Doc ) int {
  line, width := text.GetLine( str )
  init        := width

  var re regexp3.RE
  re.Match( line, "#^:.:.:b*<[:w:-_]+>:b*<[^:>]*>:>:b*<.*>" )

  command := strings.ToLower( re.GetCatch( 1 ) )
  options := re.GetCatch( 2 )
  arg     := re.GetCatch( 3 )

  var head string
  head, width = text.DragTextByIndent( str[init:], 2 )
  arg         = text.Linelize( text.SpaceSwap( arg + head, " ") )
  init       += width

  setSetupCommand( command, options, arg, docInfo )

  return  init
}

func setSetupCommand( command, options, arg string, docInfo *Doc ){
  switch command {
  case "title"      : docInfo.Title, _, _    = MarkupParser( arg, MarkupTitle, 0 )
  case "style"      : docInfo.Style          = append( docInfo.Style, arg )
  case "subtitle"   : docInfo.Subtitle, _, _ = MarkupParser( arg, MarkupSubTitle, 0 )
  case "author"     : docInfo.Author         = append( docInfo.Author, arg )
  case "translator" : docInfo.Translator     = append( docInfo.Translator, arg )
  case "licence"    : docInfo.Licence        = arg
  case "id"         : docInfo.Id             = arg
  case "date"       : docInfo.Date           = arg
  case "tags"       : docInfo.Tags           = arg
  case "description": docInfo.Description    = arg
  case "mail"       : docInfo.Mail           = arg
  case "copy"       : docInfo.Copy           = arg
  case "options"    : docInfo.Options        = append( docInfo.Options, arg )
  case "lang"       : docInfo.Lang           = arg
  case "language"   : docInfo.Lang           = arg
  }
}
