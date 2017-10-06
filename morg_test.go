package main

import (
	"io/ioutil"
	"path/filepath"
	"testing"

  "github.com/nasciiboy/morg/biskana"
  "github.com/nasciiboy/morg/katana"
  "github.com/nasciiboy/morg/porg"
)

const errorSuffix = ".out"

func TestBiskana( t *testing.T ) {
  files := []string{
    "00_empty",
    "01_conf_basic",
    "02_conf_markupTitle",
    "03_conf_markup_options",
    "04_conf_markup_toc",
    "05_conf_markup_TOC!",
    "06_conf_markup_textoc",
    "07_conf_markup_fulltextoc",
    "10_block-figure",
    "11_block_code",
    "12_block_fancyCode",
    "13_block_fancyCodeInline",
    "14_block_figure&Code&Tok&Text",
    "15_block_media",
    "16_block_wrap",
    "17_block_cols",
    "18_block_pret",
    "19_block_brick",
    "20_block_quote",
    "21_block_srci",
    "22_block_srci-fancy",
    "23_block_srci-fancy-style",
    "30_about",
    "34_list_basic",
    "35_list_inner",
    "36_list_mixed",
    "37_list_def",
    "38_list_complex",
    "44_table_basic",
    "45_table_multiline-cels",
    "70_all_basic",
  }

  testFiles( t, files, ".morg", ".html" )
}

func testFiles( t *testing.T, files []string, inSuffix, outSuffix string ) {
	for _, basename := range files {
		filename := filepath.Join("testdata", basename + inSuffix )
		inputBytes, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Errorf("Couldn't open '%s', error: %v\n", filename, err)
			continue
		}

		filename = filepath.Join( "testdata", basename + outSuffix )
		expectedBytes, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Errorf("Couldn't open '%s', error: %v\n", filename, err)
      expectedBytes = []byte{}
		}
		expected := string(expectedBytes)

    doc, errs := katana.Parse( basename, string(inputBytes) )
    if errs != "" { t.Errorf( "[%s] morg:%s", basename, errs ) }

    var to uint
    switch outSuffix {
      case ".html": to = biskana.HTML
    }

		actual := biskana.Export( doc, to )
		if actual != expected {
			t.Errorf( "\nbiskanaDiff: [%s] != [%s]\n   Gen [%s]",
        basename + inSuffix, basename + outSuffix, basename + outSuffix + errorSuffix )
      ioutil.WriteFile( filepath.Join( "testdata", basename ) + outSuffix + errorSuffix, []byte(actual), 0666 )
		}
	}
}

func TestUnPorg( t *testing.T ) {
  files := []string{
    "00_empty",
    "90_unporg_basic",
  }

  doUnPorgTestsReference( t, files )
}

func doUnPorgTestsReference( t *testing.T, files []string ) {
	for _, basename := range files {
		filename := filepath.Join("testdata", basename + porgSuffix )
		inputBytes, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Errorf("Couldn't open '%s', error: %v\n", filename, err)
			continue
		}
		input := string(inputBytes)

		filename = filepath.Join( "testdata", basename + morgSuffix )
		expectedBytes, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Errorf("Couldn't open '%s', error: %v\n", filename, err)
      expectedBytes = []byte{}
		}
		expected := string(expectedBytes)

		actual := porg.UnPorg( string(input) )
		if actual != expected {
			t.Errorf( "\nTestUnPorg: [%s] != [%s]\n   Gen [%s]",
        basename + porgSuffix, basename + morgSuffix, basename + porgSuffix + errorSuffix )
      ioutil.WriteFile( filepath.Join( "testdata", basename ) + morgSuffix + errorSuffix, []byte(actual), 0666 )
		}
	}
}
