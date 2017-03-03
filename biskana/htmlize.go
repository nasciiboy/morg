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

func marckupTrigger( str string ) (string, string, int) {
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

  strLize, strText, i := labelize( str[3:], label, operator )
  return strLize, strText, i + 3
}

func labelize( str string, label, operator byte ) (string, string, int) {
  body, text, custom, i := marckupParser( str, operator )

  return ToLabel( label, body, spaceSwap( custom, " " ) ), text, i
}

func marckupParser( str string, operator byte ) (body, text, custom string, i int) {
  for forward := 0; i < len( str ); i += forward {
    forward = 1

    if len(str[i:]) >= 2 && str[i:i+2] == "<>" {
      forward    = 2
      custom     = text
      text, body = "", ""
    } else {
      if str[ i ] == operator { i++; break }

      switch str[ i ] {
      case '<': body += "&lt;"
                text += "&lt;"
      case '>': body += "&gt;"
                text += "&gt;"
      case '&': body += "&amp;"
                text += "&amp;"
      case '"': body += "&quot;"
                text += "&quot;"
      case '@':
        var tmpBody, tmpText string
        tmpBody, tmpText, forward = marckupTrigger( str[i:] )
        body += tmpBody
        text += tmpText
      default : body += str[i:i+1]
                text += str[i:i+1]
      }
    }
  }

  if custom == "" { custom = text }

  return body, text, custom, i
}

func ToLabel( label byte, body, text string ) string {
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
  case 'N' : return "<span class=\"defnote\" id=\"" + ToLink(body) + "\" >" + body + "</span>"
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
    if text[0] == '#' && body[0] == '#' { body = body[1:] }
    return "<a href=\"" + ToLink( text ) + "\" >" + body + "</a>"
  case 'm' : return "<span class=\"math\" >" + body + "</span>"
  case 'n' : return "<span class=\"note\" ><sup><a href=\"#" + ToLink(body) + "\" >" + body + "</a></sup></span>"
  case 'o' : return body
  case 'p' : return body
  case 'q' : return "<q>" + body + "</q>"
  case 'r' : return body // ref
  case 's' : return body
  case 't' : return "<span id=\"" + text + "\" >" + body + "</span>"
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
  re.Match( link, "#^:s*<:s*[^:s]+>+" )
  t := re.GetCatch( 1 )
  re.Match( t, "<:s>" )
  return re.RplCatch( "-", 1 )
}
