package appview

import (
	"cfasuite/pkg/comp"
	"cfasuite/pkg/model/usermod"
	"cfasuite/pkg/mw"
	"cfasuite/pkg/util"
	"net/http"
)

// generates html for update profile page
func Bio(w http.ResponseWriter, r *http.Request) {
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
	bioFormErr := r.URL.Query().Get("BioFormErr")
    b := util.PageBuilder{
        Title: "CFA Suite - User Details",
    }
    components := []string{
		comp.Header("App Bio Page"),
		comp.TeamNav(),
		comp.UserDetails(user),
		comp.UserBioForm(user, r, bioFormErr),
    }
    b.AddComponents(components)
    w.Write(b.HtmlBytes())
}