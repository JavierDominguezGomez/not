package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JavierDominguezGomez/not/db"
	"github.com/JavierDominguezGomez/not/models"
)

/*RecordMsg record a message in database. */
func RecordMsg(w http.ResponseWriter, r *http.Request) {
	var message models.Msg
	err := json.NewDecoder(r.Body).Decode(&message)

	registro := models.RecordMsg{
		UserID:   IDUser,
		Message:  message.Message,
		Datetime: time.Now(),
	}

	_, status, err := db.InsertMsg(registro)
	if err != nil {
		http.Error(w, "An error occurred while trying to insert the register. "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "The message could not be inserted.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
