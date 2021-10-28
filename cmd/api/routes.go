package main

// postgres 12345 5432
import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *blog) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckHandler)

	router.HandlerFunc(http.MethodPost, "/v1/blog", app.createBlogHandler)
	router.HandlerFunc(http.MethodGet, "/v1/blog/:id", app.showBlogHandler)

	return router
}
