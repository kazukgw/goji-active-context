package activecontext

import (
	"html/template"
	"io"

	"github.com/yosssi/ace"
)

func (ac *ActiveContext) RenderTemplate(data interface{}, paths ...string) error {
	tpl, err := ac.Template.Load(paths...)
	if err != nil {
		ac.ErrorLog(err)
		return err
	}

	if err := ac.Template.Execute(tpl, ac.Writer, data); err != nil {
		ac.ErrorLog(err)
		return err
	}

	return nil
}

var TemplatePath500 = "views/error/500"
var TemplatePath404 = "views/error/404"
var TemplatePath403 = "views/error/403"

func (ac *ActiveContext) Render500(data interface{}) {
	ac.RenderTemplate(data, TemplatePath500)
}

func (ac *ActiveContext) Render404(data interface{}) {
	ac.RenderTemplate(data, TemplatePath404)
}

func (ac *ActiveContext) Render403(data interface{}) {
	ac.RenderTemplate(data, TemplatePath403)
}

type AceRenderer struct {
	Options ace.Options
}

func (a *AceRenderer) Load(paths ...string) (*template.Template, error) {
	var base, inner string
	if len(paths) > 0 {
		base = paths[0]
	} else if len(paths) > 1 {
		inner = paths[1]
	}
	return ace.Load(base, inner, &a.Options)
}

func (a *AceRenderer) Execute(t *template.Template, w io.Writer, data interface{}) error {
	return t.Execute(w, data)
}
