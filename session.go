package activecontext

import (
	"github.com/gorilla/sessions"
)

func (ac *ActiveContext) SaveSession() error {
	if err := ac.Session.Save(ac.Request, ac.Writer); err != nil {
		ac.ErrorLog(err)
		return err
	}
	return nil
}

type GorillaSession struct {
	*sessions.Session
}

func (s *GorillaSession) Get(key string) interface{} {
	return s.Session.Values[key]
}

func (s *GorillaSession) Set(key string, value interface{}) {
	s.Session.Values[key] = value
}
