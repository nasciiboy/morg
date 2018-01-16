package katana

const (
  NodeEmpty = iota
  NodeHeadline
  NodeText
  NodeBlock
  NodeBrick
  NodeAbout
  NodeSeparator
  NodeComment
  NodeFigure
  NodeCode
  NodeMedia
  NodeWrap
  NodeColumns
  NodePret
  NodeQuote
  NodeSrci

  NodeTable
  NodeTableRow
  NodeTableCell

  NodeList
  NodeListElement
)

type Node interface {
  NodeType() int
}

type DocNode struct {
  Node
  Cont []DocNode
}

func (d *DocNode) AddNode( n Node ) {
  d.Cont = append( d.Cont, DocNode{ Node: n, Cont: nil } )
}

func (d *DocNode) Add( n DocNode ) {
  d.Cont = append( d.Cont, n )
}

type Headline struct {
  Mark  Markup
  Level int
}

const (
  NodeListEmpty = iota
  NodeListMinus
  NodeListPlus
  NodeListNum
  NodeListAlpha
  NodeListMDef
  NodeListPDef
  NodeListDialog
)

type List struct {
  ListType  int
}

type ListElement struct {
  Mark   Markup
  Prefix string
}

type About struct {
  Mark Markup
}

type Separator struct {
}

type Figure struct {
  Args  ArgMap
  Title Markup
}

type Code struct {
  Lang        string
  Body        string
  Style       string
  Indexed     bool
  IndexNum    int

  Args        ArgMap
  SBody       Scanner
  SBodyIndent int
}

type BinariString struct {
  On   bool
  Str  string
}

type Srci struct {
  Lang        string
  Prompt      string
  Style       string
  Body        []BinariString
  Indexed     bool
  IndexNum    int

  Args        ArgMap
}

type Brick struct {
  Type       string
  Head       string
  Body       string

  Args        ArgMap
}

type Media struct {
  Type   string
  Ext    string
  Src    string

  Args   ArgMap
}

type Wrap struct {
  Type   string
  Head   string

  Args   ArgMap
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
  Type int
}

type TableCell struct {
  RawData string
  ColSpan int
  RowSpan int
}

type Columns struct {
  Head string
  Args ArgMap
}

type Pret struct {
  IndentMarkup Markup
  Indent       int
  Args         ArgMap
}

type Quote struct {
  Quotex []Markup
}

func (h Headline    ) NodeType() int { return NodeHeadline     }
func (m Markup      ) NodeType() int { return NodeText         }
func (b Brick       ) NodeType() int { return NodeBrick        }
func (a About       ) NodeType() int { return NodeAbout        }
func (f Figure      ) NodeType() int { return NodeFigure       }
func (c Code        ) NodeType() int { return NodeCode         }
func (m Media       ) NodeType() int { return NodeMedia        }
func (w Wrap        ) NodeType() int { return NodeWrap         }
func (c Columns     ) NodeType() int { return NodeColumns      }
func (p Pret        ) NodeType() int { return NodePret         }
func (q Quote       ) NodeType() int { return NodeQuote        }
func (s Srci        ) NodeType() int { return NodeSrci         }
func (l List        ) NodeType() int { return NodeList         }
func (l ListElement ) NodeType() int { return NodeListElement  }
func (t Table       ) NodeType() int { return NodeTable        }
func (t TableRow    ) NodeType() int { return NodeTableRow     }
func (t TableCell   ) NodeType() int { return NodeTableCell    }
func (s Separator   ) NodeType() int { return NodeSeparator    }
