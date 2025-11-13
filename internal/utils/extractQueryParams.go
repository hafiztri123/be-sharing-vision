package utils

import (
	"net/http"
	"strconv"
)

type PaginationParams struct {
	Limit  int
	Offset int
}

func ExtractPathValue(r *http.Request) (*PaginationParams, []ValidationErrorPayload) {
	var validationErrors []ValidationErrorPayload
	var queryParams PaginationParams

	if v, err := strconv.Atoi(r.PathValue("limit")); err == nil {
		queryParams.Limit = v
	} else {
		validationErrors = append(validationErrors, ValidationErrorPayload{
			Key: "limit",
			Message: "Limit is missing",
		})
	}

	if v, err := strconv.Atoi(r.PathValue("offset")); err == nil {
		queryParams.Offset = v
	} else {
		validationErrors = append(validationErrors, ValidationErrorPayload{
			Key: "offset",
			Message: "Offset is missing",
		})
	}

	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	return &queryParams, nil
}
