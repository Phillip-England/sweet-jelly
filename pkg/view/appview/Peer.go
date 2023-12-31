package appview

import (
	"cfasuite/pkg/comp"
	"cfasuite/pkg/database"
	"cfasuite/pkg/model/usermod"
	"cfasuite/pkg/mw"
	"cfasuite/pkg/util"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

func Peer(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value(mw.DbKey).(*sql.DB)
	if !ok {
		http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
		return
	}
	user, ok := r.Context().Value(mw.AuthKey).(*usermod.Model)
	if !ok {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}
	peerIDStr := r.URL.Path[len("/app/peer/"):]
	peerID, err := strconv.Atoi(peerIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	peer := usermod.NewUserModel(db)
	peer.ID = peerID
	err = database.DbGetById(peer)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to load user by id: %s", err), http.StatusInternalServerError)
		return
	}
	if user.LocationNumber != peer.LocationNumber {
		http.Error(w, "do not have access to requested user", http.StatusForbidden)
		return
	}
	b := util.PageBuilder{
		Title: "CFA Suite - Peer Details",
	}
	components := []string{
		comp.Header("App Peer Page"),
		comp.TeamNav(),
		comp.UserDetails(peer),
		comp.UserPhoto(peer),
		comp.UserBio(peer), // Include the new UserBio component
	}
	b.AddComponents(components)
	w.Write(b.HtmlBytes())
}
