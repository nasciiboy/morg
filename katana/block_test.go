package katana

import (
  "bytes"
  "fmt"
  "testing"
)

func TestGetBlock( t *testing.T ){
  data := [...]struct{
    in string
    body string
    comm string
    args ArgMap
    head string
    attach []string
    w int
  } {

    { "..title param > hol1", "", "title", ArgMap{ "param": nil }, " hol1", nil, 20 },
    { "..title param > hol2", "", "title", ArgMap{ "param": nil }, " hol2", nil, 20 },
    { "..title param > hol3\n", "", "title", ArgMap{ "param": nil }, " hol3\n", nil, 21 },
    { "..title param > hola loli", "", "title", ArgMap{ "param": nil }, " hola loli", nil, 25 },
    { "..title param > hola loli   \n", "", "title", ArgMap{ "param": nil }, " hola loli   \n", nil, 29 },
    { "..title param(} > hola loli   \n", "", "title", ArgMap{ "param": nil }, " hola loli   \n", nil, 31 },
    { "..title param(] > hola loli   \n", "", "title", ArgMap{ "param": nil }, " hola loli   \n", nil, 31 },
    { "..title param(  > hola loli   \n", "", "", nil, "", nil, 0 },
    { "..title param[) > hola loli   \n", "", "title", ArgMap{ "param": nil }, " hola loli   \n", nil, 31 },
    { "..title param{  > hola   loli   \n", "", "", nil, "", nil, 0 },
    { "..title param< > hola loli   \n", "", "", nil, "", nil, 0 },
    { "..title > hola", "", "title", nil, " hola", nil, 14 },
    { "..title > hola\n", "", "title", nil, " hola\n", nil, 15 },
    { "..title > hola\t\n", "", "title", nil, " hola\t\n", nil, 16 },
    { "..title > hola  \t\n", "", "title", nil, " hola  \t\n", nil, 18 },
    { "..title >\t\t  hola  \t\n", "", "title", nil, "\t\t  hola  \t\n", nil, 21 },
    { "..title >\t\t  hola \t hey  \t\n", "", "title", nil, "\t\t  hola \t hey  \t\n", nil, 27 },
    { "..title > hola\nhey", "", "title", nil, " hola\n", nil, 15 },
    { "..title > hola\n  hey", "", "title", nil, " hola\n  hey", nil, 20 },
    { "..title > hola\t  \n  \they\t", "", "title", nil, " hola\t  \n  \they\t", nil, 25 },
    { "..title param > hola\t  \n  \they\t", "", "title", ArgMap{ "param": nil }, " hola\t  \n  \they\t", nil, 31 },
    { "..title param[] > hola\t  \n  \they\t", "", "title", ArgMap{ "param": nil }, " hola\t  \n  \they\t", nil, 33 },
    { "..title param() parom(21) > hola\t  \n  \they\t", "", "title",
      ArgMap{
        "param": nil,
        "parom": []Arg{{"21", Int}}},
      " hola\t  \n  \they\t", nil, 43 },
    { "..title num[125] emp(i, \"o\", `yes`, 0.235) l<5> rex{true} > hola\t  \n  \they\t", "", "title",
      ArgMap{
         "num": []Arg{{"125", Int}},
         "emp": []Arg{{"i", Ident},{`o`, String},{"yes", RawString},{"0.235", Float}},
         "l": []Arg{{"5", Int}},
         "rex": []Arg{{"true", Bool}}},
      " hola\t  \n  \they\t", nil, 75 },

//// BlockDefault

    {
`..code > arg
  body`,
"  body", "code", nil, " arg\n", nil, 19 },
    {
`..code > arg
  body

`,
"  body\n", "code", nil, " arg\n", nil, 21 },
    {
`..code > arg
  body



^-{the empty zone}
`,
"  body\n", "code", nil, " arg\n", nil, 23 },
    {
`..code > arg
  body

  xterm

^-{the empty zone}
`,
"  body\n\n  xterm\n", "code", nil, " arg\n", nil, 30 },
    {
`..code > arg
  body

  xterm`,
"  body\n\n  xterm", "code", nil, " arg\n", nil, 28 },
    {
`..code > arg
  body
  e
< code..`,
"  body\n  e\n", "code", nil, " arg\n", nil, 32 },
    {
`..code > arg

< code..`,
"\n", "code", nil, " arg\n", nil, 22 },
    {
`..code num[125] rex{true} l<5>  emp(i, "o", 0xF57F, 0.235) > arg
  hey morg
  blog
< code..`,
      "  hey morg\n  blog\n", "code",
      ArgMap{
        "num" : []Arg{{"125", Int}},
        "rex" : []Arg{{"true", Bool}},
        "l"   : []Arg{{"5", Int}},
        "emp" : []Arg{{"i", Ident},{`o`, String},{"0xF57F", Hexadecimal},{"0.235", Float}}},
      " arg\n", nil, 91 },
    {
`..code > empty



`,
"", "code", nil, " empty\n", nil, 18 },

////// BlockArgsBody
    {
`..figure > arg
  arge argo
  arge2`,
"", "figure", nil, " arg\n  arge argo\n  arge2", nil, 34 },
    {
`..figure > arg
  arge argo
  arge2
`,
"", "figure", nil, " arg\n  arge argo\n  arge2\n", nil, 35 },
    {
`..figure > arg
  body

`,
"", "figure", nil, " arg\n  body\n", nil, 23 },
    {
`..figure > arg

  body`,
"\n  body", "figure", nil, " arg\n", nil, 22 },
    {
`..figure > arg

  body
  e
< figure..`,
"\n  body\n  e\n", "figure", nil, " arg\n", nil, 37 },
    {
`..figure > arg

< figure..`,
"\n", "figure", nil, " arg\n", nil, 26 },
    {
`..figure > arg


< figure..`,
"\n\n", "figure", nil, " arg\n", nil, 27 },

    {
`..figure > arg
  arge argo
  arge2
hey listen`,
"", "figure", nil, " arg\n  arge argo\n  arge2\n", nil, 35 },
    {
`..figure > arg
  arge argo
  arge2
 hey listen
`,
"", "figure", nil, " arg\n  arge argo\n  arge2\n", nil, 35 },
    {
`..figure > arg
  body

..toposoftwarier >
`,
"", "figure", nil, " arg\n  body\n", nil, 23 },
    {
`..figure > arg

  body
wanderlast`,
"\n  body\n", "figure", nil, " arg\n", nil, 23 },
    {
`..figure > arg

  body
  e
< figure..
çø§no
`, "\n  body\n  e\n", "figure", nil, " arg\n", nil, 38 },
    {
`..figure > arg

< figure..

trigueriteromo ateh ua
e  werdertum
< taretth..
`,
"\n", "figure", nil, " arg\n", nil, 27 },
    {
`..figure > Lorem ipsum es el texto que se usa habitualmente en diseño gráfico en
  demostraciones de tipografías o de borradores de diseño para probar el diseño
  visual antes de insertar el texto final.

..figure > 02`,
      "",
      "figure", nil, " Lorem ipsum es el texto que se usa habitualmente en diseño gráfico en\n  demostraciones de tipografías o de borradores de diseño para probar el diseño\n  visual antes de insertar el texto final.\n", nil, 210 },
    {
           "..figure >\n  demostraciones de tipografías o de borradores de diseño para probar el diseño\n  visual antes de insertar el texto final.\n..figure > 02",
      "", "figure", nil, "\n  demostraciones de tipografías o de borradores de diseño para probar el diseño\n  visual antes de insertar el texto final.\n", nil, 137 },
    {
`..figure > 01 Lorem
  texto final.


..figure > 02`, "", "figure", nil, " 01 Lorem\n  texto final.\n", nil, 37 },

/// attach
    {
`..code > go
  fmt.Printf( "hola mundo" )
<>
  output
<>
  other output
< code..
`,

"  fmt.Printf( \"hola mundo\" )\n", "code", nil, " go\n",
[]string{ "  output\n", "  other output\n" }, 80 },
    {
`..code > go
  fmt.Printf( "hola mundo" )
<>
  output
<>
  other output



hola que hace
`,

"  fmt.Printf( \"hola mundo\" )\n", "code", nil, " go\n",
[]string{ "  output\n", "  other output\n" }, 74 },



/// alien

    { "..comm >", "", "comm", nil, "", nil, 8 },
    { "..comm>", "", "comm", nil, "", nil, 7 },
    { "..comm>slim", "", "comm", nil, "slim", nil, 11 },
    { "..åøgø >", "", "åøgø", nil, "", nil, 11 },
    { "..12 >", "", "", nil, "", nil, 0 },
    { "..1param >", "", "", nil, "", nil, 0 },
    { "..®param >", "", "", nil, "", nil, 0 },
    { "..p®ram >", "", "", nil, "", nil, 0 },
    { "..p®ram >", "", "", nil, "", nil, 0 },
    { "..世界 ra > multi", "", "世界", ArgMap{ "ra": nil }, " multi", nil, 19 },
  }

  for _, d := range data {
    skan := NewScanner( d.in ).QuietSplash().Init()
    block := skan.GetBlock()

    if block == nil && d.w == 0 { continue }

    // if block == nil {
    //   t.Errorf( "TestGetBlockArgsBody( %q ) == nil", d.in )
    //   continue
    // }

    e := new( bytes.Buffer )
    if block.Comm.Text != d.comm {
      fmt.Fprintf( e, "Command ouput %q\n     expected %q\n", block.Comm.Text, d.comm )
    }
    if commArgMapHasDifferent( block.Args, d.args ) {
      fmt.Fprintf( e, "Args    ouput %v\n     expected %v\n", block.Args, d.args )
    }
    if block.Head.Text() != d.head {
      fmt.Fprintf( e, "Head    ouput %q\n     expected %q\n", block.Head.Text(), d.head )
    }
    if block.Body.Text() != d.body {
      fmt.Fprintf( e, "Body    ouput %q\n     expected %q\n", block.Body.Text(), d.body )
    }

    if skan.RunePos != d.w {
      fmt.Fprintf( e, "Len     ouput %d\n     expected %d\n", skan.RunePos, d.w )
    }
    if attachHasDifferent( t, block.Attach, d.attach ) {
      fmt.Fprintf( e, "Attach  ouput != expected\n" )
    }

    if e.Len() > 0 {
      t.Errorf( "getSetupCommand( %q )\n%s", d.in, e )
    }
  }
}

func commArgMapHasDifferent( a, b ArgMap ) bool {
  if len( a ) != len( b ) { return true }

  for key, args := range a {
    if argsHasDifferent( args, b[key] ) { return true }
  }

  return false
}

func attachHasDifferent( t *testing.T, a []MiniBlock, b []string ) bool {
  if len( a ) != len( b ) { return true }

  f := false
  for i, mb := range a {
    if mb.Body.Text() != b[i] {
      t.Errorf( "commarg\na: %q\nb: %q", mb.Body.Text(), b[i] )
      f = true
    }
  }

  return f
}
