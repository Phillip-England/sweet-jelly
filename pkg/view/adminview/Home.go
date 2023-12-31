package adminview

import (
	"cfasuite/pkg/comp"
	"cfasuite/pkg/util"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	b := util.PageBuilder{
		Title: "CFA Suite - Home",
	}
	components := []string{
		comp.Header("Admin Home Page"),
		comp.AdminNav(),
	}
	b.AddComponents(components)
	w.Write(b.HtmlBytes())
}