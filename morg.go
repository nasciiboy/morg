package main

import (
  "os"
  "fmt"
  "path"
  "strings"
  "io/ioutil"

  "github.com/nasciiboy/morg/biskana"
  "github.com/nasciiboy/morg/nirvana"
  "github.com/nasciiboy/morg/porg"
)

const (
  morgSuffix = ".morg"
  porgSuffix = ".porg"
  htmlSuffix = ".html"
)

func main(){
  if len(os.Args) < 2 {
    fmt.Fprintf( os.Stderr, "morg: morg command file-A file-B ...\n" )
    os.Exit( 1 )
  }

  switch os.Args[1] {
  case "tui"   : toNirvana( os.Args[2:] )
  case "toHtml": toBiskana( os.Args[2:], htmlSuffix, biskana.HTML )
  case "toTxt" :
    fmt.Println( "＼_(-_- 彡 -_-)_／☆･ ･ ･ ‥……━━●~*" )
    toBiskana( os.Args[2:], ".txt" , biskana.TXT  )
  case "unPorg" :
    fmt.Println( "＼_(-_- 彡 -_-)_／☆･ ･ ･ ‥……━━●~*" )
    toUnPorg( os.Args[2:] )
  case "help"  :
    fmt.Printf( "Usage   : morg command file-A file-B ...\n\n" )
    fmt.Printf( "Commands: \"tui\"    show file\n" )
    fmt.Printf( "          \"ToHtml\" export file to Html\n" )
  default:
    fmt.Fprintf( os.Stderr, "Command: %s no found\n", os.Args[1] )
    fmt.Fprintf( os.Stderr, "Available commands: \"ToHtml\" and \"tui\"\n" )
    os.Exit( 1 )
  }
}

func toNirvana( files []string ){
  for _, inputFileName := range files {
    inputBytes, err := ioutil.ReadFile( inputFileName )
    if err != nil {
      fmt.Fprintf( os.Stderr, "morg: Couldn't open '%s', error: %v\n", inputFileName, err )
      continue
    }

    nirvana.Show( string(inputBytes) )
  }
}

func toUnPorg( files []string ){
  for _, inputFileName := range files {
    inputBytes, err := ioutil.ReadFile( inputFileName )
    if err != nil {
      fmt.Fprintf( os.Stderr, "morg: Couldn't open '%s', error: %v\n", inputFileName, err )
      continue
    }

    pwd, _          := os.Getwd()
    outputBaseName  := path.Base( inputFileName )
    if strings.HasSuffix( outputBaseName, porgSuffix ) {
      outputBaseName = strings.TrimSuffix( outputBaseName, porgSuffix )
    }

    outputFileName := path.Join( pwd, outputBaseName + morgSuffix )
    outputBytes    := []byte( porg.UnPorg( string(inputBytes) ) )
    err             = ioutil.WriteFile( outputFileName, outputBytes, 0666 )
    if err != nil {
      fmt.Fprintf( os.Stderr, "morg: %v\n", err )
    }
  }
}

func toBiskana( files []string, outputPrefix string, to uint ){
  for _, inputFileName := range files {
    inputBytes, err := ioutil.ReadFile( inputFileName )
    if err != nil {
      fmt.Fprintf( os.Stderr, "morg: Couldn't open '%s', error: %v\n", inputFileName, err )
      continue
    }

    pwd, _          := os.Getwd()
    outputBaseName  := path.Base( inputFileName )
    if strings.HasSuffix( outputBaseName, morgSuffix ) {
      outputBaseName = strings.TrimSuffix( outputBaseName, morgSuffix )
    }

    outputFileName := path.Join( pwd, outputBaseName + htmlSuffix )
    outputBytes    := []byte( biskana.Export( string(inputBytes), to ) )
    err             = ioutil.WriteFile( outputFileName, outputBytes, 0666 )
    if err != nil {
      fmt.Fprintf( os.Stderr, "morg: %v\n", err )
    }
  }
}
