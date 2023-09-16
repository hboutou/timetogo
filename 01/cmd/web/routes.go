package main

import (
	"net/http"

	"01/ui"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	// public routes
	router.Handler(http.MethodGet, "/", app.sessionManager.LoadAndSave(noSurf(app.authenticate(http.HandlerFunc(app.home)))))
	router.Handler(http.MethodGet, "/snippet/view/:id", app.sessionManager.LoadAndSave(noSurf(app.authenticate(http.HandlerFunc(app.snippetView)))))
	router.Handler(http.MethodGet, "/user/signup", app.sessionManager.LoadAndSave(noSurf(app.authenticate(http.HandlerFunc(app.userSignup)))))
	router.Handler(http.MethodPost, "/user/signup", app.sessionManager.LoadAndSave(noSurf(app.authenticate(http.HandlerFunc(app.userSignupPost)))))
	router.Handler(http.MethodGet, "/user/login", app.sessionManager.LoadAndSave(noSurf(app.authenticate(http.HandlerFunc(app.userLogin)))))
	router.Handler(http.MethodPost, "/user/login", app.sessionManager.LoadAndSave(noSurf(app.authenticate(http.HandlerFunc(app.userLoginPost)))))

	// auth protected routes
	router.Handler(http.MethodGet, "/snippet/create", app.sessionManager.LoadAndSave(noSurf(app.authenticate(app.requireAuthentication(http.HandlerFunc(app.snippetCreate))))))
	router.Handler(http.MethodPost, "/snippet/create", app.sessionManager.LoadAndSave(noSurf(app.authenticate(app.requireAuthentication(http.HandlerFunc(app.snippetCreatePost))))))
	router.Handler(http.MethodPost, "/user/logout", app.sessionManager.LoadAndSave(noSurf(app.authenticate(app.requireAuthentication(http.HandlerFunc(app.userLogoutPost))))))

	return app.recoverPanic(app.logRequest(secureHeaders(router)))
}
