package constants

type ResponseStatus int
type Headers int
type General int

// Constant Api
const (
	Created ResponseStatus = iota + 1
	Success
	DataNotFound
	UnknownError
	InvalidRequest
	Unauthorized
	BadRequest
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{"CREATED", "SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED", "BAD_REQUEST"}[r-1]
}

func (r ResponseStatus) GetResponseKey() string {
	return [...]string{"x-created", "x-success", "x-data-not-found", "x-unknown-error", "x-invalid-request", "x-unauthorized", "x-bad-request"}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"Created", "Success", "Data Not Found", "Unknown Error", "Invalid Request", "Unauthorized", "Bad Request"}[r-1]
}
