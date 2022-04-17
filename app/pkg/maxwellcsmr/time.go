package maxwellcsmr

import (
	"bytes"
	"time"
	"unsafe"

	"github.com/volatiletech/null/v8"
)

const format = "\"2006-01-02 15:04:05\""

type Time time.Time
type NullTime null.Time

// UnmarshalJSON implements encoding.UnmarshalJSON interface
func (t *Time) UnmarshalJSON(in []byte) error {
	_t, err := time.Parse(format, *(*string)(unsafe.Pointer(&in)))
	if err != nil {
		return err
	}

	*t = Time(_t)
	return nil
}

// UnmarshalJSON implements encoding.UnmarshalJSON interface
func (t *NullTime) UnmarshalJSON(in []byte) error {
	if bytes.Equal(in, []byte("null")) {
		*t = NullTime(null.NewTime(time.Time{}, false))
		return nil
	}

	_t, err := time.Parse(format, *(*string)(unsafe.Pointer(&in)))
	if err != nil {
		return err
	}

	*t = NullTime(null.NewTime(_t, true))
	return nil
}
