package main

import (
  "os"
  "fmt"
  "path"
  "strings"
  "io/ioutil"

  "github.com/nasciiboy/morg/biskana"
  "github.com/nasciiboy/morg/nirvana"
)

const suffix = ".morg"

func main(){
  if len(os.Args) < 2 {
    fmt.Fprintf( os.Stderr, "morg: morg command file-A file-B ...\n" )
    os.Exit( 1 )
  }

  switch os.Args[1] {
  case "export": toBiskana( os.Args[2:] )
  case "tui"   : toKatana ( os.Args[2:] )
    default:
    fmt.Fprintf( os.Stderr, "command: %s no found\n", os.Args[1] )
    fmt.Fprintf( os.Stderr, "available commands: \"export\" and \"tui\"\n" )

    os.Exit( 1 )
  }
}

func toKatana( files []string ){
  for _, inputFileName := range files {
    inputBytes, err := ioutil.ReadFile( inputFileName )
    if err != nil {
      fmt.Fprintf( os.Stderr, "morg: Couldn't open '%s', error: %v\n", inputFileName, err )
      continue
    }

    tui.Show( string(inputBytes) )
  }
}

func toBiskana( files []string ){
  for _, inputFileName := range files {
    inputBytes, err := ioutil.ReadFile( inputFileName )
    if err != nil {
      fmt.Fprintf( os.Stderr, "morg: Couldn't open '%s', error: %v\n", inputFileName, err )
      continue
    }

    pwd, _          := os.Getwd()
    outputBaseName  := path.Base( inputFileName )
    if strings.HasSuffix( outputBaseName, suffix ) {
      outputBaseName = strings.TrimSuffix( outputBaseName, suffix )
    }

    outputFileName  := path.Join( pwd, outputBaseName + ".html" )

    outputFile, err := os.Create( outputFileName )
    if err != nil {
      fmt.Fprintf( os.Stderr, "morg: Couldn't create '%s', error: %v\n", outputFileName, err )
      continue
    }

    outputBytes := []byte( biskana.Export( string(inputBytes), biskana.HTML ) )

    _, err = outputFile.Write( outputBytes )
    if err != nil {
      fmt.Fprintf( os.Stderr, "morg: Couldn't write '%s', error: %v\n", outputFileName, err )
      continue
    }

    outputFile.Close()
  }
}
