package exception

import (
	"go-fiber-postgres/model"
	"net/http"
)

var (
	OUTLET_NOT_FOUND                = NewError(http.StatusBadRequest, model.BAD_REQUEST, "OUTLET_NOT_FOUND")
	NO_CITIES_FOUND                 = NewError(http.StatusBadRequest, model.BAD_REQUEST, "NO_CITIES_FOUND")
	NO_DISTRICTS_FOUND              = NewError(http.StatusBadRequest, model.BAD_REQUEST, "NO_DISTRICTS_FOUND")
	NO_OUTLET_CATEGORIES_FOUND      = NewError(http.StatusBadRequest, model.BAD_REQUEST, "NO_OUTLET_CATEGORIES_FOUND")
	NO_OUTLET_AREA_CATEGORIES_FOUND = NewError(http.StatusBadRequest, model.BAD_REQUEST, "NO_OUTLET_AREA_CATEGORIES_FOUND")
	OWNER_IS_UNDERAGE               = NewError(http.StatusBadRequest, model.BAD_REQUEST, "OWNER_IS_UNDERAGE")
	//CONTENT_NOT_FOUND          = NewError(http.StatusBadRequest, model.BAD_REQUEST, "CONTENT_NOT_FOUND")
	BAD_REQUEST = NewError(http.StatusBadRequest, model.BAD_REQUEST, "BAD_REQUEST")
	//OUTLET_NOT_VALID           = NewError(http.StatusBadRequest, model.BAD_REQUEST, "OUTLET_NOT_VALID")
	CONFIRM_PASSWORD_NOT_MATCH = NewError(http.StatusBadRequest, model.BAD_REQUEST, "CONFIRM_PASSWORD_NOT_MATCH")
	START_END_DATE_NOT_VALID   = NewError(http.StatusBadRequest, model.BAD_REQUEST, "START_END_DATE_NOT_VALID")
)
