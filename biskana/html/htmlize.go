package html

import (
  "github.com/nasciiboy/utils/text"
  "github.com/nasciiboy/morg/katana"
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

func ToLink( link string ) string {
  return text.SpaceSwap( link, "-" )
}

func AtCommand( body, custom string, label byte ) string {
  switch label {
  case katana.MarkupNil, katana.MarkupEsc, katana.MarkupErr: return body
  case katana.MarkupHeadline, katana.MarkupTitle, katana.MarkupSubTitle,
       katana.MarkupList, katana.MarkupDialog, katana.MarkupComment,
       katana.MarkupAbout, katana.MarkupCode, katana.MarkupText: return body
  case '!' : return body
  case '"' : return "<q>" + body + "</q>"
  case '#' : return `<span class="path" >` + body + "</span>"
  case '$' : return `<code class="command" >` + body + "</code>"
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
  case 'A' : return `<span class="acronym" >` + body + "</span>"
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
  case 'N' : return `<span class="defnote" id="` + ToLink(custom) + `" >` + body + "</span>"
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
  case 'f' : return `<span class="file" >` + body + "</span>"
  case 'g' : return body
  case 'h' : return body
  case 'i' : return "<i>" + body + "</i>"
  case 'j' : return body
  case 'k' : return "<kbd>" + body + "</kbd>"
  case 'l' :
    if custom != "" && custom[0] == '#' && body != "" && body[0] == '#' { body = body[1:] }
    return `<a href="` + ToLink( custom ) + `" >` + body + "</a>"
  case 'm' : return `<span class="math" >` + body + "</span>"
  case 'n' : return `<span class="note" ><sup><a href="#` + ToLink(custom) + `" >` + body + "</a></sup></span>"
  case 'o' : return body
  case 'p' : return body
  case 'q' : return "<q>" + body + "</q>"
  case 'r' : return body // ref
  case 's' : return body
  case 't' : return `<span id="` + custom + `" >` + body + "</span>"
  case 'u' : return "<u>" + body + "</u>"
  case 'v' : return `<code class="verbatim" >` + body + "</code>"
  case 'w' : return body
  case 'x' : return body
  case 'y' : return body
  case 'z' : return body
  case '|' : return body
  case '~' : return body
  }

  return body
}
