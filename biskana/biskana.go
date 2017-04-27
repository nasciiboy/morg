package biskana

import (
  "github.com/nasciiboy/morg/biskana/html"
  "github.com/nasciiboy/morg/biskana/txt"
)

const (
  HTML uint = iota
  TXT
)

func Export( str string, to uint ) string {
  if str == "" { return "" }

  switch to {
  case HTML: return html.MakeHtml( str )
  case TXT : return  txt.MakeTxt ( str )
  }

  return ""
}

func ExportPartial( str string, to uint ) string {
  if str == "" { return "" }

  switch to {
  case HTML: return html.MakeHtmlBody( str )
  case TXT : return  txt.MakeTxtBody ( str )
  }

  return ""
}
