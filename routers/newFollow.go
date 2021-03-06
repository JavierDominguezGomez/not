package routers

import (
	"net/http"

	"github.com/JavierDominguezGomez/not/db"
	"github.com/JavierDominguezGomez/not/models"
)

/*NewFollow register a new relation between two users, */
func NewFollow(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the id parameter.", http.StatusBadRequest)
		return
	}

	var t models.Follow
	t.UserID = IDUser
	t.UserFollowedID = ID

	status, err := db.InsertFollow(t)
	if err != nil {
		http.Error(w, "Failed to save relation to database! "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "Relation couldn't be saved: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
