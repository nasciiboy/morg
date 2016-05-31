#include <stdio.h>
#include "exportToHtml.h"

int main( int argc, char **argv ){
  int e = 0;

  while( --argc > 0 && (*++argv)[0] == '-' )
     while( *++argv[ 0 ] )
      switch( **argv ){
      case 'e': e = 1; break;
      default:
        fprintf( stderr, "morg: illegal option %c\n", **argv );
        return 0;
      }

  if( argc == 0 ){
    fprintf( stderr, "Usage: morg -e file\n" );
    return 0;
  }

  exportToHtml( *argv++ );

  return 0;
}
