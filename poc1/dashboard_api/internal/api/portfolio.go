package api

import (
	"net/http"

	"github.com/dxps/opa_showcase/poc1/dashboard_svc/internal/app"
)

func (api *API) getSubjectPortfolio(w http.ResponseWriter, r *http.Request) {

	// Getting the subject's external ID provided as URL param.
	subjectID, err := api.readUUIDParam(r)
	if err != nil {
		api.logger.Print("getSubjectAttributesHandler > readUUIDParam (id) error: ", err)
		api.badRequestResponse(w, r, app.ErrSubjectIDInvalid)
		return
	}
	api.logger.Println("[api > getSubjectPortfolio] subjectID =", subjectID)

	// TODO: to be cont'd
}
