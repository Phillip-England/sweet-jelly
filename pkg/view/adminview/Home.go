package adminview

import (
	"cfasuite/pkg/util"
	"net/http"
)

type HomeData struct {
	Title string
}

func Home(w http.ResponseWriter, r *http.Request) {
	util.RenderTemplate(w, "./pkg/view/adminview/Home.html", HomeData{
		Title: "CFA Suite - Admin Home",
	})


	// b := util.PageBuilder{
	// 	Title: "CFA Suite - Home",
	// }
	// components := []string{
	// 	comp.Header("Admin Home Page"),
	// 	comp.AdminNav(),
	// }
	// b.AddComponents(components)
	// w.Write(b.HtmlBytes())
}