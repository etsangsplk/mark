package mark

import (
	"github.com/a8m/expect"
	"testing"
)

func TestRender(t *testing.T) {
	expect := expect.New(t)
	expect(Render("foobar")).To.Equal("<p>foobar</p>")
	expect(Render("foo  \nbar")).To.Equal("<p>foo<br>bar</p>")
	expect(Render("__bar__ foo")).To.Equal("<p><strong>bar</strong> foo</p>")
	expect(Render("**bar** foo __bar__")).To.Equal("<p><strong>bar</strong> foo <strong>bar</strong></p>")
	expect(Render("**bar**__baz__")).To.Equal("<p><strong>bar</strong><strong>baz</strong></p>")
	expect(Render("**bar**foo__bar__")).To.Equal("<p><strong>bar</strong>foo<strong>bar</strong></p>")
	expect(Render("_bar_baz")).To.Equal("<p><em>bar</em>baz</p>")
	expect(Render("_foo_~~bar~~ baz")).To.Equal("<p><em>foo</em><del>bar</del> baz</p>")
	expect(Render("~~baz~~ _baz_")).To.Equal("<p><del>baz</del> <em>baz</em></p>")
	expect(Render("1  \n2  \n3")).To.Equal("<p>1<br>2<br>3</p>")
	expect(Render("1\n\n2")).To.Equal("<p>1</p>\n<p>2</p>")
	expect(Render("1\n\n\n2")).To.Equal("<p>1</p>\n\n<p>2</p>")
}
