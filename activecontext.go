package activecontext

import (
	"html/template"
	"io"
	"net/http"

	"github.com/zenazn/goji/web"
)

type ActiveContext struct {
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
) *ActiveContext {
	env := c.Env["env"].(string)
	tmpl, ok := c.Env["template"].(Template)
	if !ok {
		tmpl = DefaultTemplate
	}

	sess, ok := c.Env["session"].(Session)
	if !ok {
		sess = DefaultSession
	}

	logger, ok := c.Env["logger"].(Logger)
	if !ok {
		logger = DefaultLogger
	}
	return &ActiveContext{
		Contect:  c,
		Writer:   w,
		Request:  r,
		Template: tmpl,
		Env:      env,
		Session:  sess,
		Logger:   logger,
	}
}

func (ac *ActiveContext) Head(code int) {
	ac.Writer.WriteHeader(code)
}

func (ac *ActiveContext) SendFile(path string) {
	http.ServeFile(ac.Writer, ac.Request, path)
}

func (ac *ActiveContext) Redirect(path string, status int) {
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

func (ac *ActiveContext) RedirectTo500Page() {
	http.Redirect(ac.Writer, ac.Request, Path500, http.StatusFound)
}

func (ac *ActiveContext) RedirectTo404Page() {
	http.Redirect(ac.Writer, ac.Request, Path404, http.StatusFound)
}

func (ac *ActiveContext) RedirectTo403Page() {
	http.Redirect(ac.Writer, ac.Request, Path403, http.StatusFound)
}
