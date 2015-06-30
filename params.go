package appctx

func (ac *AppContext) GetQueryParam(key string) string {
	return ac.Request.URL.Query().Get(key)
}

func (ac *AppContext) ParseParams(params interface{}) error {
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

func (ac *AppContext) ParseJsonParams(params interface{}) error {
	if derr := json.NewDecoder(ac.Request.Body).Decode(params); derr != nil {
		ac.ErrorLog(derr)
		ac.ErrorJson(derr.Error())
		return derr
	}
	ac.ParamsLog(params)
	return nil
}
