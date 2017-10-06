package katana

import (
  "testing"
)

func TestMarkup( t *testing.T ){
  data := []struct {
    in, out string
    errors int
  } {
    { "", "", 0 },
    { "texto", "texto", 0 },
    { "texto @e(emph) @b(bold)", "texto emph bold", 0 },
    { "@e(algo<>emph) @b(bold)", "emph bold", 0 },
    { "n", "n", 0 },
    { "b", "b", 0 },
    { "@@", "@", 0 },
    { "@{", "{", 0 },
    { "@}", "}", 0 },
    { "@((", "((", 0 },
    { "@e[", "", 1 },
    { "@e{", "", 1 },
    { "@e}", "@e}", 1 },
    { "@ [", "@ [", 1 },
    { "@e}h", "@e}h", 1 },
    { "@e>h", "@e>h", 1 },
    { "hey @e[", "hey ", 1 },
    { "aloha bye@e{", "aloha bye", 1 },
    { "@\t>h", "@\t>h", 1 },
    { "@b{h", "h", 1 },
    { "@e{h", "h", 1 },
    { "@\"{h", "h", 1 },
    { "@\"{cite @e(emph)}", "cite emph", 0 },
    { "@q<quote @e<emph @i[italic]>>", "quote emph italic", 0 },
    { "@b<something <@> something-more>", "something <> something-more", 0 },
    { "@b<something <>b>", "b", 0 },

    { "@be(hola @l<link<>@c[main()]>)", "hola main()", 0 },
    { "hela @aei {hey", "hela @aei {hey", 1 },
    { "hela @{aei @({hey", "hela {aei ({hey", 0 },
    { "hela @{aei @\n{hey", "hela {aei @\n{hey", 1 },
    { "hela @i} @x}@{aei @e{hey", "hela @i} @x}{aei hey", 3 },
    { "hela @i}@x}@{aei @e{hey", "hela @i}@x}{aei hey", 3 },
  }

  for _, d := range data {
    x := new(Scanner).NewSrc( d.in ).Init()
    eCount := 0
    x.CustomError = func( s *Scanner, msg string ){ eCount++ }
    m := x.GetMarkup()

    result := m.String()
    if result != d.out || d.errors != eCount {
      t.Errorf( "Markup.String( %q )\nresult   %q [%d errors]\nexpected %q [%d errors]", d.in, result, d.errors, d.out, eCount )
    }
  }
}

func TestRebuild( t *testing.T ){
  data := [...]string {
    "",
    "hola",
    "@b(hola)",
    "@b<hola>",
    "@b[hola]",
    "@b{hola}",
    "text @b(left<>right)",
    "text @b<right> b[left<>right]",
    "text @b<left<>right> b[left<>right]",
    "@b(left @i(it @e(ith<>oth)<>oth) ot <> et)",
    "@q[left @i(ith-left @i{ith-left2<>oth 2} ith left-1.2 <> ith right) left 1.2 <> right @c(right left @i{&¤}<>right)]",
    "@b(hola @e(emph @i(it) ot) ul)",

    "@@{",
    "@e",
    "@a(hola)",
    "@ab(hola)",
    "@abcd(hola)",
    "hola @abcd(hola)",
    "hola @abcd(hola) bye",
    "hola @abcd(hey) [byte 0<>byte 1<>byte 2]",

    "@be(hola @l<link<>@c[main()]>)",
  }

  for i, d := range data {
    x := new(Scanner).NewSrc( d ).Init()
    x.CustomError = quietSplash
    m := x.GetMarkup()

    result := m.Rebuild()
    if result != d {
      t.Errorf( "[%d] Markup.Rebuild( %q )\nresult   %q\nexpected %q\n%v", i, d, result, d, m )
    }
  }
}

func TestMultiMarkupRebuild( t *testing.T ){
  data := [...]struct {
    in, out string
  } {
    { "@\"(@c(func))", "@\"c(func)" },
    { "hola @\"(@c(func))", "hola @\"c(func)" },
    { "hola @\"(@e(@c(func)))", "hola @\"ec(func)" },
    { "hola @\"(@e(@c(func <>toMark)))", "hola @\"ec(func <>toMark)" },
    { "hola @i(@e(hey<>@c(func <>toMark)))", "hola @ie(hey<>@c(func <>toMark))" },
    { "hola @i(@q[@e(hey<>@c[func <>toMark()])])", "hola @iqe(hey<>@c[func <>toMark()])" },
  }

  for i, d := range data {
    x := new(Scanner).NewSrc( d.in ).Init()
    x.CustomError = quietSplash
    m := x.GetMarkup()

    result := m.Rebuild()
    if result != d.out {
      t.Errorf( "[%d] Markup.Rebuild( %q )\nresult   %q\nexpected %q", i, d.in, result, d.out )
    }
  }
}
