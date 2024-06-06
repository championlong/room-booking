package constants

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Error struct {
	Code    int
	Message string
}

var (
	ErrBadRequest          = &Error{Code: 400000, Message: "Bad Request"}
	ErrInternalServerError = &Error{Code: 500000, Message: "Internal Server Error"}
)

func (e *Error) Error() string {
	return fmt.Sprintf("code:%d,message:%s", e.Code, e.Message)
}

func (e *Error) StatusCode() int {
	return e.Code / 1000
}

func ErrorToGrpcError(err error) error {
	if err == nil {
		return nil
	}
	var e *Error
	if errors.As(err, &e) {
		return status.Error(codes.Code(e.Code), e.Message)
	} else {
		return status.Error(codes.Code(ErrBadRequest.Code), ErrBadRequest.Message)
	}
}

func ErrorToRestError(err error) *Error {
	if err == nil {
		return nil
	}
	var e *Error
	if errors.As(err, &e) {
		return e
	} else {
		return ErrInternalServerError
	}

}

func GrpcErrorToError(err error) *Error {
	st, ok := status.FromError(err)
	if ok && st.Code() >= 400000 {
		return &Error{int(st.Code()), st.Message()}
	} else {
		return ErrInternalServerError
	}
}
