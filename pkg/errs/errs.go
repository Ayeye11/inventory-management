package errs

import "errors"

var (
	//4xx
	ErrBadRequest error = errors.New("bad request")
	ErrNotFound   error = errors.New("not found")
)

func ErrFunc(err error) {
	switch err {
	//4xx
	case ErrNotFound:

	}
}

func CheckErr(err error) {

}
