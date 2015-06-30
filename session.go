package appctx

func (ac *AppContext) SaveSession(sess map[string]interface{}) error {
	for k, v := range sess {
		ac.Session.Values[k] = v
	}

	if serr := ac.Session.Save(ac.Request, ac.Writer); serr != nil {
		ac.ErrorLog(serr)
		ac.RedirectTo500Page()
		return serr
	}
	return nil
}
