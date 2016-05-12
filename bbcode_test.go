package bbcode_test

import (
	"github.com/sirclo-solution/bbcode"
	"strings"
	"testing"
)

func TestConvergence(t *testing.T) {
	s := `
	<strong>bolded text</strong>
	<u>underlined text</u>
	<br/><br/>
	<s>strikethrough text</s>
	<del>strikethrough text</del>
	<p>
		Links:
	</p>
	<ul>
		<li>
			<a href="foo">Foo</a>
		<li>
			<a href="/bar/baz">BarBaz</a>
	</ul>
	<table><tr><td>table 1</td><td>table 2</td></tr><tr><td>table 3</td><td>table 4</td></tr></table>
	<img src="Face-smile.svg" alt=":-)">
	<img src="http://upload.wikimedia.org/wikipedia/commons/thumb/7/7c/Go-home.svg/100px-Go-home.svg.png" alt="" />
	<a href="http://example.com">Example</a>
	<a href="http://example.org">http://example.org</a>
	<pre>monospaced text</pre>
	<blockquote><p>quoted text</p></blockquote>
	`

	firstBB, err := bbcode.HTMLToBBCode(strings.NewReader(s))
	if err != nil {
		t.Errorf("HTMLToBBCode error %v", err)
	}

	secondHTML, _ := bbcode.BBCodeToHTML(firstBB)

	thirdBB, err := bbcode.HTMLToBBCode(strings.NewReader(secondHTML))
	if err != nil {
		t.Errorf("HTMLToBBCode error %v", err)
	}

	fourthHTML, _ := bbcode.BBCodeToHTML(thirdBB)

	if firstBB != thirdBB {
		t.Errorf("BBCode changed after two times HTML conversion")
	}
	if secondHTML != fourthHTML {
		t.Errorf("HTML changed after two times BBCode conversion")
	}
}
