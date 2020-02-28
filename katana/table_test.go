package katana

import (
  "testing"
  "fmt"
  "bytes"
  "github.com/nasciiboy/txt"
)

func TestGridify( t *testing.T ){
  data := []struct {
    in string
    out [][]rune
  } {

    { `| 0 | 1 | 2 |
|---|---|---|
| 3 | 4 | 5 |
|---|---|---|
| 6 | 7 | 8 |
|---|---|---|`,

      [][]rune {
        []rune( "| 0 | 1 | 2 |" ),
        []rune( "|---|---|---|" ),
        []rune( "| 3 | 4 | 5 |" ),
        []rune( "|---|---|---|" ),
        []rune( "| 6 | 7 | 8 |" ),
        []rune( "|---|---|---|" ),
      },
    },

    { `| 0 | 1 | 2 |
|---|---|----|
| 3 |4 | 5 |
|---|---|--|
| 6 | 7 | 8 |
|---|----|---|`,

      [][]rune {
        []rune( "| 0 | 1 | 2 |" ),
        []rune( "|---|---|----|" ),
        []rune( "| 3 |4 | 5 |" ),
        []rune( "|---|---|--|" ),
        []rune( "| 6 | 7 | 8 |" ),
        []rune( "|---|----|---|" ),
      },
    },

  }

  for i, d := range data {
    r := gridify( d.in )
    if diff, err := diffGrids( r, d.out ); diff {
      t.Errorf( "TestGridify() #%d\n %q\n", i, err )
    }
  }
}

func diffGrids( a, b [][]rune ) (diff bool, line string) {
  if len( a ) != len( b ) { return true, fmt.Sprintf( "lens a <- %d != %d -> b", len( a ), len( b ) ) }
  for i, l := range a {
    if len( l ) != len( b[i] ) { return true, fmt.Sprintf( "line[%d] lens %d != %d", i, len( l ), len( b[i] ) ) }

    for n, r := range l {
      if r != b[i][n] { return true, fmt.Sprintf( "line[%d][%d] %c != %c", i, n, r, b[i][n] ) }
    }
  }

  return false, ""
}

func TestRangeCells( t *testing.T ){
  data := [...]struct{
    a, b, out int
    cells []int
  } {
    {  0,  9, 1, []int{ 0, 9, 18, 27, 36, 45, 54, 63, 72, 81, 90 } },
    {  0,  5, 1, []int{ 0, 9, 18, 27, 36, 45, 54, 63, 72, 81, 90 } },
    {  2,  9, 1, []int{ 0, 9, 18, 27, 36, 45, 54, 63, 72, 81, 90 } },
    {  2,  8, 1, []int{ 0, 9, 18, 27, 36, 45, 54, 63, 72, 81, 90 } },
    {  0, 18, 2, []int{ 0, 9, 18, 27, 36, 45, 54, 63, 72, 81, 90 } },
    {  9, 18, 1, []int{ 0, 9, 18, 27, 36, 45, 54, 63, 72, 81, 90 } },
    {  7, 18, 2, []int{ 0, 9, 18, 27, 36, 45, 54, 63, 72, 81, 90 } },
    {  5, 45, 5, []int{ 0, 9, 18, 27, 36, 45, 54, 63, 72, 81, 90 } },
    {  5, 46, 6, []int{ 0, 9, 18, 27, 36, 45, 54, 63, 72, 81, 90 } },
    { 15, 90, 9, []int{ 0, 9, 18, 27, 36, 45, 54, 63, 72, 81, 90 } },
  }

  for _, d := range data {
    if r := rangeCells( d.a, d.b, d.cells ); r != d.out {
      t.Errorf( "Test rangeCells( %d, %d, %v )\nresult %d != %d expected", d.a, d.b, d.cells, r, d.out )
    }
  }
}

const r0 = "012\n345\n678"

func TestVisit( t *testing.T ){
  data := []struct {
    txt  string
    p    point
    r    rune
  } {
    { "", point{ y:  0, x:  0}, OUTSIDE },
    { "", point{ y:  1, x:  0}, OUTSIDE },
    { "", point{ y:  0, x: -1}, OUTSIDE },
    { r0, point{ y:  0, x:  0}, '0' },
    { r0, point{ y:  1, x:  1}, '4' },
    { r0, point{ y:  2, x:  2}, '8' },
    { r0, point{ y: -1, x:  0}, OUTSIDE },
    { r0, point{ y:  0, x: -1}, OUTSIDE },
    { r0, point{ y:  3, x:  0}, OUTSIDE },
    { r0, point{ y:  0, x:  3}, OUTSIDE },
    { r0, point{ y:  5, x:  5}, OUTSIDE },
    { r0, point{ y: -5, x: -5}, OUTSIDE },
    { "1", point{ y:  0, x:  0}, '1' },
  }

  for i, d := range data {
    g := gridify( d.txt )
    r := new( grid ).init( g ).visit( d.p )

    if r != d.r {
      t.Errorf( "TestVisid() # %d, %c != %c\n", i, d.r, r )
    }
  }
}

const r1 = "+-+\n| |\n+-+"
const r1a = "+-o\n| |\n+-+"
const r1b = "+-+\n| |\no-+"
const r2 = "+---+\n|   |\n|   |\n|   |\n+---+"

func TestFindCorner( t *testing.T ){
  data := []struct {
    txt  string
    init point
    dir  byte
    end  point
    r    rune
  } {
    { r1,  point{ y: 0, x: 0}, RIGHT, point{ y: 0, x: 2}, '+' },
    { r1,  point{ y: 0, x: 0}, DOWN , point{ y: 2, x: 0}, '+' },
    { r1,  point{ y: 2, x: 2}, LEFT , point{ y: 2, x: 0}, '+' },
    { r1,  point{ y: 2, x: 2}, UP   , point{ y: 0, x: 2}, '+' },
    { r1a, point{ y: 0, x: 0}, RIGHT, point{ y: 0, x: 2}, 'o' },
    { r1b, point{ y: 0, x: 0}, DOWN , point{ y: 2, x: 0}, 'o' },
    { r1b, point{ y: 2, x: 2}, LEFT , point{ y: 2, x: 0}, 'o' },
    { r1a, point{ y: 2, x: 2}, UP   , point{ y: 0, x: 2}, 'o' },
    { r2,  point{ y: 0, x: 0}, RIGHT, point{ y: 0, x: 4}, '+' },
    { r2,  point{ y: 0, x: 0}, DOWN , point{ y: 4, x: 0}, '+' },
    { r2,  point{ y: 4, x: 4}, LEFT , point{ y: 4, x: 0}, '+' },
    { r2,  point{ y: 4, x: 4}, UP   , point{ y: 0, x: 4}, '+' },
  }

  for i, d := range data {
    in    := gridify( d.txt )
    grid := new( grid ).init( in )
    end, r := grid.findCorner( d.init, d.dir )

    if r != d.r || end.y != d.end.y || end.x != d.end.x {
      t.Errorf( "TestGridify() # %d\nin  [%03d,%03d,%c]\nout [%03d,%03d,%c]\n",
        i, d.end.y, d.end.x, d.r, end.y, end.x, r )
    }
  }
}

func TestSquareTourist( t *testing.T ){
  data := []struct {
    txt  string
    init point
    out  [4]point
  } {
    { "+-+\n| |\n+-+",
      point{ y: 0, x: 0},
      [4]point{ {0,0}, {0,2}, {2,2}, {2,0}, } },
    { "+---+\n|   |\n|   |\n|   |\n+---+",
      point{ y: 0, x: 0},
      [4]point{ {0,0}, {0,4}, {4,4}, {4,0}, } },
    { "+---+--+\n|   |  |\n|   |  |\n|   |  |\n+---+--+",
      point{ y: 0, x: 4},
      [4]point{ {0,4}, {0,7}, {4,7}, {4,4}, } },
  }

  for i, d := range data {
    in   := gridify( d.txt )
    quad := new( grid ).init( in ).newQuad(d.init)
    quad.tourist()

    if quad.diffPoints( d.out ) {
      t.Errorf( "TestSquareTourist() # %d\nin  [%03d,%03d][%03d,%03d][%03d,%03d][%03d,%03d]\nout [%03d,%03d][%03d,%03d][%03d,%03d][%03d,%03d]\n",
        i,
        quad.cornerD.y, quad.cornerD.x, quad.cornerC.y, quad.cornerC.x, quad.cornerB.y, quad.cornerB.x, quad.cornerA.y, quad.cornerA.x,
        d.out[0].y, d.out[0].x, d.out[1].y, d.out[1].x, d.out[2].y, d.out[2].x, d.out[3].y, d.out[3].x )
    }
  }
}

func TestSquareTouristX( t *testing.T ){
  data := []struct {
    txt  string
    init point
    out  [4]point
    err  string
  } {
// +-+-+--+
// |      |
// +      +
// |      |
// +--+---+
    { "+-+-+--+\n|      |\n+      +\n|      |\n+--+---+",
      point{ y: 0, x: 0},
      [4]point{ {0,0}, {0,7}, {4,7}, {4,0}, },
      "" },
//     v
// +-+-X--+
// |      |
// +      +
// |      |
// X--+---+
    { "+-+-+--+\n|      |\n+      +\n|      |\n+--+---+",
      point{ y: 0, x: 4},
      [4]point{ {0,4}, {0,7}, {4,7}, {4,0}, },
      "quad (top-left) {0 4} != {4 0} (bottom-left)" },
//        v
// +-+-+--X
// |      |
// +      +
// |      |
// +--+---+
    { "+-+-+--+\n|      |\n+      +\n|      |\n+--+---+",
      point{ y: 0, x: 7},
      [4]point{ {0,7}, {0,8}, {0,0}, {0,0}, },
      "quad.cornerC {0 8} != '+'" },
// v
// X-+-+--+
// |      |
// +      +
// |      |
//  --+---+
    { "+-+-+--+\n|      |\n+      +\n|      |\n --+---+",
      point{ y: 0, x: 0},
      [4]point{ {0,0}, {0,7}, {4,7}, {4,0}, },
      "quad.cornerA {4 0} != '+'" },
// v
// X-+-+--+
// |      |
// +      +
// |
// +--+---+
    { "+-+-+--+\n|      |\n+      +\n|       \n+--+---+",
      point{ y: 0, x: 0},
      [4]point{ {0,0}, {0,7}, {3,7}, {0,0} },
      "quad.cornerB {3 7} != '+'" },
// +---+--+
// |   |  |
// |   |  |
// |   |  |
// +---+--+
    { "+---+--+\n|   |  |\n|   |  |\n|   |  |\n+---+--+",
      point{ y: 0, x: 4},
      [4]point{ {0,4}, {0,7}, {4,7}, {4,4} },
      "" },
// +-+-+--+---+---+
// |      |   |   |
// +   +--+---+---+
// |   |      |   |
// +--++      +---+
// |   |      |   |
// +   +--+---+   +
// |      |       |
// +--+---+--+----+
    { "+-+-+--+---+---+\n|      |   |   |\n+   +--+---+---+\n|   |      |   |\n+--++      +---+\n|   |      |   |\n+   +--+---+   +\n|      |       |\n+--+---+--+----+",
      point{ y: 2, x: 4},
      [4]point{ {2,4}, {2,11}, {6,11}, {6,4}, },
      "" },
  }

  for i, d := range data {
    quad := new( grid ).init( gridify( d.txt ) ).newQuad(d.init)
    err  := quad.tourist()

    if quad.diffPoints( d.out ) || (err != nil && d.err != err.Error()) {
      t.Errorf( "\nTestQuadTouristX() # %d, err: %v\nin  [%03d,%03d][%03d,%03d][%03d,%03d][%03d,%03d]\nout [%03d,%03d][%03d,%03d][%03d,%03d][%03d,%03d]\n",
        i, err,
        d.out[0].y, d.out[0].x, d.out[1].y, d.out[1].x, d.out[2].y, d.out[2].x, d.out[3].y, d.out[3].x,
        quad.cornerD.y, quad.cornerD.x, quad.cornerC.y, quad.cornerC.x, quad.cornerB.y, quad.cornerB.x, quad.cornerA.y, quad.cornerA.x )
    }
  }
}

func TestTextTable2HtmlTable( t *testing.T ){
  data := []struct {
    in, out, err string
  } {

    { `+-+-+--+
|      |
+      +
|      |
+--+---+`,
      `<table border="1">
<tbody>
  <tr><td><p> </p></td></tr>
</tbody>
</table>
`, "" },

    { `+---+---+---+
| 0 | 1 | 2 |
+---+---+---+
| 3 | 4 | 5 |
+---+---+---+
| 6 | 7 | 8 |
+---+---+---+`,
    `<table border="1">
<tbody>
  <tr><td><p> 0 </p></td><td><p> 1 </p></td><td><p> 2 </p></td></tr>
  <tr><td><p> 3 </p></td><td><p> 4 </p></td><td><p> 5 </p></td></tr>
  <tr><td><p> 6 </p></td><td><p> 7 </p></td><td><p> 8 </p></td></tr>
</tbody>
</table>
`, "" },

    { `+---+---+---+
| 0 | 1 | 2 |
+---+---+---+
|   3   | 4 |
+-------+---+
|   5   | 6 |
+-------+---+`,
    `<table border="1">
<tbody>
  <tr><td><p> 0 </p></td><td><p> 1 </p></td><td><p> 2 </p></td></tr>
  <tr><td colspan="2"><p> 3 </p></td><td><p> 4 </p></td></tr>
  <tr><td colspan="2"><p> 5 </p></td><td><p> 6 </p></td></tr>
</tbody>
</table>
`, "" },

    { `+---+---+---+
| 0 | 1 | 2 |
+---+---+---+
|   3   |   |
+-------+ 4 |
|   5   |   |
+-------+---+`,
      `<table border="1">
<tbody>
  <tr><td><p> 0 </p></td><td><p> 1 </p></td><td><p> 2 </p></td></tr>
  <tr><td colspan="2"><p> 3 </p></td><td rowspan="2"><p> 4 </p></td></tr>
  <tr><td colspan="2"><p> 5 </p></td></tr>
</tbody>
</table>
`, "" },

    { `+---+---+---+
| 0 | 1 | 2 |
+---+---+---+
|       |   |
|   3   | 4 |
|       |   |
+-------+---+`,
    `<table border="1">
<tbody>
  <tr><td><p> 0 </p></td><td><p> 1 </p></td><td><p> 2 </p></td></tr>
  <tr><td colspan="2"><p> 3 </p></td><td><p> 4 </p></td></tr>
</tbody>
</table>
`, "" },

    { `+----------------------+
|   mi tabla compleja  |
+-------+----+---------+
|  a    | b  |    c    |
+---+---+----+----+----+
| d | e | f  |  g |  h |
|   |   +----+----+----+
|   |   | i  |    j    |
|   +---+----+---------+
|   | k | l  |         |
+---+   +----+    m    |
|   |   | o  |         |
| n |   +----+---------+
|   |   | p  |    q    |
+---+---+----+---------+`,
    `<table border="1">
<tbody>
  <tr><td colspan="5"><p> mi tabla compleja </p></td></tr>
  <tr><td colspan="2"><p> a </p></td><td><p> b </p></td><td colspan="2"><p> c </p></td></tr>
  <tr><td rowspan="3"><p> d </p></td><td rowspan="2"><p> e </p></td><td><p> f </p></td><td><p> g </p></td><td><p> h </p></td></tr>
  <tr><td><p> i </p></td><td colspan="2"><p> j </p></td></tr>
  <tr><td rowspan="3"><p> k </p></td><td><p> l </p></td><td colspan="2" rowspan="2"><p> m </p></td></tr>
  <tr><td rowspan="2"><p> n </p></td><td><p> o </p></td></tr>
  <tr><td><p> p </p></td><td colspan="2"><p> q </p></td></tr>
</tbody>
</table>
`, "" },

    { `+------------+---------+
|     a      |    b    |
+-------+----+----+----+
|   c   | d  |  e |    |
+---+---+----+----+ f  |
| g |      h      |    |
+---+-+---------+-+----+
|     |    j    |      |
|  i  +-----+---+   k  |
|     |  l  | m |      |
+-----+-----+---+------+`,
    `<table border="1">
<tbody>
  <tr><td colspan="5"><p> a </p></td><td colspan="3"><p> b </p></td></tr>
  <tr><td colspan="3"><p> c </p></td><td colspan="2"><p> d </p></td><td colspan="2"><p> e </p></td><td rowspan="2"><p> f </p></td></tr>
  <tr><td><p> g </p></td><td colspan="6"><p> h </p></td></tr>
  <tr><td colspan="2" rowspan="2"><p> i </p></td><td colspan="4"><p> j </p></td><td colspan="2" rowspan="2"><p> k </p></td></tr>
  <tr><td colspan="2"><p> l </p></td><td colspan="2"><p> m </p></td></tr>
</tbody>
</table>
`, "" },
    { `+--------+-----+--------+
| header       |  Yhee! |
+========+=====+========+
| body-A | B   |   C    |
+~~~~~~~~+~~~~~+~~~~~~~~+
| Yhak!  |    footer    |
+--------+-----+--------+
`, `<table border="1">
<thead>
  <tr><th colspan="2"><p> header </p></th><th><p> Yhee! </p></th></tr>
</thead>
<tbody>
  <tr><td><p> body-A </p></td><td><p> B </p></td><td><p> C </p></td></tr>
</tbody>
<tfoot>
  <tr><td><p> Yhak! </p></td><td colspan="2"><p> footer </p></td></tr>
</tfoot>
</table>
`, "" },
    { `+--------+-----+--------+
| header       |  Yhee! |
+========+=====+========+
`, `<table border="1">
<thead>
  <tr><th><p> header </p></th><th><p> Yhee! </p></th></tr>
</thead>
</table>
`, "" },
    { `+~~~~~~~~+~~~~~+~~~~~~~~+
| Yhak!  |    footer    |
+--------+-----+--------+
`, `<table border="1">
<tfoot>
  <tr><td><p> Yhak! </p></td><td><p> footer </p></td></tr>
</tfoot>
</table>
`, "" },
    { `+--------+-----+--------+
| header       |  Yhee! |
+========+=====+========+
| body-A | B   |   C    |
+--------+-----+--------+
`, `<table border="1">
<thead>
  <tr><th colspan="2"><p> header </p></th><th><p> Yhee! </p></th></tr>
</thead>
<tbody>
  <tr><td><p> body-A </p></td><td><p> B </p></td><td><p> C </p></td></tr>
</tbody>
</table>
`, "" },
    { `+--------+-----+--------+
| body-A | B   |   C    |
+~~~~~~~~+~~~~~+~~~~~~~~+
| Yhak!  |    footer    |
+--------+-----+--------+
`, `<table border="1">
<tbody>
  <tr><td><p> body-A </p></td><td><p> B </p></td><td><p> C </p></td></tr>
</tbody>
<tfoot>
  <tr><td><p> Yhak! </p></td><td colspan="2"><p> footer </p></td></tr>
</tfoot>
</table>
`, "" },
    { `+-+-+--+
|      |
+      +
|      |
+--X---+`,
"", "Text2Table: square.tourist: quad.cornerA {4 3} != '+'" },
    { "", "", "Text2Table: empty input" },

    { `+-------------------------------------------------+
| una tabla realmente dificil de dibujar a mano   |
+------+----+----+----+---------------------------+
|  jo  | jo | jo | jo | y mas jos como encabezado |
+======+====+====+====+====+======================+
|             a            |          b           |
+------------------+-------+------+---------------+
|        c         |   d   |  e   |               |
+--------+---------+-------+------+      f        |
|   g    |      h                 |               |
+--------+--+----------------+----+-+-------------+
|           |        j       |      |             |
|  i        +--------+-------+   k  |      l      |
|           |    m   |   n   |      |             |
+~~~~~~~~~~~+~~~~~~~~+~~~~~~~+~~+~~~+~~~~~~~~~~~~~+
|      con un pie de pagina     |   o algo asi    |
+-------------------------------+-----------------+`,
      `<table border="1">
<thead>
  <tr><th colspan="13"><p> una tabla realmente dificil de dibujar a mano </p></th></tr>
  <tr><th><p> jo </p></th><th colspan="2"><p> jo </p></th><th><p> jo </p></th><th colspan="3"><p> jo </p></th><th colspan="6"><p> y mas jos como encabezado </p></th></tr>
</thead>
<tbody>
  <tr><td colspan="8"><p> a </p></td><td colspan="5"><p> b </p></td></tr>
  <tr><td colspan="5"><p> c </p></td><td colspan="3"><p> d </p></td><td colspan="3"><p> e </p></td><td colspan="2" rowspan="2"><p> f </p></td></tr>
  <tr><td colspan="2"><p> g </p></td><td colspan="9"><p> h </p></td></tr>
  <tr><td colspan="3" rowspan="2"><p> i </p></td><td colspan="6"><p> j </p></td><td colspan="3" rowspan="2"><p> k </p></td><td rowspan="2"><p> l </p></td></tr>
  <tr><td colspan="3"><p> m </p></td><td colspan="3"><p> n </p></td></tr>
</tbody>
<tfoot>
  <tr><td colspan="10"><p> con un pie de pagina </p></td><td colspan="3"><p> o algo asi </p></td></tr>
</tfoot>
</table>
`, "" },
    { `+----------------------------------------------------------------------------------------------------+
|                                           Proposito                                                |
+-----------------+-----------------------+---------------------------+------------------------------+
|                 | De Creacion           | Estructurales             | De Comportamiento            |
+========+========+=======================+===========================+==============================+
| Ambito | Clase  | @l(#Factory Method)   | @l(#Adapter) (de clases)  | @l(#Interpreter)             |
|        |        +-----------------------+---------------------------+------------------------------+
|        |        |                       |                           | @l(#Template Method)         |
|        +--------+-----------------------+---------------------------+------------------------------+
|        | Objeto | @l(#Abstract Factory) | @l(#Adapter) (de objetos) | @l(#Chain of Responsibility) |
|        |        +-----------------------+---------------------------+------------------------------+
|        |        | @l(#Builder)          | @l(#Bridge)               | @l(#Command)                 |
|        |        +-----------------------+---------------------------+------------------------------+
|        |        | @l(#Prototype)        | @l(#Compositive)          | @l(#Iterator)                |
|        |        +-----------------------+---------------------------+------------------------------+
|        |        | @l(#Singleton)        | @l(#Decorator)            | @l(#Mediator)                |
|        |        +-----------------------+---------------------------+------------------------------+
|        |        |                       | @l(#Facade)               | @l(#Memento)                 |
|        |        +-----------------------+---------------------------+------------------------------+
|        |        |                       | @l(#Flyweight)            | @l(#Observer)                |
|        |        +-----------------------+---------------------------+------------------------------+
|        |        |                       | @l(#Proxy)                | @l(#State)                   |
|        |        +-----------------------+---------------------------+------------------------------+
|        |        |                       |                           | @l(#Visitor)                 |
+--------+--------+-----------------------+---------------------------+------------------------------+`,
      `<table border="1">
<thead>
  <tr><th colspan="5"><p> Proposito </p></th></tr>
  <tr><th colspan="2"><p> </p></th><th><p> De Creacion </p></th><th><p> Estructurales </p></th><th><p> De Comportamiento </p></th></tr>
</thead>
<tbody>
  <tr><td rowspan="10"><p> Ambito </p></td><td rowspan="2"><p> Clase </p></td><td><p> @l(#Factory Method) </p></td><td><p> @l(#Adapter) (de clases) </p></td><td><p> @l(#Interpreter) </p></td></tr>
  <tr><td><p> </p></td><td><p> </p></td><td><p> @l(#Template Method) </p></td></tr>
  <tr><td rowspan="8"><p> Objeto </p></td><td><p> @l(#Abstract Factory) </p></td><td><p> @l(#Adapter) (de objetos) </p></td><td><p> @l(#Chain of Responsibility) </p></td></tr>
  <tr><td><p> @l(#Builder) </p></td><td><p> @l(#Bridge) </p></td><td><p> @l(#Command) </p></td></tr>
  <tr><td><p> @l(#Prototype) </p></td><td><p> @l(#Compositive) </p></td><td><p> @l(#Iterator) </p></td></tr>
  <tr><td><p> @l(#Singleton) </p></td><td><p> @l(#Decorator) </p></td><td><p> @l(#Mediator) </p></td></tr>
  <tr><td><p> </p></td><td><p> @l(#Facade) </p></td><td><p> @l(#Memento) </p></td></tr>
  <tr><td><p> </p></td><td><p> @l(#Flyweight) </p></td><td><p> @l(#Observer) </p></td></tr>
  <tr><td><p> </p></td><td><p> @l(#Proxy) </p></td><td><p> @l(#State) </p></td></tr>
  <tr><td><p> </p></td><td><p> </p></td><td><p> @l(#Visitor) </p></td></tr>
</tbody>
</table>
`, "" },
  }

  for _, d := range data {
    out, err := TextTable2HtmlTable( d.in )
    if out != d.out || (err != nil && d.err != err.Error()) {
      t.Errorf( "\nTest rangeCells() (err: %v)\n in =\n%s\n out =\n%s\n exp =\n%s\n", err, d.in, out, d.out )
    }
  }
}

func TestPullNextPoint( t *testing.T ){
  data := []struct {
    in map[point]bool
    out point
  } {
    { map[point]bool{ point{0,0}: true, }, point{0,0} },
    { map[point]bool{
      point{2,5}: true,
      point{2,7}: true,
      point{2,4}: true,
      point{2,1}: true,
      point{2,9}: true,
    },
      point{2,1} },
    { map[point]bool{
      point{7,3}: true,
      point{4,3}: true,
      point{0,3}: true,
      point{3,3}: true,
      point{1,3}: true,
    },
      point{0,3} },
    { map[point]bool{
      point{1,3}: true,
      point{2,3}: true,
      point{2,4}: true,
      point{2,0}: true,
      point{4,7}: true,
      point{4,3}: true,
      point{4,4}: true,
    },
      point{1,3} },
  }

  for _, d := range data {
    if out := pullNextPoint( d.in ); *out != d.out {
      t.Errorf( "Test pullNextPoint( %v )\n  out = %v\n  dout = %v\n", d.in, *out, d.out )
    }
  }
}


func (s *quad) diffPoints( points [4]point ) bool {
  if s.cornerD !=  points[0] { return true }
  if s.cornerC !=  points[1] { return true }
  if s.cornerB !=  points[2] { return true }
  if s.cornerA !=  points[3] { return true }


  return false
}

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
