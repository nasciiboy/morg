#include <string.h>
#include <ctype.h>

#include "fileUtils.h"

int countChars( char * reccord ){
  int i = 0;
  while( reccord[ i ] && (isgraph( reccord[ i ] ) || !isascii( reccord[ i ] )) ) i++;

  return i;
}

int countSpaces( char * reccord ){
  int i = 0;
  while( reccord[ i ] && isspace( reccord[ i ] ) ) i++;

  return i;
}


int fileTok( char * reccord, unsigned int size, FILE *stream ){
  static FILE *file = 0;
  int c, len = 0;
  if( stream != file ) file = stream;

  if( file == 0 ) c = EOF;
  else
    while( (c = fgetc( file )) != '\n' && c != EOF && size-- )
      reccord[ len++ ] = c;

  reccord[ len ] = '\0';

  if( len || c != EOF ) return 1;

  file = 0;
  return 0;
}



int reccordTok( char *reccord, unsigned int size, char * field ){
  static int i = 0;

  if( reccord[ i ] ){
    int ispc  = countSpaces( reccord + i  ) ;
    int ichar = countChars ( reccord + i + ispc );

    if( ichar > size ) ichar = size;
    if( ichar ){
      strncpy( field, reccord + i + ispc, ichar );
      field[ ichar ] = '\0';
      i += ispc + ichar;
      return 1;
    }
  }

  *field = '\0';
  return (i = 0);
}
