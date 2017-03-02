package main

import (
  "os"
  "fmt"
  "path"
  "strings"
  "io/ioutil"
  "github.com/nasciiboy/morg/biskana"
)

const suffix = ".morg"

func main(){
  if len(os.Args) == 1 {
    fmt.Fprintf( os.Stderr, "morg: morg file-A file-B ...\n" )
    os.Exit( 1 )
  }

  for _, inputFileName := range os.Args[1:] {
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

    fileOutput, err := os.Create( outputFileName )
    if err != nil {
			fmt.Fprintf( os.Stderr, "morg: Couldn't open '%s', error: %v\n", outputFileName, err )
      continue
    }

    outputBytes := []byte( biskana.MakeHtml( string(inputBytes), outputBaseName ) )

    _, err = fileOutput.Write( outputBytes )
    if err != nil {
			fmt.Fprintf( os.Stderr, "morg: Couldn't write '%s', error: %v\n", outputFileName, err )
      continue
    }

    fileOutput.Close()
  }
}
