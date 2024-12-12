package status

import "net/http"

const (
	BadRequest       = http.StatusBadRequest          // 400
	MethodNotAllowed = http.StatusMethodNotAllowed    // 405
	NotFound         = http.StatusNotFound            // 404
	OK               = http.StatusOK                  // 200
	Created          = http.StatusCreated             // 201
	InternalError    = http.StatusInternalServerError // 500
)
