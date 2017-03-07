package biskana

import (
  "github.com/nasciiboy/regexp3"
)

func ToSafeHtml( str string ) (result string) {
  for i := 0; i < len( str ); i++ {
    switch  str[ i ] {
    case '<': result += "&lt;"
    case '>': result += "&gt;"
    case '&': result += "&amp;"
    case '"': result += "&quot;"
    default : result += str[i:i+1]
    }
  }

  return result
}

func ToHtml( str string ) (result string) {
  for tmpStr, forward, i :=  "", 0, 0; i < len( str ); i += forward {
    forward = 1;

    switch  str[ i ] {
    case '<': tmpStr = "&lt;"
    case '>': tmpStr = "&gt;"
    case '&': tmpStr = "&amp;"
    case '"': tmpStr = "&quot;"
    case '@': tmpStr, _, forward = marckupTrigger( str[i:] )
    default : tmpStr = str[i:i+1]
    }

    result += tmpStr
  }

  return result
}

func ToText( str string ) (result string) {
  for tmpStr, forward, i :=  "", 0, 0; i < len( str ); i += forward {
    forward = 1;

    switch  str[ i ] {
    case '<': tmpStr = "&lt;"
    case '>': tmpStr = "&gt;"
    case '&': tmpStr = "&amp;"
    case '"': tmpStr = "&quot;"
    case '@': _, tmpStr, forward = marckupTrigger( str[i:] )
    default : tmpStr = str[i:i+1]
    }

    result += tmpStr
  }

  return result
}

func marckupTrigger( str string ) (result, text string, width int) {
  switch len( str ) {
  case 0: return  "",  "", 0
  case 1: return str, str, 1
  }

  switch str[1] {
  case '(', ')', '[', ']', '{', '}',
       '@': return str[1:2], str[1:2], 2
  case '<': return   "&lt;",   "&lt;", 2
  case '>': return   "&gt;",   "&gt;", 2
  }

  switch len( str ) {
  case 2: return str, str, 2
  case 3: return str, str, 3
  }

  label, operator := str[ 1 ], str[ 2 ];

  switch operator {
  case '(': operator = ')'
  case '[': operator = ']'
  case '{': operator = '}'
  case '<': operator = '>'
  default : return str[:3], str[:3], 3
  }

  result, text, width = labelize( str[3:], label, operator )
  return result, text, width + 3
}

func labelize( str string, label, operator byte ) (result, text string, width int) {
  result, text, custom, width := marckupParser( str, operator )

  return ToLabel( result, spaceSwap( custom, " " ), label ), text, width
}

func marckupParser( str string, operator byte ) (result, text, custom string, i int) {
  for forward := 0; i < len( str ); i += forward {
    forward = 1

    if len(str[i:]) >= 2 && str[i:i+2] == "<>" {
      forward      = 2
      custom       = text
      text, result = "", ""
    } else {
      if str[ i ] == operator { i++; break }

      switch str[ i ] {
      case '<': result, text = result + "&lt;"  , text + "&lt;"
      case '>': result, text = result + "&gt;"  , text + "&gt;"
      case '&': result, text = result + "&amp;" , text + "&amp;"
      case '"': result, text = result + "&quot;", text + "&quot;"
      case '@':
        var tmpResult, tmpText string
        tmpResult, tmpText, forward = marckupTrigger( str[i:] )
        result += tmpResult
        text   += tmpText
      default : result += str[i:i+1]
                text   += str[i:i+1]
      }
    }
  }

  if custom == "" { custom = text }

  return result, text, custom, i
}

func ToLabel( body, custom string, label byte ) string {
  switch label {
  case '!' : return body
  case '"' : return "<cite>" + body + "</cite>"
  case '#' : return "<span class=\"path\" >" + body + "</span>"
  case '$' : return "<code class=\"command\" >" + body + "</code>"
  case '%' : return body // "parentesis"
  case '&' : return body // "simbol"
  case '\'': return "<samp>" + body + "</samp>"
  case '*' : return body
  case '+' : return body
  case ',' : return body
  case '-' : return "––" + body + "––"
  case '.' : return body
  case '/' : return body
  case '0' : return body
  case '1' : return body
  case '2' : return body
  case '3' : return body
  case '4' : return body
  case '5' : return body
  case '6' : return body
  case '7' : return body
  case '8' : return body
  case '9' : return body
  case ':' : return "<dfn>" + body + "</dfn>"
  case ';' : return body
  case '=' : return body
  case '?' : return body
  case 'A' : return "<span class=\"acronym\" >" + body + "</span>"
  case 'B' : return body
  case 'C' : return body // "smallCaps"
  case 'D' : return body
  case 'E' : return body // "error"
  case 'F' : return body // "Func"
  case 'G' : return body
  case 'H' : return body
  case 'I' : return body
  case 'J' : return body
  case 'K' : return body // "keyword"
  case 'L' : return body // "label"
  case 'M' : return body
  case 'N' : return "<span class=\"defnote\" id=\"" + ToLink(custom) + "\" >" + body + "</span>"
  case 'O' : return body
  case 'P' : return body
  case 'Q' : return body
  case 'R' : return body // "result"
  case 'S' : return body
  case 'T' : return body // "radiotarget"
  case 'U' : return body
  case 'V' : return body // "var"
  case 'W' : return body // "warning"
  case 'X' : return body
  case 'Y' : return body
  case 'Z' : return body
  case '\\': return body
  case '^' : return "<sup>" + body + "</sup>"
  case '_' : return "<sub>" + body + "</sub>"
  case '`' : return body
  case 'a' : return "<abbr>" + body + "</abbr>"
  case 'b' : return "<b>" + body + "</b>"
  case 'c' : return "<code>" + body + "</code>"
  case 'd' : return body // data
  case 'e' : return "<em>" + body + "</em>"
  case 'f' : return "<span class=\"file\" >" + body + "</span>"
  case 'g' : return body
  case 'h' : return body
  case 'i' : return "<i>" + body + "</i>"
  case 'j' : return body
  case 'k' : return "<kbd>" + body + "</kbd>"
  case 'l' :
    if custom[0] == '#' && body[0] == '#' { body = body[1:] }
    return "<a href=\"" + ToLink( custom ) + "\" >" + body + "</a>"
  case 'm' : return "<span class=\"math\" >" + body + "</span>"
  case 'n' : return "<span class=\"note\" ><sup><a href=\"#" + ToLink(custom) + "\" >" + body + "</a></sup></span>"
  case 'o' : return body
  case 'p' : return body
  case 'q' : return "<q>" + body + "</q>"
  case 'r' : return body // ref
  case 's' : return body
  case 't' : return "<span id=\"" + custom + "\" >" + body + "</span>"
  case 'u' : return "<u>" + body + "</u>"
  case 'v' : return "<code class=\"verbatim\" >" + body + "</code>"
  case 'w' : return body
  case 'x' : return body
  case 'y' : return body
  case 'z' : return body
  case '|' : return body
  case '~' : return body
  // case '@', '(', ')', '[', ']', '<', '>', '{', '}'
  }

  return "[unknow-Label/]" + body + "[/unknow-Label]"
}

func ToLink( link string ) string {
  var re regexp3.RE
  re.Match( linelize( link ), "<:s+>" )
  return re.RplCatch( "-", 1 )
}
