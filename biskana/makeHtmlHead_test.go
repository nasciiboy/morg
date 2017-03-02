package biskana

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestMakeHtmlHead( t *testing.T ) {
	files := []string{
    "00_empty",
    "30_simple-head",
	}

  doHeadTestsReference( t, files )
}

func doHeadTestsReference( t *testing.T, files []string ) {
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
		actual, _ := MakeHtmlHead( input )
		if actual != expected {
			t.Errorf( "\nTest makeHtmlHead: [%s] != [%s]", basename+".morg", basename+".html" )
      ioutil.WriteFile( filepath.Join("testdata", basename ) + ".html" + ".out", []byte(actual), 0666 )
		}
	}
}
