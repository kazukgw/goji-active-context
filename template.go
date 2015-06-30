package appctx

import (
	"html/template"
	"io"

	"github.com/yosssi/ace"
)

func (ac *AppContext) RenderTemplate(data interface{}, paths []string) error {
	tpl, err := ac.Template.Load(basePath, innerPath, a.Options)
	if err != nil {
		ac.ErrorLog(err)
		return err
	}

	if err := ac.Template.Execute(tpl, ac.Writer, data); err != nil {
		ac.ErrorLog(err)
		return err
	}
}

var TemplatePath500 = "views/error/500"
var TemplatePath404 = "views/error/404"
var TemplatePath403 = "views/error/403"

func (ac *AppContext) Render500() {
	ac.Template.Render(TemplatePath500, "", nil)
}

func (ac *AppContext) Render404() {
	ac.Template.Render(TemplatePath404, "", nil)
}

func (ac *AppContext) Render403() {
	ac.Template.Render(TemplatePath403, "", nil)
}

type AceRender struct {
	Options *ace.Options
}

func (a *AceRender) GetOptions(paths []string) {
	return a.Options
}

func (a *AceRender) Load(paths []string) (*template.Template, error) {
	return ace.Load(paths[0], paths[1], a.Options)
}

func (a *AceRender) Execute(t *template.Template, w io.Writer, data interface{}) {
	return t.Execute(w, data)
}
