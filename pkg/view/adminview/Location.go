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
	"strconv"
)

func Location(w http.ResponseWriter, r *http.Request) {
    db, ok := r.Context().Value(mw.DbKey).(*sql.DB)
    if !ok {
        http.Error(w, "failed to get database connection", http.StatusInternalServerError)
        return
    }
    locationID := r.URL.Path[len("/admin/location/"):]
    locationIDInt, err := strconv.Atoi(locationID)
    if err != nil {
        http.Error(w, "Invalid location ID", http.StatusBadRequest)
        return
    }

    location := &locationmod.Model{
        DB: db,
        ID: locationIDInt,
    }

    err = database.DbGetById(location)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get location details: %s", err), http.StatusInternalServerError)
        return
    }

    b := util.PageBuilder{
        Title: "CFA Suite - Location Details",
    }

    components := []string{
        comp.AdminNav(),
        comp.LocationDetails(location),
    }
    b.AddComponents(components)
    w.Write(b.HtmlBytes())
}