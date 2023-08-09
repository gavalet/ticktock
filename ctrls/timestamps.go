package ctrls

import (
	"cmd/ticktock/models"
	"cmd/ticktock/utils"
	u "cmd/ticktock/utils"
	"net/http"

	l "cmd/ticktock/utils/logger"
)

func GetTimestamps(w http.ResponseWriter, r *http.Request) {
	reqID := utils.NewUUIDV4()
	log := l.NewLogger(reqID)

	err, pr, tz, t1, t2 := utils.ParseURLQuery(r)
	log.Debug("Filters - ", "pr:", pr, " tz: ", tz, " t1:", t1, " t2:", t2)

	if err != nil {
		log.Error("Got error: ", err)
		u.HttpError(w, http.StatusBadRequest, err)
	}

	data, status, err := models.GetTimestamps(reqID, pr, tz, t1, t2)
	if err != nil {
		log.Error("Got error: ", err)
		u.HttpError(w, status, err)
	}
	u.HttpRespond(w, status, data)
}
