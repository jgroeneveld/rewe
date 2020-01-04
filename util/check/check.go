// Package check provides function to panic on unrecoverable errors and on unexpected state.
// Basically deal with things that should never happen.
// Example:
//
//  import "rewe/util/check"
//
//  object, err := someCall()
//  check.Error(err)
//  check.Equal(len(object.Vertices), 2, "Number of Vertices")
//
// Output:
//   panic: IllegalState: Number of Vertices
//   Expected: 2
//   Actual: 4
package check

import "fmt"

const (
	msgIllegalState = "IllegalState"
	msgActual       = "  Actual"
	msgExpected     = "Expected"
)

// Error checks if there is an error and raises if needed.
// Used in places where we do not expect an error or have a situation where we can not recover from
func Error(err error) {
	if err != nil {
		panic(err)
	}
}

// Equal checks if two values are equal otherwise raises.
func Equal(actual interface{}, expected interface{}, msgf ...interface{}) {
	if expected != actual {
		msg := fmt.Sprintf("%s: %s\n%s: %#v\n%s: %#v", msgIllegalState, titleOrMsgf("Not Equal", msgf), msgExpected, expected, msgActual, actual)
		panic(msg)
	}
}

// True checks if a value is true
func True(result bool, msgf ...interface{}) {
	if !result {
		msg := msgIllegalState
		if len(msgf) > 0 {
			msg += ": " + msgfToString(msgf)
		}

		panic(msg)
	}
}

func titleOrMsgf(title string, msgf []interface{}) string {
	if len(msgf) > 0 {
		return msgfToString(msgf)
	}

	return title
}

func msgfToString(args []interface{}) string {
	if len(args) == 1 {
		return fmt.Sprintf("%s", args[0])
	}

	return fmt.Sprintf(args[0].(string), args[1:]...)
}
