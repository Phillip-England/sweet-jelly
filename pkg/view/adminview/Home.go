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
}