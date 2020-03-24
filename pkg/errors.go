package pkg

import "errors"

var (
	// ErrAlreadyExists Returned when a resource already exists
	ErrAlreadyExists = errors.New("error: This resource already exists")

	// ErrNotFound Returned when a resource is not found
	ErrNotFound      = errors.New("error: Unable to find resource")

	// ErrDatabase Generic error message for database errors
	ErrDatabase      = errors.New("error: Something went wrong with the database")

	// ErrInvalidSlug Invalid data was sent as a request
	ErrInvalidSlug   = errors.New("error: Invalid json data")

	// ErrUnauthorized The requesting user is not authorized to access this resource
	ErrUnauthorized  = errors.New("error: Unauthorized")

	// ErrForbidden The requesting user is forbidden from accessing this resource
	ErrForbidden     = errors.New("error: You are forbidden from accessing this resource")
)