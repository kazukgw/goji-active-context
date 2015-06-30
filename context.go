package appctx

import (
	"html/template"
	"io"
	"net/http"

	"github.com/zenazn/goji/web"
)

type AppContext struct {
	Context web.C
	Writer  http.ResponseWriter
	Request *http.Request
	Env     string
	Template
	Session
	Logger
}

type Template interface {
	GetOptions(...string) map[string]interface{}
	Load(...string) (*template.Template, error)
	Execute(*template.Template, io.Writer, interface{}) error
}

type Session interface {
	Save(http.ResponseWriter, *http.Request) error
}

type Logger interface {
	ParamsLog(interface{})
	ErrorWithFields(error, map[string]interface{})
	InfoWithFields(string, map[string]interface{})
}

var EnvDevelopment = "development"

func New(
	c web.C,
	w http.ResponseWriter,
	r *http.Request,
) *AppContext {
	env := c.Env["env"].(string)
	tmpl, _ := c.Env["template"].(Template)
	sess, _ := c.Env["session"].(Session)
	logger, _ := c.Env["logger"].(Logger)
	return &AppContext{
		Contect:  c,
		Writer:   w,
		Request:  r,
		Template: tmpl,
		Env:      env,
		Session:  sess,
		Logger:   logger,
	}
}

func (ac *AppContext) Head(code int) {
	ac.Writer.WriteHeader(code)
}

func (ac *AppContext) SendFile(path string) {
	http.ServeFile(ac.Writer, ac.Request, path)
}

func (ac *AppContext) Redirect(path string, status int) {
	http.Redirect(
		ac.Writer,
		ac.Request,
		path,
		status,
	)
}

var Path500 = "/500"
var Path404 = "/404"
var Path403 = "/403"

func (ac *AppContext) RedirectTo500Page() {
	http.Redirect(ac.Writer, ac.Request, Path500, http.StatusFound)
}

func (ac *AppContext) RedirectTo404Page() {
	http.Redirect(ac.Writer, ac.Request, Path404, http.StatusFound)
}

func (ac *AppContext) RedirectTo403Page() {
	http.Redirect(ac.Writer, ac.Request, Path403, http.StatusFound)
}
