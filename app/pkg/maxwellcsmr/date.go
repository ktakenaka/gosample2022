package maxwellcsmr

import (
	"time"
	"unsafe"
)

const (
	dFormat = "\"2006-01-02\""
)

type Date time.Time

// UnmarshalJSON implements encoding.UnmarshalJSON interface
func (d *Date) UnmarshalJSON(in []byte) error {
	t, err := time.Parse(dFormat, *(*string)(unsafe.Pointer(&in)))
	if err != nil {
		return err
	}

	*d = Date(t)
	return nil
}
