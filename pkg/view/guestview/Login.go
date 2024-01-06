package guestview

import (
	"cfasuite/pkg/util"
	"net/http"
)

type NotFoundData struct {
	Title string
}

type LoginData struct {
	Title string
	LoginFormErr string
	EmailInputValue string
	PasswordInputValue string
}

func Login(w http.ResponseWriter, r *http.Request) {
	// handling all 404s
	if r.URL.Path != "/" {
		util.RenderTemplate(w, "./pkg/view/guestview/404.html", NotFoundData{
			Title: "404 Not Found",
		})
		return
	}
	util.RenderTemplate(w, "./pkg/view/guestview/Login.html", LoginData{
		Title: "CFA Suite - Login",
		LoginFormErr: r.URL.Query().Get("LoginFormErr"),
		EmailInputValue: util.IfDevModeThen("test@gmail.com"),
		PasswordInputValue: util.IfDevModeThen("aspoaspo"),
	})
}