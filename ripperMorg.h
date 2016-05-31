#include <stdio.h>

enum morg_SECTYPE {
  morg_ZERO, morg_EMPTY, morg_COMMENT,
  morg_TITLE, morg_SUBTITLE, morg_AUTHOR, morg_MAIL,
  morg_LANGUAGE, morg_CHARSET,
  morg_CSS, morg_KEYWORDS, morg_DOCDESCRIPTION,
  morg_HEADLINE, morg_SUBHEADLINE, morg_TEXT,
  morg_BLOCK, morg_END_BLOCK, morg_TABLE,
  morg_ULIST, morg_OLIST, morg_DIALOG, morg_DESCRIPTION, morg_MEDIA,
  morg_UNICORN
};

enum morg_SECTYPE ripperMorg( FILE *morgFile, char *section, int *deep );
