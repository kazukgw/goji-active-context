package activecontext

import (
	"encoding/json"

	"github.com/goji/param"
)

func (ac *ActiveContext) GetQueryParam(key string) string {
	return ac.Request.URL.Query().Get(key)
}

func (ac *ActiveContext) ParseParams(params interface{}) error {
	if rerr := ac.Request.ParseForm(); rerr != nil {
		ac.ErrorLog(rerr)
		ac.RedirectTo500Page()
		return rerr
	}

	if perr := param.Parse(ac.Request.PostForm, params); perr != nil {
		ac.ErrorLog(perr)
		ac.RedirectTo500Page()
		return perr
	}
	ac.ParamsLog(params)
	return nil
}

func (ac *ActiveContext) ParseJsonParams(params interface{}) error {
	if derr := json.NewDecoder(ac.Request.Body).Decode(params); derr != nil {
		ac.ErrorLog(derr)
		ac.ErrorJson(derr.Error())
		return derr
	}
	ac.ParamsLog(params)
	return nil
}
