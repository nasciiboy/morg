package biskana

import (
  "github.com/nasciiboy/morg/biskana/html"
)

type bType uint8

const (
  HTML bType = iota
)

func Export( str string, to bType ) string {
  if str == "" { return "" }

  switch to {
  case HTML: return html.MakeHtml( str )
  }

  return ""
}

func ExportPartial( str string, to bType ) string {
  if str == "" { return "" }

  switch to {
  case HTML: return html.MakeHtmlBody( str )
  }

  return ""
}
