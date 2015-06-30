package activecontext

import (
	"encoding/json"
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/zenazn/goji/web/middleware"
)

func (ac *ActiveContext) InfoLog(msg string) {
	rID := middleware.GetReqID(ac.Context)
	ac.Logger.InfoWithFields(msg, map[string]interface{}{
		"req_id": rID,
		"uri":    ac.Request.RequestURI,
	})
}

func (ac *ActiveContext) ErrorLog(e error) {
	rID := middleware.GetReqID(ac.Context)
	if ac.Env == EnvDevelopment {
		stack := make([]byte, 4096)
		read := runtime.Stack(stack, false)
		ac.Logger.ErrorWithFields(e, map[string]interface{}{
			"datetime":   time.Now().Format("2006-01-02 15:04:05"),
			"req_id":     rID,
			"uri":        ac.Request.RequestURI,
			"stacktrace": string(stack[:read]),
		})
	} else {
		ac.Logger.ErrorWithFields(e, map[string]interface{}{
			"datetime": time.Now().Format("2006-01-02 15:04:05"),
			"req_id":   rID,
			"uri":      ac.Request.RequestURI,
		})
	}
}

func (ac *ActiveContext) ParamsLog(params interface{}) {
	rid := middleware.GetReqID(ac.Context)
	uri := ac.Request.RequestURI
	ac.Logger.ParamsWithFields(params, map[string]interface{}{
		"datetime": time.Now().Format("2006-01-02 15:04:05"),
		"req_id":   rid,
		"uri":      uri,
	})
}

type LogrusLogger struct {
	*logrus.Logger
}

func (l *LogrusLogger) ErrorWithFields(e error, f map[string]interface{}) {
	f["msg"] = e.Error()
	logrusF := logrus.Fields(f)
	l.Logger.WithFields(logrusF).Error("error")
}

func (l *LogrusLogger) InfoWithFields(msg string, f map[string]interface{}) {
	logrusF := logrus.Fields(f)
	l.Logger.WithFields(logrusF).Info(msg)
}

func (l *LogrusLogger) ParamsWithFields(params interface{}, f map[string]interface{}) {
	var pstr string
	pb, err := json.Marshal(params)
	if err != nil {
		pstr = "json marshal error"
	} else {
		pstr = string(pb)
	}

	f["params"] = pstr
	l.Logger.WithFields(logrus.Fields(f)).Info("params_log")
}
