package biskana

import "testing"

func TestToHtml(t *testing.T) {
  marckupTriggerTest( t )
  swapTest( t )
  marckupTest( t )
  customTest( t )
  humanTest( t )
  complexTest( t )
}

func marckupTriggerTest( t *testing.T ){
  data := []struct {
    input, hOutput, tOutput string
    n int
  }{
    { "", "", "", 0 },
    { "n", "n", "n", 1 },
    { "b", "b", "b", 1 },
    { "@@", "@", "@", 2 },
    { "@{", "{", "{", 2 },
    { "@}", "}", "}", 2 },
    { "@((", "(", "(", 2 },
    { "@e[", "@e[", "@e[", 3 },
    { "@e{", "@e{", "@e{", 3 },
    { "@e}", "@e}", "@e}", 3 },
    { "@ [", "@ [", "@ [", 3 },
    { "@e}h", "@e}", "@e}", 3 },
    { "@e>h", "@e>", "@e>", 3 },
    { "@\t>h", "@\t>", "@\t>", 3 },
    { "@b{h", "<b>h</b>", "h", 4 },
    { "@e{h", "<em>h</em>", "h", 4 },
    { "@\"{h", "<q>h</q>", "h", 4 },
    { "@\"{cite @e(emph)}", "<q>cite <em>emph</em></q>", "cite emph", 17 },
    { "@q<quote @e<emph @i[italic]>>", "<q>quote <em>emph <i>italic</i></em></q>", "quote emph italic", 29 },
    { "@b<something <@> something-more>", "<b>something &lt;&gt; something-more</b>", "something &lt;&gt; something-more", 32 },
    { "@b<something <>b>", "<b>b</b>", "b", 17 },
    { "@l<something>", "<a href=\"something\" >something</a>", "something", 13 },
    { "@l<something \n more>", "<a href=\"something-more\" >something \n more</a>", "something \n more", 20 },
    { "@l<something<>b>", "<a href=\"something\" >b</a>", "b", 16 },
    { "@l<something <& <>b>", "<a href=\"something-&lt;&amp;\" >b</a>", "b", 20 },
    { "@l<something <& <>@b(b)>", "<a href=\"something-&lt;&amp;\" ><b>b</b></a>", "b", 24 },
    { "@l<something <& <>@b(bold) @e[emph]>", "<a href=\"something-&lt;&amp;\" ><b>bold</b> <em>emph</em></a>", "bold emph", 36 },
    { "@l<something <& <>@b(bold @e[emph])>", "<a href=\"something-&lt;&amp;\" ><b>bold <em>emph</em></b></a>", "bold emph", 36 },
    { "@l<something \n more<>b>", "<a href=\"something-more\" >b</a>", "b", 23 },
    { "@l<\t  something \n more \t <>b>", "<a href=\"something-more\" >b</a>", "b", 29 },
    { "@l<@b(bold<>@e[emph])>", "<a href=\"emph\" ><b><em>emph</em></b></a>", "emph", 22 },
    { "@l<something @b(bold<>@e[emph])>", "<a href=\"something-emph\" >something <b><em>emph</em></b></a>", "something emph", 32 },
    { "@l<something @l(bold<>@e[emph])>", "<a href=\"something-emph\" >something <a href=\"bold\" ><em>emph</em></a></a>", "something emph", 32 },
    { "@l<@b(bold)<>text>", "<a href=\"bold\" >text</a>", "text", 18 },
    { "@l<@b(fis<>baz)<>text>", "<a href=\"baz\" >text</a>", "text", 22 },
  }

  for _, c := range data {
    hOutput, tOutput, n := marckupTrigger( c.input )
    if hOutput != c.hOutput || tOutput != c.tOutput || n != c.n {
      t.Errorf( "Error in marckupTrigger\n   input ==> %s\nexpected ==> [%s][%d]\n  result ==> [%s][%d]\n    text ==> [%s]\n   rText ==> [%s]\n", c.input, c.hOutput, c.n, hOutput, n, c.tOutput, tOutput )
    }
  }
}

func swapTest( t *testing.T ){
  data := []struct {
    input, output string
  }{
    { "@@", "@" },
    { "nasciiboy@@test", "nasciiboy@test" },
    { "test", "test" },
    { "<test>", "&lt;test&gt;" },
    { "test & test", "test &amp; test" },
    { "\"test\"", "&quot;test&quot;" },
  }

  for _, c := range data {
    output := ToHtml( c.input )
    if output != c.output {
      t.Errorf( "Error in htmlize_test.go\n   input ==> %s\nexpected ==> [%s]\n  result ==> [%s]\n", c.input, c.output, output )
    }
  }
}

func marckupTest( t *testing.T ){
  data := []struct {
    input, output string
  }{
    // { "@!(test)", "test" },
    { "@\"(test)", "<q>test</q>" },
    // { "@#(test)", "pathtest" },
    { "@$(test)", "<code class=\"command\" >test</code>" },
    { "@$[test]", "<code class=\"command\" >test</code>" },
    { "@${test}", "<code class=\"command\" >test</code>" },
    { "@$<test>", "<code class=\"command\" >test</code>" },
    // { "@%(test)", "parentesistest" },
    // { "@&(test)", "simboltest" },
    { "@'(test)", "<samp>test</samp>" },
    { "@'[test]", "<samp>test</samp>" },
    { "@'{test}", "<samp>test</samp>" },
    { "@'<test>", "<samp>test</samp>" },
    // { "@*(test)", "test" },
    // { "@+(test)", "test" },
    // { "@,(test)", "test" },
    { "@-(test)", "––test––" },
    { "@-[test]", "––test––" },
    { "@-{test}", "––test––" },
    { "@-<test>", "––test––" },
    // { "@.(test)", "test" },
    // { "@/(test)", "test" },
    // { "@0(test)", "test" },
    // { "@1(test)", "test" },
    // { "@2(test)", "test" },
    // { "@3(test)", "test" },
    // { "@4(test)", "test" },
    // { "@5(test)", "test" },
    // { "@6(test)", "test" },
    // { "@7(test)", "test" },
    // { "@8(test)", "test" },
    // { "@9(test)", "test" },
    { "@:(test)", "<dfn>test</dfn>" },
    // { "@;(test)", "test" },
    // { "@=(test)", "test" },
    // { "@?(test)", "test" },
    // { "@A(test)", "test" },
    // { "@B(test)", "test" },
    // { "@C(test)", "smallCapstest" },
    // { "@D(test)", "test" },
    // { "@E(test)", "errortest" },
    // { "@F(test)", "Functest" },
    // { "@G(test)", "test" },
    // { "@H(test)", "test" },
    // { "@I(test)", "test" },
    // { "@J(test)", "test" },
    // { "@K(test)", "keywordtest" },
    // { "@L(test)", "marckuptest" },
    // { "@M(test)", "test" },
    // { "@N(test)", "defNotetest" },
    // { "@O(test)", "test" },
    // { "@P(test)", "test" },
    // { "@Q(test)", "test" },
    // { "@R(test)", "resulttest" },
    // { "@S(test)", "test" },
    // { "@T(test)", "radiotargettest" },
    // { "@U(test)", "test" },
    // { "@V(test)", "vartest" },
    // { "@W(test)", "warningtest" },
    // { "@X(test)", "test" },
    // { "@Y(test)", "test" },
    // { "@Z(test)", "test" },
    // { "@\(test)", "test" },
    { "@^(test)", "<sup>test</sup>" },
    { "@_(test)", "<sub>test</sub>" },
    // { "@`(test)", "test" },
    { "@a(test)", "<abbr>test</abbr>" },
    { "@a[test]", "<abbr>test</abbr>" },
    { "@a{test}", "<abbr>test</abbr>" },
    { "@a<test>", "<abbr>test</abbr>" },
    { "@b(test)", "<b>test</b>" },
    { "@c(test)", "<code>test</code>" },
    // { "@d(test)", "<data>test</data>" },
    { "@e(test)", "<em>test</em>" },
    { "@f(test)", "<span class=\"file\" >test</span>" },
    { "@f[test]", "<span class=\"file\" >test</span>" },
    { "@f{test}", "<span class=\"file\" >test</span>" },
    { "@f<test>", "<span class=\"file\" >test</span>" },
    // { "@g(test)", "test" },
    // { "@h(test)", "test" },
    { "@i(test)", "<i>test</i>" },
    // { "@j(test)", "test" },
    { "@k(test)", "<kbd>test</kbd>" },
    { "@l(test)", "<a href=\"test\" >test</a>" },
    { "@l(test to to)", "<a href=\"test-to-to\" >test to to</a>" },
    { "@l( test )", "<a href=\"test\" > test </a>" },
    { "@l(  test to to  )", "<a href=\"test-to-to\" >  test to to  </a>" },
    { "@m(test)", "<span class=\"math\" >test</span>" },
    // { "@n(test)", "notetest" },
    // { "@o(test)", "test" },
    // { "@p(test)", "test" },
    { "@q(test)", "<q>test</q>" },
    // { "@r(test)", "reftest" },
    // { "@s(test)", "test" },
    // { "@t(test)", "targettest" },
    { "@u(test)", "<u>test</u>" },
    { "@v(test)", "<code class=\"verbatim\" >test</code>" },
    // { "@w(test)", "test" },
    // { "@x(test)", "test" },
    // { "@y(test)", "test" },
    // { "@z(test)", "test" },
    // { "@|(test)", "test" },
    // { "@~(test)", "test" },
    // { "@@(test)", "test" },
    // { "@((test)", "test" },
    // { "@)(test)", "test" },
    // { "@[(test)", "test" },
    // { "@](test)", "test" },
    // { "@<(test)", "test" },
    // { "@>(test)", "test" },
    // { "@{(test)", "test" },
    // { "@}(test)", "test" },
  }

  for _, c := range data {
    output := ToHtml( c.input )
    if output != c.output {
      t.Errorf( "Error in htmlize_test.go\n   input ==> %s\nexpected ==> [%s]\n  result ==> [%s]\n", c.input, c.output, output )
    }
  }
}

func customTest( t *testing.T ){
  data := []struct {
    input, output string
  }{
    { "@l(<>test)", "<a href=\"test\" >test</a>" },
    { "@l(texnho<>test to to)", "<a href=\"texnho\" >test to to</a>" },
    { "@l(@b(bold))", "<a href=\"bold\" ><b>bold</b></a>" },
    { "@l(  test to to  )", "<a href=\"test-to-to\" >  test to to  </a>" },
  }

  for _, c := range data {
    output := ToHtml( c.input )
    if output != c.output {
      t.Errorf( "Error in htmlize_test.go\n   input ==> %s\nexpected ==> [%s]\n  result ==> [%s]\n", c.input, c.output, output )
    }
  }
}

func humanTest( t *testing.T ){
  data := []struct {
    input, output string
  }{
    { "", "" },
    { "123", "123" },
    { "@", "@" },
    { "@e", "@e" },
    { "123 @", "123 @" },
    { "123 @e", "123 @e" },
    { "123 @e(", "123 @e(" },
    { "123 @e(a", "123 <em>a</em>" },
    { "@e(test", "<em>test</em>" },
    { "@(test", "(test" },
    { "@[test", "[test" },
    { "@{test", "{test" },
    { "@<test", "&lt;test" },
    { "@>test", "&gt;test" },
    { "e(test", "e(test" },
    { "e[test", "e[test" },
    { "e{test", "e{test" },
    { "e<test", "e&lt;test" },
    { "@@(test", "@(test" },
    { "@@{test", "@{test" },
    { "@l<", "@l<" },
    { "@l<a", "<a href=\"a\" >a</a>" },
    { "@e<emph @(", "<em>emph (</em>" },
    { "@e<emph @e(", "<em>emph @e(</em>" },
    { "@e<emph @e(other emph", "<em>emph <em>other emph</em></em>" },
    { "@e<emph @b(bold @i[italic", "<em>emph <b>bold <i>italic</i></b></em>" },
    { "@e<emph @b(bold @i[italic)", "<em>emph <b>bold <i>italic)</i></b></em>" },
    { "@e<emph @b(bold) @i[italic", "<em>emph <b>bold</b> <i>italic</i></em>" },
    { "@e<emph> @b(bold @i[italic)", "<em>emph</em> <b>bold <i>italic)</i></b>" },
    { "@e<emph> @b(bold) @i[italic", "<em>emph</em> <b>bold</b> <i>italic</i>" },
    { "@e<emph> @b(bold) @i[italic]", "<em>emph</em> <b>bold</b> <i>italic</i>" },
  }

  for _, c := range data {
    output := ToHtml( c.input )
    if output != c.output {
      t.Errorf( "Error in htmlize_test.go\n   input ==> %s\nexpected ==> [%s]\n  result ==> [%s]\n", c.input, c.output, output )
    }
  }
}

func complexTest( t *testing.T ){
  data := []struct {
    input, output string
  }{
    { "@e(Who's Theme)", "<em>Who's Theme</em>" },
    { "@\"(@e(THE RAPTOR) @c(Code) @b(Book))", "<q><em>THE RAPTOR</em> <code>Code</code> <b>Book</b></q>" },
    { "@e(THE @q(RAPTOR)) @c(Code) @b(Book)", "<em>THE <q>RAPTOR</q></em> <code>Code</code> <b>Book</b>" },
    { "@e(@b(@i(T))HE RAPTOR) @c(Code) @b(Book)", "<em><b><i>T</i></b>HE RAPTOR</em> <code>Code</code> <b>Book</b>" },
    { "@e(The Good,) @b(The Bad &) @k(The Queen)", "<em>The Good,</em> <b>The Bad &amp;</b> <kbd>The Queen</kbd>" },
    { "@l(@b(bold))", "<a href=\"bold\" ><b>bold</b></a>" },
    { "@l(@b(bold) & @e(emph))", "<a href=\"bold-&amp;-emph\" ><b>bold</b> &amp; <em>emph</em></a>" },
    { "@l(@b(bold & @e(emph)))", "<a href=\"bold-&amp;-emph\" ><b>bold &amp; <em>emph</em></b></a>" },
    { "@l(@b(bold & @e(emph) & @i(italic)))", "<a href=\"bold-&amp;-emph-&amp;-italic\" ><b>bold &amp; <em>emph</em> &amp; <i>italic</i></b></a>" },
    { "@l(@b(bold & @e(emph & @i(italic))))", "<a href=\"bold-&amp;-emph-&amp;-italic\" ><b>bold &amp; <em>emph &amp; <i>italic</i></em></b></a>" },
    { "@l(@b(bold & @l(link & @i(italic))))", "<a href=\"bold-&amp;-link-&amp;-italic\" ><b>bold &amp; <a href=\"link-&amp;-italic\" >link &amp; <i>italic</i></a></b></a>" },
  }

  for _, c := range data {
    output := ToHtml( c.input )
    if output != c.output {
      t.Errorf( "Error in htmlize_test.go\n   input ==> %s\nexpected ==> [%s]\n  result ==> [%s]\n", c.input, c.output, output )
    }
  }
}
