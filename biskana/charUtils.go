package biskana

import (
  "fmt"
  "github.com/nasciiboy/regexp3"
)

func getLine( str string ) (string, int) {
  for i, c := range str {
    if c == '\n' {
      return str[:i], i + 1
    }
  }

  return str, len( str )
}

func clearSpacesAtEnd( str string ) string {
  var re regexp3.RE
  if re.Match( str, "#$<:s+>" ) > 0 {
    return str[:re.GpsCatch( 1 )]
  }

  return str
}

func linelize( str string ) (result string){
  result = str

  var re regexp3.RE
  if re.Match( str, "#^<:s+>" ) > 0 {
    result = str[re.LenCatch( 1 ):]
  }

  if re.Match( result, "<:b*\n:b*>" ) > 0 {
    tmp := re.RplCatch( " ", 1 )
    result = tmp
  }

  return clearSpacesAtEnd( result )
}

func spaceSwap( str, swap string ) (result string){
  var re regexp3.RE
  if re.Match( str, "<:s+>" ) > 0 {
    return re.RplCatch( swap, 1 )
  }

  return str
}

func rmIndent( str string, indentLevel int ) (result string) {
  var re regexp3.RE

  if re.Match( str, fmt.Sprintf( "#^:b{%d}", indentLevel )) > 0 {
    str = str[ indentLevel : ]
  }

  re.Match( str, fmt.Sprintf( "<\n:b{%d}>", indentLevel ))
  result = re.RplCatch( "\n", 1 )

  return result
}

func countIndentSpaces( str string ) int {
  var re regexp3.RE
  re.Match( str, "#^<:b*>" )

  return len(re.GetCatch( 1 ))
}

func dragTextByIndent( str string, indent int ) (string, int) {
  var re regexp3.RE
  strIndent := fmt.Sprintf( "#^:b{%d,}:S", indent )

  for init, width, line := 0, 0, ""; len(str[init:]) > 0; {
    line, width = getLine( str[init:] )

    if re.Match( line, strIndent ) == 0 {
      return str[:init], init
    }

    init += width
  }

  return str, len(str)
}

func dragAllTextByIndent( str string, indent int ) (string, int) {
  var re regexp3.RE
  strIndent := fmt.Sprintf( "#^:b{%d,}:S", indent )

  for init, width, line := 0, 0, ""; len(str[init:]) > 0; {
    line, width = getLine( str[init:] )

    if re.Match( line, strIndent ) == 0  {
      switch whoIsThere( line ) {
      default: return str[:init], init
      case EMPTY, COMMENT:
        init += width
        continue
      }
    }

    init += width
  }

  return str, len(str)
}
