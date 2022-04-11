package debeziumcsmr

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

type Decimal decimal.Decimal

// TODO: Enable to handle numbers with values with non-2 digits aftere the decimal point
const (
	digitsAfterDecimalPoint = -2
	baseNumBit              = 2
	bitSize64               = 64
)

func (d *Decimal) UnmarshalJSON(in []byte) error {
	// decimal is based64 encoded by Debezium connector
	// https://debezium.io/documentation/reference/stable/connectors/mysql.html#mysql-decimal-types
	decoded, err := base64.StdEncoding.DecodeString(string(in[1 : len(in)-1]))
	if err != nil {
		return err
	}

	// TODO: Investigate more efficient way
	// decimal to binary string with 8 digits (0 padding if the length is not enough)
	// Concatenate string, and convert it to decimal
	var valStr string
	for i := 0; i <= len(decoded)-1; i++ {
		valStr += fmt.Sprintf("%08b", decoded[i])
	}

	// Convert binary with string format to decimal number
	v, err := strconv.ParseInt(valStr, baseNumBit, bitSize64)
	if err != nil {
		return err
	}

	*d = Decimal(decimal.New(v, digitsAfterDecimalPoint))
	return nil
}

func (d Decimal) String() string {
	return decimal.Decimal(d).String()
}
