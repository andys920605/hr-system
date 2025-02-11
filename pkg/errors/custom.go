// nolint:lll
package errors

// Client 809101xxx
var (
	RouteNotFound        = NewCustomError(809101001, StatusNotFound, "route not found")
	InvalidRequest       = NewCustomError(809101002, StatusBadRequest, "invalid request")
	MissAuthorization    = NewCustomError(809101003, StatusUnauthorized, "miss authorization")
	InvalidAuthorization = NewCustomError(809101004, StatusUnauthorized, "invalid authorization")
	AccessDenied         = NewCustomError(809101005, StatusForbidden, "you do not have permission to access this resource")
	ScheduleTimeInvalid  = NewCustomError(809101006, StatusBadRequest, "the time difference between end time and start time must be at least 1 hour")
)

// Server 809102xxx
var (
	InternalServerPanic = NewCustomError(809102001, StatusInternalServerError, "internal server panic")
	InternalServerError = NewCustomError(809102002, StatusInternalServerError, "internal server error")
)

// Logic 809103xxx
var (
	MaterialNotApproved   = NewCustomError(809103001, StatusConflict, "material has not been approved yet")
	ResourceAlreadyExists = NewCustomError(809103002, StatusConflict, "resource already exists and cannot be created again")
	ResourceNotFound      = NewCustomError(809103003, StatusNotFound, "resource to edit does not exist")
	StatusTransitionError = NewCustomError(809103004, StatusConflict, "status transition is not allowed")
	InsufficientCredit    = NewCustomError(809103005, StatusBadRequest, "insufficient credit")
)
