package biskana

import (
  "github.com/nasciiboy/regexp3"
)

type Options struct {
  Toc           bool
  Highlight     bool
  Pygments      bool
  HShift        int
}

type DocInfo struct {
  Title         string
  Subtitle      string
  Author        string
  Translator    string
  Mail          string
  Licence       string
  Id            string
  Style         string
  Date          string
  Tags          string
  Description   string
  Lang          string
  Options       string
  OptionsData   Options
}

func MakeHtml( str, title string ) string {
  head, docInfo := MakeHtmlHead( str )

  if docInfo.Title == "" {
    if title != "" {
      head = "<title>" +   title    + "</title>\n" + head
      docInfo.Title = title
    } else {
      head = "<title>" + "untitled" + "</title>\n" + head
    }
  }

  if docInfo.Lang == "" { docInfo.Lang = "en" }
  html :=
    "<?xml version=\"1.0\" encoding=\"utf-8\"?>\n" +
    "<!DOCTYPE html>\n" +
    "<html xmlns=\"http://www.w3.org/1999/xhtml\" lang=\"" + docInfo.Lang + "\" xml:lang=\"" + docInfo.Lang + "\">\n"

  if docInfo.OptionsData.Highlight {
    html += "<head>\n" + head +
      "<meta  http-equiv=\"Content-Type\" content=\"text/html;charset=utf-8\" />\n" +
      "<link rel=\"stylesheet\" href=\"highlight/styles/atelier-forest-dark.css\" />\n" +
      "<script src=\"highlight/highlight.pack.js\" ></script>\n" +
      "<script>hljs.initHighlightingOnLoad();</script>\n" +
      "</head>\n"
  } else {
    html += "<head>\n" + head +
      "<meta  http-equiv=\"Content-Type\" content=\"text/html;charset=utf-8\" />\n" +
      "</head>\n"
  }

  docInfo.OptionsData.HShift = 1
  body, toc := MakeHtmlBodyWithOptions( str, docInfo.OptionsData )

  html += "<body>\n"

  if docInfo.OptionsData.Toc {
    html += "<div id=\"toc\">\n" +
      "<p>index</p>\n" +
      "<div id=\"toc-contents\">\n" +
      toc +
      "</div>\n" +
      "</div>\n"
  }

  if docInfo.Title != "" {
    html += "<h1>" + ToHtml( docInfo.Title ) + "</h1>\n" +
      body + "</body>\n"
  } else {
    html += body + "</body>\n"
  }

  html += "</html>\n"

  return html
}

const ( COMMAND = iota; HEADLINE; TABLE; LIST; DEFINITION; ABOUT; TEXT; COMMENT; EMPTY )

func whoIsThere( line string ) uint {
  var re regexp3.RE
  if len(line) == 0                                               { return EMPTY
  } else if re.Match( line, "#^$:s+"                        ) > 0 { return EMPTY
  } else if re.Match( line, "#^:@(:s)"                      ) > 0 { return COMMENT
  } else if re.Match( line, "#^:*+:b"                       ) > 0 { return HEADLINE
  } else if re.Match( line, "#^$:b*:|([^|]+:|)+:s*"         ) > 0 { return TABLE
  } else if re.Match( line, "#^:b*(-|:+|(:d+|:a+)[.)]):b+:S") > 0 { return LIST
  } else if re.Match( line, "#^:b*:>:b+:S"                  ) > 0 { return LIST
  } else if re.Match( line, "#^:b*::{2}:b+:S"               ) > 0 { return ABOUT
  } else if re.Match( line, "#^:b*:.:.:b*[:w:-:_]+[^:>]*:>" ) > 0 { return COMMAND
  } else                                                          { return TEXT
  }
}

const ( LIST_ERR = iota; LIST_MINUS; LIST_PLUS; LIST_NUM; LIST_ALPHA; LIST_MDEF; LIST_PDEF; LIST_DIALOG )

func whatListIsThere( list string ) uint {
  var re regexp3.RE

  if        re.Match( list, "#^:b*:>:b+:S"            ) > 0 { return LIST_DIALOG
  } else if re.Match( list, "#^:b*-:b+<:S>"           ) > 0 {
    if re.Match( list[re.GpsCatch( 1 ):], "#?:b::{2}" ) > 0 { return LIST_MDEF  }
                                                              return LIST_MINUS
  } else if re.Match( list, "#^:b*:+:b+<:S>"          ) > 0 {
    if re.Match( list[re.GpsCatch( 1 ):], "#?:b::{2}" ) > 0 { return LIST_PDEF  }
                                                              return LIST_PLUS
  } else if re.Match( list, "#^:b*:d+[.)]:b+:S"       ) > 0 { return LIST_NUM
  } else if re.Match( list, "#^:b*:a+[.)]:b+:S"       ) > 0 { return LIST_ALPHA
  }

  return LIST_ERR
}