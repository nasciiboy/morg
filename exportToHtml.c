#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "exportToHtml.h"
#include "ripperMorg.h"
#include "regexp3.h"

#define SECSIZE 65536

static      int          SecLevel;
static enum morg_SECTYPE SecType;

static char mMORG[ SECSIZE ];
static char mTEXT[ SECSIZE ];
static char mCOMM[ SECSIZE ];
static char mLINK[ SECSIZE ];
static char mSIGN[ SECSIZE ];
static char mHTML[ SECSIZE ];
static char mTMP [ SECSIZE ];

static FILE *srcFile     = 0;
static FILE *tocSec      = 0;
static FILE *bodySec     = 0;

static char *title       = 0;
static char *subtitle    = 0;
static char *author      = 0;
static char *mail        = 0;
static char *language    = 0;
static char *charset     = 0;
static char *css         = 0;
static char *keywords    = 0;
static char *description = 0;

static void walk_morg    ( int level, int );
static void walk_morgList( int level       );
static void make_section ( enum morg_SECTYPE secType, int level );
static void setup();
static void open_media(int level);
static void open_section(enum morg_SECTYPE secType, int level);
static void close_section(enum morg_SECTYPE secType, int level);
static void clean_setup();

static char * to_mLINK( char * section );
static char * to_mTEXT( char * section );
static char * to_mHTML( char * section );

static char * translateMarkup( char * section );

static void make_head();
static void make_toc();
static void make_body();

void exportToHtml( char *fileName ){
  if( (srcFile = fopen( fileName, "r" )) && (bodySec = tmpfile()) ){
    SecType  = morg_ZERO;

    walk_morg( EOF, 0 );

    make_head();
    make_body();

    fclose( srcFile );
    fclose( bodySec );
    clean_setup();
  } else fprintf( stderr, "Error open: >>%s<<\n", fileName );
}

static void walk_morg( int level, int breakInHeadline ){
  while( SecType || (SecType = ripperMorg( srcFile, mMORG, &SecLevel )) )
    switch( SecType ){
    case morg_EMPTY: case morg_COMMENT:
      SecType = morg_ZERO; break;
    case morg_ULIST : case morg_OLIST: case morg_DESCRIPTION:
      if ( SecLevel > level ){ walk_morgList( SecLevel ); break; }
      else return;
    case morg_HEADLINE:
      if( breakInHeadline ) return;
    default:
      if ( SecLevel > level ) make_section( SecType, SecLevel );
      else return;
    }
}

static void walk_morgList( int level ){
  enum morg_SECTYPE listType = SecType;
  char *label = SecType == morg_DESCRIPTION ? "dl" : SecType == morg_OLIST ? "ol" : "ul";

  fprintf( bodySec, "%*s<%s>\n", level * 2 + 6, "", label );

  while( SecType || (SecType = ripperMorg( srcFile, mMORG, &SecLevel )) )
    if( SecType == morg_EMPTY || SecType == morg_COMMENT ) SecType = morg_ZERO;
    else if( SecType == morg_HEADLINE || SecLevel < level ||
             (SecLevel == level && SecType != listType) ) break;
    else make_section( SecType, SecLevel );

  fprintf( bodySec, "%*s</%s>\n", level * 2 + 6, "", label );
}

static void make_section(enum morg_SECTYPE secType, int level ){
  switch(secType){
  case morg_ZERO    : case morg_EMPTY   : case morg_COMMENT: break;
  case morg_TITLE   : case morg_SUBTITLE: case morg_AUTHOR :
  case morg_MAIL    : case morg_LANGUAGE: case morg_CHARSET:
  case morg_CSS     : case morg_KEYWORDS: case morg_DOCDESCRIPTION:
    setup();
    break;
  case morg_HEADLINE:
    open_section(secType, level );
    SecType = morg_ZERO;
    walk_morg( level, 0 );
    close_section(secType, level );
    return;
  case morg_BLOCK:
    open_section(secType, level );
    close_section(secType, level );
    break;
  case morg_ULIST:
  case morg_OLIST:
  case morg_DESCRIPTION:
  case morg_MEDIA:
    open_section(secType, level );
    SecType = morg_ZERO;
    walk_morg( level, 1 );
    close_section(secType, level );
   return;
  case morg_DIALOG:
    open_section(secType, level );
    close_section(secType, level );
    break;
  default :
    open_section(secType, level );
    close_section(secType, level );
    break;
  }

  SecType = morg_ZERO;
}

static void setup(){
  switch( SecType ){
  case morg_TITLE         : title       = strdup( mMORG ); break;
  case morg_SUBTITLE      : subtitle    = strdup( mMORG ); break;
  case morg_AUTHOR        : author      = strdup( mMORG ); break;
  case morg_MAIL          : mail        = strdup( mMORG ); break;
  case morg_LANGUAGE      : language    = strdup( mMORG ); break;
  case morg_CHARSET       : charset     = strdup( mMORG ); break;
  case morg_CSS           : css         = strdup( mMORG ); break;
  case morg_KEYWORDS      : keywords    = strdup( mMORG ); break;
  case morg_DOCDESCRIPTION: description = strdup( mMORG ); break;
  default:                                                  break;
  }
}

static void open_section(enum morg_SECTYPE secType, int level ){
  switch( secType ){
  case morg_HEADLINE:
    fprintf( bodySec, "%*s<h%d id=\"%s\"",
             level * 2 + 4, "", level + 1,
             to_mLINK( mMORG ) );
    fprintf( bodySec, ">%s</h%d>\n"
             "%*s<div class=\"outline-text-%d\" id=\"text-%d\">\n",
             translateMarkup( mMORG ),
             level + 1,
             level * 2 + 4, "",
             level + 2, level + 1);
    break;
  case morg_SUBHEADLINE:
    fprintf( bodySec, "%*s<h1 class=\"subheadline\">%s</h1>\n", level * 2 + 6, "", translateMarkup( mMORG ) );
    break;
  case morg_TEXT:
    fprintf( bodySec, "%*s<p>%s", level * 2 + 6, "", translateMarkup( mMORG ) );
    break;
  case morg_BLOCK:
    fprintf( bodySec, "%*s<div class=\"block\">\n"
             "%s",
             level * 2 + 8, "",
             mMORG );
    break;
  case morg_TABLE: break;
  case morg_ULIST: case morg_OLIST:
    fprintf( bodySec, "%*s<li>%s\n", level * 2 + 8, "", translateMarkup( mMORG ) );
    break;
  case morg_DIALOG:
    fprintf( bodySec, "%*s<div class=\"dialog\">\n"
             "%*s%s\n",
             level * 2 + 6, "",
             level * 2 + 6, "",
             mMORG );
    break;
  case morg_DESCRIPTION:
    translateMarkup( mMORG );
    regexp3( mMORG, "^\\s*::\\s+<!( (::|\\<:|:\\>|\\>:\\<) )+> <::|\\<:|:\\>|\\>:\\<> <.*>" );
    fprintf( bodySec, "%*s<dt>%.*s</dt><dd>%s\n",
             level * 2 + 8, "",
             lenCatch( 1 ), gpsCatch( 1 ),
             gpsCatch( 3 ) );
    break;
  case morg_MEDIA:
    open_media( level );
    break;
  default              : break;
  }
}

enum DIR { UP, LEFT, DOWN, RIGTH };

static void open_media( int level ){
  enum DIR pos[2] = { 0 };
  fprintf( bodySec, "%*s<div class=\"media-",
           level * 2 + 6, "" );

  for( int i = 0; i < 2; i++ )
    switch( mMORG[ i ] ){
    case '^' : fprintf( bodySec, "Up"    ); pos[i] = UP;    break;
    case '<' : fprintf( bodySec, "Left"  ); pos[i] = LEFT;  break;
    case '_' : fprintf( bodySec, "Down"  ); pos[i] = DOWN;  break;
    case '>' : fprintf( bodySec, "Rigth" ); pos[i] = RIGTH; break;
    defauld: break;
    }
  translateMarkup( mMORG );
  regexp3( mMORG, "^..\\s+<!( (\\<:|::|:\\>|\\>:\\<) )+> <\\<:|::|:\\>|\\>:\\<> <.*>" );

  fprintf( bodySec, "\">\n"
           "%*s<div class=\"media-dt\">%.*s</div>\n"
           "%*s<div class=\"media-dd\">%s\n",
           level * 2 + 8, "",
           lenCatch( 1 ), gpsCatch( 1 ),
           level * 2 + 8, "",
           gpsCatch( 3 ) );

}

static void close_section(enum morg_SECTYPE secType, int level ){
  switch( secType ){
  case morg_HEADLINE: fprintf( bodySec, "%*s</div>\n", level * 2 + 4, "" ); break;
  case morg_SUBHEADLINE: break;
  case morg_TEXT: fprintf( bodySec, "</p>\n"  ); break;
  case morg_BLOCK:
    fprintf( bodySec, "%*s</dd>\n", level * 2 + 8, "" );
    break;
  case morg_TABLE: break;
  case morg_ULIST: case morg_OLIST:
    fprintf( bodySec, "%*s</li>\n", level * 2 + 8, "" ); break;
  case morg_DIALOG:
    fprintf( bodySec, "%*s</div>\n", level * 2 + 6, "" );
    break;
  case morg_DESCRIPTION:
    fprintf( bodySec, "%*s</dd>\n", level * 2 + 8, "" );
    break;
  case morg_MEDIA:
    fprintf( bodySec, "%*s</div>\n"
             "%*s</div>\n",
             level * 2 + 8, "",
             level * 2 + 6, "" ); break;
  default              : break;
  }
}

static void make_head(){
  printf( "<!DOCTYPE html/>\n" );
  printf( "<html lang=\"%s\">\n",                              language    ? language    : "en"          );
  printf( "  <head>\n" );
  printf( "    <title>%s</title>\n",                           title       ? to_mTEXT( title )       : "Insert Coin" );
  printf( "    <meta charset=\"%s\"/>\n",                      charset     ? charset     : "UTF-8"       );
  printf( "    <meta name=\"description\" content=\"%s\"/>\n", description ? description : ""            );
  printf( "    <meta name=\"keywords\" content=\"%s\"/>\n",    keywords    ? keywords    : ""            );
  printf( "    <meta name=\"%s\"/>\n",                         author      ? author      : ""            );
  printf( "    <link rel=\"stylesheet\" href=\"%s\">\n",       css         ? css         : ""            );
  printf( "  </head>\n" );
  printf( "  <body>\n"
          "    <div id=\"content\">\n"
          "    <h1 id=\"%s\">", title ? to_mLINK( title ) : "Insert Coint\n" );
  printf( "%s</h1>\n", title ? translateMarkup(title) : "Insert Coint\n" );
}

static void make_body(){
  char c;
  rewind( bodySec);

  while( (c = fgetc( bodySec )) != EOF )
    putchar( c );

  printf( "    </div>\n"
          "  </body>\n"
          "</html>\n" );
}

static void clean_setup(){
  if( title       ) free( title       );
  if( subtitle    ) free( subtitle    );
  if( author      ) free( author      );
  if( mail        ) free( mail        );
  if( language    ) free( language    );
  if( charset     ) free( charset     );
  if( css         ) free( css         );
  if( keywords    ) free( keywords    );
  if( description ) free( description );
}

static void make_toc(){}

#include "charUtils.h"

static char * makeCommand( char *section, char *command, char *arg ){
  static char keyStyle;

  if( regexp3( section, "?<@{1,2}><[^@\x01-\x20\\&\\{\\(\\<\\[]+><[\\{\\(\\<\\[]>" ) ){
    if( lenCatch( 1 ) == 1 ){
      cpyCatch( command, 2 );

      switch( *(gpsCatch( 3 )) ){
      case '[': keyStyle = ']'; break;
      case '(': keyStyle = ')'; break;
      case '{': keyStyle = '}'; break;
      case '<': keyStyle = '>'; break;
      default: break;
      }

      strnCpy( arg, gpsCatch( 3 ) + 1, strchr( gpsCatch( 3 ) + 1, keyStyle ) - (gpsCatch( 3 ) + 1) );
    }

    return gpsCatch( 3 ) + 1;
  }

  return 0;
}

static void mvchars( char *arg, int pos ){
  for( int len = strlen( arg ) + 1; len >= 0; len-- )
    arg[ len + pos ] = arg[ len ];
}

static void sandwich( char *filled, char *sides ){
  int len = strlen( sides );
  mvchars( filled, len + 2 );
  *filled = '<';
  strncpy( filled + 1, sides, len );
  filled[ 1 + len ] = '>';
  sprintf( filled + strlen( filled ), "</%s>", sides );
}

static void markupCommand( char * command, char *arg ){
  for( ; *command; command++ ){
    switch( *command ){
    case '!' : break;
    case '"' : break;
    case '#' : break;
    case '$' : break;
    case '%' : break;
    case '&' : break;
    case '\'': sandwich( arg, "samp" ); break;
    case '(' : break;
    case ')' : break;
    case '*' : break;
    case '+' : break;
    case ',' : break;
    case '-' : break;
    case '.' : break;
    case '/' : break;
    case '0' : break;
    case '1' : break;
    case '2' : break;
    case '3' : break;
    case '4' : break;
    case '5' : break;
    case '6' : break;
    case '7' : break;
    case '8' : break;
    case '9' : break;
    case ':' : sandwich( arg, "dfn" ); break;
    case ';' : break;
    case '<' : break;
    case '=' : break;
    case '>' : break;
    case '?' : break;
    case 'a' : sandwich( arg, "abbr" ); break;
    case 'A' : break;
    case 'b' : sandwich( arg, "b" ); break;
    case 'B' : sandwich( arg, "b" ); break;
    case 'c' : sandwich( arg, "code" ); break;
    case 'C' : break;
    case 'd' : break;
    case 'D' : break;
    case 'e' : sandwich( arg, "em" ); break;
    case 'E' : break;
    case 'f' : break;
    case 'F' : break;
    case 'g' : break;
    case 'G' : break;
    case 'h' : break;
    case 'H' : break;
    case 'i' : sandwich( arg, "i" ); break;
    case 'I' : break;
    case 'j' : break;
    case 'J' : break;
    case 'k' : sandwich( arg, "kbd" ); break;
    case 'K' : break;
    case 'l' : break;
    case 'L' : break;
    case 'm' : break;
    case 'M' : break;
    case 'n' : break;
    case 'N' : break;
    case 'o' : break;
    case 'O' : break;
    case 'p' : break;
    case 'P' : break;
    case 'q' : sandwich( arg, "cite" ); break;
    case 'Q' : break;
    case 'r' : break;
    case 'R' : break;
    case 's' : break;
    case 'S' : sandwich( arg, "strike" ); break;
    case 't' : break;
    case 'T' : break;
    case 'u' : sandwich( arg, "u" ); break;
    case 'U' : break;
    case 'v' : break;
    case 'V' : break;
    case 'w' : break;
    case 'W' : break;
    case 'x' : break;
    case 'X' : break;
    case 'y' : break;
    case 'Y' : break;
    case 'z' : break;
    case 'Z' : break;
    case '\\': break;
    case '^' : sandwich( arg, "sup" ); break;
    case '_' : sandwich( arg, "sub" ); break;
    case '`' : break;
    case '|' : break;
    case '~' : break;
    default  : break;
    }
  }
}

static char * to_mTEXT( char * section ){
  char keyStyle, *omega = mTEXT;
  strcpy( omega, section );

  while( regexp3( omega, "?<@{1,2}>[^@\x01-\x20\\&\\{\\(\\<\\[]+<[\\{\\(\\<\\[]><.>" ) ){
    if( lenCatch( 1 ) == 1 ){
      keyStyle = *(gpsCatch( 2 ));
      strCpy( gpsCatch( 1 ), gpsCatch( 3 ) );

      switch( keyStyle ){
      case '[': keyStyle = ']'; break;
      case '(': keyStyle = ')'; break;
      case '{': keyStyle = '}'; break;
      case '<': keyStyle = '>'; break;
      }

      char *t = strChr( gpsCatch( 1 ), keyStyle );
      strCpy( t, t + 1 );
      omega = t;
    } else omega += 2;
  }

  return mTEXT;
}


static char * to_mLINK( char * section ){
  to_mTEXT( section );
  regexp3( mTEXT, "<\\s>" );
  return replaceCatch( mLINK, "-", 1 );
}

static char * translateMarkup( char * section ){
  char *origin = section, *workPoint;
  static char keyStyle;

  char *end = mHTML;
  strCpy( mHTML, section );

  while( regexp3( section, "?<@{1,2}><[^@\x01-\x20\\&\\{\\(\\<\\[]+><[\\{\\(\\<\\[]>" ) ){
    if( lenCatch( 1 ) == 1 ){
      cpyCatch( mCOMM, 2 );

      switch( *(gpsCatch( 3 )) ){
      case '[': keyStyle = ']'; break;
      case '(': keyStyle = ')'; break;
      case '{': keyStyle = '}'; break;
      case '<': keyStyle = '>'; break;
      default: break;
      }

      workPoint = strchr( gpsCatch( 3 ) + 1, keyStyle );
      strnCpy( mTMP, gpsCatch( 3 ) + 1, workPoint - (gpsCatch( 3 ) + 1) );

      markupCommand( mCOMM, mTMP );
      workPoint++;
      strnCpy( end, section, gpsCatch( 1 ) - section );
      strcat( end, mTMP );
      end += strlen( end );

      section = workPoint;
    } else section = gpsCatch( 3 ) + 1;
  }

  strCpy( mCOMM, mHTML );
  return mCOMM;
}
