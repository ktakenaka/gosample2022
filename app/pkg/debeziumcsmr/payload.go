package debeziumcsmr

/*
Following Debezium spec
https://debezium.io/documentation/reference/stable/connectors/mysql.html
*/

type Payload struct {
	After       interface{}
	Transaction interface{}
}
