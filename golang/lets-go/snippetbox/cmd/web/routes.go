package main

import (
	"net/http"

	"blue.study.golang/ui"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.FS(ui.Files))
	
	router.Handler(http.MethodGet, "/static/*filepath",  fileServer)

	dynamicMiddleware := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	router.Handler(http.MethodGet, "/", dynamicMiddleware.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/snippet/view/:id", dynamicMiddleware.ThenFunc(app.snippetView))
	router.Handler(http.MethodGet, "/user/signup", dynamicMiddleware.ThenFunc(app.userSignUp))
	router.Handler(http.MethodPost, "/user/signup", dynamicMiddleware.ThenFunc(app.userSignUpPost))
	router.Handler(http.MethodGet, "/user/login", dynamicMiddleware.ThenFunc(app.userLogin))
	router.Handler(http.MethodPost, "/user/login", dynamicMiddleware.ThenFunc(app.userLoginPost))
	
	protectedMiddleware := dynamicMiddleware.Append(app.requireAuthentication)

	router.Handler(http.MethodGet, "/snippet/create", protectedMiddleware.ThenFunc(app.snippetCreate))
	router.Handler(http.MethodPost, "/snippet/create", protectedMiddleware.ThenFunc(app.snippetCreatePost))
	router.Handler(http.MethodPost, "/user/logout", protectedMiddleware.ThenFunc(app.userLogoutPost))

	standardMiddlewares := alice.New(app.recoverPanic, app.logRequest, secureHeader)

	// with out using alice package
	// return app.recoverPanic(app.logRequest(secureHeader(mux)))
	return standardMiddlewares.Then(router)
}