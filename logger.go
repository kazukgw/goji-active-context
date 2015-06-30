package activecontext

func (ac *ActiveContext) InfoLog(msg string) {
	if ac.Logger == nil {
		return
	}
	rID := middleware.GetReqID(ac.Context)
	ac.Logger.InfoWithFields(msg, map[string]interface{}{
		"req_id": rID,
		"uri":    ac.Request.RequestURI,
	})
}

func (ac *ActiveContext) ErrorLog(e error) {
	if ac.Logger == nil {
		return
	}
	rID := middleware.GetReqID(ac.Context)
	if ac.Env == EnvDevelopment {
		stack := make([]byte, 4096)
		read := runtime.Stack(stack, false)
		ac.Logger.ErrorWithFields(e, map[string]interface{}{
			"req_id":     rID,
			"uri":        ac.Request.RequestURI,
			"stacktrace": string(stack[:read]),
		})
	} else {
		ac.Logger.ErrorWithFields(e, map[string]interface{}{
			"req_id": rID,
			"uri":    ac.Request.RequestURI,
		})
	}
}

func (ac *ActiveContext) ParamsLog(params interface{}) {
	if ac.Logger == nil {
		return
	}
	rid := middleware.GetReqID(ac.Context)
	uri := ac.Request.RequestURI
	log.ParamsLog(rid, uri, params)
}

type LogrusLogger struct {
	*logrus.Logger
}

func (l *LogrusLogger) ErrorWithFields(e error, f map[string]interface{}) {
	f["msg"] = e.Error()
	logrusF := logrus.Fields(f)
	Logger.WithFields(logrusF).Error("error")
}

func (l *LogrusLogger) InfoWithFields(msg string, f map[string]interface{}) {
	logrusF := logrus.Fields(f)
	Logger.WithFields(logrusF).Info(msg)
}

func (l *LogrusLogger) ParamsLog(params interface{}) {
	var pstr string
	pb, err := json.Marshal(params)
	if err != nil {
		pstr = "json marshal error"
	} else {
		pstr = string(pb)
	}

	l.Logger.WithFields(logrus.Fields{
		"datetime": time.Now().Format("2006-01-02 15:04:05"),
		"req_id":   rid,
		"uri":      uri,
		"params":   pstr,
	}).Info("params_log")
}
