package utils

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func ParseURLQuery(r *http.Request) (err error, period, tz, t1, t2 string) {

	period = r.URL.Query().Get("period")
	tz = r.URL.Query().Get("tz")
	t1 = r.URL.Query().Get("t1")
	t2 = r.URL.Query().Get("t2")
	errStr := ""
	if len(period) == 0 {
		errStr = errStr + "Period is empty. "
	}
	if len(tz) == 0 {
		errStr = errStr + "Time zone is empty. "
	}
	if len(t1) == 0 {
		errStr = errStr + "T1 is empty. "
	}
	if len(t2) == 0 {
		errStr = errStr + "T2 is empty. "
	}
	if len(errStr) == 0 {
		return nil, period, tz, t1, t2
	} else {
		return errors.New(errStr), period, tz, t1, t2
	}
}

func HttpRespond(w http.ResponseWriter, header int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(data)
	w.Write(b)

	w.WriteHeader(header)
}

func HttpError(w http.ResponseWriter, status int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	resp := map[string]interface{}{
		"desc":   err.Error(),
		"status": "error",
	}
	b, _ := json.Marshal(resp)
	w.Write(b)

}

func NewUUIDV4() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return ""
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}
