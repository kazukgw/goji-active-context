package activecontext

import (
	"encoding/json"
	"net/http"
)

func (ac *ActiveContext) Json(d interface{}) {
	json, err := json.Marshal(d)
	if err != nil {
		http.Error(ac.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	ac.Writer.Header().Set("Content-Type", "application/json")
	ac.Writer.Write(json)
}

func (ac *ActiveContext) ErrorJson(e interface{}) {
	ej := map[string]interface{}{
		"status":  "error",
		"message": e,
	}
	ac.Json(ej)
}

func (ac *ActiveContext) SuccessJson(d interface{}) {
	sj := map[string]interface{}{
		"status": "success",
		"data":   d,
	}
	ac.Json(sj)
}
