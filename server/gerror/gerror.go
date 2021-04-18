package gerror

import "fmt"

// Wrap wraps an error, given a message and a set of parameters
func Wrap(err error, msg string, args ...interface{}) error {
	return fmt.Errorf("%s %w", fmt.Sprintf(msg, args...), err)
}
