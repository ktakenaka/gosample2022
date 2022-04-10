package debeziumcsmr

/*
Following Debezium spec
https://debezium.io/documentation/reference/stable/connectors/mysql.html
*/
type Transaction struct {
	ID                  string `json:"id"`
	TotalOrder          uint   `json:"total_order"`
	DataCollectionOrder uint   `json:"data_collection_order"`
}

// TODO: Move to proper place. It's not natural to define this in pkg.
type SamplePayload struct {
	Payload struct {
		Before      Sample      `json:"before"`
		After       Sample      `json:"after"`
		Transaction Transaction `json:"transaction"`
	} `json:"payload"`
}
type Sample struct {
	ID        uint     `json:"id"`
	Biid      string   `json:"biid"`
	OfficeId  string   `json:"office_id"`
	Code      string   `json:"code"`
	Category  string   `json:"category"`
	Amount    Decimal  `json:"amount"`
	ValidFrom Date     `json:"valid_from"`
	ValidTo   Date     `json:"valid_to"`
	CreatedAt Time     `json:"created_at"`
	DeletedAt NullTime `json:"deleted_at"`
}

type TransactionPayload struct {
	Payload struct {
		Status          string `json:"status"`
		ID              string `json:"id"`
		EventCount      uint   `json:"event_count"`
		DataCollections []struct {
			DataCollection string `json:"data_collection"`
			EventCount     uint   `json:"event_count"`
		} `json:"data_collections"`
	} `json:"payload"`
}

/*
  "payload": {
    "status": "END",
    "id": "f64b158e-b7fb-11ec-89da-0242ac150004:49",
    "event_count": 1,
    "data_collections": [
      {
        "data_collection": "gosample2022_development.samples",
        "event_count": 1
      }
    ]
  }
*/
