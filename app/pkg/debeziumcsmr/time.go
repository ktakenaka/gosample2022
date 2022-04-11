package debeziumcsmr

import (
	"bytes"
	"time"
	"unsafe"

	"github.com/volatiletech/null/v8"
)

type Time time.Time

// UnmarshalJSON implements encoding.UnmarshalJSON interface
func (t *Time) UnmarshalJSON(in []byte) error {
	// Remove double quotes
	in = in[1 : len(in)-1]
	_t, err := time.Parse(time.RFC3339, *(*string)(unsafe.Pointer(&in)))
	if err != nil {
		return err
	}
	*t = Time(_t)
	return nil
}

// String implements Stringer interface
func (t Time) String() string {
	return time.Time(t).Format(time.RFC3339)
}

type NullTime null.Time

// UnmarshalJSON implements encoding.UnmarshalJSON interface
func (t *NullTime) UnmarshalJSON(in []byte) error {
	if bytes.Equal(in, []byte("null")) {
		*t = NullTime(null.NewTime(time.Time{}, false))
		return nil
	}
	if err := t.Time.UnmarshalJSON(in); err != nil {
		return err
	}
	t.Valid = true
	return nil
}

// String implements Stringer interface
func (t NullTime) String() string {
	if !t.Valid {
		return "null"
	}
	return t.Time.Format(time.RFC3339)
}
