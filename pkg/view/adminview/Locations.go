package adminview

import (
	"cfasuite/pkg/database"
	"cfasuite/pkg/model/locationmod"
	"cfasuite/pkg/mw"
	"cfasuite/pkg/util"
	"database/sql"
	"fmt"
	"net/http"
)

type LocationsData struct {
    Title string
    LocationFormErr string
    Locations []locationmod.Model
}

func Locations(w http.ResponseWriter, r *http.Request) {
    db, ok := r.Context().Value(mw.DbKey).(*sql.DB)
    if !ok {
        http.Error(w, "failed to get database connection", http.StatusInternalServerError)
        return
    }
    locationRepo := &locationmod.Repo{
        DB: db,
    }
    if err := database.DbGetAll(locationRepo); err != nil {
        http.Error(w, fmt.Sprintf("Failed to get locations: %s", err), http.StatusInternalServerError)
        return
    }
    util.RenderTemplate(w, "./pkg/view/adminview/Locations.html", LocationsData{
        Title: "CFA Suite - Locations",
        LocationFormErr: r.URL.Query().Get("LocationFormErr"),
        Locations: locationRepo.Locations,
    })
    // b := util.PageBuilder{
    //     Title: "CFA Suite - User Details",
    // }
    // components := []string{
    //     comp.Header("Admin Locations Page"),
    //     comp.AdminNav(),
    //     comp.CreateLocationForm(locationFormErr),
    //     comp.LocationList(locationRepo.Locations),
    // }
    // b.AddComponents(components)
    // w.Write(b.HtmlBytes())
}