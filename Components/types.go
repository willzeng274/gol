package Components

type Component struct {
	Type      string
	InnerText string
	Children  []Component
	Class     []string
	Id        string
	Style     string
	OnClick   JSFunction
}

var Types = struct {
	Div       string
	Paragraph string
}{
	"div",
	"p",
}

type JSFunction struct {
	Parameters []string
	Returns    string
	Name       string
	Body       string
}

func DummyFunction(fn JSFunction) bool {
	if len(fn.Parameters) == 0 && fn.Returns == "" && fn.Name == "" && fn.Body == "" {
		return true
	} else {
		return false
	}
}
