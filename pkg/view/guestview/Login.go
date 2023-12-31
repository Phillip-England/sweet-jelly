package guestview

import (
	"cfasuite/pkg/comp"
	"cfasuite/pkg/util"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// handling all 404s
	if r.URL.Path != "/" {
		b := util.PageBuilder{
			Title: "CFA Suite - 404 Not Found",
		}
		components := []string{
			"<h2>404 Not Found</h2>",
		}
		b.AddComponents(components)
		w.Write(b.HtmlBytes())
		return
	}
	// handling login page / index page
	b := util.PageBuilder{
		Title: "CFA Suite - Login",
	}
	components := []string{
		"<h1>Login Page</h1>",
		comp.LoginForm(r.URL.Query().Get("LoginFormErr")),
	}
	b.AddComponents(components)
	w.Write(b.HtmlBytes())
}