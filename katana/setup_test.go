package katana

import (
  "testing"
)

func TestSetupTags( t *testing.T ){
  data := []struct{
    in string
    tags []string
  } {
    { `..tags > a, b, c, 15, "hola"`, []string{ "a", "b", "c", "15", "hola" }},
    { "..tags > a, b, c\n..tags > 15, \"hola\"", []string{ "a", "b", "c", "15", "hola" }},
    { `..tags > , b, c, 15, "hola"`, []string{ "b", "c", "15", "hola" }},
    { `..tags > ", b, c, 15, "hola"`, []string{ ", b, c, 15, ", "" }},
    { `..tags > uno dos, tres`, []string{ "uno", "tres" }},
  }

  for _, d := range data {
    r, _ := Parse( "testSetuptags", d.in )
    if cmpStringArray( r.Tags, d.tags ) {
      t.Errorf( "SetupTag( %q )\nresult   %v\nexpected %v", d.in, r.Tags, d.tags )
    }
  }
}

func TestSetupOptions( t *testing.T ){
  name := "testSetupOptions"

  in, bexpected := "", false
  r, err := Parse( name, in  )
  if r.BoolOptions[ "toc" ] != bexpected {
    t.Errorf( "SetupOptions( %q )\nresult   %t\nexpected %t", in, r.BoolOptions[ "toc" ], bexpected )
  }
  if err != "" { t.Errorf( "SetupOptions( %q )\nError: %q", in, err ) }

  in, bexpected = "..options > toc", true
  r, err = Parse( name, in  )
  if r.BoolOptions[ "toc" ] != bexpected {
    t.Errorf( "SetupOptions( %q )\nresult   %t\nexpected %t", in, r.BoolOptions[ "toc" ], bexpected )
  }
  if err != "" { t.Errorf( "SetupOptions( %q )\nError: %q", in, err ) }

  in, bexpected = "..options > toc(true)\n..options > toc( false )", false
  r, err = Parse( name, in  )
  if r.BoolOptions[ "toc" ] != bexpected {
    t.Errorf( "SetupOptions( %q )\nresult   %t\nexpected %t", in, r.BoolOptions[ "toc" ], bexpected )
  }
  if err != "" { t.Errorf( "SetupOptions( %q )\nError: %q", in, err ) }

  in, bexpected = "..options > toc(true)\n..options > toc( false )\n..options > toc()", true
  r, err = Parse( name, in  )
  if r.BoolOptions[ "toc" ] != bexpected {
    t.Errorf( "SetupOptions( %q )\nresult   %t\nexpected %t", in, r.BoolOptions[ "toc" ], bexpected )
  }
  if err != "" { t.Errorf( "SetupOptions( %q )\nError: %q", in, err ) }

  in, bexpected = "..options > toc(true)\n..options > toc( false )\n..options > toc()", true
  r, err = Parse( name, in  )
  if r.BoolOptions[ "toc" ] != bexpected {
    t.Errorf( "SetupOptions( %q )\nresult   %t\nexpected %t", in, r.BoolOptions[ "toc" ], bexpected )
  }
  if err != "" { t.Errorf( "SetupOptions( %q )\nError: %q", in, err ) }

  in, bexpected = "..options > toc(true) fancyCode", true
  r, err = Parse( name, in  )
  if r.BoolOptions[ "toc" ] != bexpected {
    t.Errorf( "SetupOptions( %q )\nresult   %t\nexpected %t", in, r.BoolOptions[ "toc" ], bexpected )
  }
  if err != "" { t.Errorf( "SetupOptions( %q )\nError: %q", in, err ) }

  in, bexpected = "..options > fancyCode toc(true)", true
  r, err = Parse( name, in  )
  if r.BoolOptions[ "toc" ] != bexpected {
    t.Errorf( "SetupOptions( %q )\nresult   %t\nexpected %t", in, r.BoolOptions[ "toc" ], bexpected )
  }
  if err != "" { t.Errorf( "SetupOptions( %q )\nError: %q", in, err ) }

  in, bexpected = "..options > fancyCode toc", true
  r, err = Parse( name, in  )
  if r.BoolOptions[ "toc" ] != bexpected {
    t.Errorf( "SetupOptions( %q )\nresult   %t\nexpected %t", in, r.BoolOptions[ "toc" ], bexpected )
  }
  if err != "" { t.Errorf( "SetupOptions( %q )\nError: %q", in, err ) }

  in, bexpected = "..options > toc(true) ", false
  r, err = Parse( name, in  )
  if r.BoolOptions[ "fancyCode" ] != bexpected {
    t.Errorf( "SetupOptionsFancy( %q )\nresult   %t\nexpected %t", in, r.BoolOptions[ "fancyCode" ], bexpected )
  }
  if err != "" { t.Errorf( "SetupOptions( %q )\nError: %q", in, err ) }

  in, bexpected = "..options > toc(true) fancyCode", true
  r, err = Parse( name, in  )
  if r.BoolOptions[ "fancyCode" ] != bexpected {
    t.Errorf( "SetupOptionsFancy( %q )\nresult   %t\nexpected %t", in, r.BoolOptions[ "fancyCode" ], bexpected )
  }
  if err != "" { t.Errorf( "SetupOptions( %q )\nError: %q", in, err ) }

  in, bexpected = "..options > fancyCode toc(true)", true
  r, err = Parse( name, in  )
  if r.BoolOptions[ "fancyCode" ] != bexpected {
    t.Errorf( "SetupOptionsFancy( %q )\nresult   %t\nexpected %t", in, r.BoolOptions[ "fancyCode" ], bexpected )
  }
  if err != "" { t.Errorf( "SetupOptions( %q )\nError: %q", in, err ) }

  in, bexpected = "..options > toc(true) fancyCode( \"nascii\" )", true
  r, err = Parse( name, in  )
  if r.BoolOptions[ "fancyCode" ] != bexpected {
    t.Errorf( "SetupOptionsFancy( %q )\nresult   %t\nexpected %t", in, r.BoolOptions[ "fancyCode" ], bexpected )
  }
  if err != "" { t.Errorf( "SetupOptions( %q )\nError: %q", in, err ) }

}

func TestSyncOpt( t *testing.T ){
  data := []struct{
    in, out ArgType
    cannon  bool
  } {
    { ArgType{ "fancyCode", nil }, ArgType{ "fancyCode", []Arg{{ "nascii", String }}}, true },
    { ArgType{ "fancyCode", []Arg{{ "nascii", String }}}, ArgType{ "fancyCode", []Arg{{ "nascii", String }}}, true },
    { ArgType{ "fancyCode", []Arg{{ "believe in Emacs", String }}}, ArgType{ "fancyCode", []Arg{{ "believe in Emacs", String }}}, true },
    { ArgType{ "fancyCode", []Arg{{ "rat nation", String }}}, ArgType{ "fancyCode", []Arg{{ "rat nation", String }}}, true },
    { ArgType{ "fancyCode", []Arg{{ "", Empty }}}, ArgType{ "fancyCode", []Arg{{ "nascii", String }}}, true },
    { ArgType{ "fancyCode", []Arg{{ "rat nation", Int }}}, ArgType{ "fancyCode", []Arg{{ "nascii", String }}}, true },
    { ArgType{ "fancyCode", []Arg{{ "1524", Octal }}}, ArgType{ "fancyCode", []Arg{{ "nascii", String }}}, true },
    { ArgType{ "fancyCode", []Arg{{ "rat nation", String },{ "1524", Octal }}}, ArgType{ "fancyCode", []Arg{{ "rat nation", String }}}, true },
    { ArgType{ "hShif", nil }, ArgType{ "hShif", []Arg{{ "0", Int }}}, true },
    { ArgType{ "hShif", []Arg{{ "nascii", Int }}}, ArgType{ "hShif", []Arg{{ "nascii", Int }}}, true },
    { ArgType{ "hShif", []Arg{{ "believe in Emacs", Int }}}, ArgType{ "hShif", []Arg{{ "believe in Emacs", Int }}}, true },
    { ArgType{ "hShif", []Arg{{ "rat nation", Int }}}, ArgType{ "hShif", []Arg{{ "rat nation", Int }}}, true },
    { ArgType{ "hShif", []Arg{{ "", Empty }}}, ArgType{ "hShif", []Arg{{ "0", Int }}}, true },
    { ArgType{ "hShif", []Arg{{ "rat nation", Int }}}, ArgType{ "hShif", []Arg{{ "rat nation", Int }}}, true },
    { ArgType{ "hShif", []Arg{{ "1524", Octal }}}, ArgType{ "hShif", []Arg{{ "0", Int }}}, true },
    { ArgType{ "hShif", []Arg{{ "rat nation", String },{ "1524", Octal }}}, ArgType{ "hShif", []Arg{{ "0", Int }}}, true },
    { ArgType{ "complex", nil }, ArgType{ "complex", []Arg{{ ".125", Float}, { "true", Bool}, { "type", String }, { "func", Ident }}}, true },
    { ArgType{ "complex", []Arg{{ "459.54", Float}}}, ArgType{ "complex", []Arg{{ "459.54", Float}, { "true", Bool}, { "type", String }, { "func", Ident }}}, true },
    { ArgType{ "complex", []Arg{{ "125", Int}, { "tarauma", Ident}, { "third", String }, { "func", Ident }}}, ArgType{ "complex", []Arg{{ ".125", Float}, { "true", Bool}, { "third", String }, { "func", Ident }}}, true },
    { ArgType{ "complex", []Arg{{ "", Empty}, { "false", Bool}, { "", Empty }, { "func", Ident }}}, ArgType{ "complex", []Arg{{ ".125", Float}, { "false", Bool}, { "type", String }, { "func", Ident }}}, true },
    { ArgType{ "complex", []Arg{{ "57.00", Float}, { "false", Bool}, { "blue", String }, { "Medical", Ident }}}, ArgType{ "complex", []Arg{{ "57.00", Float}, { "false", Bool}, { "blue", String }, { "Medical", Ident }}}, true },
    { ArgType{ "30+1", []Arg{{ "-0.5", Float},{ "Medical", Ident }}}, ArgType{ "30+1", []Arg{{ "-0.5", Float},{ "Medical", Ident }}}, false },
    { ArgType{ "yo can (not) be a cannon", nil }, ArgType{ "yo can (not) be a cannon", nil }, false },
  }

  m := map[string][]Arg {
    "fancyCode": {{ "nascii", String }},
    "hShif": {{ "0", Int }},
    "complex": {{ ".125", Float}, { "true", Bool}, { "type", String }, { "func", Ident }},
  }

  for _, d := range data {
    ds := new(doc)
    ds.Scanner = new(Scanner)
    ds.CustomError = quietSplash
    out, cnn := ds.syncOpt( d.in, m )

    if cmpArgType( out, d.out ) {
      t.Errorf( "syncOpt( %v )\nresult   %v\nexpected %v", d.in, out, d.out )
    }
    if cnn != d.cannon {
      t.Errorf( "syncOpt( %v )\nresult   %t\nexpected %t", d.in, cnn, d.cannon )
    }
  }
}

func cmpArgType( a, b ArgType ) bool {
  if a.Name != b.Name { return true }

  return argsHasDifferent( a.Args, b.Args )
}

func cmpStringArray( a, b []string ) bool {
  if len( a ) != len( b ) { return true }

  for i, d := range a {
    if d != b[i] { return true }
    if d != b[i] { return true }
  }

  return false
}
