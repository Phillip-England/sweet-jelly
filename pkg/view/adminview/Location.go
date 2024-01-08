package adminview

import (
	"cfasuite/pkg/database"
	"cfasuite/pkg/model/locationmod"
	"cfasuite/pkg/model/usermod"
	"cfasuite/pkg/mw"
	"cfasuite/pkg/util"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type LocationData struct {
    Title string
    Location locationmod.Model
    Users []usermod.Model
    RegisterUserErr string
    SessionToken string
}

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
    userRepo := usermod.NewUserRepo(db)
    err = userRepo.GetAllByLocationNumber(location.Number)
    if err != nil {
        http.Error(w, fmt.Sprintf("failed to load users by location: %s", err), http.StatusInternalServerError)
    }
    util.RenderTemplate(w, "./pkg/view/adminview/Location.html", LocationData{
        Title: "CFA Suite - Location",
        Location: *location,
        Users: userRepo.Users,
        RegisterUserErr: r.URL.Query().Get("RegisterUserErr"),
        SessionToken: os.Getenv("ADMIN_SESSION_TOKEN"),
    })
}