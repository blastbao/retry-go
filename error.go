package retry

import (
	"github.com/juju/errgo"
)

var (
	TimeoutError           = errgo.New("Operation aborted. Timeout occured")
	MaxRetriesReachedError = errgo.New("Operation aborted. Too many errors.")
)

// IsTimeout returns true if the cause of the given error is a TimeoutError.
func IsTimeout(err error) bool {
	return errgo.Cause(err) == TimeoutError
}

// IsMaxRetriesReached returns true if the cause of the given error is a MaxRetriesReachedError.
func IsMaxRetriesReached(err error) bool {
	return errgo.Cause(err) == MaxRetriesReachedError
}

type Err struct {
	Code uint32
	Msg string
}

func (e Err) Error() string {
	return ""
}



func Contains(a []uint32, x uint32) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
