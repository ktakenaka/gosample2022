# Debezium Setup
Register a connector
```
curl --location --request POST 'http://localhost:8083/connectors' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "gosample2022-connector",
    "config": {
        "connector.class": "io.debezium.connector.mysql.MySqlConnector",
        "tasks.max": "1",
        "database.hostname": "mysql",
        "database.port": "3306",
        "database.user": "root",
        "database.password": "root",
        "database.include.list": "gosample2022_development",
        "table.include.list": "gosample2022_development.samples",
        "database.server.name": "gosample2022_dbserver",
        "database.history.kafka.bootstrap.servers": "kafka:29092",
        "database.history.kafka.topic": "schema-changes.gosample2022",
        "provide.transaction.metadata": true
    }
}'
```

Confirm the configuration
```
curl -H "Accept:application/json" http://localhost:8083/connectors/gosample2022-connector | jq
```

Monitor the topic
```
kcat -b localhost:9092 -t dbserver1.gosample2022_development.samples
```

Delete connector
```
curl -i -X DELETE localhost:8083/connectors/gosample2022-connector
```
