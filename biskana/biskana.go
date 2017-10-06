package biskana

import (
  "github.com/nasciiboy/morg/katana"
  "github.com/nasciiboy/morg/biskana/html"
)

const (
  HTML uint = iota
)

func Export( doc *katana.Doc, to uint ) string {
  switch to {
  case HTML: return html.MakeHtml( doc )
  }

  return ""
}

func ExportPartial( doc *katana.Doc, to uint ) string {
  switch to {
  case HTML: return html.MakeHtmlBody( doc )
  }

  return ""
}
