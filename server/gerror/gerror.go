package gerror

import "fmt"

func Wrap(err error, msg string, args ...interface{}) error {
	return fmt.Errorf("%s %w", fmt.Sprintf(msg, args...), err)
}
