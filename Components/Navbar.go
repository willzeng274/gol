package Components

var Navbar Component = Component{
	Type:      Types.Div,
	InnerText: "Hello World!",
	Children:  make([]Component, 0),
	Class:     make([]string, 0),
	Id:        "",
	Style:     "",
	OnClick: JSFunction{
		Parameters: []string{},
		Returns:    "",
		Name:       "call",
		Body:       "alert('hi')",
	},
}

//  func() string {
// 		return "(function call(){alert(\"hi\")})()"
// 	}
