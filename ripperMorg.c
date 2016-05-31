#include "regexp3.h"
#include "fileUtils.h"
#include "charUtils.h"
#include "ripperMorg.h"

#define BUFSIZE 65536

static void strCat( char *dest, char *src ){
  while( *dest ) dest++;
  while( (*dest++ = *src++) );
}

static void cpySetup( char *section, char *reccord ){
  regexp3( reccord, "^@\\w+\\s+<(\\s*[^\\s]+)*>");
  cpyCatch( section, 1 );
}

static void cpyComment( char *section, char *reccord ){
  regexp3( reccord, "^@\\s+<(\\s*[^\\s]+)*>");
  if( strLen( section ) ){
    strCat( section, "\n" );
    cpyCatch( section + strLen( section ), 1 );
  }
  else
    cpyCatch( section, 1 );
}

static void cpyTable( char *section, char *reccord ){
  regexp3( reccord, "^ *<\\|(((=|\\-)+|[^\\|]+)\\|)+>\\s*$");
  if( *section ){
    strCat( section, "\n" );
    cpyCatch( section + strLen( section ), 1 );
  }
  else
    cpyCatch( section, 1 );
}

static void cpyHeadline( char *section, char *reccord ){
  regexp3( reccord, "^\\*+\\s+<(\\s*[^\\s]+)*>");
  cpyCatch( section, 1 );
}

static void cpySubHeadline( char *section, char *reccord ){
  regexp3( reccord, "^\\s+@\\s+<(\\s*[^\\s]+)*>");
  cpyCatch( section, 1 );
}

static void cpyText( char *section, char *reccord ){
  regexp3( reccord, "^\\s*<(\\s*[^\\s]+)*>");
  cpyCatch( section, 1 );
}

static void cpyUList( char *section, char *reccord ){
  regexp3( reccord, "^\\s*(\\-|\\+|\\>)\\s+<(\\s*[^\\s]+)*>" );
  cpyCatch( section, 1 );
}

static void cpyOList( char *section, char *reccord ){
  regexp3( reccord, "^\\s*(\\d+|\\w)(\\.|\\))\\s+<(\\s*[^\\s]+)*>" );
  cpyCatch( section, 1 );
}

static void cpyBlockContent( char *section, char *reccord, int deep ){
  strCpy( section + strLen( section ), reccord + deep );
  strCat( section, "\n" );
}

static void toSection( char *section, char *reccord ){
  regexp3( reccord, "^\\s*<(\\s*[^\\s]+)*>");
  cpyCatch( section + strLen( section ), 1 );
}

static void addToSection( char *section, char *reccord ){
  regexp3( reccord, "^\\s*<(\\s*[^\\s]+)*>");
  strCat( section, " " );
  cpyCatch( section + strLen( section ), 1 );
}

static enum morg_SECTYPE whoIsThere( char *section ){
  if ( *section == 0 ) return morg_EMPTY;

  if     (regexp3( section, "^@title\\s"       )) return morg_TITLE;
  else if(regexp3( section, "^@subtitle\\s"    )) return morg_SUBTITLE;
  else if(regexp3( section, "^@author\\s"      )) return morg_AUTHOR;
  else if(regexp3( section, "^@mail\\s"        )) return morg_MAIL;
  else if(regexp3( section, "^@language\\s"    )) return morg_LANGUAGE;
  else if(regexp3( section, "^@charset\\s"     )) return morg_CHARSET;
  else if(regexp3( section, "^@css\\s"         )) return morg_CSS;
  else if(regexp3( section, "^@keywords\\s"    )) return morg_KEYWORDS;
  else if(regexp3( section, "^@description\\s" )) return morg_DOCDESCRIPTION;
  else if(regexp3( section, "^(@\\s+.*|@\\s*)$")) return morg_COMMENT;
  else if(regexp3( section, "^\\*+\\s"         )) return morg_HEADLINE;
  else if(regexp3( section, "^ +@\\s+"         )) return morg_SUBHEADLINE;
  else if(regexp3( section, "^ *\\>\\s+."      )) return morg_DIALOG;
  else if(regexp3( section, "^ *(\\-|\\+)\\s+" )) return morg_ULIST;
  else if(regexp3( section, "^ *(\\d+|\\w)(\\.|\\))\\s+" )) return morg_OLIST;
  else if(regexp3( section, "^ *\\|(((=|\\-)+|[^\\|]+)\\|)+\\s*$" )) return morg_TABLE;
  else if(regexp3( section, "^ *:: !( (::|\\<:|:\\>|\\>:\\<) )+ (::|\\<:|:\\>|\\>:\\<) .*" )) return morg_DESCRIPTION;
  else if(regexp3( section, "^ *[\\<^_\\>]{2}\\s+!( (\\<:|::|:\\>|\\>:\\<) )+ (\\<:|::|:\\ >|\\>:\\<) .*" )) return morg_MEDIA;
  else if(regexp3( section, "^ *@{1,2}[^@\x01-\x20\\&\\{\\(\\<\\[]+[\\{\\(\\<\\[]" )) return morg_TEXT;
  else if(regexp3( section, "^ *@end\\s*"      )) return morg_END_BLOCK;
  else if(regexp3( section, "^ *@\\w+\\s*"     )) return morg_BLOCK;
  else if(regexp3( section, "^\\s*$"           )) return morg_EMPTY;
  else                                            return morg_TEXT;
}

static int getDeep( char *section, enum morg_SECTYPE type ){
  switch( type ){
  case morg_ZERO: case morg_EMPTY: case morg_COMMENT:
  case morg_TITLE: case morg_SUBTITLE: case morg_AUTHOR: case morg_MAIL:
  case morg_LANGUAGE: case morg_CHARSET:
  case morg_CSS: case morg_KEYWORDS: case morg_DOCDESCRIPTION:
    return 0;
  case morg_HEADLINE:
    return countChars( section );
  case morg_SUBHEADLINE:
  case morg_TEXT: case morg_BLOCK: case morg_END_BLOCK: case morg_TABLE:
  case morg_ULIST: case morg_OLIST: case morg_DIALOG: case morg_DESCRIPTION: case morg_MEDIA:
  case morg_UNICORN:
    return countSpaces( section );
  }
}

enum morg_SECTYPE nextSection( FILE *morgFile, char *section ){
  if( fileTok( section, BUFSIZE, morgFile ) )
    return whoIsThere( section );
  else return 0;
}

enum morg_SECTYPE ripperMorg( FILE *morgFile, char *section, int *deep ){
  static char stock[ BUFSIZE ] = { 0 };
  enum morg_SECTYPE type;
  *section = 0; *deep = 0;

  if( *stock || fileTok( stock, BUFSIZE, morgFile ) ){
    type = whoIsThere( stock );
    switch( type ){
    case morg_TITLE: case morg_SUBTITLE: case morg_AUTHOR: case morg_MAIL:
    case morg_LANGUAGE: case morg_CHARSET:
    case morg_CSS: case morg_KEYWORDS: case morg_DOCDESCRIPTION:
      cpySetup( section, stock );
      *stock = 0;
      break;
    case morg_EMPTY:
      while( nextSection( morgFile, stock ) == morg_EMPTY );
      break;
    case morg_COMMENT:
      cpyComment( section, stock );
      while( nextSection( morgFile, stock ) == morg_COMMENT )
        cpyComment( section, stock );
      break;
    case morg_HEADLINE:
      *deep = getDeep( stock, type );
      cpyHeadline( section, stock );
      while( nextSection( morgFile, stock ) == morg_TEXT &&
             *deep == getDeep( stock, morg_TEXT ) )
        addToSection( section, stock );
      break;
    case morg_SUBHEADLINE:
      *deep = getDeep( stock, type );
      cpySubHeadline( section, stock );
      while( nextSection( morgFile, stock ) == morg_TEXT &&
             getDeep( stock, morg_TEXT ) >= *deep + 2 )
        addToSection( section, stock );
      break;
    case morg_TEXT:
      *deep = getDeep( stock, type );
      cpyText( section, stock );
      while( nextSection( morgFile, stock ) == morg_TEXT &&
             *deep == getDeep( stock, morg_TEXT ) )
        addToSection( section, stock );
      break;
    case morg_ULIST: case morg_DIALOG:
      *deep = getDeep( stock, type );
      cpyUList( section, stock );
      while( nextSection( morgFile, stock ) == morg_TEXT &&
             getDeep( stock, morg_TEXT ) >= *deep + 1 )
        addToSection( section, stock );
      break;
    case morg_OLIST:
      *deep = getDeep( stock, type );
      cpyOList( section, stock );
      while( nextSection( morgFile, stock ) == morg_TEXT &&
             getDeep( stock, morg_TEXT ) >= *deep + 1 )
        addToSection( section, stock );
      break;
    case morg_DESCRIPTION: case morg_MEDIA:
      *deep = getDeep( stock, type );
      toSection( section, stock );
      while( nextSection( morgFile, stock ) == morg_TEXT &&
             getDeep( stock, morg_TEXT ) >= *deep + 2 )
        addToSection( section, stock );
      break;
    case morg_TABLE:
      *deep = getDeep( stock, type );
      cpyTable( section, stock );
      while( nextSection( morgFile, stock ) == morg_TABLE &&
             *deep == getDeep( stock, type ) )
        cpyTable( section, stock );
      break;
    case morg_BLOCK:
      *deep = getDeep( stock, type );
      toSection( section, stock );
      strCat( section, "\n" );
      while( !(nextSection( morgFile, stock ) == morg_END_BLOCK && getDeep( stock,  morg_END_BLOCK ) == *deep) )
        cpyBlockContent( section, stock, *deep ? *deep + 2 + 1 : 2 );
      *stock = 0;
      break;
    case morg_ZERO:
      return morg_ZERO;
    }

    return type;
  }

  return morg_ZERO;
}
