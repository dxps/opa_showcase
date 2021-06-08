package api

import (
	"net/http"
)

func (api *API) logError(r *http.Request, err error) {
	api.logger.Printf("Error '%s' encountered on request '%s %s'", err, r.Method, r.URL.String())
}

func (api *API) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}

	err := api.writeJSON(w, status, env, nil)
	if err != nil {
		api.logError(r, err)
		w.WriteHeader(500)
	}
}

func (api *API) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	api.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	api.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (api *API) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	api.errorResponse(w, r, http.StatusNotFound, message)
}

func (api *API) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	api.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (api *API) invalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	api.errorResponse(w, r, http.StatusUnauthorized, message)
}
