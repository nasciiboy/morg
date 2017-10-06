package katana

import (
  "testing"
  "bytes"
  "fmt"
)

func TestStr2Arg( t *testing.T ){
  data := [...]struct{
    in, name string
    args []Arg
  }{
    { "", "", []Arg{} },
    { "p", "p", []Arg{} },
    { "p()", "p", []Arg{} },
    { "p[]", "p", []Arg{} },
    { "p<>", "p", []Arg{} },
    { "p{}", "p", []Arg{} },
    { "p(true)", "p", []Arg{{"true", Bool},} },
    { `p("hola")`, "p", []Arg{{`hola`, String},} },
    { "p(`hola`)", "p", []Arg{{"hola", RawString},} },
    { "p(123456)", "p", []Arg{{"123456", Int},} },
    { "p(0x5f5f)", "p", []Arg{{"0x5f5f", Hexadecimal},} },
    { "p(0x5478)", "p", []Arg{{"0x5478", Hexadecimal},} },
    { "p('\\0')", "p", []Arg{{"'\\0'", Char},} },
    { "p('\\c')", "p", []Arg{{"'\\c'", Char},} },
    { "p(.012345)", "p", []Arg{{".012345", Float},} },

    { " \t_param \t", "_param", []Arg{} },
    { " \t_param \t()", "_param", []Arg{} },
    { " \t_param \t(true)", "_param", []Arg{{"true", Bool},} },
    { " \t_param \t(\"hola\")", "_param", []Arg{{`hola`, String},} },
    { " \t_param \t(`hola`)", "_param", []Arg{{"hola", RawString},} },
    { " \t_param \t(123456)", "_param", []Arg{{"123456", Int},} },
    { " \t_param \t(0x5f5f)", "_param", []Arg{{"0x5f5f", Hexadecimal},} },
    { " \t_param \t(0x5478)", "_param", []Arg{{"0x5478", Hexadecimal},} },
    { " \t_param \t('\\0')", "_param", []Arg{{"'\\0'", Char},} },
    { " \t_param \t('\\c')", "_param", []Arg{{"'\\c'", Char},} },
    { " \t_param \t(.012345)", "_param", []Arg{{".012345", Float},} },

    { " _param", "_param", []Arg{} },
    { " _param[]", "_param", []Arg{} },
    { " _param[true]", "_param", []Arg{{"true", Bool},} },
    { " _param<\"hola\">", "_param", []Arg{{`hola`, String},} },
    { " _param{`hola`}", "_param", []Arg{{"hola", RawString},} },

    { " _param( \"hola\", 123 )", "_param", []Arg{{`hola`, String},{"123", Int}, } },
    { " _param( float, \"hola\", 123 )", "_param", []Arg{{"float", Ident},{`hola`, String},{"123", Int}, } },
    { " _param( float, .21e+12, \"hola\", 123 )", "_param", []Arg{{"float", Ident},{".21e+12", Float},{`hola`, String},{"123", Int}, } },
    { " _param( true, float, .21e+12, \"hola\", 123 )", "_param", []Arg{{"true", Bool},{"float", Ident},{".21e+12", Float},{`hola`, String},{"123", Int}, } },
    { " _param( truer, `raw`, float, .21e+12, \"hola\", 123 )", "_param", []Arg{{"truer", Ident},{"raw", RawString},{"float", Ident},{".21e+12", Float},{`hola`, String},{"123", Int}, } },
    { " _param( /*truer, `raw`, float, .21e+12, \"hola\", 123*/ )", "_param", []Arg{} },

    { " _param( , )", "_param", []Arg{{"", Empty},{"", Empty}} },
    { " _param( time, )", "_param", []Arg{{"time", Ident},{"", Empty}} },
    { " _param( time,,true )", "_param", []Arg{{"time", Ident},{"", Empty},{"true", Bool}} },
    { " _param( ,,12546 )", "_param", []Arg{{"", Empty},{"", Empty},{"12546", Int}} },

    { "p{", "p", []Arg{} },
    { "p[", "p", []Arg{} },
    { "p<", "p", []Arg{} },
    { "p(", "p", []Arg{} },
    { "p}", "p", []Arg{} },
    { "p]", "p", []Arg{} },
    { "p>", "p", []Arg{} },
    { "p)", "p", []Arg{} },

    { "p{]", "p", []Arg{} },
    { "p[}", "p", []Arg{} },
    { "p<)", "p", []Arg{} },
    { "p(}", "p", []Arg{} },

    { "p > hela", "p", []Arg{} },
    { "param > hola", "param", []Arg{} },
  }

  for _, d := range data {
    w, _ := new(Scanner).NewSrc( d.in ).QuietSplash().Init().GetArgType()
    if w == nil {
      if d.in != "" { t.Errorf( "TestStr2Arg( %q ) == nil", d.in ) }
      continue
    }

    e := new( bytes.Buffer )
    if w.Name != d.name {
      fmt.Fprintf( e, "Name    ouput %q\n     expected %q\n", w.Name, d.name )
    }
    if argsHasDifferent( w.Args, d.args ) {
      fmt.Fprintf( e, "argData ouput %v\n     expected %v\n", w.Args, d.args )
    }

    if e.Len() > 0 {
      t.Errorf( "TestStr2Arg( %q )\n%s", d.in, e )
    }
  }

}

func TestGetArgs( t *testing.T ){
  data := [...]struct{
    in string
    args []ArgType
  }{
    { "argo)",
      []ArgType{
        { Name: "argo"},
      },
    },
    { "param()",
      []ArgType{
        { Name: "param"},
      },
    },
    { "algo alguien algul",
      []ArgType{
        { Name: "algo"},
        { Name: "alguien"},
        { Name: "algul"},
      },
    },
    { "algo alguien[ 15, .15 ] algul",
      []ArgType{
        { Name: "algo"},
        { Name: "alguien", Args: []Arg{{"15", Int},{".15", Float}}},
        { Name: "algul"},
      },
    },
    { "alguien[ 15, .15 ] algo algul()",
      []ArgType{
        { Name: "alguien", Args: []Arg{{"15", Int},{".15", Float}}},
        { Name: "algo"},
        { Name: "algul"},
      },
    },
    { "alguien[ 15, .15 ] algo() algul()",
      []ArgType{
        { Name: "alguien", Args: []Arg{{"15", Int},{".15", Float}}},
        { Name: "algo"},
        { Name: "algul"},
      },
    },
    { "argo ergo() ergi(21)",
      []ArgType{
        { Name: "argo"},
        { Name: "ergo"},
        { Name: "ergi", Args: []Arg{{"21", Int}}},
      },
    },
    { "toc(true) fancyCode",
      []ArgType{
        { Name: "toc", Args: []Arg{{"true", Bool}}},
        { Name: "fancyCode" },
      },
    },
    { "toc(true) fancyCode( \"nascii\" )",
      []ArgType{
        { Name: "toc", Args: []Arg{{"true", Bool}}},
        { Name: "fancyCode", Args: []Arg{{ `nascii`, String}}},
      },
    },
    { `Complex( .123, true, "type", func ) hShif() hShif( 5 ) hShif( -78 ) fancyCode( "nascii" )`,
      []ArgType{
        { Name: "Complex",
          Args: []Arg{{".123", Float},{"true", Bool},{`type`, String},{"func", Ident}} },
        { Name: "hShif" },
        { Name: "hShif", Args: []Arg{{"5", Int}}},
        { Name: "hShif", Args: []Arg{{"-78", Int}}},
        { Name: "fancyCode", Args: []Arg{{`nascii`, String}}},
      },
    },
  }

  for _, d := range data {
    args := new(Scanner).NewSrc( d.in ).QuietSplash().Init().GetArgs()

    e := new( bytes.Buffer )
    if commArgsHasDifferent( args, d.args ) {
      fmt.Fprintf( e, "Args    ouput %v\n     expected %v\n", args, d.args )
    }

    if e.Len() > 0 {
      t.Errorf( "TestGetArgs( %q )\n%s", d.in, e )
    }
  }

}

func argsHasDifferent( a, b []Arg ) bool {
  if len( a ) != len( b ) { return true }

  for i, d := range a {
    if d.Data != b[i].Data { return true }
    if d.Type != b[i].Type { return true }
  }

  return false
}

func commArgsHasDifferent( a, b []ArgType ) bool {
  if len( a ) != len( b ) { return true }

  for i, d := range a {
    if d.Name != b[i].Name { return true }
    if len( d.Args ) != len( b[i].Args ) { return true }

    for x, y := range d.Args {
      if y.Data != b[i].Args[x].Data { return true }
      if y.Type != b[i].Args[x].Type { return true }
    }
  }

  return false
}
