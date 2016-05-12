package web

import (
	"github.com/Cloakaac/cloak/controllers"
	"github.com/Cloakaac/cloak/util"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	admin = iota
	logged
	guest
	pass
)

func newRouter() *httprouter.Router {
	router := httprouter.New()
	registerRoutes(router)
	return router
}

func registerRoutes(router *httprouter.Router) {
	base := &controllers.BaseController{}
	router.GET("/", route(base.Home, base, pass))
	router.GET("/account/create", route(base.Register, base, guest))
	router.POST("/account/create", route(base.CreateAccount, base, guest))
	router.GET("/account/login", route(base.Login, base, guest))
	router.POST("/account/login", route(base.SignIn, base, guest))
	router.GET("/account/manage", route(base.AccountManage, base, logged))
	router.POST("/account/logout", route(base.AccountLogout, base, logged))
	router.GET("/character/view/:name", route(base.CharacterView, base, pass))
	router.GET("/character/signature/:name", route(base.SignatureView, base, pass))
}

func route(controller func(w http.ResponseWriter, req *http.Request, ps httprouter.Params), base *controllers.BaseController, mode ...int) func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		if req.Method == http.MethodPost {
			err := req.ParseForm()
			if err != nil {
				util.HandleError("Something occured on req.ParseForm()", err)
				return
			}
		}
		session, err := util.GetSession(req, "cloaka")
		if err != nil {
			util.HandleError("Something occured while getting a session", err)
			return
		}
		for i := range mode {
			if mode[i] == admin {

			}
			if mode[i] == logged && session.GetInt("logged") == 0 {
				http.Redirect(w, req, "/account/login", http.StatusMovedPermanently)
				return
			}
			if mode[i] == guest && session.GetInt("logged") == 1 {
				http.Redirect(w, req, "/account/manage", http.StatusMovedPermanently)
				return
			}
		}
		base.Session = session
		controller(w, req, ps)
	}
}
