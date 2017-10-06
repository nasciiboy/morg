package html

import (
  "bytes"
  "fmt"
  "strconv"

  "github.com/nasciiboy/morg/katana"
  "github.com/nasciiboy/txt"
)

func MakeHtml( doc *katana.Doc ) string {
  w := new( bytes.Buffer )
  w.Grow( 65536 )

  w.WriteString( `<?xml version="1.0" encoding="utf-8"?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd" />
<html xmlns="http://www.w3.org/1999/xhtml" ` )

  if doc.Lang != "" {
    w.WriteString( `lang="` + doc.Lang + `" xml:lang="` + doc.Lang + "\" >\n" )
  } else {
    w.WriteString( "lang=\"en\" xml:lang=\"en\" >\n" )
  }

  w.WriteString ( "<head>\n" )
  if doc.Title.HasSomething() { w.WriteString ( "  <title>" + UnFontify( doc.Title ) + "</title>\n" ) }
  w.WriteString ( `  <meta http-equiv="Content-Type" content="text/html;charset=utf-8" />` + "\n" )
  if doc.Subtitle.HasSomething() {
    w.WriteString( `  <meta name="subtitle" content="` + UnFontify( doc.Subtitle ) + "\" />\n" )
  }
  for _, author := range doc.Author {
    w.WriteString( `  <meta name="author" content="` + ToSafeHtml( author ) + "\" />\n" )
  }
  for _, translator := range doc.Translator {
    w.WriteString( `  <meta name="translator" content="` + ToSafeHtml( translator ) + "\" />\n" )
  }
  for _, source := range doc.Source {
    w.WriteString( `  <meta name="source" content="` + ToSafeHtml( source ) + "\" />\n" )
  }
  if doc.Licence != "" {
    w.WriteString( `  <meta name="licence" content="` + ToSafeHtml( doc.Licence ) + "\" />\n" )
  }
  if doc.ID != "" {
    w.WriteString( `  <meta name="id" content="` + ToSafeHtml( doc.ID ) + "\" />\n" )
  }
  if doc.Date != "" {
    w.WriteString( `  <meta name="date" content="` + ToSafeHtml( doc.Date ) + "\" />\n" )
  }
  if len( doc.Tags ) != 0 {
    w.WriteString( `  <meta name="keywords" content="` + ToSafeHtml( doc.Tags[ 0 ] ) )
    for _, tag := range doc.Tags[ 1: ] {
      w.WriteString( "," + ToSafeHtml( tag ) )
    }
    w.WriteString( "\" />\n" )
  }
  if doc.Description != "" {
    w.WriteString( `  <meta name="description" content="` + ToSafeHtml( doc.Description ) + "\" />\n" )
  }
  for _, style := range doc.Style {
    w.WriteString( `  <link rel="stylesheet" type="text/css" href="` + ToSafeHtml( style ) + "\" />\n" )
  }
  if doc.TextOptions[ "fancyCode" ] != "" { WriteStyle( w, doc.TextOptions[ "fancyCode" ] ) }

  if doc.BoolOptions[ "mathJax" ] {
    w.WriteString( `  <script type="text/javascript" async
    src="https://cdnjs.cloudflare.com/ajax/libs/mathjax/2.7.1/MathJax.js?config=TeX-MML-AM_CHTML">
  </script>` + "\n" )
  }

  w.WriteString( "</head>\n\n<body>\n" )

  doc.HShift = 1
  if doc.BoolOptions[ "toc" ] { writeToc( w, doc ) }

  if doc.Title.HasSomething() { w.WriteString( "<h1>" + Fontify( doc.Title ) + "</h1>\n" ) }

  w.WriteString( MakeHtmlBody( doc ) )

  w.WriteString( "</body>\n</html>\n" )

  return w.String()
}

func writeToc( w *bytes.Buffer, doc *katana.Doc ){
  if len( doc.Toc ) <= 1 { return }
  w.WriteString( "<div id=\"toc\">\n  <p>index</p>\n  <div id=\"toc-contents\">\n" )

  index := 1
  var toToc func( level int )
  toToc = func( level int ){
    spcs := fmt.Sprintf( "%*.s", level * 2, "" )
    fmt.Fprintf( w, "%s<ul>\n", spcs )

    for run := true; run;  {
      if index >= len( doc.Toc ) { break }
      h := doc.Toc[index].Node.(katana.Headline)
      if h.Level > level {
        toToc( level + 1 )
        continue
      } else if h.Level < level { break }

      if index++; index >= len( doc.Toc ) { run = false }
      fmt.Fprintf( w,  "%s<li><a class=\"h%d\" href=\"#%s\" >%s</a></li>\n",
        spcs, h.Level + doc.HShift, ToLink( ToSafeHtml( h.Mark.MakeLeft() )), Fontify( h.Mark ) )
    }

    fmt.Fprintf( w, "%s</ul>\n", spcs )
  }

  toToc( 1 )
  w.WriteString( "  </div>\n</div>\n\n" )
}

func MakeHtmlBody( doc *katana.Doc ) string {
  w := new( bytes.Buffer )
  w.Grow( 65536 )

  makeHtmlBody( w, doc )

  return w.String()
}

func makeHtmlBody( w *bytes.Buffer, doc *katana.Doc ){
  for _, docNode := range doc.Toc {
    headline := docNode.Node.(katana.Headline)
    if headline.Level == 0 {
    } else {
      fmt.Fprintf( w, "<h%d id=\"%s\" >%s</h%[1]d>\n",
        headline.Level + doc.HShift, ToLink( ToSafeHtml( headline.Mark.MakeLeft() )), Fontify( headline.Mark ) )
    }

    if len( docNode.Cont ) > 0 {
      fmt.Fprintf( w, "<div class=\"hBody-%d\" >\n", headline.Level + doc.HShift )
      walkContent( w, docNode.Cont, doc )
      w.WriteString( "</div>\n" )
    }
  }
}

// func makeHtmlBody( w *bytes.Buffer, doc *katana.Doc ){
//   out := make( []*bytes.Buffer, len(doc.Toc) )

//   done := make(chan struct{})
//   for n, sec := range doc.Toc {
//     go func( i int, docNode katana.DocNode ){
//       out[i] = new(bytes.Buffer)
//       w := out[i]
//       w.Grow( 4096 )

//       headline := docNode.Node.(katana.Headline)
//       if headline.Level == 0 {
//       } else {
//         fmt.Fprintf( w, "<h%d id=\"%s\" >%s</h%[1]d>\n",
//           headline.Level + doc.HShift, ToLink( ToSafeHtml( headline.Mark.MakeLeft() )), Fontify( headline.Mark ) )
//       }

//       if len( docNode.Cont ) > 0 {
//         fmt.Fprintf( w, "<div class=\"hBody-%d\" >\n", headline.Level + doc.HShift )
//         walkContent( w, docNode.Cont, doc )
//         w.WriteString( "</div>\n" )
//       }

//       done <- struct{}{}
//     }( n, sec )
//   }

//   for range doc.Toc { <-done }
//   for _, sw := range out {
//     w.ReadFrom( sw )
//   }
// }

func walkContent( w *bytes.Buffer, cont []katana.DocNode, doc *katana.Doc ) {
  for _, docNode := range cont {
    switch docNode.NodeType() {
    case katana.NodeEmpty     :
    case katana.NodeComment   :
    case katana.NodeTable     : makeTable  ( w, docNode.Node.(katana.Table   ), docNode.Cont, doc )
    case katana.NodeList      : makeList   ( w, docNode.Node.(katana.List    ), docNode.Cont, doc )
    case katana.NodeAbout     : makeAbout  ( w, docNode.Node.(katana.About   ), docNode.Cont, doc )
    case katana.NodeCode      : makeCode   ( w, docNode.Node.(katana.Code    ), doc.BoolOptions[ "fancyCode" ], doc )
    case katana.NodeSrci      : makeSrci   ( w, docNode.Node.(katana.Srci    ), doc.BoolOptions[ "fancyCode" ], doc )
    case katana.NodeFigure    : makeFigure ( w, docNode.Node.(katana.Figure  ), docNode.Cont, doc )
    case katana.NodeMedia     : makeMedia  ( w, docNode.Node.(katana.Media   ), docNode.Cont, doc )
    case katana.NodeWrap      : makeWrap   ( w, docNode.Node.(katana.Wrap    ), docNode.Cont, doc )
    case katana.NodeColumns   : makeColumns( w, docNode.Node.(katana.Columns ), docNode.Cont, doc )
    case katana.NodePret      : makePret   ( w, docNode.Node.(katana.Pret    ), doc )
    case katana.NodeBrick     : makeBrick  ( w, docNode.Node.(katana.Brick   ), doc )
    case katana.NodeQuote     : makeQuote  ( w, docNode.Node.(katana.Quote   ), doc )
    case katana.NodeText      :
      w.WriteString( "<p>" )
      w.WriteString( Fontify( docNode.Node.(katana.Markup) ) )
      w.WriteString( "</p>\n" )
    }
  }
}

func makeAbout( w *bytes.Buffer, about katana.About, cont []katana.DocNode, doc *katana.Doc ) {
  w.WriteString( "<div class=\"about\" >\n" )
  fmt.Fprintf( w, "<div class=\"about-dt\" >%s</div>\n", Fontify( about.Mark ) )
  if len( cont ) > 0 {
    w.WriteString( "<div class=\"about-dd\" >\n" )
    walkContent( w, cont, doc )
    w.WriteString( "</div>\n" )
  }
  w.WriteString( "</div>\n" )
}

func makeList( w *bytes.Buffer, list katana.List, cont []katana.DocNode, doc *katana.Doc ){
  switch list.ListType {
  case katana.NodeListMinus, katana.NodeListPlus:
    w.WriteString( "<ul>\n" )
    makeListNodes( w, cont, doc )
    w.WriteString( "</ul>\n" )
  case katana.NodeListNum:
    w.WriteString( "<ol class=\"num\" >\n" )
    makeListNodes( w, cont, doc )
    w.WriteString( "</ol>\n" )
  case katana.NodeListAlpha:
    w.WriteString( "<ol class=\"alpha\" >\n" )
    makeListNodes( w, cont, doc )
    w.WriteString( "</ol>\n" )
  case katana.NodeListMDef, katana.NodeListPDef:
    w.WriteString( "<dl>\n" )
    makeDlListNodes( w, cont, doc )
    w.WriteString( "</dl>\n" )
  case katana.NodeListDialog:
    w.WriteString( "<ul class=\"dialog\" >\n" )
    makeListNodes( w, cont, doc )
    w.WriteString( "</ul>\n" )
  }
}

func makeListNodes( w *bytes.Buffer, cont []katana.DocNode, doc *katana.Doc ){
  for _, element := range cont {
    w.WriteString( "<li>\n" )
    walkContent( w, element.Cont, doc )
    w.WriteString( "</li>\n" )
  }

  return
}

func makeDlListNodes( w *bytes.Buffer, cont []katana.DocNode, doc *katana.Doc ){
  for _, element := range( cont ) {
    fmt.Fprintf( w, "<dt>%s</dt>\n", Fontify( element.Node.(katana.ListElement).Mark ) )

    w.WriteString( "<dd>\n" )
    walkContent( w, element.Cont, doc )
    w.WriteString( "</dd>\n" )
  }
}

func makeTable( w *bytes.Buffer, table katana.Table, cont []katana.DocNode, doc *katana.Doc ) {
  w.WriteString( "<table>\n" )
  i := 0
  for ;i < len( cont ); i++ {
    d := cont[i].Node.(katana.TableRow)

    if i == 0  {
      if d.Type == katana.TableHead {
        w.WriteString( "<thead>\n" )
      } else { break }
    }

    if d.Type != katana.TableHead {
      w.WriteString( "</thead>\n" )
      break
    }

    makeTableRow( w, d.Type, cont[i].Cont, doc )
  }

  if i < len( cont ) {
    w.WriteString( "<tbody>\n" )
    for _, row := range cont[i:] {
      d := row.Node.(katana.TableRow)
      makeTableRow( w, d.Type, row.Cont, doc )
    }
    w.WriteString( "</tbody>\n" )
  }

  w.WriteString( "</table>\n" )
}

func makeTableRow( w *bytes.Buffer, Type int, cells []katana.DocNode, doc *katana.Doc ){
  w.WriteString( "<tr>" )
  for _, cell := range cells {
    switch Type {
    case katana.TableHead: fmt.Fprintf( w, "<th>%s</th>", Fontify( cell.Node.(katana.TableCell).Mark ) )
    case katana.TableBody: fmt.Fprintf( w, "<td>%s</td>", Fontify( cell.Node.(katana.TableCell).Mark ) )
    case katana.TableFoot: fmt.Fprintf( w, "<td>%s</td>", Fontify( cell.Node.(katana.TableCell).Mark ) )
    }
  }
  w.WriteString( "</tr>\n" )
}

func makeCode( w *bytes.Buffer, c katana.Code, fancyCode bool, doc *katana.Doc ) {
  if fancyCode {
    if c.Style == doc.TextOptions[ "fancyCode" ] { c.Style = "" }
    FancyCode( w, c )
    return
  }

  if c.Indexed {
    nCode( w, c );
    return
  }

  fmt.Fprintf( w, "<pre class=\"code\" ><code class=\"%s\">%s</code></pre>\n", c.Lang, ToSafeHtml( c.Body )  )
}

func nCode( w *bytes.Buffer, c katana.Code ) {
  fmt.Fprintf( w, "<pre class=\"code\" ><code class=\"%s\">", c.Lang  )

  lines := txt.GetRawLines( c.Body )
  width := len( strconv.Itoa( len( lines ) ) )
  n     := c.IndexNum
  if n < 0 && n + len( lines ) > 0 { width++ }
  for _, line := range lines {
    fmt.Fprintf( w, "<span class=\"index\" >%*d  </span>%s", width, n, ToSafeHtml( line ) )
    n++
  }

  fmt.Fprintf( w, "<pre class=\"code\" ><code class=\"%s\">%s</code></pre>\n", c.Lang, ToSafeHtml( c.Body )  )
}

func makeSrci( w *bytes.Buffer, srci katana.Srci, fancyCode bool, doc *katana.Doc ) {
  fmt.Fprintf( w, "<pre class=\"srci\" ><code class=\"%s\" >", srci.Lang  )
  for _, binary := range srci.Body {
    if binary.On {
      if fancyCode {
        style := srci.Style
        if style == doc.TextOptions[ "fancyCode" ] { style = "" }
        fmt.Fprintf( w, "<span class=\"in\" ><span class=\"prompt\" >%s</span>", ToSafeHtml( srci.Prompt ) )
        FancySrci( w, binary.Str, srci.Lang, style )
        fmt.Fprintf( w, "</span>" )
      } else {
        fmt.Fprintf( w, "<span class=\"in\" ><span class=\"prompt\" >%s</span>%s</span>", ToSafeHtml( srci.Prompt ), ToSafeHtml( binary.Str ) )
      }
    } else {
      fmt.Fprintf( w, "<span class=\"out\" >%s</span>", ToSafeHtml( binary.Str ) )
    }
  }

  fmt.Fprintf( w, "</code></pre>\n" )
}

func makeBrick( w *bytes.Buffer, brick katana.Brick, doc *katana.Doc ){
  switch brick.Type {
  case "math"   : makeMath  ( w, brick, doc )
  default:
    fmt.Fprintf( w, "<div class=\"%s-block\" >\n", brick.Type )
    fmt.Fprintf( w, "<pre class=\"%s\" >", brick.Type )
    w.WriteString( ToSafeHtml( brick.Body ) )
    w.WriteString( "</pre>\n</div>\n" )
  }
}

func makeMath( w *bytes.Buffer, brick katana.Brick, doc *katana.Doc ){
  w.WriteString( "<div class=\"mathjax\" >\n$$" )
  w.WriteString( txt.RmSpacesToTheSides( brick.Body ) )
  w.WriteString( "$$\n</div>\n" )
}

func makeFigure( w *bytes.Buffer, fig katana.Figure, cont []katana.DocNode, doc *katana.Doc ) {
  fmt.Fprintf( w, "<div class=\"figure\" >\n<h1 class=\"figure\">%s</h1>\n", Fontify( fig.Title ) )
  walkContent( w, cont, doc )
  fmt.Fprintf( w, "</div>\n" )
}

func makeColumns( w *bytes.Buffer, columns katana.Columns, cont []katana.DocNode, doc *katana.Doc ){
  fmt.Fprintf( w, "<div class=\"cols\" style=\"width: 100%%; display: inline-flex; flex-flow: row nowrap; flex-direction: row; \">\n" )

  width := 100
  if len(cont) != 0 { width = 100 / len(cont) }
  for i, c := range cont  {
    fmt.Fprintf( w, "<div class=\"cols-element\" style=\" order: %d; width: %d%%; \">\n", i + 1, width )
    walkContent( w, c.Cont, doc )
    fmt.Fprintf( w, "</div>\n" )
  }
  fmt.Fprintf( w, "</div>\n" )
}

func makeMedia( w *bytes.Buffer, media katana.Media, cont []katana.DocNode, doc *katana.Doc ){
  switch media.Type {
  case "img"   : makeImg( w, media, cont, doc )
  case "video" : makeVideo( w, media, cont, doc )
  case "audio" : makeAudio( w, media, cont, doc )
  }
}

func makeImg( w *bytes.Buffer, media katana.Media, cont []katana.DocNode, doc *katana.Doc ){
  fmt.Fprintf( w, "<figure>\n" )
  fmt.Fprintf( w, "<img src=\"%s\" />\n", media.Src )

  if len( cont ) > 0 {
    fmt.Fprintf( w, "<figcaption>\n" )
    walkContent( w, cont, doc )
    fmt.Fprintf( w, "</figcaption>\n" )
  }

  fmt.Fprintf( w, "</figure>\n" )
}

func makeVideo( w *bytes.Buffer, media katana.Media, cont []katana.DocNode, doc *katana.Doc ){
  fmt.Fprintf( w, "<div class=\"video\">\n<video controls >\n" )
  fmt.Fprintf( w, "<source src=\"%s\" type=\"video/%s\" >\n",  media.Src, media.Ext )
  fmt.Fprintf( w, "Your browser does not support <em>.%s</em> video\n", media.Ext )
  fmt.Fprintf( w, "</video>\n" )
  walkContent( w, cont, doc )
  fmt.Fprintf( w, "</div>\n" )
}

func makeAudio( w *bytes.Buffer, media katana.Media, cont []katana.DocNode, doc *katana.Doc ){
  fmt.Fprintf( w, "<div class=\"audio\">\n<audio controls>\n" )
  if media.Ext == "mp3" {
    fmt.Fprintf( w, "<source src=\"%s\" type=\"audio/mpeg\" >\n",  media.Src )
  } else {
    fmt.Fprintf( w, "<source src=\"%s\" type=\"audio/%s\" >\n",  media.Src, media.Ext )
  }
  fmt.Fprintf( w, "Your browser does not support <em>.%s</em> audio\n", media.Ext )
  fmt.Fprintf( w, "</audio>\n" )
  walkContent( w, cont, doc )
  fmt.Fprintf( w, "</div>\n" )
}

func makeWrap( w *bytes.Buffer, wrap katana.Wrap, cont []katana.DocNode, doc *katana.Doc ) {
  fmt.Fprintf( w, "<div class=\"%s\" >\n", wrap.Type )
  walkContent( w, cont, doc )
  fmt.Fprintf( w, "</div>\n" )
}

func makeQuote( w *bytes.Buffer, quote katana.Quote, doc *katana.Doc ){
  w.WriteString( "<blockquote>\n" )

  for _, mark := range quote.Quotex {
    if mark.Type == 'q' {
      mark.Type = 0
      w.WriteString( "<div class=\"quote-author\" >\n<p>" )
      w.WriteString( Fontify( mark  ) )
      w.WriteString( "</p>\n</div>\n" )
    } else {
      w.WriteString( "<p>" )
      w.WriteString( Fontify( mark ) )
      w.WriteString( "</p>\n" )
    }
  }

  w.WriteString( "</blockquote>\n" )
}

func makePret( w *bytes.Buffer, pret katana.Pret, doc *katana.Doc ) {
  fmt.Fprintf( w, "<div class=\"pret\" >\n%s</div>\n", txt.RmIndent( Fontify( pret.IndentMarkup ), pret.Indent ) )
}
