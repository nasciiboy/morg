package html

import (
  "testing"

  "github.com/nasciiboy/morg/katana"
)

func TestToSafeHtml( t *testing.T ){
  data := [...]struct {
    in, out string
  } {
    { "", "" },
    { "text", "text" },
    { "main()\n", "main()\n" },
    { "2 < 5", "2 &lt; 5" },
    { "5 > 2", "5 &gt; 2" },
    { "m&m", "m&amp;m" },
    { "comma's", "comma&#39;s" },
    { "\"esterno\"", "&#34;esterno&#34;" },
    { `<'&">`, "&lt;&#39;&amp;&#34;&gt;" },
    { `<Aliquam erat volutpat.  Nunc 'eleifend&">`, "&lt;Aliquam erat volutpat.  Nunc &#39;eleifend&amp;&#34;&gt;" },
  }

  for _, d := range data {
    out := ToSafeHtml( d.in )
    if out != d.out {
      t.Errorf( "TestToSafeHtml( %q ) == %q, expected %q", d.in, out, d.out )
    }
  }
}

func TestFontify( t *testing.T ){
  data := [...]struct {
    in, out string
  } {
    { "", "" },
    { "text", "text" },
    { "main()\n", "main()\n" },
    { "2 < 5", "2 &lt; 5" },
    { "5 > 2", "5 &gt; 2" },
    { "m&m", "m&amp;m" },
    { "comma's", "comma&#39;s" },
    { "\"esterno\"", "&#34;esterno&#34;" },
    { `<'&">`, "&lt;&#39;&amp;&#34;&gt;" },
    { `<Aliquam erat volutpat.  Nunc 'eleifend&">`, "&lt;Aliquam erat volutpat.  Nunc &#39;eleifend&amp;&#34;&gt;" },

    { "simple @e(emph)", "simple <em>emph</em>" },
    { "simple @i(italic)", "simple <i>italic</i>" },
    { "simple @e(emph & emph-@i(italic))", "simple <em>emph &amp; emph-<i>italic</i></em>" },
    { "simple @ei(emph-italic)", "simple <em><i>emph-italic</i></em>" },
    { "simple @eil(lol<>emph-italic)", "simple <em><i><a href=\"lol\" >emph-italic</a></i></em>" },
    { "simple @lei(emph-italic)", "simple <a href=\"emph-italic\" ><em><i>emph-italic</i></em></a>" },
    { "simple @l(lol<>@ei(emph-italic))", "simple <a href=\"lol\" ><em><i>emph-italic</i></em></a>" },
    { "@&(omega) @&(real) simple @&(leftarrow) @&(inf)", "ω ℜ simple ← ∞" },
    { "morg @%q(multiple org)", "morg (<q>multiple org</q>)" },
    { `code @%"c{main()}`, "code (<q><code>main()</code></q>)" },
    { `@n(1545<>FootNote 1545)`, `<span class="note" ><sup><a href="#1545" >FootNote 1545</a></sup></span>` },
  }

  for _, d := range data {
    m := new(katana.Scanner).NewSrc(d.in).QuietSplash().Init().GetMarkup()
    out := Fontify( m )
    if out != d.out {
      t.Errorf( "TestFontify( %q )\nresult   %q\nexpected %q\n%v", d.in, out, d.out, m )
    }
  }
}
