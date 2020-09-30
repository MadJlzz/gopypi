package web
//
//import (
//	"github.com/MadJlzz/gopypi/internal/pkg/backend"
//	"github.com/MadJlzz/gopypi/internal/pkg/template"
//	"github.com/gorilla/mux"
//	"github.com/sirupsen/logrus"
//	"net/http"
//)
//
//type Controller struct {
//	localStorage *backend.LocalStorage
//	template *template.SimpleRepositoryTemplate
//}
//
//func New(ls *backend.LocalStorage, tmpl *template.SimpleRepositoryTemplate) *Controller {
//	return &Controller{
//		localStorage: ls,
//		template: tmpl,
//	}
//}
//
//func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
//	pkgs, err := c.localStorage.Load()
//	if err != nil {
//		logrus.Warnf("could not load packages. [%v]", err)
//		// Some fancy HTTP error code that is user friendly
//	}
//	if err = c.template.Execute(w, "index", pkgs); err != nil {
//		logrus.Errorf("could not execute template [index]. [%v]\n", err)
//		//Some fancy HTTP error code that is user friendly
//	}
//}
//
//func (c *Controller) Package(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	if _, found := vars["name"]; !found {
//		logrus.Errorln("routing variable [name] wasn't provided")
//		//Some fancy HTTP error code
//	}
//	// TODO: put in place a cache instead of loading everytime from the source.
//	pkgs, err := c.localStorage.Load()
//	if err != nil {
//		logrus.Warnf("could not load packages. [%v]", err)
//	}
//	pkg, found := pkgs[vars["name"]]
//	if !found {
//		logrus.Errorln("package [%s] is not available anymore...", vars["name"])
//		//Some fancy HTTP error code
//	}
//	//pkg.Files[0] = "C:/DefaultStorage/example-pkg/example-pkg-0.0.1.tar.gz"
//	if err = c.template.Execute(w, "package", pkg); err != nil {
//		logrus.Errorf("could not execute template [package]. [%v]\n", err)
//		//Some fancy HTTP error code that is user friendly
//	}
//}
