// nolint:lll
package errors

// Client 809101xxx
var (
	RouteNotFound        = NewCustomError(809101001, StatusNotFound, "route not found")
	InvalidRequest       = NewCustomError(809101002, StatusBadRequest, "invalid request")
	MissAuthorization    = NewCustomError(809101003, StatusUnauthorized, "miss authorization")
	InvalidAuthorization = NewCustomError(809101004, StatusUnauthorized, "invalid authorization")
	AccessDenied         = NewCustomError(809101005, StatusForbidden, "you do not have permission to access this resource")
)

// Server 809102xxx
var (
	InternalServerPanic = NewCustomError(809102001, StatusInternalServerError, "internal server panic")
	InternalServerError = NewCustomError(809102002, StatusInternalServerError, "internal server error")
)

// Logic 809103xxx
var (
	EmailAlreadyExists      = NewCustomError(809103001, StatusConflict, "email already exists")
	NotFound                = NewCustomError(809103002, StatusNotFound, "not found")
	EmployeeAlreadyResigned = NewCustomError(809103003, StatusNotFound, "employee is resigned")
)
