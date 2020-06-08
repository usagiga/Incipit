package errors

import (
	"fmt"
	"golang.org/x/xerrors"
)

// DistinctError represents error which has error code and its detail
// It is compatible with `golang.org/x/xerrors` and Go 1.13 `errors`
type DistinctError struct {
	PrimaryCode   PrimaryErrorCode
	SecondaryCode SecondaryErrorCode
	ErrMsg        string
	Next          error
	Detail        interface{}
	frame         xerrors.Frame
}

func NewDistinctError(
	msg string,
	pCode PrimaryErrorCode,
	sCode SecondaryErrorCode,
	detail interface{},
) *DistinctError {
	return &DistinctError{
		PrimaryCode:   pCode,
		SecondaryCode: sCode,
		ErrMsg:        msg,
		Next:          nil,
		Detail:        detail,
		frame:         xerrors.Caller(1),
	}
}

func (e DistinctError) Wrap(next error) error {
	e.Next = next
	e.frame = xerrors.Caller(1)
	return &e
}

func (e *DistinctError) Error() string {
	errMsg := fmt.Sprintf("ERR%s-%s: %s", e.PrimaryCode.String(), e.SecondaryCode.String(), e.ErrMsg)
	return errMsg
}

func (e *DistinctError) Unwrap() (err error) {
	return e.Next
}

func (e *DistinctError) Is(err error) (equal bool) {
	// Cast to DistinctError
	var dErr *DistinctError
	ok := xerrors.As(err, &dErr)
	if !ok {
		return false
	}

	// Compare
	matchPrim := dErr.PrimaryCode == e.PrimaryCode
	matchSecond := dErr.SecondaryCode == e.SecondaryCode

	return matchPrim && matchSecond
}

func (e *DistinctError) Format(s fmt.State, c rune) { xerrors.FormatError(e, s, c) }

func (e *DistinctError) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.Error())
	e.frame.Format(p)
	return e.Next
}
