package biskana

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestWhoIsThere( t *testing.T ) {
  rr := []string{ "COMMAND", "HEADLINE", "LIST", "DEFINITION", "ABOUT", "TEXT", "COMMENT", "EMPTY" }

  data := []struct {
    expected uint
    input string
  } {
    {    EMPTY, "" },
    {    EMPTY, "\n" },
    {    EMPTY, " \t\n" },
    {    EMPTY, "\n\n\n" },
    {    EMPTY, "               " },
    {    EMPTY, "              \n" },
    {    EMPTY, "              \n \n" },
    {    EMPTY, "              \n\t\n" },
    {    EMPTY, "              \n\t\n\r" },
    {  COMMENT, "@ " },
    {  COMMENT, "@\n" },
    {  COMMENT, "@ comment" },
    {  COMMENT, "@\tcomment" },
    {  COMMAND, "..command >" },
    {  COMMAND, "..command > green " },
    {  COMMAND, ".. command > greeeen" },
    {  COMMAND, "..  command  > green" },
    {  COMMAND, "..\tcommand  > green" },
    {  COMMAND, ".. stupit-command-command > green" },
    {  COMMAND, " ..command >" },
    {  COMMAND, "\t..command > green " },
    {  COMMAND, "\t\t.. command > greeeen" },
    {  COMMAND, "  ..\tcommand > greeeen" },
    {  COMMAND, " \t..  command  > green" },
    {  COMMAND, "    \t .. stupit-command-command > green" },
    { HEADLINE, "* " },
    { HEADLINE, "* headline" },
    { HEADLINE, "*\t " },
    { HEADLINE, "*\t headline" },
    { HEADLINE, "** " },
    { HEADLINE, "** headline" },
    { HEADLINE, "**\t " },
    { HEADLINE, "**\t headline" },
    { HEADLINE, "*** " },
    { HEADLINE, "*** headline" },
    { HEADLINE, "***\t " },
    { HEADLINE, "***\t headline" },
    {     LIST, "- ." },
    {     LIST, "- list" },
    {     LIST, " - e" },
    {     LIST, " - list" },
    {     LIST, " \t- a" },
    {     LIST, "+ e" },
    {     LIST, "+ list" },
    {     LIST, " + e" },
    {     LIST, " + list" },
    {     LIST, "2. e" },
    {     LIST, "3. list" },
    {     LIST, " 100. e" },
    {     LIST, " 7777. list" },
    {     LIST, " 79. u" },
    {     LIST, " \t8878. u\n" },
    {     LIST, "2) a" },
    {     LIST, "3) list" },
    {     LIST, " 100) u" },
    {     LIST, " 7777) list" },
    {     LIST, " 79) a" },
    {     LIST, " \t8878) a\n" },
    {     LIST, "A. u" },
    {     LIST, "B. e" },
    {     LIST, "C. list" },
    {     LIST, " ab. ue" },
    {     LIST, " abc. u" },
    {     LIST, " abcd. list" },
    {     LIST, " Aa. u\n" },
    {     LIST, " \tACEH. a\n" },
    {     LIST, "A) o" },
    {     LIST, "B) e" },
    {     LIST, "C) list" },
    {     LIST, " abc) o" },
    {     LIST, " zyxw) list" },
    {     LIST, " xa) o\n" },
    {     LIST, " \taoeu) e\n" },
    {    ABOUT, ":: a" },
    {    ABOUT, ":: about" },
    {    ABOUT, " :: abount" },
    {    ABOUT, "\t :: about" },
    {     TEXT, "::" },
    {     TEXT, "::" },
    {     TEXT, "1" },
    {     TEXT, "2a. " },
    {     TEXT, "3a. list" },
    {     TEXT, " 10a." },
    {     TEXT, " 100a. " },
    {     TEXT, " 7777a. list" },
    {     TEXT, " 79a.\n" },
    {     TEXT, " \t8878a. \n" },
    {     TEXT, "1a)" },
    {     TEXT, "2a) " },
    {     TEXT, "3a) list" },
    {     TEXT, " 10a)" },
    {     TEXT, " 100a) " },
    {     TEXT, " 7777a) list" },
    {     TEXT, " 79a)\n" },
    {     TEXT, " \t8878a) \n" },
    {     TEXT, "text" },
    {     TEXT, "\ntext" },
    {     TEXT, "@" },
    {     TEXT, "@text" },
    {     TEXT, " \n\t\n\r\a" },
    {     TEXT, ":< no setup " },
    {     TEXT, " :<setup> hahah" },
    {     TEXT, "text\ntext" },
    {     TEXT, " * " },
    {     TEXT, " *** headline" },
    {     TEXT, " ***\t " },
    {     TEXT, " ***\t headline" },
  }

  for _, c := range data {
    x := whoIsThere( c.input )

    if x != c.expected {
      t.Errorf( "TestWhoIsThere( \"%s\" )\nexpected %s\nresult   %s\n", c.input, rr[c.expected], rr[x] )
    }
  }

}

func TestWhatListIsThere( t *testing.T ) {
  rr := []string{ "LIST_ERR", "LIST_MINUS", "LIST_PLUS", "LIST_NUM", "LIST_ALPHA", "LIST_MDEF", "LIST_PDEF", "LIST_DIALOG" }

  data := []struct {
    expected uint
    input string
  } {
    { LIST_MINUS, "- ." },
    { LIST_MINUS, "- list" },
    { LIST_MINUS, " - e" },
    { LIST_MINUS, " - list" },
    { LIST_MINUS, " \t- a" },
    { LIST_PLUS , "+ e" },
    { LIST_PLUS , "+ list" },
    { LIST_PLUS , " + e" },
    { LIST_PLUS , " + list" },
    { LIST_NUM  , "2. e" },
    { LIST_NUM  , "3. list" },
    { LIST_NUM  , " 100. e" },
    { LIST_NUM  , " 7777. list" },
    { LIST_NUM  , " 79. u" },
    { LIST_NUM  , " \t8878. u\n" },
    { LIST_NUM  , "2) a" },
    { LIST_NUM  , "3) list" },
    { LIST_NUM  , " 100) u" },
    { LIST_NUM  , " 7777) list" },
    { LIST_NUM  , " 79) a" },
    { LIST_NUM  , " \t8878) a\n" },
    { LIST_ALPHA, "A. u" },
    { LIST_ALPHA, "B. e" },
    { LIST_ALPHA, "C. list" },
    { LIST_ALPHA, " ab. ue" },
    { LIST_ALPHA, " abc. u" },
    { LIST_ALPHA, " abcd. list" },
    { LIST_ALPHA, " Aa. u\n" },
    { LIST_ALPHA, " \tACEH. a\n" },
    { LIST_ALPHA, "A) o" },
    { LIST_ALPHA, "B) e" },
    { LIST_ALPHA, "C) list" },
    { LIST_ALPHA, " abc) o" },
    { LIST_ALPHA, " zyxw) list" },
    { LIST_ALPHA, " xa) o\n" },
    { LIST_ALPHA, " \taoeu) e\n" },
    { LIST_MDEF , "- . ::" },
    { LIST_MDEF , "- list :: def" },
    { LIST_MDEF , " - e :: Hex" },
    { LIST_MDEF , " - list\n   :: list" },
    { LIST_MDEF , " \t- a\n \t  :: def" },
    { LIST_PDEF , "+ . ::" },
    { LIST_PDEF , "+ list :: def" },
    { LIST_PDEF , " + e :: Hex" },
    { LIST_PDEF , " + list\n   :: list" },
    { LIST_PDEF , " \t+ a\n \t  :: def" },
    { LIST_ERR  , "" },
    { LIST_ERR  , "\n" },
    { LIST_ERR  , " \t\n" },
    { LIST_ERR  , "\n\n\n" },
    { LIST_ERR  , "  " },
    { LIST_ERR  , "hey" },
    { LIST_ERR  , "-" },
    { LIST_ERR  , "- " },
    { LIST_ERR  , "+" },
    { LIST_ERR  , "+ " },
    { LIST_ERR  , "..data >" },
    { LIST_ERR  , "1" },
    { LIST_ERR  , "2a. " },
    { LIST_ERR  , "3a. list" },
    { LIST_ERR  , " 10a." },
    { LIST_ERR  , " 100a. " },
    { LIST_ERR  , " 7777a. list" },
    { LIST_ERR  , " 79a.\n" },
    { LIST_ERR  , " \t8878a. \n" },
    { LIST_ERR  , "1a)" },
    { LIST_ERR  , "2a) " },
    { LIST_ERR  , "3a) list" },
    { LIST_ERR  , " 10a)" },
    { LIST_ERR  , " 100a) " },
    { LIST_ERR  , " 7777a) list" },
    { LIST_ERR  , " 79a)\n" },
    { LIST_ERR  , " \t8878a) \n" },
    { LIST_MINUS, "- . :" },
    { LIST_MINUS, "- list:: def" },
    { LIST_MINUS, " - e \n:: Hex" },
    { LIST_MDEF , " - list\n :: list" },
    { LIST_MDEF , " \t- a\n :: def" },
    { LIST_DIALOG, "> d" },
    { LIST_DIALOG, " > dialog" },
    { LIST_DIALOG, " \t> dialog" },
  }

  for _, c := range data {
    x := whatListIsThere( c.input )

    if x != c.expected {
      t.Errorf( "whatListIsThere( \"%s\" )\nexpected %s\nresult   %s\n", c.input, rr[c.expected], rr[x] )
    }
  }

}

func TestMakeHtml( t *testing.T ) {
	files := []string{
    "70_empty",
    "71_basic-html",
    "72_code-html",
    "73_code-html",
    "79_full-morg",
	}

  doTestsReference( t, files )
}

func doTestsReference( t *testing.T, files []string ) {
	for _, basename := range files {
		filename := filepath.Join("testdata", basename + ".morg")
		inputBytes, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Errorf("Couldn't open '%s', error: %v\n", filename, err)
			continue
		}
		input := string(inputBytes)

		filename = filepath.Join("testdata", basename+".html")
		expectedBytes, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Errorf("Couldn't open '%s', error: %v\n", filename, err)
			continue
		}
		expected := string(expectedBytes)

		// fmt.Fprintf(os.Stderr, "processing %s ...", filename)
		actual := MakeHtml( input, "" )
		if actual != expected {
			t.Errorf( "\nTest biskana: [%s] != [%s]", basename+".morg", basename+".html" )
      ioutil.WriteFile( filepath.Join("testdata", basename ) + ".html" + ".out", []byte(actual), 0666 )
		}
	}
}
