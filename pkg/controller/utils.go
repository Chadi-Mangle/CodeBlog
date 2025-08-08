package controller

import (
	"net/http"

	"github.com/Chadi-Mangle/CodeBlog/pkg/logging"
	"github.com/Chadi-Mangle/CodeBlog/pkg/view"
)

type Action func(http.ResponseWriter, *http.Request)

type Middleware func(Action) Action

type MiddlewareProvider interface {
	Middleware() Middleware
}

type Controller struct {
	View       *view.View
	Middleware Middleware
}

func defaultMiddleware(a Action) Action {
	return func(w http.ResponseWriter, r *http.Request) {
		logging.Request.Printf("Test")
		a(w, r)
	}
}

func WarpMiddleware(m Middleware) Middleware {
	return func(a Action) Action {
		return defaultMiddleware(m(a))
	}
}

func getMiddleware(c *Controller) Middleware {
	if c.Middleware == nil {
		return defaultMiddleware
	}
	return WarpMiddleware(c.Middleware)
}

func NewController(v *view.View, m Middleware) *Controller {
	return &Controller{
		View:       v,
		Middleware: m,
	}
}

func (c *Controller) Index(dataFunc func(*http.Request) any) Action {
	return getMiddleware(c)(func(w http.ResponseWriter, r *http.Request) {
		data := dataFunc(r)
		if err := c.View.Index.Render(w, data); err != nil {
			logging.Error.Printf("Template error : %s", err)
		}
	})
}

func (c *Controller) Show(dataFunc func(*http.Request) any) Action {
	return getMiddleware(c)(func(w http.ResponseWriter, r *http.Request) {
		data := dataFunc(r)
		if err := c.View.Show.Render(w, data); err != nil {
			logging.Error.Printf("Template error : %s", err)
		}
	})
}

func (c *Controller) New(dataFunc func(*http.Request) any) Action {
	return getMiddleware(c)(func(w http.ResponseWriter, r *http.Request) {
		data := dataFunc(r)
		if err := c.View.New.Render(w, data); err != nil {
			logging.Error.Printf("Template error : %s", err)
		}
	})
}

func (c *Controller) Edit(dataFunc func(*http.Request) any) Action {
	return getMiddleware(c)(func(w http.ResponseWriter, r *http.Request) {
		data := dataFunc(r)
		if err := c.View.Edit.Render(w, data); err != nil {
			logging.Error.Printf("Template error : %s", err)
		}
	})
}
