package adminview

import (
	"cfasuite/pkg/comp"
	"cfasuite/pkg/database"
	"cfasuite/pkg/model/locationmod"
	"cfasuite/pkg/mw"
	"cfasuite/pkg/util"
	"database/sql"
	"fmt"
	"net/http"
)

func Locations(w http.ResponseWriter, r *http.Request) {
    db, ok := r.Context().Value(mw.DbKey).(*sql.DB)
    if !ok {
        http.Error(w, "failed to get database connection", http.StatusInternalServerError)
        return
    }
	locationFormErr := r.URL.Query().Get("BioFormErr")
    locationRepo := &locationmod.Repo{
        DB: db,
    }
    if err := database.DbGetAll(locationRepo); err != nil {
        http.Error(w, fmt.Sprintf("Failed to get locations: %s", err), http.StatusInternalServerError)
        return
    }
    b := util.PageBuilder{
        Title: "CFA Suite - User Details",
    }
    components := []string{
        comp.Header("Admin Locations Page"),
        comp.AdminNav(),
        comp.CreateLocationForm(locationFormErr),
        comp.LocationList(locationRepo.Locations),
    }
    b.AddComponents(components)
    w.Write(b.HtmlBytes())
}