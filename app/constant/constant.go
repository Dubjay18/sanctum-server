package constant

import "net/http"

const (
	INTERNAL_SERVER_ERROR = http.StatusInternalServerError
	NOT_FOUND             = http.StatusNotFound
	OK                    = http.StatusOK
	NO_CONTENT            = http.StatusNoContent
	BAD_REQUEST           = http.StatusBadRequest
	UNAUTHORIZED          = http.StatusUnauthorized
	FORBIDDEN             = http.StatusForbidden
)

func GetStatusMessage(status int) string {
	switch status {
	case INTERNAL_SERVER_ERROR:
		return "Internal Server Error"
	case NOT_FOUND:
		return "Not Found"
	case OK:
		return "Success"
	case NO_CONTENT:
		return "No Content"
	case BAD_REQUEST:
		return "Bad Request"
	case UNAUTHORIZED:
		return "Unauthorized"
	case FORBIDDEN:
		return "Forbidden"
	default:
		return "Unknown Status"
	}
}
