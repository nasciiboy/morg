package html

import (
  "bytes"
  "strings"

  "github.com/nasciiboy/txt"
  "github.com/nasciiboy/morg/katana"
)

func UnFontify( m katana.Markup ) (str string) {
  return ToSafeHtml( m.String() )
}

func Fontify( m katana.Markup ) (str string){
  if len( m.Data ) > 0 {
    if m.Type != 0 {
      safe := ToSafeHtml( m.Data )
      switch m.Type {
      case 'l', 'N', 'n', 't' :
        return AtCommand( safe, safe, m.Type )
      }
      return AtCommand( "", safe, m.Type )
    }

    return ToSafeHtml( m.Data )
  }

  var left, right bytes.Buffer
  for _, c := range m.Left  {  left.WriteString( Fontify( c ) ) }
  for _, c := range m.Right { right.WriteString( Fontify( c ) ) }

  if left.Len() == 0 {
    switch m.Type {
    case 'l', 'N', 'n', 't' :
      left.WriteString( ToSafeHtml( m.MakeLeft() ) )
    }
  }

  return AtCommand( left.String(), right.String(), m.Type )
}


func ToSafeHtml( s string ) string {
  if !strings.ContainsAny( s, "'\"&<>\000" ) { return s }

  return escapeHtml( s )
}

func escapeHtml( s string ) string {
  w, last, html := new(bytes.Buffer), 0, ""
  for i, l := 0, len( s ); i < l; i++ {
    switch s[i] {
    case 0   : html = "\uFFFD"
    case '"' : html = "&#34;"
    case '\'': html = "&#39;"
    case '&' : html = "&amp;"
    case '<' : html = "&lt;"
    case '>' : html = "&gt;"
    default  : continue
    }
    w.WriteString( s[last:i] )
    w.WriteString( html )
    last = i + 1
  }

  w.WriteString( s[last:] )
  return w.String()
}

func ToLink( link string ) string {
  return txt.SpaceSwap( link, "-" )
}

func AtCommand( left, right string, label byte ) string {
  switch label {
  case katana.MarkupNil, katana.MarkupEsc: return right
  case katana.MarkupErr: return `<span class="bad-markup" >` + right + "</span>"
  case '!' : return right
  case '"' : return "<q>" + right + "</q>"
  case '#' : return `<span class="path" >` + right + "</span>"
  case '$' : return `<code class="command" >` + right + "</code>"
  case '%' : return "(" + right + ")"
  case '&' : return atCommandSymbol( right ) // "simbol"
  case '\'': return "<samp>" + right + "</samp>"
  case '*' : return right
  case '+' : return right
  case ',' : return right
  case '-' : return "––" + right + "––"
  case '.' : return right
  case '/' : return right
  case '0' : return right
  case '1' : return right
  case '2' : return right
  case '3' : return right
  case '4' : return right
  case '5' : return right
  case '6' : return right
  case '7' : return right
  case '8' : return right
  case '9' : return right
  case ':' : return "<dfn>" + right + "</dfn>"
  case ';' : return right
  case '=' : return right
  case '?' : return right
  case 'A' : return `<span class="acronym" >` + right + "</span>"
  case 'B' : return right
  case 'C' : return right // "smallCaps"
  case 'D' : return right
  case 'E' : return right // "error"
  case 'F' : return right // "Func"
  case 'G' : return right
  case 'H' : return right
  case 'I' : return right
  case 'J' : return right
  case 'K' : return right // "keyword"
  case 'L' : return right // "label"
  case 'M' : return "\\(" + right + "\\)"
  case 'N' : return `<span class="defnote" id="` + ToLink(left) + `" >` + right + "</span>"
  case 'O' : return right
  case 'P' : return right
  case 'Q' : return right
  case 'R' : return right // "result"
  case 'S' : return right
  case 'T' : return right // "radiotarget"
  case 'U' : return right
  case 'V' : return right // "var"
  case 'W' : return right // "warning"
  case 'X' : return right
  case 'Y' : return right
  case 'Z' : return right
  case '\\': return right
  case '^' : return "<sup>" + right + "</sup>"
  case '_' : return "<sub>" + right + "</sub>"
  case '`' : return right
  case 'a' : return "<abbr>" + right + "</abbr>"
  case 'b' : return "<b>" + right + "</b>"
  case 'c' : return "<code>" + right + "</code>"
  case 'd' : return right // data
  case 'e' : return "<em>" + right + "</em>"
  case 'f' : return `<span class="file" >` + right + "</span>"
  case 'g' : return right
  case 'h' : return right
  case 'i' : return "<i>" + right + "</i>"
  case 'j' : return right
  case 'k' : return "<kbd>" + right + "</kbd>"
  case 'l' :
    if left != "" && left[0] == '#' && right != "" && right[0] == '#' { right = right[1:] }
    return `<a href="` + ToLink( left ) + `" >` + right + "</a>"
  case 'm' : return `<span class="math" >` + right + "</span>"
  case 'n' : return `<span class="note" ><sup><a href="#` + ToLink(left) + `" >` + right + "</a></sup></span>"
  case 'o' : return right
  case 'p' : return right
  case 'q' : return "<q>" + right + "</q>"
  case 'r' : return right // ref
  case 's' : return "<s>" + right + "</s>"
  case 't' : return `<span id="` + ToLink( ToSafeHtml( left ) ) + `" >` + right + "</span>"
  case 'u' : return "<u>" + right + "</u>"
  case 'v' : return `<code class="verbatim" >` + right + "</code>"
  case 'w' : return right
  case 'x' : return right
  case 'y' : return right
  case 'z' : return right
  case '|' : return right
  case '~' : return right
  }

  return right
}

func atCommandSymbol( sym string ) string {
  if s, ok := symbols[ sym ]; ok { return s }

  return `<span class="bad-symbol" >` + sym + "</span>"
}

var symbols = map[string]string  {
  "aa"              : `Å`,
  "aacute"          : `á`,
  "Aacute"          : `Á`,
  "acirc"           : `â`,
  "Acirc"           : `Â`,
  "acute"           : `´`,
  "aelig"           : `æ`,
  "AElig"           : `Æ`,
  "agrave"          : `à`,
  "Agrave"          : `À`,
  "alefsym"         : `ℵ`,
  "aleph"           : `ℵ`,
  "alpha"           : `α`,
  "Alpha"           : `Α`,
  "amp"             : `&`,
  "ang"             : `∠`,
  "angle"           : `∠`,
  "approx"          : `≈`,
  "aring"           : `å`,
  "Aring"           : `Å`,
  "asciicirc"       : `^`,
  "ast"             : `∗`,
  "asymp"           : `≈`,
  "atilde"          : `ã`,
  "Atilde"          : `Ã`,
  "auml"            : `ä`,
  "Auml"            : `Ä`,
  "bdquo"           : `„`,
  "because"         : `∵`,
  "beta"            : `β`,
  "Beta"            : `Β`,
  "beth"            : `ℶ`,
  "blacksmile"      : `☻`,
  "brvbar"          : `¦`,
  "bull"            : `•`,
  "bullet"          : `•`,
  "cap"             : `∩`,
  "ccedil"          : `ç`,
  "Ccedil"          : `Ç`,
  "cdot"            : `⋅`,
  "cdots"           : `⋯`,
  "cedil"           : `¸`,
  "cent"            : `¢`,
  "check"           : `✓`,
  "checkmark"       : `✓`,
  "chi"             : `χ`,
  "Chi"             : `Χ`,
  "circ"            : `ˆ`,
  "clubs"           : `♣`,
  "clubsuit"        : `♣`,
  "colon"           : `:`,
  "cong"            : `≅`,
  "copy"            : `©`,
  "crarr"           : `↵`,
  "cup"             : `∪`,
  "curren"          : `¤`,
  "dag"             : `†`,
  "dagger"          : `†`,
  "Dagger"          : `‡`,
  "dalet"           : `ℸ`,
  "darr"            : `↓`,
  "dArr"            : `⇓`,
  "ddag"            : `‡`,
  "deg"             : `°`,
  "delta"           : `δ`,
  "Delta"           : `Δ`,
  "diamond"         : `⋄`,
  "Diamond"         : `⋄`,
  "diamondsuit"     : `♦`,
  "diams"           : `♦`,
  "div"             : `÷`,
  "dots"            : `…`,
  "downarrow"       : `↓`,
  "Downarrow"       : `⇓`,
  "eacute"          : `é`,
  "Eacute"          : `É`,
  "ecirc"           : `ê`,
  "Ecirc"           : `Ê`,
  "egrave"          : `è`,
  "Egrave"          : `È`,
  "ell"             : `ℓ`,
  "empty"           : `∅`,
  "emptyset"        : `∅`,
  "epsilon"         : `ε`,
  "Epsilon"         : `Ε`,
  "equal"           : `=`,
  "equiv"           : `≡`,
  "eta"             : `η`,
  "Eta"             : `Η`,
  "eth"             : `ð`,
  "ETH"             : `Ð`,
  "euml"            : `ë`,
  "Euml"            : `Ë`,
  "EUR"             : `€`,
  "EURcr"           : `€`,
  "EURdig"          : `€`,
  "EURhv"           : `€`,
  "euro"            : `€`,
  "EURtm"           : `€`,
  "exist"           : `∃`,
  "exists"          : `∃`,
  "fnof"            : `ƒ`,
  "forall"          : `∀`,
  "frac12"          : `½`,
  "frac14"          : `¼`,
  "frac34"          : `¾`,
  "frasl"           : `⁄`,
  "frown"           : `⌢`,
  "gamma"           : `γ`,
  "Gamma"           : `Γ`,
  "ge"              : `≥`,
  "geq"             : `≥`,
  "gets"            : `←`,
  "gg"              : `≫`,
  "Gg"              : `⋙`,
  "ggg"             : `⋙`,
  "gimel"           : `ℷ`,
  "gt"              : `>`,
  "harr"            : `↔`,
  "hArr"            : `⇔`,
  "hbar"            : `ℏ`,
  "hearts"          : `♥`,
  "heartsuit"       : `♥`,
  "hellip"          : `…`,
  "hookleftarrow"   : `↵`,
  "iacute"          : `í`,
  "Iacute"          : `Í`,
  "icirc"           : `î`,
  "Icirc"           : `Î`,
  "iexcl"           : `¡`,
  "igrave"          : `ì`,
  "Igrave"          : `Ì`,
  "image"           : `ℑ`,
  "imath"           : `ı`,
  "in"              : `∈`,
  "inf"             : `∞`,
  "infin"           : `∞`,
  "infty"           : `∞`,
  "int"             : `∫`,
  "iota"            : `ι`,
  "Iota"            : `Ι`,
  "iquest"          : `¿`,
  "isin"            : `∈`,
  "iuml"            : `ï`,
  "Iuml"            : `Ï`,
  "jmath"           : `ȷ`,
  "kappa"           : `κ`,
  "Kappa"           : `Κ`,
  "lambda"          : `λ`,
  "Lambda"          : `Λ`,
  "land"            : `∧`,
  "lang"            : `⟨`,
  "laquo"           : `«`,
  "larr"            : `←`,
  "lArr"            : `⇐`,
  "lceil"           : `⌈`,
  "ldquo"           : `“`,
  "le"              : `≤`,
  "leftarrow"       : `←`,
  "Leftarrow"       : `⇐`,
  "leftrightarrow"  : `↔`,
  "Leftrightarrow"  : `⇔`,
  "leq"             : `≤`,
  "lesseqgtr"       : `⋚`,
  "lessgtr"         : `≶`,
  "lfloor"          : `⌊`,
  "ll"              : `≪`,
  "Ll"              : `⋘`,
  "lll"             : `⋘`,
  "lor"             : `∨`,
  "lowast"          : `∗`,
  "loz"             : `◊`,
  "lrm"             : `0`,
  "lsaquo"          : `‹`,
  "lsquo"           : `‘`,
  "lt"              : `<`,
  "macr"            : `¯`,
  "mdash"           : `—`,
  "mho"             : `℧`,
  "micro"           : `µ`,
  "middot"          : `·`,
  "minus"           : `−`,
  "mu"              : `μ`,
  "Mu"              : `Μ`,
  "nabla"           : `∇`,
  "ndash"           : `–`,
  "ne"              : `≠`,
  "neg"             : `¬`,
  "neq"             : `≠`,
  "nexist"          : `∃`,
  "nexists"         : `∃`,
  "ni"              : `∋`,
  "not"             : `¬`,
  "notin"           : `∉`,
  "nsub"            : `⊄`,
  "nsup"            : `⊅`,
  "ntilde"          : `ñ`,
  "Ntilde"          : `Ñ`,
  "nu"              : `ν`,
  "Nu"              : `Ν`,
  "oacute"          : `ó`,
  "Oacute"          : `Ó`,
  "ocirc"           : `ô`,
  "Ocirc"           : `Ô`,
  "odot"            : `o`,
  "oelig"           : `œ`,
  "OElig"           : `Œ`,
  "ograve"          : `ò`,
  "Ograve"          : `Ò`,
  "oline"           : `‾`,
  "omega"           : `ω`,
  "Omega"           : `Ω`,
  "omicron"         : `ο`,
  "Omicron"         : `Ο`,
  "oplus"           : `⊕`,
  "ordf"            : `ª`,
  "ordm"            : `º`,
  "oslash"          : `ø`,
  "Oslash"          : `Ø`,
  "otilde"          : `õ`,
  "Otilde"          : `Õ`,
  "otimes"          : `⊗`,
  "ouml"            : `ö`,
  "Ouml"            : `Ö`,
  "para"            : `¶`,
  "partial"         : `∂`,
  "permil"          : `‰`,
  "perp"            : `⊥`,
  "phi"             : `φ`,
  "Phi"             : `Φ`,
  "piv"             : `ϖ`,
  "pi"              : `π`,
  "Pi"              : `Π`,
  "plus"            : `+`,
  "plusmn"          : `±`,
  "pm"              : `±`,
  "pound"           : `£`,
  "prec"            : `≺`,
  "preccurlyeq"     : `≼`,
  "preceq"          : `≼`,
  "prime"           : `′`,
  "Prime"           : `″`,
  "prod"            : `∏`,
  "prop"            : `∝`,
  "propto"          : `∝`,
  "psi"             : `ψ`,
  "Psi"             : `Ψ`,
  "quot"            : `"`,
  "radic"           : `√`,
  "rang"            : `⟩`,
  "raquo"           : `»`,
  "rarr"            : `→`,
  "rArr"            : `⇒`,
  "rceil"           : `⌉`,
  "rdquo"           : `”`,
  "real"            : `ℜ`,
  "reg"             : `®`,
  "rfloor"          : `⌋`,
  "rho"             : `ρ`,
  "Rho"             : `Ρ`,
  "rightarrow"      : `→`,
  "Rightarrow"      : `⇒`,
  "rsaquo"          : `›`,
  "rsquo"           : `’`,
  "S"               : `§`,
  "sad"             : `☹`,
  "sbquo"           : `‚`,
  "scaron"          : `š`,
  "Scaron"          : `Š`,
  "sdot"            : `⋅`,
  "sect"            : `§`,
  "setminus"        : `∖`,
  "shy"             : `­`,
  "sigmaf"          : `ς`,
  "sigma"           : `σ`,
  "Sigma"           : `Σ`,
  "sim"             : `∼`,
  "simeq"           : `≅`,
  "slash"           : `/`,
  "smile"           : `⌣`,
  "smiley"          : `☺`,
  "spades"          : `♠`,
  "spadesuit"       : `♠`,
  "star"            : `*`,
  "sub"             : `⊂`,
  "sube"            : `⊆`,
  "subset"          : `⊂`,
  "succ"            : `≻`,
  "succcurlyeq"     : `≽`,
  "succeq"          : `≽`,
  "sum"             : `∑`,
  "sup"             : `⊃`,
  "sup1"            : `¹`,
  "sup2"            : `²`,
  "sup3"            : `³`,
  "supe"            : `⊇`,
  "supset"          : `⊃`,
  "szlig"           : `ß`,
  "tau"             : `τ`,
  "Tau"             : `Τ`,
  "there4"          : `∴`,
  "therefore"       : `∴`,
  "thetasym"        : `ϑ`,
  "theta"           : `θ`,
  "Theta"           : `Θ`,
  "thorn"           : `þ`,
  "THORN"           : `Þ`,
  "tilde"           : `~`,
  "times"           : `×`,
  "to"              : `→`,
  "trade"           : `™`,
  "triangleq"       : `≜`,
  "uacute"          : `ú`,
  "Uacute"          : `Ú`,
  "uarr"            : `↑`,
  "uArr"            : `⇑`,
  "ucirc"           : `û`,
  "Ucirc"           : `Û`,
  "ugrave"          : `ù`,
  "Ugrave"          : `Ù`,
  "uml"             : `¨`,
  "under"           : `_`,
  "uparrow"         : `↑`,
  "Uparrow"         : `⇑`,
  "upsih"           : `ϒ`,
  "upsilon"         : `υ`,
  "Upsilon"         : `Υ`,
  "uuml"            : `ü`,
  "Uuml"            : `Ü`,
  "varepsilon"      : `ε`,
  "varphi"          : `ϕ`,
  "varpi"           : `ϖ`,
  "varsigma"        : `ς`,
  "vartheta"        : `ϑ`,
  "vee"             : `∨`,
  "vert"            : `|`,
  "wedge"           : `∧`,
  "weierp"          : `℘`,
  "xi"              : `ξ`,
  "Xi"              : `Ξ`,
  "yacute"          : `ý`,
  "Yacute"          : `Ý`,
  "yen"             : `¥`,
  "yuml"            : `ÿ`,
  "Yuml"            : `Ÿ`,
  "zeta"            : `ζ`,
  "Zeta"            : `Ζ`,
}
