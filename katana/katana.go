package katana

import (
  "github.com/nasciiboy/regexp3"
)

const (
  EmptyNode = iota
  CommandNode
  HeadlineNode
  TableNode
  TableRowNode
  TableCellNode
  ListNode
  ListElementNode
  AboutNode
  TextNode
  CommentNode
)

type FullData struct {
  Comm    string
  Params  string
  Arg     string
  Data    string
  Mark    Markup
  Type    int
  N       int
  N2      int
}

type Node interface {
  Type() int
  Get () FullData
}

type DocNode struct {
  Node
  Cont []DocNode
}

type Options struct {
  Toc           bool
  Highlight     bool
  Pygments      bool
  HShift        int
}

type Doc struct {
  Toc         []DocNode
  Title         string
  Subtitle      string
  Author      []string
  Translator  []string
  Mail          string
  Licence       string
  Id            string
  Style       []string
  Date          string
  Tags          string
  Description   string
  Lang          string
  Options     []string
  OptionsData   Options
}

func (d *Doc) Parse( str string ){
  var i int
  *d, i = GetSetup( str )
  d.Toc = GetToc  ( str[i:] )
}

func (d *DocNode) AddNode( n Node ) {
  d.Cont = append( d.Cont, DocNode{ n, nil } )
}

func (d *DocNode) AddNodeToLast( n Node ) {
  last := len( d.Cont )

  if last > 0 {
    last--
    d.Cont[ last ].AddNode( n )
  } else {
    d.AddNode( n )
  }
}

func (d *DocNode) Add( n DocNode ) {
  d.Cont = append( d.Cont, n )
}

func (d *DocNode) AddToLast( n DocNode ) {
  last := len( d.Cont )

  if last > 0 {
    last--
    d.Cont[ last ].Add( n )
  } else {
    d.Add( n )
  }
}

func (d *DocNode) GetLast() *DocNode {
  last := len( d.Cont )

  if last == 0 {
    d.AddNode( nil )
    return &d.Cont[ 0 ]
  }

  last--
  return &d.Cont[ last ]
}

const (
  TextSimple = iota
  TextQuoteAuthor
)

type Text struct {
  TextType int
  Mark     Markup
}

type Headline struct {
  Mark  Markup
  Level int
}

const (
  ListMinusNode = iota
  ListPlusNode
  ListNumNode
  ListAlphaNode
  ListMdefNode
  ListPdefNode
  ListDialogNode
)

type List struct {
  ListType int
}

type ListElement struct {
  Mark   Markup
  Prefix string
}

type Command struct {
  Comm   string
  Params string
  Arg    string
  Mark   Markup
  Body   string
}

type About struct {
  Mark Markup
}

type Table struct {
  Rows int
  Cols int
}

const (
  TableHead = iota
  TableBody
  TableFoot
)

type TableRow struct {
  Kind int
}

type TableCell struct {
  Mark   Markup
  Wide   int
  Length int
}

func (h Headline    ) Type() int { return HeadlineNode     }
func (t Text        ) Type() int { return TextNode         }
func (c Command     ) Type() int { return CommandNode      }
func (a About       ) Type() int { return AboutNode        }
func (l List        ) Type() int { return ListNode         }
func (l ListElement ) Type() int { return ListElementNode  }
func (t Table       ) Type() int { return TableNode        }
func (t TableRow    ) Type() int { return TableRowNode     }
func (t TableCell   ) Type() int { return TableCellNode    }

func (h Headline    ) Get() FullData { return FullData{ Mark: h.Mark, N: h.Level } }
func (t Text        ) Get() FullData { return FullData{ Mark: t.Mark, N: t.TextType } }
func (l List        ) Get() FullData { return FullData{ N: l.ListType } }
func (l ListElement ) Get() FullData { return FullData{ Mark: l.Mark, Data: l.Prefix } }
func (c Command     ) Get() FullData { return FullData{ Comm: c.Comm, Params: c.Params, Arg: c.Arg, Mark: c.Mark, Data: c.Body } }
func (a About       ) Get() FullData { return FullData{ Mark: a.Mark } }
func (t Table       ) Get() FullData { return FullData{ N: t.Rows, N2: t.Cols } }
func (t TableRow    ) Get() FullData { return FullData{ N: t.Kind } }
func (t TableCell   ) Get() FullData { return FullData{ Mark: t.Mark, N: t.Wide, N2: t.Length } }

func whoIsThere( line string ) int {
  var re regexp3.RE
  if len(line) == 0                                               { return EmptyNode
  } else if re.Match( line, "#^$:s+"                        ) > 0 { return EmptyNode
  } else if re.Match( line, "#^:@(:s)"                      ) > 0 { return CommentNode
  } else if re.Match( line, "#^:*+:b"                       ) > 0 { return HeadlineNode
  } else if re.Match( line, "#^$:b*:|([^|]+:|)+:s*"         ) > 0 { return TableNode
  } else if re.Match( line, "#^:b*(-|:+|(:d+|:a+)[.)]):b+:S") > 0 { return ListNode
  } else if re.Match( line, "#^:b*:>:b+:S"                  ) > 0 { return ListNode
  } else if re.Match( line, "#^:b*::{2}:b+:S"               ) > 0 { return AboutNode
  } else if re.Match( line, "#^:b*:.:.:b*[:w:-:_]+[^:>]*:>" ) > 0 { return CommandNode
  } else                                                          { return TextNode
  }
}

func whatListIsThere( list string ) int {
  var re regexp3.RE

  if        re.Match( list, "#^:b*:>:b+:S"            ) > 0 { return ListDialogNode
  } else if re.Match( list, "#^:b*-:b+<:S>"           ) > 0 {
    if re.Match( list[re.GpsCatch( 1 ):], "#?:b::{2}" ) > 0 { return ListMdefNode  }
                                                              return ListMinusNode
  } else if re.Match( list, "#^:b*:+:b+<:S>"          ) > 0 {
    if re.Match( list[re.GpsCatch( 1 ):], "#?:b::{2}" ) > 0 { return ListPdefNode  }
                                                              return ListPlusNode
  } else if re.Match( list, "#^:b*:d+[.)]:b+:S"       ) > 0 { return ListNumNode
  } else if re.Match( list, "#^:b*:a+[.)]:b+:S"       ) > 0 { return ListAlphaNode
  }

  return EmptyNode
}
