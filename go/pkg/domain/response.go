package domain

// BadRequest 400エラーの処理
func BadRequest(err error) Error {
	return newError(StatusBadRequest, err)
}

// Unauthorized 401エラーの処理
func Unauthorized(err error) Error {
	return newError(StatusUnauthorized, err)
}

// MethodNotAllowed 405エラーの処理
func MethodNotAllowed(err error) Error {
	return newError(StatusMethodNotAllowed, err)
}

// InternalServerError 500エラーの処理
func InternalServerError(err error) Error {
	return newError(StatusInternalServerError, err)
}

func newError(code int, err error) Error {
	if err != nil {
		return &httpError{
			Code:  code,
			error: err,
		}
	}
	return nil
}

// Error statusCode付きのerror wrapper
type Error interface {
	Error() string
	GetStatusCode() int
}

func (he *httpError) Error() string {
	return he.error.Error()
}

func (he *httpError) GetStatusCode() int {
	return he.Code
}

type httpError struct {
	Code int
	error
}

const (
	StatusDefault             = 500
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusMethodNotAllowed    = 405
	StatusInternalServerError = 500
)
