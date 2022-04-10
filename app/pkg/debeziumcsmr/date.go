package debeziumcsmr

import (
	"strconv"
	"time"
	"unsafe"
)

const (
	dFormat = "2006-01-02"
)

type Date time.Time

var (
	epochTime = func() time.Time {
		t, err := time.Parse(dFormat, "1970-01-01")
		if err != nil {
			panic(err)
		}
		return t
	}()
)

// UnmarshalJSON implements encoding.UnmarshalJSON interface
func (d *Date) UnmarshalJSON(in []byte) error {
	// Convert decimal letters to int
	// io.debezium.time.Date. Represents the number of days since the epoch.
	// https://debezium.io/documentation/reference/stable/connectors/mysql.html#mysql-temporal-types
	daysFromEpoch, err := strconv.Atoi(*(*string)(unsafe.Pointer(&in)))
	if err != nil {
		return err
	}
	*d = Date(epochTime.AddDate(0, 0, daysFromEpoch))
	return nil
}

// String implements Stringer interface
func (d Date) String() string {
	return time.Time(d).Format(dFormat)
}
