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
  case "tui"   : toNirvana( os.Args[2:] )
  case "toHtml": toBiskana( os.Args[2:], ".html", biskana.HTML )
  case "toTxt" :
    fmt.Println( "＼_(-_- 彡 -_-)_／☆･ ･ ･ ‥……━━●~*" )
    toBiskana( os.Args[2:], ".txt" , biskana.TXT  )
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

func toBiskana( files []string, outputPrefix string, to uint ){
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

    outputFileName  := path.Join( pwd, outputBaseName + outputPrefix )

    outputFile, err := os.Create( outputFileName )
    if err != nil {
      fmt.Fprintf( os.Stderr, "morg: Couldn't create '%s', error: %v\n", outputFileName, err )
      continue
    }

    outputBytes := []byte( biskana.Export( string(inputBytes), to ) )

    _, err = outputFile.Write( outputBytes )
    if err != nil {
      fmt.Fprintf( os.Stderr, "morg: Couldn't write '%s', error: %v\n", outputFileName, err )
      continue
    }

    outputFile.Close()
  }
}
