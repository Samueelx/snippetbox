package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

/**serverError writes an error message and the stack trace to the errorLog, then sends a generic 500 internal server error response to the user.*/
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

/**clientError sends a specific status code and corresponding description to the user.*/
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

/**notFound  is a wrapper around clientError which sends a 404 Not Found response to the user*/
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
