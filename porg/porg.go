package porg

import (
  "bytes"
  "github.com/nasciiboy/txt"
)

func UnPorg( file string ) string {
  buff := new(bytes.Buffer)
  buff.Grow( len( file ) )
  for _, line := range txt.GetLines( file ) {
    if len( line ) > 0 {
      if line[0] != '#' {
        buff.WriteString( line )
        buff.WriteByte( '\n' )
      }
    } else {
      buff.WriteByte( '\n' )
    }
  }
  return buff.String()
}
