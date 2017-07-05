package html

import (
  "text/template"
  "bytes"
  "fmt"

  "github.com/nasciiboy/morg/katana"
  "github.com/nasciiboy/regexp3"
  "github.com/nasciiboy/pygments"
)

const HtmlTemplate =
`<?xml version="1.0" encoding="utf-8"?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" lang="{{ .Lang }}" xml:lang="{{ .Lang }}" >
  <head>
    <title>{{ UnFontify .Title }}</title>
    <meta  http-equiv="Content-Type" content="text/html;charset=utf-8" />
    {{ if hasSomething .Subtitle }}<meta name="subtitle"    content="{{ UnFontify .Subtitle }}" />{{ end }}
    {{ range .Author     }}<meta name="author"      content="{{ .                }}" />{{ end }}
    {{ range .Translator }}<meta name="translator"  content="{{ .                }}" />{{ end }}
    {{ if .Licence       }}<meta name="licence"     content="{{ .Licence         }}" />{{ end }}
    {{ if .Id            }}<meta name="id"          content="{{ .Id              }}" />{{ end }}
    {{ if .Date          }}<meta name="date"        content="{{ .Date            }}" />{{ end }}
    {{ if .Tags          }}<meta name="tags"        content="{{ .Tags            }}" />{{ end }}
    {{ if .Description   }}<meta name="description" content="{{ .Description     }}" />{{ end }}
    {{ if .Mail          }}<meta name="mail"        content="{{ .Mail            }}" />{{ end }}

    {{ range .Style    }}<link rel="stylesheet" type="text/css" href="{{ . }}" />{{ end }}

    {{ if .OptionsData.Highlight }}
    <link rel="stylesheet" href="highlight/styles/atelier-forest-dark.css" />
    <script src="highlight/highlight.pack.js" ></script>
    <script>hljs.initHighlightingOnLoad();</script>
    {{ end }}
    {{ if .OptionsData.Mathjax }}
    <script type="text/javascript" async
      src="https://cdnjs.cloudflare.com/ajax/libs/mathjax/2.7.1/MathJax.js?config=TeX-MML-AM_CHTML">
    </script>
    {{ end }}
  </head>
  <body>


{{ if .OptionsData.Toc }}
<div id="toc">
  <p>index</p>
  <div id="toc-contents">
  {{ ToToc .Toc .OptionsData }}
  </div>
</div>
{{ end }}

{{ if hasSomething .Title }}<h1>{{ Fontify .Title }}</h1>{{ end }}

{{ ToBody .Toc .OptionsData }}

  </body>
</html>
`

func MakeHtml( str string ) string {
  var document katana.Doc
  document.Parse( str )

  document.OptionsData.HShift = 1

	t, err := template.New("HtmlTemplate").
    Funcs(template.FuncMap{ "ToBody"       : ToBody }).
    Funcs(template.FuncMap{ "ToToc"        : ToToc  }).
    Funcs(template.FuncMap{ "Fontify"      : fontify  }).
    Funcs(template.FuncMap{ "UnFontify"    : unfontify  }).
    Funcs(template.FuncMap{ "hasSomething" : hasSomething  }).
    Parse(HtmlTemplate)
  if err != nil { panic( err ) }

  out := bytes.Buffer{}
  err = t.Execute( &out, document);
  if err != nil { panic( err ) }

  return out.String()
}

func MakeHtmlBody( str string ) string {
  var document katana.Doc
  document.Parse( str )

  return ToBody( document.Toc, document.OptionsData )
}

func hasSomething( m katana.Markup ) bool {
  return m.HasSomething()
}

func fontify( m katana.Markup ) (str string) {
  if len( m.Custom ) == 0 && len( m.Body ) == 0 {
    return ToSafeHtml( m.Data )
  }

  var custom, body string

  for _, c := range m.Custom {
    custom += fontify( c )
  }

  for _, c := range m.Body {
    body   += fontify( c )
  }

  if custom == "" {
    switch m.Type {
    case 'l', 'N', 'n', 't' :
      custom = ToSafeHtml( m.MakeCustom() )
    }
  }

  return AtCommand( body, custom, m.Type )
}

func unfontify( m katana.Markup ) (str string) {
  return m.String()
}

func ToToc( toc []katana.DocNode, options katana.Options ) (str string) {
  for _, h := range( toc[1:] ) {
    full := h.Get()
    str += fmt.Sprintf( "<span class=\"toc\" ><a class=\"h%d\" href=\"#%s\" >%s</a></span>\n",
      full.N + options.HShift,
      ToLink( ToSafeHtml( full.Mark.MakeCustom() )),
      fontify( full.Mark ) )
  }

  return str
}

func ToBody( toc []katana.DocNode, options katana.Options ) (str string) {
  for _, h := range( toc ) {
    full := h.Get()
    if full.N == 0 {
    } else {
      str += fmt.Sprintf( "<h%d id=\"%s\" >%s</h%[1]d>\n",
      full.N + options.HShift,
      ToLink( ToSafeHtml( full.Mark.MakeCustom() )),
      fontify( full.Mark ) )
    }

    if len( h.Cont ) > 0 {
      str += fmt.Sprintf( "<div class=\"hBody-%d\" >\n", full.N + options.HShift )
      str += walkContent( h.Cont, options )
      str += "</div>\n"
    }
  }

  return str
}

func walkContent( doc []katana.DocNode, options katana.Options ) (str string) {
  for _, c := range doc {
    full := c.Get()
    switch c.Type() {
    case katana.EmptyNode     :
    case katana.CommentNode   :
    case katana.CommandNode   : str += makeCommand( full, c.Cont, options )
    case katana.HeadlineNode  :
    case katana.TableNode     : str += makeTable  ( full, c.Cont, options )
    case katana.ListNode      : str += makeList   ( full, c.Cont, options )
    case katana.AboutNode     : str += makeAbout  ( full, c.Cont, options )
    case katana.TextNode      :
      str += "<p>"
      str += fontify( full.Mark )
      str += "</p>\n"
    }
  }

  return str
}

func makeAbout( data katana.FullData, cont []katana.DocNode, options katana.Options ) (str string) {
  str += "<div class=\"about\" >\n"
  str += "<div class=\"about-dt\" >" + fontify( data.Mark ) + "</div>\n"
  str += "<div class=\"about-dd\" >\n"
  str += walkContent( cont, options )
  str += "</div>\n"
  str += "</div>\n"

  return
}

func makeList( data katana.FullData, cont []katana.DocNode, options katana.Options ) (str string) {
  switch data.N {
  case katana.ListMinusNode, katana.ListPlusNode :
    str += "<ul>\n"
    str += makeListNodes( cont, options )
    str += "</ul>\n"
  case katana.ListNumNode,   katana.ListAlphaNode:
    str += "<ol>\n"
    str += makeListNodes( cont, options )
    str += "</ol>\n"
  case katana.ListMdefNode,  katana.ListPdefNode :
    str += "<dl>\n"
    str += makeDlListNodes( cont, options )
    str += "</dl>\n"
  case katana.ListDialogNode:
    str += "<ul class=\"dialog\" >\n"
    str += makeListNodes( cont, options )
    str += "</ul>\n"
  }

  return
}

func makeListNodes( cont []katana.DocNode, options katana.Options ) (str string) {
  for _, element := range( cont ) {
    str += "<li>\n"
    str += walkContent( element.Cont, options )
    str += "</li>\n"
  }

  return
}

func makeDlListNodes( cont []katana.DocNode, options katana.Options ) (str string) {
  for _, element := range( cont ) {
    full := element.Get()

    str += "<dt>"
    str += fontify( full.Mark )
    str += "</dt>\n"

    str += "<dd>\n"
    str += walkContent( element.Cont, options )
    str += "</dd>\n"
  }

  return
}

func makeTable( data katana.FullData, rows []katana.DocNode, options katana.Options ) (str string) {
  str += "<table>\n"

  for i, row := range( rows ) {
    d := row.Get()

    if i == 0  {
      if d.N == katana.TableHead {
        str += "<thead>\n"
      } else { break }
    }

    if i > 0 && d.N != katana.TableHead {
      str += "</thead>\n"
      rows = rows[i:]
      break
    }

    str += makeTableRow( d.N, row.Cont )
  }

  if len( rows ) > 0 {
    str += "<tbody>\n"

    for _, row := range( rows ) {
      d := row.Get()
      str += makeTableRow( d.N, row.Cont )
    }

    str += "</tbody>\n"
  }

  str += "</table>\n"
  return
}

func makeTableRow( Type int, cells []katana.DocNode ) (str string) {
  str += "<tr>"

  for _, cell := range( cells ) {
    switch Type {
    case katana.TableHead: str += "<th>"
    case katana.TableBody: str += "<td>"
    case katana.TableFoot: str += "<td>"
    }

    data := cell.Get()
    str += fontify( data.Mark )

    switch Type {
    case katana.TableHead: str += "</th>"
    case katana.TableBody: str += "</td>"
    case katana.TableFoot: str += "</td>"
    }
  }

  str += "</tr>\n"
  return
}

func makeCommand( data katana.FullData, cont []katana.DocNode, options katana.Options ) string {
  switch data.Comm {
  case "src"    : return makeCommandSrc   ( data, options )
  case "srci"   : return makeCommandSrci  ( data, cont, options )
  case "figure" : return makeCommandFigure( data, cont, options )
  case "cols"   : return makeCommandCols  ( data, cont, options )
  case "img"    : return makeCommandImg   ( data, cont, options )
  case "video"  : return makeCommandVideo ( data, cont, options )
  case "quote"  : return makeCommandQuote ( data, cont, options )
  case "example", "pre":
    return makeCommandPre( data, options )
  case "diagram", "art":
    return makeCommandArt( data, options )
  case "center", "bold", "emph", "verse", "tab", "italic":
    return makeCommandFont( data, cont, options )
  case "pret":
    return makeCommandPret( data, cont, options )
  case "math":
    return makeCommandMath( data, cont, options )
  }

  return ""
}

func makeCommandSrc( comm katana.FullData, options katana.Options ) (str string) {
  if options.Pygments {
    pygCode, make := pygments.Highlight( comm.Data, comm.Arg, "html", "utf-8" )

    if make == false { goto simple }
    return pygCode
  }

simple:
  str += fmt.Sprintf( "<pre class=\"code\" ><code class=\"%s\">", comm.Arg )
  str += ToSafeHtml( comm.Data )
  str += "</code></pre>\n"
  return
}

func makeCommandSrci( comm katana.FullData, body []katana.DocNode, options katana.Options ) (str string) {
  str += fmt.Sprintf( "<pre class=\"srci\" ><code class=\"%s\">", comm.Arg )

  for _, c := range( body ) {
    d := c.Get()

    if d.N == katana.TextCode {
      str += makeSrciCode( d.Mark.Data, comm.Arg, options )
    } else {
      str += "<span class=\"srci-text\">" + ToSafeHtml( d.Mark.Data ) + "</span>\n"
    }
  }

  str += "</code></pre>\n"
  return
}

func makeSrciCode( code, lang string, options katana.Options ) string {
  if options.Pygments {
    pygCode, make := pygments.HighlightNoWrap( code, lang, "html", "utf-8" )

    if make == false { return ToSafeHtml( code ) + "\n" }
    return pygCode
  }

  return ToSafeHtml( code ) + "\n"
}

func makeCommandPre( comm katana.FullData, options katana.Options ) (str string) {
  str += fmt.Sprintf( "<div class=\"%s-block\" >\n", comm.Comm )
  str += fmt.Sprintf( "<pre class=\"%s\" >", comm.Comm )
  str += ToSafeHtml( comm.Data )
  str += "</pre></div>\n"
  return
}

func makeCommandFigure( comm katana.FullData, body []katana.DocNode, options katana.Options ) (str string) {
  str += "<div class=\"figure\" >\n"
  str += "<p class=\"title\">" + fontify( comm.Mark ) + "</p>\n"
  str += walkContent( body, options )
  str += "</div>\n"

  return str
}

func makeCommandCols( comm katana.FullData, body []katana.DocNode, options katana.Options ) (str string) {
  str += "<div class=\"cols\" style=\"width: 100%; display: inline-flex; flex-flow: row nowrap; flex-direction: row; \">\n"

  width := 100
  if len(body) != 0 { width = 100 / len(body) }
  for i, c := range( body ) {
    str += "<div class=\"cols-element\" "
    str += fmt.Sprintf( "style=\" order: %d; width: %d%%; \">\n", i + 1, width )
    str += walkContent( c.Cont, options )
    str += "</div>\n"
  }
  str += "</div>\n"

  return
}

func makeCommandImg( comm katana.FullData, body []katana.DocNode, options katana.Options ) (str string) {
  str += "<figure>\n"
  str += "<img src=\"" + comm.Arg + "\" />\n"

  if len( body ) != 0 {
    str += "<figcaption>\n"
    str += walkContent( body, options )
    str += "</figcaption>\n"
  }

  str += "</figure>\n"
  return
}

func makeCommandVideo( comm katana.FullData, body []katana.DocNode, options katana.Options ) (str string) {
  str += "<div class=\"video\">\n"
  str += "<video controls >\n"
  str += "<source src=\"" + comm.Arg + "\""

  var re regexp3.RE
  if re.Match( comm.Arg, "#$:.<ogg|mp4>" ) > 0 {
    str +=  " type=\"video/" + re.GetCatch( 1 ) + "\""
  }

  str += " >\n"
  str += "Your browser does not support HTML5 video\n"
  str += "</video>\n"
  str += walkContent( body, options )
  str += "</div>\n"
  return
}

func makeCommandFont( comm katana.FullData, body []katana.DocNode, options katana.Options ) (str string) {
  str += "<div class=\"" + comm.Comm + "\" >\n"
  str += walkContent( body, options )
  str += "</div>\n"
  return
}

func makeCommandMath( comm katana.FullData, body []katana.DocNode, options katana.Options ) (str string) {
  str += "<div class=\"mathjax\" >\n"
  str += "$$" + comm.Data + "$$"
  str += "</div>\n"
  return
}

func makeCommandArt( comm katana.FullData, options katana.Options ) (str string) {
  str += fmt.Sprintf( "<div class=\"%s-block\" >\n", comm.Comm )
  str += fmt.Sprintf( "<pre class=\"%s\" >", comm.Comm )
  str += ToSafeHtml( comm.Data )
  str += "</pre>\n"
  str += "</div>\n"
  return
}

func makeCommandQuote( comm katana.FullData, body []katana.DocNode, options katana.Options ) (str string) {
  str += "<blockquote>\n"

  for _, c := range( body ) {
    nodeData := c.Get()
    switch nodeData.N {
    case katana.TextQuoteAuthor:
      str += "<div class=\"quote-author\" >\n"
      str += "<p>";
      str += fontify( nodeData.Mark )
      str += "</p>\n";
      str += "</div>\n"
    case katana.TextSimple:
      str += "<p>";
      str += fontify( nodeData.Mark );
      str += "</p>\n";
    }
  }

  str += "</blockquote>\n"
  return
}

func makeCommandPret( comm katana.FullData, body []katana.DocNode, options katana.Options ) (str string) {
  str += "<div class=\"pret\" >\n"

  for _, c := range( body ) {
    nodeData := c.Get()
    str += fontify( nodeData.Mark );
    str += "<br>\n";
  }

  str += "</div>\n"
  return
}
