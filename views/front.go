package views

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func FrontPage() g.Node {
	return Page(
		"Canvas",
		"/",
		H1(g.Text(`Solutions to problems.`)),
		P(g.Text(`Do you have problems? We also had problems.`)),
		P(g.Raw(`Then we created the <em>canvas</em> app, and now we don't! 😬`)),
		H2(g.Text(`Do you want to know more?`)),
		P(g.Text(`Sign up to our newsletter below.`)),
		FormEl(Action("/newsletter/signup"), Method("post"),
			Input(Type("email"), Name("email"), AutoComplete("email"), Required()),
			Button(Type("submit"), g.Text("Sign up")),
		),
	)
}
