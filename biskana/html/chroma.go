package html

import (
  "bytes"

	"github.com/nasciiboy/morg/katana"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func FancyCode( w *bytes.Buffer, c katana.Code ) {
	builder := styles.Get( c.Style ).Builder()
	style, err := builder.Build()
  if err != nil {
    //....
  }

  var options []html.Option

  if c.Indexed {
    options = append( options, html.WithLineNumbers(true))
    options = append( options, html.BaseLineNumber( c.IndexNum ))
  } else {
    options = append( options, html.BaseLineNumber( 1 ))
  }

  options = append( options, html.TabWidth( 4 ))

  if c.Style == "" { options = append( options, html.WithClasses(true) ) }// inlinestyle
  // options = append( options, html.Standalone() )  // htmlOnly

  formatters.Register("html", html.New(options...))
  format(w, style, lex( c.Lang, c.Body ))
}

func FancySrci( w *bytes.Buffer, body, lang, sstyle string ) {
	builder := styles.Get( sstyle ).Builder()
	style, err := builder.Build()
  if err != nil {
    //....
  }

  var options []html.Option

  options = append( options, html.TabWidth( 4 ))
  if sstyle == "" { options = append( options, html.WithClasses(true) ) }// inlinestyle

  formatters.Register("html", html.New(options...))
  format(w, style, lex( lang, body ))
}

func WriteStyle( w *bytes.Buffer, sstyle string ){
  w.WriteString( "<style type=\"text/css\">\n")

	builder := styles.Get( sstyle ).Builder()
	style, _ := builder.Build()

  formatter := html.New(html.WithClasses(true))
  formatter.WriteCSS(w, style)
  w.WriteString( "</style>\n" )
}

func lex( lang, contents string) chroma.Iterator {
	lexer := lexers.Get( lang )

	if lexer == nil {
    lexer = lexers.Analyse(contents)
	}
	if rel, ok := lexer.(*chroma.RegexLexer); ok {
		rel.Trace( false )
	}
	lexer = chroma.Coalesce(lexer)
	it, err := lexer.Tokenise(nil, string(contents))
  if err != nil {
    // ...
  }
	return it
}

func format(w *bytes.Buffer, style *chroma.Style, it chroma.Iterator) {
	formatter := formatters.Get( "html" )
	err := formatter.Format(w, style, it)
  if err != nil { panic( err ) }
}
