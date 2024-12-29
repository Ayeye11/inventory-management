package errs

import "errors"

var (
	//4xx
	ErrBadRequest error = errors.New("bad request")
	ErrNotFound   error = errors.New("not found")
)

func CheckErr(err error) {

}
