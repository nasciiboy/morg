package katana

import (
  "github.com/nasciiboy/txt"
  "fmt"
  "sort"
  "bytes"
  "errors"
)

const OUTSIDE = '\uFFFD' // unicode.RuneError = '\uFFFD'
const CORNER  = '+'      //  +----+     //  +--------+-----+--------+
const VEDGE   = '|'      //  |    |     //  | header       |  Yhee! |
const HEDGE   = '-'      //  +----+     //  +========+=====+========+
                                        //  | body-A | B   |   C    |
                                        //  +~~~~~~~~+~~~~~+~~~~~~~~+
const HEDGEH  = '='  // h == head       //  | Yhak!  |    footer    |
const HEDGEF  = '~'  // f == foot       //  +--------+-----+--------+

func txt2rTable( txtTable string ) (t rTable, e error) {
  runeTable := gridify( txtTable )
  if runeTable == nil { return t, errors.New( "empty input" ) }

  var yHead, yFoot int = 0, -1
  for y, row := range runeTable {
    if isHead( row ) {
      if yHead ==  0 {
        yHead = y
      } else { return t, fmt.Errorf( "there can only be one header, repetition in row %d", y ) }
    }
    if isFoot( row ) {
      if yFoot == -1 {
        yFoot = y
      } else { return t, fmt.Errorf( "there can only be one footer, repetition in row %d", y ) }
    }
  }

  if yHead !=  0 { normalizeRuneTableRow( runeTable, yHead ) }
  if yFoot != -1 { normalizeRuneTableRow( runeTable, yFoot ) }

  g         := new( grid ).init( runeTable )
  rows      := make( [][]*quad, 0, 10 )
  row       := make(   []*quad, 0,  5 )
  hotPoints := make( map[point]bool )
  xEdges, yEdges := map[int]bool{ 0: true }, map[int]bool{ 0: true }

  for next, y := new(point), 0; next != nil;  {
    s := g.newQuad(*next)

    if e = s.tourist(); e != nil {
      return t, fmt.Errorf( "square.tourist: %v", e )
    }

    xEdges[ s.cornerC.x ] = true
    yEdges[ s.cornerA.y ] = true

    if y != s.cornerD.y {
      rows = append( rows, row)
      row  = make( []*quad, 0, 5 )
      y    = s.cornerD.y
    }

    row = append( row, s )

    if g.visit( s.cornerA.nmv( DOWN ) ) == '|' {
      hotPoints[ s.cornerA ] = true
    }

    if _, exist := hotPoints[ s.cornerC ]; exist {
      delete( hotPoints, s.cornerC )
      next = &s.cornerC
    } else if g.visit( s.cornerC.nmv( RIGHT ) ) == '-' {
      next = &s.cornerC
    } else {
      next = pullNextPoint( hotPoints )

      if next == nil { rows = append( rows, row) }
    }
  }

  xS, yS :=  mapToSortIntArray(xEdges), mapToSortIntArray(yEdges)
  table := make( [][]TableCell, len( rows ) )
  for y, row := range rows {
    table[y] = make([]TableCell, len( row ) )
    for x, q := range row {
      table[y][x].RowSpan = rangeCells( q.cornerD.y, q.cornerA.y, yS )
      table[y][x].ColSpan = rangeCells( q.cornerD.x, q.cornerB.x, xS )
      table[y][x].RawData = cutCellTxt( g.data, q.cornerD, q.cornerB )
    }

    if yHead != 0 && (row[0].cornerA.y == yHead) {
      yHead  = y + 1
      t.Head = table[:yHead]
    }
    if yFoot != -1 && (row[0].cornerD.y == yFoot) {
      yFoot  = y
      t.Foot = table[yFoot:]
    }
  }
  if yFoot == -1 { yFoot = len( table ) }
  if yFoot - yHead > 0 { t.Body = table[yHead:yFoot] }

  return
}

type rTable struct {
  Head [][]TableCell
  Body [][]TableCell
  Foot [][]TableCell
}

type point struct { y, x int }

type grid struct {
  data [][]rune
  w, h int
}

func (g *grid) init( data [][]rune ) *grid {
  g.data = data
  g.h = len( g.data )
  if g.h > 0 {
    g.w = len( g.data[0] )
  } else { g.w = 0 }

  return g
}

func (g *grid) visit( p point ) rune {
  if p.x < 0 || p.y < 0 || p.x >= g.w || p.y >= g.h ||
     g.h < 1 || g.w < 1 {
    return OUTSIDE
  }

  return g.data[p.y][p.x]
}

const ( RIGHT = iota; UP; LEFT; DOWN )

func (p *point) mv( dir byte ) point {
  switch dir {
  case RIGHT: p.x += 1
  case UP   : p.y -= 1
  case LEFT : p.x -= 1
  case DOWN : p.y += 1
  }

  return *p
}

func (p point) nmv( dir byte ) point { return p.mv( dir ) }

type quad struct {                          // D        C
  cornerA, cornerB, cornerC, cornerD point  //  +------+
  *grid                                     //  |      |
}                                           //  +------+
                                            // A        B

func (g *grid) newQuad( p point ) *quad {
  return &quad{ grid: g, cornerD: p }
}

func (s *quad) tourist() error {
  if s.visit( s.cornerD ) != '+' {
    return fmt.Errorf( "quad.cornerD %v != '+'", s.cornerD )
  }
  var c rune
  if s.cornerC, c = s.findUpCorner( s.cornerD ); c != '+' {
    return fmt.Errorf( "quad.cornerC %v != '+'", s.cornerC )
  }
  if s.cornerB, c = s.findRightCorner( s.cornerC ); c != '+' {
    return fmt.Errorf( "quad.cornerB %v != '+'", s.cornerB )
  }
  if s.cornerA, c = s.findDownCorner ( s.cornerB ); c != '+' {
    return fmt.Errorf( "quad.cornerA %v != '+'", s.cornerA )
  }

  if s.cornerA.x != s.cornerD.x {
    return fmt.Errorf( "quad (top-left) %v != %v (bottom-left)", s.cornerD, s.cornerA )
  }

  return nil
}

func (g *grid) findCorner( p point, dir byte ) (point, rune) {
  edge := HEDGE
  if dir == UP || dir == DOWN { edge = VEDGE }

  for {
    switch r := g.visit( p.mv( dir ) ); r {
    case CORNER : return p, r
    case edge   : continue
    default     : return p, r // error
    }
  }
}

func (q *quad) findUpCorner( c point) (p point, r rune) {
  for p, r = q.findCorner( c, RIGHT ); q.visit( p.nmv( DOWN ) ) != VEDGE; p, r = q.findCorner( p, RIGHT ) {
    if r != CORNER { return }
  }

  return
}

func (q *quad) findDownCorner( c point) (p point, r rune) {
  for p, r = q.findCorner( c, LEFT ); q.visit( p.nmv( UP ) ) != VEDGE; p, r = q.findCorner( p, LEFT ) {
    if r != CORNER { return }
  }

  return
}

func (q *quad) findRightCorner( c point) (p point, r rune) {
  for p, r = q.findCorner( c, DOWN ); q.visit( p.nmv( LEFT ) ) != HEDGE; p, r = q.findCorner( p, DOWN ) {
    if r != CORNER { return }
  }

  return
}

var zs = [3]string{ "thead", "tbody", "tfoot" }
var ys = [3]string{    "th",    "td",    "td" }

func TextTable2HtmlTable( data string ) (string, error) {
  t, e := txt2rTable( data )
  if e != nil { return "", fmt.Errorf( "Text2Table: %v", e ) }

  buf := new( bytes.Buffer )
  fmt.Fprintf( buf,  "<table border=\"1\">\n" )

  for i, sec := range [3][][]TableCell{ t.Head, t.Body, t.Foot } {
    if sec == nil { continue }

    fmt.Fprintf( buf,  "<%s>\n", zs[i] )
      for _, row := range sec {
        buf.WriteString( "  <tr>" )
        for _, cell := range row {
          fmt.Fprintf( buf, "<%s", ys[ i ] )
          if cell.ColSpan != 1 { fmt.Fprintf( buf, " colspan=\"%d\"", cell.ColSpan ) }
          if cell.RowSpan != 1 { fmt.Fprintf( buf, " rowspan=\"%d\"", cell.RowSpan ) }
          fmt.Fprintf( buf, "><p>%s</p></%s>", txt.SpaceSwap( cell.RawData, " " ), ys[ i ] )
        }
        buf.WriteString( "</tr>\n" )
      }
    fmt.Fprintf( buf,  "</%s>\n", zs[i] )
  }

  fmt.Fprintf( buf, "</table>\n" )

  return buf.String(), nil
}

func gridify( data string ) (grid [][]rune) {
  if data == "" { return }

  lines := txt.GetLines( data )
  grid   = make( [][]rune, len( lines ) )

  for i, l := range lines { grid[ i ] = []rune( l ) }

  return
}

func pullNextPoint( m map[point]bool ) *point {
  if len( m ) == 0 { return nil }

  n := point{}
  for key, _ := range m { // random first value
    n = key
    break
  }

  for key, _ := range m {
    if key.y < n.y {
      n = key
    } else if key.y == n.y && key.x < n.x  {
      n = key
    }
  }

  delete( m, n )
  return &n
}

func rangeCells( a, b int, cells []int ) int {
  min := 0
  for min < len( cells ) && a > cells[ min ] { min++ }
  max := min
  for max < len( cells ) && b > cells[ max ] { max++ }

  if min == max {
    return 1
  } else if a < cells[min] {
    return max - min + 1
  }

  return max - min
}

func cutCellTxt( grid [][]rune, d, b point ) string {
  buf := new( bytes.Buffer )

  for y := d.y + 1; y < b.y; y++ {
    buf.WriteString( string(grid[y][d.x+1:b.x]) )
    buf.WriteByte( '\n' )
  }

  return buf.String()
}

func mapToSortIntArray( m map[int]bool ) []int {
  arry, i := make([]int, len( m )), 0
  for key := range m { arry[i] = key; i++ }
  sort.Ints( arry )
  return arry
}

func isHead( h []rune ) bool {
  for _, c := range h {
    if c != CORNER && c != HEDGEH { return false }
  }
  return true
}

func isFoot( h []rune ) bool {
  for _, c := range h {
    if c != CORNER && c != HEDGEF { return false }
  }
  return true
}

func normalizeRuneTableRow( t [][]rune, y int ){
  for x := range t[y] {
    if t[y][x] != CORNER { t[y][x] = HEDGE }
  }
}
