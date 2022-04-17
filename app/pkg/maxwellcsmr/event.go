package maxwellcsmr

import "fmt"

type Event struct {
	Database string `json:"database"`
	Table    string `json:"table"`
	Type     string `json:"type"`
	TS       uint   `json:"ts"`
	XID      uint   `json:"xid"`
	XOffset  uint   `json:"xoffset"`
	Commit   bool   `json:"commit"`
	Data     Data   `json:"data"`
	Old      Data   `json:"old"`
}

type Data []byte

func (d *Data) UnmarshalJSON(in []byte) error {
	*d = in
	return nil
}

// TODO: Move to proper place
type Sample struct {
	ID        uint     `json:"id"`
	Biid      string   `json:"biid"`
	OfficeID  string   `json:"office_id"`
	Code      string   `json:"code"`
	Category  string   `json:"category"`
	Amount    Decimal  `json:"amount"`
	ValidFrom Date     `json:"valid_from"`
	ValidTo   Date     `json:"valid_to"`
	CreatedAt Time     `json:"created_at"`
	DeletedAt NullTime `json:"deleted_at"`
	Version   uint8    `json:"version"`
}

func CacheKey(xid uint) string {
	return fmt.Sprintf("%d-records", xid)
}
