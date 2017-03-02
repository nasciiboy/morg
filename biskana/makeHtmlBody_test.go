package biskana

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestMakeHtmlBody( t *testing.T ) {
	files := []string{
    "00_empty",
    "01_simple-text",
    "02_text-plus-emph",
    "03_simple-headlines",
    "04_headlines-and-spaces",
    "05_headlines-and-text",
    "06_headlines-and-text-emph",
    "07_basic-links",
    "10_img-commad",
    "11_src-commad",
    "12_figure-commad",
    "13_emphs-command",
    "20_basic-list",
    "21_inner-list",
    "22_mixed-list",
    "23_def-list",
    "25_basic-note",
	}

  doBodyTestsReference( t, files )
}

func doBodyTestsReference( t *testing.T, files []string ) {
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
		actual, _ := MakeHtmlBody( input )
		if actual != expected {
			t.Errorf( "\nTest makeHtmlBody: [%s] != [%s]\n   Gen [%s]",
        basename+".morg", basename+".html", basename+".morg"+".out" )
      ioutil.WriteFile( filepath.Join("testdata", basename ) + ".html" + ".out", []byte(actual), 0666 )
		}
	}
}
