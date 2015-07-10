package activecontext

import (
	"encoding/json"
	"net/http"
)

func (ac *ActiveContext) Json(d interface{}, status int) {
	json, err := json.Marshal(d)
	if err != nil {
		http.Error(ac.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	ac.Writer.Header().Set("Content-Type", "application/json")
	ac.Writer.WriteHeader(status)
	ac.Writer.Write(json)
}
