package errors

import (
	"fmt"
	"io/fs"

	stderrors "errors"

	"github.com/pkg/errors"
)

func wrapErrorStd() error {
	// %w returns wrapError which implements Unwrap method
	return fmt.Errorf("something failed %w", errors.New("treasure"))
}

func wrapFSErorr() error {
	return fmt.Errorf("something failed %w", fs.ErrClosed)
}

func foo() error {
	// use WithMessage instead if stacktrace is not needed
	return errors.Wrap(errors.Wrap(errors.New("treasure"), "layer one"), "layer two")
}

func bar() error {
	return errors.Wrap(foo(), "call to foo failed")
}

func run() {
	// errors from std errors pkg
	e1 := wrapErrorStd()
	fmt.Println(e1.Error())
	// unwrap the outer layer
	fmt.Println(stderrors.Unwrap(e1))
	// nothing left to unwrap
	fmt.Println(stderrors.Unwrap(stderrors.Unwrap(e1)))

	fmt.Println(stderrors.Is(wrapFSErorr(), fs.ErrClosed))

	// errors from github.com/pkg/errors
	e2 := bar()
	// errors from this pkg implement fmt.Formatter
	// %v verb is also implemented, so Error() method need not be invoked for printing
	// %s, %v recursively prints the stack by unwrapping the errors
	fmt.Println(e2)

	// returns the topmost error that doesn't implement causer interface
	fmt.Println(errors.Cause(e2))

	// Prints stack trace of the err
	fmt.Printf("%+v\n", errors.Cause(e2))
}
