package historydate

import (
	"bytes"
	"time"
	"unsafe"
)

type Date time.Time

const format = "2006-01-02"

var maxDate = func() Date {
	t, err := time.Parse(format, "9999-12-31")
	if err != nil {
		panic(err)
	}
	return Date(t)
}()

func ParseDate(s string) (Date, error) {
	t, err := time.ParseInLocation(format, s, time.UTC)
	if err != nil {
		return Date{}, err
	}
	return Date(t), nil
}

func (d Date) String() string {
	return time.Time(d).Format(format)
}

func (d Date) ToTime() time.Time {
	return time.Time(d)
}

func (d Date) IsMax() bool {
	return d.ToTime().Equal(maxDate.ToTime())
}

func (d Date) MarshalJSON() ([]byte, error) {
	if d.IsMax() {
		return []byte(`null`), nil
	}

	return []byte(`"` + d.String() + `"`), nil
}

func (d *Date) UnmarshalJSON(in []byte) error {
	if bytes.Equal(in, []byte(`null`)) {
		*d = maxDate
		return nil
	}

	if len(in) > 1 {
		in = in[1 : len(in)-1]
	}
	_d, err := ParseDate(*(*string)(unsafe.Pointer(&in)))
	if err != nil {
		return err
	}
	*d = _d
	return nil
}
