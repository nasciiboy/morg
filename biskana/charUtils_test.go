package biskana

import (
	"testing"
)

func TestDragTextByIndent( t *testing.T ) {
	data := []struct {
    input, expected string
    indent int
  } {
    { "- 1st element", "- 1st element", 0 },
    { " - 1st element", " - 1st element", 0 },
    { " - 1st element", " - 1st element", 1 },
    { " - 1st element", "", 2 },
    { " - 1st element\n - 2st element", "", 2 },
    { " - 1st element\n - 2st element", " - 1st element\n - 2st element", 1 },
    { " - 1st element\n- 2st element", " - 1st element\n", 1 },
    { "- 1st element\n - 2st element", "", 1 },
    { "- 1st element\n- 2st element", "", 1 },
    { "  3st element\n  element", "  3st element\n  element", 2 },
	}

  for _, c := range data {
    x, _ := dragTextByIndent( c.input, c. indent )

    if x != c.expected {
      t.Errorf( "TestDragTextByIndent( \"%s\" )\nexpected [%s]\nresult   [%s]\n", c.input, c.expected, x )
    }
  }
}
