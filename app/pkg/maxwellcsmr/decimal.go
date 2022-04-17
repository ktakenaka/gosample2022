package maxwellcsmr

import (
	"unsafe"

	"github.com/shopspring/decimal"
)

// "amount":12.34

type Decimal decimal.Decimal

func (d *Decimal) UnmarshalJSON(in []byte) error {
	_d, err := decimal.NewFromString(*(*string)(unsafe.Pointer(&in)))
	if err != nil {
		return err
	}
	*d = Decimal(_d)
	return nil
}
