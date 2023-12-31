package appview

import (
	"cfasuite/pkg/comp"
	"cfasuite/pkg/model/usermod"
	"cfasuite/pkg/mw"
	"cfasuite/pkg/util"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    // db, ok := r.Context().Value(dbKey).(*sql.DB)
    // if !ok {
    //     http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
    //     return
    // }
	user, ok := r.Context().Value(mw.AuthKey).(*usermod.Model)
	if !ok {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}
    b := util.PageBuilder{
        Title: "CFA Suite - User Details",
    }
    components := []string{
		comp.Header("App Home Page"),
		comp.TeamNav(),
		comp.UserDetails(user),
		comp.UserPhoto(user),
    }
    b.AddComponents(components)
    w.Write(b.HtmlBytes())
}