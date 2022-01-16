package parsers

import (
	"fmt"
	"strings"
	C "test/Components"
)

func Parse(s *string, comp C.Component) {
	st := *s
	i := strings.Index(*s, "<script>") + len("<script>")
	c := Factory(comp)
	st = st[:i] + fmt.Sprintf("\ndocument.getElementById(\"root\").innerHTML = `%s`;document.body.removeChild(document.querySelector('script'))", strings.Replace(c, "`", "\\`", -1)) + st[i:] + "\n"
	*s = st
}

func Factory(c C.Component) string {
	base := fmt.Sprintf("<%s ", c.Type)
	if len(c.Class) != 0 {
		base += "class=\""
		for _, s := range c.Class {
			base += s + " "
		}
		base += "\" "
	}
	if c.Id != "" {
		base += fmt.Sprintf("id = \"%s\" ", c.Id)
	}
	if c.Style != "" {
		base += fmt.Sprintf("style = \"%s\" ", c.Style)
	}
	if !C.DummyFunction(c.OnClick) {
		base += fmt.Sprintf("onclick = \"%s\" ", JSFunctionFactory(c.OnClick))
	}
	base += ">" + c.InnerText
	if len(c.Children) != 0 {
		for _, co := range c.Children {
			base += Factory(co)
		}
	}
	base += fmt.Sprintf("</%s>", c.Type)
	return base
}

func JSFunctionFactory(fn C.JSFunction) string {
	if fn.Name == "" {
		fn.Name = "_0x83758910923"
	}
	final := "(function " + fn.Name + "("
	if len(fn.Parameters) != 0 {
		for _, s := range fn.Parameters {
			final += s + ","
		}
	} else {
		final += "){"
	}
	if fn.Body != "" {
		final += strings.Replace(fn.Body, "\"", "'", -1) + ";"
	} else {
		return "(function _0x82081098409(){})()"
	}
	if fn.Returns != "" {
		final += "return " + fn.Returns + ";"
	}
	return final + "})()"
}
