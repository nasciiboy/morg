package katana

import (
  "errors"
)

const (
  MarkupNil byte = iota
  MarkupEsc
  MarkupErr
  MarkupHeadline
  MarkupTitle
  MarkupSubTitle
  MarkupList
  MarkupDialog
  MarkupComment
  MarkupAbout
  MarkupCode
  MarkupText
)

type Markup struct {
  Body      []Markup
  Custom    []Markup

	Type      byte
  Data      string
}

func (m Markup) HasSomething() bool {
  if len( m.Body ) != 0 || len(m.Body) != 0 || len(m.Custom) != 0 {
    return true
  }

  return false
}

func (m Markup) String() (str string) {
  if len( m.Body ) == 0 {
    return m.Data
  }

  for _, c := range m.Body {
    str += c.String()
  }

  return str
}

func (m Markup) GetCustom() (str string) {
  if len( m.Custom ) == 0 {
    return m.Data
  }

  for _, c := range m.Custom {
    str += c.String()
  }

  return str
}

func (m Markup) MakeCustom() (str string) {
  if m.Data != "" {
    return m.Data
  }

  if len( m.Custom ) > 0 {
    for _, c := range m.Custom {
      str += c.String()
    }
    return str
  }

  for _, c := range m.Body {
    str += c.String()
  }
  return str
}

func (m Markup) Rebuild() (str string) {
  if len( m.Custom ) == 0 && len( m.Body ) == 0 {
    return m.Data
  }

  if m.Type > 33 {
    str = "@" + string(m.Type) + "("
  }

  for _, c := range m.Custom {
    str += c.Rebuild()
  }

  if len( m.Custom ) > 0 {
    str += "<>"
  }

  for _, c := range m.Body {
    str += c.Rebuild()
  }

  if m.Type > 10 {
    str += ")"
  }

  return str
}

func (m *Markup) Parse( str string ) error {
  m.Type = MarkupNil
  m.Body = make([]Markup, 0, 4)

  last , i := 0, 0
  for i < len( str ) {
    if str[ i ] == '@' {
      if last < i {
        m.Body = append( m.Body, Markup{ Type: MarkupNil, Data: str[last:i] } )
      }

      node, forward, err := MarckupTrigger( str[i:] )

      if err != nil { return err }

      m.Body = append(m.Body, node)
      i     += forward
      last   = i
      continue
    }

    i++
  }

  if last < len( str ) {
    m.Body = append( m.Body, Markup{ Type: MarkupNil, Data: str[last:] } )
  }

  return nil
}

func MarckupTrigger( str string ) (Markup, int, error) {
  if len( str ) == 1 {
    return Markup{ Type: MarkupErr, Data: str }, 1, nil
  }

  switch str[1] {
  case '(', ')', '[', ']', '{', '}', '@':
    return Markup{ Type: MarkupEsc, Data: str[1:2] }, 2, nil
  }

  switch len( str ) {
  case 2: return Markup{ Type: MarkupErr, Data: str }, 2, errors.New( "markup.go: incorret len" )
  case 3: return Markup{ Type: MarkupErr, Data: str }, 3, errors.New( "markup.go: markup without data" )
  }

  label, operator := str[ 1 ], str[ 2 ];

  switch operator {
  case '(': operator = ')'
  case '[': operator = ']'
  case '{': operator = '}'
  case '<': operator = '>'
  default : return Markup{ Type: MarkupErr, Data: str[:3] }, 3, errors.New( "markup.go: incorret operator" )
  }

  node, width, err := MarkupParser( str[3:], label, operator )
  return node, width + 3, err
}

func MarkupParser( str string, label, operator byte ) (node Markup, i int, err error) {
  node.Type = label
  node.Body = make( []Markup, 0, 2 )
  last     := 0

  for i < len( str ) {
    if len(str[i:]) >= 2 && str[i:i+2] == "<>" {
      if last < i {
        node.Body = append( node.Body, Markup{ Type:  MarkupNil, Data: str[last:i] } )
      }

      if len( node.Custom ) > 0 {
        return node, i, errors.New( "markup.go: more than one custom data" )
      }

      node.Custom  = node.Body
      node.Body    = make( []Markup, 0, 2)
      i           += 2
      last         = i
    } else {
      if str[ i ] == operator {
        if last < i {
          node.Body = append( node.Body, Markup{ Type:  MarkupNil, Data: str[last:i] } )
        }
        return  node, i + 1, nil
      }

      if str[ i ] == '@' {
        if last < i {
          node.Body = append( node.Body, Markup{ Type:  MarkupNil, Data: str[last:i] } )
        }

        iNode, forward, err := MarckupTrigger( str[i:] )

        if err != nil {
          return node, i, err
        }

        node.Body = append( node.Body, iNode )
        i        += forward
        last      = i
        continue
      }

      i++
    }
  }

  if last < len(str) {
    node.Body = append(node.Body, Markup{ Type:  MarkupNil, Data: str[last:], Custom: nil, Body: nil } )
  }

  return node, len( str ), nil
}

func StrToMark( str string ) (m Markup) {
  m.Parse( str )
  return
}
